package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var changerecoveryCmd = &cobra.Command{
	Use:   "changerecovery",
	Short: "Change or add recovery key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(changerecoveryCmd).Standalone()
	rootCmd.AddCommand(changerecoveryCmd)

	changerecoveryCmd.Flags().Bool("inputplist", false, "Read configuration from stdin")
	changerecoveryCmd.Flags().Bool("institutional", false, "Specify an institutional recovery key")
	changerecoveryCmd.Flags().Bool("norecoverykey", false, "Do not return a recovery key")
	changerecoveryCmd.Flags().Bool("outputplist", false, "Output key and computer info to stdout")
	changerecoveryCmd.Flags().Bool("personal", false, "Specify a personal recovery key")
	changerecoveryCmd.Flags().Bool("verbose", false, "Enable verbose mode")
}
