package cmd

import (
	"fmt"

	"github.com/carapace-sh/carapace"
	spec "github.com/carapace-sh/carapace-spec"
	"github.com/spf13/cobra"
)

var schemaCmd = &cobra.Command{
	Use:   "--schema",
	Short: "json schema for spec files",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), spec.Schema())
	},
}

func init() {
	carapace.Gen(schemaCmd).Standalone()

}
