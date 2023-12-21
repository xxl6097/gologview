# 基于Golang语言的Web端查看日志

## 添加依赖

```shell

go get -u github.com/xxl6097/gologview

```
## 使用步骤

```shell

func main() {
	logview.New().Start(8080)
}

```
