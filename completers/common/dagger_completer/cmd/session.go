package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sessionCmd = &cobra.Command{
	Use:    "session",
	Short:  "",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sessionCmd).Standalone()

	sessionCmd.Flags().String("label", "", "label that identifies the source of this session (e.g, --label 'dagger.io/sdk.name:python' --label 'dagger.io/sdk.version:0.5.2' --label 'dagger.io/sdk.async:true')")
	rootCmd.AddCommand(sessionCmd)
}
