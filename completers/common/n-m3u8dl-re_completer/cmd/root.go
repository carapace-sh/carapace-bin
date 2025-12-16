package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "n-m3u8dl-re",
	Short: "Cross-Platform, modern and powerful stream downloader for MPD/M3U8/ISM",
	Long:  "https://github.com/nilaoda/N_m3u8DL-RE",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("?", "?", false, "Show help and usage information")
	rootCmd.Flags().String("ad-keyword", "", "Set URL keywords (regular expressions) for AD segments")
	rootCmd.Flags().Bool("allow-hls-multi-ext-map", false, "Allow multiple #EXT-X-MAP in HLS (experimental)")
	rootCmd.Flags().Bool("append-url-params", false, "Add Params of input Url to segments, useful for some websites, such as kakao.com")
	rootCmd.Flags().Bool("auto-select", false, "Automatically selects the best tracks of all types")
	rootCmd.Flags().Bool("auto-subtitle-fix", false, "Automatically fix subtitles")
	rootCmd.Flags().String("base-url", "", "Set BaseURL")
	rootCmd.Flags().Bool("binary-merge", false, "Binary merge")
	rootCmd.Flags().Bool("check-segments-count", false, "Check if the actual number of segments downloaded matches the expected number")
	rootCmd.Flags().BoolP("concurrent-download", "mt", false, "Concurrently download the selected audio, video and subtitles")
	rootCmd.Flags().String("custom-hls-iv", "", "Set the HLS decryption iv")
	rootCmd.Flags().String("custom-hls-key", "", "Set the HLS decryption key")
	rootCmd.Flags().String("custom-hls-method", "", "Set HLS encryption method")
	rootCmd.Flags().String("custom-proxy", "", "Set web request proxy")
	rootCmd.Flags().String("custom-range", "", "Download only part of the segments")
	rootCmd.Flags().String("decryption-binary-path", "", "Full path to the tool used for MP4 decryption")
	rootCmd.Flags().String("decryption-engine", "", "Set the third-party program used for decryption")
	rootCmd.Flags().Bool("del-after-done", false, "Delete temporary files when done")
	rootCmd.Flags().Bool("disable-update-check", false, "Disable version update check")
	rootCmd.Flags().String("download-retry-count", "", "The number of retries when download segment error")
	rootCmd.Flags().StringP("drop-audio", "da", "", "Drop audio streams by regular expressions")
	rootCmd.Flags().StringP("drop-subtitle", "ds", "", "Drop subtitle streams by regular expressions")
	rootCmd.Flags().StringP("drop-video", "dv", "", "Drop video streams by regular expressions")
	rootCmd.Flags().String("ffmpeg-binary-path", "", "Full path to the ffmpeg binary")
	rootCmd.Flags().Bool("force-ansi-console", false, "Force assuming the terminal is ANSI-compatible and interactive")
	rootCmd.Flags().StringP("header", "H", "", "Pass custom header(s) to server")
	rootCmd.Flags().BoolP("help", "h", false, "Show help and usage information")
	rootCmd.Flags().String("http-request-timeout", "", "Timeout duration for HTTP requests (in seconds)")
	rootCmd.Flags().StringSlice("key", nil, "Set decryption key(s) to mp4decrypt/shaka-packager/ffmpeg")
	rootCmd.Flags().String("key-text-file", "", "Set the kid-key file, the program will search the KEY with KID from the file")
	rootCmd.Flags().Bool("live-fix-vtt-by-audio", false, "Correct VTT sub by reading the start time of the audio file")
	rootCmd.Flags().Bool("live-keep-segments", false, "Keep segments when recording a live (liveRealTimeMerge enabled)")
	rootCmd.Flags().Bool("live-perform-as-vod", false, "Download live streams as vod")
	rootCmd.Flags().Bool("live-pipe-mux", false, "Real-time muxing to TS file through pipeline + ffmpeg (liveRealTimeMerge enabled)")
	rootCmd.Flags().Bool("live-real-time-merge", false, "Real-time merge into file when recording live")
	rootCmd.Flags().String("live-record-limit", "", "Recording time limit when recording live")
	rootCmd.Flags().String("live-take-count", "", "Manually set the number of segments downloaded for the first time when recording live")
	rootCmd.Flags().String("live-wait-time", "", "Manually set the live playlist refresh interval")
	rootCmd.Flags().String("log-level", "", "Set log level")
	rootCmd.Flags().StringP("max-speed", "R", "", "Set speed limit")
	rootCmd.Flags().String("morehelp", "", "Set more help info about one option")
	rootCmd.Flags().Bool("mp4-real-time-decryption", false, "Decrypt MP4 segments in real time")
	rootCmd.Flags().StringP("mux-after-done", "M", "", "When all works is done, try to mux the downloaded streams")
	rootCmd.Flags().StringSlice("mux-import", nil, "When MuxAfterDone enabled, allow to import local media files")
	rootCmd.Flags().Bool("no-ansi-color", false, "Remove ANSI colors")
	rootCmd.Flags().Bool("no-date-info", false, "Date information is not written during muxing")
	rootCmd.Flags().Bool("no-log", false, "Disable log file output")
	rootCmd.Flags().String("save-dir", "", "Set output directory")
	rootCmd.Flags().String("save-name", "", "Set output filename")
	rootCmd.Flags().StringP("select-audio", "sa", "", "Select audio streams by regular expressions")
	rootCmd.Flags().StringP("select-subtitle", "ss", "", "Select subtitle streams by regular expressions")
	rootCmd.Flags().StringP("select-video", "sv", "", "Select video streams by regular expressions")
	rootCmd.Flags().Bool("skip-download", false, "Skip download")
	rootCmd.Flags().Bool("skip-merge", false, "Skip segments merge")
	rootCmd.Flags().String("sub-format", "", "Subtitle output format")
	rootCmd.Flags().Bool("sub-only", false, "Select only subtitle tracks")
	rootCmd.Flags().String("task-start-at", "", "Task execution will not start before this time")
	rootCmd.Flags().String("thread-count", "", "Set download thread count")
	rootCmd.Flags().String("tmp-dir", "", "Set temporary file directory")
	rootCmd.Flags().String("ui-language", "", "Set UI language")
	rootCmd.Flags().String("urlprocessor-args", "", "Give these arguments to the URL Processors.")
	rootCmd.Flags().Bool("use-ffmpeg-concat-demuxer", false, "When merging with ffmpeg, use the concat demuxer instead of the concat protocol")
	rootCmd.Flags().Bool("use-system-proxy", false, "Use system default proxy")
	rootCmd.Flags().Bool("version", false, "Show version information")
	rootCmd.Flags().Bool("write-meta-json", false, "Write meta json after parsed")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"custom-hls-method":      carapace.ActionValues("AES_128", "AES_128_ECB", "CENC", "CHACHA20", "NONE", "SAMPLE_AES", "SAMPLE_AES_CTR", "UNKNOWN"),
		"decryption-binary-path": carapace.ActionFiles(),
		"decryption-engine":      carapace.ActionValues("FFMPEG", "MP4DECRYPT", "SHAKA_PACKAGER"),
		"log-level":              carapace.ActionValues("DEBUG", "ERROR", "INFO", "OFF", "WARN").StyleF(style.ForLogLevel),
		"morehelp":               carapace.ActionValues("custom-range", "mux-after-done", "mux-import", "select-audio", "select-subtitle", "select-video"),
		"mux-after-done": carapace.ActionMultiParts(":", func(cColon carapace.Context) carapace.Action {
			return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					usedKeys := make([]string, len(cColon.Parts))
					for index, pair := range cColon.Parts {
						usedKeys[index] = strings.Split(pair, "=")[0]
					}

					return carapace.ActionValuesDescribed(
						"format", "set container",
						"muxer", "set muxer",
						"bin_path", "set binary file path",
						"skip_sub", "set whether or not skip subtitle files",
						"keep", "set whether or not keep files",
					).Invoke(c).Filter(usedKeys...).Suffix("=").ToA()
				case 1:
					switch c.Parts[0] {
					case "format":
						return carapace.ActionValues("mkv", "mp4", "ts")
					case "muxer":
						return carapace.ActionValues("ffmpeg", "mkvmerge")
					case "bin_path":
						return carapace.ActionFiles()
					case "keep", "skip_sub":
						return carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
					default:
						return carapace.ActionValues()
					}
				default:
					return carapace.ActionValues()
				}
			})
		}),
		"mux-import": carapace.ActionMultiParts(":", func(cColon carapace.Context) carapace.Action {
			return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					usedKeys := make([]string, len(cColon.Parts))
					for index, pair := range cColon.Parts {
						usedKeys[index] = strings.Split(pair, "=")[0]
					}

					return carapace.ActionValuesDescribed(
						"path", "set file path",
						"lang", "set media language code (not required)",
						"name", "set description (not required)",
					).Invoke(c).Filter(usedKeys...).Suffix("=").ToA()
				case 1:
					switch c.Parts[0] {
					case "path":
						return carapace.ActionFiles()
					default:
						return carapace.ActionValues()
					}
				default:
					return carapace.ActionValues()
				}
			})
		}),
		"save-dir":        carapace.ActionDirectories(),
		"select-audio":    ActionSelectStreams(),
		"select-subtitle": ActionSelectStreams(),
		"select-video":    ActionSelectStreams(),
		"sub-format":      carapace.ActionValues("SRT", "VTT"),
		"tmp-dir":         carapace.ActionDirectories(),
		"ui-language":     carapace.ActionValues("en-US", "zh-CN", "zh-TW"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}

func ActionSelectStreams() carapace.Action {
	return carapace.ActionMultiParts(":", func(cColon carapace.Context) carapace.Action {
		return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				usedKeys := make([]string, len(cColon.Parts))
				for index, pair := range cColon.Parts {
					usedKeys[index] = strings.Split(pair, "=")[0]
				}

				return carapace.ActionValues(
					"bwMax",
					"bwMin",
					"ch",
					"codecs",
					"for",
					"frame",
					"id",
					"lang",
					"name",
					"plistDurMax",
					"plistDurMin",
					"range",
					"res",
					"role",
					"segsMax",
					"segsMin",
					"url",
				).Invoke(c).Filter(usedKeys...).Suffix("=").ToA()
			case 1:
				switch c.Parts[0] {
				case "for":
					return carapace.ActionValues("best", "worst", "all")
				default:
					return carapace.ActionValues()
				}
			default:
				return carapace.ActionValues()
			}
		})
	})
}
