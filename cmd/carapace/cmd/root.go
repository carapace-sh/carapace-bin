package cmd

//go:generate go run ../../carapace-generate/gen.go

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/action"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/completers"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/completers/bridges"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/lazyinit"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/shim"
	"github.com/rsteube/carapace-bin/internal/env"
	"github.com/rsteube/carapace-bin/pkg/actions"
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
	Args:               cobra.MinimumNArgs(1),
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		// since flag parsing is disabled do this manually
		switch args[0] {
		case "--condition":
			conditionCmd.SetArgs(args[1:])
			conditionCmd.Execute()
		case "--macro":
			macroCmd.SetArgs(args[1:])
			macroCmd.Execute()
		case "-h", "--help":
			cmd.Help()
		case "-v", "--version":
			println(cmd.Version)
		case "--list":
			listCmd.SetArgs(args[1:])
			listCmd.Execute()
		case "--run":
			runCmd.SetArgs(args[1:])
			runCmd.Execute()
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
				println(err.Error()) // TODO fail / exit 1 ?
			}

			if err := createOverlayDir(); err != nil { // TODO do this only if needed
				println(err.Error())
			}

			uniqueNames := make(map[string]bool)
			// TODO enable/order by CARAPACE_BRIDGE
			for _, name := range completers.Names() {
				uniqueNames[name] = true
			}
			for _, b := range env.Bridges() {
				switch b {
				case "fish":
					for _, name := range bridges.Fish() {
						uniqueNames[name] = true
					}
				case "zsh":
					for _, name := range bridges.Zsh() {
						uniqueNames[name] = true
					}
				}
			}

			completerNames := make([]string, 0)
			for name := range uniqueNames {
				completerNames = append(completerNames, name)
			}
			sort.Strings(completerNames)

			if os.Getenv("CARAPACE_ENV") == "0" {
				filtered := make([]string, 0, len(completerNames))
				for _, name := range completerNames {
					switch name {
					case "get-env", "set-env", "unset-env":
					default:
						filtered = append(filtered, name)

					}
				}
				completerNames = filtered
			}

			switch shell {
			case "bash":
				fmt.Println(lazyinit.Bash(completerNames))
			case "bash-ble":
				fmt.Println(lazyinit.BashBle(completerNames))
			case "elvish":
				fmt.Println(lazyinit.Elvish(completerNames))
			case "fish":
				fmt.Println(lazyinit.Fish(completerNames))
			case "nushell":
				fmt.Println(lazyinit.Nushell(completerNames))
			case "oil":
				fmt.Println(lazyinit.Oil(completerNames))
			case "powershell":
				fmt.Println(lazyinit.Powershell(completerNames))
			case "tcsh":
				fmt.Println(lazyinit.Tcsh(completerNames))
			case "xonsh":
				fmt.Println(lazyinit.Xonsh(completerNames))
			case "zsh":
				fmt.Println(lazyinit.Zsh(completerNames))
			default:
				// TODO
				// println("could not determine shell")
				return rootCmd.Execute()
			}
			return nil
		}
	}
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().Bool("codegen", false, "generate code for spec file")
	rootCmd.Flags().Bool("condition", false, "list or execute condition")
	rootCmd.Flags().BoolP("help", "h", false, "help for carapace")
	rootCmd.Flags().Bool("list", false, "list completers")
	rootCmd.Flags().Bool("macro", false, "list or execute macros")
	rootCmd.Flags().Bool("run", false, "run spec")
	rootCmd.Flags().Bool("schema", false, "json schema for spec files")
	rootCmd.Flags().Bool("style", false, "set style")
	rootCmd.Flags().BoolP("version", "v", false, "version for carapace")

	rootCmd.MarkFlagsMutuallyExclusive(
		"codegen",
		"condition",
		"help",
		"list",
		"macro",
		"run",
		"schema",
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
			return carapace.ActionMultiPartsN("/", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return action.ActionCompleters()
				default:
					return carapace.ActionStyledValues(
						"bash", "#d35673",
						"carapace", style.Default,
						"carapace-bin", style.Default,
						"clap", style.Default,
						"click", style.Default,
						"cobra", style.Default,
						"complete", style.Default,
						"fish", "#7ea8fc",
						"inshellisense", style.Default,
						"kingpin", style.Default,
						"powershell", style.Default,
						"urfavecli", style.Default,
						"yargs", style.Default,
						"zsh", "#efda53",
					)
				}
			})
		}),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !strings.HasPrefix(c.Args[0], "-") {
				return carapace.ActionExecute(invokeCmd)
			}
			switch c.Args[0] {
			case "--codegen":
				return carapace.ActionExecute(codegenCmd).Shift(1)
			case "--condition":
				return carapace.ActionExecute(conditionCmd).Shift(1)
			case "--help":
				return carapace.ActionValues()
			case "--list":
				return carapace.ActionExecute(listCmd).Shift(1).Usage("list")
			case "--macro":
				return carapace.ActionExecute(macroCmd).Shift(1)
			case "--run":
				return carapace.ActionExecute(runCmd).Shift(1)
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

	for m, f := range actions.MacroMap {
		spec.AddMacro(m, f)
	}
	spec.Register(rootCmd)
}
