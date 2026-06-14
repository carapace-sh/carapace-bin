package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	shlex "github.com/carapace-sh/carapace-shlex"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/carapace-sh/carapace/third_party/golang.org/x/sys/execabs"
	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "jj",
	Short: "Jujutsu (An experimental VCS)",
	Long:  "https://github.com/martinvonz/jj",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("at-op", "", "Operation to load the repo at")
	rootCmd.PersistentFlags().String("at-operation", "", "Operation to load the repo at")
	rootCmd.PersistentFlags().String("color", "", "When to colorize output")
	rootCmd.PersistentFlags().StringSlice("config", nil, "Additional configuration options (can be repeated)")
	rootCmd.PersistentFlags().StringSlice("config-file", nil, "Additional configuration files (can be repeated)")
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug logging")
	rootCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.PersistentFlags().Bool("ignore-immutable", false, "Allow rewriting immutable commits")
	rootCmd.PersistentFlags().Bool("ignore-working-copy", false, "Don't snapshot the working copy, and don't update it")
	rootCmd.PersistentFlags().Bool("no-integrate-operation", false, "Run the command as usual but don't integrate any operations")
	rootCmd.PersistentFlags().Bool("no-pager", false, "Disable the pager")
	rootCmd.PersistentFlags().Bool("quiet", false, "Silence non-primary command output")
	rootCmd.PersistentFlags().StringP("repository", "R", "", "Path to repository to operate on")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose logging")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	rootCmd.MarkFlagsMutuallyExclusive("at-op", "at-operation")
	rootCmd.MarkFlagsMutuallyExclusive("quiet", "verbose")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"at-op":        jj.ActionOperations(100),
		"at-operation": jj.ActionOperations(100),
		"color":        carapace.ActionValues("always", "never", "debug", "auto").StyleF(style.ForKeyword),
		"config-file":  carapace.ActionFiles(),
		"repository":   carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if f := rootCmd.Flag("at-op"); f.Changed {
				c.Setenv("JJ_OPERATION", f.Value.String()) // TODO pseudo environment variable (jj doesn't have one for this)
			}
			if f := rootCmd.Flag("at-operation"); f.Changed {
				c.Setenv("JJ_OPERATION", f.Value.String()) // TODO pseudo environment variable (jj doesn't have one for this)
			}
			if f := rootCmd.Flag("repository"); f.Changed {
				repository, err := c.Abs(f.Value.String())
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}
				c.Setenv("JJ_REPOSITORY", repository) // TODO pseudo environment variable (jj doesn't have one for this)
			}
			return action.Invoke(c).ToA()
		})
	})

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		output, err := execabs.Command("jj", "config", "list", "--include-defaults", "aliases").Output()
		if err != nil {
			carapace.LOG.Println(err.Error())
			return
		}

		var config struct {
			Aliases map[string][]string
		}
		if err = toml.Unmarshal(output, &config); err != nil {
			carapace.LOG.Println(err.Error())
			return
		}

		for name, alias := range config.Aliases {
			aliasCmd := &cobra.Command{
				Use:                name,
				Short:              shlex.Join(alias),
				GroupID:            "alias",
				DisableFlagParsing: true,
				Run:                func(cmd *cobra.Command, args []string) {},
			}
			cmd.Root().AddCommand(aliasCmd)

			carapace.Gen(aliasCmd).PositionalAnyCompletion(
				carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					c.Args = append(alias, c.Args...)
					return bridge.ActionCarapaceBin("jj").Invoke(c).ToA()
				}),
			)
		}
	})
}
