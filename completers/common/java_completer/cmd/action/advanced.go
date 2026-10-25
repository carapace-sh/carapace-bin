package action

import "github.com/carapace-sh/carapace"

func ActionAdvanced() carapace.Action {
	return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionAdvancedKeys()
		case 1:
			return ActionAdvancedValues(c.Parts[0])
		default:
			return carapace.ActionValues()
		}
	})
}

func ActionAdvancedKeys() carapace.Action {
	return carapace.ActionValuesDescribed(
		"+AggressiveHeap", "Enables Java heap optimization",
		"+AggressiveOpts", "Enables the use of aggressive performance optimization features",
		"+AllowUserSignalHandlers", "Enables installation of signal handlers by the application",
		"+AlwaysPreTouch", "Enables touching of every page on the Java heap during JVM initialization",
		"+BackgroundCompilation", "Enables background compilation",
		"+CMSClassUnloadingEnabled", "Enables class unloading when using the concurrent mark-sweep garbage collector",
		"+CMSScavengeBeforeRemark", "Enables scavenging attempts before the CMS remark step",
		"+DisableAttachMechanism", "Enables the option that disables the mechanism that lets tools attach to the JVM",
		"+DisableExplicitGC", "Enables the option that disables processing of calls to System.gc()",
		"+DoEscapeAnalysis", "Enables the use of escape analysis",
		"+ExplicitGCInvokesConcurrent", "Enables invoking of concurrent GC by using the System.gc() request",
		"+ExplicitGCInvokesConcurrentAndUnloadsClasses", "Enables invoking of concurrent GC",
		"+ExtendedDTraceProbes", "Enables additional dtrace tool probes that impact the performance",
		"+FailOverToOldVerifier", "Enables automatic failover to the old verifier when the new type checker fails",
		"+G1PrintHeapRegions", "Enables the printing of information about which regions are allocated",
		"+HeapDumpOnOutOfMemory", "Enables the dumping of the Java heap to a file in the current directory",
		"+Inline", "Enables method inlining",
		"+LogCompilation", "Enables logging of compilation activity to a file named hotspot.log in the current working directory",
		"+OptimizeStringConcat", "Enables the optimization of String concatenation operations",
		"+ParallelRefProcEnabled", "Enables parallel reference processing",
		"+PerfDataSaveToFile", "If enabled, saves jstat binary data when the Java application exits",
		"+PrintAdaptiveSizePolicy", "Enables printing of information about adaptive generation sizing",
		"+PrintAssembly", "Enables printing of assembly code for bytecoded and native methods by using the external disassembler.so library",
		"+PrintClassHistogram", "Enables printing of a class instance histogram after a Control+C event",
		"+PrintCommandLineFlags", "Enables printing of ergonomically selected JVM flags that appeared on the command line",
		"+PrintCompilation", "Enables verbose diagnostic output from the JVM by printing a message to the console every time a method is compiled",
		"+PrintConcurrentLocks", "Enables printing of locks after a event",
		"+PrintGC", "Enables printing of messages at every GC",
		"+PrintGCApplicationConcurrentTime", "Enables printing of how much time elapsed since the last pause",
		"+PrintGCApplicationStoppedTime", "Enables printing of how much time the pause lasted",
		"+PrintGCDateStamps", "Enables printing of a date stamp at every GC",
		"+PrintGCDetails", "Enables printing of detailed messages at every GC",
		"+PrintGCTaskTimeStamps", "Enables printing of time stamps for every individual GC worker thread task",
		"+PrintGCTimeStamps", "Enables printing of time stamps at every GC",
		"+PrintInlining", "Enables printing of inlining decisions",
		"+PrintNMTStatistics", "Enables printing of collected native memory tracking data at JVM exit when native memory tracking is enabled",
		"+PrintStringDeduplicationStatistics", "Prints detailed deduplication statistics",
		"+PrintTenuringDistribution", "Enables printing of tenuring age information",
		"+RelaxAccessControlCheck", "Decreases the amount of access control checks in the verifier",
		"+ScavengeBeforeFullGC", "Enables GC of the young generation before each full GC",
		"+ShowMessageBoxOnError", "Enables displaying of a dialog box when the JVM experiences an irrecoverable error",
		"+TraceClassLoading", "Enables tracing of classes as they are loaded",
		"+TraceClassLoadingPreorder", "Enables tracing of all loaded classes in the order in which they are referenced",
		"+TraceClassResolution", "Enables tracing of constant pool resolutions",
		"+TraceClassUnloading", "Enables tracing of classes as they are unloaded",
		"+TraceLoaderConstraints", "Enables tracing of the loader constraints recording",
		"+UnlockDiagnosticVMOptions", "Unlocks the options intended for diagnosing the JVM",
		"+UseAES", "Enables hardware-based AES intrinsics for Intel, AMD, and SPARC hardware",
		"+UseAESIntrinsics", "UseAES and UseAESIntrinsics flags are enabled by default and are supported only for Java HotSpot Server VM 32-bit and 64-bit",
		"+UseAdaptiveSizePolicy", "Enables the use of adaptive generation sizing",
		"+UseAltSigs", "Enables the use of alternative signals instead of SIGUSR1 and SIGUSR2 for JVM internal signals",
		"+UseCMSInitiatingOccupancyOnly", "Enables the use of the occupancy value as the only criterion for initiating the CMS collector",
		"+UseCodeCacheFlushing", "Enables flushing of the code cache before shutting down the compiler",
		"+UseConcMarkSweepGC", "Enables the use of the CMS garbage collector for the old generation",
		"+UseCondCardMark", "Enables checking of whether the card is already marked before updating the card table",
		"+UseG1GC", "Enables the use of the garbage-first garbage collector",
		"+UseGCOverheadLimit", "Enables the use of a policy that limits the proportion of time spent by the JVMon GC before an OutOfMemoryError exception is thrown",
		"+UseHugeTLBFS", "This option pre-allocates all large pages up-front",
		"+UseLargePages", "Enables the use of large page memory",
		"+UseMembar", "Enables issuing of membars on thread state transitions",
		"+UseNUMA", "Enables performance optimization of an application on a machine with nonuniform memory architecture",
		"+UseParNewGC", "Enables the use of parallel threads for collection in the young generation",
		"+UseParallelGC", "Enables the use of the parallel scavenge garbage collector",
		"+UseParallelOldGC", "Enables the use of the parallel garbage collector for full GCs",
		"+UsePerfData", "Enables the perfdata feature",
		"+UseRTMDeopt", "Auto-tunes RTM locking depending on the abort ratio",
		"+UseRTMLocking", "Generate Restricted Transactional Memory locking code for all inflated locks",
		"+UseSHA", "Enables hardware-based intrinsics for SHA crypto hash functions for SPARC hardware",
		"+UseSHA1Intrinsics", "Enables intrinsics for SHA-1 crypto hash function",
		"+UseSHA256Intrinsics", "Enables intrinsics for SHA-224 and SHA-256 crypto hash functions",
		"+UseSHA512Intrinsics", "Enables intrinsics for SHA-384 and SHA-512 crypto hash functions",
		"+UseSHM", "On Linux, enables the JVM to use shared memory to setup large pages",
		"+UseSerialGC", "Enables the use of the serial garbage collector",
		"+UseStringDeduplication", "Enables string deduplication",
		"+UseSuperWord", "Enables the transformation of scalar operations into superword operations",
		"+UseTLAB", "Enables the use of thread-local allocation blocks in the young generation space",
		"+UseTransparentHugePages", "On Linux, enables the use of large pages that can dynamically grow or shrink",
		"-TieredCompilation", "Disables the use of tiered compilation",
		"-UseBiasedLocking", "Disables the use of biased locking",
		"-UseCompressedOops", "Disables the use of compressed pointers",
		"AllocateInstancePrefetchLines=", "Sets the number of lines to prefetch ahead of the instance allocation pointer",
		"AllocatePrefetchDistance=", "Sets the size, of the prefetch distance for object allocation",
		"AllocatePrefetchInstr=", "Sets the prefetch instruction to prefetch ahead of the allocation pointer",
		"AllocatePrefetchLines=", "Sets the number of cache lines to load after the last object allocation by using the prefetch instructions generated in compiled code",
		"AllocatePrefetchStepSize=", "Sets the step size for sequential prefetch instructions",
		"AllocatePrefetchStyle=", "Sets the generated code style for prefetch instructions",
		"CICompilerCount=", "Sets the number of compiler threads to use for compilation",
		"CMSExpAvgFactor=", "Sets the percentage of time used to weight the current sample",
		"CMSInitiatingOccupancyFraction=", "Sets the percentage of the old generation occupancy at which to start a CMS collection cycle",
		"CMSTriggerRatio=", "Sets the percentage of the value specified by -XX:MinHeapFreeRatio that is allocated before a CMS collection cycle commences",
		"CodeCacheMinimumFreeSpace=", "Sets the minimum free space required for compilation",
		"CompileCommand=", "Specifies a command to perform on a method",
		"CompileCommandFile=", "Sets the file from which JIT compiler commands are read",
		"CompileOnly=", "Sets the list of methods to which compilation should be restricted",
		"CompileThreshold=", "Sets the number of interpreted method invocations before compilation",
		"ConcGCThreads=", "Sets the number of threads used for concurrent GC",
		"ErrorFile=", "Specifies the path and file name to which error data is written when an irrecoverable error occurs",
		"G1HeapRegionSize=", "Sets the size of the regions into which the Java heap is subdivided",
		"G1ReservePercent=", "Sets the percentage of the heap that is reserved",
		"HeapDumpPath=", "Sets the path and file name for writing the heap dump provided by the heap profiler",
		"InitialCodeCacheSize=", "Sets the initial code cache size",
		"InitialHeapSize=", "Sets the initial size of the memory allocation pool",
		"InitialSurvivorRatio=", "Sets the initial survivor space ratio used by the throughput garbage collector",
		"InitiatingHeapOccupancyPercent=", "Sets the percentage of the heap occupancy",
		"InlineSmallCode=", "Sets the maximum code size for compiled methods that should be inlined",
		"LargePageSizeInBytes=", "On Solaris, sets the maximum size for large pages used for Java heap",
		"LogFile=", "Sets the path and file name where log data is written",
		"MaxDirectMemorySize=", "Sets the maximum total size direct-buffer allocations",
		"MaxGCPauseMillis=", "Sets a target for the maximum GC pause time",
		"MaxHeapFreeRatio=", "Sets the maximum allowed percentage of free heap space",
		"MaxHeapSize=", "Sets the maximum size of the memory allocation pool",
		"MaxInlineSize=", "Sets the maximum bytecode size  of a method to be inlined",
		"MaxMetaspaceSize=", "Sets the maximum amount of native memory that can be allocated for class metadata",
		"MaxNewSize=", "Sets the maximum size of the heap for the young generation",
		"MaxNodeLimit=", "Sets the maximum number of nodes to be used during single method compilation",
		"MaxTenuringThreshold=", "Sets the maximum tenuring threshold for use in adaptive GC sizing",
		"MaxTrivialSize=", "Sets the maximum bytecode size of a trivial method to be inlined",
		"MetaspaceSize=", "Sets the size of the allocated class metadata space that will trigger a garbage collection",
		"MinHeapFreeRatio=", "Sets the minimum allowed percentage of free heap space after a GC event",
		"NativeMemoryTracking=", "Specifies the mode for tracking JVM native memory usage",
		"NewRatio=", "Sets the ratio between young and old generation sizes",
		"NewSize=", "Sets the initial size of the heap for the young generation",
		"ObjectAlignmentInBytes=", "Sets the memory alignment of Java objects",
		"OnError=", "Sets a custom command or a series of semicolon-separated commands to run when an irrecoverable error occurs",
		"OnOutOfMemoryError=", "Sets a custom command or a series of semicolon-separated commands to run when an OutOfMemoryError exception is first thrown",
		"ParallelGCThreads=", "Sets the number of threads used for parallel garbage collection in the young and old generations",
		"RTMAbortRatio=", "The RTM abort ratio is specified as a percentage of all executed RTM transactions",
		"RTMRetryCount=", "RTM locking code will be retried, when it is aborted or busy, the number of times specified by this option",
		"ReservedCodeCacheSize=", "Sets the maximum code cache size  for JIT-compiled code",
		"SoftRefLRUPolicyMSPerMB=", "Sets the amount of time a softly reachable object is kept active on the heap",
		"StringDeduplicationAgeThreshold=", "String objects reaching the specified age are considered candidates for deduplication",
		"SurvivorRatio=", "Sets the ratio between eden space size and survivor space size",
		"TLABSize=", "Sets the initial size",
		"TargetSurvivorRatio=", "Sets the desired percentage of survivor space used after young garbage collection",
		"ThreadStackSize=", "Sets the thread stack size",
		"nced Garbage Collection Options", "These options control how garbage collection is performed by the Java HotSpot VM.",
	)
}

