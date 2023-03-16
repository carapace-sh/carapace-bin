package cmd

//go:generate go run ../../generate/gen.go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/completers"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/shim"
	spec "github.com/rsteube/carapace-spec"
	"github.com/rsteube/carapace/pkg/ps"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/rsteube/carapace/pkg/xdg"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "carapace [flags] [COMPLETER] [bash|elvish|fish|nushell|oil|powershell|tcsh|xonsh|zsh]",
	Short: "multi-shell multi-command argument completer",
	Example: fmt.Sprintf(`  All completers and specs:
    bash:       source <(carapace _carapace bash)
    elvish:     eval (carapace _carapace elvish | slurp)
    fish:       carapace _carapace fish | source
    nushell:    carapace _carapace nushell
    oil:        source <(carapace _carapace oil)
    powershell: carapace _carapace powershell | Out-String | Invoke-Expression
    tcsh:       eval `+"`"+`carapace _carapace tcsh`+"`"+`
    xonsh:      exec($(carapace _carapace xonsh))
    zsh:        source <(carapace _carapace zsh)

  Single completer:
    bash:       source <(carapace chmod bash)
    elvish:     eval (carapace chmod elvish | slurp)
    fish:       carapace chmod fish | source
    nushell:    carapace chmod nushell
    oil:        source <(carapace chmod oil)
    powershell: carapace chmod powershell | Out-String | Invoke-Expression
    tcsh:       eval `+"`"+`carapace _chmod tcsh`+"`"+`
    xonsh:      exec($(carapace chmod xonsh))
    zsh:        source <(carapace chmod zsh)

  Bridge completion:
    bash:       source <(carapace --bridge vault/posener)
    elvish:     eval (carapace --bridge vault/posener|slurp)
    fish:       carapace --bridge vault/posener | source
    nushell:    carapace --bridge vault/posener
    oil:        source <(carapace --bridge vault/posener)
    powershell: carapace --bridge vault/posener | Out-String | Invoke-Expression
    tcsh:       eval `+"`"+`carapace --bridge vault/posener`+"`"+`
    xonsh:      exec($(carapace --bridge vault/posener))
    zsh:        source <(carapace --bridge vault/posener)
  
  Spec completion:
    bash:       source <(carapace --spec example.yaml)
    elvish:     eval (carapace --spec example.yaml|slurp)
    fish:       carapace --spec example.yaml | source
    oil:        source <(carapace --spec example.yaml)
    nushell:    carapace --spec example.yaml
    powershell: carapace --spec example.yaml | Out-String | Invoke-Expression
    tcsh:       eval `+"`"+`carapace --spec example.yaml`+"`"+`
    xonsh:      exec($(carapace --spec example.yaml))
    zsh:        source <(carapace --spec example.yaml)

  Style:
    set:        carapace --style 'carapace.Value=bold,magenta'
    clear:      carapace --style 'carapace.Description='

  Shell parameter is optional and if left out carapace will try to detect it by parent process name.
  Some completions are cached at [%v/carapace].
  Config is written to [%v/carapace].
  Specs are loaded from [%v/carapace/specs].
  `, suppressErr(xdg.UserCacheDir), suppressErr(xdg.UserConfigDir), suppressErr(xdg.UserConfigDir)),
	Args:      cobra.MinimumNArgs(1),
	ValidArgs: completers.Names(),
	Run: func(cmd *cobra.Command, args []string) {
		// since flag parsing is disabled do this manually
		switch args[0] {
		case "--bridge":
			if len(args) > 1 {
				// TODO support multiple (comma separated)
				if splitted := strings.SplitN(args[1], "/", 2); len(splitted) == 2 {
					bridgeCompletion(splitted[0], splitted[1], args[2:]...)
				}
			}
		case "--spec":
			if len(args) > 1 {
				out, err := specCompletion(args[1], args[2:]...)
				if err != nil {
					fmt.Fprintln(cmd.ErrOrStderr(), err.Error())
					return
				}
				fmt.Fprintln(cmd.OutOrStdout(), out)
			}
		case "--macros":
			if len(args) > 1 {
				printMacro(args[1])
			} else {
				printMacros()
			}
		case "-h":
			cmd.Help()
		case "--help":
			cmd.Help()
		case "-v":
			println(cmd.Version)
		case "--version":
			println(cmd.Version)
		case "--list":
			printCompleters()
		case "--list=json":
			printCompletersJson()
		case "--run":
			_, spec, err := loadSpec(args[1])
			if err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), err.Error())
				os.Exit(1)
			}

			cmd := spec.ToCobra()
			cmd.SetArgs(args[2:])
			cmd.Execute() // TODO handle error?
		case "--schema":
			if schema, err := spec.Schema(); err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), err.Error()) // TODO fail / exit 1 ?
			} else {
				fmt.Fprintln(cmd.OutOrStdout(), schema)
			}
		case "--scrape":
			if len(args) > 1 {
				scrape(args[1])
			}
		case "--style":
			if len(args) > 1 {
				if err := setStyle(args[1]); err != nil {
					fmt.Fprintln(cmd.ErrOrStderr(), err.Error())
				}
			}
		case "_carapace":
			shell := ps.DetermineShell()
			if len(args) > 1 {
				shell = args[1]
			}
			if len(args) <= 2 {
				if err := shim.Update(); err != nil {
					fmt.Fprintln(cmd.ErrOrStderr(), err.Error()) // TODO fail / exit 1 ?
				}

				if err := updateSchema(); err != nil { // TODO do this only if needed
					fmt.Fprintln(cmd.ErrOrStderr(), err.Error())
				}

				if err := createOverlayDir(); err != nil { // TODO do this only if needed
					fmt.Fprintln(cmd.ErrOrStderr(), err.Error())
				}
			}
			switch shell {
			case "bash":
				fmt.Fprintln(cmd.OutOrStdout(), bash_lazy(completers.Names()))
			case "bash-ble":
				fmt.Fprintln(cmd.OutOrStdout(), bash_ble_lazy(completers.Names()))
			case "elvish":
				fmt.Fprintln(cmd.OutOrStdout(), elvish_lazy(completers.Names()))
			case "fish":
				fmt.Fprintln(cmd.OutOrStdout(), fish_lazy(completers.Names()))
			case "nushell":
				fmt.Fprintln(cmd.OutOrStdout(), nushell_lazy(completers.Names()))
			case "oil":
				fmt.Fprintln(cmd.OutOrStdout(), oil_lazy(completers.Names()))
			case "powershell":
				fmt.Fprintln(cmd.OutOrStdout(), powershell_lazy(completers.Names()))
			case "tcsh":
				fmt.Fprintln(cmd.OutOrStdout(), tcsh_lazy(completers.Names()))
			case "xonsh":
				fmt.Fprintln(cmd.OutOrStdout(), xonsh_lazy(completers.Names()))
			case "zsh":
				fmt.Fprintln(cmd.OutOrStdout(), zsh_lazy(completers.Names()))
			default:
				fmt.Fprintln(os.Stderr, "could not determine shell")
			}
		default:
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
					out = strings.Replace(out, fmt.Sprintf("--spec '%v'", specPath), args[0], 1)
					out = strings.Replace(out, fmt.Sprintf("'--spec', '%v'", specPath), fmt.Sprintf("'%v'", args[0]), 1) // xonsh callback
					fmt.Fprintln(cmd.OutOrStdout(), out)
				}
				fmt.Println(invokeCompleter(args[0]))
			}
		}

	},
	FParseErrWhitelist: cobra.FParseErrWhitelist{
		UnknownFlags: true,
	},
	DisableFlagParsing: true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func suppressErr(f func() (string, error)) string { s, _ := f(); return s }

