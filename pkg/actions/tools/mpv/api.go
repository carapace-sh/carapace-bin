package mpv

import "github.com/rsteube/carapace"

// ActionHardwareDecodingAPIs completes hardware decoding apis
//
//	auto (forcibly enable any hw decoder found)
//	auto-copy (enable best hw decoder with copy-back)
func ActionHardwareDecodingAPIs() carapace.Action {
	return carapace.ActionValuesDescribed(
		"no", "always use software decoding (default)",
		"auto", "forcibly enable any hw decoder found",
		"yes", "exactly the same as auto",
		"auto-safe", "enable any whitelisted hw decoder",
		"auto-copy", "enable best hw decoder with copy-back",
		"vdpau", "requires --vo=gpu with X11, or --vo=vdpau (Linux only)",
		"vdpau-copy", "copies video back into system RAM (Linux with some GPUs only)",
		"vaapi", "requires --vo=gpu or --vo=vaapi (Linux only)",
		"vaapi-copy", "copies video back into system RAM (Linux with some GPUs only)",
		"videotoolbox", "requires --vo=gpu (macOS 10.8 and up), or --vo=libmpv (iOS 9.0 and up)",
		"videotoolbox-copy", "copies video back into system RAM (macOS 10.8 or iOS 9.0 and up)",
		"dxva2", "requires --vo=gpu with --gpu-context=d3d11, --gpu-context=angle or --gpu-context=dxinterop (Windows only)",
		"dxva2-copy", "copies video back to system RAM (Windows only)",
		"d3d11va", "requires --vo=gpu with --gpu-context=d3d11 or --gpu-context=angle (Windows 8+ only)",
		"d3d11va-copy", "copies video back to system RAM (Windows 8+ only)",
		"mediacodec", "requires --vo=mediacodec_embed (Android only)",
		"mediacodec-copy", "copies video back to system RAM (Android only)",
		"mmal", "requires --vo=gpu (Raspberry Pi only - default if available)",
		"mmal-copy", "copies video back to system RAM (Raspberry Pi only)",
		"nvdec", "requires --vo=gpu (Any platform CUDA is available)",
		"nvdec-copy", "copies video back to system RAM (Any platform CUDA is available)",
		"cuda", "requires --vo=gpu (Any platform CUDA is available)",
		"cuda-copy", "copies video back to system RAM (Any platform CUDA is available)",
		"crystalhd", "copies video back to system RAM (Any platform supported by hardware)",
		"rkmpp", "requires --vo=gpu (some RockChip devices only)",
	)
}
