package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Create a new tape file by recording your actions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(recordCmd).Standalone()

	recordCmd.Flags().StringP("shell", "s", "", "shell for recording")
	rootCmd.AddCommand(recordCmd)

	carapace.Gen(recordCmd).FlagCompletion(carapace.ActionMap{
		"shell": os.ActionShells(),
	})
}
