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
	"github.com/carapace-sh/carapace/pkg/uid"
	"github.com/spf13/cobra"
)

var invokeCmd = &cobra.Command{
	Use:   "invoke",
	Short: "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// if completer := args[0]; strings.Contains(completer, "/") { // TODO this isn't correct any more in regards to variants
		// 	// patch to base if resolves to path executable (#2971)
		// 	base := filepath.Base(completer)
		// 	if abs, err := execabs.LookPath(base); err == nil && args[0] == abs {
		// 		args[0] = base
		// 	}
		// }
		// invoke(cmd, args)
		fmt.Print(invokeCompleter(args[0]))
	},
}

// func invoke(cmd *cobra.Command, args []string) {
// 	if overlayPath, err := overlayPath(args[0]); err == nil && len(args) > 2 { // and arg[1] is a known shell
// 		cmd := &cobra.Command{
// 			DisableFlagParsing: true,
// 			CompletionOptions: cobra.CompletionOptions{
// 				DisableDefaultCmd: true,
// 			},
// 		}

// 		// TODO yuck
// 		command := args[0]
// 		shell := args[1]
// 		args[0] = "_carapace"
// 		args[1] = "export"
// 		os.Args[1] = "_carapace"
// 		os.Args[2] = "export"
// 		os.Setenv("CARAPACE_LENIENT", "1")

// 		carapace.Gen(cmd).PositionalAnyCompletion(
// 			carapace.ActionCallback(func(c carapace.Context) carapace.Action {
// 				batch := carapace.Batch()
// 				specPath, err := completers.SpecPath(command)
// 				if err != nil {
// 					batch = append(batch, carapace.ActionImport([]byte(invokeCompleter(command))))
// 				} else {
// 					out, err := specCompletion(specPath, args[1:]...)
// 					if err != nil {
// 						return carapace.ActionMessage(err.Error())
// 					}

// 					batch = append(batch, carapace.ActionImport([]byte(out)))
// 				}

// 				batch = append(batch, overlayCompletion(overlayPath, args[1:]...))
// 				return batch.ToA()
// 			}),
// 		)

// 		cmd.SetArgs(append([]string{"_carapace", shell}, args[2:]...))
// 		cmd.Execute()
// 	} else {
// 		if specPath, err := completers.SpecPath(args[0]); err == nil {
// 			out, err := specCompletion(specPath, args[1:]...)
// 			if err != nil {
// 				fmt.Fprintln(cmd.ErrOrStderr(), err.Error())
// 				return
// 			}

// 			// TODO revert the patching from specCompletion to use the integrated version for overlay to work (should move this somewhere else - best in specCompletion)
// 			// TODO only patch completion script
// 			// TODO this isn't correct anymore
// 			out = strings.ReplaceAll(out, fmt.Sprintf("--spec '%v'", specPath), args[0])
// 			out = func() string {
// 				var (
// 					s, old, new string = out, fmt.Sprintf("'--spec', '%v'", specPath), fmt.Sprintf("'%v'", args[0])
// 					n           int    = -1
// 				)
// 				if old == new || n == 0 {
// 					return s
// 				}
// 				if m := strings.Count(s, old); m == 0 {
// 					return s
// 				} else if n < 0 || m < n {
// 					n = m
// 				}
// 				var b strings. // Apply replacements to buffer.
// 						Builder
// 				b.Grow(len(s) + n*(len(new)-len(old)))
// 				start := 0
// 				if len(old) > 0 {
// 					for range n {
// 						j := start + strings.Index(s[start:], old)
// 						b.WriteString(s[start:j])
// 						b.WriteString(new)
// 						start = j + len(old)
// 					}
// 				} else {
// 					b.WriteString(new)
// 					for range n - 1 {
// 						_, wid := utf8.DecodeRuneInString(s[start:])
// 						j := start + wid
// 						b.WriteString(s[start:j])
// 						b.WriteString(new)
// 						start = j
// 					}
// 				}
// 				b.WriteString(s[start:])
// 				return b.String()
// 			}() // xonsh callback
// 			fmt.Fprint(cmd.OutOrStdout(), out)
// 		} else if _, err := completers.Lookup(args[0]); err == nil {
// 			fmt.Print(invokeCompleter(args[0]))
// 		} else {
// 			if _, ok := bridges.Bridges()[args[0]]; ok {
// 				_bridgeCmd := &cobra.Command{
// 					Use:                args[0],
// 					DisableFlagParsing: true,
// 				}

// 				carapace.Gen(_bridgeCmd).PositionalAnyCompletion(
// 					bridge.ActionBridges(args[0]),
// 				)
// 				carapace.Gen(_bridgeCmd).Standalone()

// 				out, err := cmdCompletion(_bridgeCmd, args[1:]...)
// 				if err != nil {
// 					fmt.Fprintln(cmd.ErrOrStderr(), err.Error())
// 					return
// 				}
// 				fmt.Fprint(cmd.OutOrStdout(), out)
// 			}
// 			return
// 		}
// 	}
// }

func init() {
	carapace.Gen(invokeCmd).Standalone()
	invokeCmd.Flags().SetInterspersed(false)

	carapace.Gen(invokeCmd).PositionalCompletion(
		carapacebin.ActionCompleters(),
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

func invokeCompleter(nameVariant string) string {
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

	os.Args[1] = "_carapace"
	if completer, err := completers.Lookup(nameVariant); err == nil { // TODO handle error
		switch completer.Execute {
		case nil:
			// TODO bridges have no Execute set yet (panic)
			carapace.LOG.Printf("Completer.Execute of %q is nil", nameVariant)
		default:
			completer.Execute() // TODO Execute should handle the stdout wrapping
		}
	}

	w.Close()
	out := <-outC
	os.Stdout = old

	executable := uid.Executable()
	patched := strings.ReplaceAll(string(out), fmt.Sprintf("%v _carapace", executable), fmt.Sprintf("%v %v", executable, nameVariant))      // general callback
	patched = strings.ReplaceAll(patched, fmt.Sprintf("'%v', '_carapace'", executable), fmt.Sprintf("'%v', '%v'", executable, nameVariant)) // xonsh callback
	return patched
}

// TODO use specCompletion and extract common code with invokeCompleter
// func cmdCompletion(cmd *cobra.Command, args ...string) (string, error) {
// 	old := os.Stdout
// 	r, w, _ := os.Pipe()
// 	os.Stdout = w

// 	outC := make(chan string)
// 	// copy the output in a separate goroutine so printing can't block indefinitely
// 	go func() {
// 		var buf bytes.Buffer
// 		io.Copy(&buf, r)
// 		outC <- buf.String()
// 	}()

// 	a := []string{"_carapace"}
// 	a = append(a, args...)
// 	cmd.SetArgs(a)
// 	cmd.Execute()

// 	w.Close()
// 	out := <-outC
// 	os.Stdout = old

// 	executable, err := os.Executable()
// 	if err != nil {
// 		return "", err
// 	}

// 	executableName := filepath.Base(executable)
// 	patched := strings.ReplaceAll(string(out), fmt.Sprintf("%v _carapace", executableName), fmt.Sprintf("%v %v", executableName, cmd.Name()))      // general callback
// 	patched = strings.ReplaceAll(patched, fmt.Sprintf("'%v', '_carapace'", executableName), fmt.Sprintf("'%v', '%v'", executableName, cmd.Name())) // xonsh callback
// 	return patched, nil
// }
