package packageprac

import (
	"fmt"
	"runtime"
)

func Goversion() {
	fmt.Println(runtime.Version());
}