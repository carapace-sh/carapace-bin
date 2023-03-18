package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	spec "github.com/rsteube/carapace-spec"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func loadSpec(path string) (string, *spec.Command, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", nil, err
	}

	content, err := os.ReadFile(abs)
	if err != nil {
		return "", nil, err
	}

	var specCmd spec.Command
	if err := yaml.Unmarshal(content, &specCmd); err != nil {
		return "", nil, err
	}
	return abs, &specCmd, nil
}

func addAliasCompletion(cmd *cobra.Command, specCommand spec.Command) {
	if specCommand.Run != "" &&
		len(specCommand.Flags) == 0 &&
		len(specCommand.PersistentFlags) == 0 &&
		len(specCommand.Completion.Positional) == 0 &&
		len(specCommand.Completion.PositionalAny) == 0 &&
		len(specCommand.Completion.Dash) == 0 &&
		len(specCommand.Completion.DashAny) == 0 {

		cmd.DisableFlagParsing = true
		carapace.Gen(cmd).PositionalAnyCompletion(
			carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				switch {
				case regexp.MustCompile(`^\[.*\]$`).MatchString(string(specCommand.Run)):
					var mArgs []string
					if err := yaml.Unmarshal([]byte(specCommand.Run), &mArgs); err != nil {
						return carapace.ActionMessage(err.Error())
					}
					if len(mArgs) == 0 {
						return carapace.ActionMessage("empty alias: %#v", specCommand.Run)
					}

					execArgs := []string{mArgs[0], "export", mArgs[0]}
					execArgs = append(execArgs, mArgs[1:]...)
					execArgs = append(execArgs, c.Args...)
					execArgs = append(execArgs, c.CallbackValue)
					return carapace.ActionExecCommand("carapace", execArgs...)(func(output []byte) carapace.Action {
						return carapace.ActionImport(output)
					})

				default:
					return carapace.ActionValues()
				}
			}),
		)
	}

	for _, subCommand := range specCommand.Commands {
		if subCmd, _, err := cmd.Find([]string{subCommand.Name}); err == nil {
			addAliasCompletion(subCmd, subCommand)
		}
	}
}

func scrape(path string) {
	// TODO yuck - all this needs some cleanup
	if _, spec, err := loadSpec(path); err == nil {
		spec.Scrape()
	}
}

func specCompletion(path string, args ...string) (string, error) {
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

	abs, spec, err := loadSpec(path)
	if err != nil {
		return "", err
	}

	cmd := spec.ToCobra()
	addAliasCompletion(cmd, *spec) // TODO put this somewhere else

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
	patched := strings.Replace(string(out), fmt.Sprintf("%v _carapace", executableName), fmt.Sprintf("%v --spec '%v'", executableName, abs), -1)       // general callback
	patched = strings.Replace(patched, fmt.Sprintf("'%v', '_carapace'", executableName), fmt.Sprintf("'%v', '--spec', '%v'", executableName, abs), -1) // xonsh callback
	return patched, nil
}
