package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nfsstat",
	Short: "display NFS statistics",
	Long:  "https://keith.github.io/xcode-manpages/nfsstat.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("3", "3", false, "Print only NFS v3 statistics")
	rootCmd.Flags().BoolS("4", "4", false, "Print only NFS v4 statistics")
	rootCmd.Flags().BoolS("E", "E", false, "Show NFS protocol errors statistics")
	rootCmd.Flags().BoolS("c", "c", false, "Show NFS client statistics")
	rootCmd.Flags().BoolS("e", "e", false, "Show NFS server exported directory statistics")
	rootCmd.Flags().BoolS("s", "s", false, "Show NFS server statistics")
	rootCmd.Flags().BoolS("u", "u", false, "Show NFS server active user statistics")
	rootCmd.Flags().BoolS("v", "v", false, "Show additional information")
	rootCmd.Flags().BoolS("z", "z", false, "Zero NFS client and server statistics")

	rootCmd.Flags().StringS("f", "f", "", "Output format (JSON)")
	rootCmd.Flags().StringS("n", "n", "", "Select NFS statistics by user/net")
	rootCmd.Flags().StringS("w", "w", "", "Wait interval between statistics display")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"f": carapace.ActionValues("JSON"),
		"n": carapace.ActionValues("user", "net"),
	})
}
