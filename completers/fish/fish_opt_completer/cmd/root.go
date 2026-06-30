package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fish_opt",
	Short: "Create option spec for argparse",
	Long:  "https://fishshell.com/docs/current/cmds/fish_opt.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("d", "d", false, "delete option")
	rootCmd.Flags().Bool("delete", false, "delete option")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().StringS("l", "l", "", "long flag name")
	rootCmd.Flags().String("long", "", "long flag name")
	rootCmd.Flags().Bool("long-only", false, "only long flag")
	rootCmd.Flags().BoolS("m", "m", false, "accumulate multiple values")
	rootCmd.Flags().Bool("multiple-vals", false, "accumulate multiple values")
	rootCmd.Flags().BoolS("o", "o", false, "optional value")
	rootCmd.Flags().Bool("optional-val", false, "optional value")
	rootCmd.Flags().BoolS("r", "r", false, "required value")
	rootCmd.Flags().Bool("required-val", false, "required value")
	rootCmd.Flags().StringS("s", "s", "", "short flag character")
	rootCmd.Flags().String("short", "", "short flag character")
	rootCmd.Flags().StringS("v", "v", "", "validate command")
	rootCmd.Flags().String("validate", "", "validate command")
}
