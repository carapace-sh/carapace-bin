package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/action"
	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/completers"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace-bridge/pkg/bridges"
	"github.com/carapace-sh/carapace/pkg/uid"
	"github.com/spf13/cobra"
)

// TODO
// func hasUserSpec(s string) bool {
// 	_, err := completers.SpecPath(s)
// 	return err != nil
// }

var invokeCmd = &cobra.Command{
	Use:   "invoke",
	Short: "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO newInvoke(cmd, args)
		oldInvoke(cmd, args)
	},
}

// func newInvoke(cmd *cobra.Command, args []string) {
// 	command, _bridge, explicit := strings.Cut(args[0], "/")
// 	args = args[1:]

// 	batch := carapace.Batch()
// 	switch {
// 	case explicit:
// 		switch _bridge {
// 		// TODO use given bridge
// 		}
// 	case hasUserSpec(command): // user spec
// 		if path, err := completers.SpecPath(command); err == nil { // TODO optimistic double check
// 			batch = append(batch, spec.ActionSpec(path))
// 		}
// 	case false: // TODO system spec (not yet supported)
// 	case slices.Contains(completers.Names(), command): // internal completer
// 		batch = append(batch, carapace.ActionImport([]byte(invokeCompleter(command))))
// 	case false: // TODO bridge current shell (if bridge exists)
// 	default: // implicit bridge
// 		batch = append(batch, bridge.ActionBridges(command))
// 	}

// 	if path, err := overlayPath(command); err == nil {
// 		batch = carapace.Batch(
// 			carapace.ActionCallback(func(c carapace.Context) carapace.Action {
// 				c.Setenv("CARAPACE_LENIENT", "1")
// 				return append(batch, overlayCompletion(path, args...)).Invoke(c).Merge().ToA()
// 			}),
// 			// TODO
// 			// append(batch, overlayCompletion(path, args...)).ToA().
// 			// 	Setenv("CARAPACE_LENIENT", "1"),
// 		)
// 	}

// 	_cmd := &cobra.Command{
// 		DisableFlagParsing: true,
// 	}
// 	_cmd.SetOut(cmd.OutOrStdout())
// 	_cmd.SetErr(cmd.OutOrStderr())
// 	_cmd.SetArgs(append([]string{"_carapace", "export"}, args...))
// 	carapace.Gen(_cmd).Standalone()

// 	carapace.Gen(_cmd).PositionalAnyCompletion(
// 		batch.ToA(),
// 	)
// 	_cmd.Execute() // TODO error handling?
// }

func oldInvoke(cmd *cobra.Command, args []string) {
	if overlayPath, err := overlayPath(args[0]); err == nil && len(args) > 2 { // and arg[1] is a known shell
		cmd := &cobra.Command{
			DisableFlagParsing: true,
			CompletionOptions: cobra.CompletionOptions{
				DisableDefaultCmd: true,
			},
		}

		// TODO yuck
		command := args[0]
		shell := args[1]
		args[0] = "_carapace"
		args[1] = "export"
		os.Args[1] = "_carapace"
		os.Args[2] = "export"
		os.Setenv("CARAPACE_LENIENT", "1")

		carapace.Gen(cmd).PositionalAnyCompletion(
			carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				batch := carapace.Batch()
				specPath, err := completers.SpecPath(command)
				if err != nil {
					batch = append(batch, carapace.ActionImport([]byte(invokeCompleter(command))))
				} else {
					out, err := specCompletion(specPath, args[1:]...)
					if err != nil {
						return carapace.ActionMessage(err.Error())
					}

					batch = append(batch, carapace.ActionImport([]byte(out)))
				}

				batch = append(batch, overlayCompletion(overlayPath, args[1:]...))
				return batch.ToA()
			}),
		)

		cmd.SetArgs(append([]string{"_carapace", shell}, args[2:]...))
		cmd.Execute()
	} else {
		if specPath, err := completers.SpecPath(args[0]); err == nil {
			out, err := specCompletion(specPath, args[1:]...)
			if err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), err.Error())
				return
			}

			// TODO revert the patching from specCompletion to use the integrated version for overlay to work (should move this somewhere else - best in specCompletion)
			// TODO only patch completion script
			// TODO this isn't correct anymore
			out = strings.Replace(out, fmt.Sprintf("--spec '%v'", specPath), args[0], -1)
			out = strings.Replace(out, fmt.Sprintf("'--spec', '%v'", specPath), fmt.Sprintf("'%v'", args[0]), -1) // xonsh callback
			fmt.Fprint(cmd.OutOrStdout(), out)
		} else if slices.Contains(completers.Names(), args[0]) {
			fmt.Print(invokeCompleter(args[0]))
		} else {
			if _, ok := bridges.Bridges()[args[0]]; ok {
				_bridgeCmd := &cobra.Command{
					Use:                args[0],
					DisableFlagParsing: true,
				}

				carapace.Gen(_bridgeCmd).PositionalAnyCompletion(
					bridge.ActionBridges(args[0]),
				)
				carapace.Gen(_bridgeCmd).Standalone()

				out, err := cmdCompletion(_bridgeCmd, args[1:]...)
				if err != nil {
					fmt.Fprintln(cmd.ErrOrStderr(), err.Error())
					return
				}
				fmt.Fprint(cmd.OutOrStdout(), out)
			}
			return
		}
	}
}

func init() {
	carapace.Gen(invokeCmd).Standalone()
	invokeCmd.Flags().SetInterspersed(false)

	carapace.Gen(invokeCmd).PositionalCompletion(
		action.ActionCompleters(action.CompleterOpts{}.Default()),
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
				return carapace.ActionValues(c.Args[0])
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
				return bridge.ActionCarapaceBin(c.Args[0]).Shift(3)
			default:
				return carapace.ActionValues()
			}
		}),
	)
}

func invokeCompleter(completer string) string {
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
	executeCompleter(completer)

	w.Close()
	out := <-outC
	os.Stdout = old

	executable := uid.Executable()
	patched := strings.Replace(string(out), fmt.Sprintf("%v _carapace", executable), fmt.Sprintf("%v %v", executable, completer), -1)      // general callback
	patched = strings.Replace(patched, fmt.Sprintf("'%v', '_carapace'", executable), fmt.Sprintf("'%v', '%v'", executable, completer), -1) // xonsh callback
	return patched
}

// TODO use specCompletion and extract common code with invokeCompleter

func cmdCompletion(cmd *cobra.Command, args ...string) (string, error) {
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

	a := []string{"_carapace"}
	a = append(a, args...)
	cmd.SetArgs(a)
	cmd.Execute()

	w.Close()
	out := <-outC
	os.Stdout = old

	executable, err := os.Executable()
	if err != nil {
		return "", err
	}

	executableName := filepath.Base(executable)
	patched := strings.Replace(string(out), fmt.Sprintf("%v _carapace", executableName), fmt.Sprintf("%v %v", executableName, cmd.Name()), -1)      // general callback
	patched = strings.Replace(patched, fmt.Sprintf("'%v', '_carapace'", executableName), fmt.Sprintf("'%v', '%v'", executableName, cmd.Name()), -1) // xonsh callback
	return patched, nil
}
