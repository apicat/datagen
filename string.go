package datagen

import (
	"fmt"
	"strings"
	"unicode"
)

// String 生成字符串
// mode
func String(model string, n ...int) string {
	x := randCount(n...)
	if x == -1 {
		x = 10
	}
	switch model {
	case "number":
		return Regexp(fmt.Sprintf(`\d{%d}`, x))
	case "upper":
		return Regexp(fmt.Sprintf(`[A-Z]{%d}`, x))
	case "letter":
		return Regexp(fmt.Sprintf(`[a-z]{%d}`, x))
	}
	return Regexp(fmt.Sprintf(`[a-zA-Z0-9_\-@\*]{%d}`, x))
}

type TypographyObject struct {
	o *i18nText
}

func Typography(local ...string) *TypographyObject {
	return &TypographyObject{
		o: GetLocale(local...).Text,
	}
}

func (t *TypographyObject) Word(n ...int) string {
	return t.o.RandomWord(n...)
}

func (t *TypographyObject) Phrase(n ...int) string {
	return t.o.RandomPhrase(n...)
}

func (t *TypographyObject) Sentence(n ...int) string {
	return t.o.RandomSentence(n...)
}

func (t *TypographyObject) Paragraph(n ...int) string {
	return t.o.RandomParagraph(n...)
}

func (t *TypographyObject) Title(n ...int) string {
	return t.o.RandomTitle(n...)
}

func (t *TypographyObject) Markdown() string {
	var buf strings.Builder
	buf.WriteString("# ")
	buf.WriteString(t.Title(4, 10))
	buf.WriteByte(0x0a)
	buf.WriteByte(0x0a)
	buf.WriteString(t.Paragraph(4, 10))
	buf.WriteByte(0x0a)
	buf.WriteByte(0x0a)
	buf.WriteString("> ")
	buf.WriteString(t.Paragraph(4, 10))
	buf.WriteByte(0x0a)
	buf.WriteByte(0x0a)
	listn := randInt(2, 5)
	for i := 0; i < listn; i++ {
		buf.WriteString("- ")
		buf.WriteString(t.Phrase(2, 5))
		buf.WriteByte(0x0a)
	}
	return buf.String()
}

type i18nText struct {
	Chars         string `json:"chars"`
	WordSplit     string `json:"wordSplit"`
	SentenceSplit string `json:"sentenceSplit"`
	SentenceEnd   string `json:"sentenceEnd"`
	//
	_charRuns []rune
}

func (t *i18nText) getn(n ...int) int {
	x := randCount(n...)
	if x == -1 {
		// 默认6-12
		return 6 + grand.Intn(6)
	}
	return x
}

// RandomWord 随机生成词语
func (t *i18nText) RandomWord(n ...int) string {
	var b strings.Builder
	for i := 0; i < t.getn(n...); i++ {
		b.WriteRune(pick(t._charRuns))
	}
	return b.String()
}

// RandomPhrase 随机生成一段短语
// 短语是由1个多单词组成的文本
func (t *i18nText) RandomPhrase(n ...int) string {
	var b strings.Builder
	for i := 0; i < t.getn(n...); i++ {
		if i > 0 {
			b.WriteString(t.WordSplit)
		}
		b.WriteString(t.RandomWord(2, 10))
	}
	return b.String()
}

// RandomSentence 随机生成一个语句
// 首位大写 并且有句号
func (t *i18nText) RandomSentence(n ...int) string {
	var b strings.Builder
	b.WriteRune(unicode.ToTitle(pick(t._charRuns)))
	b.WriteString(t.RandomPhrase(n...))
	b.WriteString(t.SentenceEnd)
	return b.String()
}

// RandomParagraph 随机生成一个段落
// 多个短语组成，第一个短语的首位大写 每个短语由句号或逗号分割，最后一定以句号结尾
func (t *i18nText) RandomParagraph(n ...int) string {
	var b strings.Builder
	b.WriteRune(unicode.ToTitle(pick(t._charRuns)))
	for i := 0; i < t.getn(n...); i++ {
		if i > 0 {
			// 随机给句号或逗号
			b.WriteString(t.SentenceSplit)
		}
		b.WriteString(t.RandomPhrase(2, 10))
	}
	b.WriteString(t.SentenceEnd)
	return b.String()
}

// RandomTitle 随机生成一个标题
// 这个貌似只有英文才有区别
func (t *i18nText) RandomTitle(n ...int) string {
	var b strings.Builder
	for i := 0; i < t.getn(n...); i++ {
		if i > 0 {
			b.WriteString(t.WordSplit)
		}
		b.WriteRune(unicode.ToTitle(pick(t._charRuns)))
		b.WriteString(t.RandomWord(1, 10))
	}
	return b.String()
}
