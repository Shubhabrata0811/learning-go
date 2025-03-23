package golang

import (
	"fmt"
	"runtime"
)

func Goversion() {
	fmt.Println(runtime.Version());
}