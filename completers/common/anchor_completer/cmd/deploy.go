package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:    "deploy",
	Short:  "Deploys each program in the workspace",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deployCmd).Standalone()

	deployCmd.Flags().BoolP("help", "h", false, "Print help")
	deployCmd.Flags().Bool("no-idl", false, "Don't upload IDL during deployment (IDL is uploaded by default)")
	deployCmd.Flags().String("program-keypair", "", "Keypair of the program (filepath) (requires program-name)")
	deployCmd.Flags().StringP("program-name", "p", "", "Only deploy this program")
	deployCmd.Flags().BoolP("verifiable", "v", false, "If true, deploy from path target/verifiable")
	rootCmd.AddCommand(deployCmd)
}
