package gosms

import "fmt"

type ClientInterface interface {
	SendMessage(mobile string, templateId string, params []string) (*SendMessageResponse, error)
}

type SendMessageResponse struct {
	// 短信总计费条数
	TotalFee int64
	// 发送情况
	SendMessageData
}

type SendMessageData struct {
	Mobile string
	// 第三方ID
	Sid string
}

type InnerErr struct {
	Msg  string
	Code int64
}

func NewInnerError(code int64, msg string) *InnerErr {
	return &InnerErr{
		Code: code,
		Msg:  msg,
	}
}
func (innerErr *InnerErr) Error() string {
	return fmt.Sprintf("%d|%s", innerErr.Code, innerErr.Msg)
}

func IsErr(e error) (innerErr *InnerErr, ok bool) {
	innerErr, ok = e.(*InnerErr)
	return
}
