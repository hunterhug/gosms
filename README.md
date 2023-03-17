# 短信发送 Golang SDK

目前支持云之讯 open2 接口：https://office2.ucpaas.com 。

# 如何使用

```go
go get -u -v github.com/hunterhug/gosms
```

例子：

```go
func TestUSmsClient_SendMessage(t *testing.T) {
	miner.SetLogLevel(miner.Level(miner.DEBUG))

	clientId := ""
	password := ""

	client, err := NewUSmsClient(clientId, password)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	mobile := ""

	templateId := "107"
	params := []string{"1", "2"}

	message, err := client.SendMessage(mobile, templateId, params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(message)
}
```