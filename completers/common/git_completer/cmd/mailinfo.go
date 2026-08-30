package cmd

import (
	"github.com/carapace-sh/carapace"
	net_http "github.com/carapace-sh/carapace-bin/pkg/actions/net/http"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var mailinfoCmd = &cobra.Command{
	Use:     "mailinfo",
	Short:   "Extracts patch and authorship from a single e-mail message",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(mailinfoCmd).Standalone()

	mailinfoCmd.Flags().BoolS("b", "b", false, "keep non patch brackets in subject")
	mailinfoCmd.Flags().String("encoding", "", "re-code metadata to this encoding")
	mailinfoCmd.Flags().Bool("inbody-headers", false, "use headers in message's body")
	mailinfoCmd.Flags().BoolS("k", "k", false, "keep subject")
	mailinfoCmd.Flags().BoolP("message-id", "m", false, "copy Message-ID to the end of commit message")
	mailinfoCmd.Flags().BoolS("n", "n", false, "disable charset re-coding of metadata")
	mailinfoCmd.Flags().String("quoted-cr", "", "action when quoted CR is found")
	mailinfoCmd.Flags().Bool("scissors", false, "use scissors")
	mailinfoCmd.Flags().BoolS("u", "u", false, "re-code metadata to i18n.commitEncoding")
	mailinfoCmd.Flag("inbody-headers").Hidden = true
	rootCmd.AddCommand(mailinfoCmd)

	carapace.Gen(mailinfoCmd).FlagCompletion(carapace.ActionMap{
		"encoding":  net_http.ActionCharsets(),
		"quoted-cr": git.ActionQuotedCrModes(),
	})
}
