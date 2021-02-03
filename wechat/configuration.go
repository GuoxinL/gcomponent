/*
  Created by guoxin in 2020/6/5 10:48 上午
*/
package wechat

//
//import (
//    "fmt"
//    "github.com/GuoxinL/gcomponent/core/tools"
//    "github.com/GuoxinL/gcomponent/environment"
//    "github.com/GuoxinL/gcomponent/logging"
//    "time"
//)
//
//func init() {
//	new(Configuration).Initialize()
//}
//
//const (
//	content_title = "### **%s**\n"
//	content_field = ">**%s**: <font color=\"black\">%v</font>\n"
//)
//
//var warningInstances *Configuration
//
//type Configuration struct {
//	Url          string `mapstructure:"url"`
//	Enable       bool   `mapstructure:"enable"`
//	LocalAddress string `mapstructure:"localAddress"`
//}
//
//func (this *Configuration) Initialize(params ...interface{}) interface{} {
//	logging.Info("GComponent [wechat-warning]初始化接口")
//	err := environment.GetProperty("components.wechat-warning", &this)
//	if err != nil {
//		logging.Exitf("GComponent [wechat-warning]读取配置异常, 退出程序！！！")
//	}
//	if this.Enable {
//		_ = logging.Warn("GComponent [wechat-warning]pprof未开启，如需开启请配置'components.wechat-warning.enable=true'")
//		return nil
//	}
//	ip, err := tools.GetLocalIP()
//	if err != nil {
//		this.LocalAddress = fmt.Sprint(content_field, "告警主机IP", "获得告警主机IP异常"+err.Error())
//	} else {
//		this.LocalAddress = fmt.Sprint(content_field, "告警主机IP", ip)
//	}
//	warningInstances = this
//	logging.Info("GComponent [wechat-warning] init success")
//	return nil
//}
//
//type messageType string
//
//const (
//	Markdown messageType = "markdown"
//	Text     messageType = "text"
//)
//
//type MessageTemplate struct {
//	// 标题
//	Title string
//	// 告警类型
//	MessageType messageType
//	// 这里放 DemoMessage
//	Message interface{}
//}
//
//type DemoMessage struct {
//	ErrorMessage  string `json:"异常信息"`
//	ErrorCode     int    `json:"异常码"`
//	InterfaceName string `json:"接口名称"`
//	SNID          string `json:"SNID"`
//}
//
///**
//Example
//wechat_warning.SendMessage(wechat_warning.MessageTemplate{
//	Title:       "标题",
//	MessageType: wechat_warning.Markdown,
//	Message: wechat_warning.DemoMessage{
//		ErrorMessage:  "我错了",
//		ErrorCode:     32137839213891,
//		InterfaceName: "测试接口",
//		SNID:          "372193789129",
//	},
//})
//实际打印：
//标题
// 异常信息: 我错了
// 异常码: 3.2137839213891e+13
// 接口名称: 测试接口
// SNID: 372193789129
//
//自带默认字段：
//告警时间：2006-01-02 15:04:05
//告警主机IP：2006-01-02 15:04:05
//*/
//func SendMessage(mt MessageTemplate) {
//	warningInstances.SendMessage(mt)
//}
//
//func (this *Configuration) SendMessage(mt MessageTemplate) {
//	if !this.Enable {
//		_ = logging.Warn("GComponent [wechat-warning]发送告警未开启，如需开启请配置'components.wechat-warning.enable=true'")
//		return
//	}
//	messageRequest := new(messageRequest)
//	messageResponse := new(messageResponse)
//	messageRequest.MsgType = string(mt.MessageType)
//	switch mt.MessageType {
//	case Markdown:
//		messageRequest.Markdown = this.getContentWrapper(mt)
//	case Text:
//		messageRequest.Text = this.getContentWrapper(mt)
//	default:
//		messageRequest.Text = this.getContentWrapper(mt)
//	}
//
//	err := tools.Post(this.Url, messageRequest, messageResponse, 10*time.Second)
//	if err != nil {
//		logging.Error0("GComponent [wechat-warning]发送告警信息异常，异常信息：%v", err.Error())
//	} else {
//		_ = logging.Warn("GComponent [wechat-warning]发送告警信息成功! Code: %v, Message: %v", messageResponse.Code, messageResponse.Message)
//	}
//}
//
//type contentWrapper struct {
//	Content string `json:"content"`
//	// userid的列表，提醒群中的指定成员(@某个成员)，@all表示提醒所有人，如果开发者获取不到userid，可以使用mentioned_mobile_list
//	MentionedList string `json:"mentioned_list,omitempty"`
//	// 手机号列表，提醒手机号对应的群成员(@某个成员)，@all表示提醒所有人
//	MentionedMobileList string `json:"mentioned_mobile_list,omitempty"`
//}
//
//type messageRequest struct {
//	MsgType  string         `json:"msgtype"`
//	Text     contentWrapper `json:"text,omitempty"`
//	Markdown contentWrapper `json:"markdown,omitempty"`
//}
//
///**
//{
//   "errcode": 0,
//   "errmsg": "ok"，
//   "type": "file",
//   "media_id": "1G6nrLmr5EC3MMb_-zK1dDdzmd0p7cNliYu9V5w7o8K0",
//   "created_at": "1380000000"
//}
//
//type		媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件(file)
//media_id	媒体文件上传后获取的唯一标识，3天内有效
//created_at	媒体文件上传时间戳
//
//*/
//type messageResponse struct {
//	Code      int    `json:"errcode,omitempty"`
//	Message   string `json:"errmsg,omitempty"`
//	Type      string `json:"type,omitempty"`
//	MsgType   string `json:"msgtype,omitempty"`
//	MediaId   string `json:"media_id,omitempty"`
//	CreatedAt string `json:"created_at,omitempty"`
//}
//
//func (this *Configuration) getContentWrapper(mt MessageTemplate) (wrapper contentWrapper) {
//	toJson, err := tools.ToJson(mt.Message)
//	if err != nil {
//		wrapper.Content += "Message ToJson转换异常"
//		return
//	}
//	wrapper = *new(contentWrapper)
//	wrapper.Content += fmt.Sprintf(content_title, mt.Title)
//
//	fieldMap := make(map[string]interface{})
//	err = tools.ToObject(toJson, &fieldMap)
//	if err != nil {
//		wrapper.Content += "Message ToObject转换异常：" + err.Error()
//		return
//	}
//	for fieldName, fieldValue := range fieldMap {
//		wrapper.Content += fmt.Sprintf(content_field, fieldName, fieldValue)
//	}
//	wrapper.Content += fmt.Sprint(content_field, "告警时间", time.Now().Format("2006-01-02 15:04:05"))
//	wrapper.Content += fmt.Sprint(content_field, "告警主机IP", this.LocalAddress)
//
//	return
//}