func ActionAdvancedValues(key string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch key {
		case "ErrorFile":
			return carapace.ActionFiles()
		case "LargePageSizeInBytes":
			return carapace.ActionValues()
		case "MaxDirectMemorySize":
			return carapace.ActionValues()
		case "NativeMemoryTracking":
			return carapace.ActionValuesDescribed(
				"detail", "additionally track memory usage by individual CallSite",
				"off", "Do not track JVM native memory usage",
				"summary", "Only track memory usage by JVM subsystems",
			)
		case "ObjectAlignmentInBytes":
			return carapace.ActionValues()
		case "OnError":
			return carapace.ActionValues()
		case "OnOutOfMemoryError":
			return carapace.ActionValues()
		case "ThreadStackSize":
			return carapace.ActionValues()
		case "AllocateInstancePrefetchLines":
			return carapace.ActionValues()
		case "AllocatePrefetchDistance":
			return carapace.ActionValues()
		case "AllocatePrefetchInstr":
			return carapace.ActionValues()
		case "AllocatePrefetchLines":
			return carapace.ActionValues()
		case "AllocatePrefetchStepSize":
			return carapace.ActionValues()
		case "AllocatePrefetchStyle":
			return carapace.ActionValuesDescribed(
				"0", "Do not generate prefetch instructions",
				"1", "Execute prefetch instructions after each allocation",
				"2", "Use the thread-local allocation block (TLAB) watermark pointer to determine when prefetch instructions are executed",
				"3", "Use BIS instruction on SPARC for allocation prefetch",
			)
		case "CICompilerCount":
			return carapace.ActionValues()
		case "CodeCacheMinimumFreeSpace":
			return carapace.ActionValues()
		case "CompileCommand":
			return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"break", "Set a breakpoint when debugging the JVM to stop at the beginning of compilation of the specified method",
						"compileonly", "Exclude all methods from compilation except for the specified method",
						"dontinline", "Prevent inlining of the specified method",
						"exclude", "Exclude the specified method from compilation",
						"help", "Print a help message for the -XX:CompileCommand option",
						"inline", "Attempt to inline the specified method",
						"log", "Exclude compilation logging for all methods except for the specified method",
						"option", "This command can be used to pass a JIT compilation option to the specified method",
						"print", "Print generated assembler code after compilation of the specified method",
						"quiet", "Do not print the compile commands",
					).NoSpace()
				default:
					return carapace.ActionValues()
				}
			})
		case "CompileCommandFile":
			return carapace.ActionFiles()
		case "CompileOnly":
			return carapace.ActionValues()
		case "CompileThreshold":
			return carapace.ActionValues()
		case "InitialCodeCacheSize":
			return carapace.ActionValues()
		case "InlineSmallCode":
			return carapace.ActionValues()
		case "MaxInlineSize":
			return carapace.ActionValues()
		case "MaxNodeLimit":
			return carapace.ActionValues()
		case "MaxTrivialSize":
			return carapace.ActionValues()
		case "ReservedCodeCacheSize":
			return carapace.ActionValues()
		case "RTMAbortRatio":
			return carapace.ActionValues()
		case "RTMRetryCount":
			return carapace.ActionValues()
		case "HeapDumpPath":
			return carapace.ActionFiles()
		case "LogFile":
			return carapace.ActionFiles()
		case "CMSExpAvgFactor":
			return carapace.ActionValues()
		case "CMSInitiatingOccupancyFraction":
			return carapace.ActionValues()
		case "CMSTriggerRatio":
			return carapace.ActionValues()
		case "ConcGCThreads":
			return carapace.ActionValues()
		case "G1HeapRegionSize":
			return carapace.ActionValues()
		case "G1ReservePercent":
			return carapace.ActionValues()
		case "InitialHeapSize":
			return carapace.ActionValues()
		case "InitialSurvivorRatio":
			return carapace.ActionValues()
		case "InitiatingHeapOccupancyPercent":
			return carapace.ActionValues()
		case "MaxGCPauseMillis":
			return carapace.ActionValues()
		case "MaxHeapSize":
			return carapace.ActionValues()
		case "MaxHeapFreeRatio":
			return carapace.ActionValues()
		case "MaxMetaspaceSize":
			return carapace.ActionValues()
		case "MaxNewSize":
			return carapace.ActionValues()
		case "MaxTenuringThreshold":
			return carapace.ActionValues()
		case "MetaspaceSize":
			return carapace.ActionValues()
		case "MinHeapFreeRatio":
			return carapace.ActionValues()
		case "NewRatio":
			return carapace.ActionValues()
		case "NewSize":
			return carapace.ActionValues()
		case "ParallelGCThreads":
			return carapace.ActionValues()
		case "SoftRefLRUPolicyMSPerMB":
			return carapace.ActionValues()
		case "StringDeduplicationAgeThreshold":
			return carapace.ActionValues()
		case "SurvivorRatio":
			return carapace.ActionValues()
		case "TargetSurvivorRatio":
			return carapace.ActionValues()
		case "TLABSize":
			return carapace.ActionValues()
		default:
			return carapace.ActionValues()
		}
	})
}
