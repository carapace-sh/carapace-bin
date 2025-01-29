package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secret_lsCmd = &cobra.Command{
	Use:     "ls [options]",
	Short:   "List secrets",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_lsCmd).Standalone()

	secret_lsCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter secret output")
	secret_lsCmd.Flags().String("format", "", "Format volume output using Go template")
	secret_lsCmd.Flags().BoolP("noheading", "n", false, "Do not print headers")
	secret_lsCmd.Flags().BoolP("quiet", "q", false, "Print secret IDs only")
	secretCmd.AddCommand(secret_lsCmd)
}
