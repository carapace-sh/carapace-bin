package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var program_deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy an upgradeable program",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(program_deployCmd).Standalone()

	program_deployCmd.Flags().String("buffer", "", "Buffer account to use for deployment")
	program_deployCmd.Flags().Bool("final", false, "Make the program immutable after deployment (cannot be upgraded)")
	program_deployCmd.Flags().BoolP("help", "h", false, "Print help")
	program_deployCmd.Flags().String("max-len", "", "Maximum transaction length (BPF loader upgradeable limit)")
	program_deployCmd.Flags().Bool("no-idl", false, "Don't upload IDL during deployment (IDL is uploaded by default)")
	program_deployCmd.Flags().String("program-id", "", "Program id to deploy to (derived from program-keypair if not specified)")
	program_deployCmd.Flags().String("program-keypair", "", "Program keypair filepath (defaults to target/deploy/{program_name}-keypair.json)")
	program_deployCmd.Flags().StringP("program-name", "p", "", "Program name to deploy (from workspace). Used when program_filepath is not provided")
	program_deployCmd.Flags().String("upgrade-authority", "", "Upgrade authority keypair (defaults to configured wallet)")
	program_deployCmd.Flags().Bool("use-rpc", false, "Send write transactions through RPC instead of TPU")
	programCmd.AddCommand(program_deployCmd)
}
