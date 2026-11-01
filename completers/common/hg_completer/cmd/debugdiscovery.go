package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var debugdiscoveryCmd = &cobra.Command{
	Use:    "debugdiscovery",
	Short:  "runs the changeset discovery protocol in isolation",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugdiscoveryCmd).Standalone()

	debugdiscoveryCmd.Flags().Bool("insecure", false, "do not verify server certificate (ignoring web.cacerts config)")
	debugdiscoveryCmd.Flags().String("local-as-revs", "", "treat local has having these revisions only")
	debugdiscoveryCmd.Flags().Bool("nonheads", false, "use old-style discovery with non-heads included")
	debugdiscoveryCmd.Flags().Bool("old", false, "use old-style discovery")
	debugdiscoveryCmd.Flags().String("remote-as-revs", "", "use local as remote, with only these revisions")
	debugdiscoveryCmd.Flags().String("remotecmd", "", "specify hg command to run on the remote side")
	debugdiscoveryCmd.Flags().String("rev", "", "restrict discovery to this set of revs")
	debugdiscoveryCmd.Flags().String("seed", "", "specify the random seed use for discovery")
	debugdiscoveryCmd.Flags().StringP("ssh", "e", "", "specify ssh command to use")
	debugdiscoveryCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(debugdiscoveryCmd)

	carapace.Gen(debugdiscoveryCmd).FlagCompletion(carapace.ActionMap{
		"remotecmd": carapace.ActionExecutables(),
		"rev":       hg.ActionRevisions(),
	})
}
