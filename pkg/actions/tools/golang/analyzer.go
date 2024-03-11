package golang

import "github.com/carapace-sh/carapace"

// ActionAnalyzers completes analyzers
//
//	assign (check for useless assignments)
//	atomic (check for common mistakes using the sync/atomic package)
func ActionAnalyzers() carapace.Action {
	return carapace.ActionValuesDescribed(
		"asmdecl", "report mismatches between assembly files and Go declarations",
		"assign", "check for useless assignments",
		"atomic", "check for common mistakes using the sync/atomic package",
		"bools", "check for common mistakes involving boolean operators",
		"buildtag", "check //go:build and // +build directives",
		"cgocall", "detect some violations of the cgo pointer passing rules",
		"composites", "check for unkeyed composite literals",
		"copylocks", "check for locks erroneously passed by value",
		"directive", "check Go toolchain directives such as //go:debug",
		"errorsas", "report passing non-pointer or non-error values to errors.As",
		"framepointer", "report assembly that clobbers the frame pointer before saving it",
		"httpresponse", "check for mistakes using HTTP responses",
		"ifaceassert", "detect impossible interface-to-interface type assertions",
		"loopclosure", "check references to loop variables from within nested functions",
		"lostcancel", "check cancel func returned by context.WithCancel is called",
		"nilfunc", "check for useless comparisons between functions and nil",
		"printf", "check consistency of Printf format strings and arguments",
		"shift", "check for shifts that equal or exceed the width of the integer",
		"sigchanyzer", "check for unbuffered channel of os.Signal",
		"slog", "check for invalid structured logging calls",
		"stdmethods", "check signature of methods of well-known interfaces",
		"stringintconv", "check for string(int) conversions",
		"structtag", "check that struct field tags conform to reflect.StructTag.Get",
		"testinggoroutine", "report calls to (*testing.T).Fatal from goroutines started by a test.",
		"tests", "check for common mistaken usages of tests and examples",
		"timeformat", "check for calls of (time.Time).Format or time.Parse with 2006-02-01",
		"unmarshal", "report passing non-pointer or non-interface values to unmarshal",
		"unreachable", "check for unreachable code",
		"unsafeptr", "check for invalid conversions of uintptr to unsafe.Pointer",
		"unusedresult", "check for unused results of calls to some functions",
	)
}
