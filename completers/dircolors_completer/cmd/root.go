package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dircolors",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("csh,", "c", "", "output C shell code to set LS_COLORS")
	rootCmd.Flags().String("c-shell", "", "output C shell code to set LS_COLORS")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("print-database", "p", false, "output defaults")
	rootCmd.Flags().StringP("sh,", "b", "", "output Bourne shell code to set LS_COLORS")
	rootCmd.Flags().String("bourne-shell", "", "output Bourne shell code to set LS_COLORS")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).PositionalCompletion(carapace.ActionFiles(""))
}
