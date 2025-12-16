package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/ffmpeg"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "ffmpeg",
	Short:              "Hyper fast Audio and Video encoder",
	Long:               "https://ffmpeg.org/",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.HasPrefix(c.Value, "-") {
				return actionFlags()
			}

			if len(c.Args) > 0 {
				if previous := c.Args[len(c.Args)-1]; strings.HasPrefix(previous, "-") {
					return actionFlagArguments(previous)
				}
			}
			return carapace.ActionFiles()
		}),
	)
}

func actionFlags() carapace.Action {
	// TODO highlighting, colon suffix where needed, complex flag completion
	return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.Batch(
				carapace.ActionValuesDescribed(
					"-ab", "audio bitrate",
					"-acodec", "force audio codec",
					"-af", "set audio filters",
					"-ar", "set audio sampling rate",
					"-c", "codec name",
					"-codec", "codec name",
					"-f", "force format",
					"-h", "show help",
					"-i", "input file",
					"-hwaccels", "show available HW acceleration methods",
					"-?", "show help",
					"-help", "show help",
					"-loglevel", "set logging level",
					"--help", "show help",
					"-scodec", "force subtitle codec",
					"-sinks", "list sinks of the output device",
					"-sources", "list sources of the input device",
					"-vcodec", "force video codec",
					"-vf", "set video filters",
				).Style(style.Carapace.FlagArg),
				carapace.ActionValuesDescribed(
					"-b", "bitrate",
				).Style(style.Carapace.FlagOptArg),
				carapace.ActionValuesDescribed(
					"-L", "show license",
					"-version", "show version",
					"-buildconf", "show build configuration",
					"-formats", "show available formats",
					"-muxers", "show available muxers",
					"-demuxers", "show available demuxers",
					"-devices", "show available devices",
					"-codecs", "show available codecs",
					"-decoders", "show available decoders",
					"-encoders", "show available encoders",
					"-bsfs", "show available bit stream filters",
					"-protocols", "show available protocols",
					"-filters", "show available filters",
					"-pix_fmts", "show available pixel formats",
					"-layouts", "show standard channel layouts",
					"-sample_fmts", "show available audio sample formats",
					"-dispositions", "show available stream dispositions",
					"-colors", "show available color names",
					"-v", "set logging level",
					"-report", "generate a report",
					"-max_alloc", "set maximum size of a single allocated block",
					"-y", "overwrite output files",
					"-n", "never overwrite output files",
					"-ignore_unknown", "Ignore unknown stream types",
					"-filter_threads", "number of non-complex filter threads",
					"-filter_complex_threads", "number of threads for -filter_complex",
					"-stats", "print progress report during encoding",
					"-max_error_rate", "ratio of decoding errors above which ffmpeg returns an error instead of success",
					"-cpuflags", "force specific cpu flags",
					"-cpucount", "force specific cpu count",
					"-hide_banner", "do not show program banner",
					"-copy_unknown", "Copy unknown stream types",
					"-recast_media", "allow recasting stream type in order to force a decoder of different media type",
					"-benchmark", "add timings for benchmarking",
					"-benchmark_all", "add timings for each task",
					"-progress", "write program-readable progress information",
					"-stdin", "enable or disable interaction on standard input",
					"-timelimit", "set max runtime in seconds in CPU user time",
					"-dump", "dump each input packet",
					"-hex", "when dumping packets, also dump the payload",
					"-vsync", "set video sync method globally; deprecated, use -fps_mode",
					"-frame_drop_threshold", "frame drop threshold",
					"-adrift_threshold", "audio drift threshold",
					"-copyts", "copy timestamps",
					"-start_at_zero", "shift input timestamps to start at 0 when using copyts",
					"-copytb", "copy input stream time base when stream copying",
					"-dts_delta_threshold", "timestamp discontinuity delta threshold",
					"-dts_error_threshold", "timestamp error delta threshold",
					"-xerror", "exit on error",
					"-abort_on", "abort on the specified condition flags",
					"-filter_complex", "create a complex filtergraph",
					"-lavfi", "create a complex filtergraph",
					"-filter_complex_script", "read complex filtergraph description from a file",
					"-auto_conversion_filters", "enable automatic conversion filters globally",
					"-stats_period", "set the period at which ffmpeg updates stats and -progress output",
					"-debug_ts", "print timestamp debugging info",
					"-psnr", "calculate PSNR of compressed frames",
					"-vstats", "dump video coding statistics to file",
					"-vstats_file", "dump video coding statistics to file",
					"-vstats_version", "Version of the vstats format to use.",
					"-qphist", "show QP histogram",
					"-sdp_file", "specify a file in which to print sdp information",
					"-vaapi_device", "set VAAPI hardware device",
					"-qsv_device", "set QSV hardware device",
					"-init_hw_device", "initialise hardware device",
					"-filter_hw_device", "set hardware device used when filtering",
					"-pre", "preset name",
					"-map_metadata", "set metadata information of outfile from infile",
					"-t", "record or transcode \"duration\" seconds of audio/video",
					"-to", "record or transcode stop time",
					"-fs", "set the limit file size in bytes",
					"-ss", "set the start time offset",
					"-sseof", "set the start time offset relative to EOF",
					"-seek_timestamp", "enable/disable seeking by timestamp with -ss",
					"-timestamp", "set the recording timestamp",
					"-metadata", "add metadata",
					"-program", "add program with specified streams",
					"-target", "specify target file type",
					"-apad", "audio pad",
					"-frames", "set the number of frames to output",
					"-filter", "set stream filtergraph",
					"-filter_script", "read stream filtergraph description from a file",
					"-reinit_filter", "reinit filtergraph on input parameter changes",
					"-discard", "discard",
					"-disposition", "disposition",
					"-map", "set input stream mapping",
					"-map_channel", "map an audio channel from one stream to another",
					"-map_chapters", "set chapters mapping",
					"-accurate_seek", "enable/disable accurate seeking with -ss",
					"-isync", "ref     Indicate the input index for sync reference",
					"-itsoffset", "set the input ts offset",
					"-itsscale", "set the input ts scale",
					"-dframes", "set the number of data frames to output",
					"-re", "read input at native frame rate; equivalent to -readrate 1",
					"-readrate", "read input at specified rate",
					"-shortest", "finish encoding within shortest input",
					"-shortest_buf_duration", "maximum buffering duration for the -shortest option",
					"-bitexact", "bitexact mode",
					"-copyinkf", "copy initial non-keyframes",
					"-copypriorss", "copy or discard frames before start time",
					"-tag", "force codec tag/fourcc",
					"-q", "use fixed quality scale",
					"-qscale", "use fixed quality scale",
					"-profile", "set profile",
					"-attach", "add an attachment to the output file",
					"-dump_attachment", "extract an attachment into a file",
					"-stream_loop", "count  set number of times input stream shall be looped",
					"-thread_queue_size", "set the maximum number of queued packets from the demuxer",
					"-find_stream_info", "read and decode the streams to fill missing information with heuristics",
					"-bits_per_raw_sample", "set the number of bits per raw sample",
					"-stats_enc_pre", "write encoding stats before encoding",
					"-stats_enc_post", "write encoding stats after encoding",
					"-stats_mux_pre", "write packets stats before muxing",
					"-stats_enc_pre_fmt", "format of the stats written with -stats_enc_pre",
					"-stats_enc_post_fmt", "format of the stats written with -stats_enc_post",
					"-stats_mux_pre_fmt", "format of the stats written with -stats_mux_pre",
					"-autorotate", "automatically insert correct rotate filters",
					"-autoscale", "automatically insert a scale filter at the end of the filter graph",
					"-muxdelay", "set the maximum demux-decode delay",
					"-muxpreload", "set the initial demux-decode delay",
					"-time_base", "set the desired time base hint for output stream",
					"-enc_time_base", "set the desired time base for the encoder,-1 = match source time base",
					"-bsf", "A comma-separated list of bitstream filters",
					"-fpre", "set options from indicated preset file",
					"-max_muxing_queue_size", "maximum number of packets that can be buffered while waiting for all streams to initialize",
					"-muxing_queue_data_threshold", "set the threshold after which max_muxing_queue_size is taken into account",
					"-dcodec", "force data codec",
					"-vframes", "set the number of video frames to output",
					"-r", "set frame rate",
					"-fpsmax", "set max frame rate",
					"-s", "set frame size",
					"-aspect", "set aspect ratio",
					"-display_rotation", "set pure counter-clockwise rotation in degrees for strea",
					"-display_hflip", "set display horizontal flip for strea",
					"-display_vflip", "set display vertical flip for strea",
					"-vn", "disable video",
					"-timecode", "set initial TimeCode value",
					"-pass", "select the pass number",
					"-dn", "disable data",
					"-pix_fmt", "set pixel format",
					"-rc_override", "rate control override for specific intervals",
					"-passlogfile", "select two pass log file name prefix",
					"-psnr", "calculate PSNR of compressed frames",
					"-vstats", "dump video coding statistics to file",
					"-vstats_file", "dump video coding statistics to file",
					"-vstats_version", "Version of the vstats format to use.",
					"-intra_matrix", "specify intra matrix coeffs",
					"-inter_matrix", "specify inter matrix coeffs",
					"-chroma_intra_matrix", "specify intra matrix coeffs",
					"-top", "top=1/bottom=0/auto=-1 field first",
					"-vtag", "force video tag/fourcc",
					"-qphist", "show QP histogram",
					"-fps_mode", "set framerate mode for matching video streams; overrides vsync",
					"-force_fps", "force the selected framerate, disable the best supported framerate selection",
					"-streamid", "set the value of an outfile streamid",
					"-force_key_frames", "force key frames at specified timestamps",
					"-hwaccel", "name  use HW accelerated decoding",
					"-hwaccel_device", "select a device for HW acceleration",
					"-hwaccel_output_format", "select output format used with HW accelerated decoding",
					"-fix_sub_duration_heartbeat", "set this video output stream to be a heartbeat stream for fix_sub_duration",
					"-vbsf", "bitstream_filters  deprecated",
					"-vpre", "set the video options to the indicated preset",
					"-aframes", "set the number of audio frames to output",
					"-aq", "set audio quality",
					"-ac", "set number of audio channels",
					"-an", "disable audio",
					"-atag", "force audio tag/fourcc",
					"-sample_fmt", "set sample format",
					"-channel_layout", "set channel layout",
					"-ch_layout", "set channel layout",
					"-guess_layout_max", "set the maximum number of channels to try to guess the channel layout",
					"-absf", "bitstream_filters  deprecated",
					"-apre", "set the audio options to the indicated preset",
					"-s", "set frame size",
					"-sn", "disable subtitle",
					"-stag", "force subtitle tag/fourcc",
					"-fix_sub_duration", "fix subtitles duration",
					"-canvas_size", "set canvas size",
					"-spre", "set the subtitle options to the indicated preset",
				).Style(style.Carapace.FlagNoArg),
			).ToA()
		default:
			switch c.Parts[0] {
			case "-b":
				return carapace.ActionValuesDescribed(
					"a", "audio",
					"v", "video",
				)
			case "-c", "-codec":
				return carapace.ActionValuesDescribed(
					"a", "audio",
					"v", "video",
					"s", "subtitle",
				)
			default:
				return carapace.ActionValues()
			}
		}
	}).Tag("flags")
}

