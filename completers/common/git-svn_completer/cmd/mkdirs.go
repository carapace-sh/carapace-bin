package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mkdirsCmd = &cobra.Command{
	Use:   "mkdirs",
	Short: "Recreate empty directories after a checkout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mkdirsCmd).Standalone()

	mkdirsCmd.Flags().IntP("revision", "r", 0, "Refer to a specific revision")
	rootCmd.AddCommand(mkdirsCmd)
}
