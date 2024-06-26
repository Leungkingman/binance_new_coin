package service

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type AliSMS struct {
	AccessKeyId  string
	AccessSecret string
	Sign         string
}

//type SendSmsRequest struct {
//	*requests.RpcRequest
//	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
//	SmsUpExtendCode      string           `position:"Query" name:"SmsUpExtendCode"`
//	SignName             string           `position:"Query" name:"SignName"`
//	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
//	PhoneNumbers         string           `position:"Query" name:"PhoneNumbers"`
//	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
//	OutId                string           `position:"Query" name:"OutId"`
//	TemplateCode         string           `position:"Query" name:"TemplateCode"`
//	TemplateParam        string           `position:"Query" name:"TemplateParam"`
//}

func (ali AliSMS) SendSms(tplCode string, phone string) (*dysmsapi.SendSmsResponse, error) {
	client, _ := dysmsapi.NewClientWithAccessKey("ap-northeast-1", ali.AccessKeyId, ali.AccessSecret)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phone
	request.SignName = ali.Sign
	request.TemplateCode = tplCode
	response, err := client.SendSms(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