func printCompleters() {
	maxlen := 0
	for _, name := range completers.Names() {
		if len := len(name); len > maxlen {
			maxlen = len
		}
	}

	for _, name := range completers.Names() {
		fmt.Printf("%-"+strconv.Itoa(maxlen)+"v %v\n", name, completers.Description(name))
	}
}

func createOverlayDir() error {
	configDir, err := xdg.UserConfigDir()
	if err != nil {
		return err
	}
	return os.MkdirAll(fmt.Sprintf("%v/carapace/overlays", configDir), os.ModePerm)
}

func overlayPath(command string) (string, error) {
	configDir, err := xdg.UserConfigDir()
	if err != nil {
		return "", err
	}

	overlayPath := fmt.Sprintf("%v/carapace/overlays/%v.yaml", configDir, command)
	if _, err = os.Stat(overlayPath); err != nil {
		return "", err
	}
	return overlayPath, nil
}

func overlayCompletion(overlayPath string, args ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args[0] = "export" // TODO length check
		out, err := specCompletion(overlayPath, args...)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return carapace.ActionImport([]byte(out))
	})
}

func printCompletersJson() {
	// TODO move to completers package
	type _completer struct {
		Name        string
		Description string
		Spec        string `json:",omitempty"`
	}

	_completers := make([]_completer, 0)
	for _, name := range completers.Names() {
		specPath, _ := completers.SpecPath(name) // TODO handle error (log?)
		_completers = append(_completers, _completer{
			Name:        name,
			Description: completers.Description(name),
			Spec:        specPath,
		})
	}
	if m, err := json.Marshal(_completers); err == nil { // TODO handle error (log?)
		fmt.Println(string(m))
	}
}

