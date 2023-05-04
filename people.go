package datagen

import (
	"fmt"
)

type PeopleObject struct {
	o *i18nPeople
}

func People(local ...string) *PeopleObject {
	return &PeopleObject{
		o: GetLocale(local...).People,
	}
}

func (p *PeopleObject) Name() string {
	return p.o.Name.RandomString()
}

func (p *PeopleObject) FirstName() string {
	return p.o.Name.RandomFirstName()
}

func (p *PeopleObject) LastName() string {
	return p.o.Name.RandomLastName()
}

func (p *PeopleObject) Gender() string {
	return pick(p.o.Gender)
}

// Phone 生成电话号码
// 现在基本都是手机了各国基本也都是手机和座机号码基本规则一致
func (p *PeopleObject) Phone() string {
	return NumberPattern(p.o.Phone)
}

// IDCard 公民唯一号码
// 中国:身份证号
// 美国:社会统一保险号
func (p *PeopleObject) IDCard() string {
	return NumberPattern(p.o.IDCard)
}

type i18nPeople struct {
	Name   i18nName `json:"name"`
	Phone  string   `json:"phone"`
	IDCard string   `json:"idNumber"`
	Gender []string `json:"gender"`
}

type i18nName struct {
	First []string `json:"first"`
	Last  []string `json:"last"`
	Split string   `json:"split"`
}

func (n *i18nName) RandomFirstName() string {
	return pick(n.First)
}

func (n *i18nName) RandomLastName() string {
	return pick(n.Last)
}

func (n *i18nName) RandomString() string {
	return fmt.Sprintf("%s%s%s", n.RandomFirstName(), n.Split, n.RandomLastName())
}
