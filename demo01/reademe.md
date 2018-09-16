# 测试用例  
## 发送 HTTP GET 请求
```
$ curl -XGET http://127.0.0.1:8080/sd/health
OK

$ curl -XGET http://127.0.0.1:8080/sd/disk
OK - Free space: 16321MB (15GB) / 51200MB (50GB) | Used: 31%

$ curl -XGET http://127.0.0.1:8080/sd/cpu
CRITICAL - Load average: 2.39, 2.13, 1.97 | Cores: 2

$ curl -XGET http://127.0.0.1:8080/sd/ram
OK - Free space: 455MB (0GB) / 8192MB (8GB) | Used: 5%...

```
