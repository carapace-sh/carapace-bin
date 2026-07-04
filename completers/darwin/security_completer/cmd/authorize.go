package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var authorizeCmd = &cobra.Command{
	Use:   "authorize",
	Short: "Perform authorization operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(authorizeCmd).Standalone()
	authorizeCmd.Flags().BoolP("destroy", "d", false, "Destroy acquired rights")
	authorizeCmd.Flags().BoolP("externalize", "e", false, "Externalize authref to stdout")
	authorizeCmd.Flags().BoolP("internalize", "i", false, "Internalize authref passed on stdin")
	authorizeCmd.Flags().BoolP("least-privileged", "l", false, "Operate in least privileged mode")
	authorizeCmd.Flags().BoolP("partial-rights", "p", false, "Allow returning partial rights")
	authorizeCmd.Flags().BoolP("pre-authorize", "P", false, "Pre-authorize rights only")
	authorizeCmd.Flags().BoolP("user-interaction", "u", false, "Allow user interaction")
	authorizeCmd.Flags().BoolP("wait", "w", false, "Wait while holding AuthorizationRef until stdout is closed")
	rootCmd.AddCommand(authorizeCmd)
}
