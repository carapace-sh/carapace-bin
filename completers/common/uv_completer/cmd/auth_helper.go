package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_helperCmd = &cobra.Command{
	Use:    "helper",
	Short:  "Act as a credential helper for external tools",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_helperCmd).Standalone()

	auth_helperCmd.Flags().String("protocol", "", "The credential helper protocol to use")
	auth_helperCmd.MarkFlagRequired("protocol")
	authCmd.AddCommand(auth_helperCmd)
	carapace.Gen(auth_helperCmd).FlagCompletion(carapace.ActionMap{
		"protocol": carapace.ActionValuesDescribed("bazel", "Bazel credential helper protocol as described in [the spec](https://github.com/bazelbuild/proposals/blob/main/designs/2022-06-07-bazel-credential-helpers.md)"),
	})
}
