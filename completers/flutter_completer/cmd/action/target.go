package action

import "github.com/rsteube/carapace"

func ActionTargets() carapace.Action {
	return carapace.ActionValues(
		"copy_assets",
		"kernel_snapshot",
		"aot_elf_profile",
		"aot_elf_release",
		"aot_assembly_profile",
		"aot_assembly_release",
		"debug_macos_framework",
		"debug_macos_bundle_flutter_assets",
		"debug_bundle_linux_x64_assets",
		"debug_bundle_linux_arm64_assets",
		"profile_bundle_linux_x64_assets",
		"profile_bundle_linux_arm64_assets",
		"release_bundle_linux_x64_assets",
		"release_bundle_linux_arm64_assets",
		"web_service_worker",
		"release_android_application",
		"copy_flutter_bundle",
		"debug_android_application",
		"profile_android_application",

		// TODO flutter/packages/flutter_tools/lib/src/commands/assemble.dart
		//  // Android ABI specific AOT rules.
		//  androidArmProfileBundle,
		//  androidArm64ProfileBundle,
		//  androidx64ProfileBundle,
		//  androidArmReleaseBundle,
		//  androidArm64ReleaseBundle,
		//  androidx64ReleaseBundle,
		//  // Deferred component enabled AOT rules
		//  androidArmProfileDeferredComponentsBundle,
		//  androidArm64ProfileDeferredComponentsBundle,
		//  androidx64ProfileDeferredComponentsBundle,
		//  androidArmReleaseDeferredComponentsBundle,
		//  androidArm64ReleaseDeferredComponentsBundle,
		//  androidx64ReleaseDeferredComponentsBundle,

		"debug_ios_bundle_flutter_assets",
		"profile_ios_bundle_flutter_assets",
		"release_ios_bundle_flutter_assets",
		"unpack_windows",
		"debug_bundle_windows_assets",
		"profile_bundle_windows_assets",
		"release_bundle_windows_assets",
		"debug_bundle_windows_assets_uwp",
		"profile_bundle_windows_assets_uwp",
		"release_bundle_windows_assets_uwp",
	)
}
