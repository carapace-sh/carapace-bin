package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pass"
	"github.com/spf13/cobra"
)

var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "insert new password",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(insertCmd).Standalone()
	insertCmd.Flags().BoolP("echo", "e", false, "echo the password back to the console")
	insertCmd.Flags().BoolP("force", "f", false, "overwrite existing entry without prompt")
	insertCmd.Flags().BoolP("multiline", "m", false, "multiline entry")
	rootCmd.AddCommand(insertCmd)

	carapace.Gen(insertCmd).PositionalCompletion(
		pass.ActionPasswords(),
	)
}
