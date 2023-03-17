package gosms

import (
	"fmt"
	"github.com/hunterhug/marmot/miner"
	"testing"
)

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
