package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/openrc"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rc-status [flags] [runlevel]",
	Short: "show status info about OpenRC runlevels",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "Show all runlevels and their services")
	rootCmd.Flags().BoolP("crashed", "c", false, "List all crashed services")
	rootCmd.Flags().StringP("format", "f", "", "Select an output format")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information and exit")
	rootCmd.Flags().StringArrayP("in-state", "i", nil, "Show services in the given state")
	rootCmd.Flags().BoolP("list", "l", false, "List all defined runlevels")
	rootCmd.Flags().BoolP("manual", "m", false, "Show all manually started services")
	rootCmd.Flags().BoolP("nocolor", "C", false, "Disable color output")
	rootCmd.Flags().BoolP("quiet", "q", false, "Run quietly")
	rootCmd.Flags().BoolP("runlevel", "r", false, "Print the current runlevel name")
	rootCmd.Flags().BoolP("servicelist", "s", false, "Show all services")
	rootCmd.Flags().BoolP("supervised", "S", false, "Show all supervised services")
	rootCmd.Flags().BoolP("unused", "u", false, "Show services not assigned to any runlevel")
	rootCmd.Flags().BoolP("user", "U", false, "Operate on user services and runlevels")
	rootCmd.Flags().BoolP("verbose", "v", false, "Run verbosely")
	rootCmd.Flags().BoolP("version", "V", false, "Display software version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"format":   carapace.ActionValues("ini"),
		"in-state": openrc.ActionServiceStates(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		openrc.ActionRunlevels(),
	)
}
