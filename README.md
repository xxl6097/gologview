# 基于Golang语言的Web端查看日志

## 添加依赖

```shell

go get -u github.com/xxl6097/gologview

```
## 使用步骤

```shell
func onLogHook(log []byte) {
	logview.GetLogApi().Send(log)
}

func init() {
	glog.Hook(onLogHook)
	logview.GetLogApi().SetUser("admin", "admin")
	go func() {
		logview.GetLogApi().Start(8080)
	}()
}

func main() {
  for {
		fmt.Println("aaaaaa")
		time.Sleep(time.Second)
	}
}

```
