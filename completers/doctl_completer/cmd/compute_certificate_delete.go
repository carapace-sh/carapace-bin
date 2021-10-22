package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_certificate_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the specified certificate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_certificate_deleteCmd).Standalone()
	compute_certificate_deleteCmd.Flags().BoolP("force", "f", false, "Delete the certificate without a confirmation prompt")
	compute_certificateCmd.AddCommand(compute_certificate_deleteCmd)
}
