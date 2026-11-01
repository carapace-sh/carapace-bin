package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "smbutil",
	Short: "interface to the SMB requester",
	Long:  "https://keith.github.io/xcode-manpages/smbutil.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Print a short help message")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
}
