package utils

import (
	"os/exec"
	"runtime"
)

var (

	OStool 	string
	Cmd		exec.Cmd
	CmdArgs	[]string

)

// OsDetection checks user's OS and selects appropriate command and arguments for CMD.
func OsDetection() {

	var err error
	var toolPath string

	switch runtime.GOOS {
	case "linux":

		OStool = "xdg-open"
		toolPath, err = exec.LookPath(OStool)
		CmdArgs = append(CmdArgs, OStool)

		Cmd = exec.Cmd{
			Path: toolPath,
			Args: CmdArgs,
		}

	case "windows":

		OStool = "rundll32"
		toolPath, err = exec.LookPath(OStool)
		OStoolArg := "url.dll,FileProtocolHandler"
		CmdArgs = append(CmdArgs, OStool, OStoolArg)

		Cmd = exec.Cmd{
			Path: toolPath,
			Args: CmdArgs,
		}

	case "darwin":

		OStool = "open"
		toolPath, err = exec.LookPath(OStool)
		CmdArgs = append(CmdArgs, OStool)

		Cmd = exec.Cmd{
			Path: toolPath,
			Args: CmdArgs,
		}

	default:
		ErrorLogger.Println("unsupported platform")
	}

	if err != nil {
		ErrorLogger.Println(err)
	}
}
