package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbacklight",
	Short: "adjust backlight brightness using RandR extension",
	Long:  "https://linux.die.net/man/1/xbacklight",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	carapace.Override(carapace.Opts{
		LongShorthand: true,
	})
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("d", "", "Display to set backlight for.")
	rootCmd.Flags().String("dec", "", "Decreases brightness by the specified amount.")
	rootCmd.Flags().String("display", "", "Display to set backlight for.")
	rootCmd.Flags().Bool("get", false, "Print  out  the  current backlight brightness of each output with such a control.")
	rootCmd.Flags().Bool("help", false, "Print out a summary of the usage and exit.")
	rootCmd.Flags().String("inc", "", "Increases brightness by the specified amount.")
	rootCmd.Flags().String("set", "", "Sets each backlight brightness to the specified level.")
	rootCmd.Flags().String("steps", "", "Number of steps to take while fading.")
	rootCmd.Flags().String("time", "", "Length of time to spend fading the backlight between old and new value.")
	rootCmd.Flags().Bool("version", false, "Print out the program version and exit.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"d":       os.ActionDisplays(),
		"display": os.ActionDisplays(),
	})
}