func printMacros() {
	maxlen := 0
	names := make([]string, 0)
	for name := range macros {
		names = append(names, name)
		if len := len(name); len > maxlen {
			maxlen = len
		}
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%-"+strconv.Itoa(maxlen)+"v %v\n", name, macroDescriptions[name])
	}
}

func printMacro(name string) {
	if m, ok := macros[name]; ok {
		path := strings.Replace(name, ".", "/", -1)
		signature := ""
		if s := m.Signature(); s != "" {
			signature = fmt.Sprintf("(%v)", s)
		}

		fmt.Printf(`signature:   $_%v%v
description: %v
reference:   https://pkg.go.dev/github.com/rsteube/carapace-bin/pkg/actions/%v#Action%v
`, name, signature, macroDescriptions[name], filepath.Dir(path), filepath.Base(path))
	}
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

	executable, err := os.Executable()
	if err != nil {
		panic(err.Error()) // TODO exit with error message
	}
	executableName := filepath.Base(executable)
	patched := strings.Replace(string(out), fmt.Sprintf("%v _carapace", executableName), fmt.Sprintf("%v %v", executableName, completer), -1)      // general callback
	patched = strings.Replace(patched, fmt.Sprintf("'%v', '_carapace'", executableName), fmt.Sprintf("'%v', '%v'", executableName, completer), -1) // xonsh callback
	return patched

}

func setStyle(s string) error {
	if splitted := strings.SplitN(s, "=", 2); len(splitted) == 2 {
		return style.Set(splitted[0], splitted[1])
	}
	return fmt.Errorf("invalid format: '%v'", s)
}

func updateSchema() error {
	confDir, err := xdg.UserConfigDir()
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%v/carapace/schema.json", confDir)

	infoSchema, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	executable, err := os.Executable()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	infoCarapace, err := os.Stat(executable)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if infoSchema == nil || !infoSchema.ModTime().Equal(infoCarapace.ModTime()) {
		schema, err := spec.Schema()
		if err != nil && !os.IsNotExist(err) {
			return err
		}

		carapace.LOG.Printf("writing json schema to %#v", path)
		if err := os.WriteFile(path, []byte(schema), os.ModePerm); err != nil {
			return err
		}

		if err := os.Chtimes(path, time.Now(), infoCarapace.ModTime()); err != nil {
			return err
		}
	}
	return nil
}

func Execute(version string) error {
	rootCmd.Version = version
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().String("bridge", "", "bridge completion")
	rootCmd.Flags().Bool("list", false, "list completers")
	rootCmd.Flags().String("run", "", "run spec")
	rootCmd.Flags().String("scrape", "", "scrape spec to go code")
	rootCmd.Flags().String("spec", "", "spec completion")
	rootCmd.Flags().String("style", "", "set style")

	for m, f := range macros {
		spec.AddMacro(m, f)
	}
}
