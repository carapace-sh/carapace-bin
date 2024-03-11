package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secret_lsCmd = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Short:   "List secrets",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_lsCmd).Standalone()

	secret_lsCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	secret_lsCmd.Flags().String("format", "", "Format output using a custom template:")
	secret_lsCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	secretCmd.AddCommand(secret_lsCmd)
}
