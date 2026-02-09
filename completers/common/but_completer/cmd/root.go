package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	shlex "github.com/carapace-sh/carapace-shlex"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/carapace-sh/carapace/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "but",
	Short: "A GitButler CLI tool",
	Long:  "https://github.com/gitbutlerapp/gitbutler/tree/master/crates/but",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "inspection"},
		&cobra.Group{ID: "branching and committing"},
		&cobra.Group{ID: "server interactions"},
		&cobra.Group{ID: "editing commits"},
		&cobra.Group{ID: "operation history"},
		&cobra.Group{ID: "alias"},
	)

	rootCmd.Flags().StringP("current-dir", "C", "", "Run as if gitbutler-cli was started in PATH instead of the current working directory")
	rootCmd.Flags().StringP("format", "f", "", "Explicitly control how output should be formatted")
	rootCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.PersistentFlags().BoolP("json", "j", false, "Whether to use JSON output format")
	rootCmd.PersistentFlags().Bool("status-after", false, "After a mutation command completes, also output workspace status")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"current-dir": carapace.ActionDirectories(),
		"format":      carapace.ActionValues("human", "shell", "json", "none"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCliIds(but.CliIdsOpts{}.Default()),
			carapace.ActionDirectories(),
		).ToA(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.Args[0]) {
				return carapace.ActionValues()
			}
			return carapace.Batch(
				but.ActionCliIds(but.CliIdsOpts{Branches: true, Stacks: true}),
				but.ActionLocalBranches(),
			).ToA().FilterArgs()
		}),
	)

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, _ *pflag.Flag, action carapace.Action) carapace.Action {
		return action.ChdirF(traverse.Flag(cmd.Flag("current-dir")))
	})

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		if aliases, err := but.Aliases(); err == nil {
			m := aliases.Default
			for _, alias := range aliases.User {
				m[alias.Name] = alias.Value
			}
			for key, value := range m {
				aliasCmd := &cobra.Command{
					Use:                key,
					Short:              strings.TrimSpace(value),
					GroupID:            "alias",
					DisableFlagParsing: true,
					Run:                func(cmd *cobra.Command, args []string) {},
				}
				cmd.Root().AddCommand(aliasCmd)

				carapace.Gen(aliasCmd).PositionalAnyCompletion(
					carapace.ActionCallback(func(c carapace.Context) carapace.Action {
						splitted, err := shlex.Split(aliasCmd.Short)
						if err != nil {
							return carapace.ActionMessage(err.Error())
						}
						c.Args = append(splitted.Strings(), c.Args...)
						return bridge.ActionCarapaceBin("but").Invoke(c).ToA()
					}),
				)
			}
		}

	})
}
