package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hasinstitutionalrecoverykeyCmd = &cobra.Command{
	Use:   "hasinstitutionalrecoverykey",
	Short: "Check if an institutional recovery key has been set",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hasinstitutionalrecoverykeyCmd).Standalone()
}
