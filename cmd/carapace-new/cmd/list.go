package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		println("TODO list completers")
	},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().StringP("format", "f", "txt", "output format")

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "txt").StyleF(func(s string, sc style.Context) string {
			return style.ForPathExt("."+s, sc)
		}),
	})
}
