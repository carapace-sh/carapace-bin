package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "hf",
	Short:              "Hugging Face CLI",
	Long:               "https://huggingface.co/docs/huggingface_hub/guides/cli",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	// everything except `hf cache` is delegated to the click completion of hf.
	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionClick("hf"),
	)
}
