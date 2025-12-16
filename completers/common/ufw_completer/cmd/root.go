package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ufw",
	Short: "program for managing a netfilter firewall",
	Long:  "https://launchpad.net/ufw",
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("dry-run", false, "don't modify anything, just show the changes")
	rootCmd.Flags().Bool("force", false, "")
	rootCmd.Flags().BoolP("help", "h", false, "show help message and exit")
	rootCmd.Flags().Bool("version", false, "show program's version number and exit")
}
