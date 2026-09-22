package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "build",
	Short:   "Builds the workspace",
	Aliases: []string{"b"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().String("arch", "v3", "SBPF architecture to pass to `cargo build-sbf`")
	buildCmd.Flags().StringP("bootstrap", "b", "none", "Bootstrap docker image from scratch, installing all requirements for verifiable builds. Only works for debian-based images")
	buildCmd.Flags().StringP("docker-image", "d", "", "Docker image to use. For --verifiable builds only")
	buildCmd.Flags().StringSliceP("env", "e", nil, "Environment variables to pass into the docker container")
	buildCmd.Flags().BoolP("help", "h", false, "Print help")
	buildCmd.Flags().StringP("idl", "i", "", "Output directory for the IDL")
	buildCmd.Flags().StringP("idl-ts", "t", "", "Output directory for the TypeScript IDL")
	buildCmd.Flags().Bool("ignore-keys", false, "Skip checking for program ID mismatch between keypair and declare_id")
	buildCmd.Flags().Bool("no-docs", false, "Suppress doc strings in IDL output")
	buildCmd.Flags().Bool("no-idl", false, "Do not build the IDL")
	buildCmd.Flags().StringP("program-name", "p", "", "Name of the program to build")
	buildCmd.Flags().Bool("skip-lint", false, "True if the build should not fail even if there are no \"CHECK\" comments")
	buildCmd.Flags().StringP("solana-version", "s", "", "Version of the Solana toolchain to use. For --verifiable builds only")
	buildCmd.Flags().String("tools-version", "v1.57", "Platform tools version to pass to `cargo build-sbf`")
	buildCmd.Flags().BoolP("verifiable", "v", false, "True if the build artifact needs to be deterministic and verifiable")
	rootCmd.AddCommand(buildCmd)
}
