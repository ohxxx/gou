package bridge

import "fmt"

//【创建型】
// 桥接模式
// 		意图：将抽象部分与实现部分分离，使它们都可以独立的变化
// 		解决：在有多种可能会变化的情况下，用继承会造成类爆炸问题，扩展起来不灵活
// 		实例：有抽象的消息和消息的实现

// 抽象的消息接口
type MessageAbstract interface {
	SendMessage(msg, to string)
}

// 实现的消息接口
type MessageImpl interface {
	Send(msg, to string)
}

// SMS 消息
type MessageSMS struct{}

func (m *MessageSMS) Send(msg, to string) {
	fmt.Printf("Send %s to %s via SMS\n", msg, to)
}

func ViaSMS() MessageImpl {
	return &MessageSMS{}
}

// Email 消息
type MessageEmail struct{}

func (m *MessageEmail) Send(msg, to string) {
	fmt.Printf("Send %s to %s via Email\n", msg, to)
}

func ViaEmail() MessageImpl {
	return &MessageEmail{}
}

// 抽象的实现
type CommonMessage struct {
	method MessageImpl
}

func NewCommonMessage(method MessageImpl) *CommonMessage {
	return &CommonMessage{
		method: method,
	}
}

func (m *CommonMessage) SendMessage(msg, to string) {
	m.method.Send(msg, to)
}

type UrgencyMessage struct {
	method MessageImpl
}

func NewUrgencyMessage(method MessageImpl) *UrgencyMessage {
	return &UrgencyMessage{
		method: method,
	}
}

func (m *UrgencyMessage) SendMessage(msg, to string) {
	m.method.Send(fmt.Sprintf("[Urgency] %s", msg), to)
}
