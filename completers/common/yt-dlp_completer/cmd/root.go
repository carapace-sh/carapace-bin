package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/http"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/ytdlp"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yt-dlp",
	Short: "A youtube-dl fork with additional features and fixes",
	Long:  "https://github.com/yt-dlp/yt-dlp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("abort-on-error", false, "Abort downloading of further videos if an error occurs")
	rootCmd.Flags().Bool("abort-on-unavailable-fragments", false, "Abort download if a fragment is unavailable")
	rootCmd.Flags().StringSlice("add-headers", nil, "Specify a custom HTTP header and its value")
	rootCmd.Flags().String("age-limit", "", "Download only videos suitable for the given age")
	rootCmd.Flags().StringSlice("alias", nil, "Create aliases for an option string")
	rootCmd.Flags().Bool("allow-dynamic-mpd", false, "Process dynamic DASH manifests")
	rootCmd.Flags().Bool("ap-list-mso", false, "List all supported multiple-system operators")
	rootCmd.Flags().String("ap-mso", "", "Adobe Pass multiple-system operator")
	rootCmd.Flags().String("ap-password", "", "Multiple-system operator account password")
	rootCmd.Flags().String("ap-username", "", "Multiple-system operator account login")
	rootCmd.Flags().String("audio-format", "", "Format to convert the audio to when -x is used")
	rootCmd.Flags().Bool("audio-multistreams", false, "Allow multiple audio streams to be merged into a single file")
	rootCmd.Flags().String("audio-quality", "", "Specify ffmpeg audio quality to use when converting the audio with -x")
	rootCmd.Flags().StringP("batch-file", "a", "", "File containing URLs to download")
	rootCmd.Flags().Bool("bidi-workaround", false, "Work around terminals that lack bidirectional text support")
	rootCmd.Flags().String("break-match-filters", "", "Same as \"--match-filters\" but stops the download process when a video is rejected")
	rootCmd.Flags().Bool("break-on-existing", false, "Stop the download process when encountering a file that is in the archive")
	rootCmd.Flags().Bool("break-per-input", false, "Reset per input URL")
	rootCmd.Flags().String("buffer-size", "", "Size of download buffer")
	rootCmd.Flags().String("cache-dir", "", "Location in the filesystem where yt-dlp can store some downloaded information")
	rootCmd.Flags().Bool("check-all-formats", false, "Check all formats for whether they are actually downloadable")
	rootCmd.Flags().Bool("check-formats", false, "Make sure formats are selected only from those that are actually downloadable")
	rootCmd.Flags().Bool("clean-info-json", false, "Remove some private fields such as filenames from the infojson")
	rootCmd.Flags().String("client-certificate", "", "Path to client certificate file in PEM format")
	rootCmd.Flags().String("client-certificate-key", "", "Path to private key file for client certificate")
	rootCmd.Flags().String("client-certificate-password", "", "Password for client certificate private key")
	rootCmd.Flags().String("compat-options", "", "Options that can help keep compatibility with youtube-dl or youtube-dlc")
	rootCmd.Flags().String("concat-playlist", "", "Concatenate videos in a playlist")
	rootCmd.Flags().StringP("concurrent-fragments", "N", "", "Number of fragments of a dash/hlsnative video that should be downloaded concurrently")
	rootCmd.Flags().StringSlice("config-locations", nil, "Location of the main configuration file")
	rootCmd.Flags().Bool("console-title", false, "Display progress in console titlebar")
	rootCmd.Flags().BoolP("continue", "c", false, "Resume partially downloaded files/fragments")
	rootCmd.Flags().String("convert-subs", "", "Convert the subtitles to another format")
	rootCmd.Flags().String("convert-thumbnails", "", "Convert the thumbnails to another format")
	rootCmd.Flags().String("cookies", "", "Netscape formatted file to read cookies from and dump cookie jar in")
	rootCmd.Flags().String("cookies-from-browser", "", "The name of the browser to load cookies from")
	rootCmd.Flags().String("date", "", "Download only videos uploaded on this date")
	rootCmd.Flags().String("dateafter", "", "Download only videos uploaded on or after this date")
	rootCmd.Flags().String("datebefore", "", "Download only videos uploaded on or before this date")
	rootCmd.Flags().String("default-search", "", "Use this prefix for unqualified URLs")
	rootCmd.Flags().String("download-archive", "", "Download only videos not listed in the archive file")
	rootCmd.Flags().String("download-sections", "", "Download only chapters whose title matches the given regular expression")
	rootCmd.Flags().String("downloader", "", "Name or path of the external downloader to use")
	rootCmd.Flags().String("downloader-args", "", "Give these arguments to the external downloader")
	rootCmd.Flags().BoolP("dump-json", "j", false, "Quiet, but print JSON information for each video")
	rootCmd.Flags().Bool("dump-pages", false, "Print downloaded pages encoded using base64 to debug problems")
	rootCmd.Flags().BoolP("dump-single-json", "J", false, "Quiet, but print JSON information for each url or infojson passed")
	rootCmd.Flags().Bool("dump-user-agent", false, "Display the current user-agent and exit")
	rootCmd.Flags().Bool("embed-chapters", false, "Add chapter markers to the video file")
	rootCmd.Flags().Bool("embed-info-json", false, "Embed the infojson as an attachment to mkv/mka video files")
	rootCmd.Flags().Bool("embed-metadata", false, "Embed metadata to the video file")
	rootCmd.Flags().Bool("embed-subs", false, "Embed subtitles in the video")
	rootCmd.Flags().Bool("embed-thumbnail", false, "Embed thumbnail in the video as cover art")
	rootCmd.Flags().Bool("enable-file-urls", false, "Enable file:// URLs")
	rootCmd.Flags().String("encoding", "", "Force the specified encoding")
	rootCmd.Flags().String("exec", "", "Execute a command")
	rootCmd.Flags().BoolP("extract-audio", "x", false, "Convert video files to audio-only files")
	rootCmd.Flags().String("extractor-args", "", "Pass ARGS arguments to the IE_KEY extractor")
	rootCmd.Flags().Bool("extractor-descriptions", false, "Output descriptions of all supported extractors and exit")
	rootCmd.Flags().String("extractor-retries", "", "Number of retries for known extractor errors")
	rootCmd.Flags().String("ffmpeg-location", "", "Location of the ffmpeg binary")
	rootCmd.Flags().String("file-access-retries", "", "Number of times to retry on file access error")
	rootCmd.Flags().String("fixup", "", "Automatically correct known faults of the file")
	rootCmd.Flags().Bool("flat-playlist", false, "Do not extract the videos of a playlist, only list them")
	rootCmd.Flags().BoolP("force-ipv4", "4", false, "Make all connections via IPv4")
	rootCmd.Flags().BoolP("force-ipv6", "6", false, "Make all connections via IPv6")
	rootCmd.Flags().Bool("force-keyframes-at-cuts", false, "Force keyframes at cuts when downloading/splitting/removing sections")
	rootCmd.Flags().Bool("force-overwrites", false, "Overwrite all video and metadata files")
	rootCmd.Flags().Bool("force-write-archive", false, "Force download archive entries to be written as far as no errors occur")
	rootCmd.Flags().StringP("format", "f", "", "Video format code")
	rootCmd.Flags().StringP("format-sort", "S", "", "Sort the formats by the fields given")
	rootCmd.Flags().Bool("format-sort-force", false, "Force user specified sort order to have precedence over all fields")
	rootCmd.Flags().String("fragment-retries", "", "Number of retries for a fragment")
	rootCmd.Flags().Bool("geo-bypass", false, "Bypass geographic restriction via faking X-Forwarded-For HTTP header")
	rootCmd.Flags().String("geo-bypass-country", "", "Force bypass geographic restriction with explicitly provided two-letter ISO 3166-2 country code")
	rootCmd.Flags().String("geo-bypass-ip-block", "", "Force bypass geographic restriction with explicitly provided IP block in CIDR notation")
	rootCmd.Flags().String("geo-verification-proxy", "", "Use this proxy to verify the IP address for some geo-restricted sites")
	rootCmd.Flags().BoolP("help", "h", false, "Print this help text and exit")
	rootCmd.Flags().Bool("hls-split-discontinuity", false, "Split HLS playlists to different formats at discontinuities such as ad breaks")
	rootCmd.Flags().Bool("hls-use-mpegts", false, "Use the mpegts container for HLS videos")
	rootCmd.Flags().String("http-chunk-size", "", "Size of a chunk for chunk-based HTTP downloading")
	rootCmd.Flags().Bool("ignore-config", false, "Don't load any more configuration files")
	rootCmd.Flags().Bool("ignore-dynamic-mpd", false, "Do not process dynamic DASH manifests")
	rootCmd.Flags().BoolP("ignore-errors", "i", false, "Ignore download and postprocessing errors")
	rootCmd.Flags().Bool("ignore-no-formats-error", false, "Ignore \"No video formats\" error")
	rootCmd.Flags().Bool("keep-fragments", false, "Keep downloaded fragments on disk after downloading is finished")
	rootCmd.Flags().BoolP("keep-video", "k", false, "Keep the intermediate video file on disk after post-processing")
	rootCmd.Flags().Bool("lazy-playlist", false, "Process entries in the playlist as they are received")
	rootCmd.Flags().Bool("legacy-server-connect", false, "Explicitly allow HTTPS connection to servers")
	rootCmd.Flags().StringP("limit-rate", "r", "", "Maximum download rate in bytes per second")
	rootCmd.Flags().Bool("list-extractors", false, "List all supported extractors and exit")
	rootCmd.Flags().BoolP("list-formats", "F", false, "List available formats of each video")
	rootCmd.Flags().Bool("list-subs", false, "List available subtitles of each video")
	rootCmd.Flags().Bool("list-thumbnails", false, "List available thumbnails of each video")
	rootCmd.Flags().Bool("live-from-start", false, "Download livestreams from the start")
	rootCmd.Flags().String("load-info-json", "", "JSON file containing the video information")
	rootCmd.Flags().Bool("mark-watched", false, "Mark videos watched")
	rootCmd.Flags().String("match-filters", "", "Generic video filter")
	rootCmd.Flags().String("max-downloads", "", "Abort after downloading NUMBER files")
	rootCmd.Flags().String("max-filesize", "", "Abort download if filesize is larger than SIZE")
	rootCmd.Flags().String("max-sleep-interval", "", "Maximum number of seconds to sleep")
	rootCmd.Flags().String("merge-output-format", "", "Containers that may be used when merging formats")
	rootCmd.Flags().String("min-filesize", "", "Abort download if filesize is smaller than SIZE")
	rootCmd.Flags().Bool("mtime", false, "Use the Last-modified header to set the file modification time")
	rootCmd.Flags().BoolP("netrc", "n", false, "Use .netrc authentication data")
	rootCmd.Flags().String("netrc-location", "", "Location of .netrc authentication data")
	rootCmd.Flags().Bool("newline", false, "Output progress bar as new lines")
	rootCmd.Flags().Bool("no-abort-on-error", false, "Continue with next video on download errors")
	rootCmd.Flags().Bool("no-audio-multistreams", false, "Only one audio stream is downloaded for each output file")
	rootCmd.Flags().Bool("no-batch-file", false, "Do not read URLs from batch file")
	rootCmd.Flags().Bool("no-break-match-filters", false, "Do not use any --break-match-filters")
	rootCmd.Flags().Bool("no-break-per-input", false, "--break-on-existing and similar options terminates the entire download queue")
	rootCmd.Flags().Bool("no-cache-dir", false, "Disable filesystem caching")
	rootCmd.Flags().Bool("no-check-certificates", false, "Suppress HTTPS certificate validation")
	rootCmd.Flags().Bool("no-check-formats", false, "Do not check that the formats are actually downloadable")
	rootCmd.Flags().Bool("no-clean-info-json", false, "Write all fields to the infojson")
	rootCmd.Flags().Bool("no-colors", false, "Do not emit color codes in output")
	rootCmd.Flags().Bool("no-config-locations", false, "Do not load any custom configuration files")
	rootCmd.Flags().Bool("no-continue", false, "Do not resume partially downloaded fragments")
	rootCmd.Flags().Bool("no-cookies", false, "Do not read/dump cookies from/to file")
	rootCmd.Flags().Bool("no-cookies-from-browser", false, "Do not load cookies from browser")
	rootCmd.Flags().Bool("no-download-archive", false, "Do not use archive file")
	rootCmd.Flags().Bool("no-embed-chapters", false, "Do not add chapter markers")
	rootCmd.Flags().Bool("no-embed-info-json", false, "Do not embed the infojson as an attachment to the video file")
	rootCmd.Flags().Bool("no-embed-metadata", false, "Do not add metadata to file")
	rootCmd.Flags().Bool("no-embed-subs", false, "Do not embed subtitles")
	rootCmd.Flags().Bool("no-embed-thumbnail", false, "Do not embed thumbnail")
	rootCmd.Flags().Bool("no-exec", false, "Remove any previously defined --exec")
	rootCmd.Flags().Bool("no-flat-playlist", false, "Extract the videos of a playlist")
	rootCmd.Flags().Bool("no-force-keyframes-at-cuts", false, "Do not force keyframes around the chapters when cutting/splitting")
	rootCmd.Flags().Bool("no-force-overwrites", false, "Do not overwrite the video, but overwrite related files")
	rootCmd.Flags().Bool("no-format-sort-force", false, "Some fields have precedence over the user specified sort order")
	rootCmd.Flags().Bool("no-geo-bypass", false, "Do not bypass geographic restriction via faking X-Forwarded-For")
	rootCmd.Flags().Bool("no-hls-split-discontinuity", false, "Do not split HLS playlists to different formats at discontinuities such as ad breaks")
	rootCmd.Flags().Bool("no-hls-use-mpegts", false, "Do not use the mpegts container for HLS videos")
	rootCmd.Flags().Bool("no-ignore-no-formats-error", false, "Throw error when no downloadable video formats are found")
	rootCmd.Flags().Bool("no-keep-fragments", false, "Delete downloaded fragments after downloading is finished")
	rootCmd.Flags().Bool("no-keep-video", false, "Delete the intermediate video file after post-processing")
	rootCmd.Flags().Bool("no-lazy-playlist", false, "Process videos in the playlist only after the entire playlist is parsed")
	rootCmd.Flags().Bool("no-live-from-start", false, "Download livestreams from the current time")
	rootCmd.Flags().Bool("no-mark-watched", false, "Do not mark videos watched")
	rootCmd.Flags().Bool("no-match-filter", false, "Do not use any --match-filter")
	rootCmd.Flags().Bool("no-mtime", false, "Do not use the Last-modified header to set the file modification time")
	rootCmd.Flags().BoolP("no-overwrites", "w", false, "Do not overwrite any files")
	rootCmd.Flags().Bool("no-part", false, "Do not use .part files - write directly into output file")
	rootCmd.Flags().Bool("no-playlist", false, "Download only the video")
	rootCmd.Flags().Bool("no-post-overwrites", false, "Do not overwrite post-processed files")
	rootCmd.Flags().Bool("no-prefer-free-formats", false, "Don't give any special preference to free containers")
	rootCmd.Flags().Bool("no-progress", false, "Do not print progress bar")
	rootCmd.Flags().Bool("no-remove-chapters", false, "Do not remove any chapters from the file")
	rootCmd.Flags().Bool("no-resize-buffer", false, "Do not automatically adjust the buffer size")
	rootCmd.Flags().Bool("no-restrict-filenames", false, "Allow Unicode characters")
	rootCmd.Flags().Bool("no-simulate", false, "Download the video even if printing/listing options are used")
	rootCmd.Flags().Bool("no-split-chapters", false, "Do not split video based on chapters")
	rootCmd.Flags().Bool("no-sponsorblock", false, "Disable both --sponsorblock-mark and --sponsorblock-remove")
	rootCmd.Flags().Bool("no-update", false, "Do not check for updates")
	rootCmd.Flags().Bool("no-video-multistreams", false, "Only one video stream is downloaded for each output file")
	rootCmd.Flags().Bool("no-wait-for-video", false, "Do not wait for scheduled streams")
	rootCmd.Flags().Bool("no-warnings", false, "Ignore warnings")
	rootCmd.Flags().Bool("no-windows-filenames", false, "Make filenames Windows-compatible only if using Windows")
	rootCmd.Flags().Bool("no-write-auto-subs", false, "Do not write auto-generated subtitles")
	rootCmd.Flags().Bool("no-write-comments", false, "Do not retrieve video comments unless the extraction is known to be quick")
	rootCmd.Flags().Bool("no-write-description", false, "Do not write video description")
	rootCmd.Flags().Bool("no-write-info-json", false, "Do not write video metadata")
	rootCmd.Flags().Bool("no-write-playlist-metafiles", false, "Do not write playlist metadata")
	rootCmd.Flags().Bool("no-write-subs", false, "Do not write subtitle file")
	rootCmd.Flags().Bool("no-write-thumbnail", false, "Do not write thumbnail image to disk")
	rootCmd.Flags().StringP("output", "o", "", "Output filename template")
	rootCmd.Flags().String("output-na-placeholder", "", "Placeholder for unavailable fields in \"OUTPUT TEMPLATE\"")
	rootCmd.Flags().String("parse-metadata", "", "Parse additional metadata like title/artist from other fields")
	rootCmd.Flags().Bool("part", false, "Use .part files instead of writing directly into output file")
	rootCmd.Flags().StringP("password", "p", "", "Account password")
	rootCmd.Flags().StringP("paths", "P", "", "The paths where the files should be downloaded")
	rootCmd.Flags().StringP("playlist-items", "I", "", "Comma separated playlist_index of the items to download")
	rootCmd.Flags().Bool("playlist-random", false, "Download playlist videos in random order")
	rootCmd.Flags().Bool("post-overwrites", false, "Overwrite post-processed files")
	rootCmd.Flags().String("postprocessor-args", "", "Give these arguments to the postprocessors")
	rootCmd.Flags().Bool("prefer-free-formats", false, "Prefer video formats with free containers over non-free ones of same quality")
	rootCmd.Flags().Bool("prefer-insecure", false, "Use an unencrypted connection to retrieve information about the video")
	rootCmd.Flags().StringP("print", "O", "", "Field name or output template to print to screen")
	rootCmd.Flags().StringSlice("print-to-file", nil, "Append given template to the file")
	rootCmd.Flags().Bool("print-traffic", false, "Display sent and read HTTP traffic")
	rootCmd.Flags().Bool("progress", false, "Show progress bar, even if in quiet mode")
	rootCmd.Flags().String("progress-template", "", "Template for progress outputs")
	rootCmd.Flags().String("proxy", "", "Use the specified HTTP/HTTPS/SOCKS proxy")
	rootCmd.Flags().BoolP("quiet", "q", false, "Activate quiet mode")
	rootCmd.Flags().String("recode-video", "", "Re-encode the video into another format if necessary")
	rootCmd.Flags().String("remove-chapters", "", "Remove chapters whose title matches the given regular expression")
	rootCmd.Flags().String("remux-video", "", "Remux the video into another container if necessary")
	rootCmd.Flags().StringSlice("replace-in-metadata", nil, "Replace text in a metadata field using the given regex")
	rootCmd.Flags().Bool("resize-buffer", false, "The buffer size is automatically resized from an initial value")
	rootCmd.Flags().Bool("restrict-filenames", false, "Restrict filenames to only ASCII characters")
	rootCmd.Flags().StringP("retries", "R", "", "Number of retries")
	rootCmd.Flags().String("retry-sleep", "", "Time to sleep between retries in seconds")
	rootCmd.Flags().Bool("rm-cache-dir", false, "Delete all filesystem cache files")
	rootCmd.Flags().BoolP("simulate", "s", false, "Do not download the video and do not write anything to disk")
	rootCmd.Flags().Bool("skip-download", false, "Do not download the video but write all related files")
	rootCmd.Flags().String("skip-playlist-after-errors", "", "Number of allowed failures until the rest of the playlist is skipped")
	rootCmd.Flags().Bool("skip-unavailable-fragments", false, "Skip unavailable fragments for DASH, hlsnative and ISM downloads")
	rootCmd.Flags().String("sleep-interval", "", "Number of seconds to sleep before each download")
	rootCmd.Flags().String("sleep-requests", "", "Number of seconds to sleep between requests during data extraction")
	rootCmd.Flags().String("sleep-subtitles", "", "Number of seconds to sleep before each subtitle download")
	rootCmd.Flags().String("socket-timeout", "", "Time to wait before giving up, in seconds")
	rootCmd.Flags().String("source-address", "", "Client-side IP address to bind to")
	rootCmd.Flags().Bool("split-chapters", false, "Split video into multiple files based on internal chapters")
	rootCmd.Flags().String("sponsorblock-api", "", "SponsorBlock API location")
	rootCmd.Flags().String("sponsorblock-chapter-title", "", "An output template for the title of the SponsorBlock chapters")
	rootCmd.Flags().String("sponsorblock-mark", "", "SponsorBlock categories to create chapters for")
	rootCmd.Flags().String("sponsorblock-remove", "", "SponsorBlock categories to be removed from the video file")
	rootCmd.Flags().String("sub-format", "", "Subtitle format")
	rootCmd.Flags().String("sub-langs", "", "Languages of the subtitles to download")
	rootCmd.Flags().String("throttled-rate", "", "Minimum download rate in bytes per second")
	rootCmd.Flags().String("trim-filenames", "", "Limit the filename length")
	rootCmd.Flags().StringP("twofactor", "2", "", "Two-factor authentication code")
	rootCmd.Flags().BoolP("update", "U", false, "Check if updates are available")
	rootCmd.Flags().String("update-to", "", "Upgrade/downgrade to a specific version")
	rootCmd.Flags().String("use-extractors", "", "Extractor names to use separated by commas")
	rootCmd.Flags().String("use-postprocessor", "", "The (case sensitive) name of plugin postprocessors to be enabled")
	rootCmd.Flags().StringP("username", "u", "", "Login with this account ID")
	rootCmd.Flags().BoolP("verbose", "v", false, "Print various debugging information")
	rootCmd.Flags().Bool("version", false, "Print program version and exit")
	rootCmd.Flags().Bool("video-multistreams", false, "Allow multiple video streams to be merged into a single file")
	rootCmd.Flags().String("video-password", "", "Video password")
	rootCmd.Flags().String("wait-for-video", "", "Wait for scheduled streams to become available")
	rootCmd.Flags().Bool("windows-filenames", false, "Force filenames to be Windows-compatible")
	rootCmd.Flags().Bool("write-all-thumbnails", false, "Write all thumbnail image formats to disk")
	rootCmd.Flags().Bool("write-auto-subs", false, "Write automatically generated subtitle file")
	rootCmd.Flags().Bool("write-comments", false, "Retrieve video comments to be placed in the infojson")
	rootCmd.Flags().Bool("write-description", false, "Write video description to a .description file")
	rootCmd.Flags().Bool("write-desktop-link", false, "Write a .desktop Linux internet shortcut")
	rootCmd.Flags().Bool("write-info-json", false, "Write video metadata to a .info.json file")
	rootCmd.Flags().Bool("write-link", false, "Write an internet shortcut file")
	rootCmd.Flags().Bool("write-pages", false, "Write downloaded intermediary pages to files in the current directory")
	rootCmd.Flags().Bool("write-playlist-metafiles", false, "Write playlist metadata in addition to the video metadata")
	rootCmd.Flags().Bool("write-subs", false, "Write subtitle file")
	rootCmd.Flags().Bool("write-thumbnail", false, "Write thumbnail image to disk")
	rootCmd.Flags().Bool("write-url-link", false, "Write a .url Windows internet shortcut")
	rootCmd.Flags().Bool("write-webloc-link", false, "Write a .webloc macOS internet shortcut")
	rootCmd.Flags().Bool("xattr-set-filesize", false, "Set file xattribute ytdl.filesize with expected file size")
	rootCmd.Flags().Bool("xattrs", false, "Write metadata to the video file's xattrs")
	rootCmd.Flags().Bool("yes-playlist", false, "Download the playlist")

	rootCmd.Flag("alias").Nargs = 2
	rootCmd.Flag("print-to-file").Nargs = 2
	rootCmd.Flag("replace-in-metadata").Nargs = 3

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"add-headers": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !strings.Contains(c.Value, ":") {
				return http.ActionRequestHeaderNames().Suffix(":").NoSpace(':')
			}
			header := strings.SplitN(c.Value, ":", 2)[0]
			c.Value = strings.TrimPrefix(c.Value, header+":")
			return http.ActionRequestHeaderValues(header).Invoke(c).Prefix(header + ":").ToA()
		}),
		"batch-file":             carapace.ActionFiles(),
		"cache-dir":              carapace.ActionDirectories(),
		"client-certificate":     carapace.ActionFiles(),
		"client-certificate-key": carapace.ActionFiles(),
		"config-locations":       carapace.ActionFiles(),
		"convert-subs":           ytdlp.ActionSubtitleFormats(),
		"convert-thumbnails":     ytdlp.ActionThumbnailFormats(),
		"cookies-from-browser":   ytdlp.ActionBrowsers(),
		"date":                   time.ActionDate(),
		"dateafter":              time.ActionDate(),
		"datebefore":             time.ActionDate(),
		"download-archive":       carapace.ActionFiles(),
		"downloader": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !strings.Contains(c.Value, ":") {
				return carapace.Batch(
					ytdlp.ActionProtocols().Suffix(":").NoSpace(':'),
					carapace.ActionFiles(),
				).ToA()
			}

			protocol := strings.SplitN(c.Value, ":", 2)[0]
			c.Value = strings.SplitN(c.Value, ":", 2)[1]
			return carapace.ActionFiles().Invoke(c).Prefix(protocol + ":").ToA()
		}),
		"ffmpeg-location": carapace.ActionFiles(),
		"format": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			batch := carapace.Batch()
			for _, arg := range c.Args {
				batch = append(batch, ytdlp.ActionFormats(arg))
			}
			return batch.ToA()
		}),
		"load-info-json":      carapace.ActionFiles(),
		"merge-output-format": ytdlp.ActionOutputFormats().UniqueList("/"),
		"output": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			index := strings.LastIndex(c.Value, "%(")
			if index < 0 {
				return carapace.ActionFiles()
			}

			prefix := c.Value[:index+2]
			c.Value = c.Value[index+2:]
			batch := carapace.Batch(
				ytdlp.ActionFields(),
				ytdlp.ActionNumericMetaFields(),
				ytdlp.ActionStringMetaFields(),
				ytdlp.ActionChapterFields(),
				ytdlp.ActionEpisodeFields(),
				ytdlp.ActionTrackFields(),
			)

			if rootCmd.Flag("download-sections").Changed { // TODO also enabled when chapter: prefix?
				batch = append(batch, ytdlp.ActionSectionFields())
			}
			return batch.Invoke(c).Merge().Prefix(prefix).ToA().NoSpace()
		}),
		"paths": carapace.ActionFiles(),
		"print-to-file": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 1:
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()
			}
		}),
		"recode-video": ytdlp.ActionVideoFormats(),
		"remux-video":  ytdlp.ActionVideoFormats(),
		"sponsorblock-mark": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			a := ytdlp.ActionSponsorblockCategories()
			for index, part := range c.Parts {
				c.Parts[index] = strings.TrimPrefix(part, "-")
			}
			a = a.Invoke(c).Filter(c.Parts...).ToA()

			if strings.HasPrefix(c.Value, "-") {
				a = a.Prefix("-")
			}
			return a
		}).NoSpace(),
		"sponsorblock-remove": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			a := ytdlp.ActionSponsorblockCategories()
			for index, part := range c.Parts {
				c.Parts[index] = strings.TrimPrefix(part, "-")
			}
			a = a.Invoke(c).Filter(c.Parts...).ToA()

			if strings.HasPrefix(c.Value, "-") {
				a = a.Prefix("-")
			}
			return a
		}).NoSpace(),
		"sub-langs": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			batch := carapace.Batch()
			for _, arg := range c.Args {
				batch = append(batch, ytdlp.ActionSubtitles(arg))
			}
			return batch.ToA()
		}).UniqueList(","),
		"update-to": carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("stable", "nightly")
			default:
				return carapace.ActionValues()
			}
		}),
		"use-extractors": ytdlp.ActionExtractors().UniqueList(","),
	})
}
