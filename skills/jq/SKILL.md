---
name: jq
description: >
  Use when working with the jq command-line JSON processor — filters, builtins,
  operators, path expressions, assignments, variables, functions, regex, format
  strings, dates, modules, invocation flags, and exit codes. Triggers on:
  "jq", "jq filter", "jq program", "jq query", "jq language", "jq manual",
  "jq builtin", "jq regex", "jq module", "jq reduce", "jq foreach",
  "jq select", "jq walk", "jq recurse", "jq map", "jq inputs", "jq slurp",
  "jq --arg", "jq --argjson", "jq --null-input", "jq --raw-output",
  "jq @base64", "jq @csv", "jq @sh", "jq @uri", "jq tojson", "jq fromjson",
  "jq todate", "jq fromdate", "jq strptime", "jq strftime", "jq now",
  "jq $ENV", "jq env", "jq getpath", "jq setpath", "jq del", "jq paths",
  "jq sort", "jq group_by", "jq sub", "jq gsub", "jq test", "jq match",
  "jq with_entries", "jq map_values", "jq modulemeta", "jq import",
  "jq input", "jq --stream", "jq --exit-status",
  "jq --compact-output", "jq //", "jq destructuring",
  "jq label", "jq break", "jq try", "jq catch", "jq INDEX", "jq JOIN".
user-invocable: true
---

# jq — Command-line JSON Processor In-Depth Reference

Comprehensive reference for the jq language and CLI — filters, builtin functions, operators, path expressions, assignments, variables, functions, regex, format strings, dates, modules, and invocation. Covers jq 1.7. Source: <https://jqlang.github.io/jq/manual/>.

## Data Flow

```
Input (stdin / files / --null-input)
  → JSON parser (stream of whitespace-separated values)
    → filter program (pipeline of filters, one input → zero or more outputs)
      → output formatter (--compact-output, --raw-output, --tab, --indent, --sort-keys)
        → stdout (newline-separated JSON texts)
```

A jq program is a single filter expression. Every filter takes one input and produces zero or more outputs. Filters compose with `|` (pipe), `,` (comma), and operators. Multiple outputs from one stage are fed one-by-one into the next — this is the generator model at the heart of jq.

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| invocation, command line, flags, options, --null-input, -n, --raw-input, -R, --slurp, -s, --compact-output, -c, --raw-output, -r, --raw-output0, --join-output, -j, --ascii-output, -a, --sort-keys, -S, --color-output, -C, --monochrome-output, -M, --tab, --indent, --unbuffered, --stream, --stream-errors, --seq, --from-file, -f, -L, --library-path, --arg, --argjson, --slurpfile, --rawfile, --args, --jsonargs, --exit-status, -e, --binary, -b, --version, -V, --build-configuration, --help, -h, --run-tests, --, exit status, exit code, JQ_COLORS, NO_COLOR, $ENV, env, $ARGS, shell quoting, stdin, files | [references/cli.md](references/cli.md) |
| filter, identity, ., .foo, .foo.bar, .foo?, optional, object index, .["foo"], array index, .[0], negative index, slice, .[2:4], .[:3], .[-2:], value iterator, .[], .[]?, comma, pipe, |, parenthesis, recursive descent, .., basic filter | [references/filters.md](references/filters.md) |
| type, value, number, string, boolean, null, array, object, array construction, [], object construction, {}, shortcut syntax, {user, title}, key expression, variable key, {$foo}, IEEE754, double precision, number literal, precision | [references/types.md](references/types.md) |
| addition, +, subtraction, -, multiplication, *, division, /, modulo, %, abs, length, utf8bytelength, arithmetic, array concatenation, string concatenation, object merge, null addition, comparison, ==, !=, >, >=, <=, <, and, or, not, truthiness, alternative operator, //, default value, sort order | [references/operators.md](references/operators.md) |
| builtin, function, keys, keys_unsorted, has, in, map, map_values, pick, select, arrays, objects, iterables, booleans, numbers, normals, finites, strings, nulls, values, scalars, empty, error, halt, halt_error, $__loc__, paths, add, any, all, flatten, range, floor, sqrt, tonumber, tostring, type, infinite, nan, isnan, isinfinite, isfinite, isnormal, sort, sort_by, group_by, min, max, min_by, max_by, unique, unique_by, reverse, contains, indices, index, rindex, inside, startswith, endswith, combinations, ltrimstr, rtrimstr, explode, implode, split, join, ascii_downcase, ascii_upcase, while, repeat, until, recurse, walk, transpose, bsearch, builtins, getpath, setpath, delpaths, to_entries, from_entries, with_entries, path, del | [references/builtins.md](references/builtins.md) |
| conditional, if, then, else, elif, end, try, catch, ?, error suppression, optional operator, reduce, foreach, label, break, breaking out, control structure, generators, iterator, tail call, recursion | [references/control.md](references/control.md) |
| variable, binding, as, $name, destructuring, pattern, ?//, alternative, def, function definition, function argument, scoping, lexical scope, filter argument, callback, value argument, addvalue, isempty, limit, skip, first, last, nth | [references/variables-functions.md](references/variables-functions.md) |
| path, path expression, path(f), getpath, setpath, delpaths, del, pick, assignment, =, update assignment, |=, arithmetic update, +=, -=, *=, /=, %=, //=, complex assignment, LHS, RHS, immutable, deep copy, posts, comments | [references/paths-assignment.md](references/paths-assignment.md) |
| string, interpolation, \(exp), tojson, fromjson, tostring, tonumber, length, utf8bytelength, split, join, ltrimstr, rtrimstr, startswith, endswith, explode, implode, ascii_downcase, ascii_upcase, ascii, test, match, capture, scan, sub, gsub, splits | [references/strings.md](references/strings.md) |
| regex, regular expression, Oniguruma, Perl NG, test, match, capture, scan, split, splits, sub, gsub, flags, g, i, m, n, p, s, l, x, named group, capturing group, offset, length | [references/regex.md](references/regex.md) |
| format string, @text, @json, @html, @uri, @csv, @tsv, @sh, @base64, @base64d, escaping, percent-encoding, CSV, TSV, shell escaping, HTML entity, string interpolation, @foo | [references/formats.md](references/formats.md) |
| date, time, fromdate, todate, fromdateiso8601, todateiso8601, now, strptime, strftime, strflocaltime, mktime, gmtime, localtime, broken down time, Unix epoch, ISO 8601, UTC, timezone | [references/dates.md](references/dates.md) |
| module, import, include, module directive, modulemeta, .jq, search path, -L, $ORIGIN, ~, metadata, deps, defs, I/O, input, inputs, input_filename, input_line_number, debug, stderr, INDEX, JOIN, IN, SQL, $ENV, env, $JQ_BUILD_CONFIGURATION | [references/modules.md](references/modules.md) |

