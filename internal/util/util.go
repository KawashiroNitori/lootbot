package util

import "io"

func Close(closer io.Closer) {
	_ = closer.Close()
}
