package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var srpCmd = &cobra.Command{
	Use:     "srp",
	Short:   "Maintain SRP password file",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(srpCmd).Standalone()

	srpCmd.Flags().BoolS("add", "add", false, "Add a user and SRP verifier")
	srpCmd.Flags().StringS("config", "config", "", "A config file")
	srpCmd.Flags().BoolS("delete", "delete", false, "Delete user from verifier file")
	srpCmd.Flags().StringS("engine", "engine", "", "Use engine, possibly a hardware device")
	srpCmd.Flags().StringS("gn", "gn", "", "Set g and N values to be used for new verifier")
	srpCmd.Flags().BoolS("list", "list", false, "List users")
	srpCmd.Flags().BoolS("modify", "modify", false, "Modify the SRP verifier of an existing user")
	srpCmd.Flags().StringS("name", "name", "", "The particular srp definition to use")
	srpCmd.Flags().StringS("passin", "passin", "", "Input file pass phrase source")
	srpCmd.Flags().StringS("passout", "passout", "", "Output file pass phrase source")
	srpCmd.Flags().StringS("srpvfile", "srpvfile", "", "The srp verifier file name")
	srpCmd.Flags().StringS("userinfo", "userinfo", "", "Additional info to be set for user")
	srpCmd.Flags().BoolS("verbose", "verbose", false, "Talk a lot while doing things")
	rootCmd.AddCommand(srpCmd)

	carapace.Gen(srpCmd).FlagCompletion(carapace.ActionMap{
		"config":   carapace.ActionFiles(),
		"engine":   action.ActionEngines(),
		"srpvfile": carapace.ActionFiles(),
	})
}