## Quick Guide

- **How do I invoke jq and what do the flags do?** → [references/cli.md](references/cli.md)
- **How do I write a basic filter (`.`, `.foo`, `.[]`, `|`)?** → [references/filters.md](references/filters.md)
- **What types does jq support and how do I construct arrays/objects?** → [references/types.md](references/types.md)
- **How do arithmetic, comparison, and boolean operators work?** → [references/operators.md](references/operators.md)
- **What builtin functions are available (map, select, add, sort, etc.)?** → [references/builtins.md](references/builtins.md)
- **How do `if`/`try`/`reduce`/`foreach` and `label`/`break` work?** → [references/control.md](references/control.md)
- **How do I bind variables (`as $x`) and define functions (`def`)?** → [references/variables-functions.md](references/variables-functions.md)
- **How do path expressions and assignment operators (`=`, `|=`, `+=`) work?** → [references/paths-assignment.md](references/paths-assignment.md)
- **What string functions exist (split, join, interpolation, etc.)?** → [references/strings.md](references/strings.md)
- **How do jq regular expressions work (test, match, sub, gsub, flags)?** → [references/regex.md](references/regex.md)
- **What are the `@foo` format strings and how do they escape?** → [references/formats.md](references/formats.md)
- **How do I parse and format dates in jq?** → [references/dates.md](references/dates.md)
- **How do modules, I/O builtins, and SQL-style operators work?** → [references/modules.md](references/modules.md)
- **How do I pass variables from the command line (`--arg`, `--argjson`)?** → [references/cli.md](references/cli.md)
- **What exit codes does jq use?** → [references/cli.md](references/cli.md)
- **How do I read input from files or use `--null-input`?** → [references/cli.md](references/cli.md)
- **How do I configure output formatting (compact, raw, tab, indent)?** → [references/cli.md](references/cli.md)
- **How does `--stream` mode work?** → [references/cli.md](references/cli.md)
- **How do I debug a jq program (`debug`, `stderr`)?** → [references/modules.md](references/modules.md)
- **What is the `//` alternative operator?** → [references/operators.md](references/operators.md)
- **How does destructuring with `?//` work?** → [references/variables-functions.md](references/variables-functions.md)
- **How do I break out of a `reduce` or `foreach`?** → [references/control.md](references/control.md)

## Cross-Project References

- For carapace-specific jq integration (the `jq` completer in `completers/common/jq_completer/`), see the **carapace-dev** skill → `references/action.md`. The completer itself only covers CLI flags and file/directory arguments — it does not complete the jq filter language.
- For shell quoting rules that affect how jq programs are passed on the command line, see the **bash**, **zsh**, or **fish** skills as appropriate for the target shell.
