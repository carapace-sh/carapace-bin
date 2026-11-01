package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/devenv_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/devenv"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devenv",
	Short: "Fast, Declarative, Reproducible, and Composable Developer Environments",
	Long:  "https://devenv.sh",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringSliceP("clean", "c", nil, "Ignore existing environment variables when entering the shell")
	rootCmd.PersistentFlags().StringP("cores", "u", "", "Number of CPU cores available to each build")
	rootCmd.PersistentFlags().Bool("eval-cache", false, "Enable caching of Nix evaluation results (default)")
	rootCmd.PersistentFlags().String("from", "", "Source for devenv.nix")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print help")
	rootCmd.PersistentFlags().BoolP("impure", "i", false, "Relax the hermeticity of the environment")
	rootCmd.PersistentFlags().StringP("max-jobs", "j", "", "Maximum number of Nix builds to run concurrently")
	rootCmd.PersistentFlags().Bool("nix-debugger", false, "Enter the Nix debugger on failure")
	rootCmd.PersistentFlags().StringSlice("nix-option", nil, "Pass additional options to nix commands")
	rootCmd.PersistentFlags().Bool("no-eval-cache", false, "Disable caching of Nix evaluation results")
	rootCmd.PersistentFlags().Bool("no-impure", false, "Force a hermetic environment, overriding config")
	rootCmd.PersistentFlags().Bool("no-reload", false, "Disable auto-reload when config files change")
	rootCmd.PersistentFlags().Bool("no-tui", false, "Disable the interactive terminal interface")
	rootCmd.PersistentFlags().Bool("offline", false, "Disable substituters and consider all previously downloaded files up-to-date")
	rootCmd.PersistentFlags().StringSliceP("option", "O", nil, "Override configuration options with typed values")
	rootCmd.PersistentFlags().StringSliceP("override-input", "o", nil, "Override inputs in devenv.yaml")
	rootCmd.PersistentFlags().StringArrayP("profile", "P", nil, "Activate one or more profiles defined in devenv.nix")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Silence all logs")
	rootCmd.PersistentFlags().Bool("refresh-eval-cache", false, "Force a refresh of the Nix evaluation cache")
	rootCmd.PersistentFlags().Bool("refresh-task-cache", false, "Force a refresh of the task cache")
	rootCmd.PersistentFlags().Bool("reload", false, "Enable auto-reload when config files change (default)")
	rootCmd.PersistentFlags().String("secretspec-profile", "", "Override the secretspec profile")
	rootCmd.PersistentFlags().String("secretspec-provider", "", "Override the secretspec provider")
	rootCmd.PersistentFlags().String("shell", "", "Shell to use for interactive sessions")
	rootCmd.PersistentFlags().StringP("system", "s", "", "Override the target system")
	rootCmd.PersistentFlags().StringArray("trace-to", nil, "Enable tracing")
	rootCmd.PersistentFlags().String("tui", "", "Enable the interactive terminal interface (default when interactive)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable additional debug logs")

	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	rootCmd.Flag("clean").NoOptDefVal = " "
	rootCmd.Flag("nix-option").Nargs = 2
	rootCmd.Flag("option").Nargs = 2
	rootCmd.Flag("override-input").Nargs = 2

	rootCmd.MarkFlagsMutuallyExclusive("eval-cache", "no-eval-cache")
	rootCmd.MarkFlagsMutuallyExclusive("impure", "no-impure")
	rootCmd.MarkFlagsMutuallyExclusive("reload", "no-reload")
	rootCmd.MarkFlagsMutuallyExclusive("tui", "no-tui")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"clean": env.ActionNames().UniqueList(","),
		"from":  actionSources(),
		"nix-option": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return nix.ActionConfigKeys()
			case 1:
				return nix.ActionConfigValues(c.Parts[0])
			default:
				return carapace.ActionValues()
			}
		}),
		"option": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 1:
						return carapace.ActionValuesDescribed(
							"bool", "boolean value",
							"float", "floating point value",
							"int", "integer value",
							"path", "path value",
							"pkg", "single package",
							"pkgs", "list of packages (appends)",
							"pkgs!", "list of packages (replaces)",
							"string", "string value",
						)
					default:
						return carapace.ActionValues()
					}
				})
			default:
				return carapace.ActionValues()
			}
		}),
		"override-input": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return devenv.ActionInputs()
			case 1:
				return actionSources()
			default:
				return carapace.ActionValues()
			}
		}),
		"profile":  action.ActionProfiles(rootCmd),
		"shell":    actionShells(),
		"system":   nix.ActionSystems(),
		"trace-to": devenv.ActionTraceTargets(),
		"tui":      carapace.ActionValues("false", "true"),
	})
}

// actionExecutionModes completes the execution mode for tasks
func actionExecutionModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"after", "run the task and its dependents",
		"all", "run the task with its dependencies and dependents",
		"before", "run the task and its dependencies",
		"single", "run the task on its own",
	).Tag("execution modes")
}

// actionShells completes shells supported for interactive sessions
func actionShells() carapace.Action {
	return carapace.ActionValues("bash", "fish", "nu", "zsh").Tag("shells")
}

// actionSources completes flake references and `path:` prefixed local paths
func actionSources() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if strings.HasPrefix(c.Value, "path:") {
			return carapace.ActionDirectories().Prefix("path:")
		}

		return carapace.Batch(
			carapace.ActionValues("path:").NoSpace(':'),
			nix.ActionFlakes(),
		).ToA()
	}).Tag("sources")
}
