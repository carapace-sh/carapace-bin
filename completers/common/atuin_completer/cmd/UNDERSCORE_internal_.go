package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var __internal_Cmd = &cobra.Command{
	Use:    "__internal_",
	Short:  "We want to exclude the `__internal` subcommand from Clap's `infer_subcommands`; otherwise, a user could access it simply by typing `atuin _`. However, Clap has no way to disable `infer_subcommands` for a single command. As a workaround, we define a dummy command with the same name but with an extra understore, which forces `__internal` to be typed out in entirety, since any prefix of the name would be ambiguous",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(__internal_Cmd).Standalone()

	rootCmd.AddCommand(__internal_Cmd)
}
