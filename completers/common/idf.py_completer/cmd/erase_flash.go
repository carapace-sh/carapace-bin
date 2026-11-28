package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eraseFlashCmd = &cobra.Command{
	Use:     "erase-flash",
	Aliases: []string{"erase_flash"},
	Short:   "Erase entire flash chip",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eraseFlashCmd).Standalone()
	rootCmd.AddCommand(eraseFlashCmd)
}