func actionFlagArguments(flag string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		splitted := strings.Split(strings.TrimLeft(flag, "-"), ":")
		beforeInput := true
		for _, arg := range c.Args {
			if arg == "-i" {
				beforeInput = false
				break
			}
		}

		switch splitted[0] {
		case "ab":
			return carapace.ActionValues("16", "32", "64", "128", "192", "256", "320")
		case "acodec":
			if beforeInput {
				return carapace.Batch(
					ffmpeg.ActionDecodableCodecs(ffmpeg.CodecOpts{Audio: true}),
					ffmpeg.ActionDecoders(ffmpeg.DecoderOpts{Audio: true}),
				).ToA()
			}
			return carapace.Batch(
				ffmpeg.ActionEncodableCodecs(ffmpeg.CodecOpts{Audio: true}),
				ffmpeg.ActionEncoders(ffmpeg.EncoderOpts{Audio: true}),
			).ToA()
		case "af":
			return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return ffmpeg.ActionFilters().NoSpace()
					default:
						return carapace.ActionValues()
					}
				})
			})
		case "ar":
			return carapace.ActionValues("22050", "44100", "48000")
		case "b":
			if len(splitted) > 1 {
				switch splitted[1] {
				case "a":
					return carapace.ActionValues("16", "32", "64", "128", "192", "256", "320")
				case "v":
					return carapace.ActionValues() // video bitrate
				}
			}
			return carapace.ActionValues() // TODO invalid flag (missing a/v))
		case "c", "codec":
			audio := true
			subtitle := true
			video := true
			if len(splitted) > 1 {
				audio = splitted[1] == "a"
				subtitle = splitted[1] == "s"
				video = splitted[1] == "v"
			}
			if beforeInput {
				return carapace.Batch(
					ffmpeg.ActionDecodableCodecs(ffmpeg.CodecOpts{Audio: audio, Subtitle: subtitle, Video: video}),
					ffmpeg.ActionDecoders(ffmpeg.DecoderOpts{Audio: audio, Subtitle: subtitle, Video: video}),
				).ToA()
			}
			return carapace.Batch(
				ffmpeg.ActionEncodableCodecs(ffmpeg.CodecOpts{Audio: audio, Subtitle: subtitle, Video: video}),
				ffmpeg.ActionEncoders(ffmpeg.EncoderOpts{Audio: audio, Subtitle: subtitle, Video: video}),
			).ToA()
		case "f":
			return ffmpeg.ActionFormats()
		case "h", "?", "help":
			return ffmpeg.ActionHelpTopics()
		case "hwaccel":
			return ffmpeg.ActionHardwareAccelerations()
		case "i":
			return carapace.ActionFiles()
		case "loglevel":
			return ffmpeg.ActionLogLevels()
		case "scodec":
			if beforeInput {
				return carapace.Batch(
					ffmpeg.ActionDecodableCodecs(ffmpeg.CodecOpts{Subtitle: true}),
					ffmpeg.ActionDecoders(ffmpeg.DecoderOpts{Subtitle: true}),
				).ToA()
			}
			return carapace.Batch(
				ffmpeg.ActionEncodableCodecs(ffmpeg.CodecOpts{Subtitle: true}),
				ffmpeg.ActionEncoders(ffmpeg.EncoderOpts{Subtitle: true}),
			).ToA()
		case "sinks":
			return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return ffmpeg.ActionDevices(ffmpeg.DeviceOpts{Demuxing: true}).NoSpace()
				default:
					return carapace.ActionValues()
				}
			})
		case "sources":
			return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return ffmpeg.ActionDevices(ffmpeg.DeviceOpts{Muxing: true}).NoSpace()
				default:
					return carapace.ActionValues()
				}
			})
		case "vcodec":
			if beforeInput {
				return carapace.Batch(
					ffmpeg.ActionDecodableCodecs(ffmpeg.CodecOpts{Video: true}),
					ffmpeg.ActionDecoders(ffmpeg.DecoderOpts{Video: true}),
				).ToA()
			}
			return carapace.Batch(
				ffmpeg.ActionEncodableCodecs(ffmpeg.CodecOpts{Video: true}),
				ffmpeg.ActionEncoders(ffmpeg.EncoderOpts{Video: true}),
			).ToA()
		case "vf":
			return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return ffmpeg.ActionFilters().NoSpace()
					default:
						return carapace.ActionValues()
					}
				})
			})
		default:
			//return carapace.ActionValues() // TODO
			return carapace.ActionFiles() // default file completion for now (positional)
		}
	})
}
