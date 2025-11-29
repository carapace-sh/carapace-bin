package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_credentialsCmd = &cobra.Command{
	Use:    "credentials",
	Short:  "Output S3 cache credentials",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_credentialsCmd).Standalone()

	cache_credentialsCmd.Flags().String("format", "", "Output format, either json or sh")
	cacheCmd.AddCommand(cache_credentialsCmd)

	carapace.Gen(cache_credentialsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "sh"),
	})
}
