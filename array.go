package util

type Array []interface{}

func NewArray() Array {
	return make(Array, 0)
}

func (l Array) Len() int {
	return len(l)
}

func (l Array) Get(i int) interface{} {
	if i >= 0 && i < len(l) {
		return l[i]
	}
	return nil
}

func (l Array) Add(t interface{}) {
	l = append(l, t)
}

func (l Array) Exist(t interface{}) bool {
	for _, name := range l {
		if name == t {
			return true
		}
	}
	return false
}

func (l Array) Remove(t interface{}) {
	for i, name := range l {
		if name == t {
			l = append(l[:i], l[i+1:]...)
			break
		}
	}
}

func (l Array) Fetch(callback func(interface{}) error) error {
	for _, s := range l {
		err := callback(s)
		if err != nil {
			return err
		}
	}
	return nil
}
