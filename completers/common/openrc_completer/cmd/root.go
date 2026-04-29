package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openrc",
	Short: "dependency based init system",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().BoolP("no-stop", "n", false, "do not stop any services")
	rootCmd.Flags().BoolP("override", "o", false, "override the inactive runlevel warning")
	rootCmd.Flags().BoolP("quiet", "q", false, "suppress non-error output")
	rootCmd.Flags().StringP("runlevel", "r", "", "change to the specified runlevel")
	rootCmd.Flags().BoolP("service", "s", false, "run a single service")
	rootCmd.Flags().BoolP("sys", "S", false, "output the system type")
	rootCmd.Flags().BoolP("verbose", "v", false, "show extra output")
	rootCmd.Flags().BoolP("version", "V", false, "show version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"runlevel": actionRunlevels(),
	})
}

func actionRunlevels() carapace.Action {
	return carapace.ActionValuesDescribed(
		"boot", "system boot runlevel",
		"default", "default runlevel",
		"nonetwork", "runlevel without networking",
		"shutdown", "system shutdown runlevel",
		"single", "single-user runlevel",
		"sysinit", "system initialization runlevel",
	)
}
