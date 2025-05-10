package cmd

import (
	"os"

	"github.com/carapace-sh/carapace"
	osAction "github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mpv"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "mpv",
	Short: "a media player",
	Long:  "https://mpv.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("ab-loop-a", "", "Set loop point")
	rootCmd.Flags().String("ab-loop-b", "", "Set loop Point")
	rootCmd.Flags().String("ab-loop-count", "", "Run A-B only N times")
	rootCmd.Flags().String("access-references", "", "Follow any references in the file being opened")
	rootCmd.Flags().String("ad", "", "Specify a priority list of audio decoders to be used")
	rootCmd.Flags().String("ad-lavc-ac3drc", "", "Select the Dynamic Range Compression level for AC-3 audio streams")
	rootCmd.Flags().String("ad-lavc-downmix", "", "Whether  to request audio channel downmixing from the decoder")
	rootCmd.Flags().String("ad-lavc-o", "", "Pass AVOptions to libavcodec decoder")
	rootCmd.Flags().String("ad-lavc-o-add", "", "")
	rootCmd.Flags().String("ad-lavc-o-append", "", "")
	rootCmd.Flags().String("ad-lavc-o-clr", "", "")
	rootCmd.Flags().String("ad-lavc-o-del", "", "")
	rootCmd.Flags().String("ad-lavc-o-remove", "", "")
	rootCmd.Flags().String("ad-lavc-o-set", "", "")
	rootCmd.Flags().String("ad-lavc-threads", "", "Number  of threads to use for decoding")
	rootCmd.Flags().String("ad-queue-enable", "", "Enable running the audio decoder on a separate thread ")
	rootCmd.Flags().String("ad-queue-max-bytes", "", "Maximum approximate allowed size of the queue")
	rootCmd.Flags().String("ad-queue-max-samples", "", "Maximum number samples of the queue")
	rootCmd.Flags().String("ad-queue-max-secs", "", "Maximum number of seconds of media in the queue")
	rootCmd.Flags().String("af", "", "Specify a list of audio filters to  apply  to  the  audio stream")
	rootCmd.Flags().String("af-add", "", "")
	rootCmd.Flags().String("af-append", "", "")
	rootCmd.Flags().String("af-clr", "", "")
	rootCmd.Flags().String("af-help", "", "")
	rootCmd.Flags().String("af-pre", "", "")
	rootCmd.Flags().String("af-remove", "", "")
	rootCmd.Flags().String("af-set", "", "")
	rootCmd.Flags().String("af-toggle", "", "")
	rootCmd.Flags().String("aid", "", "Select audio track")
	rootCmd.Flags().String("alang", "", "Specify a prioritized list of audio languages to use, as IETF language tags")
	rootCmd.Flags().String("alang-add", "", "")
	rootCmd.Flags().String("alang-append", "", "")
	rootCmd.Flags().String("alang-clr", "", "")
	rootCmd.Flags().String("alang-del", "", "")
	rootCmd.Flags().String("alang-pre", "", "")
	rootCmd.Flags().String("alang-remove", "", "")
	rootCmd.Flags().String("alang-set", "", "")
	rootCmd.Flags().String("alang-toggle", "", "")
	rootCmd.Flags().String("allow-delayed-peak-detect", "", "When using --hdr-compute-peak, allow delaying the detected peak by a frame when beneficial for performance")
	rootCmd.Flags().String("alsa-buffer-time", "", "Set the requested buffer time in microseconds")
	rootCmd.Flags().String("alsa-ignore-chmap", "", "Don't read or set the channel map of the ALSA device")
	rootCmd.Flags().String("alsa-mixer-device", "", "Set the mixer device used with ao-volume")
	rootCmd.Flags().String("alsa-mixer-index", "", "Set the index of the mixer channel")
	rootCmd.Flags().String("alsa-mixer-name", "", "Set the name of the mixer element")
	rootCmd.Flags().String("alsa-non-interleaved", "", "Allow output of non-interleaved formats")
	rootCmd.Flags().String("alsa-periods", "", "Number of periods requested from the ALSA API")
	rootCmd.Flags().String("alsa-resample", "", "Enable ALSA resampling plugin")
	rootCmd.Flags().String("ao", "", "Specify a priority list of audio output drivers to be used")
	rootCmd.Flags().String("ao-add", "", "")
	rootCmd.Flags().String("ao-append", "", "")
	rootCmd.Flags().String("ao-clr", "", "")
	rootCmd.Flags().String("ao-help", "", "")
	rootCmd.Flags().String("ao-null-broken-delay", "", "Simulate broken audio drivers, which don't report latency correctly")
	rootCmd.Flags().String("ao-null-broken-eof", "", "Simulate broken audio drivers, which always add the fixed device latency to the reported audio playback position")
	rootCmd.Flags().String("ao-null-buffer", "", "Simulated buffer length in seconds")
	rootCmd.Flags().String("ao-null-channel-layouts", "", "If not empty, this is a , separated list of channel layouts the AO allows")
	rootCmd.Flags().String("ao-null-format", "", "Force the audio output format the AO will accept")
	rootCmd.Flags().String("ao-null-latency", "", "Simulated device latency")
	rootCmd.Flags().String("ao-null-outburst", "", "Simulated chunk size in samples")
	rootCmd.Flags().String("ao-null-speed", "", "Simulated audio playback speed as a multiplier")
	rootCmd.Flags().String("ao-null-untimed", "", "Do not simulate timing of a perfect audio device")
	rootCmd.Flags().String("ao-pcm-append", "", "Append to the file, instead of overwriting it")
	rootCmd.Flags().String("ao-pcm-file", "", "Write the sound to <filename> instead of the default audiodump.wav")
	rootCmd.Flags().String("ao-pcm-waveheader", "", "Include or do not include the WAVE header")
	rootCmd.Flags().String("ao-pre", "", "")
	rootCmd.Flags().String("ao-remove", "", "")
	rootCmd.Flags().String("ao-set", "", "")
	rootCmd.Flags().String("ao-toggle", "", "")
	rootCmd.Flags().String("archive-exts", "", "Archive file extensions to try to match when using --autocreate-playlist or --directory-filter-types")
	rootCmd.Flags().String("archive-exts-add", "", "")
	rootCmd.Flags().String("archive-exts-append", "", "")
	rootCmd.Flags().String("archive-exts-clr", "", "")
	rootCmd.Flags().String("archive-exts-del", "", "")
	rootCmd.Flags().String("archive-exts-pre", "", "")
	rootCmd.Flags().String("archive-exts-remove", "", "")
	rootCmd.Flags().String("archive-exts-set", "", "")
	rootCmd.Flags().String("archive-exts-toggle", "", "")
	rootCmd.Flags().String("audio", "", "alias for --aid")
	rootCmd.Flags().String("audio-backward-batch", "", "Number of keyframe ranges to decode at once when backward decoding")
	rootCmd.Flags().String("audio-backward-overlap", "", "Number of overlapping keyframe ranges to use for backward decoding")
	rootCmd.Flags().String("audio-buffer", "", "Set the audio output minimum buffer")
	rootCmd.Flags().String("audio-channels", "", "Control which audio channels are output")
	rootCmd.Flags().String("audio-client-name", "", "The application name the player reports to the audio API")
	rootCmd.Flags().String("audio-delay", "", "Audio delay in seconds")
	rootCmd.Flags().String("audio-demuxer", "", "Use this audio demuxer type when using --audio-file")
	rootCmd.Flags().String("audio-device", "", "Use the given audio device")
	rootCmd.Flags().String("audio-display", "", "Determines whether to display cover art when playing audio files and with what priority")
	rootCmd.Flags().String("audio-exclusive", "", "Enable exclusive output mode")
	rootCmd.Flags().String("audio-exts", "", "Audio file extensions to try to match when using --audio-file-auto, --autocreate-playlist or --directory-filter-types")
	rootCmd.Flags().String("audio-exts-add", "", "")
	rootCmd.Flags().String("audio-exts-append", "", "")
	rootCmd.Flags().String("audio-exts-clr", "", "")
	rootCmd.Flags().String("audio-exts-del", "", "")
	rootCmd.Flags().String("audio-exts-pre", "", "")
	rootCmd.Flags().String("audio-exts-remove", "", "")
	rootCmd.Flags().String("audio-exts-set", "", "")
	rootCmd.Flags().String("audio-exts-toggle", "", "")
	rootCmd.Flags().String("audio-fallback-to-null", "", "If no audio device can be opened, behave as if --ao=null was given")
	rootCmd.Flags().String("audio-file", "", "CLI/config file only alias for --audio-files-append")
	rootCmd.Flags().String("audio-file-auto", "", "Load additional audio files matching the video filename")
	rootCmd.Flags().String("audio-file-auto-exts", "", "alias for audio-exts")
	rootCmd.Flags().String("audio-file-paths", "", "Analogous to --sub-file-paths option, but for auto-loaded audio files")
	rootCmd.Flags().String("audio-file-paths-add", "", "")
	rootCmd.Flags().String("audio-file-paths-append", "", "")
	rootCmd.Flags().String("audio-file-paths-clr", "", "")
	rootCmd.Flags().String("audio-file-paths-del", "", "")
	rootCmd.Flags().String("audio-file-paths-pre", "", "")
	rootCmd.Flags().String("audio-file-paths-remove", "", "")
	rootCmd.Flags().String("audio-file-paths-set", "", "")
	rootCmd.Flags().String("audio-file-paths-toggle", "", "")
	rootCmd.Flags().String("audio-files", "", "Play audio from an external file while viewing a video")
	rootCmd.Flags().String("audio-files-add", "", "")
	rootCmd.Flags().String("audio-files-append", "", "")
	rootCmd.Flags().String("audio-files-clr", "", "")
	rootCmd.Flags().String("audio-files-del", "", "")
	rootCmd.Flags().String("audio-files-pre", "", "")
	rootCmd.Flags().String("audio-files-remove", "", "")
	rootCmd.Flags().String("audio-files-set", "", "")
	rootCmd.Flags().String("audio-files-toggle", "", "")
	rootCmd.Flags().String("audio-format", "", "Select the sample format used for output from the audio filter layer to the sound card")
	rootCmd.Flags().String("audio-normalize-downmix", "", "Enable/disable normalization if surround audio is downmixed to stereo")
	rootCmd.Flags().String("audio-pitch-correction", "", "If this is enabled (default), playing with a speed different from normal automatically inserts the scaletempo2 audio filter")
	rootCmd.Flags().String("audio-resample-cutoff", "", "Cutoff frequency")
	rootCmd.Flags().String("audio-resample-filter-size", "", "Length of the filter with respect to the lower sampling rate")
	rootCmd.Flags().String("audio-resample-linear", "", "If set then filters will be linearly interpolated between polyphase entries")
	rootCmd.Flags().String("audio-resample-max-output-size", "", "Limit maximum size of audio frames filtered at once, in ms")
	rootCmd.Flags().String("audio-resample-phase-shift", "", "Log2 of the number of polyphase entries")
	rootCmd.Flags().String("audio-reversal-buffer", "", "For backward decoding")
	rootCmd.Flags().String("audio-samplerate", "", "Select the output sample rate to be used ")
	rootCmd.Flags().String("audio-spdif", "", "List of codecs for which compressed audio passthrough should be used")
	rootCmd.Flags().String("audio-stream-silence", "", "Cash-grab consumer audio hardware (such as A/V receivers) often ignore initial audio sent over HDMI")
	rootCmd.Flags().String("audio-swresample-o", "", "Set AVOptions on the SwrContext or AVAudioResampleContext")
	rootCmd.Flags().String("audio-swresample-o-add", "", "")
	rootCmd.Flags().String("audio-swresample-o-append", "", "")
	rootCmd.Flags().String("audio-swresample-o-clr", "", "")
	rootCmd.Flags().String("audio-swresample-o-del", "", "")
	rootCmd.Flags().String("audio-swresample-o-remove", "", "")
	rootCmd.Flags().String("audio-swresample-o-set", "", "")
	rootCmd.Flags().String("audio-wait-open", "", "This makes sense for use with --audio-stream-silence=yes. If this option is given, the player will wait for the given amount of seconds after opening the audio device before sending actual audio data to it")
	rootCmd.Flags().String("auto-window-resize", "", "By default, mpv will automatically resize itself if the video's size changes")
	rootCmd.Flags().String("autocreate-playlist", "", "When opening a local file, act as if the parent directory is opened and create a playlist automatically")
	rootCmd.Flags().String("autofit", "", "Set the initial window size to a maximum size specified by WxH, without changing the window's aspect ratio")
	rootCmd.Flags().String("autofit-larger", "", "This option behaves exactly like --autofit, except that it sets the maximum size of the window")
	rootCmd.Flags().String("autofit-smaller", "", "This option behaves exactly like --autofit, except that it sets the minimum size of the window")
	rootCmd.Flags().String("autoload-files", "", "Automatically load/select external files")
	rootCmd.Flags().String("autosync", "", "Gradually adjusts the A/V sync based on audio delay measurements")
	rootCmd.Flags().String("background", "", "If the frame has an alpha component, decide what kind of background, if any, to blend it with")
	rootCmd.Flags().String("background-color", "", "Color used to draw parts of the mpv window not covered by video")
	rootCmd.Flags().String("blend-subtitles", "", "Blend subtitles directly onto upscaled video frames, before interpolation and/or color management")
	rootCmd.Flags().String("bluray-device", "", "Specify the Blu-ray disc location")
	rootCmd.Flags().String("border", "", "Play video with window border and decorations")
	rootCmd.Flags().String("border-background", "", "Same as --background but only applies to the black bar/border area of the window")
	rootCmd.Flags().String("brightness", "", "Adjust the brightness of the video signal")
	rootCmd.Flags().String("cache", "", "Decide whether to use network cache settings")
	rootCmd.Flags().String("cache-on-disk", "", "Write packet data to a temporary file, instead of keeping them in memory")
	rootCmd.Flags().String("cache-pause", "", "Whether the player should automatically pause when the cache runs out of data and stalls decoding/playback")
	rootCmd.Flags().String("cache-pause-initial", "", "Enter \"buffering\" mode before starting playback")
	rootCmd.Flags().String("cache-pause-wait", "", "Number of seconds the packet cache should have buffered before starting playback again if \"buffering\" was entered")
	rootCmd.Flags().String("cache-secs", "", "How many seconds of audio/video to prefetch if the cache is active")
	rootCmd.Flags().String("cdda-cdtext", "", "Print CD text")
	rootCmd.Flags().String("cdda-device", "", "Specify the CD device for CDDA playback")
	rootCmd.Flags().String("cdda-overlap", "", "Force minimum overlap search during verification to <value> sectors")
	rootCmd.Flags().String("cdda-paranoia", "", "Set paranoia level")
	rootCmd.Flags().String("cdda-sector-size", "", "Set atomic read size")
	rootCmd.Flags().String("cdda-skip", "", "(Never) accept imperfect data reconstruction")
	rootCmd.Flags().String("cdda-speed", "", "Set CD spin speed")
	rootCmd.Flags().String("cdda-toc-offset", "", "Add <value> sectors to the values reported when addressing tracks")
	rootCmd.Flags().String("chapter-merge-threshold", "", "Threshold for merging almost consecutive ordered chapter parts in milliseconds")
	rootCmd.Flags().String("chapter-seek-threshold", "", "Distance in seconds from the beginning of a chapter within which a backward chapter seek will go to the previous chapter")
	rootCmd.Flags().String("chapters-file", "", "Load chapters from this file, instead of using the chapter metadata found in the main file")
	rootCmd.Flags().String("clipboard-backends", "", "Specify a priority list of the clipboard backends to be used")
	rootCmd.Flags().String("clipboard-backends-add", "", "")
	rootCmd.Flags().String("clipboard-backends-append", "", "")
	rootCmd.Flags().String("clipboard-backends-clr", "", "")
	rootCmd.Flags().String("clipboard-backends-help", "", "")
	rootCmd.Flags().String("clipboard-backends-pre", "", "")
	rootCmd.Flags().String("clipboard-backends-remove", "", "")
	rootCmd.Flags().String("clipboard-backends-set", "", "")
	rootCmd.Flags().String("clipboard-backends-toggle", "", "")
	rootCmd.Flags().String("clipboard-monitor", "", "Enable clipboard monitoring so that the clipboard property can be observed for content changes")
	rootCmd.Flags().String("config", "", "")
	rootCmd.Flags().String("config-dir", "", "Force a different configuration directory")
	rootCmd.Flags().String("container-fps-override", "", "Override video framerate")
	rootCmd.Flags().String("contrast", "", "Adjust the contrast of the video signal")
	rootCmd.Flags().String("cookies", "", "Support cookies when making HTTP requests")
	rootCmd.Flags().String("cookies-file", "", "Read HTTP cookies from <filename>")
	rootCmd.Flags().String("corner-rounding", "", "If set to a value above 0.0, the output will be rendered with rounded corners, as if an alpha transparency mask had been applied")
	rootCmd.Flags().String("correct-downscaling", "", "When using convolution based filters, extend the filter size when downscaling")
	rootCmd.Flags().String("correct-pts", "", "--correct-pts=no switches mpv to a mode where video timing is determined using a fixed framerate value")
	rootCmd.Flags().String("cover-art-auto", "", "Whether to load _external_ cover art automatically")
	rootCmd.Flags().String("cover-art-auto-exts", "", "alias for image-exts")
	rootCmd.Flags().StringSlice("cover-art-file", nil, "CLI/config file only alias for --cover-art-files-append")
	rootCmd.Flags().String("cover-art-files", "", "Use an external file as cover art while playing audio")
	rootCmd.Flags().String("cover-art-files-add", "", "")
	rootCmd.Flags().String("cover-art-files-append", "", "")
	rootCmd.Flags().String("cover-art-files-clr", "", "")
	rootCmd.Flags().String("cover-art-files-del", "", "")
	rootCmd.Flags().String("cover-art-files-pre", "", "")
	rootCmd.Flags().String("cover-art-files-remove", "", "")
	rootCmd.Flags().String("cover-art-files-set", "", "")
	rootCmd.Flags().String("cover-art-files-toggle", "", "")
	rootCmd.Flags().String("cover-art-whitelist", "", "Filenames to load as cover art, sorted by descending priority")
	rootCmd.Flags().String("cover-art-whitelist-add", "", "")
	rootCmd.Flags().String("cover-art-whitelist-append", "", "")
	rootCmd.Flags().String("cover-art-whitelist-clr", "", "")
	rootCmd.Flags().String("cover-art-whitelist-del", "", "")
	rootCmd.Flags().String("cover-art-whitelist-pre", "", "")
	rootCmd.Flags().String("cover-art-whitelist-remove", "", "")
	rootCmd.Flags().String("cover-art-whitelist-set", "", "")
	rootCmd.Flags().String("cover-art-whitelist-toggle", "", "")
	rootCmd.Flags().String("cscale", "", "As --scale, but for interpolating chroma information")
	rootCmd.Flags().String("cscale-antiring", "", "Set the antiringing strength")
	rootCmd.Flags().String("cscale-blur", "", "Kernel scaling factor")
	rootCmd.Flags().String("cscale-clamp", "", "Specifies a weight bias to multiply into negative coefficients")
	rootCmd.Flags().String("cscale-param1", "", "Set filter parameters")
	rootCmd.Flags().String("cscale-param2", "", "Set filter parameters")
	rootCmd.Flags().String("cscale-radius", "", "Set radius for tunable filters, must be a float number between 0.5 and 16.0")
	rootCmd.Flags().String("cscale-taper", "", "Kernel/window taper factor")
	rootCmd.Flags().String("cscale-window", "", "Choose a custom windowing function for the kernel")
	rootCmd.Flags().String("cscale-wparam", "", "Configure the parameter for the window function given by --scale-window etc")
	rootCmd.Flags().String("cscale-wtaper", "", "Kernel/window taper factor")
	rootCmd.Flags().String("cuda-decode-device", "", "Choose the GPU device used for decoding when using the cuda or nvdec hwdecs with the OpenGL GPU backend, and with the cuda-copy or nvdec-copy hwdecs in all cases")
	rootCmd.Flags().String("cursor-autohide", "", "Make mouse cursor automatically hide after given number of milliseconds")
	rootCmd.Flags().String("cursor-autohide-fs-only", "", "If this option is given, the cursor is always visible in windowed mode")
	rootCmd.Flags().String("deband", "", "Enable the debanding algorithm")
	rootCmd.Flags().String("deband-grain", "", "Add some extra noise to the image")
	rootCmd.Flags().String("deband-iterations", "", "The number of debanding steps to perform per sample")
	rootCmd.Flags().String("deband-range", "", "The debanding filter's initial radius")
	rootCmd.Flags().String("deband-threshold", "", "The debanding filter's cut-off threshold")
	rootCmd.Flags().String("deinterlace", "", "Enable or disable deinterlacing")
	rootCmd.Flags().String("deinterlace-field-parity", "", "Specify the field parity/order when deinterlacing")
	rootCmd.Flags().String("demuxer", "", "Force demuxer type")
	rootCmd.Flags().String("demuxer-backward-playback-step", "", "Number of seconds the demuxer should seek back to get new packets during backward playback")
	rootCmd.Flags().String("demuxer-cache-dir", "", "Directory where to create temporary files")
	rootCmd.Flags().String("demuxer-cache-unlink-files", "", "Whether or when to unlink cache files")
	rootCmd.Flags().String("demuxer-cache-wait", "", "Before starting playback, read data until either the end of the file was reached, or the demuxer cache has reached maximum capacity")
	rootCmd.Flags().String("demuxer-donate-buffer", "", "Whether to let the back buffer use part of the forward buffer")
	rootCmd.Flags().String("demuxer-hysteresis-secs", "", "Once the demuxer limit is reached (--demuxer-max-bytes, --demuxer-readahead-secs or --cache-secs), this value can be used to specify a hysteresis before the demuxer will buffer ahead again")
	rootCmd.Flags().String("demuxer-lavf-allow-mimetype", "", "Allow deriving the format from the HTTP MIME type")
	rootCmd.Flags().String("demuxer-lavf-analyzeduration", "", "Maximum length in seconds to analyze the stream properties")
	rootCmd.Flags().String("demuxer-lavf-buffersize", "", "Size  of the stream read buffer allocated for libavformat in bytes")
	rootCmd.Flags().String("demuxer-lavf-format", "", "Force a specific libavformat demuxer")
	rootCmd.Flags().String("demuxer-lavf-hacks", "", "By default, some formats will be handled differently from other formats by explicitly checking for them")
	rootCmd.Flags().String("demuxer-lavf-linearize-timestamps", "", "Attempt to linearize timestamp resets in demuxed streams")
	rootCmd.Flags().String("demuxer-lavf-o", "", "Pass AVOptions to libavformat demuxer")
	rootCmd.Flags().String("demuxer-lavf-o-add", "", "")
	rootCmd.Flags().String("demuxer-lavf-o-append", "", "")
	rootCmd.Flags().String("demuxer-lavf-o-clr", "", "")
	rootCmd.Flags().String("demuxer-lavf-o-del", "", "")
	rootCmd.Flags().String("demuxer-lavf-o-remove", "", "")
	rootCmd.Flags().String("demuxer-lavf-o-set", "", "")
	rootCmd.Flags().String("demuxer-lavf-probe-info", "", "Whether to probe stream information")
	rootCmd.Flags().String("demuxer-lavf-probescore", "", "Minimum required libavformat probe score")
	rootCmd.Flags().String("demuxer-lavf-probesize", "", "Maximum amount of data to probe during the detection phase")
	rootCmd.Flags().String("demuxer-lavf-propagate-opts", "", "Propagate FFmpeg-level options to recursively opened connections")
	rootCmd.Flags().String("demuxer-max-back-bytes", "", "This controls how much past data the demuxer is allowed to preserve")
	rootCmd.Flags().String("demuxer-max-bytes", "", "This controls how much the demuxer is allowed to buffer ahead")
	rootCmd.Flags().String("demuxer-mkv-crop-compat", "", "Enable compatibility mode for files that do not fully comply with the Matroska specification")
	rootCmd.Flags().String("demuxer-mkv-probe-start-time", "", "Check the start time of Matroska files")
	rootCmd.Flags().String("demuxer-mkv-probe-video-duration", "", "When  opening  the file, seek to the end of it, and check what timestamp the last video packet has, and report that as file duration")
	rootCmd.Flags().String("demuxer-mkv-subtitle-preroll", "", "Try harder to show embedded soft subtitles when seeking somewhere")
	rootCmd.Flags().String("demuxer-mkv-subtitle-preroll-secs", "", "See --demuxer-mkv-subtitle-preroll")
	rootCmd.Flags().String("demuxer-mkv-subtitle-preroll-secs-index", "", "See --demuxer-mkv-subtitle-preroll")
	rootCmd.Flags().String("demuxer-rawaudio-channels", "", "Number of channels (or channel layout) if --demuxer=rawaudio is used")
	rootCmd.Flags().String("demuxer-rawaudio-format", "", "Sample format for --demuxer=rawaudio")
	rootCmd.Flags().String("demuxer-rawaudio-rate", "", "Sample rate for --demuxer=rawaudio")
	rootCmd.Flags().String("demuxer-rawvideo-codec", "", "Set the video codec instead of selecting the rawvideo codec when using --demuxer=rawvideo")
	rootCmd.Flags().String("demuxer-rawvideo-format", "", "Color space (fourcc) in hex or string for --demuxer=rawvideo")
	rootCmd.Flags().String("demuxer-rawvideo-fps", "", "Rate in frames per second for --demuxer=rawvideo")
	rootCmd.Flags().String("demuxer-rawvideo-h", "", "Image dimension in pixels for --demuxer=rawvideo")
	rootCmd.Flags().String("demuxer-rawvideo-mp-format", "", "Color space by internal video format for --demuxer=rawvideo")
	rootCmd.Flags().String("demuxer-rawvideo-size", "", "Frame size in bytes when using --demuxer=rawvideo")
	rootCmd.Flags().String("demuxer-rawvideo-w", "", "Image dimension in pixels for --demuxer=rawvideo")
	rootCmd.Flags().String("demuxer-readahead-secs", "", "If --demuxer-thread is enabled, this controls how much the demuxer should buffer ahead in seconds")
	rootCmd.Flags().String("demuxer-seekable-cache", "", "Debugging option to control whether seeking can use the demuxer cache")
	rootCmd.Flags().String("demuxer-termination-timeout", "", "Number of seconds the player should wait to shutdown the demuxer")
	rootCmd.Flags().String("demuxer-thread", "", "Run the demuxer in a separate thread, and let it prefetch a certain amount of packets")
	rootCmd.Flags().String("directory-filter-types", "", "Media file types to filter when opening directory")
	rootCmd.Flags().String("directory-filter-types-add", "", "")
	rootCmd.Flags().String("directory-filter-types-append", "", "")
	rootCmd.Flags().String("directory-filter-types-clr", "", "")
	rootCmd.Flags().String("directory-filter-types-del", "", "")
	rootCmd.Flags().String("directory-filter-types-pre", "", "")
	rootCmd.Flags().String("directory-filter-types-remove", "", "")
	rootCmd.Flags().String("directory-filter-types-set", "", "")
	rootCmd.Flags().String("directory-filter-types-toggle", "", "")
	rootCmd.Flags().String("directory-mode", "", "When opening a directory, open subdirectories lazily, recursively or not at all")
	rootCmd.Flags().String("display-fps-override", "", "Set the display FPS used with the --video-sync=display-* modes")
	rootCmd.Flags().String("display-tags", "", "Set the list of tags that should be displayed on the terminal and stats")
	rootCmd.Flags().String("display-tags-add", "", "")
	rootCmd.Flags().String("display-tags-append", "", "")
	rootCmd.Flags().String("display-tags-clr", "", "")
	rootCmd.Flags().String("display-tags-del", "", "")
	rootCmd.Flags().String("display-tags-pre", "", "")
	rootCmd.Flags().String("display-tags-remove", "", "")
	rootCmd.Flags().String("display-tags-set", "", "")
	rootCmd.Flags().String("display-tags-toggle", "", "")
	rootCmd.Flags().String("dither", "", "Select dithering algorithm")
	rootCmd.Flags().String("dither-depth", "", "Set dither target depth to N")
	rootCmd.Flags().String("dither-size-fruit", "", "Set the size of the dither matrix")
	rootCmd.Flags().String("drag-and-drop", "", "Controls the default behavior of drag and drop on platforms that support this")
	rootCmd.Flags().String("drm-connector", "", "Select the connector to use")
	rootCmd.Flags().String("drm-device", "", "Select the DRM device file to use")
	rootCmd.Flags().String("drm-draw-plane", "", "Select the DRM plane to which video and OSD is drawn to, under normal circumstances")
	rootCmd.Flags().String("drm-draw-surface-size", "", "Sets the size of the surface used on the draw plane")
	rootCmd.Flags().String("drm-drmprime-video-plane", "", "Select the DRM plane to use for video with the drmprime-overlay hwdec interop")
	rootCmd.Flags().String("drm-format", "", "Select the DRM format to use")
	rootCmd.Flags().String("drm-mode", "", "Mode to use")
	rootCmd.Flags().String("drm-vrr-enabled", "", "Toggle use of Variable Refresh Rate (VRR), aka Freesync or Adaptive Sync on compatible systems")
	rootCmd.Flags().String("dscale", "", "Like --scale, but apply these filters on downscaling instead")
	rootCmd.Flags().String("dscale-antiring", "", "Set the antiringing strength")
	rootCmd.Flags().String("dscale-blur", "", "Kernel scaling factor")
	rootCmd.Flags().String("dscale-clamp", "", "Specifies a weight bias to multiply into negative coefficients")
	rootCmd.Flags().String("dscale-param1", "", "Set filter parameters")
	rootCmd.Flags().String("dscale-param2", "", "Set filter parameters")
	rootCmd.Flags().String("dscale-radius", "", "Set radius for tunable filters, must be a float number between 0.5 and 16.0")
	rootCmd.Flags().String("dscale-taper", "", "Kernel/window taper factor")
	rootCmd.Flags().String("dscale-window", "", "Choose a custom windowing function for the kernel")
	rootCmd.Flags().String("dscale-wparam", "", "")
	rootCmd.Flags().String("dscale-wtaper", "", "Kernel/window taper factor")
	rootCmd.Flags().String("dump-stats", "", "Write certain statistics to the given file")
	rootCmd.Flags().String("dvbin-card", "", "Specifies using card number 0-15")
	rootCmd.Flags().String("dvbin-channel-switch-offset", "", "This value is not meant for setting via configuration, but used in channel switching")
	rootCmd.Flags().String("dvbin-file", "", "Instructs mpv to read the channels list from <filename>")
	rootCmd.Flags().String("dvbin-full-transponder", "", "Apply no filters on program PIDs, only tune to frequency and pass full transponder to demuxer")
	rootCmd.Flags().String("dvbin-prog", "", "This defines the program to tune to")
	rootCmd.Flags().String("dvbin-timeout", "", "Maximum number of seconds to wait when trying to tune a frequency before giving up")
	rootCmd.Flags().String("dvd-angle", "", "Some DVDs contain scenes that can be viewed from multiple angles")
	rootCmd.Flags().String("dvd-device", "", "Specify the DVD device or .iso filename")
	rootCmd.Flags().String("dvd-speed", "", "Try to limit DVD speed")
	rootCmd.Flags().String("edition", "", "(Matroska files only) Specify the edition (set of chapters) to use, where 0 is the first")
	rootCmd.Flags().String("egl-config-id", "", "(EGL only) Select EGLConfig with specific EGL_CONFIG_ID")
	rootCmd.Flags().String("egl-output-format", "", "(EGL only) Select a specific EGL output format to utilize for OpenGL rendering")
	rootCmd.Flags().String("embeddedfonts", "", "Use fonts embedded in Matroska container files and ASS scripts")
	rootCmd.Flags().String("end", "", "Stop at given time")
	rootCmd.Flags().String("error-diffusion", "", "The error diffusion kernel to use when --dither=error-diffusion is set")
	rootCmd.Flags().String("external-file", "", "CLI/config file only alias for --external-files-append")
	rootCmd.Flags().String("external-files", "", "Load a file and add all of its tracks")
	rootCmd.Flags().String("external-files-add", "", "")
	rootCmd.Flags().String("external-files-append", "", "")
	rootCmd.Flags().String("external-files-clr", "", "")
	rootCmd.Flags().String("external-files-del", "", "")
	rootCmd.Flags().String("external-files-pre", "", "")
	rootCmd.Flags().String("external-files-remove", "", "")
	rootCmd.Flags().String("external-files-set", "", "")
	rootCmd.Flags().String("external-files-toggle", "", "")
	rootCmd.Flags().String("fbo-format", "", "Selects the internal format of textures used for FBOs")
	rootCmd.Flags().String("focus-on", "", "(macOS only) Focus the video window and make it the front most window on specific events")
	rootCmd.Flags().String("force-media-title", "", "Force the contents of the media-title property to this value.")
	rootCmd.Flags().String("force-render", "", "Forces mpv to always render frames regardless of the visibility of the window")
	rootCmd.Flags().String("force-rgba-osd-rendering", "", "Change how some video outputs render the OSD and text subtitles")
	rootCmd.Flags().String("force-seekable", "", "If the player thinks that the media is not seekable (e.g. playing from a pipe, or it's an http stream with a server that doesn't support range requests), seeking will be disabled")
	rootCmd.Flags().String("force-window", "", "Create a video output window even if there is no video")
	rootCmd.Flags().String("force-window-position", "", "Forcefully move mpv's video output window to default location whenever there is a change in video parameters, video stream or file")
	rootCmd.Flags().String("framedrop", "", "Skip displaying some frames to maintain A/V sync on slow systems, or playing high framerate video on video outputs that have an upper framerate limit")
	rootCmd.Flags().String("frames", "", "Play/convert only first <number> video frames, then quit")
	rootCmd.Flags().String("fs", "", "fullscreen playback")
	rootCmd.Flags().String("fs-screen", "", "In multi-monitor configurations (i.e. a single desktop that spans across multiple displays), this option tells mpv which screen to go fullscreen to")
	rootCmd.Flags().String("fs-screen-name", "", "In multi-monitor configurations, this option tells mpv which screen to go fullscreen to based on the screen name from the video backend")
	rootCmd.Flags().String("fullscreen", "", "Fullscreen playback")
	rootCmd.Flags().String("gamma", "", "Adjust the gamma of the video signal")
	rootCmd.Flags().String("gamma-factor", "", "Set an additional raw gamma factor")
	rootCmd.Flags().String("gamut-mapping-mode", "", "Specifies the algorithm used for reducing the gamut of images for the target display, after any tone mapping is done")
	rootCmd.Flags().String("gapless-audio", "", "Try to play consecutive audio files with no silence or disruption at the point of file change")
	rootCmd.Flags().String("geometry", "", "Adjust the initial window position or size")
	rootCmd.Flags().StringSlice("glsl-shader", nil, "Custom GLSL hooks")
	rootCmd.Flags().String("glsl-shader-opts", "", "Specifies the options to use for tunable shader parameters")
	rootCmd.Flags().String("glsl-shader-opts-add", "", "")
	rootCmd.Flags().String("glsl-shader-opts-append", "", "")
	rootCmd.Flags().String("glsl-shader-opts-clr", "", "")
	rootCmd.Flags().String("glsl-shader-opts-del", "", "")
	rootCmd.Flags().String("glsl-shader-opts-remove", "", "")
	rootCmd.Flags().String("glsl-shader-opts-set", "", "")
	rootCmd.Flags().String("glsl-shaders", "", "Custom GLSL hooks")
	rootCmd.Flags().String("glsl-shaders-add", "", "")
	rootCmd.Flags().String("glsl-shaders-append", "", "")
	rootCmd.Flags().String("glsl-shaders-clr", "", "")
	rootCmd.Flags().String("glsl-shaders-del", "", "")
	rootCmd.Flags().String("glsl-shaders-pre", "", "")
	rootCmd.Flags().String("glsl-shaders-remove", "", "")
	rootCmd.Flags().String("glsl-shaders-set", "", "")
	rootCmd.Flags().String("glsl-shaders-toggle", "", "")
	rootCmd.Flags().String("gpu-api", "", "Specify a priority list of accepted graphics APIs")
	rootCmd.Flags().String("gpu-api-add", "", "")
	rootCmd.Flags().String("gpu-api-append", "", "")
	rootCmd.Flags().String("gpu-api-clr", "", "")
	rootCmd.Flags().String("gpu-api-help", "", "")
	rootCmd.Flags().String("gpu-api-pre", "", "")
	rootCmd.Flags().String("gpu-api-remove", "", "")
	rootCmd.Flags().String("gpu-api-set", "", "")
	rootCmd.Flags().String("gpu-api-toggle", "", "")
	rootCmd.Flags().String("gpu-context", "", "Specify a priority list of the GPU contexts to be used")
	rootCmd.Flags().String("gpu-context-add", "", "")
	rootCmd.Flags().String("gpu-context-append", "", "")
	rootCmd.Flags().String("gpu-context-clr", "", "")
	rootCmd.Flags().String("gpu-context-help", "", "")
	rootCmd.Flags().String("gpu-context-pre", "", "")
	rootCmd.Flags().String("gpu-context-remove", "", "")
	rootCmd.Flags().String("gpu-context-set", "", "")
	rootCmd.Flags().String("gpu-context-toggle", "", "")
	rootCmd.Flags().String("gpu-debug", "", "Enables GPU debugging")
	rootCmd.Flags().String("gpu-dumb-mode", "", "This mode is extremely restricted, and will disable most extended features. That includes high quality scalers and custom shaders!")
	rootCmd.Flags().String("gpu-hwdec-interop", "", "This option is for troubleshooting hwdec interop issues")
	rootCmd.Flags().String("gpu-shader-cache", "", "Store and load compiled GLSL shaders in the cache directory")
	rootCmd.Flags().String("gpu-shader-cache-dir", "", "The directory where gpu shader cache is stored")
	rootCmd.Flags().String("gpu-sw", "", "Continue even if a software renderer is detected")
	rootCmd.Flags().String("gpu-tex-pad-x", "", "Enlarge the video source textures by this many pixels")
	rootCmd.Flags().String("gpu-tex-pad-y", "", "Enlarge the video source textures by this many pixels")
	rootCmd.Flags().String("h", "", "Show short summary of options")
	rootCmd.Flags().String("hdr-compute-peak", "", "Compute the HDR peak and frame average brightness per-frame instead of relying on tagged metadata")
	rootCmd.Flags().String("hdr-contrast-recovery", "", "Enables the HDR contrast recovery algorithm, which is to designed to enhance contrast of HDR video after tone mapping")
	rootCmd.Flags().String("hdr-contrast-smoothness", "", "Enables the HDR contrast recovery algorithm, which is to designed to enhance contrast of HDR video after tone mapping")
	rootCmd.Flags().String("hdr-peak-decay-rate", "", "The decay rate used for the HDR peak detection algorithm")
	rootCmd.Flags().String("hdr-peak-percentile", "", "Which percentile of the input image brightness histogram to consider as the true peak of the scene")
	rootCmd.Flags().String("hdr-scene-threshold-high", "", "The upper thresholds (in dB) for a brightness difference to be considered a scene change")
	rootCmd.Flags().String("hdr-scene-threshold-low", "", "The lower thresholds (in dB) for a brightness difference to be considered a scene change")
	rootCmd.Flags().String("help", "", "Show short summary of options")
	rootCmd.Flags().String("hidpi-window-scale", "", "Scale the window size according to the backing DPI scale factor from the OS")
	rootCmd.Flags().String("hls-bitrate", "", "If HLS streams are played, this option controls what streams are selected by default")
	rootCmd.Flags().String("hr-seek", "", "Select when to use precise seeks that are not limited to keyframes")
	rootCmd.Flags().String("hr-seek-demuxer-offset", "", "This option exists to work around failures to do precise seeks (as in --hr-seek) caused by bugs or limitations in the demuxers for some file formats")
	rootCmd.Flags().String("hr-seek-framedrop", "", "Allow the video decoder to drop frames during seek, if these frames are before the seek target")
	rootCmd.Flags().String("http-header-fields", "", "Set custom HTTP fields when accessing HTTP stream")
	rootCmd.Flags().String("http-header-fields-add", "", "")
	rootCmd.Flags().String("http-header-fields-append", "", "")
	rootCmd.Flags().String("http-header-fields-clr", "", "")
	rootCmd.Flags().String("http-header-fields-del", "", "")
	rootCmd.Flags().String("http-header-fields-pre", "", "")
	rootCmd.Flags().String("http-header-fields-remove", "", "")
	rootCmd.Flags().String("http-header-fields-set", "", "")
	rootCmd.Flags().String("http-header-fields-toggle", "", "")
	rootCmd.Flags().String("http-proxy", "", "URL of the HTTP/HTTPS proxy")
	rootCmd.Flags().String("hue", "", "Adjust the hue of the video signal")
	rootCmd.Flags().String("hwdec", "", "Specify the hardware video decoding API that should be used if possible")
	rootCmd.Flags().String("hwdec-add", "", "")
	rootCmd.Flags().String("hwdec-append", "", "")
	rootCmd.Flags().String("hwdec-clr", "", "")
	rootCmd.Flags().String("hwdec-codecs", "", "")
	rootCmd.Flags().String("hwdec-del", "", "")
	rootCmd.Flags().String("hwdec-extra-frames", "", "")
	rootCmd.Flags().String("hwdec-image-format", "", "")
	rootCmd.Flags().String("hwdec-pre", "", "")
	rootCmd.Flags().String("hwdec-remove", "", "")
	rootCmd.Flags().String("hwdec-set", "", "")
	rootCmd.Flags().String("hwdec-software-fallback", "", "Fallback to software decoding if the hardware-accelerated decoder fails")
	rootCmd.Flags().String("hwdec-toggle", "", "")
	rootCmd.Flags().String("icc-3dlut-size", "", "Size of the 3D LUT generated from the ICC profile in each dimension")
	rootCmd.Flags().String("icc-cache", "", "Store and load 3DLUTs created from the ICC profile on disk in the cache directory")
	rootCmd.Flags().String("icc-cache-dir", "", "The directory where icc cache is stored")
	rootCmd.Flags().String("icc-force-contrast", "", "Override the target device's detected contrast ratio by a specific value")
	rootCmd.Flags().String("icc-intent", "", "Specifies the ICC intent used for the color transformation")
	rootCmd.Flags().String("icc-profile", "", "Load an ICC profile and use it to transform video RGB to screen output")
	rootCmd.Flags().String("icc-profile-auto", "", "Automatically select the ICC display profile currently specified by the display settings of the operating system")
	rootCmd.Flags().String("icc-use-luma", "", "Use ICC profile luminance value")
	rootCmd.Flags().String("idle", "", "Makes mpv wait idly instead of quitting when there is no file to play")
	rootCmd.Flags().String("ignore-path-in-watch-later-config", "", "Ignore path (i.e. use filename only) when using watch later feature")
	rootCmd.Flags().String("image-display-duration", "", "If the current file is an image, play the image for the given amount of seconds")
	rootCmd.Flags().String("image-exts", "", "Image file extensions to try to match when using --cover-art-auto, --autocreate-playlist or --directory-filter-types")
	rootCmd.Flags().String("image-exts-add", "", "")
	rootCmd.Flags().String("image-exts-append", "", "")
	rootCmd.Flags().String("image-exts-clr", "", "")
	rootCmd.Flags().String("image-exts-del", "", "")
	rootCmd.Flags().String("image-exts-pre", "", "")
	rootCmd.Flags().String("image-exts-remove", "", "")
	rootCmd.Flags().String("image-exts-set", "", "")
	rootCmd.Flags().String("image-exts-toggle", "", "")
	rootCmd.Flags().String("image-lut", "", "Specifies a custom LUT file (in Adobe .cube format) to apply to the colors during image decoding")
	rootCmd.Flags().String("image-lut-type", "", "Controls the interpretation of color values fed to and from the LUT specified as --image-lut")
	rootCmd.Flags().String("image-subs-video-resolution", "", "Override the image subtitle resolution with the video resolution")
	rootCmd.Flags().String("include", "", "Specify configuration file to be parsed after the default ones")
	rootCmd.Flags().String("index", "", "Controls how to seek in files")
	rootCmd.Flags().String("initial-audio-sync", "", "When starting a video file or after events such as seeking, mpv will by default modify the audio stream to make it start from the same timestamp as video, by either inserting silence at the start or cutting away the first samples")
	rootCmd.Flags().String("input-ar-delay", "", "Delay in milliseconds before we start to autorepeat a key")
	rootCmd.Flags().String("input-ar-rate", "", "Number of key presses to generate per second on autorepeat")
	rootCmd.Flags().String("input-builtin-bindings", "", "Enable loading of built-in key bindings during start-up")
	rootCmd.Flags().String("input-builtin-dragging", "", "Enable the built-in window-dragging behavior")
	rootCmd.Flags().String("input-cmdlist", "", "Prints all commands that can be bound to keys")
	rootCmd.Flags().String("input-commands", "", "Define a list of commands for mpv to run")
	rootCmd.Flags().String("input-commands-add", "", "")
	rootCmd.Flags().String("input-commands-append", "", "")
	rootCmd.Flags().String("input-commands-clr", "", "")
	rootCmd.Flags().String("input-commands-del", "", "")
	rootCmd.Flags().String("input-commands-pre", "", "")
	rootCmd.Flags().String("input-commands-remove", "", "")
	rootCmd.Flags().String("input-commands-set", "", "")
	rootCmd.Flags().String("input-commands-toggle", "", "")
	rootCmd.Flags().String("input-conf", "", "Specify input configuration file other than the default location in the mpv configuration directory")
	rootCmd.Flags().String("input-cursor", "", "Permit mpv to receive pointer events reported by the video output driver")
	rootCmd.Flags().String("input-cursor-passthrough", "", "Tell the backend windowing system to allow pointer events to passthrough the mpv window")
	rootCmd.Flags().String("input-default-bindings", "", "Enable default-level (\"weak\") key bindings")
	rootCmd.Flags().String("input-doubleclick-time", "", "Time in milliseconds to recognize two consecutive button presses as a double-click")
	rootCmd.Flags().String("input-dragging-deadzone", "", "Begin the built-in window dragging when the mouse moves outside a deadzone of N pixels while the mouse button is being held down")
	rootCmd.Flags().String("input-ime", "", "Enable keyboard input via an active input method (IME) connected to the VO")
	rootCmd.Flags().String("input-ipc-client", "", "Connect a single IPC client to the given FD")
	rootCmd.Flags().String("input-ipc-server", "", "Enable the IPC support and create the listening socket at the given path")
	rootCmd.Flags().String("input-key-fifo-size", "", "Specify the size of the FIFO that buffers key events")
	rootCmd.Flags().String("input-keylist", "", "Prints all keys that can be bound to commands")
	rootCmd.Flags().String("input-media-keys", "", "On systems where mpv can choose between receiving media keys or letting the system handle them - this option controls whether mpv should receive them")
	rootCmd.Flags().String("input-preprocess-wheel", "", "Preprocess WHEEL_* events so that while scrolling on the horizontal or vertical direction, the events aren't generated for another direction even when the two directions are scrolled together")
	rootCmd.Flags().String("input-right-alt-gr", "", "(macOS and Windows only) Use the right Alt key as Alt Gr to produce special characters")
	rootCmd.Flags().String("input-terminal", "", "--input-terminal=no prevents the player from reading key events from standard input")
	rootCmd.Flags().String("input-test", "", "Input test mode")
	rootCmd.Flags().String("input-touch-emulate-mouse", "", "When multi-touch support is enabled (either required by the platform, or enabled by --native-touch), emulate mouse move and button presses for the touch events")
	rootCmd.Flags().String("input-vo-keyboard", "", "Disable all keyboard input on for VOs which can't participate in proper keyboard input dispatching")
	rootCmd.Flags().String("interpolation", "", "Reduce stuttering caused by mismatches in the video fps and display refresh rate")
	rootCmd.Flags().String("interpolation-preserve", "", "Preserve the previous frames' interpolated results even when renderer parameters are changed - with the exception of options related to cropping and video placement, which always invalidate the cache")
	rootCmd.Flags().String("interpolation-threshold", "", "Threshold below which frame ratio interpolation gets disabled")
	rootCmd.Flags().String("inverse-tone-mapping", "", "If set, allows inverse tone mapping (expanding SDR to HDR)")
	rootCmd.Flags().String("jack-autostart", "", "Automatically start jackd if necessary")
	rootCmd.Flags().String("jack-connect", "", "Automatically create connections to output ports")
	rootCmd.Flags().String("jack-name", "", "Client name that is passed to JACK")
	rootCmd.Flags().String("jack-port", "", "Connects to the ports with the given name")
	rootCmd.Flags().String("jack-std-channel-layout", "", "Select the standard channel layout")
	rootCmd.Flags().String("js-memory-report", "", "Enable memory reporting for javascript scripts in the stats overlay")
	rootCmd.Flags().String("keep-open", "", "Do not terminate when playing or seeking beyond the end of the file, and there is no next file to be played (and --loop is not used). Instead, pause the player.")
	rootCmd.Flags().String("keep-open-pause", "", "If set to no, instead of pausing when --keep-open is active, just stop at end of file and continue playing forward when you seek backwards until end where it stops again")
	rootCmd.Flags().String("keepaspect", "", "--keepaspect=no will always stretch the video to window size, and will disable the window manager hints that force the window aspect ratio")
	rootCmd.Flags().String("keepaspect-window", "", "--keepaspect-window=yes (the default) will lock the window size to the video aspect")
	rootCmd.Flags().String("lavfi-complex", "", "Set a \"complex\" libavfilter filter, which means a single filter graph can take input from multiple source audio and video tracks")
	rootCmd.Flags().String("length", "", "Stop after a given time relative to the start time")
	rootCmd.Flags().String("libplacebo-opts", "", "Passes extra raw option to the libplacebo rendering backend")
	rootCmd.Flags().String("libplacebo-opts-add", "", "")
	rootCmd.Flags().String("libplacebo-opts-append", "", "")
	rootCmd.Flags().String("libplacebo-opts-clr", "", "")
	rootCmd.Flags().String("libplacebo-opts-del", "", "")
	rootCmd.Flags().String("libplacebo-opts-remove", "", "")
	rootCmd.Flags().String("libplacebo-opts-set", "", "")
	rootCmd.Flags().String("linear-downscaling", "", "Scale in linear light when downscaling")
	rootCmd.Flags().String("linear-upscaling", "", "Scale in linear light when upscaling")
	rootCmd.Flags().Bool("list-options", false, "list all mpv options")
	rootCmd.Flags().String("list-properties", "", "Print a list of the available properties")
	rootCmd.Flags().String("list-protocols", "", "Print a list of the supported protocols")
	rootCmd.Flags().String("load-auto-profiles", "", "Enable the builtin script that does auto profiles")
	rootCmd.Flags().String("load-commands", "", "Enable the built-in script to enter commands in the console")
	rootCmd.Flags().String("load-console", "", "Enable the built-in script to handle textual input")
	rootCmd.Flags().String("load-positioning", "", "Enable the builtin script that provides various keybindings to pan videos and images")
	rootCmd.Flags().String("load-scripts", "", "If set to no, don't auto-load scripts from the scripts configuration subdirectory")
	rootCmd.Flags().String("load-select", "", "Enable the builtin script that lets you select from lists of items")
	rootCmd.Flags().String("load-stats-overlay", "", "Enable the builtin script that shows useful playback information on a key binding")
	rootCmd.Flags().String("load-unsafe-playlists", "", "Load URLs from playlists which are considered unsafe")
	rootCmd.Flags().String("log-file", "", "Opens the given path for writing, and print log messages to it")
	rootCmd.Flags().String("loop", "", "Loop a single file N times")
	rootCmd.Flags().String("loop-file", "", "Loop a single file N times")
	rootCmd.Flags().String("loop-playlist", "", "Loops playback N times")
	rootCmd.Flags().String("lut", "", "Specifies a custom LUT (in Adobe .cube format) to apply to the colors as part of color conversion")
	rootCmd.Flags().String("lut-type", "", "Controls the interpretation of color values fed to and from the LUT specified as --lut")
	rootCmd.Flags().String("mc", "", "Maximum A-V sync correction per frame")
	rootCmd.Flags().String("media-controls", "", "(Windows only) Enable integration of media control interface SystemMediaTransportControls")
	rootCmd.Flags().String("merge-files", "", "Pretend that all files passed to mpv are concatenated into a single, big file")
	rootCmd.Flags().String("metadata-codepage", "", "Codepage for various input metadata")
	rootCmd.Flags().String("mf-fps", "", "Framerate used when decoding from multiple PNG or JPEG files with mf://")
	rootCmd.Flags().String("mf-type", "", "Input file type for mf://")
	rootCmd.Flags().String("monitoraspect", "", "Set the aspect ratio of your monitor or TV screen")
	rootCmd.Flags().String("monitorpixelaspect", "", "Set the aspect of a single pixel of your monitor or TV screen")
	rootCmd.Flags().String("msg-color", "", "Enable colorful console output on terminals")
	rootCmd.Flags().String("msg-level", "", "Control verbosity directly for each module")
	rootCmd.Flags().String("msg-module", "", "Prepend module name to each console message")
	rootCmd.Flags().String("msg-time", "", "Prepend timing information to each console message")
	rootCmd.Flags().String("mute", "", "Set startup audio mute status")
	rootCmd.Flags().String("native-fs", "", "(macOS only) Uses the native fullscreen mechanism of the OS")
	rootCmd.Flags().String("native-keyrepeat", "", "Use system settings for keyrepeat delay and rate, instead of --input-ar-delay and --input-ar-rate")
	rootCmd.Flags().String("native-touch", "", "(Windows only) For platforms which send emulated mouse inputs for touch-unaware clients, such as Windows, use system native touch events, instead of receiving them as emulated mouse events")
	rootCmd.Flags().String("network-timeout", "", "Specify the network timeout in seconds")
	rootCmd.Flags().Bool("no-audio", false, "do not play sound")
	rootCmd.Flags().Bool("no-config", false, "Do not load default configuration or any user files")
	rootCmd.Flags().Bool("no-video", false, "do not play video")
	rootCmd.Flags().String("o", "", "Enables encoding mode and specifies the output file name")
	rootCmd.Flags().String("oac", "", "Specifies the output audio codec")
	rootCmd.Flags().String("oacopts", "", "Specifies the output audio codec options for libavcodec")
	rootCmd.Flags().String("oacopts-add", "", "")
	rootCmd.Flags().String("oacopts-append", "", "")
	rootCmd.Flags().String("oacopts-clr", "", "")
	rootCmd.Flags().String("oacopts-del", "", "")
	rootCmd.Flags().String("oacopts-remove", "", "")
	rootCmd.Flags().String("oacopts-set", "", "")
	rootCmd.Flags().String("ocopy-metadata", "", "Copy metadata from input files to output files when encoding")
	rootCmd.Flags().String("of", "", "Specifies the output format")
	rootCmd.Flags().String("ofopts", "", "Specifies the output format options for libavformat")
	rootCmd.Flags().String("ofopts-add", "", "")
	rootCmd.Flags().String("ofopts-append", "", "")
	rootCmd.Flags().String("ofopts-clr", "", "")
	rootCmd.Flags().String("ofopts-del", "", "")
	rootCmd.Flags().String("ofopts-remove", "", "")
	rootCmd.Flags().String("ofopts-set", "", "")
	rootCmd.Flags().String("on-all-workspaces", "", "(X11 and macOS only) Show the video window on all virtual desktops")
	rootCmd.Flags().String("ontop", "", "Makes the player window stay on top of other windows")
	rootCmd.Flags().String("ontop-level", "", "(macOS only) Sets the level of an on-top window")
	rootCmd.Flags().String("openal-direct-channels", "", "Enable OpenAL Soft's direct channel extension when available to avoid tinting the sound with ambisonics or HRTF")
	rootCmd.Flags().String("openal-num-buffers", "", "Specify the number of audio buffers to use")
	rootCmd.Flags().String("openal-num-samples", "", "Specify the number of complete samples to use for each buffer")
	rootCmd.Flags().String("opengl-check-pattern-a", "", "")
	rootCmd.Flags().String("opengl-check-pattern-b", "", "")
	rootCmd.Flags().String("opengl-early-flush", "", "Call glFlush() after rendering a frame and before attempting to display it")
	rootCmd.Flags().String("opengl-es", "", "Controls which type of OpenGL context will be accepted")
	rootCmd.Flags().String("opengl-glfinish", "", "Call glFinish() before swapping buffers")
	rootCmd.Flags().String("opengl-pbo", "", "Enable use of PBOs")
	rootCmd.Flags().String("opengl-rectangle-textures", "", "Force use of rectangle textures")
	rootCmd.Flags().String("opengl-swapinterval", "", "Interval in displayed frames between two buffer swaps")
	rootCmd.Flags().String("opengl-waitvsync", "", "Call glXWaitVideoSyncSGI after each buffer swap")
	rootCmd.Flags().String("orawts", "", "Copies input pts to the output video")
	rootCmd.Flags().String("ordered-chapters", "", "Enable support for Matroska ordered chapters")
	rootCmd.Flags().String("ordered-chapters-files", "", "Loads the given file as playlist, and tries to use the files contained in it as reference files when opening a Matroska file that uses ordered chapters")
	rootCmd.Flags().String("oremove-metadata", "", "Specifies metadata to exclude from the output file when copying from the input file")
	rootCmd.Flags().String("oremove-metadata-add", "", "")
	rootCmd.Flags().String("oremove-metadata-append", "", "")
	rootCmd.Flags().String("oremove-metadata-clr", "", "")
	rootCmd.Flags().String("oremove-metadata-del", "", "")
	rootCmd.Flags().String("oremove-metadata-pre", "", "")
	rootCmd.Flags().String("oremove-metadata-remove", "", "")
	rootCmd.Flags().String("oremove-metadata-set", "", "")
	rootCmd.Flags().String("oremove-metadata-toggle", "", "")
	rootCmd.Flags().String("osc", "", "Whether to load the on-screen-controller")
	rootCmd.Flags().String("osd-align-x", "", "Control to which corner of the screen OSD should be aligned to")
	rootCmd.Flags().String("osd-align-y", "", "Vertical position (default: top). Details see --osd-align-x")
	rootCmd.Flags().String("osd-back-color", "", "See --sub-color. Color used for OSD text background")
	rootCmd.Flags().String("osd-bar", "", "Enable display of the OSD bar")
	rootCmd.Flags().String("osd-bar-align-x", "", "Position of the OSD bar")
	rootCmd.Flags().String("osd-bar-align-y", "", "Position of the OSD bar")
	rootCmd.Flags().String("osd-bar-border-size", "", "alias for --osd-bar-outline-size")
	rootCmd.Flags().String("osd-bar-h", "", "Height of the OSD bar, in percentage of the screen height")
	rootCmd.Flags().String("osd-bar-marker-min-size", "", "Minimum OSD bar marker size")
	rootCmd.Flags().String("osd-bar-marker-scale", "", "Factor for the OSD bar marker size relative to the OSD bar outline size")
	rootCmd.Flags().String("osd-bar-marker-style", "", "Set the OSD bar marker style")
	rootCmd.Flags().String("osd-bar-outline-size", "", "Size of the outline of the OSD bar in scaled pixels")
	rootCmd.Flags().String("osd-bar-w", "", "Width of the OSD bar, in percentage of the screen width")
	rootCmd.Flags().String("osd-blur", "", "Gaussian blur factor applied to the OSD font border")
	rootCmd.Flags().String("osd-bold", "", "Format text on bold")
	rootCmd.Flags().String("osd-border-color", "", "alias for --osd-outline-color")
	rootCmd.Flags().String("osd-border-size", "", "alias for --osd-outline-size")
	rootCmd.Flags().String("osd-border-style", "", "See --sub-border-style. Style used for OSD text border")
	rootCmd.Flags().String("osd-color", "", "Specify the color used for OSD")
	rootCmd.Flags().String("osd-duration", "", "Set the duration of the OSD messages in ms")
	rootCmd.Flags().String("osd-font", "", "Specify font to use for OSD")
	rootCmd.Flags().String("osd-font-provider", "", "See --sub-font-provider for details and accepted values")
	rootCmd.Flags().String("osd-font-size", "", "Specify the OSD font size")
	rootCmd.Flags().String("osd-fonts-dir", "", "See --sub-fonts-dir for details")
	rootCmd.Flags().String("osd-fractions", "", "Show OSD times with fractions of seconds")
	rootCmd.Flags().String("osd-italic", "", "Format text on italic")
	rootCmd.Flags().String("osd-justify", "", "")
	rootCmd.Flags().String("osd-level", "", "Specifies which mode the OSD should start in")
	rootCmd.Flags().String("osd-margin-x", "", "Left and right screen margin for the OSD in scaled pixels")
	rootCmd.Flags().String("osd-margin-y", "", "Top and bottom screen margin for the OSD in scaled pixels")
	rootCmd.Flags().String("osd-msg1", "", "Show this string as message on OSD with OSD level 1")
	rootCmd.Flags().String("osd-msg2", "", "Similar to --osd-msg1, but for OSD level 2")
	rootCmd.Flags().String("osd-msg3", "", "Similar to --osd-msg1, but for OSD level 3")
	rootCmd.Flags().String("osd-on-seek", "", "Set what is displayed on the OSD during seeks")
	rootCmd.Flags().String("osd-outline-color", "", "See --sub-color. Color used for the OSD font outline")
	rootCmd.Flags().String("osd-outline-size", "", "Size of the OSD font outline in scaled pixels")
	rootCmd.Flags().String("osd-playing-msg", "", "Show a message on OSD when playback starts")
	rootCmd.Flags().String("osd-playing-msg-duration", "", "Set the duration of osd-playing-msg in ms")
	rootCmd.Flags().String("osd-playlist-entry", "", "Whether to display the media title, filename, or both")
	rootCmd.Flags().String("osd-scale", "", "OSD font size multiplier, multiplied with --osd-font-size value")
	rootCmd.Flags().String("osd-scale-by-window", "", "Whether to scale the OSD with the window size")
	rootCmd.Flags().String("osd-selected-color", "", "The color of the selected item in lists")
	rootCmd.Flags().String("osd-selected-outline-color", "", "The outline color of the selected item in lists")
	rootCmd.Flags().String("osd-shadow-color", "", "alias for --osd-back-color")
	rootCmd.Flags().String("osd-shadow-offset", "", "Displacement of the OSD shadow in scaled pixels")
	rootCmd.Flags().String("osd-spacing", "", "Horizontal OSD/sub font spacing in scaled pixels")
	rootCmd.Flags().String("osd-status-msg", "", "Show a custom string during playback instead of the standard status text")
	rootCmd.Flags().String("oset-metadata", "", "Specifies metadata to include in the output file")
	rootCmd.Flags().String("oset-metadata-add", "", "")
	rootCmd.Flags().String("oset-metadata-append", "", "")
	rootCmd.Flags().String("oset-metadata-clr", "", "")
	rootCmd.Flags().String("oset-metadata-del", "", "")
	rootCmd.Flags().String("oset-metadata-remove", "", "")
	rootCmd.Flags().String("oset-metadata-set", "", "")
	rootCmd.Flags().String("ovc", "", "Specifies the output video codec")
	rootCmd.Flags().String("ovcopts", "", "Specifies the output video codec options for libavcodec")
	rootCmd.Flags().String("ovcopts-add", "", "")
	rootCmd.Flags().String("ovcopts-append", "", "")
	rootCmd.Flags().String("ovcopts-clr", "", "")
	rootCmd.Flags().String("ovcopts-del", "", "")
	rootCmd.Flags().String("ovcopts-remove", "", "")
	rootCmd.Flags().String("ovcopts-set", "", "")
	rootCmd.Flags().String("panscan", "", "Enables pan-and-scan functionality")
	rootCmd.Flags().String("pause", "", "Start the player in paused state")
	rootCmd.Flags().String("pipewire-buffer", "", "Set the audio buffer size in milliseconds")
	rootCmd.Flags().String("pipewire-remote", "", "Specify the PipeWire remote daemon name to connect to via local UNIX sockets")
	rootCmd.Flags().String("pipewire-volume-mode", "", "Specify if the ao-volume property should apply to the channel volumes or the global volume")
	rootCmd.Flags().String("pitch", "", "Raise or lower the audio's pitch by the factor given as parameter")
	rootCmd.Flags().String("play-direction", "", "Control the playback direction")
	rootCmd.Flags().String("player-operation-mode", "", "For enabling \"pseudo GUI mode\", which means that the defaults for some options are changed")
	rootCmd.Flags().String("playlist", "", "specify playlist file")
	rootCmd.Flags().String("playlist-exts", "", "Playlist file extensions to try to match when using --autocreate-playlist or --directory-filter-types")
	rootCmd.Flags().String("playlist-exts-add", "", "")
	rootCmd.Flags().String("playlist-exts-append", "", "")
	rootCmd.Flags().String("playlist-exts-clr", "", "")
	rootCmd.Flags().String("playlist-exts-del", "", "")
	rootCmd.Flags().String("playlist-exts-pre", "", "")
	rootCmd.Flags().String("playlist-exts-remove", "", "")
	rootCmd.Flags().String("playlist-exts-set", "", "")
	rootCmd.Flags().String("playlist-exts-toggle", "", "")
	rootCmd.Flags().String("playlist-start", "", "Set which file on the internal playlist to start playback with")
	rootCmd.Flags().String("prefetch-playlist", "", "Prefetch next playlist entry while playback of the current entry is ending")
	rootCmd.Flags().String("profile", "", "Use the given profile(s)")
	rootCmd.Flags().String("profile-add", "", "")
	rootCmd.Flags().String("profile-append", "", "")
	rootCmd.Flags().String("profile-clr", "", "")
	rootCmd.Flags().String("profile-del", "", "")
	rootCmd.Flags().String("profile-pre", "", "")
	rootCmd.Flags().String("profile-remove", "", "")
	rootCmd.Flags().String("profile-set", "", "")
	rootCmd.Flags().String("profile-toggle", "", "")
	rootCmd.Flags().String("pulse-allow-suspended", "", "Allow mpv to use PulseAudio even if the sink is suspended")
	rootCmd.Flags().String("pulse-buffer", "", "Set the audio buffer size in milliseconds")
	rootCmd.Flags().String("pulse-host", "", "Specify the host to use")
	rootCmd.Flags().String("pulse-latency-hacks", "", "Enable hacks to workaround PulseAudio timing bugs")
	rootCmd.Flags().String("quiet", "", "Make console output less verbose")
	rootCmd.Flags().String("rar-list-all-volumes", "", "When opening multi-volume rar files, open all volumes to create a full list of contained files")
	rootCmd.Flags().String("really-quiet", "", "Display even less output and status messages than with --quiet")
	rootCmd.Flags().String("rebase-start-time", "", "Whether to move the file start time to 00:00:00")
	rootCmd.Flags().String("referrer", "", "Specify a referrer path or URL for HTTP requests")
	rootCmd.Flags().String("replaygain", "", "Adjust volume gain according to replaygain values stored in the file metadata")
	rootCmd.Flags().String("replaygain-clip", "", "Allow the volume gain to clip")
	rootCmd.Flags().String("replaygain-fallback", "", "Gain in dB to apply if the file has no replay gain tags")
	rootCmd.Flags().String("replaygain-preamp", "", "Pre-amplification gain in dB to apply to the selected replaygain gain")
	rootCmd.Flags().String("reset-on-next-file", "", "Normally, mpv will try to keep all settings when playing the next file on the playlist, even if they were changed by the user during playback")
	rootCmd.Flags().String("reset-on-next-file-add", "", "")
	rootCmd.Flags().String("reset-on-next-file-append", "", "")
	rootCmd.Flags().String("reset-on-next-file-clr", "", "")
	rootCmd.Flags().String("reset-on-next-file-del", "", "")
	rootCmd.Flags().String("reset-on-next-file-pre", "", "")
	rootCmd.Flags().String("reset-on-next-file-remove", "", "")
	rootCmd.Flags().String("reset-on-next-file-set", "", "")
	rootCmd.Flags().String("reset-on-next-file-toggle", "", "")
	rootCmd.Flags().String("resume-playback", "", "Restore playback position from the watch_later configuration subdirectory usually ~/.config/mpv/watch_later/")
	rootCmd.Flags().String("resume-playback-check-mtime", "", "Only restore the playback position from the watch_later configuration subdirectory if the file's modification time is the same as at the time of saving")
	rootCmd.Flags().String("rtsp-transport", "", "Select RTSP transport method")
	rootCmd.Flags().String("saturation", "", "Adjust the saturation of the video signal")
	rootCmd.Flags().String("save-position-on-quit", "", "Always save the current playback position on quit, and also when the loadfile command is used to replace the current playlist")
	rootCmd.Flags().String("save-watch-history", "", "Whether to save which files are played")
	rootCmd.Flags().String("scale", "", "The filter function to use when upscaling video")
	rootCmd.Flags().String("scale-antiring", "", "Set the antiringing strength")
	rootCmd.Flags().String("scale-blur", "", "Kernel scaling factor")
	rootCmd.Flags().String("scale-clamp", "", "Specifies a weight bias to multiply into negative coefficients")
	rootCmd.Flags().String("scale-param1", "", "Set filter parameters")
	rootCmd.Flags().String("scale-param2", "", "Set filter parameters")
	rootCmd.Flags().String("scale-radius", "", "Set radius for tunable filters, must be a float number between 0.5 and 16.0")
	rootCmd.Flags().String("scale-taper", "", "Kernel/window taper factor")
	rootCmd.Flags().String("scale-window", "", "Choose a custom windowing function for the kernel")
	rootCmd.Flags().String("scale-wparam", "", "Configure the parameter for the window function given by --scale-window etc")
	rootCmd.Flags().String("scale-wtaper", "", "Kernel/window taper factor")
	rootCmd.Flags().String("scaler-resizes-only", "", "Disable the scaler if the video image is not resized")
	rootCmd.Flags().String("screen", "", "In multi-monitor configurations (i.e. a single desktop that spans across multiple displays), this option tells mpv which screen to display the video on")
	rootCmd.Flags().String("screen-name", "", "In multi-monitor configurations, this option tells mpv which screen to display the video on based on the screen name from the video backend")
	rootCmd.Flags().String("screenshot-avif-encoder", "", "Specify the AV1 encoder to be used by libavcodec for encoding avif screenshots")
	rootCmd.Flags().String("screenshot-avif-opts", "", "Specifies libavcodec options for selected encoder")
	rootCmd.Flags().String("screenshot-avif-opts-add", "", "")
	rootCmd.Flags().String("screenshot-avif-opts-append", "", "")
	rootCmd.Flags().String("screenshot-avif-opts-clr", "", "")
	rootCmd.Flags().String("screenshot-avif-opts-del", "", "")
	rootCmd.Flags().String("screenshot-avif-opts-remove", "", "")
	rootCmd.Flags().String("screenshot-avif-opts-set", "", "")
	rootCmd.Flags().String("screenshot-avif-pixfmt", "", "Specify the pixel format for the libavcodec encoder")
	rootCmd.Flags().String("screenshot-dir", "", "Store screenshots in this directory")
	rootCmd.Flags().String("screenshot-directory", "", "alias for --screenshot-dir")
	rootCmd.Flags().String("screenshot-format", "", "Set the image file type used for saving screenshots")
	rootCmd.Flags().String("screenshot-high-bit-depth", "", "If possible, write screenshots with a bit depth similar to the source video")
	rootCmd.Flags().String("screenshot-jpeg-quality", "", "Set the JPEG quality level")
	rootCmd.Flags().String("screenshot-jpeg-source-chroma", "", "Write JPEG files with the same chroma subsampling as the video")
	rootCmd.Flags().String("screenshot-jxl-distance", "", "Set the JPEG XL Butteraugli distance")
	rootCmd.Flags().String("screenshot-jxl-effort", "", "Set the JPEG XL compression effort")
	rootCmd.Flags().String("screenshot-png-compression", "", "Set the PNG compression level")
	rootCmd.Flags().String("screenshot-png-filter", "", "Set the filter applied prior to PNG compression")
	rootCmd.Flags().String("screenshot-sw", "", "Whether to use software rendering for screenshots")
	rootCmd.Flags().String("screenshot-tag-colorspace", "", "Tag screenshots with the appropriate colorspace")
	rootCmd.Flags().String("screenshot-template", "", "Specify the filename template used to save screenshots")
	rootCmd.Flags().String("screenshot-webp-compression", "", "Set the WebP compression level")
	rootCmd.Flags().String("screenshot-webp-lossless", "", "Write lossless WebP files")
	rootCmd.Flags().String("screenshot-webp-quality", "", "Set the WebP quality level")
	rootCmd.Flags().String("script", "", "Load a Lua script")
	rootCmd.Flags().StringSlice("script-opt", nil, "Set options for scripts")
	rootCmd.Flags().String("script-opts", "", "Set options for scripts")
	rootCmd.Flags().String("script-opts-add", "", "")
	rootCmd.Flags().String("script-opts-append", "", "")
	rootCmd.Flags().String("script-opts-clr", "", "")
	rootCmd.Flags().String("script-opts-del", "", "")
	rootCmd.Flags().String("script-opts-remove", "", "")
	rootCmd.Flags().String("script-opts-set", "", "")
	rootCmd.Flags().String("scripts", "", "Load a Lua script")
	rootCmd.Flags().String("scripts-add", "", "")
	rootCmd.Flags().String("scripts-append", "", "")
	rootCmd.Flags().String("scripts-clr", "", "")
	rootCmd.Flags().String("scripts-del", "", "")
	rootCmd.Flags().String("scripts-pre", "", "")
	rootCmd.Flags().String("scripts-remove", "", "")
	rootCmd.Flags().String("scripts-set", "", "")
	rootCmd.Flags().String("scripts-toggle", "", "")
	rootCmd.Flags().String("secondary-sid", "", "Select a secondary subtitle stream")
	rootCmd.Flags().String("secondary-sub-ass-override", "", "Control whether user secondary substyle overrides should be applied")
	rootCmd.Flags().String("secondary-sub-delay", "", "Delays secondary subtitles by <sec> seconds")
	rootCmd.Flags().String("secondary-sub-pos", "", "Specify the position of secondary subtitles on the screen")
	rootCmd.Flags().String("secondary-sub-visibility", "", "Can be used to disable display of secondary subtitles, but still select and decode them")
	rootCmd.Flags().String("sharpen", "", "If set to a value other than 0, enable an unsharp masking filter")
	rootCmd.Flags().String("show-in-taskbar", "", "(Windows and X11 only) Show mpv in the taskbar")
	rootCmd.Flags().String("show-profile", "", "Show the description and content of a profile")
	rootCmd.Flags().String("shuffle", "", "Play files in random order")
	rootCmd.Flags().String("sid", "", "Display the subtitle stream specified by <ID>")
	rootCmd.Flags().String("sigmoid-center", "", "The center of the sigmoid curve used for --sigmoid-upscaling")
	rootCmd.Flags().String("sigmoid-slope", "", "The slope of the sigmoid curve used for --sigmoid-upscaling")
	rootCmd.Flags().String("sigmoid-upscaling", "", "When upscaling, use a sigmoidal color transform to avoid emphasizing ringing artifacts")
	rootCmd.Flags().String("slang", "", "Analogous to --alang, for subtitle tracks")
	rootCmd.Flags().String("slang-add", "", "")
	rootCmd.Flags().String("slang-append", "", "")
	rootCmd.Flags().String("slang-clr", "", "")
	rootCmd.Flags().String("slang-del", "", "")
	rootCmd.Flags().String("slang-pre", "", "")
	rootCmd.Flags().String("slang-remove", "", "")
	rootCmd.Flags().String("slang-set", "", "")
	rootCmd.Flags().String("slang-toggle", "", "")
	rootCmd.Flags().String("snap-window", "", "(Windows only) Snap the player window to screen edges")
	rootCmd.Flags().String("speed", "", "Slow down or speed up playback by the factor given as parameter")
	rootCmd.Flags().String("spirv-compiler", "", "Controls which compiler is used to translate GLSL to SPIR-V")
	rootCmd.Flags().String("sstep", "", "Skip <sec> seconds after every frame")
	rootCmd.Flags().String("start", "", "seek to given (percent, seconds, or hh:mm:ss) position")
	rootCmd.Flags().String("stop-playback-on-init-failure", "", "Stop playback if either audio or video fails to initialize")
	rootCmd.Flags().String("stop-screensaver", "", "Turns off the screensaver (or screen blanker and similar mechanisms) at startup and turns it on again on exit")
	rootCmd.Flags().String("stream-buffer-size", "", "Size of the low level stream byte buffer")
	rootCmd.Flags().String("stream-dump", "", "Instead of playing a file, read its byte stream and write it to the given destination file")
	rootCmd.Flags().String("stream-lavf-o", "", "Set AVOptions on streams opened with libavformat")
	rootCmd.Flags().String("stream-lavf-o-add", "", "")
	rootCmd.Flags().String("stream-lavf-o-append", "", "")
	rootCmd.Flags().String("stream-lavf-o-clr", "", "")
	rootCmd.Flags().String("stream-lavf-o-del", "", "")
	rootCmd.Flags().String("stream-lavf-o-remove", "", "")
	rootCmd.Flags().String("stream-lavf-o-set", "", "")
	rootCmd.Flags().String("stream-record", "", "Write received/read data from the demuxer to the given output file")
	rootCmd.Flags().String("stretch-dvd-subs", "", "Stretch DVD subtitles when playing anamorphic videos for better looking fonts on badly mastered DVDs")
	rootCmd.Flags().String("stretch-image-subs-to-screen", "", "Stretch DVD and other image subtitles to the screen, ignoring the video margins")
	rootCmd.Flags().String("sub", "", "alias for --sid")
	rootCmd.Flags().String("sub-align-x", "", "Control to which corner of the screen text subtitles should be aligned to")
	rootCmd.Flags().String("sub-align-y", "", "Vertical position")
	rootCmd.Flags().String("sub-ass", "", "Render ASS subtitles natively")
	rootCmd.Flags().String("sub-ass-force-margins", "", "Enables placing toptitles and subtitles in black borders when they are available, if the subtitles are in the ASS format.")
	rootCmd.Flags().String("sub-ass-justify", "", "Applies justification as defined by --sub-justify on ASS subtitles if --sub-ass-override is not set to no")
	rootCmd.Flags().String("sub-ass-override", "", "Control whether user style overrides should be applied")
	rootCmd.Flags().String("sub-ass-prune-delay", "", "Set the delay for automatic pruning of events from memory in libass")
	rootCmd.Flags().String("sub-ass-scale-with-window", "", "Like --sub-scale-with-window, but affects subtitles in ASS format only")
	rootCmd.Flags().String("sub-ass-style-overrides", "", "Override some style or script info parameters")
	rootCmd.Flags().String("sub-ass-style-overrides-add", "", "")
	rootCmd.Flags().String("sub-ass-style-overrides-append", "", "")
	rootCmd.Flags().String("sub-ass-style-overrides-clr", "", "")
	rootCmd.Flags().String("sub-ass-style-overrides-del", "", "")
	rootCmd.Flags().String("sub-ass-style-overrides-pre", "", "")
	rootCmd.Flags().String("sub-ass-style-overrides-remove", "", "")
	rootCmd.Flags().String("sub-ass-style-overrides-set", "", "")
	rootCmd.Flags().String("sub-ass-style-overrides-toggle", "", "")
	rootCmd.Flags().String("sub-ass-styles", "", "Load all SSA/ASS styles found in the specified file and use them for rendering text subtitles")
	rootCmd.Flags().String("sub-ass-use-video-data", "", "Controls which information about the video stream is passed to libass")
	rootCmd.Flags().String("sub-ass-video-aspect-override", "", "Allows passing any arbitrary aspect ratio to libass instead of the video’s actual aspect ratio")
	rootCmd.Flags().String("sub-ass-vsfilter-color-compat", "", "Mangle colors like (xy-)vsfilter do")
	rootCmd.Flags().String("sub-auto", "", "Load additional subtitle files matching the video filename")
	rootCmd.Flags().String("sub-auto-exts", "", "Subtitle extensions to try and match when using --sub-auto")
	rootCmd.Flags().String("sub-auto-exts-add", "", "")
	rootCmd.Flags().String("sub-auto-exts-append", "", "")
	rootCmd.Flags().String("sub-auto-exts-clr", "", "")
	rootCmd.Flags().String("sub-auto-exts-del", "", "")
	rootCmd.Flags().String("sub-auto-exts-pre", "", "")
	rootCmd.Flags().String("sub-auto-exts-remove", "", "")
	rootCmd.Flags().String("sub-auto-exts-set", "", "")
	rootCmd.Flags().String("sub-auto-exts-toggle", "", "")
	rootCmd.Flags().String("sub-back-color", "", "Color used for sub text background")
	rootCmd.Flags().String("sub-blur", "", "Gaussian blur factor applied to the sub font border")
	rootCmd.Flags().String("sub-bold", "", "Format text on bold")
	rootCmd.Flags().String("sub-border-color", "", "alias for --sub-outline-color")
	rootCmd.Flags().String("sub-border-size", "", "alias for --sub-outline-size")
	rootCmd.Flags().String("sub-border-style", "", "The style of the border")
	rootCmd.Flags().String("sub-clear-on-seek", "", "Can be used to play broken mkv files with duplicate ReadOrder fields")
	rootCmd.Flags().String("sub-codepage", "", "You can use this option to specify the subtitle codepage")
	rootCmd.Flags().String("sub-color", "", "Specify the color used for unstyled text subtitles")
	rootCmd.Flags().String("sub-create-cc-track", "", "For every video stream, create a closed captions track")
	rootCmd.Flags().String("sub-delay", "", "Delays primary subtitles by <sec> seconds")
	rootCmd.Flags().String("sub-demuxer", "", "Force subtitle demuxer type for --sub-file")
	rootCmd.Flags().String("sub-file", "", "specify subtitle file to use")
	rootCmd.Flags().String("sub-file-paths", "", "Specify extra directories to search for subtitles matching the video")
	rootCmd.Flags().String("sub-file-paths-add", "", "")
	rootCmd.Flags().String("sub-file-paths-append", "", "")
	rootCmd.Flags().String("sub-file-paths-clr", "", "")
	rootCmd.Flags().String("sub-file-paths-del", "", "")
	rootCmd.Flags().String("sub-file-paths-pre", "", "")
	rootCmd.Flags().String("sub-file-paths-remove", "", "")
	rootCmd.Flags().String("sub-file-paths-set", "", "")
	rootCmd.Flags().String("sub-file-paths-toggle", "", "")
	rootCmd.Flags().String("sub-files", "", "Add a subtitle file to the list of external subtitles")
	rootCmd.Flags().String("sub-files-add", "", "")
	rootCmd.Flags().String("sub-files-append", "", "")
	rootCmd.Flags().String("sub-files-clr", "", "")
	rootCmd.Flags().String("sub-files-del", "", "")
	rootCmd.Flags().String("sub-files-pre", "", "")
	rootCmd.Flags().String("sub-files-remove", "", "")
	rootCmd.Flags().String("sub-files-set", "", "")
	rootCmd.Flags().String("sub-files-toggle", "", "")
	rootCmd.Flags().String("sub-filter-jsre", "", "Same as --sub-filter-regex but with JavaScript regular expressions")
	rootCmd.Flags().String("sub-filter-jsre-add", "", "")
	rootCmd.Flags().String("sub-filter-jsre-append", "", "")
	rootCmd.Flags().String("sub-filter-jsre-clr", "", "")
	rootCmd.Flags().String("sub-filter-jsre-del", "", "")
	rootCmd.Flags().String("sub-filter-jsre-pre", "", "")
	rootCmd.Flags().String("sub-filter-jsre-remove", "", "")
	rootCmd.Flags().String("sub-filter-jsre-set", "", "")
	rootCmd.Flags().String("sub-filter-jsre-toggle", "", "")
	rootCmd.Flags().String("sub-filter-regex", "", "Set a list of regular expressions to match on text subtitles, and remove any lines that match")
	rootCmd.Flags().String("sub-filter-regex-add", "", "")
	rootCmd.Flags().String("sub-filter-regex-append", "", "")
	rootCmd.Flags().String("sub-filter-regex-clr", "", "")
	rootCmd.Flags().String("sub-filter-regex-del", "", "")
	rootCmd.Flags().String("sub-filter-regex-enable", "", "Whether to enable regex filtering")
	rootCmd.Flags().String("sub-filter-regex-plain", "", "Whether to first convert the ASS \"Text\" field to plain-text")
	rootCmd.Flags().String("sub-filter-regex-pre", "", "")
	rootCmd.Flags().String("sub-filter-regex-remove", "", "")
	rootCmd.Flags().String("sub-filter-regex-set", "", "")
	rootCmd.Flags().String("sub-filter-regex-toggle", "", "")
	rootCmd.Flags().String("sub-filter-regex-warn", "", "Log dropped lines with warning log level, instead of verbose")
	rootCmd.Flags().String("sub-filter-sdh", "", "Applies filter removing subtitle additions for the deaf or hard-of-hearing")
	rootCmd.Flags().String("sub-filter-sdh-enclosures", "", "Specify a string of characters that --sub-filter-sdh will use to potentially remove text")
	rootCmd.Flags().String("sub-filter-sdh-harder", "", "Do harder SDH filtering")
	rootCmd.Flags().String("sub-fix-timing", "", "Adjust subtitle timing is to remove minor gaps or overlaps between subtitles")
	rootCmd.Flags().String("sub-font", "", "Specify font to use for subtitles that do not themselves specify a particular font")
	rootCmd.Flags().String("sub-font-provider", "", "Which libass font provider backend to use")
	rootCmd.Flags().String("sub-font-size", "", "Specify the sub font size")
	rootCmd.Flags().String("sub-fonts-dir", "", "Font files in this directory are used by mpv/libass for subtitles")
	rootCmd.Flags().String("sub-forced-events-only", "", "Enabling this displays only forced events within subtitle streams")
	rootCmd.Flags().String("sub-fps", "", "Specify the framerate of the subtitle file")
	rootCmd.Flags().String("sub-gauss", "", "Apply Gaussian blur to image subtitles")
	rootCmd.Flags().String("sub-gray", "", "Convert image subtitles to grayscale")
	rootCmd.Flags().String("sub-hinting", "", "Set font hinting type")
	rootCmd.Flags().String("sub-italic", "", "Format text on italic")
	rootCmd.Flags().String("sub-justify", "", "Control how multi line subs are justified irrespective of where they are aligned")
	rootCmd.Flags().String("sub-lavc-o", "", "Pass AVOptions to libavcodec decoder")
	rootCmd.Flags().String("sub-lavc-o-add", "", "")
	rootCmd.Flags().String("sub-lavc-o-append", "", "")
	rootCmd.Flags().String("sub-lavc-o-clr", "", "")
	rootCmd.Flags().String("sub-lavc-o-del", "", "")
	rootCmd.Flags().String("sub-lavc-o-remove", "", "")
	rootCmd.Flags().String("sub-lavc-o-set", "", "")
	rootCmd.Flags().String("sub-line-spacing", "", "Set line spacing value for SSA/ASS renderer")
	rootCmd.Flags().String("sub-margin-x", "", "Left and right screen margin for the subs in scaled pixels")
	rootCmd.Flags().String("sub-margin-y", "", "Top and bottom screen margin for the subs in scaled pixels")
	rootCmd.Flags().String("sub-outline-color", "", "Color used for the sub font outline")
	rootCmd.Flags().String("sub-outline-size", "", "Size of the sub font outline in scaled pixels")
	rootCmd.Flags().String("sub-past-video-end", "", "After the last frame of video, if this option is enabled, subtitles will continue to update based on audio timestamps")
	rootCmd.Flags().String("sub-pos", "", "Specify the position of subtitles on the screen")
	rootCmd.Flags().String("sub-scale", "", "Factor for the text subtitle font size")
	rootCmd.Flags().String("sub-scale-by-window", "", "Whether to scale subtitles with the window size")
	rootCmd.Flags().String("sub-scale-signs", "", "When set to yes, also apply --sub-scale to typesetting (or \"signs\")")
	rootCmd.Flags().String("sub-scale-with-window", "", "Make the subtitle font size relative to the window")
	rootCmd.Flags().String("sub-shadow-color", "", "alias for --sub-back-color")
	rootCmd.Flags().String("sub-shadow-offset", "", "Displacement of the sub text shadow in scaled pixels")
	rootCmd.Flags().String("sub-shaper", "", "Set the text layout engine used by libass")
	rootCmd.Flags().String("sub-spacing", "", "Horizontal sub font spacing in scaled pixels")
	rootCmd.Flags().String("sub-speed", "", "Multiply the subtitle event timestamps with the given value")
	rootCmd.Flags().String("sub-stretch-durations", "", "Stretch a subtitle duration so it ends when the next one starts")
	rootCmd.Flags().String("sub-use-margins", "", "Enables placing toptitles and subtitles in black borders when they are available, if the subtitles are in a plain text format")
	rootCmd.Flags().String("sub-visibility", "", "Can be used to disable display of subtitles, but still select and decode them")
	rootCmd.Flags().String("sub-vsfilter-bidi-compat", "", "Set implicit bidi detection to ltr instead of auto to match ASS' default")
	rootCmd.Flags().String("subs-fallback", "", "When autoselecting a subtitle track, if no tracks match your preferred languages, select a full track even if it doesn't match your preferred subtitle language")
	rootCmd.Flags().String("subs-fallback-forced", "", "When autoselecting a subtitle track, the default value of yes will prefer using a forced subtitle track if the subtitle language matches the audio language and matches your list of preferred languages")
	rootCmd.Flags().String("subs-match-os-language", "", "When autoselecting a subtitle track, select the track that matches the language of your OS if the audio stream is in a different language if suitable")
	rootCmd.Flags().String("subs-with-matching-audio", "", "When autoselecting a subtitle track, select it even if the selected audio stream matches you preferred subtitle language")
	rootCmd.Flags().String("swapchain-depth", "", "Allow up to N in-flight frames")
	rootCmd.Flags().String("sws-allow-zimg", "", "Allow using zimg")
	rootCmd.Flags().String("sws-bitexact", "", "Unknown functionality")
	rootCmd.Flags().String("sws-cgb", "", "Software scaler Gaussian blur filter")
	rootCmd.Flags().String("sws-chs", "", "Software scaler chroma horizontal shifting")
	rootCmd.Flags().String("sws-cs", "", "Software scaler sharpen filter")
	rootCmd.Flags().String("sws-cvs", "", "Software scaler chroma vertical shifting")
	rootCmd.Flags().String("sws-fast", "", "Allow optimizations that help with performance, but reduce quality")
	rootCmd.Flags().String("sws-lgb", "", "Software scaler Gaussian blur filter")
	rootCmd.Flags().String("sws-ls", "", "Software scaler sharpen filter")
	rootCmd.Flags().String("sws-scaler", "", "Specify the software scaler algorithm to be used with --vf=scale")
	rootCmd.Flags().String("target-colorspace-hint", "", "Automatically configure the output colorspace of the display to pass through the input values of the stream (e.g. for HDR passthrough), if possible")
	rootCmd.Flags().String("target-contrast", "", "Specifies the measured contrast of the output display")
	rootCmd.Flags().String("target-gamut", "", "Constrains the gamut of the display")
	rootCmd.Flags().String("target-lut", "", "Specifies a custom LUT file (in Adobe .cube format) to apply to the colors before display on-screen")
	rootCmd.Flags().String("target-peak", "", "Specifies the measured peak brightness of the output display, in cd/m^2 (AKA nits)")
	rootCmd.Flags().String("target-prim", "", "Specifies the primaries of the display")
	rootCmd.Flags().String("target-trc", "", "Specifies the transfer characteristics (gamma) of the display")
	rootCmd.Flags().String("taskbar-progress", "", "(Windows only) Enable/disable playback progress rendering in taskbar")
	rootCmd.Flags().String("teletext-page", "", "Select a teletext page number to decode")
	rootCmd.Flags().String("temporal-dither", "", "Enable temporal dithering")
	rootCmd.Flags().String("temporal-dither-period", "", "Determines how often the dithering pattern is updated when --temporal-dither is in use")
	rootCmd.Flags().String("term-osd", "", "Control whether OSD messages are shown on the console when no video output is available")
	rootCmd.Flags().String("term-osd-bar", "", "Enable printing a progress bar under the status line on the terminal")
	rootCmd.Flags().String("term-osd-bar-chars", "", "Customize the --term-osd-bar feature")
	rootCmd.Flags().String("term-playing-msg", "", "Print out a string after starting playback")
	rootCmd.Flags().String("term-status-msg", "", "Print out a custom string during playback instead of the standard status line")
	rootCmd.Flags().String("term-title", "", "Set the terminal title")
	rootCmd.Flags().String("terminal", "", "--terminal=no disables any use of the terminal and stdin/stdout/stderr")
	rootCmd.Flags().String("title", "", "Set the window title")
	rootCmd.Flags().String("title-bar", "", "(Windows and X11 only) Play video with the window title bar")
	rootCmd.Flags().String("tls-ca-file", "", "Certificate authority database file for use with TLS")
	rootCmd.Flags().String("tls-cert-file", "", "A file containing a certificate to use in the handshake with the peer")
	rootCmd.Flags().String("tls-key-file", "", "A file containing the private key for the certificate")
	rootCmd.Flags().String("tls-verify", "", "Verify peer certificates when using TLS (e.g. with https://...). (Silently fails with older FFmpeg versions.)")
	rootCmd.Flags().String("tone-mapping", "", "Specifies the algorithm used for tone-mapping images onto the target display")
	rootCmd.Flags().String("tone-mapping-max-boost", "", "Upper limit for how much the tone mapping algorithm is allowed to boost the average brightness by over-exposing the image")
	rootCmd.Flags().String("tone-mapping-param", "", "Set tone mapping parameters")
	rootCmd.Flags().String("tone-mapping-visualize", "", "Display a (PQ-PQ) graph of the active tone-mapping LUT")
	rootCmd.Flags().String("track-auto-selection", "", "Enable the default track auto-selection")
	rootCmd.Flags().String("tscale", "", "The filter used for interpolating the temporal axis")
	rootCmd.Flags().String("tscale-antiring", "", "Set the antiringing strength")
	rootCmd.Flags().String("tscale-blur", "", "Kernel scaling factor")
	rootCmd.Flags().String("tscale-clamp", "", "Specifies a weight bias to multiply into negative coefficients")
	rootCmd.Flags().String("tscale-param1", "", "Set filter parameters")
	rootCmd.Flags().String("tscale-param2", "", "Set filter parameters")
	rootCmd.Flags().String("tscale-radius", "", "Set radius for tunable filters, must be a float number between 0.5 and 16.0")
	rootCmd.Flags().String("tscale-taper", "", "Kernel/window taper factor")
	rootCmd.Flags().String("tscale-window", "", "Choose a custom windowing function for the kernel")
	rootCmd.Flags().String("tscale-wparam", "", "Configure the parameter for the window function given by --scale-window etc")
	rootCmd.Flags().String("tscale-wtaper", "", "Kernel/window taper factor")
	rootCmd.Flags().String("untimed", "", "Do not sleep when outputting video frames")
	rootCmd.Flags().String("use-embedded-icc-profile", "", "Load the embedded ICC profile contained in media files such as PNG images")
	rootCmd.Flags().String("use-filedir-conf", "", "Look for a file-specific configuration file in the same directory as the file that is being played")
	rootCmd.Flags().String("user-agent", "", "Use <string> as user agent for HTTP streaming")
	rootCmd.Flags().CountP("v", "v", "Increment verbosity level, one level for each -v found on the command line")
	rootCmd.Flags().String("vaapi-device", "", "Choose the DRM device for vaapi-copy")
	rootCmd.Flags().String("vd", "", "Specify a priority list of video decoders to be used")
	rootCmd.Flags().String("vd-apply-cropping", "", "Certain video codecs support cropping, meaning that only a sub-rectangle of the decoded frame is intended for display")
	rootCmd.Flags().String("vd-lavc-assume-old-x264", "", "Assume the video was encoded by an old, buggy x264 version")
	rootCmd.Flags().String("vd-lavc-bitexact", "", "Only use bit-exact algorithms in all decoding steps")
	rootCmd.Flags().String("vd-lavc-check-hw-profile", "", "Check hardware decoder profile")
	rootCmd.Flags().String("vd-lavc-dr", "", "Enable direct rendering")
	rootCmd.Flags().String("vd-lavc-fast", "", "Enable optimizations which do not comply with the format specification and potentially cause problems, like simpler dequantization, simpler motion compensation, assuming use of the default quantization matrix, assuming YUV 4:2:0 and skipping a few checks to detect damaged bitstreams")
	rootCmd.Flags().String("vd-lavc-film-grain", "", "Enables film grain application on the GPU")
	rootCmd.Flags().String("vd-lavc-framedrop", "", "Set framedropping mode used with --framedrop")
	rootCmd.Flags().String("vd-lavc-o", "", "Pass AVOptions to libavcodec decoder")
	rootCmd.Flags().String("vd-lavc-o-add", "", "")
	rootCmd.Flags().String("vd-lavc-o-append", "", "")
	rootCmd.Flags().String("vd-lavc-o-clr", "", "")
	rootCmd.Flags().String("vd-lavc-o-del", "", "")
	rootCmd.Flags().String("vd-lavc-o-remove", "", "")
	rootCmd.Flags().String("vd-lavc-o-set", "", "")
	rootCmd.Flags().String("vd-lavc-show-all", "", "Show even broken/corrupt frames")
	rootCmd.Flags().String("vd-lavc-skipframe", "", "Skips decoding of frames completely")
	rootCmd.Flags().String("vd-lavc-skipidct", "", "Skips the IDCT step")
	rootCmd.Flags().String("vd-lavc-skiploopfilter", "", "Skips the loop filter (AKA deblocking) during decoding")
	rootCmd.Flags().String("vd-lavc-threads", "", "Number of threads to use for decoding")
	rootCmd.Flags().String("vd-queue-enable", "", "Enable running the video decoder on a separate thread")
	rootCmd.Flags().String("vd-queue-max-bytes", "", "Maximum approximate allowed size of the queue")
	rootCmd.Flags().String("vd-queue-max-samples", "", "Maximum number of frames (video) or samples (audio) of the queue")
	rootCmd.Flags().String("vd-queue-max-secs", "", "Maximum number of seconds of media in the queue")
	rootCmd.Flags().StringP("version", "V", "", "Print version string and exit")
	rootCmd.Flags().String("vf", "", "Setup a chain of video filters")
	rootCmd.Flags().String("vf-add", "", "")
	rootCmd.Flags().String("vf-append", "", "")
	rootCmd.Flags().String("vf-clr", "", "")
	rootCmd.Flags().String("vf-help", "", "")
	rootCmd.Flags().String("vf-pre", "", "")
	rootCmd.Flags().String("vf-remove", "", "")
	rootCmd.Flags().String("vf-set", "", "")
	rootCmd.Flags().String("vf-toggle", "", "")
	rootCmd.Flags().String("vid", "", "Select video channel")
	rootCmd.Flags().String("video", "", "alias for --vid")
	rootCmd.Flags().String("video-align-x", "", "When the video is bigger than the window, these move the displayed rectangle to show different parts of the video. --video-align-y=-1 would display the top of the video, 0 would display the center (default), and 1 would display the bottom")
	rootCmd.Flags().String("video-align-y", "", "When the video is bigger than the window, these move the displayed rectangle to show different parts of the video. --video-align-y=-1 would display the top of the video, 0 would display the center (default), and 1 would display the bottom")
	rootCmd.Flags().String("video-aspect-method", "", "This sets the default video aspect determination method")
	rootCmd.Flags().String("video-aspect-override", "", "Override video aspect ratio, in case aspect information is incorrect or missing in the file being played")
	rootCmd.Flags().String("video-backward-batch", "", "Number of keyframe ranges to decode at once when backward decoding")
	rootCmd.Flags().String("video-backward-overlap", "", "Number of overlapping keyframe ranges to use for backward decoding")
	rootCmd.Flags().String("video-crop", "", "Crop the video by starting at the x, y offset for w, h pixels")
	rootCmd.Flags().String("video-exts", "", "Video file extensions to try to match when using --autocreate-playlist or --directory-filter-types")
	rootCmd.Flags().String("video-exts-add", "", "")
	rootCmd.Flags().String("video-exts-append", "", "")
	rootCmd.Flags().String("video-exts-clr", "", "")
	rootCmd.Flags().String("video-exts-del", "", "")
	rootCmd.Flags().String("video-exts-pre", "", "")
	rootCmd.Flags().String("video-exts-remove", "", "")
	rootCmd.Flags().String("video-exts-set", "", "")
	rootCmd.Flags().String("video-exts-toggle", "", "")
	rootCmd.Flags().String("video-latency-hacks", "", "Enable some things which tend to reduce video latency by 1 or 2 frames")
	rootCmd.Flags().String("video-margin-ratio-bottom", "", "Set extra video margin on bottom border")
	rootCmd.Flags().String("video-margin-ratio-left", "", "Set extra video margin on left border")
	rootCmd.Flags().String("video-margin-ratio-right", "", "Set extra video margin on right border")
	rootCmd.Flags().String("video-margin-ratio-top", "", "Set extra video margin on top border")
	rootCmd.Flags().String("video-osd", "", "Enabled OSD rendering on the video window")
	rootCmd.Flags().String("video-output-levels", "", "RGB color levels used with YUV to RGB conversion")
	rootCmd.Flags().String("video-pan-x", "", "Moves the displayed video rectangle by the given value in the X direction")
	rootCmd.Flags().String("video-pan-y", "", "Moves the displayed video rectangle by the given value in the Y direction")
	rootCmd.Flags().String("video-recenter", "", "Whether to behave as if --video-align-x and --video-align-y were 0 when the video becomes smaller than the window in the respective direction")
	rootCmd.Flags().String("video-reversal-buffer", "", "For backward decoding")
	rootCmd.Flags().String("video-rotate", "", "Rotate the video clockwise, in degrees")
	rootCmd.Flags().String("video-scale-x", "", "Multiply the video display size with the given value")
	rootCmd.Flags().String("video-scale-y", "", "Multiply the video display size with the given value")
	rootCmd.Flags().String("video-sync", "", "How the player synchronizes audio and video")
	rootCmd.Flags().String("video-sync-max-audio-change", "", "Maximum additional speed difference in percent that is applied to audio with --video-sync=display-...")
	rootCmd.Flags().String("video-sync-max-factor", "", "Maximum multiple for which to try to fit the video's FPS to the display's FPS")
	rootCmd.Flags().String("video-sync-max-video-change", "", "Maximum speed difference in percent that is applied to video with --video-sync=display-... ")
	rootCmd.Flags().String("video-timing-offset", "", "Control how long before video display target time the frame should be rendered")
	rootCmd.Flags().String("video-unscaled", "", "Disable scaling of the video")
	rootCmd.Flags().String("video-zoom", "", "Adjust the video display scale factor by the given value")
	rootCmd.Flags().String("vlang", "", "Analogous to --alang and --slang, for video tracks")
	rootCmd.Flags().String("vlang-add", "", "")
	rootCmd.Flags().String("vlang-append", "", "")
	rootCmd.Flags().String("vlang-clr", "", "")
	rootCmd.Flags().String("vlang-del", "", "")
	rootCmd.Flags().String("vlang-pre", "", "")
	rootCmd.Flags().String("vlang-remove", "", "")
	rootCmd.Flags().String("vlang-set", "", "")
	rootCmd.Flags().String("vlang-toggle", "", "")
	rootCmd.Flags().String("vo", "", "Specify the video output backend to be used")
	rootCmd.Flags().String("vo-add", "", "")
	rootCmd.Flags().String("vo-append", "", "")
	rootCmd.Flags().String("vo-clr", "", "")
	rootCmd.Flags().String("vo-help", "", "")
	rootCmd.Flags().String("vo-image-avif-encoder", "", "")
	rootCmd.Flags().String("vo-image-avif-opts", "", "")
	rootCmd.Flags().String("vo-image-avif-opts-add", "", "")
	rootCmd.Flags().String("vo-image-avif-opts-append", "", "")
	rootCmd.Flags().String("vo-image-avif-opts-clr", "", "")
	rootCmd.Flags().String("vo-image-avif-opts-del", "", "")
	rootCmd.Flags().String("vo-image-avif-opts-remove", "", "")
	rootCmd.Flags().String("vo-image-avif-opts-set", "", "")
	rootCmd.Flags().String("vo-image-avif-pixfmt", "", "")
	rootCmd.Flags().String("vo-image-format", "", "Select the image file format")
	rootCmd.Flags().String("vo-image-high-bit-depth", "", "")
	rootCmd.Flags().String("vo-image-jpeg-quality", "", "JPEG quality factor")
	rootCmd.Flags().String("vo-image-jpeg-source-chroma", "", "")
	rootCmd.Flags().String("vo-image-jxl-distance", "", "")
	rootCmd.Flags().String("vo-image-jxl-effort", "", "")
	rootCmd.Flags().String("vo-image-outdir", "", "Specify the directory to save the image files to")
	rootCmd.Flags().String("vo-image-png-compression", "", "PNG compression factor")
	rootCmd.Flags().String("vo-image-png-filter", "", "Filter applied prior to PNG compression")
	rootCmd.Flags().String("vo-image-tag-colorspace", "", "")
	rootCmd.Flags().String("vo-image-webp-compression", "", "WebP compression factor")
	rootCmd.Flags().String("vo-image-webp-lossless", "", "Enable writing lossless WebP files")
	rootCmd.Flags().String("vo-image-webp-quality", "", "WebP quality")
	rootCmd.Flags().String("vo-kitty-alt-screen", "", "Whether or not to use the alternate screen buffer and return the terminal to its previous state on exit")
	rootCmd.Flags().String("vo-kitty-cols", "", "Specify the terminal size in character cells, otherwise (0) read it from the terminal, or fall back to 80x25")
	rootCmd.Flags().String("vo-kitty-config-clear", "", "Whether or not to clear the terminal whenever the output is reconfigured")
	rootCmd.Flags().String("vo-kitty-height", "", "Specify the available size in pixels, otherwise (0) read it from the terminal, or fall back to 320x240")
	rootCmd.Flags().String("vo-kitty-left", "", "Specify the position in character cells where the image starts")
	rootCmd.Flags().String("vo-kitty-rows", "", "Specify the terminal size in character cells, otherwise (0) read it from the terminal, or fall back to 80x25")
	rootCmd.Flags().String("vo-kitty-top", "", "Specify the position in character cells where the image starts")
	rootCmd.Flags().String("vo-kitty-use-shm", "", "Use shared memory objects to transfer image data to the terminal")
	rootCmd.Flags().String("vo-kitty-width", "", "Specify the available size in pixels, otherwise (0) read it from the terminal, or fall back to 320x240")
	rootCmd.Flags().String("vo-null-fps", "", "Simulate display FPS")
	rootCmd.Flags().String("vo-pre", "", "")
	rootCmd.Flags().String("vo-remove", "", "")
	rootCmd.Flags().String("vo-set", "", "")
	rootCmd.Flags().String("vo-sixel-alt-screen", "", "Whether or not to use the alternate screen buffer and return the terminal to its previous state on exit")
	rootCmd.Flags().String("vo-sixel-buffered", "", "Buffers the full output sequence before writing it to the terminal")
	rootCmd.Flags().String("vo-sixel-cols", "", "Specify the terminal size in character cells, otherwise (0) read it from the terminal, or fall back to 80x25")
	rootCmd.Flags().String("vo-sixel-config-clear", "", "Whether or not to clear the terminal whenever the output is reconfigured")
	rootCmd.Flags().String("vo-sixel-dither", "", "Selects the dither algorithm which libsixel should apply")
	rootCmd.Flags().String("vo-sixel-fixedpalette", "", "Use libsixel's built-in static palette using the XTERM256 profile for dither")
	rootCmd.Flags().String("vo-sixel-height", "", "Specify the available size in pixels, otherwise (0) read it from the terminal, or fall back to 320x240")
	rootCmd.Flags().String("vo-sixel-left", "", "Specify the position in character cells where the image starts")
	rootCmd.Flags().String("vo-sixel-pad-x", "", "Specify the number of padding pixels which are included at the size which the terminal reports")
	rootCmd.Flags().String("vo-sixel-pad-y", "", "Specify the number of padding pixels which are included at the size which the terminal reports")
	rootCmd.Flags().String("vo-sixel-reqcolors", "", "Has no effect with fixed palette")
	rootCmd.Flags().String("vo-sixel-rows", "", "Specify the terminal size in character cells, otherwise (0) read it from the terminal, or fall back to 80x25")
	rootCmd.Flags().String("vo-sixel-threshold", "", "Has no effect with fixed palette")
	rootCmd.Flags().String("vo-sixel-top", "", "Specify the position in character cells where the image starts")
	rootCmd.Flags().String("vo-sixel-width", "", "Specify the available size in pixels, otherwise (0) read it from the terminal, or fall back to 320x240")
	rootCmd.Flags().String("vo-tct-256", "", "Use 256 colors - for terminals which don't support true color")
	rootCmd.Flags().String("vo-tct-algo", "", "Select how to write the pixels to the terminal")
	rootCmd.Flags().String("vo-tct-buffering", "", "Specifies the size of data batches buffered before being sent to the terminal")
	rootCmd.Flags().String("vo-tct-height", "", "Assume the terminal has the specified character height")
	rootCmd.Flags().String("vo-tct-width", "", "Assume the terminal has the specified character width")
	rootCmd.Flags().String("vo-toggle", "", "")
	rootCmd.Flags().String("vo-vaapi-scaled-osd", "", "If enabled, then the OSD is rendered at video resolution and scaled to display resolution")
	rootCmd.Flags().String("vo-vaapi-scaling", "", "")
	rootCmd.Flags().String("vo-vdpau-chroma-deint", "", "Makes temporal deinterlacers operate both on luma and chroma")
	rootCmd.Flags().String("vo-vdpau-colorkey", "", "Set the VDPAU presentation queue background color, which in practice is the colorkey used if VDPAU operates in overlay mode")
	rootCmd.Flags().String("vo-vdpau-composite-detect", "", "The player tries to detect whether a compositing window manager is active. If one is detected, the player disables timing adjustments as if the user had specified fps=-1")
	rootCmd.Flags().String("vo-vdpau-denoise", "", "Apply a noise reduction algorithm to the video")
	rootCmd.Flags().String("vo-vdpau-force-yuv", "", "Never accept RGBA input")
	rootCmd.Flags().String("vo-vdpau-fps", "", "Override autodetected display refresh rate value")
	rootCmd.Flags().String("vo-vdpau-hqscaling", "", "Use default VDPAU scaling")
	rootCmd.Flags().String("vo-vdpau-output-surfaces", "", "Allocate this many output surfaces to display video frames")
	rootCmd.Flags().String("vo-vdpau-pullup", "", "Try to apply inverse telecine, needs motion adaptive temporal deinterlacing")
	rootCmd.Flags().String("vo-vdpau-queuetime-fs", "", "Use VDPAU's presentation queue functionality to queue future video frame changes at most this many milliseconds in advance")
	rootCmd.Flags().String("vo-vdpau-queuetime-windowed", "", "Use VDPAU's presentation queue functionality to queue future video frame changes at most this many milliseconds in advance")
	rootCmd.Flags().String("vo-vdpau-sharpen", "", "For positive values, apply a sharpening algorithm to the video, for negative values a blurring algorithm")
	rootCmd.Flags().String("volume", "", "Set the startup volume")
	rootCmd.Flags().String("volume-gain", "", "Set the volume gain in dB")
	rootCmd.Flags().String("volume-gain-max", "", "Set the volume gain range in dB")
	rootCmd.Flags().String("volume-gain-min", "", "Set the volume gain range in dB")
	rootCmd.Flags().String("volume-max", "", "Set the maximum amplification level in percent")
	rootCmd.Flags().String("vulkan-async-compute", "", "Enables the use of async compute queues on supported vulkan devices")
	rootCmd.Flags().String("vulkan-async-transfer", "", "Enables the use of async transfer queues on supported vulkan devices")
	rootCmd.Flags().String("vulkan-device", "", "The name or UUID of the Vulkan device to use for rendering and presentation")
	rootCmd.Flags().String("vulkan-display-display", "", "The index of the display, on the selected Vulkan device, to present on when using the displayvk GPU context")
	rootCmd.Flags().String("vulkan-display-mode", "", "The index of the display mode, of the selected Vulkan display, to use when using the displayvk GPU contex")
	rootCmd.Flags().String("vulkan-display-plane", "", "The index of the plane, on the selected Vulkan device, to present on when using the displayvk GPU context")
	rootCmd.Flags().String("vulkan-queue-count", "", "Controls the number of VkQueues used for rendering")
	rootCmd.Flags().String("vulkan-swap-mode", "", "Controls the presentation mode of the vulkan swapchain")
	rootCmd.Flags().String("watch-history-path", "", "The path in which to store the watch history")
	rootCmd.Flags().String("watch-later-dir", "", "The directory in which to store the \"watch later\" temporary files")
	rootCmd.Flags().String("watch-later-directory", "", "alias for --watch-later-dir")
	rootCmd.Flags().String("watch-later-options", "", "The options that are saved in \"watch later\" files if they have been changed since when mpv started")
	rootCmd.Flags().String("watch-later-options-add", "", "")
	rootCmd.Flags().String("watch-later-options-append", "", "")
	rootCmd.Flags().String("watch-later-options-clr", "", "")
	rootCmd.Flags().String("watch-later-options-del", "", "")
	rootCmd.Flags().String("watch-later-options-pre", "", "")
	rootCmd.Flags().String("watch-later-options-remove", "", "")
	rootCmd.Flags().String("watch-later-options-set", "", "")
	rootCmd.Flags().String("watch-later-options-toggle", "", "")
	rootCmd.Flags().String("wayland-app-id", "", "Set the client app id for Wayland-based video output methods")
	rootCmd.Flags().String("wayland-configure-bounds", "", "Controls whether or not mpv opts into the configure bounds event if sent by the compositor")
	rootCmd.Flags().String("wayland-content-type", "", "If supported by the compositor, mpv will send a hint using the content-type protocol telling the compositor what type of content is being displayed")
	rootCmd.Flags().String("wayland-edge-pixels-pointer", "", "Defines the size of an edge border (default: 16) to initiate client side resize events in the wayland contexts with the mouse")
	rootCmd.Flags().String("wayland-edge-pixels-touch", "", "Defines the size of an edge border (default: 32) to initiate client side resizes events in the wayland contexts with touch events")
	rootCmd.Flags().String("wayland-internal-vsync", "", "Controls whether to use mpv's internal vsync for Wayland-base video outputs")
	rootCmd.Flags().String("wayland-present", "", "Enable the use of wayland's presentation time protocol for more accurate frame presentation if it is supported by the compositor")
	rootCmd.Flags().String("wid", "", "This tells mpv to attach to an existing window")
	rootCmd.Flags().String("window-dragging", "", "Move the window when clicking on it and moving the mouse pointer")
	rootCmd.Flags().String("window-maximized", "", "Whether the video window is maximized or not")
	rootCmd.Flags().String("window-minimized", "", "Whether the video window is minimized or not")
	rootCmd.Flags().String("window-scale", "", "Resize the video window to a multiple (or fraction) of the video size")
	rootCmd.Flags().String("write-filename-in-watch-later-config", "", "Prepend the watch later config files with the name of the file they refer to")
	rootCmd.Flags().String("x11-bypass-compositor", "", "If set to yes, then ask the compositor to unredirect the mpv window")
	rootCmd.Flags().String("x11-name", "", "Set the window instance name for X11-based video output methods")
	rootCmd.Flags().String("x11-netwm", "", "(X11 only) Control the use of NetWM protocol features")
	rootCmd.Flags().String("x11-present", "", "Whether or not to use presentation statistics from X11's presentation extension")
	rootCmd.Flags().String("x11-wid-title", "", "Whether or not to set the window title when mpv is embedded on X11")
	rootCmd.Flags().String("xv-adaptor", "", "Select a specific XVideo adapter")
	rootCmd.Flags().String("xv-buffers", "", "Number of image buffers to use for the internal ringbuffer")
	rootCmd.Flags().String("xv-ck", "", "Select the source from which the color key is taken")
	rootCmd.Flags().String("xv-ck-method", "", "Sets the color key drawing method")
	rootCmd.Flags().String("xv-colorkey", "", "Changes the color key to an RGB value of your choice")
	rootCmd.Flags().String("xv-port", "", "Select a specific XVideo port")
	rootCmd.Flags().String("ytdl", "", "Enable the youtube-dl hook-script")
	rootCmd.Flags().String("ytdl-format", "", "Video format/quality that is directly passed to youtube-dl")
	rootCmd.Flags().String("ytdl-raw-options", "", "Pass arbitrary options to youtube-dl")
	rootCmd.Flags().String("ytdl-raw-options-add", "", "")
	rootCmd.Flags().String("ytdl-raw-options-append", "", "")
	rootCmd.Flags().String("ytdl-raw-options-clr", "", "")
	rootCmd.Flags().String("ytdl-raw-options-del", "", "")
	rootCmd.Flags().String("ytdl-raw-options-remove", "", "")
	rootCmd.Flags().String("ytdl-raw-options-set", "", "")
	rootCmd.Flags().String("zimg-dither", "", "Dithering")
	rootCmd.Flags().String("zimg-fast", "", "Allow optimizations that help with performance, but reduce quality")
	rootCmd.Flags().String("zimg-scaler", "", "Zimg luma scaler to use")
	rootCmd.Flags().String("zimg-scaler-chroma", "", "Same as --zimg-scaler, for for chroma interpolation")
	rootCmd.Flags().String("zimg-scaler-chroma-param-a", "", "Same as --zimg-scaler-param-a, for chroma")
	rootCmd.Flags().String("zimg-scaler-chroma-param-b", "", "Same as --zimg-scaler-param-b, for chroma")
	rootCmd.Flags().String("zimg-scaler-param-a", "", "Set scaler parameters")
	rootCmd.Flags().String("zimg-scaler-param-b", "", "Set scaler parameters")
	rootCmd.Flags().String("zimg-threads", "", "Set the maximum number of threads to use for scaling")
	rootCmd.Flags().Bool("{", false, "")
	rootCmd.Flags().Bool("}", false, "")

	rootCmd.Flags().VisitAll(func(f *pflag.Flag) {
		if f.NoOptDefVal == "" {
			f.NoOptDefVal = " "
		}
	})

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"ab-loop-count":                        carapace.ActionValues("inf", "help"),
		"access-references":                    carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"ad":                                   mpv.ActionAudioDecoders().UniqueList(","),
		"ad-lavc-downmix":                      carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"ad-queue-enable":                      carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"af":                                   mpv.ActionAudioFilters().List(","),
		"aid":                                  carapace.ActionValues("no", "auto", "help").StyleF(style.ForKeyword),
		"alang":                                osAction.ActionLanguages().List(","),
		"allow-delayed-peak-detect":            carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"alsa-ignore-chmap":                    carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"alsa-non-interleaved":                 carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"alsa-resample":                        carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"ao":                                   mpv.ActionAudioOutputs().UniqueList(","),
		"ao-null-broken-delay":                 carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"ao-null-broken-eof":                   carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"ao-null-channel-layouts":              mpv.ActionAudioChannels().List(","),
		"ao-null-format":                       carapace.ActionValues("u8", "s16", "s32", "s64", "float", "double", "u8p", "s16p", "s32p", "s64p", "floatp", "doublep", "spdif-aac", "spdif-ac3", "spdif-dts", "spdif-dtshd", "spdif-eac3", "spdif-mp3", "spdif-truehd", "help"),
		"ao-null-untimed":                      carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"ao-pcm-append":                        carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"ao-pcm-file":                          carapace.ActionFiles(),
		"ao-pcm-waveheader":                    carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"archive-exts":                         carapace.ActionValues("zip", "rar", "7z", "cbz", "cbr", "tar", "gz", "xz", "bz2", "ar").UniqueList(","),
		"audio":                                carapace.ActionValues("no", "auto", "help").StyleF(style.ForKeyword),
		"audio-backward-overlap":               carapace.ActionValues("auto", "help"),
		"audio-channels":                       mpv.ActionAudioChannels().List(","),
		"audio-demuxer":                        mpv.ActionDemuxers(),
		"audio-display":                        carapace.ActionValues("no", "embedded-first", "external-first", "help").StyleF(style.ForKeyword),
		"audio-exclusive":                      carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"audio-fallback-to-null":               carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"audio-file-auto":                      carapace.ActionValues("no", "exact", "fuzzy", "all", "help").StyleF(style.ForKeyword),
		"audio-file-paths":                     carapace.ActionFiles().List(string(os.PathListSeparator)),
		"audio-files":                          carapace.ActionFiles().List(string(os.PathListSeparator)),
		"audio-normalize-downmix":              carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"audio-pitch-correction":               carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"audio-resample-linear":                carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"audio-stream-silence":                 carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"auto-window-resize":                   carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"autoload-files":                       carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"autosync":                             carapace.ActionValues("no", "help").StyleF(style.ForKeyword),
		"blend-subtitles":                      carapace.ActionValues("no", "yes", "video", "help").StyleF(style.ForKeyword),
		"bluray-device":                        carapace.ActionFiles(),
		"border":                               carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"cache":                                carapace.ActionValues("no", "auto", "yes", "help").StyleF(style.ForKeyword),
		"cache-on-disk":                        carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"cache-pause":                          carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"cache-pause-initial":                  carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"cdda-cdtext":                          carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"cdda-skip":                            carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"chapters-file":                        carapace.ActionFiles(),
		"config":                               carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"config-dir":                           carapace.ActionFiles(),
		"cookies":                              carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"cookies-file":                         carapace.ActionFiles(),
		"correct-downscaling":                  carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"correct-pts":                          carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"cover-art-auto":                       carapace.ActionValues("no", "exact", "fuzzy", "all", "help").StyleF(style.ForKeyword),
		"cover-art-files":                      carapace.ActionFiles().List(string(os.PathListSeparator)),
		"cuda-decode-device":                   carapace.ActionValues("auto", "help"),
		"cursor-autohide":                      carapace.ActionValues("no", "always", "help").StyleF(style.ForKeyword),
		"cursor-autohide-fs-only":              carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"deband":                               carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"demuxer-cache-wait":                   carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"demuxer-donate-buffer":                carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"demuxer-lavf-allow-mimetype":          carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"demuxer-lavf-hacks":                   carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"demuxer-lavf-linearize-timestamps":    carapace.ActionValues("no", "auto", "yes", "help").StyleF(style.ForKeyword),
		"demuxer-lavf-probe-info":              carapace.ActionValues("no", "yes", "auto", "nostreams", "help").StyleF(style.ForKeyword),
		"demuxer-lavf-propagate-opts":          carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"demuxer-mkv-probe-start-time":         carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"demuxer-mkv-probe-video-duration":     carapace.ActionValues("no", "yes", "full", "help").StyleF(style.ForKeyword),
		"demuxer-mkv-subtitle-preroll":         carapace.ActionValues("no", "yes", "index", "help").StyleF(style.ForKeyword),
		"demuxer-rawaudio-format":              carapace.ActionValues("u8", "s8", "u16le", "u16be", "s16le", "s16be", "u24le", "u24be", "s24le", "s24be", "u32le", "u32be", "s32le", "s32be", "floatle", "floatbe", "doublele", "doublebe", "u16", "s16", "u24", "s24", "u32", "s32", "float", "double", "help"),
		"demuxer-seekable-cache":               carapace.ActionValues("auto", "no", "yes", "help").StyleF(style.ForKeyword),
		"demuxer-thread":                       carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"dither":                               carapace.ActionValues("fruit", "ordered", "error-diffusion", "no", "help").StyleF(style.ForKeyword),
		"dither-depth":                         carapace.ActionValues("no", "auto", "help").StyleF(style.ForKeyword),
		"drm-device":                           carapace.ActionFiles(),
		"drm-draw-plane":                       carapace.ActionValues("primary", "overlay", "help"),
		"drm-drmprime-video-plane":             carapace.ActionValues("primary", "overlay", "help"),
		"drm-format":                           carapace.ActionValues("xrgb8888", "xrgb2101010", "xbgr8888", "xbgr2101010", "yuyv", "help"),
		"dump-stats":                           carapace.ActionFiles(),
		"dvbin-file":                           carapace.ActionFiles(),
		"dvbin-full-transponder":               carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"dvd-device":                           carapace.ActionFiles(),
		"edition":                              carapace.ActionValues("auto", "help"),
		"egl-output-format":                    carapace.ActionValues("auto", "rgb8", "rgba8", "rgb10", "rgb10_a2", "rgb16", "rgba16", "rgb16f", "rgba16f", "rgb32f", "rgba32f", "help"),
		"embeddedfonts":                        carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"error-diffusion":                      carapace.ActionValues("simple", "false-fs", "sierra-lite", "floyd-steinberg", "atkinson", "jarvis-judice-ninke", "stucki", "burkes", "sierra-3", "sierra-2", "help"),
		"external-files":                       carapace.ActionFiles().List(string(os.PathListSeparator)),
		"focus-on":                             carapace.ActionValues("all", "never", "open", "help"),
		"force-rgba-osd-rendering":             carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"force-seekable":                       carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"force-window":                         carapace.ActionValues("no", "yes", "immediate", "help").StyleF(style.ForKeyword),
		"force-window-position":                carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"framedrop":                            carapace.ActionValues("no", "vo", "decoder", "decoder+vo", "help").StyleF(style.ForKeyword),
		"frames":                               carapace.ActionValues("all", "help"),
		"fs-screen":                            carapace.ActionValues("all", "current", "help"),
		"fullscreen":                           carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"gamut-mapping-mode":                   carapace.ActionValues("auto", "clip", "perceptual", "relative", "saturation", "absolute", "desaturate", "darken", "warn", "linear", "help"),
		"gapless-audio":                        carapace.ActionValues("no", "yes", "weak", "help").StyleF(style.ForKeyword),
		"glsl-shaders":                         carapace.ActionFiles().List(string(os.PathListSeparator)),
		"gpu-api":                              carapace.ActionValues("auto", "opengl", "vulkan", "d3d11", "help").UniqueList(","),
		"gpu-context":                          mpv.ActionGPUContexts().UniqueList(","),
		"gpu-debug":                            carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"gpu-dumb-mode":                        carapace.ActionValues("auto", "yes", "no", "help").StyleF(style.ForKeyword),
		"gpu-shader-cache-dir":                 carapace.ActionFiles(),
		"gpu-sw":                               carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"hdr-compute-peak":                     carapace.ActionValues("auto", "yes", "no", "help").StyleF(style.ForKeyword),
		"hidpi-window-scale":                   carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"hls-bitrate":                          carapace.ActionValues("no", "min", "max", "help").StyleF(style.ForKeyword),
		"hr-seek":                              carapace.ActionValues("no", "absolute", "yes", "always", "default", "help").StyleF(style.ForKeyword),
		"hr-seek-framedrop":                    carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"hwdec":                                mpv.ActionHardwareDecodingAPIs(),
		"hwdec-software-fallback":              carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"icc-cache-dir":                        carapace.ActionFiles(),
		"icc-force-contrast":                   carapace.ActionValues("no", "inf", "help").StyleF(style.ForKeyword),
		"icc-profile":                          carapace.ActionFiles(),
		"icc-profile-auto":                     carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"idle":                                 carapace.ActionValues("no", "once", "yes", "help").StyleF(style.ForKeyword),
		"ignore-path-in-watch-later-config":    carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"image-lut":                            carapace.ActionFiles(),
		"image-lut-type":                       carapace.ActionValues("auto", "native", "normalized", "conversion", "help"),
		"image-subs-video-resolution":          carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"include":                              carapace.ActionFiles(),
		"index":                                carapace.ActionValues("default", "recreate", "help"),
		"initial-audio-sync":                   carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"input-builtin-bindings":               carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"input-conf":                           carapace.ActionFiles(),
		"input-cursor":                         carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"input-default-bindings":               carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"input-ipc-server":                     carapace.ActionFiles(),
		"input-media-keys":                     carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"input-right-alt-gr":                   carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"input-terminal":                       carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"input-test":                           carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"input-vo-keyboard":                    carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"interpolation":                        carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"jack-autostart":                       carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"jack-connect":                         carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"jack-std-channel-layout":              carapace.ActionValues("waveext", "any", "help"),
		"keep-open":                            carapace.ActionValues("no", "yes", "always", "help").StyleF(style.ForKeyword),
		"keep-open-pause":                      carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"keepaspect":                           carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"keepaspect-window":                    carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"linear-downscaling":                   carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"linear-upscaling":                     carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"list-properties":                      carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"load-auto-profiles":                   carapace.ActionValues("no", "yes", "auto", "help").StyleF(style.ForKeyword),
		"load-scripts":                         carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"load-stats-overlay":                   carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"load-unsafe-playlists":                carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"log-file":                             carapace.ActionFiles(),
		"loop-file":                            carapace.ActionValues("no", "inf", "yes", "help").StyleF(style.ForKeyword),
		"loop-playlist":                        carapace.ActionValues("no", "inf", "yes", "force", "help").StyleF(style.ForKeyword),
		"lut":                                  carapace.ActionFiles(),
		"lut-type":                             carapace.ActionValues("auto", "native", "normalized", "conversion", "help"),
		"merge-files":                          carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"msg-color":                            carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"msg-module":                           carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"msg-time":                             carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"mute":                                 carapace.ActionValues("no", "auto", "yes", "help").StyleF(style.ForKeyword),
		"native-fs":                            carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"native-keyrepeat":                     carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"o":                                    carapace.ActionFiles(),
		"oac":                                  mpv.ActionAudioCodecs(),
		"oacopts":                              mpv.ActionAudioCodecOptions(),
		"ocopy-metadata":                       carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"of":                                   mpv.ActionFormats(),
		"ofopts":                               mpv.ActionFormatOptions(),
		"on-all-workspaces":                    carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"ontop":                                carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"ontop-level":                          carapace.ActionValues("window", "system", "desktop", "help"),
		"opengl-early-flush":                   carapace.ActionValues("no", "yes", "auto", "help").StyleF(style.ForKeyword),
		"opengl-es":                            carapace.ActionValues("auto", "yes", "no", "help").StyleF(style.ForKeyword),
		"opengl-glfinish":                      carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"opengl-pbo":                           carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"opengl-rectangle-textures":            carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"opengl-waitvsync":                     carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"orawts":                               carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"ordered-chapters":                     carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"ordered-chapters-files":               carapace.ActionFiles(),
		"osc":                                  carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"osd-align-x":                          carapace.ActionValues("left", "center", "right", "help"),
		"osd-align-y":                          carapace.ActionValues("top", "center", "bottom", "help"),
		"osd-bar":                              carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"osd-bold":                             carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"osd-font-provider":                    carapace.ActionValues("auto", "none", "fontconfig", "help"),
		"osd-fractions":                        carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"osd-italic":                           carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"osd-justify":                          carapace.ActionValues("auto", "left", "center", "right", "help"),
		"osd-level":                            carapace.ActionValues("0", "1", "2", "3", "help"),
		"osd-on-seek":                          carapace.ActionValues("no", "bar", "msg", "msg-bar", "help").StyleF(style.ForKeyword),
		"osd-scale-by-window":                  carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"ovc":                                  mpv.ActionVideoCodecs(),
		"ovcopts":                              mpv.ActionVideoCodecOptions(),
		"pause":                                carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"play-direction":                       carapace.ActionValues("forward", "backward", "help"),
		"player-operation-mode":                carapace.ActionValues("cplayer", "pseudo-gui", "help"),
		"playlist":                             carapace.ActionFiles(),
		"playlist-start":                       carapace.ActionValues("auto", "no", "help").StyleF(style.ForKeyword),
		"prefetch-playlist":                    carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"profile":                              mpv.ActionProfiles().UniqueList(","),
		"pulse-allow-suspended":                carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"pulse-buffer":                         carapace.ActionValues("native", "help"),
		"pulse-latency-hacks":                  carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"quiet":                                carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"rar-list-all-volumes":                 carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"really-quiet":                         carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"rebase-start-time":                    carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"replaygain":                           carapace.ActionValues("no", "track", "album", "help").StyleF(style.ForKeyword),
		"replaygain-clip":                      carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"resume-playback":                      carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"resume-playback-check-mtime":          carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"rtsp-transport":                       carapace.ActionValues("lavf", "udp", "tcp", "http", "udp_multicast", "help"),
		"save-position-on-quit":                carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"scale-window":                         carapace.ActionValues("bartlett", "cosine", "hanning", "tukey", "hamming", "quadric", "welch", "kaiser", "blackman", "sphinx", "jinc", "help"),
		"scaler-resizes-only":                  carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"screen":                               carapace.ActionValues("default", "help"),
		"screenshot-directory":                 carapace.ActionFiles(),
		"screenshot-format":                    carapace.ActionValues("jpg", "jpeg", "png", "webp", "jxl", "avif", "help"),
		"screenshot-high-bit-depth":            carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"screenshot-jpeg-source-chroma":        carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"screenshot-sw":                        carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"screenshot-tag-colorspace":            carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"screenshot-webp-lossless":             carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"scripts":                              carapace.ActionFiles().List(string(os.PathListSeparator)),
		"secondary-sid":                        carapace.ActionValues("no", "auto", "help").StyleF(style.ForKeyword),
		"secondary-sub-visibility":             carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"shuffle":                              carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sid":                                  carapace.ActionValues("no", "auto", "help").StyleF(style.ForKeyword),
		"sigmoid-upscaling":                    carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"snap-window":                          carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"spirv-compiler":                       carapace.ActionValues("auto", "shaderc", "help"),
		"stop-playback-on-init-failure":        carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"stop-screensaver":                     carapace.ActionValues("always", "no", "yes", "help").StyleF(style.ForKeyword),
		"stream-dump":                          carapace.ActionFiles(),
		"stretch-dvd-subs":                     carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"stretch-image-subs-to-screen":         carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-align-x":                          carapace.ActionValues("left", "center", "right", "help"),
		"sub-align-y":                          carapace.ActionValues("top", "center", "bottom", "help"),
		"sub-ass":                              carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-ass-force-margins":                carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-ass-justify":                      carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-ass-override":                     carapace.ActionValues("no", "yes", "force", "scale", "strip", "help").StyleF(style.ForKeyword),
		"sub-ass-scale-with-window":            carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-ass-styles":                       carapace.ActionFiles(),
		"sub-ass-vsfilter-color-compat":        carapace.ActionValues("no", "basic", "full", "force-601", "help").StyleF(style.ForKeyword),
		"sub-auto":                             carapace.ActionValues("no", "exact", "fuzzy", "all", "help").StyleF(style.ForKeyword),
		"sub-bold":                             carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-clear-on-seek":                    carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-create-cc-track":                  carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-file-paths":                       carapace.ActionFiles().List(string(os.PathListSeparator)),
		"sub-files":                            carapace.ActionFiles().List(string(os.PathListSeparator)),
		"sub-filter-regex-enable":              carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-filter-regex-plain":               carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-filter-regex-warn":                carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-filter-sdh":                       carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-filter-sdh-harder":                carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-fix-timing":                       carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-font-provider":                    carapace.ActionValues("auto", "none", "fontconfig", "help"),
		"sub-gray":                             carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-hinting":                          carapace.ActionValues("none", "light", "normal", "native", "help"),
		"sub-italic":                           carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-justify":                          carapace.ActionValues("auto", "left", "center", "right", "help"),
		"sub-past-video-end":                   carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-scale-by-window":                  carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-scale-with-window":                carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-shaper":                           carapace.ActionValues("simple", "complex", "help"),
		"sub-use-margins":                      carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sub-visibility":                       carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"subs-with-matching-audio":             carapace.ActionValues("forced", "no", "yes", "help").StyleF(style.ForKeyword),
		"sws-allow-zimg":                       carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sws-bitexact":                         carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sws-fast":                             carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"sws-scaler":                           carapace.ActionValues("fast-bilinear", "bilinear", "bicubic", "x", "point", "area", "bicublin", "gauss", "sinc", "lanczos", "spline", "help"),
		"target-peak":                          carapace.ActionValues("auto", "help"),
		"target-prim":                          carapace.ActionValues("auto", "bt.601-525", "bt.601-625", "bt.709", "bt.2020", "bt.470m", "apple", "adobe", "prophoto", "cie1931", "dci-p3", "display-p3", "v-gamut", "s-gamut", "ebu3213", "film-c", "aces-ap0", "aces-ap1", "help"),
		"target-trc":                           carapace.ActionValues("auto", "bt.1886", "srgb", "linear", "gamma1.8", "gamma2.0", "gamma2.2", "gamma2.4", "gamma2.6", "gamma2.8", "prophoto", "pq", "hlg", "v-log", "s-log1", "s-log2", "st428", "help"),
		"taskbar-progress":                     carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"temporal-dither":                      carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"term-osd":                             carapace.ActionValues("force", "auto", "no", "help").StyleF(style.ForKeyword),
		"term-osd-bar":                         carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"terminal":                             carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"tls-ca-file":                          carapace.ActionFiles(),
		"tls-cert-file":                        carapace.ActionFiles(),
		"tls-key-file":                         carapace.ActionFiles(),
		"tls-verify":                           carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"tone-mapping":                         carapace.ActionValues("auto", "clip", "mobius", "reinhard", "hable", "gamma", "linear", "spline", "bt.2390", "bt.2446a", "st2094-40", "st2094-10", "help"),
		"track-auto-selection":                 carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"untimed":                              carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"use-embedded-icc-profile":             carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"use-filedir-conf":                     carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vd":                                   mpv.ActionVideoDecoders().UniqueList(","),
		"vd-lavc-assume-old-x264":              carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vd-lavc-bitexact":                     carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vd-lavc-check-hw-profile":             carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vd-lavc-dr":                           carapace.ActionValues("auto", "no", "yes", "help").StyleF(style.ForKeyword),
		"vd-lavc-fast":                         carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vd-lavc-framedrop":                    carapace.ActionValues("none", "default", "nonref", "bidir", "nonkey", "all", "help"),
		"vd-lavc-show-all":                     carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vd-lavc-skipframe":                    carapace.ActionValues("none", "default", "nonref", "bidir", "nonkey", "all", "help"),
		"vd-lavc-skipidct":                     carapace.ActionValues("none", "default", "nonref", "bidir", "nonkey", "all", "help"),
		"vd-lavc-skiploopfilter":               carapace.ActionValues("none", "default", "nonref", "bidir", "nonkey", "all", "help"),
		"vd-queue-enable":                      carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vf":                                   mpv.ActionVideoFilters().List(","),
		"vid":                                  carapace.ActionValues("no", "auto", "help").StyleF(style.ForKeyword),
		"video-aspect-method":                  carapace.ActionValues("bitstream", "container", "help"),
		"video-backward-overlap":               carapace.ActionValues("auto", "help"),
		"video-latency-hacks":                  carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"video-osd":                            carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"video-output-levels":                  carapace.ActionValues("auto", "limited", "full", "help"),
		"video-rotate":                         carapace.ActionValues("no", "help").StyleF(style.ForKeyword),
		"video-sync":                           carapace.ActionValues("audio", "display-resample", "display-resample-vdrop", "display-resample-desync", "display-tempo", "display-adrop", "display-vdrop", "display-desync", "desync", "help"),
		"video-unscaled":                       carapace.ActionValues("no", "yes", "downscale-big", "help").StyleF(style.ForKeyword),
		"vo":                                   mpv.ActionVideoOutputs(),
		"vo-image-format":                      carapace.ActionValues("jpg", "jpeg", "png", "webp", "jxl", "avif", "help"),
		"vo-image-high-bit-depth":              carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vo-image-jpeg-source-chroma":          carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vo-image-outdir":                      carapace.ActionFiles(),
		"vo-image-tag-colorspace":              carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vo-image-webp-lossless":               carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vo-tct-256":                           carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vo-tct-algo":                          carapace.ActionValues("plain", "half-blocks", "help"),
		"vo-vaapi-scaled-osd":                  carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vo-vaapi-scaling":                     carapace.ActionValues("default", "fast", "hq", "nla", "help"),
		"vo-vdpau-chroma-deint":                carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vo-vdpau-composite-detect":            carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vo-vdpau-force-yuv":                   carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vo-vdpau-pullup":                      carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vulkan-async-compute":                 carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vulkan-async-transfer":                carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"vulkan-swap-mode":                     carapace.ActionValues("auto", "fifo", "fifo-relaxed", "mailbox", "immediate", "help"),
		"watch-later-directory":                carapace.ActionFiles(),
		"wayland-internal-vsync":               carapace.ActionValues("auto", "no", "yes", "help").StyleF(style.ForKeyword),
		"window-dragging":                      carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"window-maximized":                     carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"window-minimized":                     carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"write-filename-in-watch-later-config": carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"x11-bypass-compositor":                carapace.ActionValues("no", "yes", "fs-only", "never", "help").StyleF(style.ForKeyword),
		"x11-netwm":                            carapace.ActionValues("auto", "no", "yes", "help").StyleF(style.ForKeyword),
		"xv-ck":                                carapace.ActionValues("use", "set", "cur", "help"),
		"xv-ck-method":                         carapace.ActionValues("none", "bg", "man", "auto", "help"),
		"ytdl":                                 carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"zimg-dither":                          carapace.ActionValues("no", "ordered", "random", "error-diffusion", "help").StyleF(style.ForKeyword),
		"zimg-fast":                            carapace.ActionValues("no", "yes", "help").StyleF(style.ForKeyword),
		"zimg-scaler":                          carapace.ActionValues("point", "bilinear", "bicubic", "spline16", "spline36", "lanczos", "help"),
		"zimg-scaler-chroma":                   carapace.ActionValues("point", "bilinear", "bicubic", "spline16", "spline36", "lanczos", "help"),
		"zimg-threads":                         carapace.ActionValues("auto", "help"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
