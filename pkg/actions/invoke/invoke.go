// package invoke contains carapace-bin internal actions
package invoke

import (
	"bytes"
	"io"
	"os"

	"github.com/rsteube/carapace"
)

// ActionInvoke invokes an internal carapace-bin completer
//   var composeCmd = &cobra.Command{
//   	Use:                "compose",
//   	Short:              "Define and run multi-container applications with Docker",
//   	Run:                func(cmd *cobra.Command, args []string) {},
//   	DisableFlagParsing: true,
//   }
//
//   func init() {
//   	carapace.Gen(composeCmd).Standalone()
//
//   	rootCmd.AddCommand(composeCmd)
//
//   	carapace.Gen(composeCmd).PositionalAnyCompletion(
//   		invoke.ActionInvokeCompleter(compose.Execute),
//   	)
//   }
func ActionInvoke(f func() error) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO experimental
		// TODO beware of carapace.Batch goroutines - is this safe? might need locking
		old := os.Args
		args := []string{"", "_carapace", "export", "_", ""}
		args = append(args, c.Args...)
		args = append(args, c.CallbackValue)
		os.Args = args
		output, err := captureStdout(f)
		os.Args = old
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return carapace.ActionImport([]byte(output))
	})
}

func captureStdout(f func() error) (string, error) {
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

	err := f()

	w.Close()
	out := <-outC
	os.Stdout = old

	return out, err
}
