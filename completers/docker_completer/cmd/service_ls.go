package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var service_lsCmd = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Short:   "List services",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_lsCmd).Standalone()

	service_lsCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	service_lsCmd.Flags().String("format", "", "Format output using a custom template:")
	service_lsCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	serviceCmd.AddCommand(service_lsCmd)
}
