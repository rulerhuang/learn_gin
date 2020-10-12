# learn_gin
+ 项目地址 https://github.com/EDDYCJY/go-gin-example
+ 文档地址 https://eddycjy.com/tags/gin/

# 依赖库
+ com库：https://github.com/unknwon/com
+ validation库：https://github.com/astaxie/beego/tree/master/validation

# tag说明
+ tag v1.0：初始化项目，搭建gin框架运行环境

# 亮点
+ 隐藏数据库连接：models/models.go中将数据库连接设为private，隐藏数据库连接，对外统一暴露的是models.XX方法而不是原始的数据库连接