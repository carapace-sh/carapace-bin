# Control Structures

Conditionals, error handling, reduction, iteration, and breaking out of control structures. These are the flow-control constructs that go beyond simple filter piping.

> **Source of truth**: <https://jqlang.github.io/jq/manual/#conditionals-and-comparisons> and <https://jqlang.github.io/jq/manual/#advanced-features>. For **variable bindings** used in `reduce`/`foreach`, see [variables-functions.md](variables-functions.md).

## Conditionals

### `if-then-else-end`

```jq
if A then B else C end
if A then B end              # else is optional, defaults to .
if A then B elif D then E else F end
```

- If `A` produces a value other than `false` or `null`, behaves as `B`
- Otherwise behaves as `C` (or `.` if `else` is omitted)
- If `A` produces multiple results, `B` runs for each truthy result and `C` for each falsey one

**Truthiness**: only `false` and `null` are falsey. Empty strings, `0`, and empty arrays are truthy — this differs from JavaScript/Python.

```jq
if .name == "" then A else B end    # correct way to test for empty string
# NOT: if .name then A else B end   # this is true even when .name is ""
```

| Program | Input | Output |
|---------|-------|--------|
| `if . == 0 then "zero" elif . == 1 then "one" else "many" end` | `2` | `"many"` |

## Error Handling

### `try-catch`

```jq
try EXP catch EXP            # run EXP, on error run catch with error message
try EXP                      # catch defaults to empty (suppress error)
```

The first expression runs; if it fails, the second runs with the error message as input. The output of the handler becomes the output, as if it came from the try expression.

| Program | Input | Output |
|---------|-------|--------|
| `try .a catch ". is not an object"` | `true` | `". is not an object"` |
| `[.[]|try .a]` | `[{}, true, {"a":1}]` | `[null, 1]` |
| `try error("some exception") catch .` | `true` | `"some exception"` |

### Error Suppression / Optional Operator: `?`

`EXP?` is shorthand for `try EXP` (with `empty` as the catch handler):

```jq
[.[] | .a?]                 # get .a from each, suppress errors
[.[] | tonumber?]           # parse numbers, skip invalid
```

| Program | Input | Output |
|---------|-------|--------|
| `[.[] | .a?]` | `[{}, true, {"a":1}]` | `[null, 1]` |
| `[.[] | tonumber?]` | `["1", "invalid", "3", 4]` | `[1, 3, 4]` |

## Reduction

### `reduce EXP as $var (INIT; UPDATE)`

Combine all results of an expression into a single accumulated value.

```
reduce EXP as $var (INIT; UPDATE)
```

- `INIT` is the initial accumulator value (run once on the input)
- For each result of `EXP`, run `UPDATE` with `.` as the current accumulator and `$var` as the current result
- The final value of `.` is the output

Example — sum of an array:

```jq
reduce .[] as $item (0; . + $item)
# equivalent to:
# 0 | 1 as $item | . + $item |
#     2 as $item | . + $item |
#     3 as $item | . + $item
```

| Program | Input | Output |
|---------|-------|--------|
| `reduce .[] as $item (0; . + $item)` | `[1,2,3,4,5]` | `15` |
| `reduce .[] as [$i,$j] (0; . + $i * $j)` | `[[1,2],[3,4],[5,6]]` | `44` |
| `reduce .[] as {$x,$y} (null; .x += $x | .y += [$y])` | `[{"x":"a","y":1},{"x":"b","y":2},{"x":"c","y":3}]` | `{"x":"abc","y":[1,2,3]}` |

Destructuring patterns work in the `as` binding — see [variables-functions.md](variables-functions.md).

## Foreach

### `foreach EXP as $var (INIT; UPDATE; EXTRACT)`

Like `reduce`, but emits intermediate results. Intended for reducers that produce outputs along the way.

```
foreach EXP as $var (INIT; UPDATE; EXTRACT)
```

- `INIT` is the initial accumulator
- For each result of `EXP`, run `UPDATE` to get the new accumulator
- After each update, run `EXTRACT` to produce output(s) from the accumulator
- If `EXTRACT` is omitted, the identity (`.`) is used — outputs the accumulator itself

```jq
foreach .[] as $item (0; . + $item; [$item, . * 2])
# intermediate accumulators: 1, 3, 6, 10, 15
# outputs: [1,2], [2,6], [3,12], [4,20], [5,30]
```

| Program | Input | Output |
|---------|-------|--------|
| `foreach .[] as $item (0; . + $item)` | `[1,2,3,4,5]` | `1`, `3`, `6`, `10`, `15` |
| `foreach .[] as $item (0; . + $item; [$item, . * 2])` | `[1,2,3,4,5]` | `[1,2]`, `[2,6]`, `[3,12]`, `[4,20]`, `[5,30]` |
| `foreach .[] as $item (0; . + 1; {index: ., $item})` | `["foo", "bar", "baz"]` | `{"index":1,"item":"foo"}`, `{"index":2,"item":"bar"}`, `{"index":3,"item":"baz"}` |

## Breaking Out of Control Structures

### `label` / `break`

jq has lexical labels for breaking out of `reduce`, `foreach`, `while`, and other loops:

```jq
label $out | ... break $out ...
```

`break $label` acts as though the nearest `label $label` (to the left) produced `empty`. The relationship is **lexical** — the label must be visible from the break.

```jq
# Break out of a reduce when condition is met
label $out | reduce .[] as $item (null; if .==false then break $out else ... end)
```

A `break` with no visible corresponding `label` is a syntax error:

```jq
break $out    # ERROR: no label $out visible
```

### Using `try`/`catch` to Break

A common idiom is to use `try`/`catch` with a sentinel error to break out of loops:

```jq
# Repeat until "break" error, then stop
try repeat(exp) catch if . == "break" then empty else error
```

This works because `error("break")` can be caught and the message checked.

## Recursion and Tail Calls

Any jq function can be recursive. Tail calls are optimized when the expression to the left of the recursive call outputs its last value (i.e., produces at most one output per input):

```jq
def recurse(f): def r: ., (f | select(. != null) | r); r;

def while(cond; update):
  def _while:
    if cond then ., (update | _while) else empty end;
  _while;

def repeat(exp):
  def _repeat: exp, _repeat;
  _repeat;
```

The comma operator is a generator. `empty` backtracks. Recursive generators are efficient when recursive calls are in tail position.

## Generators and Iterators

jq is fundamentally about generators. Several operators and builtins produce zero, one, or more values per input:

| Generator | Produces |
|-----------|----------|
| `.[]` | All values in array/object |
| `range(0; 10)` | Integers 0 through 9 |
| `EXP, EXP` | Left's outputs, then right's outputs |
| `empty` | Zero outputs (backtracks) |
| `recurse(f)` | `.`, then `.|f`, then `.|f|f`, ... |
| `combinations` | Cartesian product |

All jq functions can be generators by using these. New generators can be built with recursion and the comma operator.

| Program | Input | Output |
|---------|-------|--------|
| `def while(cond; update): def _while: if cond then ., (update | _while) else empty end; _while; [while(.<100; .*2)]` | `1` | `[1,2,4,8,16,32,64]` |

## References

- [jq manual — Conditionals and Comparisons](https://jqlang.github.io/jq/manual/#conditionals-and-comparisons)
- [jq manual — Advanced features](https://jqlang.github.io/jq/manual/#advanced-features)
- For `as` binding and destructuring patterns, see [variables-functions.md](variables-functions.md)
- For `error`/`halt`/`halt_error` builtins, see [builtins.md](builtins.md)
