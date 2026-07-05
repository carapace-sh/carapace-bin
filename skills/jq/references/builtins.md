# Builtin Functions

The core jq standard library — functions for iteration, selection, aggregation, sorting, searching, and data transformation. Many of these are themselves written in jq (see [variables-functions.md](variables-functions.md) for `def`).

> **Source of truth**: <https://jqlang.github.io/jq/manual/#builtin-operators-and-functions>. For **string-specific functions**, see [strings.md](strings.md). For **regex functions**, see [regex.md](regex.md). For **date functions**, see [dates.md](dates.md). For **path functions** (`path`, `getpath`, `setpath`, `del`, `delpaths`, `pick`), see [paths-assignment.md](paths-assignment.md).

## Keys and Values

### `keys`, `keys_unsorted`

`keys` returns an array of an object's keys, sorted by Unicode codepoint. For arrays, returns `[0, 1, ..., length-1]`. `keys_unsorted` preserves insertion order for objects.

| Program | Input | Output |
|---------|-------|--------|
| `keys` | `{"abc": 1, "abcd": 2, "Foo": 3}` | `["Foo", "abc", "abcd"]` |
| `keys` | `[42,3,35]` | `[0,1,2]` |

### `has(key)`, `in`

`has($key)` — true if the input object has the key, or the input array has an element at that index. `in` is the inverse — the input is the key/index, the argument is the container.

| Program | Input | Output |
|---------|-------|--------|
| `map(has("foo"))` | `[{"foo": 42}, {}]` | `[true, false]` |
| `map(has(2))` | `[[0,1], ["a","b","c"]]` | `[false, true]` |
| `.[] | in({"foo": 42})` | `["foo", "bar"]` | `true`, `false` |
| `map(in([0,1]))` | `[2, 0]` | `[false, true]` |

## Mapping and Transformation

### `map(f)`, `map_values(f)`

`map(f)` applies `f` to each value in an array/object and collects results into an array. `map_values(f)` preserves the input type (array→array, object→object) and uses only the first output of `f` per value.

- `map(f)` = `[.[] | f]` — always outputs an array
- `map_values(f)` = `.[] |= f` — outputs same type as input

```jq
map(.+1)              # [1,2,3] → [2,3,4]
map_values(.+1)       # {"a":1,"b":2} → {"a":2,"b":3}
map(., .)             # [1] → [1,1] (multiple outputs collected)
map_values(., .)      # [1] → [1] (only first output used)
```

| Program | Input | Output |
|---------|-------|--------|
| `map(.+1)` | `[1,2,3]` | `[2,3,4]` |
| `map_values(.+1)` | `{"a": 1, "b": 2, "c": 3}` | `{"a": 2, "b": 3, "c": 4}` |
| `map(., .)` | `[1,2]` | `[1,1,2,2]` |
| `map_values(. // empty)` | `{"a": null, "b": true, "c": false}` | `{"b":true}` |

### `select(boolean_expression)`

Produces input unchanged if the expression returns true; produces nothing otherwise. Essential for filtering.

```jq
[1,2,3] | map(select(. >= 2))    # → [2,3]
```

| Program | Input | Output |
|---------|-------|--------|
| `map(select(. >= 2))` | `[1,5,3,0,7]` | `[5,3,7]` |
| `.[] | select(.id == "second")` | `[{"id": "first", "val": 1}, {"id": "second", "val": 2}]` | `{"id": "second", "val": 2}` |

### `to_entries`, `from_entries`, `with_entries(f)`

Convert between objects and arrays of key-value pairs.

| Function | Input | Output |
|----------|-------|--------|
| `to_entries` | `{"a": 1, "b": 2}` | `[{"key":"a", "value":1}, {"key":"b", "value":2}]` |
| `from_entries` | `[{"key":"a", "value":1}, {"key":"b", "value":2}]` | `{"a": 1, "b": 2}` |
| `with_entries(f)` | object | `to_entries | map(f) | from_entries` |

