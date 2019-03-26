### k8s部署封装

#### 配置k8s api地址
```
    k8s-clinet/deploay.go
    const k8s_api = "http://k8s.****.com" //k8s api地址
```

#### 创建项目
```bash
    curl 127.0.0.1:1711/create?name=**&app=**&image=**&pods=**&port=**
```

| 参数 | 说明 | 可空 |
|:--:|:--:|:--:|
|name|服务名称|N|
|app|应用名称|N|
|image|docker镜像地址|N|
|pods|启动服务数量|N|
|port|映射端口|N|

#### 更新项目
```bash
    curl 127.0.0.1:1711/update?name=**&image=**&pods=**
```
| 参数 | 说明 | 可空 |
|:--:|:--:|:--:|
|name|服务名称|N|
|image|docker镜像地址|N|
|pods|启动服务数量|N|
