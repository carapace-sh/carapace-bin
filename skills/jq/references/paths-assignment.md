# Path Expressions and Assignment

Path expressions identify locations within a JSON document. The assignment operators (`=`, `|=`, `+=`, etc.) use path expressions on the left-hand side to modify values. jq values are immutable — assignments produce new values rather than mutating in place.

> **Source of truth**: <https://jqlang.github.io/jq/manual/#assignment-operators> and the path functions in [Builtin operators and functions](https://jqlang.github.io/jq/manual/#builtin-operators-and-functions). For **basic path syntax** (`.foo`, `.[]`, `.[0]`), see [filters.md](filters.md).

## Path Expressions

Path expressions are jq expressions that refer to a location within `.`. They include `.a`, `.a.b`, `.[0]`, `.[]`, `.a[].b`, etc.

There are two kinds:
- **Exact match** — `.a.b.c` matches at most one location
- **Pattern** — `.a[].b` can match multiple locations

Any jq expression can be a path expression when used with `path()`, `del()`, or assignment operators.

## Path Functions

### `path(path_expression)`

Outputs array representations of the paths matched by the expression. Arrays contain strings (object keys) and/or numbers (array indices).

```jq
path(.a[0].b)              # → ["a",0,"b"] (even if path doesn't exist in input)
path(..)                   # all paths in the document
path(.. | select(type=="boolean"))  # all paths to boolean values
```

`path(exact_path_expression)` produces the array form even if the path doesn't exist in `.`, as long as `.` is `null`, an array, or an object.

| Program | Input | Output |
|---------|-------|--------|
| `path(.a[0].b)` | `null` | `["a",0,"b"]` |
| `[path(..)]` | `{"a":[{"b":1}]}` | `[[],["a"],["a",0],["a",0,"b"]]` |

### `getpath(PATHS)`

Outputs the values at each path in `PATHS`. `PATHS` is an array of path arrays.

| Program | Input | Output |
|---------|-------|--------|
| `getpath(["a","b"])` | `null` | `null` |
| `[getpath(["a","b"], ["a","c"])]` | `{"a":{"b":0, "c":1}}` | `[0, 1]` |

### `setpath(PATHS; VALUE)`

Sets the path in `.` to `VALUE`, returning the modified value.

| Program | Input | Output |
|---------|-------|--------|
| `setpath(["a","b"]; 1)` | `null` | `{"a": {"b": 1}}` |
| `setpath(["a","b"]; 1)` | `{"a":{"b":0}}` | `{"a": {"b": 1}}` |
| `setpath([0,"a"]; 1)` | `null` | `[{"a":1}]` |

### `delpaths(PATHS)`

Deletes multiple paths. `PATHS` must be an array of path arrays.

| Program | Input | Output |
|---------|-------|--------|
| `delpaths([["a","b"]])` | `{"a":{"b":1},"x":{"y":2}}` | `{"a":{},"x":{"y":2}}` |

### `del(path_expression)`

Removes keys from objects or elements from arrays.

| Program | Input | Output |
|---------|-------|--------|
| `del(.foo)` | `{"foo": 42, "bar": 9001, "baz": 42}` | `{"bar": 9001, "baz": 42}` |
| `del(.[1, 2])` | `["foo", "bar", "baz"]` | `["foo"]` |

### `pick(pathexps)`

Projection — emit only the specified paths from the input, with unmatched paths set to `null`. For arrays, negative indices and `.[m:n]` should not be used.

The invariant: `(. | p)` equals `(. | pick(pathexps) | p)` for any path `p` in the specification.

| Program | Input | Output |
|---------|-------|--------|
| `pick(.a, .b.c, .x)` | `{"a": 1, "b": {"c": 2, "d": 3}, "e": 4}` | `{"a":1,"b":{"c":2},"x":null}` |
| `pick(.[2], .[0], .[0])` | `[1,2,3,4]` | `[1,null,3]` |

## Assignment Operators

jq values are **immutable**. Assignment works by computing a new replacement value for `.` with all desired assignments applied, then outputting it. The original value is not changed:

