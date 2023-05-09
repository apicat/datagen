# datagen

模拟数据生成 可以用作虚假数据生成，使用场景如mock，或者生成数据库测试数据等场景


## 特性

1. 国际化 目前仅支持英文和中文
2. 支持通过正则表达式生成数据
3. 支持jsonschema生成数据


## 常用functions

[已支持函数](function.md)


## 从字符串表达生成数据

### 字符串格式
`funcName(loc...)|args...`

* funcName 函数名
* loc 语言信息(非必须) 比如en,zh 默认en
* args 参数(非必须) args默认使用`,`隔开支持多参数，如果参数包含`,`请使用双引号或单引号包裹，强制指定未整体字符串


```
eg:

regexp|\d{4}
numberpattern|###-####
float|10,20
word(zh)
```



### jsonschema 生成数据

默认使用jsonschema内不的format,maxitems等属性生成，如果存在`x-datagen`字段 则会使用字符串表达式生成

> `x-datagen`字段 支持自定义如 `x-apicat-mock`

```

{
    "type":"object",
    "properties":{
        "children":{
            "type":"array",
            "items":{
                "type":"string",
                "x-datagen":"uuid"
            }
        }
    }
}


JSONSchemaGen(jsonschemaString, &GenOption{DatagenKey: "x-datagen"})

```

### struct 生成数据
```
type T struct {
    values map[string]string
    uid    string `datagen:"uuid"`
    info   struct {
        name    string `datagen:"name"`
        age     int    `datagen:"integer|10,40"`
        address string `datagen:"address"`
    }
}
var testt T
StructGen(testt, &GenOption{DatagenKey: "datagen"})
```


### 添加自己专用的函数

```
RegisterFunction("customgen", func(p datagen.Param)any{
    return "nihao"+p.Args.At(0)
})


CallFunction("customgen|a")
```

