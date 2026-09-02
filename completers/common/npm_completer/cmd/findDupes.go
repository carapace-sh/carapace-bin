package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var findDupesCmd = &cobra.Command{
	Use:   "find-dupes",
	Short: "Find duplication in the package tree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(findDupesCmd).Standalone()
	findDupesCmd.Flags().Bool("audit", false, "submit audit reports")
	findDupesCmd.Flags().Bool("bin-links", false, "create symlinks")
	findDupesCmd.Flags().Bool("fund", false, "Display funding message")
	findDupesCmd.Flags().Bool("global-style", false, "Use global layout")
	findDupesCmd.Flags().Bool("ignore-scripts", false, "Disable scripts")
	findDupesCmd.Flags().String("include", "", "include dependency types")
	findDupesCmd.Flags().String("install-strategy", "", "strategy for installing packages in node_modules")
	findDupesCmd.Flags().Bool("legacy-bundling", false, "Use legacy bundling")
	findDupesCmd.Flags().Bool("no-audit", false, "skip audit")
	findDupesCmd.Flags().Bool("no-bin-links", false, "skip symlinks")
	findDupesCmd.Flags().Bool("no-fund", false, "skip funding message")
	findDupesCmd.Flags().Bool("no-package-lock", false, "ignore package-lock.json")
	findDupesCmd.Flags().StringArray("omit", nil, "omit dependency type")
	findDupesCmd.Flags().Bool("package-lock", false, "ignore package-lock.json")
	findDupesCmd.Flags().Bool("strict-peer-deps", false, "Fail and abort for any conflicting `peerDependencies`")
	addWorkspaceFlags(findDupesCmd)

	rootCmd.AddCommand(findDupesCmd)

	carapace.Gen(findDupesCmd).FlagCompletion(carapace.ActionMap{
		"include":          carapace.ActionValues("prod", "dev", "optional", "peer"),
		"install-strategy": carapace.ActionValues("hoisted", "nested", "shallow", "linked"),
		"omit":             carapace.ActionValues("dev", "optional", "peer"),
	})
}
