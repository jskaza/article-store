package utils

import (
	"fmt"
	"os/exec"
)

// run this as goroutine
func ParsePaper(uuid, extension string) {
	switch extension {
	case ".md":
		cmd := exec.Command("pandoc", "-f", "markdown", "-t", "html", "-o", uuid+".html", uuid+extension)
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}
	case ".tex":
		exec.Command("pandoc", "-f", "latex", "-t", "html", "-o", uuid+".html", uuid+extension).Run()
	case ".docx":
		exec.Command("pandoc", "-f", "docx", "-t", "html", "-o", uuid+".html", uuid+extension).Run()
	default:
		return
	}
}
