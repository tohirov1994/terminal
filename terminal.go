package terminal

import (
	"os"
	"os/exec"
)


func Cleaner()  {
	commandTerminal := exec.Command("clear")
	commandTerminal.Stdout = os.Stdout
	commandTerminal.Run()
}
