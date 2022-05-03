package util

import (
	"fmt"
	"github.com/phuslu/log"
	"runtime/debug"
)

func Recover[T any](f func(arg T)) func(arg T) {
	fr := func(arg T) {
		defer func() {
			if r := recover(); r != nil {
				if e, ok := r.(error); ok {
					log.Error().Msg(fmt.Sprintf("%v\n%s", e, string(debug.Stack())))
				} else {
					log.Error().Msg(fmt.Sprintf("unknown panic\n%s", string(debug.Stack())))
				}
			}
		}()
		f(arg)
	}

	return fr
}
