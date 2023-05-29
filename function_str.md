# 字符串生成模式

## string
生成字符串

语法
```
string|{type?},{minlength?},{maxlength?}
```

参数
- **type** `string` 字符串类型 默认[a-zA-Z0-9]
  * upper 大写
  * letter 小写
  * ansic 大小写字母数字以及一些特殊符号

- **minlength** `int` 最小长度
- **maxlength** `int` 最大长度


示例
```
string // 随机长度的字符
string|number // 随机长度的数字
string|letter,10  // 长度为10的小写字母
string|letter,10,20  // 长度为10-20的小写字母
```

## boolean
生成bool值

语法
```
boolean|{value?}...
```

参数
- **value** `bool`

示例
```
boolean  // 随机boolean值
boolean|true  // true
boolean|true,false,false,false  // 随机从[true,false,false,false]中取值
```


## integer
生成数字整形

语法
```
integer|{min?},{max?}
```

参数
- min `int` 最小值
- max `int` 最大值

示例
```
integer // 随机一个整数
integer|10000 // 10000
integer|10000,20000  // 13886
```

## float
生成浮点数，支持保留几位小数点

语法
```
float|{min?},{max?},{fixed?}
```

参数
- min `float` 最小值
- max `float` 最大值
- fixed `int` 小数点保留几位

示例
```
float // 随机一个浮点数
float|102.01 // 102.01
float|10000,20000,4  // 13886.1021
```

## regexp
使用正则表达式生成数据

语法
```
regexp|{value}
```

示例
```
regexp|\d{10} // 2102301024
```

## oneof
多选一

语法
```
oneof|{value?}...
```

示例
```
oneof|男,女 // 女
oneof|1,"a b",a,b // b
```

## autoincrement
自增 仅在jsonschema，struct等模式下可用

语法
```
autoincrement|{begin?},{step?}
```

- **begin** `int` 起始值 默认 1
- **step** `int` 步长 默认 1

示例
```
autoincrement // 1,2,3,4,5...
autoincrement|100 // 100,101,102,...
autoincrement|100,2 // 100, 102,104,106...
```

## numberpattern
使用#替代数字

语法
```
numberpattern|{pattern}
```

示例
```
numberpattern|###-######-#  // 231-200312-1
```

## word, phrase, sentence, paragraph,title
单词，短语，句子，段落,标题   支持国际化

语法
```
word(lang?)|{minlength?),{maxlength?}
phrase(lang?)|{minlength?},{maxlength?}
sentence(lang?)|{minlength?},{maxlength?}
paragraph(lang?)|{minlength?},{maxlength?}
title(lang?)|{minlength?},{maxlength?}
```

参数
- lang `string` 语言
    - en (default)
    - zh
- minlength 最小长度
- maxlength 最大长度

示例
```
word
word|3  // abc
word(zh)|3  // 我爱你
word|3,5 // abcd
sentence(zh) // 这是一个中文句子。
title // Zcpk Wxtxnsa Lorn Gzfls Utegn Mrp
```


## name, firstname, lastname
```语法
name(lang?)
```

示例
```
name // Robert Robinson
name(zh) // 张爱军
firstname(zh) // 张
lastname // Robinson
```

## gender
性别 支持国际化 如果需要其他性别 请使用`oneof`

```语法
gender(lang?)
```

示例
```
gender // female
gender(zh) // 女
```

## phone
电话号码 除国内外其他基本上手机哈固定号码规则一致 phone(zh) 生成的是手机号 如果需要固话或其他格式 包括带国际编号的 请使用`numberpattern`

```语法
phone(lang?)
```

示例
```
phone // 202-328-2516
phone(zh) // 18910000000
```

## idcard
- 中国:身份证号
- 美国:社会统一保险号（默认）

语法
```
idcard(lang?)
```

示例
```
idcard // 203-23-1000
idcard(zh) // 610103199609225013
```

## ipv4, ipv6
```示例
ipv4 // 127.0.0.1
ipv6 // 2001:0db8:3c4d:0015:0000:0000:1a2f:1a2b
```

## uuid
示例
```
uuid // 9b271dc8-abb9-19b0-f5f4-43225ff7968c
```

## domain
域名

示例
```
domain // example.com.cn
```

## url
网址

```
url // https://aaa.org/somepath/radomx
```

## email
```
email // 183xas@aiwna.com
```

## date
日期 `format`采用通用的**java日期格式**

```格式
date|{format?}
```

示例
```
date // 2002-10-23
date|YYYY年MM月dd日 // 2020年01月20日
date|"YYYY MMMM/dd" // 2020 January/01
```

## time
格式
```
date|{format?}
```

示例
```
time // 12:00:00
time|HH:mm:ss  // 12:00:00 
```

## datetime
格式
```
datetime|{format?}
```

 - **format** 默认格式 RFC3339
示例
```
datetime // 2006-01-02T15:04:05Z07:00
datetime|"YYYY年MM月dd日 HH:mm" // 2020年01月20日 12:00
```

## timestamp
示例
```
timestamp // 1684853928
```

## now
当前时间 支持按format格式话当前时间 是`datetime`的别名 只是不随机

格式
```
now|{format?}
```

示例
```
now // 2006-01-02T15:04:05Z07:00
now|"YYYY年MM月dd日 HH:mm:ss"  // 2022年05月22日 13:00:32
```

## color
格式
```
color|{type?}
```

type
- rgb
- rgba
- hsl
- hex (default)

示例
```
color // #002211
color|rgb // rgb(255,0,123)
```


## httpcode
http状态码

示例
```
httpcode // 200
```

## httpmethod
http 请求方法

示例
```
httpmethod // DELETE
```

## imagedata
base64格式的图片数据

格式
```
imagedata|{width?},{height?}
```

- width 默认128
- height 默认128

示例
```
imagedata // data:image/png;base64,xx29ahjsh28wallxasdasdxxxx...
imagedata|200  // 宽高等于200的图片
imagedata|200,100 // 宽200 高100的图片
```

## imageurl
图片url

格式
```
imageurl|{width?},{height?}
```

- width 默认128
- height 默认128


示例
```
imageurl // https://dummyimage.com/128/128.png
imageurl|200  // https://dummyimage.com/200/200.png
imageurl|200,100  // https://dummyimage.com/200/100.png
```

## city
城市 支持国际话

格式
```
city(lang?)
```

示例
```
city // San Francisco
city(zh) // 西安市
```

## provinceorstate
省/邦/州

格式
```
provinceorstate(lang?)
```

示例
```
provinceorstate // California
provinceorstate(zh) // 陕西省
```


## provinceorstate&city
省/邦/州 市

格式
```
provinceorstate&city(lang?)
```

示例
```
provinceorstate&city // Ogden, Utah
provinceorstate&city(zh) // 陕西省 西安市
```

## street
街道

格式
```
street(lang?)
```

示例
```
street // 1950 ST Ciwbv Fort
street(zh) // 因看大街吴云 商场64-4号
```

## address
详细地址

格式
```
address(lang?)
```

示例
```
address // 1950 ST Ciwbv Fort, Laredo, Texas 71300
address(zh) // 河南省 郑州市 还用场集 小区1栋0单元975号 282563
```

## zipcode
邮编

格式
```
zipcode(lang?)
```

示例
```
zipcode // 71300
zipcode(zh) // 727000
```

## longitude, latitude
经度, 维度

示例
```
longitude // 116.397128
latitude // 39.916527
```


## longitude&latitude
经纬度

示例
```
longitude&latitude // 116.397128, 39.916527
```

## 