package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tload",
	Short: "graphic representation of system load average",
	Long:  "https://www.man7.org/linux/man-pages/man1/tload.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("delay", "d", "", "update delay in seconds")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().StringP("scale", "s", "", "vertical scale")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")
}