`from_entries` accepts keys named `"key"`, `"Key"`, `"name"`, `"Name"`, `"value"`, and `"Value"`.

```jq
with_entries(.key |= "KEY_" + .)    # prefix all keys
```

| Program | Input | Output |
|---------|-------|--------|
| `to_entries` | `{"a": 1, "b": 2}` | `[{"key":"a", "value":1}, {"key":"b", "value":2}]` |
| `from_entries` | `[{"key":"a", "value":1}, {"key":"b", "value":2}]` | `{"a": 1, "b": 2}` |
| `with_entries(.key |= "KEY_" + .)` | `{"a": 1, "b": 2}` | `{"KEY_a": 1, "KEY_b": 2}` |

## Aggregation

### `add`

Takes an array, adds elements together. Rules same as `+` operator (sum for numbers, concatenate for strings/arrays, merge for objects). Empty array → `null`.

| Program | Input | Output |
|---------|-------|--------|
| `add` | `["a","b","c"]` | `"abc"` |
| `add` | `[1, 2, 3]` | `6` |
| `add` | `[]` | `null` |

### `any`, `any(condition)`, `any(generator; condition)`

True if any element is true. Empty array → `false`.

| Form | Behavior |
|------|----------|
| `any` | True if any element of input array is true |
| `any(condition)` | Applies condition to each element |
| `any(generator; condition)` | Applies condition to outputs of generator |

| Program | Input | Output |
|---------|-------|--------|
| `any` | `[true, false]` | `true` |
| `any` | `[false, false]` | `false` |
| `any` | `[]` | `false` |

### `all`, `all(condition)`, `all(generator; condition)`

True if all elements are true. Empty array → `true`.

| Program | Input | Output |
|---------|-------|--------|
| `all` | `[true, false]` | `false` |
| `all` | `[true, true]` | `true` |
| `all` | `[]` | `true` |

### `flatten`, `flatten(depth)`

Recursively flatten nested arrays. Optional depth limits levels.

| Program | Input | Output |
|---------|-------|--------|
| `flatten` | `[1, [2], [[3]]]` | `[1, 2, 3]` |
| `flatten(1)` | `[1, [2], [[3]]]` | `[1, 2, [3]]` |
| `flatten` | `[[]]` | `[]` |

## Sorting and Grouping

### `sort`, `sort_by(path_expression)`

`sort` sorts the input array using jq's default ordering (see [operators.md](operators.md)). `sort_by(f)` sorts by the result of applying `f` to each element. `sort_by` can take multiple arguments for tie-breaking.

| Program | Input | Output |
|---------|-------|--------|
| `sort` | `[8,3,null,6]` | `[null,3,6,8]` |
| `sort_by(.foo)` | `[{"foo":4}, {"foo":3}, {"foo":2}]` | `[{"foo":2}, {"foo":3}, {"foo":4}]` |
| `sort_by(.foo, .bar)` | `[{"foo":4,"bar":10}, {"foo":3,"bar":20}, {"foo":2,"bar":1}, {"foo":3,"bar":10}]` | `[{"foo":2,"bar":1}, {"foo":3,"bar":10}, {"foo":3,"bar":20}, {"foo":4,"bar":10}]` |

### `group_by(path_expression)`

Groups elements by the result of the expression, sorted by group key.

| Program | Input | Output |
|---------|-------|--------|
| `group_by(.foo)` | `[{"foo":1, "bar":10}, {"foo":3, "bar":100}, {"foo":1, "bar":1}]` | `[[{"foo":1, "bar":10}, {"foo":1, "bar":1}], [{"foo":3, "bar":100}]]` |

### `min`, `max`, `min_by(path_exp)`, `max_by(path_exp)`

Find the minimum/maximum element of an array.

