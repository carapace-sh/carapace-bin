package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var sendPackCmd = &cobra.Command{
	Use:     "send-pack",
	Short:   "Push objects over Git protocol to another repository",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_synching].ID,
}

func init() {
	carapace.Gen(sendPackCmd).Standalone()

	sendPackCmd.Flags().Bool("all", false, "update all heads that locally exist")
	sendPackCmd.Flags().Bool("atomic", false, "use an atomic transaction for updating the refs")
	sendPackCmd.Flags().Bool("dry-run", false, "do everything except actually send the updates")
	sendPackCmd.Flags().String("exec", "", "same as --receive-pack=<git-receive-pack>")
	sendPackCmd.Flags().BoolP("force", "f", false, "disable ancestor check")
	sendPackCmd.Flags().Bool("force-if-includes", false, "require remote updates to be integrated locally")
	sendPackCmd.Flags().String("force-with-lease", "", "require old value of ref to be at this value")
	sendPackCmd.Flags().Bool("helper-status", false, "print status from remote helper")
	sendPackCmd.Flags().Bool("mirror", false, "mirror all refs")
	sendPackCmd.Flags().Bool("progress", false, "force progress reporting")
	sendPackCmd.Flags().StringArray("push-option", nil, "pass the specified string as a push option")
	sendPackCmd.Flags().String("receive-pack", "", "path to the git-receive-pack program on the remote end")
	sendPackCmd.Flags().String("remote", "", "remote name")
	sendPackCmd.Flags().String("signed", "", "GPG-sign the push request")

	sendPackCmd.Flag("force-with-lease").NoOptDefVal = " "
	sendPackCmd.Flag("signed").NoOptDefVal = "if-asked"
	sendPackCmd.Flags().Bool("stateless-rpc", false, "use stateless RPC protocol")
	sendPackCmd.Flags().Bool("stdin", false, "take the list of refs from stdin, one per line")
	sendPackCmd.Flags().Bool("thin", false, "send a \"thin\" pack")
	sendPackCmd.Flags().Bool("verbose", false, "run verbosely")
	rootCmd.AddCommand(sendPackCmd)

	carapace.Gen(sendPackCmd).FlagCompletion(carapace.ActionMap{
		"force-with-lease": carapace.ActionValues(), // TODO ref:value completion
		"push-option":      git.ActionPushOptions(),
		"signed":           carapace.ActionValues("true", "false", "if-asked").StyleF(style.ForKeyword),
	})

	carapace.Gen(sendPackCmd).PositionalCompletion(
		carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return net.ActionHosts().Suffix(":")
			default:
				return carapace.ActionValues()
			}
		}),
	)

	carapace.Gen(sendPackCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if sendPackCmd.Flag("all").Changed {
				return carapace.ActionValues()
			}
			return git.ActionRefs(git.RefOption{}.Default()).FilterArgs().Shift(1)
		}),
	)
}
