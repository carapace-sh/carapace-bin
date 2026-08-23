package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add user(s) to existing FileVault",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().Bool("inputplist", false, "Read configuration from stdin")
	addCmd.Flags().Bool("quiet", false, "No status during operation")
	addCmd.Flags().String("user", "", "Short user name")
	addCmd.Flags().String("usertoadd", "", "Additional user name")
	addCmd.Flags().String("uuid", "", "User UUID")
	addCmd.Flags().Bool("verbose", false, "Enable verbose mode")

	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"user":      os.ActionUsers(),
		"usertoadd": os.ActionUsers(),
	})
}
