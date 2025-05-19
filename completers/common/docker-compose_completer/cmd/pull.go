package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull [OPTIONS] [SERVICE...]",
	Short: "Pull service images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pullCmd).Standalone()

	pullCmd.Flags().Bool("ignore-buildable", false, "Ignore images that can be built")
	pullCmd.Flags().Bool("ignore-pull-failures", false, "Pull what it can and ignores images with pull failures")
	pullCmd.Flags().Bool("include-deps", false, "Also pull services declared as dependencies")
	pullCmd.Flags().Bool("no-parallel", false, "DEPRECATED disable parallel pulling")
	pullCmd.Flags().Bool("parallel", false, "DEPRECATED pull multiple images in parallel")
	pullCmd.Flags().String("policy", "", "Apply pull policy (\"missing\"|\"always\")")
	pullCmd.Flags().BoolP("quiet", "q", false, "Pull without printing progress information")
	pullCmd.Flag("no-parallel").Hidden = true
	pullCmd.Flag("parallel").Hidden = true
	rootCmd.AddCommand(pullCmd)

	carapace.Gen(pullCmd).FlagCompletion(carapace.ActionMap{
		"policy": carapace.ActionValues("missing", "always"),
	})

	carapace.Gen(pullCmd).PositionalAnyCompletion(
		action.ActionServices(pullCmd).FilterArgs(),
	)
}
