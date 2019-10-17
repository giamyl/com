# bcom

写代码中用到的一些公共的函数和方法

将平时经常用到的东西记录下来， 避免用到的时候再找浪费时间

<!-- TOC -->

- [自定义函数包](#自定义函数包)
    - [benc](#benc)
    - [bpath](#bpath)
    - [blog](#blog)
    - [bstr](#bstr)
- [常用的第三方包](#常用的第三方包)

<!-- /TOC -->

## 自定义函数包

暂时还简单， 根据实际使用慢慢完善

### benc

用于处理加密相关的操作，包含：

+ 密码加盐加密

+ 字符串转 md5 和 sha256

+ aes 加密和解密

### bpath

用来处理路径

+ 获取项目根目录

### blog

用来处理日志， 暂时做的还很基础， 需要启动项目时配合 shell 的输出重定向写入到文件

### bstr

用于处理一些字符串相关的操作

+ 生成随机字符串

+ 生成 UUID

+ 隐藏字符串中的部分内容

## 常用的第三方包

+ [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)

+ [gin-gonic/gin](https://github.com/gin-gonic/gin) -
[文档](https://gin-gonic.com/zh-cn/docs/)

+ [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) -
[文档（英文）](https://github.com/go-sql-driver/mysql/wiki)

+ [sqlx](https://github.com/jmoiron/sqlx)

+ [squirrel](https://github.com/Masterminds/squirrel)

+ [jinzhu/gorm](https://github.com/jinzhu/gorm) -
[文档](https://gorm.io/zh_CN/docs/)

+ [json-iterator](https://github.com/json-iterator/go)

+ [ethereum/go-ethereum](https://github.com/ethereum/go-ethereum) -
[官方文档（英文）](https://geth.ethereum.org/docs/) -
[中文教程](https://goethereumbook.org/zh/)

+ [gofrs/uuid](https://github.com/gofrs/uuid)

+ [mozillazg/go-pinyin](https://github.com/mozillazg/go-pinyin)

+ [robfig/cron](https://github.com/robfig/cron) -
[文档](https://godoc.org/github.com/robfig/cron#hdr-Usage)
