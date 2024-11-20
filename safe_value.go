package bngsocket

import (
	"sync"
)

func newSafeBool(v bool) _SafeBool {
	return _SafeBool{
		newSafeValue(v),
	}
}

func newSafeBytes(v []byte) _SafeBytes {
	return _SafeBytes{
		newSafeValue(v),
	}
}

func newSafeInt(v int) _SafeInt {
	return _SafeInt{
		newSafeValue(v),
	}
}

func newSafeValue[T any](v T) _SafeValue[T] {
	mutex := new(sync.Mutex)
	return _SafeValue[T]{
		value:   &v,
		lock:    mutex,
		cond:    sync.NewCond(mutex),
		changes: 0,
	}
}

func newSafeAck() _SafeAck {
	return _SafeAck{
		_SafeChan: NewSafeChan[*_AckItem](),
	}
}

func newSafeMap[X any, T any]() _SafeMap[X, T] {
	return _SafeMap[X, T]{
		Map: new(sync.Map),
	}
}
