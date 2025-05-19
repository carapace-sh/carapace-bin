package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "env",
	Short: "run a program in a modified environment",
	Long:  "https://linux.die.net/man/1/env",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().String("block-signal", "", "block delivery of SIG signal(s) to COMMAND")
	rootCmd.Flags().StringP("chdir", "C", "", "change working directory to DIR")
	rootCmd.Flags().BoolP("debug", "v", false, "print verbose information for each processing step")
	rootCmd.Flags().String("default-signal", "", "reset handling of SIG signal(s) to the default")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("ignore-environment", "i", false, "start with an empty environment")
	rootCmd.Flags().String("ignore-signal", "", "set handling of SIG signals(s) to do nothing")
	rootCmd.Flags().Bool("list-signal-handling", false, "list non default signal handling to stderr")
	rootCmd.Flags().BoolP("null", "0", false, "end each output line with NUL, not newline")
	rootCmd.Flags().StringP("split-string", "S", "", "process and split S into separate arguments;")
	rootCmd.Flags().StringP("unset", "u", "", "remove variable from the environment")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	rootCmd.Flag("block-signal").NoOptDefVal = " "
	rootCmd.Flag("default-signal").NoOptDefVal = " "
	rootCmd.Flag("ignore-signal").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"block-signal":   ps.ActionKillSignals(),
		"chdir":          carapace.ActionDirectories(),
		"default-signal": ps.ActionKillSignals(),
		"ignore-signal":  ps.ActionKillSignals(),
		"unset":          os.ActionEnvironmentVariables(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			for index, arg := range c.Args {
				if strings.Contains(arg, "=") {
					splitted := strings.SplitN(arg, "=", 2)
					c.Setenv(splitted[0], splitted[1])
				} else {
					return bridge.ActionCarapaceBin().Shift(index).Invoke(c).ToA()
				}
			}

			return carapace.Batch(
				env.ActionNameValues(false),
				carapace.ActionExecutables(),
				carapace.ActionFiles(),
			).Invoke(c).Merge().ToA()
		}),
	)

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.Chdir(rootCmd.Flag("chdir").Value.String())
	})
}
