package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_platform_addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Adds a platform(s) to the workspace file and updates the lock file",
	Aliases: []string{"a"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_platform_addCmd).Standalone()

	workspace_platform_addCmd.Flags().String("archspec", "", "Declare a `__archspec` virtual package with the given microarchitecture string, e.g. `x86-64-v3`. Valid on any subdir")
	workspace_platform_addCmd.Flags().Bool("auto-detect", false, "Detect this machine's platform (subdir and virtual packages) instead of naming a subdir. Optionally pass a single `<name>` to name it; any virtual-package flags override the detected values. The detected platform is placed at the top of the list")
	workspace_platform_addCmd.Flags().Bool("auto-detected", false, "Detect this machine's platform (subdir and virtual packages) instead of naming a subdir. Optionally pass a single `<name>` to name it; any virtual-package flags override the detected values. The detected platform is placed at the top of the list")
	workspace_platform_addCmd.Flags().String("cuda", "", "Declare a `__cuda` virtual package at the given version, e.g. `12.0`. Valid on any subdir")
	workspace_platform_addCmd.Flags().String("cuda-arch", "", "Declare a `__cuda_arch` virtual package (GPU compute capability) at the given version, e.g. `8.6`. Requires `--cuda` (or an existing `__cuda`), matching the conda CEP coupling. Serialized as `cuda = { driver, arch }`")
	workspace_platform_addCmd.Flags().Bool("current", false, "Detect this machine's platform (subdir and virtual packages) instead of naming a subdir. Optionally pass a single `<name>` to name it; any virtual-package flags override the detected values. The detected platform is placed at the top of the list")
	workspace_platform_addCmd.Flags().StringP("feature", "f", "", "The name of the feature to add the platform to")
	workspace_platform_addCmd.Flags().String("glibc", "", "Declare a `__glibc` virtual package at the given version, e.g. `2.28`. Only valid on linux subdirs")
	workspace_platform_addCmd.Flags().String("linux", "", "Declare a `__linux` virtual package at the given kernel version, e.g. `5.10`. Only valid on linux subdirs")
	workspace_platform_addCmd.Flags().String("macos", "", "Declare a `__osx` virtual package at the given macOS version, e.g. `14.0`. Only valid on osx subdirs")
	workspace_platform_addCmd.Flags().Bool("no-install", false, "Don't update the environment, only add changed packages to the lock file")
	workspace_platform_addCmd.Flags().String("osx", "", "Declare a `__osx` virtual package at the given macOS version, e.g. `14.0`. Only valid on osx subdirs")
	workspace_platform_addCmd.Flags().String("windows", "", "Declare a `__win` virtual package at the given Windows version, e.g. `10`. Only valid on win subdirs")
	workspace_platformCmd.AddCommand(workspace_platform_addCmd)

	carapace.Gen(workspace_platform_addCmd).FlagCompletion(carapace.ActionMap{
		"feature": pixi.ActionFeatures(),
	})

	carapace.Gen(workspace_platform_addCmd).PositionalAnyCompletion(
		pixi.ActionKnownPlatforms(),
	)
}
