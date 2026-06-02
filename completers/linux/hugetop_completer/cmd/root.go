package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hugetop",
	Short: "report huge page information",
	Long:  "https://man7.org/linux/man-pages/man1/hugetop.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("delay", "d", "", "delay updates")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("human", "H", false, "display human-readable output")
	rootCmd.Flags().BoolP("numa", "n", false, "display per NUMA nodes Huge pages information")
	rootCmd.Flags().BoolP("once", "o", false, "only display once, then exit")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")

	rootCmd.MarkFlagsMutuallyExclusive("delay", "once")
}
