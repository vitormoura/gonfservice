package main

import (
	"time"
)

//MessageType defines a type to messages
type MessageType int

const (

	//PlainMail plain text e-mail message
	PlainMail MessageType = iota + 1

	//RichFormatMail e-mail with some sort of rich format
	RichFormatMail
)

//Message defines the properties of a message
type Message struct {
	Type    MessageType
	From    string
	To      []string
	CC      []string
	Subject string
	Content string
}

//SendMessageResult defines the result of a send message operation
type SendMessageResult struct {
	MessageID string
	Date      time.Time
	Success   bool
	Errors    []string
}

//MessageSender defines operations to send messages
type MessageSender interface {

	//Send sends the message
	Send(m Message) (SendMessageResult, error)
}

//MessageSenderCreator defines a function capable to create message senders
type MessageSenderCreator func() MessageSender
