# grpc-todolist

## 运行
```bash
make env-up         # 启动环境
make user           # 启动用户摸块
make task           # 启动任务模块
make api            # 启动网关
```
## 项目结构

### 调用关系
```
                                      http
                           ┌────────────────────────┐
 ┌─────────────────────────┤                        ├───────────────────────────────┐
 │                         │          api           │                               │
 │      ┌─────────────────►|                        │◄──────────────────────┐       │
 │      │                  └───────────▲────────────┘                       │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                           resolve                                 │       │
 │      │                              │                                    │       │
req    resp                            │                                   resp    req
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                   ┌──────────▼─────────┐                          │       │
 │      │                   │                    │                          │       │
 │      │       ┌──────────►|        Etcd        |◄────────────────┐        │       │
 │      │       │           │                    │                 │        │       │
 │      │       │           └────────────────────┘                 │        │       │
 │      │       │                                                  │        │       │
 │      │     register                                           register   │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
┌▼──────┴───────┴───┐                                           ┌──┴────────┴───────▼─┐
│                   │───────────────── req ────────────────────►│                     │
│       user        │                                           │         task        │
│                   │◄──────────────── resp ────────────────────│                     │
└───────────────────┘                                           └─────────────────────┘
      protobuf                                                           protobuf

```
### 整体
```
.                 
├─.idea             
├─cmd               
│  ├─api                # 网关模块
│  │  ├─handler
│  │  ├─middleware
│  │  ├─routes
│  │  ├─rpc
│  │  └─types
│  ├─task               # todolist模块
│  │  ├─dal
│  │  │  └─db
│  │  ├─model
│  │  ├─pack
│  │  └─service
│  └─user               # 用户模块
│      ├─dal
│      │  └─db
│      ├─model
│      ├─pack
│      └─service
├─config                # 配置
│  └─sql                    # init.sql
├─consts                # 常量定义
├─idl                   # proto定义
│  └─pb                     # proto生成类
│      ├─task
│      └─user
└─pkg                   # 通用包
    ├─discovery             # 服务注册与发现
    ├─errno                 # 返回码
    └─utils                 # 杂项工具

## 指令列表
```bash
make env-up         # 启动环境
make env-down       # 结束环境
make proto          # 更新protoc
make user           # 启动用户摸块
make task           # 启动任务模块
make api            # 启动网关
```