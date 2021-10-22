package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var OneClick_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of 1-Click applications",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(OneClick_listCmd).Standalone()
	OneClick_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `SLUG`, `TYPE`")
	OneClick_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	OneClick_listCmd.Flags().String("type", "", "The 1-Click type. Valid types are one of the following: kubernetes, droplet")
	OneClickCmd.AddCommand(OneClick_listCmd)
}
