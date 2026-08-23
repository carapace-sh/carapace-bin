package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ipcrm",
	Short: "remove a System V interprocess communication facility",
	Long:  "https://man.freebsd.org/cgi/man.cgi?ipcrm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("M", "M", "", "Remove shared memory identifier")
	rootCmd.Flags().StringS("Q", "Q", "", "Remove message queue identifier")
	rootCmd.Flags().StringS("S", "S", "", "Remove semaphore identifier")
	rootCmd.Flags().BoolS("m", "m", false, "Remove shared memory")
	rootCmd.Flags().BoolS("q", "q", false, "Remove message queue")
	rootCmd.Flags().BoolS("s", "s", false, "Remove semaphore")
	rootCmd.Flags().BoolS("z", "z", false, "Remove all")
}
