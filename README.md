## logurs

#### 开启 Debug 模式
```go
logrus.SetLevel(logrus.DebugLevel)
```
#### json 格式
```go
logrus.SetFormatter(&logrus.JSONFormatter{})
```
输出
```yaml
{"level":"info","msg":"I am info","name":"test","namespace":"kube-system","time":"2021-08-06T16:51:11+08:00"}
```