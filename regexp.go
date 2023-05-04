package datagen

import (
	"bytes"
	"regexp/syntax"
	"sync"
)

// Regexp 通过正则表达式生成数据
func Regexp(regstr string) string {
	var buf bytes.Buffer
	v, ok := regGen.regexpCache.Load(regstr)
	if !ok {
		nodeTree, err := syntax.Parse(regstr, syntax.Perl)
		if err != nil {
			panic(err)
		}
		v = nodeTree.Simplify()
		regGen.regexpCache.Store(regstr, v)
	}
	s := v.(*syntax.Regexp)
	regGen.generate(&buf, s)
	return buf.String()
}

type regexpGen struct {
	regexpCache *sync.Map
	// 最大长度
	max int
}

var regGen = &regexpGen{
	regexpCache: &sync.Map{},
	max:         32,
}

func (r *regexpGen) randRepeatSubReg(buf *bytes.Buffer, subs []*syntax.Regexp, min, max int) {
	if max > r.max {
		max = r.max
	}
	n := randInt(min, max)
	for i := 0; i < n; i++ {
		for _, v := range subs {
			r.generate(buf, v)
		}
	}
}

func (r *regexpGen) generate(buf *bytes.Buffer, reg *syntax.Regexp) {
	switch reg.Op {
	case syntax.OpLiteral:
		buf.WriteString(string(reg.Rune))
	case syntax.OpCharClass:
		n := len(reg.Rune)
		if n == 0 || n%2 != 0 {
			return
		}
		i := grand.Intn(n / 2)
		buf.WriteRune(randRune(reg.Rune[i*2], reg.Rune[i*2+1]))
	case syntax.OpAnyCharNotNL:
		buf.WriteRune(randRune(0x20, 0x7f))
	case syntax.OpAnyChar:
		r := randRune(0x20, 0x7f+1)
		if r == 0x7f {
			r = 0x0a // 替换delete 为 换行
		}
		buf.WriteRune(r)
	case syntax.OpEndLine:
		buf.WriteRune(rune(0x0a))
	case syntax.OpWordBoundary:
		buf.WriteRune(rune(0x20))
	case syntax.OpStar:
		r.randRepeatSubReg(buf, reg.Sub, 0, -1)
	case syntax.OpPlus:
		r.randRepeatSubReg(buf, reg.Sub, 1, -1)
	case syntax.OpQuest:
		r.randRepeatSubReg(buf, reg.Sub, 0, 1)
	case syntax.OpRepeat:
		r.randRepeatSubReg(buf, reg.Sub, reg.Min, reg.Max)
	case syntax.OpConcat, syntax.OpCapture:
		for _, v := range reg.Sub {
			r.generate(buf, v)
		}
	case syntax.OpAlternate:
		r.generate(buf, pick(reg.Sub))
	}
}
