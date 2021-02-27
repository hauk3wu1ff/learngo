package version

import (
	"fmt"
	"runtime"
)

func GoVersion() {
	fmt.Println("The current golang version is", runtime.Version())
}
