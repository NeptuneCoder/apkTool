#### 项目描述
该工程用于处理apk资源合并

#### 项目结构说明
.
├── analysis  //处理包的流程
│   └── analysis.go 
├── apkTool  //程序的入口
│   ├── apkTool.go
│   └── apkTool_test.go
├── cmad       //开始执行命令
│   └── cmd.go
├── customConfig
│   └── createConfig.go
├── dom4g       //解析xml的工具类 https://github.com/donnie4w/dom4g.git
│   ├── dom4g.go
│   └── dom4g_test.go
├── env         //读取环境变量
│   └── Env.go
├── gameInfo    //读取游戏的相关配置信息
│   └── game.go
├── jar2smali   //将jar 转换成smali文件
│   └── Jar2Smali.go
├── merge
│   ├── merge.go
│   └── merge_test.go
├── model       //各种数据结构，主要构件读取的配置信息
│   └── model.go
├── parse       //解析配置信息，后期改为数据库
│   ├── parseJson.go
│   └── parseJson_test.go
├── replace     //替换文本内容
│   └── replaceContent.go
├── rjar        //有的模块可能通过R.*.* 的方式读取的资源文件，为了解决找不到对应资源id的问题，配置相关的包名，生成对应包名下的R文件
│   └── RJar.go
└── utils       //工具类
    ├── FilePath.go
    └── Utils.go
