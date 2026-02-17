package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "bazel",
	Short:   "build system",
	Long:    "https://bazel.build/reference/command-line-reference",
	Aliases: []string{"bazelisk"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("autodetect_server_javabase", false, "When --noautodetect_server_javabase is passed, Bazel does not fall back to the local JDK for running the bazel server and instead exits.")
	rootCmd.Flags().Bool("batch", false, "If set, Bazel will be run as just a client process without a server, instead of in the standard client/server mode. This is deprecated and will be removed, please prefer shutting down the server explicitly if you wish to avoid lingering servers.")
	rootCmd.Flags().Bool("batch_cpu_scheduling", false, "Only on Linux; use 'batch' CPU scheduling for Blaze. This policy is useful for workloads that are non-interactive, but do not want to lower their nice value. See 'man 2 sched_setscheduler'. If false, then Bazel does not perform a system call.")
	rootCmd.Flags().StringSlice("bazelrc", nil, "<path>The location of the user .bazelrc file containing default values of Bazel options.")
	rootCmd.Flags().Bool("block_for_lock", false, "When --noblock_for_lock is passed, Bazel does not wait for a running command to complete, but instead exits immediately.")
	rootCmd.Flags().Bool("client_debug", false, "If true, log debug information from the client to stderr. Changing this option will not cause the server to restart.")
	rootCmd.Flags().String("command_port", "", "Port to start up the gRPC command server on. If 0, let the kernel choose.")
	rootCmd.Flags().String("connect_timeout_secs", "", "The amount of time the client waits for each attempt to connect to the server")
	rootCmd.Flags().String("default_system_javabase", "", "The root of the user's local JDK install, to be used as the default target javabase and as a fall-back host_javabase. This is not the embedded JDK.")
	rootCmd.Flags().String("digest_function", "", "The hash function to use when computing file digests.")
	rootCmd.Flags().String("experimental_cgroup_parent", "", "<path>The cgroup where to start the bazel server as an absolute path. The server process will be started in the specified cgroup for each supported controller. For example, if the value of this flag is /build/bazel and the cpu and memory controllers are mounted respectively on /sys/fs/cgroup/cpu and /sys/fs/cgroup/memory, the server will be started in the cgroups /sys/fs/cgroup/cpu/build/bazel and /sys/fs/cgroup/memory/build/bazel.It is not an error if the specified cgroup is not writable for one or more of the controllers. This options does not have any effect on platforms that do not support cgroups.")
	rootCmd.Flags().Bool("experimental_run_in_user_cgroup", false, "If true, the Bazel server will be run with systemd-run, and the user will own the cgroup. This flag only takes effect on Linux.")
	rootCmd.Flags().String("extra_classpath", "", "A colon-separated list of classpath entries to be added to the classpath of the Bazel server.")
	rootCmd.Flags().String("failure_detail_out", "", "<path>If set, specifies a location to write a failure_detail protobuf message if the server experiences a failure and cannot report it via gRPC, as normal. Otherwise, the location will be ${OUTPUT_BASE}/failure_detail.rawproto.")
	rootCmd.Flags().Bool("fatal_event_bus_exceptions", false, "Whether or not to exit if an exception is thrown by an internal EventBus handler. No-op if --fatal_async_exceptions_exclusions is available; that flag's behavior is preferentially used.")
	rootCmd.Flags().Bool("home_rc", false, "Whether or not to look for the home bazelrc file at `$HOME/.bazelrc`")
	rootCmd.Flags().StringSlice("host_jvm_args", nil, "<jvm_arg>Flags to pass to the JVM executing Blaze.")
	rootCmd.Flags().String("host_jvm_debug", "", "Convenience option to add some additional JVM startup flags, which cause the JVM to wait during startup until you connect from a JDWP-compliant debugger (like Eclipse) to port 5005.")
	rootCmd.Flags().Bool("idle_server_tasks", false, "Run System.gc() when the server is idle")
	rootCmd.Flags().Bool("ignore_all_rc_files", false, "Disables all rc files, regardless of the values of other rc-modifying flags, even if these flags come later in the list of startup options.")
	rootCmd.Flags().String("install_base", "", "This launcher option is intended for use only by tests.")
	rootCmd.Flags().String("install_md5", "", "This launcher option is intended for use only by tests.")
	rootCmd.Flags().String("invocation_policy", "", "A base64-encoded-binary-serialized or text-formated invocation_policy.InvocationPolicy proto. Unlike other options, it is an error to specify --invocation_policy multiple times.")
	rootCmd.Flags().String("io_nice_level", "", "{-1,0,1,2,3,4,5,6,7}Only on Linux; set a level from 0-7 for best-effort IO scheduling using the sys_ioprio_set system call. 0 is highest priority, 7 is lowest. The anticipatory scheduler may only honor up to priority 4. If set to a negative value, then Bazel does not perform a system call.")
	rootCmd.Flags().String("local_startup_timeout_secs", "", "The maximum amount of time the client waits to connect to the server")
	rootCmd.Flags().Bool("lock_install_base", false, "Whether the server should hold a lock on the install base while running, to prevent another server from attempting to garbage collect it.")
	rootCmd.Flags().String("macos_qos_class", "", "Sets the QoS service class of the %{product} server when running on macOS. This flag has no effect on all other platforms but is supported to ensure rc files can be shared among them without changes. Possible values are: user-interactive, user-initiated, default, utility, and background.")
	rootCmd.Flags().String("max_idle_secs", "", "<integer>The number of seconds the build server will wait idling before shutting down. Zero means that the server will never shutdown. This is only read on server-startup, changing this option will not cause the server to restart.")
	rootCmd.Flags().Bool("noautodetect_server_javabase", false, "When --noautodetect_server_javabase is passed, Bazel does not fall back to the local JDK for running the bazel server and instead exits.")
	rootCmd.Flags().Bool("nobatch", false, "If set, Bazel will be run as just a client process without a server, instead of in the standard client/server mode. This is deprecated and will be removed, please prefer shutting down the server explicitly if you wish to avoid lingering servers.")
	rootCmd.Flags().Bool("nobatch_cpu_scheduling", false, "Only on Linux; use 'batch' CPU scheduling for Blaze. This policy is useful for workloads that are non-interactive, but do not want to lower their nice value. See 'man 2 sched_setscheduler'. If false, then Bazel does not perform a system call.")
	rootCmd.Flags().Bool("noblock_for_lock", false, "When --noblock_for_lock is passed, Bazel does not wait for a running command to complete, but instead exits immediately.")
	rootCmd.Flags().Bool("noclient_debug", false, "If true, log debug information from the client to stderr. Changing this option will not cause the server to restart.")
	rootCmd.Flags().Bool("noexperimental_run_in_user_cgroup", false, "If true, the Bazel server will be run with systemd-run, and the user will own the cgroup. This flag only takes effect on Linux.")
	rootCmd.Flags().Bool("nofatal_event_bus_exceptions", false, "Whether or not to exit if an exception is thrown by an internal EventBus handler. No-op if --fatal_async_exceptions_exclusions is available; that flag's behavior is preferentially used.")
	rootCmd.Flags().Bool("nohome_rc", false, "Whether or not to look for the home bazelrc file at `$HOME/.bazelrc`")
	rootCmd.Flags().Bool("noidle_server_tasks", false, "Run System.gc() when the server is idle")
	rootCmd.Flags().Bool("noignore_all_rc_files", false, "Disables all rc files, regardless of the values of other rc-modifying flags, even if these flags come later in the list of startup options.")
	rootCmd.Flags().Bool("nolock_install_base", false, "Whether the server should hold a lock on the install base while running, to prevent another server from attempting to garbage collect it.")
	rootCmd.Flags().Bool("nopreemptible", false, "If true, the command can be preempted if another command is started.")
	rootCmd.Flags().Bool("noquiet", false, "If true, no informational messages are emitted on the console, only errors. Changing this option will not cause the server to restart.")
	rootCmd.Flags().Bool("noshutdown_on_low_sys_mem", false, "If max_idle_secs is set and the build server has been idle for a while, shut down the server when the system is low on free RAM. Linux and MacOS only.")
	rootCmd.Flags().Bool("nosystem_rc", false, "Whether or not to look for the system-wide bazelrc.")
	rootCmd.Flags().Bool("nounlimit_coredumps", false, "Raises the soft coredump limit to the hard limit to make coredumps of the server (including the JVM) and the client possible under common conditions. Stick this flag in your bazelrc once and forget about it so that you get coredumps when you actually encounter a condition that triggers them.")
	rootCmd.Flags().Bool("nowindows_enable_symlinks", false, "If true, real symbolic links will be created on Windows instead of file copying. Requires Windows developer mode to be enabled and Windows 10 version 1703 or greater.")
	rootCmd.Flags().Bool("noworkspace_rc", false, "Whether or not to look for the workspace bazelrc file at `$workspace/.bazelrc`")
	rootCmd.Flags().Bool("nowrite_command_log", false, "WARNING: This option is deprecated and will be removed soon. Please use the command option instead.")
	rootCmd.Flags().String("option_sources", "", "")
	rootCmd.Flags().String("output_base", "", "<path>If set, specifies the output location to which all build output will be written. Otherwise, the location will be ${OUTPUT_ROOT}/_blaze_${USER}/${MD5_OF_WORKSPACE_ROOT}. Note: If you specify a different option from one to the next Bazel invocation for this value, you'll likely start up a new, additional Bazel server. Bazel starts exactly one server per specified output base. Typically there is one output base per workspace - however, with this option you may have multiple output bases per workspace and thereby run multiple builds for the same client on the same machine concurrently. See 'bazel help shutdown' on how to shutdown a Bazel server.")
	rootCmd.Flags().String("output_user_root", "", "<path>The user-specific directory beneath which all build outputs are written; by default, this is a function of $USER, but by specifying a constant, build outputs can be shared between collaborating users.")
	rootCmd.Flags().Bool("preemptible", false, "If true, the command can be preempted if another command is started.")
	rootCmd.Flags().String("product_name", "", "The name of the build system. It is used as part of the name of the generated directories (e.g. productName-bin for binaries) as well as for printing error messages and logging")
	rootCmd.Flags().Bool("quiet", false, "If true, no informational messages are emitted on the console, only errors. Changing this option will not cause the server to restart.")
	rootCmd.Flags().String("server_javabase", "", "<jvm path>Path to the JVM used to execute Bazel itself.")
	rootCmd.Flags().String("server_jvm_out", "", "<path>The location to write the server's JVM's output. If unset then defaults to a location in output_base.")
	rootCmd.Flags().Bool("shutdown_on_low_sys_mem", false, "If max_idle_secs is set and the build server has been idle for a while, shut down the server when the system is low on free RAM. Linux and MacOS only.")
	rootCmd.Flags().Bool("system_rc", false, "Whether or not to look for the system-wide bazelrc.")
	rootCmd.Flags().String("unix_digest_hash_attribute_name", "", "The name of an extended attribute that can be placed on files to store a precomputed copy of the file's hash, corresponding with --digest_function. This option can be used to reduce disk I/O and CPU load caused by hash computation. This extended attribute is checked on all source files and output files, meaning that it causes a significant number of invocations of the getxattr() system call.")
	rootCmd.Flags().Bool("unlimit_coredumps", false, "Raises the soft coredump limit to the hard limit to make coredumps of the server (including the JVM) and the client possible under common conditions. Stick this flag in your bazelrc once and forget about it so that you get coredumps when you actually encounter a condition that triggers them.")
	rootCmd.Flags().Bool("windows_enable_symlinks", false, "If true, real symbolic links will be created on Windows instead of file copying. Requires Windows developer mode to be enabled and Windows 10 version 1703 or greater.")
	rootCmd.Flags().String("workspace_directory", "", "The root of the workspace, that is, the directory that Bazel uses as the root of the build. This flag is only to be set by the bazel client.")
	rootCmd.Flags().Bool("workspace_rc", false, "Whether or not to look for the workspace bazelrc file at `$workspace/.bazelrc`")
	rootCmd.Flags().Bool("write_command_log", false, "WARNING: This option is deprecated and will be removed soon. Please use the command option instead.")
}
