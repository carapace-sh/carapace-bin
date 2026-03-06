package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "feh",
	Short: "a fast and light image viewer",
	Long:  "https://feh.finalrewind.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("action", "A", "", "Specify action to perform when pressing <return>")
	rootCmd.Flags().String("action1", "", "Extra action triggered with number key 1")
	rootCmd.Flags().String("action2", "", "Extra action triggered with number key 2")
	rootCmd.Flags().String("action3", "", "Extra action triggered with number key 3")
	rootCmd.Flags().String("action4", "", "Extra action triggered with number key 4")
	rootCmd.Flags().String("action5", "", "Extra action triggered with number key 5")
	rootCmd.Flags().String("action6", "", "Extra action triggered with number key 6")
	rootCmd.Flags().String("action7", "", "Extra action triggered with number key 7")
	rootCmd.Flags().String("action8", "", "Extra action triggered with number key 8")
	rootCmd.Flags().String("action9", "", "Extra action triggered with number key 9")
	rootCmd.Flags().StringP("alpha", "a", "", "Set thumbnail transparency level (0 .. 255)")
	rootCmd.Flags().Bool("auto-reload", false, "automatically reload shown image if file was changed")
	rootCmd.Flags().Bool("auto-rotate", false, "Rotate images according to Exif info (if compiled with exif=1)")
	rootCmd.Flags().BoolP("auto-zoom", "Z", false, "Zoom picture to screen size in fullscreen/geom mode")
	rootCmd.Flags().StringP("bg", "b", "", "Set montage background")
	rootCmd.Flags().String("bg-center", "", "Set FILE as centered desktop background")
	rootCmd.Flags().String("bg-fill", "", "Like --bg-scale, but preserves aspect ratio by zooming the image until it fits. May cut off corners")
	rootCmd.Flags().String("bg-max", "", "Like --bg-fill, but scale the image to the maximum size that fits the screen with black borders on one side")
	rootCmd.Flags().String("bg-scale", "", "Set FILE as scaled desktop background. This will fill the whole background, but the images' aspect ratio may not be preserved")
	rootCmd.Flags().String("bg-tile", "", "Set FILE as tiled desktop background")
	rootCmd.Flags().BoolP("borderless", "x", false, "Create borderless windows")
	rootCmd.Flags().String("cache-size", "", "imlib cache size in mebibytes (0 .. 2048)")
	rootCmd.Flags().BoolP("cache-thumbnails", "P", false, "Enable thumbnail caching for thumbnail mode. Only works with thumbnails <= 256x256 pixels")
	rootCmd.Flags().StringP("caption-path", "K", "", "Path to caption directory, enables caption display")
	rootCmd.Flags().String("class", "", "Set the X11 class hint to class")
	rootCmd.Flags().Bool("conversion-timeout", false, "INT  Load unknown files with dcraw or ImageMagick, timeout after INT seconds (0: no timeout)")
	rootCmd.Flags().StringP("customlist", "L", "", "list mode with custom output, see FORMAT SPECIFIERS")
	rootCmd.Flags().BoolP("draw-actions", "G", false, "Show the defined actions in the image window")
	rootCmd.Flags().Bool("draw-exif", false, "Show some Exif information (if compiled with exif=1)")
	rootCmd.Flags().BoolP("draw-filename", "d", false, "Show the filename in the image window")
	rootCmd.Flags().Bool("draw-tinted", false, "Show overlay texts on semi-transparent background")
	rootCmd.Flags().Bool("edit", false, "Make flip/rotation keys flip/rotate the underlying file")
	rootCmd.Flags().StringP("filelist", "f", "", "Load/save images from/to the FILE filelist")
	rootCmd.Flags().StringP("font", "e", "", "Set font for thumbnail information, in the form fontname/pointsize")
	rootCmd.Flags().StringP("fontpath", "C", "", "Specify an extra directory to look in for fonts, can be used multiple times to add multiple paths.")
	rootCmd.Flags().Bool("force-aliasing", false, "Disable antialiasing")
	rootCmd.Flags().BoolP("fullindex", "I", false, "Index mode with additional image information")
	rootCmd.Flags().BoolP("fullscreen", "F", false, "Make the window full screen")
	rootCmd.Flags().StringP("geometry", "g", "", "Limit the window size to DIMENSION[+OFFSET]")
	rootCmd.Flags().BoolP("help", "h", false, "Show help and exit")
	rootCmd.Flags().BoolP("hide-pointer", "Y", false, "Hide the pointer")
	rootCmd.Flags().BoolP("ignore-aspect", "X", false, "Set thumbnail to specified width/height without retaining aspect ratio")
	rootCmd.Flags().StringP("image-bg", "B", "", "Set background for transparent images and the like. Accepted values: default, checks, or a XColor (eg. #428bdd)")
	rootCmd.Flags().BoolP("index", "i", false, "Create an index print of all images")
	rootCmd.Flags().String("index-info", "", "Show FORMAT below images in index/thumbnail mode")
	rootCmd.Flags().String("info", "", "Run CMD and show its output in the image window")
	rootCmd.Flags().Bool("insecure", false, "Disable peer/host verification when using HTTPS.")
	rootCmd.Flags().BoolP("keep-http", "k", false, "Keep local copies when viewing HTTP/FTP files")
	rootCmd.Flags().Bool("keep-zoom-vp", false, "Keep viewport zoom and settings while changing images")
	rootCmd.Flags().StringP("limit-height", "H", "", "Limit the height of the montage in pixels (at least one of these two must be specified)")
	rootCmd.Flags().StringP("limit-width", "W", "", "Limit the width of the montage in pixels")
	rootCmd.Flags().BoolP("list", "l", false, "list mode: ls-style output with image information")
	rootCmd.Flags().BoolP("loadable", "U", false, "List all loadable files. No image display")
	rootCmd.Flags().String("max-dimension", "", "Only show images with width <= W and height <= H")
	rootCmd.Flags().StringP("menu-font", "M", "", "Use FONT for the font in menus.")
	rootCmd.Flags().String("min-dimension", "", "Only show images with width >= W and height >= H")
	rootCmd.Flags().BoolP("montage", "m", false, "Enable montage mode")
	rootCmd.Flags().BoolP("multiwindow", "w", false, "Open all files at once, one window per image")
	rootCmd.Flags().Bool("no-fehbg", false, "Do not write a ~/.fehbg file")
	rootCmd.Flags().Bool("no-jump-on-resort", false, "Don't jump to the first image when the filelist is resorted")
	rootCmd.Flags().BoolP("no-menus", "N", false, "Don't load or show any menus.")
	rootCmd.Flags().Bool("no-recursive", false, "Do not recursively expand directories (this is the default)")
	rootCmd.Flags().Bool("no-screen-clip", false, "Do not limit window size to screen size")
	rootCmd.Flags().Bool("no-xinerama", false, "Disable Xinerama support")
	rootCmd.Flags().String("on-last-slide", "", "Exit after one loop through the slide show (old --cycle-once)")
	rootCmd.Flags().StringP("output", "o", "", "Save the created montage to FILE")
	rootCmd.Flags().BoolP("output-dir", "j", false, "With -k: Output directory for saved files")
	rootCmd.Flags().StringP("output-only", "O", "", "Just save the created montage to FILE WITHOUT displaying it")
	rootCmd.Flags().BoolP("preload", "p", false, "Remove unloadable files from the internal filelist before attempting to display anything")
	rootCmd.Flags().BoolP("quiet", "q", false, "Hide non-fatal errors. May be used with --verbose")
	rootCmd.Flags().BoolP("randomize", "z", false, "Randomize the filelist")
	rootCmd.Flags().BoolP("recursive", "r", false, "Recursively expand any directories in FILE to the content of those directories")
	rootCmd.Flags().StringP("reload", "R", "", "Reload images after NUM seconds")
	rootCmd.Flags().BoolP("reverse", "n", false, "Reverse sort order")
	rootCmd.Flags().BoolP("scale-down", ".", false, "Automatically scale down images to fit screen size")
	rootCmd.Flags().String("scroll-step", "", "scroll COUNT pixels when movement key is pressed")
	rootCmd.Flags().StringP("slideshow-delay", "D", "", "Set delay between automatically changing slides")
	rootCmd.Flags().StringP("sort", "S", "", "Sort files by:")
	rootCmd.Flags().StringP("start-at", "|", "", "Start at FILENAME in the filelist")
	rootCmd.Flags().BoolP("stretch", "s", false, "Scale up images if they are smaller than the specified thumbnail size")
	rootCmd.Flags().Bool("tap-zones", false, "Enable tap zones for previous/next file in slide show mode")
	rootCmd.Flags().StringP("theme", "T", "", "Load options with name THEME")
	rootCmd.Flags().StringP("thumb-height", "E", "", "Set thumbnail height in pixels")
	rootCmd.Flags().StringP("thumb-redraw", "J", "", "Redraw thumbnail window every N images")
	rootCmd.Flags().StringP("thumb-title", "~", "", "Title for windows opened from thumbnail mode")
	rootCmd.Flags().StringP("thumb-width", "y", "", "Set thumbnail width in pixels")
	rootCmd.Flags().BoolP("thumbnails", "t", false, "Show images as clickable thumbnails")
	rootCmd.Flags().StringP("title", "^", "", "Set window title (see FORMAT SPECIFIERS)")
	rootCmd.Flags().StringP("title-font", "@", "", "Use FONT to print a title on the index, if no font is specified, a title will not be printed")
	rootCmd.Flags().BoolP("unloadable", "u", false, "List all unloadable files. No image display")
	rootCmd.Flags().BoolP("verbose", "V", false, "Show progress bars and other extra information")
	rootCmd.Flags().BoolP("version", "v", false, "Show version information and exit")
	rootCmd.Flags().Bool("version-sort", false, "Natural sort of (version) numbers within text")
	rootCmd.Flags().String("window-id", "", "Draw to an existing X11 window by its ID")
	rootCmd.Flags().String("xinerama-index", "", "Assumee that I is the active xinerama screen")
	rootCmd.Flags().String("zoom", "", "Zooms images by a PERCENT, when in full screen mode or when window geometry is fixed. If combined with --auto-zoom, zooming will be limited to the the size. Also support \"max\" and \"fill\"")
	rootCmd.Flags().String("zoom-step", "", "Zoom images in and out by PERCENT (default: 25) when using the zoom keys / buttons")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"bg": carapace.Batch(
			carapace.ActionFiles(),
			carapace.ActionValues("trans"),
		).ToA(),
		"caption-path":  carapace.ActionDirectories(),
		"filelist":      carapace.ActionFiles(),
		"fontpath":      carapace.ActionDirectories(),
		"on-last-slide": carapace.ActionValues("hold", "quit", "resume"),
		"output":        carapace.ActionFiles(),
		"output-dir":    carapace.ActionDirectories(),
		"output-only":   carapace.ActionFiles(),
		"sort":          carapace.ActionValues("name", "none", "dirname", "filename", "mtime", "width", "height", "pixels", "size", "format"),
		"zoom":          carapace.ActionValues("max", "fill"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
