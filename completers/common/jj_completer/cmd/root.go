package cmd

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	shlex "github.com/carapace-sh/carapace-shlex"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/carapace-sh/carapace/third_party/golang.org/x/sys/execabs"
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

	rootCmd.PersistentFlags().String("at-operation", "", "Operation to load the repo at")
	rootCmd.PersistentFlags().String("color", "", "When to colorize output (always, never, auto)")
	rootCmd.PersistentFlags().StringSlice("config", nil, "Additional configuration options (can be repeated)")
	rootCmd.PersistentFlags().StringSlice("config-file", nil, "Additional configuration files (can be repeated)")
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug logging")
	rootCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.PersistentFlags().Bool("ignore-immutable", false, "Allow rewriting of immutable commits")
	rootCmd.PersistentFlags().Bool("ignore-working-copy", false, "Don't snapshot the working copy, and don't update it")
	rootCmd.PersistentFlags().Bool("no-pager", false, "Disable the pager")
	rootCmd.PersistentFlags().Bool("quiet", false, "Silence non-primary output")
	rootCmd.PersistentFlags().StringP("repository", "R", "", "Path to repository to operate on")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose logging")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	rootCmd.MarkFlagsMutuallyExclusive("quiet", "verbose")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"at-operation": jj.ActionOperations(100),
		"color":        carapace.ActionValues("always", "never", "debug", "auto").StyleF(style.ForKeyword),
		"config-file":  carapace.ActionFiles(),
		"repository":   carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.ChdirF(traverse.Flag(rootCmd.Flag("repository")))
	})

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		output, err := execabs.Command("jj", "config", "get", "aliases").Output()
		if err != nil {
			carapace.LOG.Println(err.Error())
			return
		}
		s := string(output)
		s = strings.TrimLeft(s, "{ ")
		s = strings.TrimRight(s, " }\n")
		for _, alias := range strings.Split(s, ", ") {
			if name, value, ok := strings.Cut(alias, " = "); ok {
				var args []string
				if err := json.Unmarshal([]byte(value), &args); err != nil {
					carapace.LOG.Println(err.Error())
					continue
				}

				aliasCmd := &cobra.Command{
					Use:                name,
					Short:              shlex.Join(args),
					GroupID:            "alias",
					DisableFlagParsing: true,
					Run:                func(cmd *cobra.Command, args []string) {},
				}
				cmd.Root().AddCommand(aliasCmd)

				carapace.Gen(aliasCmd).PositionalAnyCompletion(
					carapace.ActionCallback(func(c carapace.Context) carapace.Action {
						c.Args = append(args, c.Args...)
						return bridge.ActionCarapaceBin("jj").Invoke(c).ToA()
					}),
				)
			}
		}
	})
}
