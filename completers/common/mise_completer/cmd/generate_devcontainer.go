// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_devcontainerCmd = &cobra.Command{
	Use:   "devcontainer",
	Short: "Generate a devcontainer configuration that installs mise and the project's tools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_devcontainerCmd).Standalone()

	generate_devcontainerCmd.Flags().BoolP("help", "h", false, "Print help")
	generate_devcontainerCmd.Flags().StringP("image", "i", "", "Base image; defaults to mcr.microsoft.com/devcontainers/base:ubuntu")
	generate_devcontainerCmd.Flags().BoolP("mount-mise-data", "m", false, "Keep installed tools in a `mise-data-volume` Docker volume so they survive rebuilds")
	generate_devcontainerCmd.Flags().StringP("name", "n", "", "Container name; defaults to mise")
	generate_devcontainerCmd.Flags().BoolP("write", "w", false, "Write to .devcontainer/devcontainer.json instead of printing")
	generateCmd.AddCommand(generate_devcontainerCmd)
}
