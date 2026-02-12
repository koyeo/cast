# Cast

适用于快速交付的本地集成和部署工具。

## 安装

```bash
go install github.com/koyeo/cast@latest
```

> 注： go install 将会把 cast 编译安装在 $GOPATH/bin 目录下， 安装前请检查 $GOPATH 指向位置，且是否添加的 $PATH 路径下。

## 初始化配置

```bash
cast init
```

1. 如果目录下不存在 `cast.yml` 文件，则创建该文件。
2. 在 `.gitignore` 添加 `.cast` 行，以忽略 Cast 临时工作目录。

## 第一个任务

通过一些配置示例，实现如下功能：

1. 本地完成构建。
2. 将构建结果发布到服务器指定位置。
3. 在服务器执行重启。

**编辑 cast.yml：**

```yml
version: 1.0
servers:
  server-1:
    comment: 示例服务器
    host: 192.168.1.10
    user: root                                 # 默认使用 ~/.ssh/id_rsa 私钥进行认证
tasks:
  task-1:                                      # 任务名称
    comment: 示例任务                           # 任务注释
    steps:
      - run: go build -o foo foo.go            # 本地执行构建
      - deploy:
          servers:
            - use: server-1                    # 部署服务器
          mappers:
            - source: ./foo                    # 本地文件路径
              target: /app/foo/bin/foo         # 服务器存放位置
          executes:
            - run: supervisorctl restart foo   # 服务器重启服务
      - run: rm foo                            # 本地清理
  hi:
    comment: 打个招呼
    steps:
      - run: echo "Hi! this is from cast~" 
```

**执行工作流：**

```
cast run task-1
```

更多用法参见文档：[https://cast.kozilla.io](https://cast.kozilla.io)。

## 贡献
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
