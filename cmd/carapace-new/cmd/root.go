package cmd

//go:generate go run ../../carapace-generate/gen.go

import (
	"fmt"
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/cmd/carapace-new/cmd/completers"
	"github.com/rsteube/carapace-bin/pkg/actions"
	"github.com/rsteube/carapace-bin/pkg/util/embed"
	spec "github.com/rsteube/carapace-spec"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/rsteube/carapace/pkg/xdg"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "carapace-new [flags] [COMPLETER] [bash|elvish|fish|nushell|oil|powershell|tcsh|xonsh|zsh]",
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
}

func suppressErr(f func() (string, error)) string { s, _ := f(); return s }

func Execute(version string) error {
	rootCmd.Version = version
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		ActionCompleters(),
	)

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
			completerCmd := &cobra.Command{
				Use: args[0],
				Run: func(cmd *cobra.Command, args []string) {
					// TODO patch os.Args and so on
					os.Args[0] = "_carapace"
					executeCompleter(args[0])
				},
			}
			cmd.AddCommand(completerCmd)
		}
	})

	embed.SubcommandsAsFlags(rootCmd,
		listCmd,
		macrosCmd,
		runCmd,
		scrapeCmd,
		styleCmd,
	)

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
