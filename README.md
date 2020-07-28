## telegram-proxy
转发telegram请求，解决TG通知无法发送的问题

示例
```bash
curl --location --request POST 'https://<your dmain>/<sender>:<telegram secret>/sendMessage' \
--header 'Content-Type: application/json' \
--data-raw '{
	"chat_id": <chat id>,
	"text": "hello"
}'
```

## 其他
### 构建 Image

```bash
docker build -t polywoo/telegram-proxy:v1 .
```

### 上传 Image

```bash
docker login
docker push polywoo/telegram-proxy:v1
```

### K8S部署

```bash
kubectl apply -f app.yaml
```
