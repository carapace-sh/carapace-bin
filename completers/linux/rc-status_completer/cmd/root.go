package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/openrc"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rc-status",
	Short: "show status of OpenRC services",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "show services from all run levels")
	rootCmd.Flags().BoolP("crashed", "c", false, "show crashed services")
	rootCmd.Flags().StringP("format", "f", "", "format status to be parsable (currently arg must be ini)")
	rootCmd.Flags().Bool("help", false, "display this help output")
	rootCmd.Flags().StringP("in-state", "i", "", "show services which are in this state")
	rootCmd.Flags().BoolP("list", "l", false, "show list of run levels")
	rootCmd.Flags().BoolP("manual", "m", false, "show manually started services")
	rootCmd.Flags().Bool("nocolor", false, "disable color output")
	rootCmd.Flags().BoolP("quiet", "q", false, "run quietly (repeat to suppress errors)")
	rootCmd.Flags().BoolP("runlevel", "r", false, "show the name of the current runlevel")
	rootCmd.Flags().BoolP("servicelist", "s", false, "show service list")
	rootCmd.Flags().BoolP("supervised", "S", false, "show supervised services")
	rootCmd.Flags().BoolP("unused", "u", false, "show services not assigned to any runlevel")
	rootCmd.Flags().Bool("user", false, "run in user mode")
	rootCmd.Flags().BoolP("verbose", "v", false, "run verbosely")
	rootCmd.Flags().BoolP("version", "V", false, "display software version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"format":   carapace.ActionValues("ini"),
		"in-state": openrc.ActionServiceStates(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		openrc.ActionRunlevels(),
	)
}