| Program | Input | Output |
|---------|-------|--------|
| `min` | `[5,4,2,7]` | `2` |
| `max_by(.foo)` | `[{"foo":1, "bar":14}, {"foo":2, "bar":3}]` | `{"foo":2, "bar":3}` |

### `unique`, `unique_by(path_exp)`

Remove duplicates, sorted. `unique_by(f)` keeps one element per group defined by `f`.

| Program | Input | Output |
|---------|-------|--------|
| `unique` | `[1,2,5,3,5,3,1,3]` | `[1,2,3,5]` |
| `unique_by(.foo)` | `[{"foo":1,"bar":2}, {"foo":1,"bar":3}, {"foo":4,"bar":5}]` | `[{"foo":1,"bar":2}, {"foo":4,"bar":5}]` |
| `unique_by(length)` | `["chunky", "bacon", "kitten", "cicada", "asparagus"]` | `["bacon", "chunky", "asparagus"]` |

### `reverse`

Reverses an array.

| Program | Input | Output |
|---------|-------|--------|
| `reverse` | `[1,2,3,4]` | `[4,3,2,1]` |

## Searching and Containment

### `contains(element)`

True if the argument is completely contained in the input:

| Type | Containment |
|------|-------------|
| strings | substring |
| arrays | all elements of argument contained in input |
| objects | all key-value pairs of argument present in input |
| other | equality |

| Program | Input | Output |
|---------|-------|--------|
| `contains("bar")` | `"foobar"` | `true` |
| `contains(["baz", "bar"])` | `["foobar", "foobaz", "blarp"]` | `true` |
| `contains({foo: 12, bar: [{barp: 12}]})` | `{"foo": 12, "bar":[1,2,{"barp":12, "blip":13}]}` | `true` |

### `inside`

Inverse of `contains` — true if the input is contained within the argument.

### `indices(s)`, `index(s)`, `rindex(s)`

`indices(s)` — array of all positions where `s` occurs. `index(s)` — first occurrence. `rindex(s)` — last occurrence.

| Program | Input | Output |
|---------|-------|--------|
| `indices(", ")` | `"a,b, cd, efg, hijk"` | `[3,7,12]` |
| `indices(1)` | `[0,1,2,1,3,1,4]` | `[1,3,5]` |
| `index(", ")` | `"a,b, cd, efg, hijk"` | `3` |
| `rindex(1)` | `[0,1,2,1,3,1,4]` | `5` |

## Iteration and Generation

### `range(upto)`, `range(from; upto)`, `range(from; upto; by)`

Generate numbers as separate outputs. Use `[range(...)]` to collect into an array.

| Program | Input | Output |
|---------|-------|--------|
| `range(2; 4)` | `null` | `2`, `3` |
| `[range(4)]` | `null` | `[0,1,2,3]` |
| `[range(0; 10; 3)]` | `null` | `[0,3,6,9]` |
| `[range(0; 10; -1)]` | `null` | `[]` |
| `[range(0; -5; -1)]` | `null` | `[0,-1,-2,-3,-4]` |

### `recurse(f)`, `recurse`, `recurse(f; condition)`

Search through recursive structures. `recurse` (no arg) = `recurse(.[]?)` = `..`.

```jq
recurse(.children[]) | .name      # all filenames in a tree
recurse(. * .; . < 20)            # 2 → 2, 4, 16
```

| Program | Input | Output |
|---------|-------|--------|
| `recurse(.foo[])` | `{"foo":[{"foo": []}, {"foo":[{"foo":[]}]}]}` | *(4 outputs, all nested objects)* |
| `recurse` | `{"a":0,"b":[1]}` | `{"a":0,"b":[1]}`, `0`, `[1]`, `1` |
| `recurse(. * .; . < 20)` | `2` | `2`, `4`, `16` |

`recurse(f)` is identical to `recurse(f; true)` and is safe regarding recursion depth. Recursive calls don't consume extra memory when `f` produces at most one output per input.

