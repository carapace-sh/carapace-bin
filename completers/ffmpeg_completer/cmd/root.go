package cmd

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ffmpeg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ffmpeg",
	Short: "Hyper fast Audio and Video encoder",
	Long:  "https://ffmpeg.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	for _, prefix := range []string{
		"-c:a",
		"-c:v",
		"-b:a",
		"-b:v",
		"-disposition",
	} {
		for index, arg := range os.Args {
			if strings.HasPrefix(arg, prefix) {
				os.Args[index] = prefix // strip stream specifier
			}
		}
	}

	carapace.Override(carapace.Opts{
		LongShorthand: true,
	})
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringSlice("ab", []string{}, "audio bitrate (please use -b:a)")
	rootCmd.Flags().StringSlice("ac", []string{}, "set number of audio channels")
	rootCmd.Flags().StringSlice("acodec", []string{}, "force audio codec ('copy' to copy stream)")
	rootCmd.Flags().StringSlice("af", []string{}, "set audio filters")
	rootCmd.Flags().StringSlice("aframes", []string{}, "set the number of audio frames to output")
	rootCmd.Flags().BoolSlice("an", []bool{}, "disable audio")
	rootCmd.Flags().BoolSlice("apad", []bool{}, "audio pad")
	rootCmd.Flags().StringSlice("aq", []string{}, "set audio quality (codec-specific)")
	rootCmd.Flags().StringSlice("ar", []string{}, "set audio sampling rate (in Hz)")
	rootCmd.Flags().StringSlice("aspect", []string{}, "set aspect ratio (4:3, 16:9 or 1.3333, 1.7777)")
	rootCmd.Flags().StringSlice("b:a", []string{}, "audio bitrate")
	rootCmd.Flags().StringSlice("b:v", []string{}, "video bitrate")
	rootCmd.Flags().StringSlice("bits_per_raw_sample", []string{}, "set the number of bits per raw sample")
	rootCmd.Flags().Bool("bsfs", false, "show available bit stream filters")
	rootCmd.Flags().Bool("buildconf", false, "show build configuration")
	rootCmd.Flags().StringSlice("c", []string{}, "codec name")
	rootCmd.Flags().StringSlice("c:a", []string{}, "audio codec")
	rootCmd.Flags().StringSlice("c:v", []string{}, "video codec")
	rootCmd.Flags().StringSlice("canvas_size", []string{}, "set canvas size (WxH or abbreviation)")
	rootCmd.Flags().StringSlice("codec", []string{}, "codec name")
	rootCmd.Flags().Bool("codecs", false, "show available codecs")
	rootCmd.Flags().Bool("colors", false, "show available color names")
	rootCmd.Flags().Bool("decoders", false, "show available decoders")
	rootCmd.Flags().Bool("demuxers", false, "show available demuxers")
	rootCmd.Flags().Bool("devices", false, "show available devices")
	rootCmd.Flags().BoolSlice("discard", []bool{}, "discard")
	rootCmd.Flags().BoolSlice("disposition", []bool{}, "disposition")
	rootCmd.Flags().BoolSlice("dn", []bool{}, "disable data")
	rootCmd.Flags().Bool("encoders", false, "show available encoders")
	rootCmd.Flags().StringSlice("f", []string{}, "Force inputor output file format")
	rootCmd.Flags().StringSlice("filter", []string{}, "set stream filtergraph")
	rootCmd.Flags().Bool("filter_complex_threads", false, "number of threads for -filter_complex")
	rootCmd.Flags().StringSlice("filter_script", []string{}, "read stream filtergraph description from a file")
	rootCmd.Flags().Bool("filter_threads", false, "number of non-complex filter threads")
	rootCmd.Flags().Bool("filters", false, "show available filters")
	rootCmd.Flags().BoolSlice("fix_sub_duration", []bool{}, "fix subtitles duration")
	rootCmd.Flags().Bool("formats", false, "show available formats")
	rootCmd.Flags().StringSlice("fpsmax", []string{}, "set max frame rate (Hz value, fraction or abbreviation)")
	rootCmd.Flags().StringSlice("frames", []string{}, "set the number of frames to output")
	rootCmd.Flags().StringSlice("fs", []string{}, "set the limit file size in bytes")
	rootCmd.Flags().String("help", "", "show help")
	rootCmd.Flags().Bool("hide_banner", false, "Suppress printing banner")
	rootCmd.Flags().StringSlice("hwaccel", []string{}, "use HW acceleration")
	rootCmd.Flags().Bool("hwaccels", false, "show available HW acceleration methods")
	rootCmd.Flags().StringSlice("i", []string{}, "input file")
	rootCmd.Flags().Bool("ignore_unknown", false, "Ignore unknown stream types")
	rootCmd.Flags().Bool("layouts", false, "show standard channel layouts")
	rootCmd.Flags().String("loglevel", "", "set logging level")
	rootCmd.Flags().StringSlice("map_metadata", []string{}, "set metadata information of outfile from infile")
	rootCmd.Flags().String("max_alloc", "", "set maximum size of a single allocated block")
	rootCmd.Flags().StringSlice("metadata", []string{}, "add metadata")
	rootCmd.Flags().Bool("muxers", false, "show available muxers")
	rootCmd.Flags().Bool("n", false, "Do not overwrite output files")
	rootCmd.Flags().StringSlice("pass", []string{}, "select the pass number (1 to 3)")
	rootCmd.Flags().Bool("pix_fmts", false, "show available pixel formats")
	rootCmd.Flags().StringSlice("pre", []string{}, "preset name")
	rootCmd.Flags().StringSlice("program", []string{}, "add program with specified streams")
	rootCmd.Flags().Bool("protocols", false, "show available protocols")
	rootCmd.Flags().BoolSlice("reinit_filter", []bool{}, "reinit filtergraph on input parameter changes")
	rootCmd.Flags().Bool("report", false, "generate a report")
	rootCmd.Flags().Bool("sample_fmts", false, "show available audio sample formats")
	rootCmd.Flags().StringSlice("scodec", []string{}, "force subtitle codec ('copy' to copy stream)")
	rootCmd.Flags().BoolSlice("seek_timestamp", []bool{}, "enable/disable seeking by timestamp with -ss")
	rootCmd.Flags().String("sinks", "", "list sinks of the output device")
	rootCmd.Flags().BoolSlice("sn", []bool{}, "disable subtitle")
	rootCmd.Flags().String("sources", "", "list sources of the input device")
	rootCmd.Flags().StringSlice("spre", []string{}, "set the subtitle options to the indicated preset")
	rootCmd.Flags().StringSlice("ss", []string{}, "set the start time offset")
	rootCmd.Flags().StringSlice("sseof", []string{}, "set the start time offset relative to EOF")
	rootCmd.Flags().StringSlice("stag", []string{}, "force subtitle tag/fourcc")
	rootCmd.Flags().Bool("stats", false, "print progress report during encoding")
	rootCmd.Flags().StringSlice("target", []string{}, "specify target file type (\"vcd\", \"svcd\", \"dvd\", \"dv\" or \"dv50\" with optional prefixes \"pal-\", \"ntsc-\" or \"film-\")")
	rootCmd.Flags().StringSlice("timecode", []string{}, "set initial TimeCode value.")
	rootCmd.Flags().StringSlice("timestamp", []string{}, "set the recording timestamp ('now' to set the current time)")
	rootCmd.Flags().StringSlice("to", []string{}, "record or transcode stop time")
	rootCmd.Flags().StringSlice("vcodec", []string{}, "force video codec ('copy' to copy stream)")
	rootCmd.Flags().Bool("version", false, "show version")
	rootCmd.Flags().StringSlice("vf", []string{}, "set video filters")
	rootCmd.Flags().StringSlice("vframes", []string{}, "set the number of video frames to output")
	rootCmd.Flags().BoolSlice("vn", []bool{}, "disable video")
	rootCmd.Flags().StringSlice("vol", []string{}, "change audio volume (256=normal)")
	rootCmd.Flags().Bool("y", false, "Overwrite output files without asking")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"ab":     carapace.ActionValues("16", "32", "64", "128", "192", "256", "320"),
		"acodec": action.ActionCodecs(), // TODO only audio
		"af": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return action.ActionFilters()
				default:
					return carapace.ActionValues()
				}
			})
		}),
		"ar":      carapace.ActionValues("22050", "44100", "48000"),
		"c":       action.ActionCodecs(),
		"c:a":     action.ActionCodecs(),
		"c:v":     action.ActionCodecs(),
		"codec":   action.ActionCodecs(),
		"f":       action.ActionFormats(),
		"help":    action.ActionHelpTopics(),
		"hwaccel": action.ActionHwAccelerations(),
		"i":       carapace.ActionFiles(),
		"loglevel": carapace.ActionValuesDescribed(
			"quiet", "Show nothing at all; be silent.",
			"panic", "Only show fatal errors which could lead the process to crash",
			"fatal", "Only show fatal errors.",
			"error", "Show all errors, including ones which can be recovered from.",
			"warning", "Show all warnings and errors.",
			"info", "Show informative messages during processing.",
			"verbose", "Same as \"info\", except more verbose.",
			"debug", "Show everything, including debugging information.",
			"trace", "",
		),
		"sinks": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionDevices()
			default:
				return carapace.ActionValues()
			}
		}),
		"sources": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionDevices()
			default:
				return carapace.ActionValues()
			}
		}),
		"vcodec": action.ActionCodecs(), // TODO only video
		"vf": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return action.ActionFilters()
				default:
					return carapace.ActionValues()
				}
			})
		}),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
