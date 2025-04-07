package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scrcpy",
	Short: "Display and control your Android device",
	Long:  "https://github.com/Genymobile/scrcpy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("G", "G", false, "Same as --gamepad=uhid, or --keyboard=aoa if --otg is set")
	rootCmd.Flags().BoolS("K", "K", false, "Same as --keyboard=uhid, or --keyboard=aoa if --otg is set.")
	rootCmd.Flags().BoolS("M", "M", false, "Same as --mouse=uhid, or --mouse=aoa if --otg is set")
	rootCmd.Flags().Bool("always-on-top", false, "Make scrcpy window always on top")
	rootCmd.Flags().String("angle", "", "Rotate the video content by a custom angle")
	rootCmd.Flags().String("audio-bit-rate", "", "Encode the audio at the given bit rate")
	rootCmd.Flags().String("audio-buffer", "", "Configure the audio buffering delay")
	rootCmd.Flags().String("audio-codec", "", "Select an audio codec")
	rootCmd.Flags().String("audio-codec-options", "", "Set options for the device audio encoder")
	rootCmd.Flags().Bool("audio-dup", false, "Duplicate audio")
	rootCmd.Flags().String("audio-encoder", "", "Use a specific MediaCodec audio encoder")
	rootCmd.Flags().Bool("audio-output-buffer", false, "ms Configure the size of the SDL audio output buffer")
	rootCmd.Flags().String("audio-source", "", "Select the audio source")
	rootCmd.Flags().String("camera-ar", "", "Select the camera size by its aspect ratio")
	rootCmd.Flags().String("camera-facing", "", "Select the device camera by its facing direction")
	rootCmd.Flags().String("camera-fps", "", "Specify the camera capture frame rate")
	rootCmd.Flags().Bool("camera-high-speed", false, "Enable high-speed camera capture mode")
	rootCmd.Flags().String("camera-id", "", "Specify the device camera id to mirror")
	rootCmd.Flags().String("camera-size", "", "Specify an explicit camera capture size")
	rootCmd.Flags().String("capture-orientation", "", "Specify capture orientation")
	rootCmd.Flags().String("crop", "", "Crop the device screen on the server")
	rootCmd.Flags().Bool("disable-screensaver", false, "Disable screensaver while scrcpy is running")
	rootCmd.Flags().String("display-id", "", "Specify the device display id to mirror")
	rootCmd.Flags().String("display-orientation", "", "Set the initial display orientation")
	rootCmd.Flags().Bool("force-adb-forward", false, "Do not attempt to use \"adb reverse\" to connect to the device")
	rootCmd.Flags().BoolP("fullscreen", "f", false, "Start in fullscreen")
	rootCmd.Flags().String("gamepad", "", "Select how to send gamepad inputs to the device")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().String("keyboard", "", "Select how to send keyboard inputs to the device")
	rootCmd.Flags().Bool("kill-adb-on-close", false, "Kill adb when scrcpy terminates")
	rootCmd.Flags().Bool("legacy-paste", false, "Inject computer clipboard text as a sequence of key events on Ctrl+v")
	rootCmd.Flags().Bool("list-apps", false, "List Android apps installed on the device")
	rootCmd.Flags().Bool("list-camera-sizes", false, "List the valid camera capture sizes")
	rootCmd.Flags().Bool("list-cameras", false, "List cameras available on the device")
	rootCmd.Flags().Bool("list-displays", false, "List displays available on the device")
	rootCmd.Flags().Bool("list-encoders", false, "List video and audio encoders available on the device")
	rootCmd.Flags().String("max-fps", "", "Limit the framerate of screen capture")
	rootCmd.Flags().StringP("max-size", "m", "", "Limit both the width and height of the video to value")
	rootCmd.Flags().String("mouse", "", "Select how to send mouse inputs to the device")
	rootCmd.Flags().String("mouse-bind", "", "Configure bindings of secondary clicks")
	rootCmd.Flags().String("new-display", "", "Create a new display with the specified resolution and density")
	rootCmd.Flags().Bool("no-audio", false, "Disable audio forwarding")
	rootCmd.Flags().Bool("no-audio-playback", false, "Disable audio playback on the computer")
	rootCmd.Flags().Bool("no-cleanup", false, "Do not restore the device state")
	rootCmd.Flags().Bool("no-clipboard-autosync", false, "Do not automatically synchronizes the computer clipboard")
	rootCmd.Flags().BoolP("no-control", "n", false, "Disable device control")
	rootCmd.Flags().Bool("no-downsize-on-error", false, "Do no automatically try again with a lower definition")
	rootCmd.Flags().Bool("no-key-repeat", false, "Do not forward repeated key events when a key is held down")
	rootCmd.Flags().Bool("no-mipmaps", false, "Disables the generation of mipmaps")
	rootCmd.Flags().Bool("no-mouse-hover", false, "Do not forward mouse hover events")
	rootCmd.Flags().BoolP("no-playback", "N", false, "Disable video and audio playback on the computer")
	rootCmd.Flags().Bool("no-power-on", false, "Do not power on the device on start")
	rootCmd.Flags().Bool("no-vd-system-decorations", false, "Disable virtual display system decorations flag")
	rootCmd.Flags().Bool("no-video", false, "Disable video forwarding")
	rootCmd.Flags().Bool("no-video-playback", false, "Disable video playback on the computer")
	rootCmd.Flags().Bool("no-window", false, "Disable scrcpy window")
	rootCmd.Flags().String("orientation", "", "Same as --display-orientation=value --record-orientation=value")
	rootCmd.Flags().Bool("otg", false, "Run in OTG mode")
	rootCmd.Flags().String("pause-on-exit", "", "Configure pause on exit")
	rootCmd.Flags().StringP("port", "p", "", "Set the TCP port (range) used by the client to listen")
	rootCmd.Flags().Bool("power-off-on-close", false, "Turn the device screen off when closing scrcpy")
	rootCmd.Flags().Bool("prefer-text", false, "Inject alpha characters and space as text events instead of key events")
	rootCmd.Flags().Bool("print-fps", false, "Start FPS counter")
	rootCmd.Flags().String("push-target", "", "Set the target directory for pushing files")
	rootCmd.Flags().Bool("raw-key-events", false, "Inject key events for all input keys, and ignore text events")
	rootCmd.Flags().StringP("record", "r", "", "Record screen to file")
	rootCmd.Flags().String("record-format", "", "Force recording format")
	rootCmd.Flags().String("record-orientation", "", "Set the record orientation")
	rootCmd.Flags().String("render-driver", "", "Request SDL to use the given render driver")
	rootCmd.Flags().Bool("require-audio", false, "Fail if audio is enabled but does not work")
	rootCmd.Flags().BoolP("select-tcpip", "e", false, "Use TCP/IP device")
	rootCmd.Flags().BoolP("select-usb", "d", false, "Use USB device")
	rootCmd.Flags().StringP("serial", "s", "", "The device serial number")
	rootCmd.Flags().String("shortcut-mod", "", "Specify the modifiers to use for scrcpy shortcuts")
	rootCmd.Flags().BoolP("show-touches", "t", false, "Enable \"show touches\" on start")
	rootCmd.Flags().String("start-app", "", "Start an Android app, by its exact package name")
	rootCmd.Flags().BoolP("stay-awake", "w", false, "Keep the device on while scrcpy is running")
	rootCmd.Flags().String("tcpip", "", "Configure and reconnect the device over TCP/IP")
	rootCmd.Flags().String("time-limit", "", "Set the maximum mirroring time")
	rootCmd.Flags().String("tunnel-host", "", "Set the IP address of the adb tunnel to reach the scrcpy server")
	rootCmd.Flags().String("tunnel-port", "", "Set the TCP port of the adb tunnel to reach the scrcpy server")
	rootCmd.Flags().BoolP("turn-screen-off", "S", false, "Turn the device screen off immediately")
	rootCmd.Flags().String("v4l2-buffer", "", "Add a buffering delay before pushing frames")
	rootCmd.Flags().String("v4l2-sink", "", "Output to v4l2loopback device")
	rootCmd.Flags().StringP("verbosity", "V", "", "Set the log level")
	rootCmd.Flags().BoolP("version", "v", false, "Print the version of scrcpy")
	rootCmd.Flags().StringP("video-bit-rate", "b", "", "Encode the video at the given bit rate")
	rootCmd.Flags().String("video-buffer", "", "Add a buffering delay before displaying video frames")
	rootCmd.Flags().String("video-codec", "", "Select a video codec")
	rootCmd.Flags().String("video-codec-options", "", "Set a list of options for the device video encoder")
	rootCmd.Flags().String("video-encoder", "", "Use a specific MediaCodec video encoder")
	rootCmd.Flags().String("video-source", "", "Select the video source")
	rootCmd.Flags().Bool("window-borderless", false, "Disable window decorations")
	rootCmd.Flags().String("window-height", "", "Set the initial window height")
	rootCmd.Flags().String("window-title", "", "Set a custom window title")
	rootCmd.Flags().String("window-width", "", "Set the initial window width")
	rootCmd.Flags().String("window-x", "", "Set the initial window horizontal position")
	rootCmd.Flags().String("window-y", "", "Set the initial window vertical position")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"angle":               carapace.ActionValues(),
		"audio-bit-rate":      carapace.ActionValues(),
		"audio-buffer":        carapace.ActionValues(),
		"audio-codec":         carapace.ActionValues(),
		"audio-codec-options": carapace.ActionValues(),
		"audio-encoder":       carapace.ActionValues(),
		"audio-source":        carapace.ActionValues(),
		"camera-ar":           carapace.ActionValues(),
		"camera-facing":       carapace.ActionValues(),
		"camera-fps":          carapace.ActionValues(),
		"camera-id":           carapace.ActionValues(),
		"camera-size":         carapace.ActionValues(),
		"capture-orientation": carapace.ActionValues(),
		"crop":                carapace.ActionValues(),
		"display-id":          carapace.ActionValues(),
		"display-orientation": carapace.ActionValues(),
		"gamepad":             carapace.ActionValues(),
		"keyboard":            carapace.ActionValues(),
		"max-fps":             carapace.ActionValues(),
		"max-size":            carapace.ActionValues(),
		"mouse":               carapace.ActionValues(),
		"mouse-bind":          carapace.ActionValues(),
		"new-display":         carapace.ActionValues(),
		"orientation":         carapace.ActionValues(),
		"pause-on-exit":       carapace.ActionValues(),
		"port":                carapace.ActionValues(),
		"push-target":         carapace.ActionValues(),
		"record":              carapace.ActionValues(),
		"record-format":       carapace.ActionValues(),
		"record-orientation":  carapace.ActionValues(),
		"render-driver":       carapace.ActionValues(),
		"serial":              carapace.ActionValues(),
		"shortcut-mod":        carapace.ActionValues(),
		"start-app":           carapace.ActionValues(),
		"tcpip":               carapace.ActionValues(),
		"time-limit":          carapace.ActionValues(),
		"tunnel-host":         carapace.ActionValues(),
		"tunnel-port":         carapace.ActionValues(),
		"v4l2-buffer":         carapace.ActionValues(),
		"v4l2-sink":           carapace.ActionValues(),
		"verbosity":           carapace.ActionValues(),
		"video-bit-rate":      carapace.ActionValues(),
		"video-buffer":        carapace.ActionValues(),
		"video-codec":         carapace.ActionValues(),
		"video-codec-options": carapace.ActionValues(),
		"video-encoder":       carapace.ActionValues(),
		"video-source":        carapace.ActionValues(),
		"window-height":       carapace.ActionValues(),
		"window-title":        carapace.ActionValues(),
		"window-width":        carapace.ActionValues(),
		"window-x":            carapace.ActionValues(),
		"window-y":            carapace.ActionValues(),
	})
}
