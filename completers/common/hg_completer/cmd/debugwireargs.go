package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugwireargsCmd = &cobra.Command{
	Use:    "debugwireargs",
	Short:  "REPO [OPTIONS]... [ONE [TWO]]",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugwireargsCmd).Standalone()

	debugwireargsCmd.Flags().String("five", "", "five")
	debugwireargsCmd.Flags().String("four", "", "four")
	debugwireargsCmd.Flags().Bool("insecure", false, "do not verify server certificate (ignoring web.cacerts config)")
	debugwireargsCmd.Flags().String("remotecmd", "", "specify hg command to run on the remote side")
	debugwireargsCmd.Flags().StringP("ssh", "e", "", "specify ssh command to use")
	debugwireargsCmd.Flags().String("three", "", "three")
	rootCmd.AddCommand(debugwireargsCmd)

	carapace.Gen(debugwireargsCmd).FlagCompletion(carapace.ActionMap{
		"remotecmd": carapace.ActionExecutables(),
	})
}
