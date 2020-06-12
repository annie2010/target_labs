// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package internal

import (
	"path"
	"runtime"
)

// FuncName returns the current function name
func FuncName(back int) string {
	pc := make([]uintptr, 1)
	runtime.Callers(back+2, pc)
	return path.Base(runtime.FuncForPC(pc[0]).Name())
}
