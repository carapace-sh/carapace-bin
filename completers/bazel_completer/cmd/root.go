package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bazel",
	Short: "a fast, scalable, multi-language and extensible build system",
	Long:  "https://bazel.build/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("autodetect_server_javabase", false, "When --noautodetect_server_javabase is passed, Bazel does not fall back to the local JDK for running the bazel server and instead exits")
	rootCmd.Flags().Bool("batch", false, "If set, Bazel will be run as just a client process without a server, instead of in the standard client/server mode")
	rootCmd.Flags().Bool("batch_cpu_scheduling", false, "Only on Linux; use 'batch' CPU scheduling for Blaze")
	rootCmd.Flags().String("bazelrc", "", "The location of the user .bazelrc file containing default values of Bazel options")
	rootCmd.Flags().Bool("block_for_lock", false, "When --noblock_for_lock is passed, Bazel does not wait for a running command to complete, but instead exits immediately")
	rootCmd.Flags().Bool("client_debug", false, "If true, log debug information from the client to stderr")
	rootCmd.Flags().String("connect_timeout_secs", "", "The amount of time the client waits for each attempt to connect to the server Tags: bazel_internal_configuration")
	rootCmd.Flags().String("digest_function", "", "The hash function to use when computing file digests")
	rootCmd.Flags().Bool("expand_configs_in_place", false, "Changed the expansion of --config flags to be done in-place, as opposed to in a fixed point expansion between normal rc options and command-line specified options")
	rootCmd.Flags().String("failure_detail_out", "", "If set, specifies a location to write a failure_detail protobuf message if the server experiences a failure and cannot report it via gRPC, as normal")
	rootCmd.Flags().Bool("home_rc", false, "Whether or not to look for the home bazelrc file at $HOME/.bazelrc Tags: changes_inputs")
	rootCmd.Flags().StringSlice("host_jvm_args", []string{}, "Flags to pass to the JVM executing Blaze")
	rootCmd.Flags().String("host_jvm_profile", "", "Convenience option to add some profiler/debugger-specific JVM startup flags")
	rootCmd.Flags().Bool("idle_server_tasks", false, "Run System.gc() when the server is idle Tags: loses_incremental_state, host_machine_resource_optimizations")
	rootCmd.Flags().Bool("ignore_all_rc_files", false, "Disables all rc files, regardless of the values of other rc-modifying flags, even if these flags come later in the list of startup options")
	rootCmd.Flags().String("io_nice_level", "", "Only on Linux; set a level from 0-7 for best-effort IO scheduling using the sys_ioprio_set system call")
	rootCmd.Flags().String("local_startup_timeout_secs", "", "The maximum amount of time the client waits to connect to the server Tags: bazel_internal_configuration")
	rootCmd.Flags().String("macos_qos_class", "", "Sets the QoS service class of the bazel server when running on macOS")
	rootCmd.Flags().String("max_idle_secs", "", "The number of seconds the build server will wait idling before shutting down")
	rootCmd.Flags().Bool("no-autodetect_server_javabase", false, "When --noautodetect_server_javabase is passed, Bazel does not fall back to the local JDK for running the bazel server and instead exits")
	rootCmd.Flags().Bool("no-batch", false, "If set, Bazel will be run as just a client process without a server, instead of in the standard client/server mode")
	rootCmd.Flags().Bool("no-batch_cpu_scheduling", false, "Only on Linux; use 'batch' CPU scheduling for Blaze")
	rootCmd.Flags().Bool("no-block_for_lock", false, "When --noblock_for_lock is passed, Bazel does not wait for a running command to complete, but instead exits immediately")
	rootCmd.Flags().Bool("no-client_debug", false, "If true, log debug information from the client to stderr")
	rootCmd.Flags().Bool("no-expand_configs_in_place", false, "Changed the expansion of --config flags to be done in-place, as opposed to in a fixed point expansion between normal rc options and command-line specified options")
	rootCmd.Flags().Bool("no-home_rc", false, "Whether or not to look for the home bazelrc file at $HOME/.bazelrc Tags: changes_inputs")
	rootCmd.Flags().Bool("no-idle_server_tasks", false, "Run System.gc() when the server is idle Tags: loses_incremental_state, host_machine_resource_optimizations")
	rootCmd.Flags().Bool("no-ignore_all_rc_files", false, "Disables all rc files, regardless of the values of other rc-modifying flags, even if these flags come later in the list of startup options")
	rootCmd.Flags().Bool("no-preemptible", false, "If true, the command can be preempted if another command is started")
	rootCmd.Flags().Bool("no-shutdown_on_low_sys_mem", false, "If max_idle_secs is set and the build server has been idle for a while, shut down the server when the system is low on free RAM")
	rootCmd.Flags().Bool("no-system_rc", false, "Whether or not to look for the system-wide bazelrc")
	rootCmd.Flags().Bool("no-unlimit_coredumps", false, "Raises the soft coredump limit to the hard limit to make coredumps of the server (including the JVM) and the client possible under common conditions")
	rootCmd.Flags().Bool("no-watchfs", false, "If true, bazel tries to use the operating system's file watch service for local changes instead of scanning every file for a change")
	rootCmd.Flags().Bool("no-windows_enable_symlinks", false, "If true, real symbolic links will be created on Windows instead of file copying")
	rootCmd.Flags().Bool("no-workspace_rc", false, "Whether or not to look for the workspace bazelrc file at $workspace/.bazelrc Tags: changes_inputs")
	rootCmd.Flags().String("output_base", "", "If set, specifies the output location to which all build output will be written")
	rootCmd.Flags().String("output_user_root", "", "The user-specific directory beneath which all build outputs are written; by default, this is a function of $USER, but by specifying a constant, build outputs can be shared between collaborating users")
	rootCmd.Flags().Bool("preemptible", false, "If true, the command can be preempted if another command is started")
	rootCmd.Flags().String("server_javabase", "", "Path to the JVM used to execute Bazel itself")
	rootCmd.Flags().String("server_jvm_out", "", "The location to write the server's JVM's output")
	rootCmd.Flags().Bool("shutdown_on_low_sys_mem", false, "If max_idle_secs is set and the build server has been idle for a while, shut down the server when the system is low on free RAM")
	rootCmd.Flags().Bool("system_rc", false, "Whether or not to look for the system-wide bazelrc")
	rootCmd.Flags().Bool("unlimit_coredumps", false, "Raises the soft coredump limit to the hard limit to make coredumps of the server (including the JVM) and the client possible under common conditions")
	rootCmd.Flags().Bool("watchfs", false, "If true, bazel tries to use the operating system's file watch service for local changes instead of scanning every file for a change")
	rootCmd.Flags().Bool("windows_enable_symlinks", false, "If true, real symbolic links will be created on Windows instead of file copying")
	rootCmd.Flags().Bool("workspace_rc", false, "Whether or not to look for the workspace bazelrc file at $workspace/.bazelrc Tags: changes_inputs")
}
