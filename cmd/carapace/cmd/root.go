package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/lazyinit"
	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/shim"
	"github.com/carapace-sh/carapace-bin/pkg/actions"
	carapacebin "github.com/carapace-sh/carapace-bin/pkg/actions/tools/carapace"
	spec "github.com/carapace-sh/carapace-spec"
	"github.com/carapace-sh/carapace/pkg/ps"
	"github.com/carapace-sh/carapace/pkg/xdg"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "carapace [flags] [COMPLETER] [bash|elvish|fish|nushell|oil|powershell|tcsh|xonsh|zsh]",
	Short: "A multi-shell completion binary",
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

  Bridges:
    set-env CARAPACE_BRIDGES 'zsh,fish,bash,inshellisense'

  Shell parameter is optional and if left out carapace will try to detect it by parent process name.
  Some completions are cached at [%v/carapace].
  Config is written to [%v/carapace].
  Specs are loaded from [%v/carapace/specs].
  `, suppressErr(xdg.UserCacheDir), suppressErr(xdg.UserConfigDir), suppressErr(xdg.UserConfigDir)),
	Args:               cobra.MinimumNArgs(1),
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		// since flag parsing is disabled do this manually
		switch args[0] {
		case "--choice":
			choiceCmd.SetArgs(args[1:])
			choiceCmd.Execute()
		case "--clear-cache":
			clearCacheCmd.SetArgs(args[1:])
			clearCacheCmd.Execute()
		case "--condition":
			conditionCmd.SetArgs(args[1:])
			conditionCmd.Execute()
		case "--detect":
			detectCmd.SetArgs(args[1:])
			detectCmd.Execute()
		case "--diff":
			diffCmd.SetArgs(args[1:])
			diffCmd.Execute()
		case "--edit":
			editCmd.SetArgs(args[1:])
			editCmd.Execute()
		case "--macro":
			macroCmd.SetArgs(args[1:])
			macroCmd.Execute()
		case "-h", "--help":
			cmd.Help()
		case "-v", "--version":
			fmt.Println("carapace-bin " + cmd.Version)
		case "--list":
			listCmd.SetArgs(args[1:])
			listCmd.Execute()
		case "--run":
			runCmd.SetArgs(args[1:])
			runCmd.Execute()
		case "--selfupdate":
			selfupdateCmd.SetArgs(args[1:])
			selfupdateCmd.Execute()
		case "--schema":
			schemaCmd.SetArgs(args[1:])
			schemaCmd.Execute()
		case "--codegen":
			codegenCmd.SetArgs(args[1:])
			codegenCmd.Execute()
		case "--style":
			styleCmd.SetArgs(args[1:])
			styleCmd.Execute()
		default:
			invokeCmd.SetArgs(args)
			invokeCmd.Execute()
		}
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func suppressErr(f func() (string, error)) string { s, _ := f(); return s }

func createOverlayDir() error {
	configDir, err := xdg.UserConfigDir()
	if err != nil {
		return err
	}
	return os.MkdirAll(fmt.Sprintf("%v/carapace/overlays", configDir), os.ModePerm)
}

func Execute(version string) error {
	rootCmd.Version = version

	if len(os.Args) > 1 {
		if os.Args[1] == "carapace" {
			os.Args[1] = "_carapace"
		} else if len(os.Args) < 4 && os.Args[1] == "_carapace" {
			shell := ps.DetermineShell()
			if len(os.Args) > 2 {
				shell = os.Args[2]
			}

			if err := shim.Update(); err != nil {
				carapace.LOG.Printf("failed to update shims: %v", err.Error())
			}

			if err := createOverlayDir(); err != nil { // TODO do this only if needed
				carapace.LOG.Printf("failed to create overlay directory: %v", err.Error())
			}

			switch shell {
			case "bash",
				"bash-ble",
				"cmd-clink",
				"elvish",
				"fish",
				"nushell",
				"oil",
				"powershell",
				"tcsh",
				"xonsh",
				"zsh":
				fmt.Println(lazyinit.Snippet(shell)) // TODO maybe just return an error for unknown shell
			default:
				// TODO
				// println("could not determine shell")
				return rootCmd.Execute() // TODO verify
			}
			return nil
		}
	}
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().Bool("choice", false, "list or edit choices")
	rootCmd.Flags().Bool("clear-cache", false, "clear caches")
	rootCmd.Flags().Bool("codegen", false, "generate code for spec file")
	rootCmd.Flags().Bool("condition", false, "list or execute condition")
	rootCmd.Flags().Bool("detect", false, "detect bridge by invoking command")
	rootCmd.Flags().Bool("diff", false, "diff completion")
	rootCmd.Flags().Bool("edit", false, "edit files")
	rootCmd.Flags().BoolP("help", "h", false, "help for carapace")
	rootCmd.Flags().Bool("list", false, "list completers")
	rootCmd.Flags().Bool("macro", false, "list or execute macros")
	rootCmd.Flags().Bool("run", false, "run spec")
	rootCmd.Flags().Bool("selfupdate", false, "update to nightly/stable")
	rootCmd.Flags().Bool("schema", false, "json schema for spec files")
	rootCmd.Flags().Bool("style", false, "set style")
	rootCmd.Flags().BoolP("version", "v", false, "version for carapace")

	rootCmd.MarkFlagsMutuallyExclusive(
		"choice",
		"clear-cache",
		"codegen",
		"condition",
		"detect",
		"diff",
		"edit",
		"help",
		"list",
		"macro",
		"run",
		"schema",
		"selfupdate",
		"style",
		"version",
	)

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.HasPrefix(c.Value, "-") {
				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(rootCmd.Flags())
				return carapace.ActionExecute(cmd)
			}
			return carapacebin.ActionCompleters()
		}),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !strings.HasPrefix(c.Args[0], "-") {
				return carapace.ActionExecute(invokeCmd)
			}
			switch c.Args[0] {
			case "--choice":
				return carapace.ActionExecute(choiceCmd).Shift(1)
			case "--clear-cache":
				return carapace.ActionExecute(clearCacheCmd).Shift(1)
			case "--codegen":
				return carapace.ActionExecute(codegenCmd).Shift(1)
			case "--condition":
				return carapace.ActionExecute(conditionCmd).Shift(1)
			case "--detect":
				return carapace.ActionExecute(detectCmd).Shift(1)
			case "--diff":
				return carapace.ActionExecute(diffCmd).Shift(1)
			case "--edit":
				return carapace.ActionExecute(editCmd).Shift(1)
			case "--help":
				return carapace.ActionValues()
			case "--list":
				return carapace.ActionExecute(listCmd).Shift(1).Usage("list")
			case "--macro":
				return carapace.ActionExecute(macroCmd).Shift(1)
			case "--run":
				return carapace.ActionExecute(runCmd).Shift(1)
			case "--selfupdate":
				return carapace.ActionExecute(selfupdateCmd).Shift(1)
			case "--schema":
				return carapace.ActionExecute(schemaCmd).Shift(1)
			case "--style":
				return carapace.ActionExecute(styleCmd).Shift(1)
			case "-v", "--version":
				return carapace.ActionValues()
			default:
				return carapace.ActionValues()
			}
		}),
	)

	for m, f := range actions.Macros {
		spec.AddMacro(m, f) // TODO just provide a reference for lookup to skip adding them all the time
	}
	spec.Register(rootCmd)
}
