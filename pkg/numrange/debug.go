package numrange

import (
	"fmt"
	"os"
)

func isDebugEnabled() bool {
	return os.Getenv("DEBUG") == "1"
}

func DebugPrintf(format string, a ...interface{}) {
	if isDebugEnabled() {
		fmt.Printf(format+"\n", a...)
	}
}
