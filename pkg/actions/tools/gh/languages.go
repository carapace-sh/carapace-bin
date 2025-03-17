package gh

import (
	"github.com/carapace-sh/carapace"
)

// ActionLanguages completes repository languages
//
//	Go (programming)
//	Go Checksums (data)
func ActionLanguages() carapace.Action {
	// curl https://raw.githubusercontent.com/github/linguist/master/lib/linguist/languages.yml | yj -yj | jq --raw-output  'to_entries[] | "\"\(.key)\", \"\(.value.type)\","'
	return carapace.ActionValuesDescribed(
		"1C Enterprise", "programming",
		"2-Dimensional Array", "data",
		"4D", "programming",
		"ABAP", "programming",
		"ABAP CDS", "programming",
		"ABNF", "data",
		"AGS Script", "programming",
		"AIDL", "programming",
		"AL", "programming",
		"AMPL", "programming",
		"ANTLR", "programming",
		"API Blueprint", "markup",
		"APL", "programming",
		"ASL", "programming",
		"ASN.1", "data",
		"ASP.NET", "programming",
		"ATS", "programming",
		"ActionScript", "programming",
		"Ada", "programming",
		"Adblock Filter List", "data",
		"Adobe Font Metrics", "data",
		"Agda", "programming",
		"Alloy", "programming",
		"Alpine Abuild", "programming",
		"Altium Designer", "data",
		"AngelScript", "programming",
		"Answer Set Programming", "programming",
		"Ant Build System", "data",
		"Antlers", "markup",
		"ApacheConf", "data",
		"Apex", "programming",
		"Apollo Guidance Computer", "programming",
		"AppleScript", "programming",
		"Arc", "programming",
		"AsciiDoc", "prose",
		"AspectJ", "programming",
		"Assembly", "programming",
		"Astro", "markup",
		"Asymptote", "programming",
		"Augeas", "programming",
		"AutoHotkey", "programming",
		"AutoIt", "programming",
		"Avro IDL", "data",
		"Awk", "programming",
		"B4X", "programming",
		"BASIC", "programming",
		"BQN", "programming",
		"Ballerina", "programming",
		"Batchfile", "programming",
		"Beef", "programming",
		"Befunge", "programming",
		"Berry", "programming",
		"BibTeX", "markup",
		"BibTeX Style", "programming",
		"Bicep", "programming",
		"Bikeshed", "markup",
		"Bison", "programming",
		"BitBake", "programming",
		"Blade", "markup",
		"BlitzBasic", "programming",
		"BlitzMax", "programming",
		"Bluespec", "programming",
		"Bluespec BH", "programming",
		"Boo", "programming",
		"Boogie", "programming",
		"Brainfuck", "programming",
		"BrighterScript", "programming",
		"Brightscript", "programming",
		"Browserslist", "data",
		"C", "programming",
		"C#", "programming",
		"C++", "programming",
		"C-ObjDump", "data",
		"C2hs Haskell", "programming",
		"CAP CDS", "programming",
		"CIL", "data",
		"CLIPS", "programming",
		"CMake", "programming",
		"COBOL", "programming",
		"CODEOWNERS", "data",
		"COLLADA", "data",
		"CSON", "data",
		"CSS", "markup",
		"CSV", "data",
		"CUE", "programming",
		"CWeb", "programming",
		"Cabal Config", "data",
		"Caddyfile", "data",
		"Cadence", "programming",
		"Cairo", "programming",
		"Cairo Zero", "programming",
		"CameLIGO", "programming",
		"Cap'n Proto", "programming",
		"Carbon", "programming",
		"CartoCSS", "programming",
		"Ceylon", "programming",
		"Chapel", "programming",
		"Charity", "programming",
		"Checksums", "data",
		"ChucK", "programming",
		"Circom", "programming",
		"Cirru", "programming",
		"Clarion", "programming",
		"Clarity", "programming",
		"Classic ASP", "programming",
		"Clean", "programming",
		"Click", "programming",
		"Clojure", "programming",
		"Closure Templates", "markup",
		"Cloud Firestore Security Rules", "data",
		"CoNLL-U", "data",
		"CodeQL", "programming",
		"CoffeeScript", "programming",
		"ColdFusion", "programming",
		"ColdFusion CFC", "programming",
		"Common Lisp", "programming",
		"Common Workflow Language", "programming",
		"Component Pascal", "programming",
		"Cool", "programming",
		"Coq", "programming",
		"Cpp-ObjDump", "data",
		"Creole", "prose",
		"Crystal", "programming",
		"Csound", "programming",
		"Csound Document", "programming",
		"Csound Score", "programming",
		"Cuda", "programming",
		"Cue Sheet", "data",
		"Curry", "programming",
		"Cycript", "programming",
		"Cylc", "data",
		"Cypher", "programming",
		"Cython", "programming",
		"D", "programming",
		"D-ObjDump", "data",
		"D2", "markup",
		"DIGITAL Command Language", "programming",
		"DM", "programming",
		"DNS Zone", "data",
		"DTrace", "programming",
		"Dafny", "programming",
		"Darcs Patch", "data",
		"Dart", "programming",
		"DataWeave", "programming",
		"Debian Package Control File", "data",
		"DenizenScript", "programming",
		"Dhall", "programming",
		"Diff", "data",
		"DirectX 3D File", "data",
		"Dockerfile", "programming",
		"Dogescript", "programming",
		"Dotenv", "data",
		"Dune", "programming",
		"Dylan", "programming",
		"E", "programming",
		"E-mail", "data",
		"EBNF", "data",
		"ECL", "programming",
		"ECLiPSe", "programming",
		"EJS", "markup",
		"EQ", "programming",
		"Eagle", "data",
		"Earthly", "programming",
		"Easybuild", "data",
		"Ecere Projects", "data",
		"Ecmarkup", "markup",
		"Edge", "markup",
		"EdgeQL", "programming",
		"EditorConfig", "data",
		"Edje Data Collection", "data",
		"Eiffel", "programming",
		"Elixir", "programming",
		"Elm", "programming",
		"Elvish", "programming",
		"Elvish Transcript", "programming",
		"Emacs Lisp", "programming",
		"EmberScript", "programming",
		"Erlang", "programming",
		"Euphoria", "programming",
		"F#", "programming",
		"F*", "programming",
		"FIGlet Font", "data",
		"FIRRTL", "programming",
		"FLUX", "programming",
		"Factor", "programming",
		"Fancy", "programming",
		"Fantom", "programming",
		"Faust", "programming",
		"Fennel", "programming",
		"Filebench WML", "programming",
		"Filterscript", "programming",
		"Fluent", "programming",
		"Formatted", "data",
		"Forth", "programming",
		"Fortran", "programming",
		"Fortran Free Form", "programming",
		"FreeBASIC", "programming",
		"FreeMarker", "programming",
		"Frege", "programming",
		"Futhark", "programming",
		"G-code", "programming",
		"GAML", "programming",
		"GAMS", "programming",
		"GAP", "programming",
		"GCC Machine Description", "programming",
		"GDB", "programming",
		"GDScript", "programming",
		"GEDCOM", "data",
		"GLSL", "programming",
		"GN", "data",
		"GSC", "programming",
		"Game Maker Language", "programming",
		"Gemfile.lock", "data",
		"Gemini", "prose",
		"Genero 4gl", "programming",
		"Genero per", "markup",
		"Genie", "programming",
		"Genshi", "programming",
		"Gentoo Ebuild", "programming",
		"Gentoo Eclass", "programming",
		"Gerber Image", "data",
		"Gettext Catalog", "prose",
		"Gherkin", "programming",
		"Git Attributes", "data",
		"Git Config", "data",
		"Git Revision List", "data",
		"Gleam", "programming",
		"Glimmer JS", "programming",
		"Glimmer TS", "programming",
		"Glyph", "programming",
		"Glyph Bitmap Distribution Format", "data",
		"Gnuplot", "programming",
		"Go", "programming",
		"Go Checksums", "data",
		"Go Module", "data",
		"Go Workspace", "data",
		"Godot Resource", "data",
		"Golo", "programming",
		"Gosu", "programming",
		"Grace", "programming",
		"Gradle", "data",
		"Gradle Kotlin DSL", "data",
		"Grammatical Framework", "programming",
		"Graph Modeling Language", "data",
		"GraphQL", "data",
		"Graphviz (DOT)", "data",
		"Groovy", "programming",
		"Groovy Server Pages", "programming",
		"HAProxy", "data",
		"HCL", "programming",
		"HLSL", "programming",
		"HOCON", "data",
		"HTML", "markup",
		"HTML+ECR", "markup",
		"HTML+EEX", "markup",
		"HTML+ERB", "markup",
		"HTML+PHP", "markup",
		"HTML+Razor", "markup",
		"HTTP", "data",
		"HXML", "data",
		"Hack", "programming",
		"Haml", "markup",
		"Handlebars", "markup",
		"Harbour", "programming",
		"Hare", "programming",
		"Haskell", "programming",
		"Haxe", "programming",
		"HiveQL", "programming",
		"HolyC", "programming",
		"Hosts File", "data",
		"Hy", "programming",
		"HyPhy", "programming",
		"IDL", "programming",
		"IGOR Pro", "programming",
		"INI", "data",
		"IRC log", "data",
		"ISPC", "programming",
		"Idris", "programming",
		"Ignore List", "data",
		"ImageJ Macro", "programming",
		"Imba", "programming",
		"Inform 7", "programming",
		"Ink", "programming",
		"Inno Setup", "programming",
		"Io", "programming",
		"Ioke", "programming",
		"Isabelle", "programming",
		"Isabelle ROOT", "programming",
		"J", "programming",
		"JAR Manifest", "data",
		"JCL", "programming",
		"JFlex", "programming",
		"JSON", "data",
		"JSON with Comments", "data",
		"JSON5", "data",
		"JSONLD", "data",
		"JSONiq", "programming",
		"Jai", "programming",
		"Janet", "programming",
		"Jasmin", "programming",
		"Java", "programming",
		"Java Properties", "data",
		"Java Server Pages", "programming",
		"Java Template Engine", "programming",
		"JavaScript", "programming",
		"JavaScript+ERB", "programming",
		"Jest Snapshot", "data",
		"JetBrains MPS", "programming",
		"Jinja", "markup",
		"Jison", "programming",
		"Jison Lex", "programming",
		"Jolie", "programming",
		"Jsonnet", "programming",
		"Julia", "programming",
		"Julia REPL", "programming",
		"Jupyter Notebook", "markup",
		"Just", "programming",
		"KDL", "data",
		"KRL", "programming",
		"Kaitai Struct", "programming",
		"KakouneScript", "programming",
		"KerboScript", "programming",
		"KiCad Layout", "data",
		"KiCad Legacy Layout", "data",
		"KiCad Schematic", "data",
		"Kickstart", "data",
		"Kit", "markup",
		"Kotlin", "programming",
		"Kusto", "data",
		"LFE", "programming",
		"LLVM", "programming",
		"LOLCODE", "programming",
		"LSL", "programming",
		"LTspice Symbol", "data",
		"LabVIEW", "programming",
		"Lark", "data",
		"Lasso", "programming",
		"Latte", "markup",
		"Lean", "programming",
		"Lean 4", "programming",
		"Less", "markup",
		"Lex", "programming",
		"LigoLANG", "programming",
		"LilyPond", "programming",
		"Limbo", "programming",
		"Linear Programming", "programming",
		"Linker Script", "programming",
		"Linux Kernel Module", "data",
		"Liquid", "markup",
		"Literate Agda", "programming",
		"Literate CoffeeScript", "programming",
		"Literate Haskell", "programming",
		"LiveCode Script", "programming",
		"LiveScript", "programming",
		"Logos", "programming",
		"Logtalk", "programming",
		"LookML", "programming",
		"LoomScript", "programming",
		"Lua", "programming",
		"Luau", "programming",
		"M", "programming",
		"M4", "programming",
		"M4Sugar", "programming",
		"MATLAB", "programming",
		"MAXScript", "programming",
		"MDX", "markup",
		"MLIR", "programming",
		"MQL4", "programming",
		"MQL5", "programming",
		"MTML", "markup",
		"MUF", "programming",
		"Macaulay2", "programming",
		"Makefile", "programming",
		"Mako", "programming",
		"Markdown", "prose",
		"Marko", "markup",
		"Mask", "markup",
		"Mathematica", "programming",
		"Maven POM", "data",
		"Max", "programming",
		"Mercury", "programming",
		"Mermaid", "markup",
		"Meson", "programming",
		"Metal", "programming",
		"Microsoft Developer Studio Project", "data",
		"Microsoft Visual Studio Solution", "data",
		"MiniD", "programming",
		"MiniYAML", "data",
		"MiniZinc", "programming",
		"MiniZinc Data", "data",
		"Mint", "programming",
		"Mirah", "programming",
		"Modelica", "programming",
		"Modula-2", "programming",
		"Modula-3", "programming",
		"Module Management System", "programming",
		"Mojo", "programming",
		"Monkey", "programming",
		"Monkey C", "programming",
		"Moocode", "programming",
		"MoonBit", "programming",
		"MoonScript", "programming",
		"Motoko", "programming",
		"Motorola 68K Assembly", "programming",
		"Move", "programming",
		"Muse", "prose",
		"Mustache", "markup",
		"Myghty", "programming",
		"NASL", "programming",
		"NCL", "programming",
		"NEON", "data",
		"NL", "data",
		"NMODL", "programming",
		"NPM Config", "data",
		"NSIS", "programming",
		"NWScript", "programming",
		"Nasal", "programming",
		"Nearley", "programming",
		"Nemerle", "programming",
		"NetLinx", "programming",
		"NetLinx+ERB", "programming",
		"NetLogo", "programming",
		"NewLisp", "programming",
		"Nextflow", "programming",
		"Nginx", "data",
		"Nim", "programming",
		"Ninja", "data",
		"Nit", "programming",
		"Nix", "programming",
		"Noir", "programming",
		"Nu", "programming",
		"NumPy", "programming",
		"Nunjucks", "markup",
		"Nushell", "programming",
		"OASv2-json", "data",
		"OASv2-yaml", "data",
		"OASv3-json", "data",
		"OASv3-yaml", "data",
		"OCaml", "programming",
		"OMNeT++ MSG", "programming",
		"OMNeT++ NED", "programming",
		"Oberon", "programming",
		"ObjDump", "data",
		"Object Data Instance Notation", "data",
		"ObjectScript", "programming",
		"Objective-C", "programming",
		"Objective-C++", "programming",
		"Objective-J", "programming",
		"Odin", "programming",
		"Omgrofl", "programming",
		"Opa", "programming",
		"Opal", "programming",
		"Open Policy Agent", "programming",
		"OpenAPI Specification v2", "data",
		"OpenAPI Specification v3", "data",
		"OpenCL", "programming",
		"OpenEdge ABL", "programming",
		"OpenQASM", "programming",
		"OpenRC runscript", "programming",
		"OpenSCAD", "programming",
		"OpenStep Property List", "data",
		"OpenType Feature File", "data",
		"Option List", "data",
		"Org", "prose",
		"OverpassQL", "programming",
		"Ox", "programming",
		"Oxygene", "programming",
		"Oz", "programming",
		"P4", "programming",
		"PDDL", "programming",
		"PEG.js", "programming",
		"PHP", "programming",
		"PLSQL", "programming",
		"PLpgSQL", "programming",
		"POV-Ray SDL", "programming",
		"Pact", "programming",
		"Pan", "programming",
		"Papyrus", "programming",
		"Parrot", "programming",
		"Parrot Assembly", "programming",
		"Parrot Internal Representation", "programming",
		"Pascal", "programming",
		"Pawn", "programming",
		"Pep8", "programming",
		"Perl", "programming",
		"Pic", "markup",
		"Pickle", "data",
		"PicoLisp", "programming",
		"PigLatin", "programming",
		"Pike", "programming",
		"Pip Requirements", "data",
		"Pkl", "programming",
		"PlantUML", "data",
		"Pod", "prose",
		"Pod 6", "prose",
		"PogoScript", "programming",
		"Polar", "programming",
		"Pony", "programming",
		"Portugol", "programming",
		"PostCSS", "markup",
		"PostScript", "markup",
		"PowerBuilder", "programming",
		"PowerShell", "programming",
		"Praat", "programming",
		"Prisma", "data",
		"Processing", "programming",
		"Procfile", "programming",
		"Proguard", "data",
		"Prolog", "programming",
		"Promela", "programming",
		"Propeller Spin", "programming",
		"Protocol Buffer", "data",
		"Protocol Buffer Text Format", "data",
		"Public Key", "data",
		"Pug", "markup",
		"Puppet", "programming",
		"Pure Data", "data",
		"PureBasic", "programming",
		"PureScript", "programming",
		"Pyret", "programming",
		"Python", "programming",
		"Python console", "programming",
		"Python traceback", "data",
		"Q#", "programming",
		"QML", "programming",
		"QMake", "programming",
		"Qt Script", "programming",
		"Quake", "programming",
		"QuickBASIC", "programming",
		"R", "programming",
		"RAML", "markup",
		"RBS", "data",
		"RDoc", "prose",
		"REALbasic", "programming",
		"REXX", "programming",
		"RMarkdown", "prose",
		"RON", "data",
		"RPC", "programming",
		"RPGLE", "programming",
		"RPM Spec", "data",
		"RUNOFF", "markup",
		"Racket", "programming",
		"Ragel", "programming",
		"Raku", "programming",
		"Rascal", "programming",
		"Raw token data", "data",
		"ReScript", "programming",
		"Readline Config", "data",
		"Reason", "programming",
		"ReasonLIGO", "programming",
		"Rebol", "programming",
		"Record Jar", "data",
		"Red", "programming",
		"Redcode", "programming",
		"Redirect Rules", "data",
		"Regular Expression", "data",
		"Ren'Py", "programming",
		"RenderScript", "programming",
		"Rez", "programming",
		"Rich Text Format", "markup",
		"Ring", "programming",
		"Riot", "markup",
		"RobotFramework", "programming",
		"Roc", "programming",
		"Roff", "markup",
		"Roff Manpage", "markup",
		"Rouge", "programming",
		"RouterOS Script", "programming",
		"Ruby", "programming",
		"Rust", "programming",
		"SAS", "programming",
		"SCSS", "markup",
		"SELinux Policy", "data",
		"SMT", "programming",
		"SPARQL", "data",
		"SQF", "programming",
		"SQL", "data",
		"SQLPL", "programming",
		"SRecode Template", "markup",
		"SSH Config", "data",
		"STAR", "data",
		"STL", "data",
		"STON", "data",
		"SVG", "data",
		"SWIG", "programming",
		"Sage", "programming",
		"SaltStack", "programming",
		"Sass", "markup",
		"Scala", "programming",
		"Scaml", "markup",
		"Scenic", "programming",
		"Scheme", "programming",
		"Scilab", "programming",
		"Self", "programming",
		"ShaderLab", "programming",
		"Shell", "programming",
		"ShellCheck Config", "data",
		"ShellSession", "programming",
		"Shen", "programming",
		"Sieve", "programming",
		"Simple File Verification", "data",
		"Singularity", "programming",
		"Slang", "programming",
		"Slash", "programming",
		"Slice", "programming",
		"Slim", "markup",
		"Slint", "markup",
		"SmPL", "programming",
		"Smali", "programming",
		"Smalltalk", "programming",
		"Smarty", "programming",
		"Smithy", "programming",
		"Snakemake", "programming",
		"Solidity", "programming",
		"Soong", "data",
		"SourcePawn", "programming",
		"Spline Font Database", "data",
		"Squirrel", "programming",
		"Stan", "programming",
		"Standard ML", "programming",
		"Starlark", "programming",
		"Stata", "programming",
		"StringTemplate", "markup",
		"Stylus", "markup",
		"SubRip Text", "data",
		"SugarSS", "markup",
		"SuperCollider", "programming",
		"Survex data", "data",
		"Svelte", "markup",
		"Sway", "programming",
		"Sweave", "prose",
		"Swift", "programming",
		"SystemVerilog", "programming",
		"TI Program", "programming",
		"TL-Verilog", "programming",
		"TLA", "programming",
		"TOML", "data",
		"TSPLIB data", "data",
		"TSQL", "programming",
		"TSV", "data",
		"TSX", "programming",
		"TXL", "programming",
		"Tact", "programming",
		"Talon", "programming",
		"Tcl", "programming",
		"Tcsh", "programming",
		"TeX", "markup",
		"Tea", "markup",
		"Terra", "programming",
		"Terraform Template", "markup",
		"Texinfo", "prose",
		"Text", "prose",
		"TextGrid", "data",
		"TextMate Properties", "data",
		"Textile", "prose",
		"Thrift", "programming",
		"Toit", "programming",
		"Tree-sitter Query", "programming",
		"Turing", "programming",
		"Turtle", "data",
		"Twig", "markup",
		"Type Language", "data",
		"TypeScript", "programming",
		"TypeSpec", "programming",
		"Typst", "programming",
		"Unified Parallel C", "programming",
		"Unity3D Asset", "data",
		"Unix Assembly", "programming",
		"Uno", "programming",
		"UnrealScript", "programming",
		"UrWeb", "programming",
		"V", "programming",
		"VBA", "programming",
		"VBScript", "programming",
		"VCL", "programming",
		"VHDL", "programming",
		"Vala", "programming",
		"Valve Data Format", "data",
		"Velocity Template Language", "markup",
		"Verilog", "programming",
		"Vim Help File", "prose",
		"Vim Script", "programming",
		"Vim Snippet", "markup",
		"Visual Basic .NET", "programming",
		"Visual Basic 6.0", "programming",
		"Volt", "programming",
		"Vue", "markup",
		"Vyper", "programming",
		"WDL", "programming",
		"WGSL", "programming",
		"Wavefront Material", "data",
		"Wavefront Object", "data",
		"Web Ontology Language", "data",
		"WebAssembly", "programming",
		"WebAssembly Interface Type", "data",
		"WebIDL", "programming",
		"WebVTT", "data",
		"Wget Config", "data",
		"Whiley", "programming",
		"Wikitext", "prose",
		"Win32 Message File", "data",
		"Windows Registry Entries", "data",
		"Witcher Script", "programming",
		"Wollok", "programming",
		"World of Warcraft Addon Data", "data",
		"Wren", "programming",
		"X BitMap", "data",
		"X Font Directory Index", "data",
		"X PixMap", "data",
		"X10", "programming",
		"XC", "programming",
		"XCompose", "data",
		"XML", "data",
		"XML Property List", "data",
		"XPages", "data",
		"XProc", "programming",
		"XQuery", "programming",
		"XS", "programming",
		"XSLT", "programming",
		"Xmake", "programming",
		"Xojo", "programming",
		"Xonsh", "programming",
		"Xtend", "programming",
		"YAML", "data",
		"YANG", "data",
		"YARA", "programming",
		"YASnippet", "markup",
		"Yacc", "programming",
		"Yul", "programming",
		"ZAP", "programming",
		"ZIL", "programming",
		"Zeek", "programming",
		"ZenScript", "programming",
		"Zephir", "programming",
		"Zig", "programming",
		"Zimpl", "programming",
		"cURL Config", "data",
		"crontab", "data",
		"desktop", "data",
		"dircolors", "data",
		"eC", "programming",
		"edn", "data",
		"fish", "programming",
		"hoon", "programming",
		"iCalendar", "data",
		"jq", "programming",
		"kvlang", "markup",
		"mIRC Script", "programming",
		"mcfunction", "programming",
		"mdsvex", "markup",
		"mupad", "programming",
		"nanorc", "data",
		"nesC", "programming",
		"ooc", "programming",
		"q", "programming",
		"reStructuredText", "prose",
		"robots.txt", "data",
		"sed", "programming",
		"templ", "markup",
		"vCard", "data",
		"wisp", "programming",
		"xBase", "programming",
	)
}