### `walk(f)`

Apply `f` recursively to every component — elements of arrays first, then the array; values of objects first, then the object.

```jq
walk(if type == "array" then sort else . end)    # sort all nested arrays
walk(if type == "object" then with_entries(.key |= sub("^_+"; "")) else . end)  # strip leading _ from all keys
```

| Program | Input | Output |
|---------|-------|--------|
| `walk(if type == "array" then sort else . end)` | `[[4, 1, 7], [8, 5, 2], [3, 6, 9]]` | `[[1,4,7],[2,5,8],[3,6,9]]` |

### `combinations`, `combinations(n)`

Cartesian product of arrays. With argument `n`, all combinations of `n` repetitions.

| Program | Input | Output |
|---------|-------|--------|
| `combinations` | `[[1,2], [3, 4]]` | `[1, 3]`, `[1, 4]`, `[2, 3]`, `[2, 4]` |
| `combinations(2)` | `[0, 1]` | `[0, 0]`, `[0, 1]`, `[1, 0]`, `[1, 1]` |

## Looping Constructs

### `while(cond; update)`

Repeatedly apply `update` to `.` until `cond` is false. Emits each intermediate value.

| Program | Input | Output |
|---------|-------|--------|
| `[while(.<100; .*2)]` | `1` | `[1,2,4,8,16,32,64]` |

### `until(cond; next)`

Repeatedly apply `next` until `cond` is true. Emits only the final value.

| Program | Input | Output |
|---------|-------|--------|
| `[.,1]|until(.[0] < 1; [.[0] - 1, .[1] * .[0]])|.[1]` | `4` | `24` (factorial) |

### `repeat(exp)`

Repeatedly apply `exp` to `.` until an error is raised.

| Program | Input | Output |
|---------|-------|--------|
| `[repeat(.*2, error)?]` | `1` | `[2]` |

## Type Conversion

### `tonumber`

Parse a string as a number. Numbers pass through. Other types error.

| Program | Input | Output |
|---------|-------|--------|
| `.[] | tonumber` | `[1, "1"]` | `1`, `1` |

### `tostring`

Strings pass through. All other values are JSON-encoded.

| Program | Input | Output |
|---------|-------|--------|
| `.[] | tostring` | `[1, "1", [1]]` | `"1"`, `"1"`, `"[1]"` |

### `tojson`, `fromjson`

`tojson` serializes as JSON text (strings get quotes, unlike `tostring`). `fromjson` parses a JSON string into a value.

| Program | Input | Output |
|---------|-------|--------|
| `[.[]|tostring]` | `[1, "foo", ["foo"]]` | `["1","foo","[\"foo\"]"]` |
| `[.[]|tojson]` | `[1, "foo", ["foo"]]` | `["1","\"foo\"","[\"foo\"]"]` |
| `[.[]|tojson|fromjson]` | `[1, "foo", ["foo"]]` | `[1,"foo",["foo"]]` |

### `type`

Returns the type name as a string: `"null"`, `"boolean"`, `"number"`, `"string"`, `"array"`, or `"object"`.

| Program | Input | Output |
|---------|-------|--------|
| `map(type)` | `[0, false, [], {}, null, "hello"]` | `["number", "boolean", "array", "object", "null", "string"]` |

## Math Functions

| Function | Description |
|----------|-------------|
| `floor` | Floor of numeric input |
| `sqrt` | Square root |
| `fabs` | Floating-point absolute value |
| `abs` | Absolute value (see [operators.md](operators.md)) |

| Program | Input | Output |
|---------|-------|--------|
| `floor` | `3.14159` | `3` |
| `sqrt` | `9` | `3` |

## Error and Control

### `empty`

Produces no results at all — not even `null`. Backtracks to the preceding generator.

| Program | Input | Output |
|---------|-------|--------|
| `1, empty, 2` | `null` | `1`, `2` |
| `[1,2,empty,3]` | `null` | `[1,2,3]` |

