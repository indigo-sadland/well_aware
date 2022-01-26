package utils

import (
	"fmt"
	"log"
	"runtime"
)

var OStool string

func OsDetection() {

	var err error
	switch runtime.GOOS {
	case "linux":
		OStool = "xdg-open"
	case "windows":
		OStool = "rundll32"
	case "darwin":
		OStool = "open"
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
