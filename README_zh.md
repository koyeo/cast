<p align="center">
  <img src="./logo.png" alt="Cast" width="200" />
</p>

# Cast

[English](./README.md)

适用于快速交付的本地集成和部署工具。

## 为什么选择 Cast？

CI/CD 的工具很多，比如 Jenkins、GitHub Actions、Makefile、Ansible、Travis CI 等，但它们太重了，不方便快速交付：

- **Jenkins** 要部署软件，通过网页使用。
- **GitHub Actions / Travis CI** 等平台性工具限定了使用场景。
- **Ansible + Makefile** 二者要相互搭配才能做到构建到部署工作流执行，但是每次都要配一遍，很麻烦。

Cast 通过一个配置文件和命令行工具解决这些痛点。

### 适用场景

- 特别适合独立的全栈开发工程师。
- 想快速将项目发布到服务器上，得到效果的反馈。
- 快速发布前端项目。

### 更优的选择

在多人协作的生产环境，可能需要严格的发版管理，此时不建议使用 Cast。

## 安装

```bash
go install github.com/koyeo/cast@latest
```

> **注：** `go install` 将会把 `cast` 编译安装在 `$GOPATH/bin` 目录下，安装前请检查 `$GOPATH` 指向位置，且是否添加到 `$PATH` 路径下。

## 快速上手

### 1. 初始化配置

```bash
cast init
```

执行后将：
1. 如果目录下不存在 `cast.yaml` 文件，则创建该文件。
2. 在 `.gitignore` 添加 `.cast` 行，以忽略 Cast 临时工作目录。

也可以指定自定义配置文件名：

```bash
cast init cast.production.yml
```

### 2. 编辑 `cast.yaml`

通过一些配置示例，实现本地构建、部署到服务器、重启服务：

```yaml
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
      - use: hi                                # 继承 hi 任务的 steps
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

### 3. 执行任务

```bash
cast run task-1
```

执行多个任务：

```bash
cast run task-1 hi
```

## CLI 参考

### `cast init`

初始化 `cast.yaml` 配置文件，并自动更新 `.gitignore` 文件。

### `cast run <task...>`

执行一个或多个任务。

### `cast list`

列出配置文件里的资源项，包括任务、服务器、环境变量等。

## 配置参考

### 服务器

```yaml
servers:
  server_1:                         # 服务器标识，可以在 deploy 任务中通过 use 引用
    comment: 第一台服务器             # 备注
    host: 192.168.1.5               # 服务器地址
    port: 2222                      # 端口，默认使用 22
    user: root                      # 服务器用户名
    password: 123456                # 服务器密码，可以由 identity_file 选项替代
    identity_file: ~/.ssh/id_rsa    # 服务私钥认证文件，默认使用 ~/.ssh/id_rsa
```

### 环境变量

```yaml
envs:
  k1: v1                            # 通过键值对配置全局变量
  k2: v2
```

### 部署文件映射

| source  | target            | 服务器存放位置               |
|:--------|:------------------|:----------------------|
| `file1` | `/app/test/file1` | `/app/test/file1`     |
| `file1` | `/app/test/file2` | `/app/test/file2`     |
| `file1` | `/app/test`       | `/app/test`           |
| `file1` | `/app/test/`      | `/app/test/file1`     |
| `dir1`  | `/app/test/dir2`  | `/app/test/dir2`      |
| `dir1`  | `/app/test/dir2/` | `/app/test/dir1/dir2` |
| `dir1`  | `/app/test/`      | `/app/test/dir1`      |
| `dir1`  | `/app/test`       | `/app/test`           |

## 反馈

如果你有使用上的问题，或想参与项目的开发，可以通过邮箱联系：koyeo@qq.com。

## 贡献

欢迎提交 Pull Request。如果是较大的改动，请先创建 Issue 讨论。

## License

[MIT](https://choosealicense.com/licenses/mit/)
