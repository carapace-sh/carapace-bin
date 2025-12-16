package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nproc",
	Short: "print the number of processing units available",
	Long:  "https://www.man7.org/linux/man-pages/man1/nproc.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("all", false, "print the number of installed processors")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().String("ignore", "", "if possible, exclude N processing units")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
}
