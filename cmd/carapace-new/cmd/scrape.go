package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var scrapeCmd = &cobra.Command{
	Use:   "--scrape",
	Short: "scrape spec to go code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scrapeCmd).Standalone()

	carapace.Gen(scrapeCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
