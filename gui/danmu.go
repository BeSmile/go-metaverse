package gui

import (
	mLive "go-metaverse/models/live/models"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <init.h>
*/
import "C"

type DanMuClient struct {
}

func (dm *DanMuClient) NewChatMessage(message mLive.ChatMsgMessage) {
	C.InitDataSource(C.CString(message.Ic), C.CString(message.Nn), C.CString(message.Txt), C.CString(string(message.Type)))
}

func (dm *DanMuClient) NewUEnterMessage(message mLive.UenterMessage) {
	C.InitDataSource(C.CString(message.Ic), C.CString(message.Nn), C.CString(message.Txt), C.CString(string(message.Type)))
}

func (dm *DanMuClient) Init() {
	C.StartApp()
}

func NewClient() *DanMuClient {
	return &DanMuClient{}
}
