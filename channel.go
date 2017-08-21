package util

import (
	//"github.com/alvinwsz/glog"
	"github.com/alvinwsz/go-libnet"
)

type SessionMap map[string]*libnet.Channel

/*
1. 允许多终端同时登陆
	router已做优化处理，负载平衡时将同用户分配到同一个imserver
	用Channel来管理同用户的所有client session
	[user]--->Channel
2. Group群聊
	同一group的在线成员可能分布在不同imserver上
	同一ip上的所有成员组成一个Channel，以便广播
	[ip]--->Channel
*/

func NewSessionMap() SessionMap {
	return make(SessionMap)
}

func (l SessionMap) Exist(key string) bool {
	_, ok := l[key]
	return ok
}

func (l SessionMap) Len() int {
	return len(l)
}

func (l SessionMap) Add(key string, session *libnet.Session) {
	if _, ok := l[key]; !ok {
		l[key] = libnet.NewChannel()
	}
	l[key].Join(session, nil)
}
/* BUG: session.Close() will invokeCloseCallbacks registered by Channel.join(), which will call channel.Exit(session).
SO NEVER called
*/
func (l SessionMap) Remove(key string, session *libnet.Session) {
	if _, ok := l[key]; ok {
		l[key].Exit(session)
		if l[key].Len() == 0 {
			delete(l, key)
		}
	}
}
func (l SessionMap) Fetch(key string, callback func(*libnet.Session)) {
	if c, ok := l[key]; ok {
		c.Fetch(callback)
	}
}

func (l SessionMap) Broadcast(key string, encoder libnet.Encoder) {
	if c, ok := l[key]; ok {
		c.Broadcast(encoder)
	}
}
