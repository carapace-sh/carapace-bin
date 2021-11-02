package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_cdn_flushCmd = &cobra.Command{
	Use:   "flush",
	Short: "Flush the cache of a CDN",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_cdn_flushCmd).Standalone()
	compute_cdn_flushCmd.Flags().StringSlice("files", []string{"*"}, "cdn files")
	compute_cdnCmd.AddCommand(compute_cdn_flushCmd)

	carapace.Gen(compute_cdn_flushCmd).FlagCompletion(carapace.ActionMap{
		"files": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
	})
}
