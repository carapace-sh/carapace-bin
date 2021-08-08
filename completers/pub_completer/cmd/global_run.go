package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var global_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an executable from a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_runCmd).Standalone()

	global_runCmd.Flags().Bool("enable-asserts", false, "Enable assert statements.")
	global_runCmd.Flags().String("enable-experiment", "", "Runs the executable in a VM with the given experiments enabled.")
	global_runCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	global_runCmd.Flags().Bool("no-enable-asserts", false, "Do not enable assert statements.")
	global_runCmd.Flags().Bool("no-sound-null-safety", false, "Do not override the default null safety execution mode.")
	global_runCmd.Flags().Bool("sound-null-safety", false, "Override the default null safety execution mode.")
	globalCmd.AddCommand(global_runCmd)

	// TODO positional completion
}
