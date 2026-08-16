package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-prefetch-url",
	Short: "download a file and print its hash into the Nix store",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-prefetch-url.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("executable", false, "Set the executable bit on the downloaded file")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().String("name", "", "Override the name of the file in the Nix store")
	rootCmd.Flags().Bool("print-path", false, "Print the store path of the downloaded file")
	rootCmd.Flags().String("type", "", "Use the specified cryptographic hash algorithm")
	rootCmd.Flags().Bool("unpack", false, "Unpack the archive and add the result to the Nix store")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("md5", "sha1", "sha256", "sha512", "blake3"),
	})
}