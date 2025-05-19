package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makePolicyCmd = &cobra.Command{
	Use:   "make:policy [-f|--force] [-m|--model [MODEL]] [-g|--guard [GUARD]] [--] <name>",
	Short: "Create a new policy class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makePolicyCmd).Standalone()

	makePolicyCmd.Flags().Bool("force", false, "Create the class even if the policy already exists")
	makePolicyCmd.Flags().String("guard", "", "The guard that the policy relies on")
	makePolicyCmd.Flags().String("model", "", "The model that the policy applies to")
	rootCmd.AddCommand(makePolicyCmd)
}
