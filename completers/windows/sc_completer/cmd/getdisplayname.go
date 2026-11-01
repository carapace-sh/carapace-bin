package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getdisplaynameCmd = &cobra.Command{
	Use:   "getdisplayname",
	Short: "get the display name for a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getdisplaynameCmd).Standalone()
	rootCmd.AddCommand(getdisplaynameCmd)
}
