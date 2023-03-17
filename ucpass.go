package gosms

import (
	"encoding/json"
	"errors"
	"github.com/hunterhug/marmot/miner"
	"github.com/hunterhug/marmot/util"
	"strings"
)

const (
	// 变量模板短信下行
	uVariableSmsUrl = "http://open2.ucpaas.com/sms-server/variablesms"
)

// 云之讯短信客户端
type uSmsClient struct {
	ClientId string
	Password string
}

// NewUSmsClient 新建云之讯客户端
func NewUSmsClient(clientId, password string) (ClientInterface, error) {
	if clientId == "" || password == "" {
		return nil, errors.New("config empty")
	}

	c := new(uSmsClient)
	c.ClientId = clientId
	c.Password = strings.ToLower(util.Md5(password))
	return c, nil
}

type uSmsMessageRequest struct {
	ClientId   string `json:"clientid"`
	Password   string `json:"password"`
	Mobile     string `json:"mobile"`
	TemplateId string `json:"templateid"`
	Param      string `json:"param"`
}

type uSmsMessageResponse struct {
	TotalFee int64             `json:"total_fee"`
	Code     int64             `json:"code"`
	Msg      string            `json:"msg"`
	Data     []uSmsMessageData `json:"data"`
}

type uSmsMessageData struct {
	Fee    int64  `json:"fee"`
	Mobile string `json:"mobile"`
	Sid    string `json:"sid"`
}

// SendMessage 发送短信
// http://open2.ucpaas.com/sms-server/variablesms
func (c *uSmsClient) SendMessage(mobile string, templateId string, params []string) (*SendMessageResponse, error) {
	if len(mobile) == 0 || templateId == "" {
		return nil, errors.New("req empty")
	}
	w := miner.NewAPI()

	req := &uSmsMessageRequest{
		ClientId:   c.ClientId,
		Password:   c.Password,
		Mobile:     mobile,
		TemplateId: templateId,
		Param:      strings.Join(params, ";"),
	}

	raw, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	miner.Logger.Debugf("uSmsClient SendMessage req: %s", string(raw))

	result, err := w.SetUrl(uVariableSmsUrl).SetBData(raw).PostJSON()
	if err != nil {
		return nil, err
	}

	miner.Logger.Debugf("uSmsClient SendMessage resp: %s", string(result))

	resp := new(uSmsMessageResponse)
	err = json.Unmarshal(result, resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, NewInnerError(resp.Code, resp.Msg)
	}

	returnData := new(SendMessageResponse)
	returnData.TotalFee = resp.TotalFee
	for _, v := range resp.Data {
		returnData.Mobile = v.Mobile
		returnData.Sid = v.Sid
	}

	return returnData, nil
}
