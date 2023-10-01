package cmd

//go:generate go run ../../carapace-generate/gen.go

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/cmd/carapace-new/cmd/completers"
	"github.com/rsteube/carapace-bin/pkg/actions"
	spec "github.com/rsteube/carapace-spec"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/rsteube/carapace/pkg/xdg"
	"github.com/spf13/cobra"
)

var subcommands = map[string]*cobra.Command{
	"--list":   listCmd,
	"--macros": macrosCmd,
	"--run":    runCmd,
	"--scrape": scrapeCmd,
	"--style":  styleCmd,
}

var rootCmd = &cobra.Command{
	Use:   "carapace-new [COMPLETER] [bash|elvish|fish|nushell|oil|powershell|tcsh|xonsh|zsh]",
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

  Single completer or spec:
    bash:       source <(carapace chmod bash)
    elvish:     eval (carapace chmod elvish | slurp)
    fish:       carapace chmod fish | source
    nushell:    carapace chmod nushell
    oil:        source <(carapace chmod oil)
    powershell: carapace chmod powershell | Out-String | Invoke-Expression
    tcsh:       eval `+"`"+`carapace chmod tcsh`+"`"+`
    xonsh:      exec($(carapace chmod xonsh))
    zsh:        source <(carapace chmod zsh)
  
  Style:
    set:        carapace --style 'carapace.Value=bold,magenta'
    clear:      carapace --style 'carapace.Description='

  Shell parameter is optional and if left out carapace will try to detect it by parent process name.
  Some completions are cached at [%v/carapace].
  Config is written to [%v/carapace].
  Specs are loaded from [%v/carapace/specs].
  `, suppressErr(xdg.UserCacheDir), suppressErr(xdg.UserConfigDir), suppressErr(xdg.UserConfigDir)),
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if subCmd := subcommands[args[0]]; subCmd != nil {
			subCmd.SetArgs(args[1:])
			return subCmd.Execute()
		}
		return errors.New("unknown") // TODO
	},
}

func suppressErr(f func() (string, error)) string { s, _ := f(); return s }

func Execute(version string) error {
	switch {
	case len(os.Args) > 1 && os.Args[1] != "_carapace":
		if subCmd := subcommands[os.Args[1]]; subCmd != nil {
			subCmd.SetArgs(os.Args[2:])
			return subCmd.Execute()
		}

		fmt.Print(invokeCompleter(os.Args[1]))
		return nil

	default:
		rootCmd.Version = version
		return rootCmd.Execute()
	}
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "help for carapace")
	rootCmd.Flags().BoolP("version", "v", false, "version for carapace")
	for name, cmd := range subcommands {
		rootCmd.Flags().Bool(strings.TrimPrefix(name, "--"), false, cmd.Short)
	}

	carapace.Gen(rootCmd).PositionalCompletion(
		ActionCompleters(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if subCmd := subcommands[c.Args[0]]; subCmd != nil {
				return carapace.ActionExecute(subCmd).Shift(1)
			}

			os.Args = []string{os.Args[1], "_carapace", "export", os.Args[1], "_carapace"}
			// if len(os.Args) > 4 {
			os.Args = append(os.Args, os.Args[5:]...)
			// }
			carapace.LOG.Printf("%#v", os.Args)
			// panic(fmt.Sprintf("%#v", os.Args))
			return carapace.ActionImport([]byte(invokeCompleter(c.Args[0])))
		}),
	)

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			cmd.DisableFlagParsing = true
		}
	})

	for m, f := range actions.MacroMap {
		spec.AddMacro(m, f) // TODO only do this when needed
	}
}

func ActionCompleters() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		vals := make([]string, 0)

		for _, name := range completers.Names() {
			s := style.Default
			if _, err := completers.SpecPath(name); err == nil {
				s = style.Blue
			}
			if _, err := completers.OverlayPath(name); err == nil {
				s = style.Of(s, style.Underlined)
			}
			vals = append(vals, name, completers.Description(name), s)
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}

func invokeCompleter(completer string) string { // TODO patching should only be necessary for the init script
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
	completers.ExecuteCompleter(completer)

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
