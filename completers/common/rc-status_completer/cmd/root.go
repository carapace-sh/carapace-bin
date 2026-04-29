package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rc-status",
	Short: "show OpenRC service status",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error { return rootCmd.Execute() }

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "show all runlevels")
	rootCmd.Flags().BoolP("crashed", "c", false, "show crashed services")
	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().BoolP("list", "l", false, "list runlevels")
	rootCmd.Flags().StringP("runlevel", "r", "", "show services in runlevel")
	rootCmd.Flags().BoolP("servicelist", "s", false, "show service list only")
	rootCmd.Flags().BoolP("unused", "u", false, "show services not assigned to any runlevel")
	rootCmd.Flags().BoolP("version", "V", false, "show version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{"runlevel": actionRunlevels()})
	carapace.Gen(rootCmd).PositionalAnyCompletion(actionRunlevels())
}

func actionRunlevels() carapace.Action {
	return carapace.ActionValues("boot", "default", "nonetwork", "shutdown", "single", "sysinit")
}
