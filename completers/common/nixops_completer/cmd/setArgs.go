package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setArgsCmd = &cobra.Command{
	Use:   "set-args",
	Short: "persistently set deployment arguments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setArgsCmd).Standalone()

	setArgsCmd.Flags().StringSlice("arg", nil, "Pass a Nix expression value")
	setArgsCmd.Flags().StringSlice("argstr", nil, "Pass a string value")
	setArgsCmd.Flags().StringArray("unset", nil, "Unset previously set argument")

	setArgsCmd.Flag("arg").Nargs = 2
	setArgsCmd.Flag("argstr").Nargs = 2

	rootCmd.AddCommand(setArgsCmd)
}
