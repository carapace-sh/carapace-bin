package cmd

import (
	"github.com/spf13/cobra"
)

var guiCmd = &cobra.Command{
	Use: "gui",
	Short: "A portable graphical interface to Git",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {

    rootCmd.AddCommand(guiCmd)
}
