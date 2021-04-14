# Attack-Seaman

简单来说，这是一个为 [ATT&CK® Navigator](https://mitre-attack.github.io/attack-navigator/) 服务的 [MITRE 知识库](https://github.com/mitre/cti/) 编辑工具。针对的人群主要是 ATT&CK 矩阵的使用用户。

目的是为解决以下问题：

1. 难以编辑大 JSON 文件；
2. 知识库字段繁多，易出错；
   - 详见 [STIX](https://oasis-open.github.io/cti-documentation/stix/intro.html)

### Usage

1. 上传本地知识库
2. 新增（子）Technique 与 Relationship
3. 发布到 navigator
   - 底层仅修改了 navigator 的 assets 来源
   - 不对 upstream 做任何破坏

为使 navigator 成功加载，你的编辑动作需满足以下条件：

1. 前端列出的必填项
2. 子技术的 kill_chase_phase 应当是 父技术的 kill_chase_phase 的子集
3. relationship 的从属不能是同一个 Technique
4. `x_mitre_is_subtechnique` 的 `true/false` 应当与定义的 relationship 从属关系一致

由于项目硬编码了设想的域名 `attack-seaman.com` 在代码中，你可能需要：

### 编辑 Hosts

| 操作系统 | Hosts 文件位置                          |
| -------- | --------------------------------------- |
| Linux    | `/etc/hosts`                            |
| Windows  | `C:\Windows\System32\drivers\etc\hosts` |
| Mac      | `/private/etc/hosts`                    |

添加一行

```bash
127.0.0.1 attack-seaman.com
```

### Docker

```sh
# 通过 docker 启动项目
docker-compose up -d

# 查看特定 container （如：app）日志
docker-compose logs app

# 访问以下地址，如数据、操作等均无误则说明启动成功
# 前端：attack-seaman.com:3000
# 后端API：attack-seaman.com:6868
# self-host navigator：attack-seaman.com:4200
```

### Dev

如果需要本地开发，确保满足如下环境：

```sh
(master)⚡ % systemctl status mongodb

● mongodb.service - High-**performance**, schema-free document-oriented database
     Loaded: loaded (/usr/lib/systemd/system/mongodb.service; disabled; vendor preset: disabled)
     Active: active (running) since Tue 2021-04-13 17:39:17 CST; 46s ago
   Main PID: 4116 (mongod)
      Tasks: 32 (limit: 17833)
     Memory: 211.2M
     CGroup: /system.slice/mongodb.service
             └─4116 /usr/bin/mongod --quiet --config /etc/mongodb.conf

4月 13 17:39:17 X systemd[1]: Started High-performance, schema-free document-oriented database.
(master)⚡ % node -v
v15.12.0
(master)⚡ % go version
go version go1.16.2 linux/amd64


```

#### /src/config.yml 
```yml
# 1. 配置好 mongodb connection
database:
  dbname: attackSeaman
  connection: mongodb://attackSeaman:cuccs1sgreat@127.0.0.1:27017
  tbname: enterprise

# 2. 上传文件的保存路径
utils:
  filepath: /app/assets/
```
#### API

```sh
cd src
go run .
```

#### frontend

```sh
# yarn 不行就试试 npm
cd frontend
yarn & yarn start
```

#### navigator
```sh
cd src/nav-app
yarn & yarn start

```


### Mitre ATT&CK 编辑
添加好 (Sub)Technique 以及 Relationship(如果需要)

#### 新建 Technique
![](./assets/attackpatterns_create.png)

#### 新建 Relationship
![](./assets/relationship_create.png)


ATT&CK® Navigator 新增的技术点

![](./assets/res.png)

# Refs

- [Ten-Minutes-App](https://github.com/Kirk-Wang/Ten-Minutes-App)

* [mongo-go-dirver offical examples](https://github.com/mongodb/mongo-go-driver/blob/master/examples/documentation_examples/examples.go)
* [Go by Example](https://gobyexample.com/)
* [gotify/server](https://github.com/gotify/server)
* [gin-jwt](https://github.com/appleboy/gin-jwt)