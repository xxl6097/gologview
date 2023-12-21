# 基于Golang语言的Web端查看日志

## 添加依赖

```shell

go get -u github.com/xxl6097/gologview

```
## 使用步骤

```go

func init() {
	go logview.GetLogApi().RunAndSetUserPass(8080, "admin", "admin")
}

func main() {
  for {
		fmt.Println("aaaaaa这个不会在web上显示哦")
		glog.Info("info...")
		time.Sleep(time.Second)
	}
}

```