```jq
{a:{b:{c:1}}} | (.a.b|=3), .    # outputs {"a":{"b":3}} then {"a":{"b":{"c":1}}}
```

All assignment operators have **path expressions** on the left-hand side (LHS). The right-hand side (RHS) provides values to set.

### Important Notes

- The LHS refers to a value **in `.`** — `$var.foo = 1` won't work; use `$var | .foo = 1`
- `.a,.b=0` does **not** set both — it's `(.a, .b=0)` which sets `.b` and outputs `.a`. Use `(.a,.b)=0` to set both.

### Update-Assignment: `|=`

The "update" operator. Takes a filter on the RHS. The old value at the LHS path is fed through the RHS filter to produce the new value.

```jq
(.foo, .bar) |= .+1     # increment both .foo and .bar
```

- LHS can be any path expression (including generators like `.[]`)
- If RHS outputs `empty`, the path is deleted (like `del(path)`)
- If RHS outputs multiple values, only the **first** is used (changed in jq 1.6; jq 1.5 used the last)

| Program | Input | Output |
|---------|-------|--------|
| `(..|select(type=="boolean")) |= if . then 1 else 0 end` | `[true,false,[5,true,[true,[false]],false]]` | `[1,0,[5,1,[1,[0]],0]]` |

### Arithmetic Update-Assignment: `+=`, `-=`, `*=`, `/=`, `%=`, `//=`

All equivalent to `a |= . op b`:

```jq
.foo += 1      # same as .foo |= . + 1
.foo -= 1      # same as .foo |= . - 1
.foo *= 2      # same as .foo |= . * 2
```

| Program | Input | Output |
|---------|-------|--------|
| `.foo += 1` | `{"foo": 42}` | `{"foo": 43}` |

### Plain Assignment: `=`

Unlike `|=`, the input to the RHS is the **same as the input to the LHS** (not the value at the LHS path). All values output by the RHS are used.

If the RHS produces multiple values, jq outputs a modified `.` for each:

```jq
(.a,.b) = range(3)    # outputs {"a":0,"b":0}, {"a":1,"b":1}, {"a":2,"b":2}
```

### `=` vs `|=` — The Key Difference

Input: `{"a": {"b": 10}, "b": 20}`

| Program | Output | Why |
|---------|--------|-----|
| `.a = .b` | `{"a":20,"b":20}` | RHS input is the whole input, so `.b` is `20` |
| `.a |= .b` | `{"a":10,"b":20}` | RHS input is the value at `.a` (i.e. `{"b":10}`), so `.b` is `10` |

| Program | Input | Output |
|---------|-------|--------|
| `.a = .b` | `{"a": {"b": 10}, "b": 20}` | `{"a":20,"b":20}` |
| `.a |= .b` | `{"a": {"b": 10}, "b": 20}` | `{"a":10,"b":20}` |
| `(.a, .b) = range(3)` | `null` | `{"a":0,"b":0}`, `{"a":1,"b":1}`, `{"a":2,"b":2}` |
| `(.a, .b) |= range(3)` | `null` | `{"a":0,"b":0}` (only first output used) |

## Complex Assignments

The LHS of an assignment can be any path expression — including ones that produce multiple paths:

```jq
.posts[0].title = "JQ Manual"

.posts[].comments |= . + ["this is great"]    # append to comments in every post
```

### Conditional Assignment with `select`

Combine `select` with assignment to target specific elements:

```jq
(.posts[] | select(.author == "stedolan") | .comments) |= . + ["terrible."]
```

This finds all posts by "stedolan" and appends to their comments arrays. The LHS produces paths to each matching post's comments field.

## References

- [jq manual — Path functions](https://jqlang.github.io/jq/manual/#path-path_expression)
- [jq manual — Assignment operators](https://jqlang.github.io/jq/manual/#assignment-operators)
- For basic path syntax, see [filters.md](filters.md)
- For `keys`, `has`, `paths`, `add`, `any`, `all`, and other builtins, see [builtins.md](builtins.md)
