# Mercury

#### 简介

消息推送平台🔥推送下发【邮件】【短信】【微信服务号】【微信小程序】【企业微信】【钉钉】等消息类型。

#### 项目描述

核心功能：统一的接口发送各种类型消息，对消息生命周期全链路追踪。

意义：只要公司内部有发送消息的需求，都应该要有类似austin的项目。消息推送平台对各类消息进行统一发送处理，这有利于对功能的收拢，以及提高业务需求开发的效率。

#### 项目启动
1. 需要搭建mysql/rabbitmq/redis服务
1. 导入数据库文件:`Mercury/sql/Mercury.sql`并修改`send_account`表中的`config`配置
2. 修改对应的mysql/rabbitmq/redis配置
- rpc服务配置:`Mercury/app/Mercury-web/rpc/etc/Mercury.yaml`
- job服务配置:`Mercury/app/Mercury-job/etc/Mercury-job.yaml`
3.
> 手动启动
- 启动api服务:`Mercury/app/Mercury-web/api/Mercury.go`
- 启动rpc服务:`Mercury/app/Mercury-web/rpc/Mercury.go`
- 启动job服务:`Mercury/app/Mercury-job/rpc/Mercury-job.go`
> docker一键启动
- 执行 deployment-shell.sh
4. 如需要测试去重服务则修改`message_template`表中的`deduplication_config`字段
```
{"deduplication_10":{"num":1,"time":300},"deduplication_20":{"num":5}}
```
5. 使用示例
> 邮件消息
```
curl --location --request POST 'http://localhost:8888/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "code": "send",
    "messageParam": {
        "receiver": "test@qq.com",
        "variables": {
            "title": "测试操作",
            "content": "Hello <b>Bob</b> and <i>Cora</i>!"
        }
    },
    "messageTemplateId": 2
}'
```

> 微信公众号消息
```
curl --location --request POST 'http://localhost:8888/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "code": "send",
    "messageParam": {
        "receiver": "openId",
        "variables": {
            "map": {
                "name":"张三12333"
            },
            "url": "https://www.baidu.com/"
        }
    },
    "messageTemplateId": 4
}'

//参数带颜色的
curl --location --request POST 'http://localhost:8888/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "code": "send",
    "messageParam": {
        "receiver": "openId",
        "variables": {
            "map": {
                "name":"张三12333|#0000FF"
            },
            "url": "https://www.baidu.com/"
        }
    },
    "messageTemplateId": 4
}'
```

> 钉钉自定义机器人
```
//艾特某些手机号
curl --location --request POST 'http://localhost:8888/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "code": "send",
    "messageParam": {
        "receiver": "13588888888,13588888887",
        "variables": {
            "content": "测试\n换行"
        }
    },
    "messageTemplateId": 5
}'

//艾特全部人
curl --location --request POST 'http://localhost:8888/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "code": "send",
    "messageParam": {
        "receiver": "@all",
        "variables": {
            "content": "测试\n换行"
        }
    },
    "messageTemplateId": 5
}'
```




#### 目录说明

```
├── app  
│   ├── Mercury-admin        消息平台管理端  
│   ├── Mercury-common       项目公用的一些结构体,接口定义  
│   ├── Mercury-job          项目消费端 mq消费/定时任务  
│   ├── Mercury-support      项目独有的一些支持方法  
│   └── Mercury-web          项目对外提供的接口  
├── common                  项目使用的公共的一些方法  
├── gen.md                  生成api/rpc的脚本 参考goctl  
└── sql                     项目sql文件  
```

#### 项目未完成功能

1. 对接管理平台

2. 实现对应的推送信息handler

3. 文件导入实时/定时推送

4. 客户端sdk