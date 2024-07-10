# 部署
```shell
docker compose up
```

# 项目结构
```
.
├── README.md
├── docker-compose.yml
├── Dockerfile
├── main.go
└── config
    ├── prometheus.yml
```

# 说明
## 端口
- 9090: Prometheus UI
- 8080: gin

## 配置
- Prometheus 配置文件: config/prometheus.yml

## 路由
- /metrics: Prometheus 指标
- /hello: Hello World