### `error`, `error(message)`

Produces an error with the input value as the message, or with the given argument. Caught by `try`/`catch` (see [control.md](control.md)).

```jq
try error("invalid value: \(.)") catch .     # input 42 → "invalid value: 42"
```

### `halt`, `halt_error(exit_code)`

Stop the jq program. `halt` exits with status 0. `halt_error` prints input to stderr (raw, no newline) and exits with the given code (default 5).

```jq
"Error: something went wrong\n" | halt_error(1)
```

### `$__loc__`

Produces `{"file": "<filename>", "line": <line>}` for the location where it appears.

```jq
try error("\($__loc__)") catch .    # → "{\"file\":\"<top-level>\",\"line\":1}"
```

## Utility Functions

### `paths`, `paths(node_filter)`

`paths` outputs all paths to elements in the input (excluding the empty path for `.` itself). `paths(f)` outputs only paths to values where `f` is true.

| Program | Input | Output |
|---------|-------|--------|
| `[paths]` | `[1,[[],{"a":2}]]` | `[[0],[1],[1,0],[1,1],[1,1,"a"]]` |
| `[paths(type == "number")]` | `[1,[[],{"a":2}]]` | `[[0],[1,1,"a"]]` |

### `transpose`

Transpose a jagged matrix (array of arrays). Rows padded with null to make rectangular.

| Program | Input | Output |
|---------|-------|--------|
| `transpose` | `[[1], [2,3]]` | `[[1,2],[null,3]]` |

### `bsearch(x)`

Binary search in a sorted array. Returns index if found, or `(-1 - insertion_point)` if not found.

| Program | Input | Output |
|---------|-------|--------|
| `bsearch(0)` | `[0,1]` | `0` |
| `bsearch(0)` | `[1,2,3]` | `-1` |
| `bsearch(4) as $ix | if $ix < 0 then .[-(1+$ix)] = 4 else . end` | `[1,2,3]` | `[1,2,3,4]` |

### `builtins`

Returns a list of all builtin functions in `name/arity` format. Functions with the same name but different arities are separate entries (`all/0`, `all/1`, `all/2`).

## Generator Utilities

These help work with generators (filters that produce multiple outputs):

| Function | Description |
|----------|-------------|
| `isempty(exp)` | True if `exp` produces no outputs |
| `limit(n; expr)` | Up to `n` outputs from `expr` |
| `skip(n; expr)` | All outputs from `expr` except the first `n` |
| `first(expr)` | First output from `expr` |
| `last(expr)` | Last output from `expr` |
| `nth(n; expr)` | The nth output from `expr` (no negative `n`) |
| `first` | First element of array at `.` |
| `last` | Last element of array at `.` |
| `nth(n)` | The nth element of array at `.` |

| Program | Input | Output |
|---------|-------|--------|
| `isempty(empty)` | `null` | `true` |
| `isempty(.[])` | `[]` | `true` |
| `isempty(.[])` | `[1,2,3]` | `false` |
| `[limit(3; .[])]` | `[0,1,2,3,4,5,6,7,8,9]` | `[0,1,2]` |
| `[skip(3; .[])]` | `[0,1,2,3,4,5,6,7,8,9]` | `[3,4,5,6,7,8,9]` |
| `[first(range(.)), last(range(.)), nth(5; range(.))]` | `10` | `[0,9,5]` |
| `[range(.)]|[first, last, nth(5)]` | `10` | `[0,9,5]` |

## References

- [jq manual — Builtin operators and functions](https://jqlang.github.io/jq/manual/#builtin-operators-and-functions)
- For string functions, see [strings.md](strings.md)
- For regex functions, see [regex.md](regex.md)
- For date functions, see [dates.md](dates.md)
- For path functions, see [paths-assignment.md](paths-assignment.md)
- For defining your own functions, see [variables-functions.md](variables-functions.md)
