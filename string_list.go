package util

import (
	"github.com/alvinwsz/glog"
)

type StringList []string

func NewStringList() StringList {
	return make(StringList, 0)
}
func (l StringList) Len() int {
	return len(l)
}
func (l StringList) Add(t string) {
	l = append(l, t)
}

func (l StringList) Exist(t string) bool {
	for _, name := range l {
		if name == t {
			return true
		}
	}
	return false
}

func (l StringList) Remove(t string) {
	for i, name := range l {
		if name == t {
			l = append(l[:i], l[i+1:]...)
			break
		}
	}
}

func (l StringList) Get(i int) string {
	if i >= 0 && i < len(l) {
		return l[i]
	}
	return ""
}
func (l StringList) Fetch(callback func(string) error) error {
	for _, s := range l {
		err := callback(s)
		if err != nil {
			glog.ErrorDepth(1, err.Error())
		}
	}
	return nil
}
