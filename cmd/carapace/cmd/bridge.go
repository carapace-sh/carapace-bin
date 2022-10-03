package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

func bridgeCompletion(command string, engine string, args ...string) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// TODO handle error
	cmd := createCmd(command, engine)

	a := []string{"_carapace"}
	a = append(a, args...)
	cmd.SetArgs(a)
	cmd.Execute()

	w.Close()
	out := <-outC
	os.Stdout = old

	executable, err := os.Executable()
	if err != nil {
		panic(err.Error()) // TODO exit with error message
	}

	executableName := filepath.Base(executable)
	patched := strings.Replace(string(out), fmt.Sprintf("%v _carapace", executableName), fmt.Sprintf("%v --bridge %v/%v", executableName, command, engine), -1)         // general callback
	patched = strings.Replace(patched, fmt.Sprintf("'%v', '_carapace'", executableName), fmt.Sprintf("'%v', '--bridge', '%v/%v'", executableName, command, engine), -1) // xonsh callback
	fmt.Print(patched)

}

func createCmd(command string, engine string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                command,
		DisableFlagParsing: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		Run: func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(cmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch engine {
			case "argcomplete":
				return bridge.ActionArgcomplete(command)
			case "click":
				return bridge.ActionClick(command)
			case "cobra":
				return bridge.ActionCobra(command)
			case "fish":
				return bridge.ActionFish(command)
			case "posener":
				return bridge.ActionPosener(command)
			default:
				return carapace.ActionValues()
			}
		}),
	)

	return cmd
}
