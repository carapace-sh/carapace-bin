package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verifies the on-chain bytecode matches the locally compiled artifact. Run this command inside a program subdirectory, i.e., in the dir containing the program's Cargo.toml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifyCmd).Standalone()

	verifyCmd.Flags().String("commit-hash", "", "The commit hash to verify against. Requires `--repo-url`")
	verifyCmd.Flags().Bool("current-dir", false, "Verify against the source code in the current directory. Conflicts with `--repo-url`")
	verifyCmd.Flags().BoolP("help", "h", false, "Print help")
	verifyCmd.Flags().String("program-name", "", "Name of the program to run the command on. Defaults to the package name")
	verifyCmd.Flags().String("repo-url", "", "The URL of the repository to verify against. Conflicts with `--current-dir`")
	rootCmd.AddCommand(verifyCmd)
}
