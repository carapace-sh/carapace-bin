package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/completers"
	carapacebin "github.com/carapace-sh/carapace-bin/pkg/actions/tools/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	spec "github.com/carapace-sh/carapace-spec"
	"github.com/carapace-sh/carapace/pkg/uid"
	"github.com/spf13/cobra"
)

var invokeCmd = &cobra.Command{
	Use:   "invoke",
	Short: "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(invokeCompleter(args[0]))
	},
}

func init() {
	carapace.Gen(invokeCmd).Standalone()
	invokeCmd.Flags().SetInterspersed(false)

	carapace.Gen(invokeCmd).PositionalCompletion(
		carapacebin.ActionCompleters(true),
		bridge.ActionCarapaceBin("_carapace", "export", "", "_carapace").Shift(1).
			Filter("macro", "style"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch c.Args[1] {
			case "bash",
				"bash-ble",
				"cmd-clink",
				"elvish",
				"export",
				"fish",
				"ion",
				"nushell",
				"oil",
				"powershell",
				"tcsh",
				"xonsh",
				"zsh":
				command := c.Args[0]
				command = strings.Split(command, "@")[0] // strip group
				command = strings.Split(command, "/")[0] // strip variant
				return carapace.ActionValues(command)
			default:
				return carapace.ActionValues()
			}
		}),
	)

	carapace.Gen(invokeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch c.Args[1] {
			case "bash",
				"bash-ble",
				"cmd-clink",
				"elvish",
				"export",
				"fish",
				"ion",
				"nushell",
				"oil",
				"powershell",
				"tcsh",
				"xonsh",
				"zsh":
				args := []string{c.Args[0], "export"}
				args = append(args, c.Args[2:]...)
				args = append(args, c.Value)
				return carapace.ActionExecCommand("carapace", args...)(func(output []byte) carapace.Action {
					if string(output) == "" {
						return carapace.ActionValues()
					}
					return carapace.ActionImport(output)
				})
			default:
				return carapace.ActionValues()
			}
		}),
	)
}

func wrap(f func()) string { // TODO get rid of this
	old := os.Stdout
	r, w, _ := os.Pipe()
	// TODO defer w.Close() ?
	os.Stdout = w

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	f() // execute and grab output

	w.Close()
	out := <-outC
	os.Stdout = old
	return out
}

func invokeCompleter(nameVariant string) string {
	var out string
	if completer, err := completers.Lookup(nameVariant); err == nil { // TODO handle error
		switch {
		case completer.Execute == nil:
			carapace.LOG.Printf("Completer.Execute of %q is nil", nameVariant)
		case completer.Overlay != "" && len(os.Args) > 2:
			// TODO yuck
			shell := os.Args[2]

			os.Setenv("CARAPACE_LENIENT", "1")
			os.Args[2] = "export"
			out = wrap(func() {
				os.Args[1] = "_carapace"
				completer.Execute() // TODO Execute should handle the stdout wrapping
			})

			os.Args[2] = shell
			out = wrap(func() {
				complete(completer.Name, carapace.Batch(
					carapace.ActionImport([]byte(out)),
					spec.ActionSpec(completer.Overlay),
				).ToA())() // TODO handle error
			})
		default:
			out = wrap(func() {
				os.Args[1] = "_carapace"
				completer.Execute() // TODO handle error
			})
		}
	}

	executable := uid.Executable()
	patched := strings.ReplaceAll(string(out), fmt.Sprintf("%v _carapace", executable), fmt.Sprintf("%v %v", executable, nameVariant))      // general callback
	patched = strings.ReplaceAll(patched, fmt.Sprintf("'%v', '_carapace'", executable), fmt.Sprintf("'%v', '%v'", executable, nameVariant)) // xonsh callback
	return patched
}

func complete(name string, action carapace.Action) func() error {
	return func() error {
		cmd := &cobra.Command{
			Use:                name,
			DisableFlagParsing: true,
		}
		carapace.Gen(cmd).Standalone()
		carapace.Gen(cmd).PositionalAnyCompletion(
			action,
		)
		return cmd.Execute()
	}
}
