package utils

import (
	"path"
	"runtime"
)

func Path(position int) string {
	_, filename, _, ok := runtime.Caller(position)

	if !ok {
		panic("unable to get the current filename")
	}
	__dirname := path.Dir(filename)

	return __dirname
}
