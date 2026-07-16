package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_platform_editCmd = &cobra.Command{
	Use:     "edit",
	Short:   "Edit an existing workspace platform's subdir and/or virtual packages",
	Aliases: []string{"e"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_platform_editCmd).Standalone()

	workspace_platform_editCmd.Flags().String("archspec", "", "Declare a `__archspec` virtual package with the given microarchitecture string, e.g. `x86-64-v3`. Valid on any subdir")
	workspace_platform_editCmd.Flags().Bool("clear-virtual-packages", false, "Clear all virtual packages before applying any add/upsert operations")
	workspace_platform_editCmd.Flags().String("cuda", "", "Declare a `__cuda` virtual package at the given version, e.g. `12.0`. Valid on any subdir")
	workspace_platform_editCmd.Flags().String("cuda-arch", "", "Declare a `__cuda_arch` virtual package (GPU compute capability) at the given version, e.g. `8.6`. Requires `--cuda` (or an existing `__cuda`), matching the conda CEP coupling. Serialized as `cuda = { driver, arch }`")
	workspace_platform_editCmd.Flags().String("glibc", "", "Declare a `__glibc` virtual package at the given version, e.g. `2.28`. Only valid on linux subdirs")
	workspace_platform_editCmd.Flags().String("linux", "", "Declare a `__linux` virtual package at the given kernel version, e.g. `5.10`. Only valid on linux subdirs")
	workspace_platform_editCmd.Flags().String("macos", "", "Declare a `__osx` virtual package at the given macOS version, e.g. `14.0`. Only valid on osx subdirs")
	workspace_platform_editCmd.Flags().Bool("no-install", false, "Don't update the environment, only refresh the lock-file")
	workspace_platform_editCmd.Flags().String("osx", "", "Declare a `__osx` virtual package at the given macOS version, e.g. `14.0`. Only valid on osx subdirs")
	workspace_platform_editCmd.Flags().StringSlice("remove-virtual-package", nil, "Remove the named virtual package from this platform. Can be repeated")
	workspace_platform_editCmd.Flags().String("subdir", "", "Set a new conda subdir for this platform")
	workspace_platform_editCmd.Flags().String("windows", "", "Declare a `__win` virtual package at the given Windows version, e.g. `10`. Only valid on win subdirs")
	workspace_platformCmd.AddCommand(workspace_platform_editCmd)
}
