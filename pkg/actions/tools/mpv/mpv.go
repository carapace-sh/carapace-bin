package mpv

import (
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionAudioOutputs completes audio outputs
//
//	alsa (ALSA audio output)
//	jack (JACK audio output)
func ActionAudioOutputs() carapace.Action {
	return carapace.ActionExecCommand("mpv", "--ao=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^\s+(?P<name>\S+)\s+(?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionVideoOutputs completes video outputs
//
//	dmabuf-wayland (Wayland dmabuf video output)
//	drm (Direct Rendering Manager (software scaling))
func ActionVideoOutputs() carapace.Action {
	return carapace.ActionExecCommand("mpv", "--vo=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^\s+(?P<name>\S+)\s+(?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionAudioChannels completes audio channels
//
//	1.0 (fc)
//	2.0 (fl-fr)
func ActionAudioChannels() carapace.Action {
	return carapace.ActionExecCommand("mpv", "--audio-channels=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^\s*(?P<name>\S+)(?:\s+\((?P<description>.*)\))?$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionDemuxers completes demuxers
//
//	cue (CUE sheet)
//	directory (Playlist dir)
func ActionDemuxers() carapace.Action {
	return carapace.ActionExecCommand("mpv", "--audio-demuxer=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^\s*(?P<name>\S+)\s+(?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines[2:] {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionGPUContexts completes GPU contexts
//
//	auto (Auto detect)
//	displayvk (VK_KHR_display)
func ActionGPUContexts() carapace.Action {
	return carapace.ActionExecCommand("mpv", "--gpu-context=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^\s+(?P<name>\S+)\s+(?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionAudioCodecs completes audio codecs
//
//	aac (AAC (Advanced Audio Coding))
//	ac3 (ATSC A/52A (AC-3))
func ActionAudioCodecs() carapace.Action {
	return carapace.ActionExecCommand("mpv", "--oac=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^\s*--oac=(?P<name>\S+)\s+(?P<description>.+)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionAudioCodecOptions completes audio codec options for libavcodec
//
//	aac_coder (Coding algorithm)
//	aac_is (Intensity stereo coding)
func ActionAudioCodecOptions() carapace.Action {
	return carapace.ActionExecCommand("mpv", "--oacopts=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^\s*--oacopts=(?P<name>\w+)(?:=<[^>\s]+>)?\s+(?P<description>.+)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionVideoCodecs completes video codecs
//
//	a64multi (Multicolor charset for Commodore 64)
//	a64multi5 (Multicolor charset for Commodore 64, extended with 5th color (colram))

func ActionVideoCodecs() carapace.Action {
	return carapace.ActionExecCommand("mpv", "--ovc=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^\s*--ovc=(?P<name>\S+)\s+(?P<description>.+)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionVideoCodecOptions completes video codec options for libavcodec
//
//	2pass (Use 2pass encoding mode)
//	8x8dct (High profile 8x8 transform.)
func ActionVideoCodecOptions() carapace.Action {
	return carapace.ActionExecCommand("mpv", "--ovcopts=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^\s*--ovcopts=(?P<name>\w+)(?:=<[^>\s]+>)?\s+(?P<description>.+)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionFormats completes output formats
//
//	3g2 (3GP2 (3GPP2 file format))
//	3gp (3GP (3GPP file format)
func ActionFormats() carapace.Action {
	return carapace.ActionExecCommand("mpv", "--of=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^\s*--of=(?P<name>\S+)\s+(?P<description>.+)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionFormatOptions completes format options for libavformat
//
//	adaptation_sets (Adaptation sets. Syntax: id=0,streams=0,1,2 id=1,streams=3,4 and so on)
//	allow_raw_vfw (allow raw VFW mode)
func ActionFormatOptions() carapace.Action {
	return carapace.ActionExecCommand("mpv", "--ofopts=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^\s*--ofopts=(?P<name>\w+)(?:=<[^>\s]+>)?\s+(?P<description>.+)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionProfiles completes profiles
//
//	enc-a-aac (AAC (libfdk-aac or FFmpeg))
//	enc-a-ac3 (AC3 (FFmpeg))
func ActionProfiles() carapace.Action {
	return carapace.ActionExecCommand("mpv", "--profile=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^\s+(?P<name>\S+)\s+(?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
