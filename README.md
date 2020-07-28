# request-proxy
转发telegram请求


# go

## build

```bash
go mod init main
go mod download
go build -o ./app main.go
```

# k8s

## 构建 Image

```bash
docker build -t polywoo/telegram-proxy:v1 .
```

## 上传 Image

```bash
docker login
docker push polywoo/telegram-proxy:v1
```

## 部署

```bash
kubectl apply -f app.yaml
```
