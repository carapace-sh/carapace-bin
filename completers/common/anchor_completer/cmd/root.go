package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "anchor",
	Short: "Anchor CLI",
	Long:  "https://www.anchor-lang.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("commitment", "", "Commitment override (valid values: processed, confirmed, finalized)")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.PersistentFlags().String("provider.cluster", "", "Cluster override")
	rootCmd.PersistentFlags().String("provider.wallet", "", "Wallet override")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
}
