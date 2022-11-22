package main

import (
	"signaling/src/framework"
	"signaling/src/action"
)

func init() {
	framework.GActionRouter["/signaling/push"] = action.NewPushAction()

	framework.GActionRouter["/xrtcclient/push"] = action.NewXrtcClientPushAction()
}
