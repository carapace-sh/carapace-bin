package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var onboardingCmd = &cobra.Command{
	Use:    "onboarding",
	Short:  "INTERNAL: First-run onboarding that shows metrics info and marks onboarding complete",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(onboardingCmd).Standalone()

	onboardingCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(onboardingCmd)
}
