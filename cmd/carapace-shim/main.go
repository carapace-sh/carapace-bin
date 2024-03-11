package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace/pkg/xdg"
)

func main() {
	executable, err := os.Executable()
	if err != nil {
		panic(err.Error())
	}

	configDir, err := xdg.UserConfigDir()
	if err != nil {
		panic(err.Error())
	}

	name := strings.TrimSuffix(filepath.Base(executable), ".exe")
	args := []string{"--run", fmt.Sprintf("%v/carapace/specs/%v.yaml", configDir, name)}
	cmd := exec.Command("carapace", append(args, os.Args[1:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		panic(err.Error)
	}
}
