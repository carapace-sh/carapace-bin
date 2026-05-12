package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/openrc"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openrc [flags] [runlevel]",
	Short: "stops and starts services for the specified runlevel",
	Long:  "https://github.com/OpenRC/openrc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Display usage information and exit")
	rootCmd.Flags().BoolP("no-stop", "n", false, "Do not stop any services")
	rootCmd.Flags().BoolP("nocolor", "C", false, "Disable color output")
	rootCmd.Flags().BoolP("override", "o", false, "Override the next runlevel")
	rootCmd.Flags().BoolP("quiet", "q", false, "Run quietly")
	rootCmd.Flags().BoolP("service", "s", false, "Run the service specified with the rest of the arguments")
	rootCmd.Flags().BoolP("sys", "S", false, "Output the RC system type")
	rootCmd.Flags().BoolP("user", "U", false, "Operate on user services and runlevels")
	rootCmd.Flags().BoolP("verbose", "v", false, "Run verbosely")
	rootCmd.Flags().BoolP("version", "V", false, "Display software version")

	carapace.Gen(rootCmd).PositionalCompletion(
		openrc.ActionRunlevels(),
	)
}
