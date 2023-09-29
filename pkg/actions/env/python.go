package env

import (
	"os"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/python"
	"github.com/rsteube/carapace-bin/pkg/conditions"
)

func init() {
	knownVariables["python"] = func() variables {
		return variables{
			Condition: conditions.ConditionPath("python"),
			Variables: map[string]string{
				"PYTHONSAFEPATH":          "prepend a potentially unsafe path to sys.path such as the current directory",
				"PYTHONHOME":              "Change  the  location  of  the  standard  Python libraries",
				"PYTHONPATH":              "Augments the default search path for module files",
				"PYTHONPLATLIBDIR":        "Override sys.platlibdir",
				"PYTHONSTARTUP":           "Python commands in this file are executed before the first prompt in interactive mode",
				"PYTHONOPTIMIZE":          "Remove assert statements and any code conditional on the value of __debug_",
				"PYTHONDEBUG":             "Turn on parser debugging output",
				"PYTHONDONTWRITEBYTECODE": "Don't write .pyc files on import",
				"PYTHONINSPECT":           "enter interactive mode after executing the script or the command",
				"PYTHONIOENCODING":        "overrides the encoding used for stdin/stdout/stderr",
				"PYTHONNOUSERSITE":        "Don't add user site directory to sys.path",
				"PYTHONUNBUFFERED":        "Force the stdout and stderr streams to be unbuffered",
				"PYTHONVERBOSE":           "Print  a  message  each  time a module is initialized",
				"PYTHONWARNINGS":          "Warning control",
				"PYTHONHASHSEED":          "value used to seed the hashes of str and bytes objects",
				"PYTHONINTMAXSTRDIGITS":   "Limit the maximum digit characters in an int value when converting from a string",
				"PYTHONMALLOC":            "Set the Python memory allocators and/or install  debug  hooks",
				"PYTHONMALLOCSTATS":       "Python will print statistics of the pymalloc memory allocator",
				"PYTHONASYNCIODEBUG":      "enable the debug mode of the asyncio module",
				"PYTHONTRACEMALLOC":       "start tracing Python memory allocations using the tracemalloc module",
				"PYTHONFAULTHANDLER":      "call faulthandler.enable() at startup",
				"PYTHONEXECUTABLE":        "set sys.argv[0] to this value instead of the value got through the C runtime",
				"PYTHONUSERBASE":          "user base directory",
				"PYTHONPROFILEIMPORTTIME": "show how long each import takes",
				"PYTHONBREAKPOINT":        "set debugger callable",
				"PYTHONTHREADDEBUG":       "print threading  debug info",
				"PYTHONDUMPREFS":          "dump objects and reference counts still alive after shutting down the interpreter",
			},
			VariableCompletion: map[string]carapace.Action{
				"PYTHONSAFEPATH":   carapace.ActionDirectories(),
				"PYTHONHOME":       carapace.ActionDirectories(),
				"PYTHONPATH":       carapace.ActionDirectories().List(string(os.PathListSeparator)),
				"PYTHONPLATLIBDIR": carapace.ActionDirectories(),
				"PYTHONSTARTUP":    carapace.ActionValues(), // TODO
				"PYTHONOPTIMIZE": carapace.ActionValuesDescribed(
					"1", "Remove assert statements and any code conditional on the value of __debug_",
					"2", "Like '1' but also discard docstrings",
				),
				"PYTHONDEBUG":             carapace.ActionValues("1"),
				"PYTHONDONTWRITEBYTECODE": carapace.ActionValues("1"),
				"PYTHONINSPECT":           carapace.ActionValues("1"),
				"PYTHONIOENCODING":        carapace.ActionValues(), // TODO
				"PYTHONNOUSERSITE":        carapace.ActionValues("1"),
				"PYTHONUNBUFFERED":        carapace.ActionValues("1"),
				"PYTHONVERBOSE":           carapace.ActionValues("1", "2"),
				"PYTHONWARNINGS":          python.ActionWarningControls().List(","),
				"PYTHONHASHSEED":          carapace.ActionValues(),
				"PYTHONINTMAXSTRDIGITS":   carapace.ActionValues(),
				"PYTHONMALLOC":            carapace.ActionValues("malloc", "pymalloc", "debug", "malloc_debug", "pymalloc_debug"), // TODO verify
				"PYTHONMALLOCSTATS":       carapace.ActionValues("1"),
				"PYTHONASYNCIODEBUG":      carapace.ActionValues("1"),
				"PYTHONTRACEMALLOC":       carapace.ActionValues("1", "5", "10", "25", "50", "100"),
				"PYTHONFAULTHANDLER":      carapace.ActionValues("1"),
				"PYTHONEXECUTABLE":        carapace.ActionValues(),
				"PYTHONUSERBASE":          carapace.ActionDirectories(),
				"PYTHONPROFILEIMPORTTIME": carapace.ActionValues("1"),
				"PYTHONBREAKPOINT":        carapace.ActionValuesDescribed("0", "disable"), // TODO debuggers
				"PYTHONTHREADDEBUG":       carapace.ActionValues("1"),
				"PYTHONDUMPREFS":          carapace.ActionValues("1"),
			},
		}
	}
}
