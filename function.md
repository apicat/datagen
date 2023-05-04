
# Functions

### Core
- [String](#string)
- [Boolean](#boolean)
- [Integer](#integer)
- [Float](#float)
- [OneOf](#oneOf)  多选一
- [Regexp](#regexp)
- [NumberStringPattern](#numberStringPattern) eg: `(###)-#-##-##` `(851)-4-98-20`

### Plus
- [Typography](#typography) i18n
    - [Word](#word) 单词
    - [Phrase](#phrase) 短语
    - [Sentence](#sentence) 句子
    - [Paragraph](#paragraph) 段落
    - [Title](#title) 标题
- [People](#people) i18n
    - [FirstName](#firstName)
    - [LastName](#lastName)
    - [Name](#name)
    - [IDCard](#idcard) 身份证/ssn
    - [Gender](#gender) 性别
    - [Phone](#phone) 电话/手机
    - [Email](#email)
- [Address](#address) i18n
    - [Address](#address) 详细地址包含省市区
    - [ProvinceOrState](#provinceOrState) 省(州/加盟国)
    - [City](#city) 城市
    - [ZipCode](#zipcode) 邮编
    - [Street](#street) 街道
    - [Longitude](#longitude) 经度
    - [Latitude](#latitude) 维度
    - [LongitudeAndLatitude](#longitudeAndLatitude) 经纬度
- [Date](#date)
    - [Date](#date)
    - [DateTime](#datetime)
    - [Time](#time)
    - [Timestamp](#timestamp)
    - [Now](#now)
- [Internet](#internet)
    - [IPv4](#ipv4)
    - [IPv6](#ipv6)
    - [UUID](#uuid)
    - [URL](#url)
    - [Domain](#domain)
- [Draw](#draw)
    - [ImageURL](#imageurl)
    - [ImageData](#imagedata)
    - [Color](#color)





# Core

## String
`String(mode,minLength,maxLength)`  
**mode**
- default (ansic|number)
- ansic
- letter
- upper
- number
```go
String() // "wjE292@AX"
String("number") // "2013818123"
String("number",3) // "201"
String("numner",3,10) // "20100304"
String("upper",4) // "HPUJ"
```

## Boolean
`Boolean(arg ...bool)`
```go
Boolean() // true or flase
Boolean(true) // true
Boolean(false) // false
Boolean(true,true,false) // random[ture,true,false]
```

## Integer
`Integer(min,max)`
```go
Integer() // 1923841
Integer(4) // 4
Integer(1000,10000) // 9828
```

## Float
`Float(min,max,fixed)`
```go
Float() // 19.123412
Float(10.2) // 10.2
Float(1,2,2) // 1.42
Float(1,2,3) // 1.423
```

## OneOf
`OneOf(v ...any)`
```go
OneOf(1,2,3,4) // 3
OneOf("abc",true,2000) // 2000
OneOf("China","USA","Russia") // "China"
```

## Regexp
`Regexp(regexp)`
```go
Regexp("\d{10}") // 2312342141
```

## NumberPattern
`NumberPattern(pattern)`
```go
NumberPattern("(###)-###-####") // (013)-2341-2938
```


# Typography
`Typography(lang...)`

- [Word](#word)
- [Phrase](#phrase) 
- [Sentence](#sentence) 
- [Paragraph](#paragraph) 
- [Title](#title) 

```go
Typography() // default en
Typography("zh") // zh
Typography("en","zh") // oneOf "en","zh"
Typography("kr") // Typography("en")  not support kr return default en
```

## Word
`Word(minLength,maxLength)`
```go
t:=Typography()
t.Word() // "abcd"
t.Word(1) // "a"
t.Word(1,3) // "ab"

t:=Typography("zh")
t.Word() // "随五破下先"
```

## Phrase
`Phrase(minLength,maxLength)`
```go
t:=Typography()
t.Phrase() // "opa iwabuwna kaso"
t.Phrase(2) // "opa iwabuwna"
t.Phrase(1,10) // "opa iwabuwna kaso wfga"
```

## Sentence
`Sentence(minLength,maxLength)`
```go
t:=Typography()
t.Sentence() // "Pasdin xkao ikwnwre."
t:=Typography("zh")
t.Sentence() // "问候话我怕是你父亲成功后。"
```

## Paragraph
`Paragraph(minLength,maxLength)`
```go
t:=Typography()
t.Paragraph() // "Pasdin xkao ikwnwre, wiahu aiw angyubgad, khguig ja xoaqhd."
t:=Typography("zh")
t.Paragraph() // "问候话我怕是你，阿斯旺，父亲成功后。"
```

## Title
`Title(minLength,maxLength)`
```go
Typography().Title() // "Hyhwi Tuao Opa Kuahnng"
```


# People
`People(lang...)`
- [FirstName](#firstName)
- [LastName](#lastName)
- [Name](#name)
- [IDCard](#idcard) 身份证/ssn
- [Gender](#gender) 性别
- [Phone](#phone) 电话/手机
- [Email](#email)

## Name
## FirstName
## LastName

```go
p:=People()
p.Name() // "James Martinez"
P:=People("zh")
p.Name() // "王翠花"
p.FirstName() // "王"
p.LastName() // "翠花"
```



## IDCard
`IDCard()`
```go
// SSN
People().IDCard() // "427617281"
// 身份证号
People("zh").IDCard() // "610819284719228391"
```

## Gender
`Gender()`
```go
People().Gender() // "female"
People("zh").Gender() // "男"
```

## Phone
`Phone(pattn)`
```go
People().Phone() // "3023202027"
People("zh").Phone() // "18966842819"
People().Phone("(###)###-####") // "(213)231-1234"
```

## Email
`Email()`
```go
People().Email() // "q2xa323@demia.com"
```

# Position
`Position(loc...)`
- [Address](#address) 详细地址包含省市区
- [ProvinceOrState](#provinceOrState) 省(州/加盟国)
- [City](#city) 城市
- [ZipCode](#zipcode) 邮编
- [Street](#street) 街道
- [Longitude](#longitude) 经度
- [Latitude](#latitude) 维度
- [LongitudeAndLatitude](#longitudeAndLatitude) 经纬度

## Address
`Address()`
```go
Position().Address() // ""
Position("zh").Address() // ""
```

## ProvinceOrState
```go
Position().ProvinceOrState() // ""
Position("zh").ProvinceOrState() // ""
```

## City
```go
Position().City() // ""
Position("zh").City() // ""
```

## ZipCode
```go
Position().ZipCode() // ""
Position("zh").ZipCode() // ""
```

## Street
```go
Position().Street() // ""
Position("zh").Street() // ""
```
## Longitude
## Latitude
## LongitudeAndLatitude
```go
Position().Longitude()
Position().Latitude()
Position().LongitudeAndLatitude()
```

# Date&Time
format support java data formt
- [Date](#date)
- [DateTime](#datetime)
- [Time](#time)
- [Timestamp](#timestamp)
- [Now](#now)

## Date
`Date(format)`
```go
Date() //
Date("YYYY-MM-dd") //
```

## Time
`Date(format)`
```go
Date() //
Date("HH") //
```

## DateTime
`Date(format)`
```go
Date() //
Date("HH") //
```

## Timestamp
`Timestamp()`
```go
Timestamp() // 1678943232
```
## Now
`Now(format)`
```go
Now()
Now("xxx")
```

# Internet
- IPv4
- IPv6
- UUID
- URL
- Domain

```go
IPv4()
IPv6()
UUID()
URL()
Domain()
```

# Draw
- [ImageURL](#imageurl)
- [ImageData](#imagedata)
- [Color](#color)

## ImageURL
## ImageData
## Color
