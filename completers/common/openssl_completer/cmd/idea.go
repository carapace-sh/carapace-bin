package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var ideaCmd = &cobra.Command{
	Use:                "idea",
	Short:              "IDEA Cipher",
	GroupID:            "cipher",
	Aliases:            []string{"idea-cbc", "idea-cfb", "idea-ecb", "idea-ofb"},
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(ideaCmd).Standalone()

	rootCmd.AddCommand(ideaCmd)

	carapace.Gen(ideaCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("openssl", "enc"),
	)
}
