# Control Flow and Exceptions

Ruby's control flow structures — conditionals, loops, iteration, break/next/redo/retry, and the `begin`/`rescue`/`ensure` exception system.

> **Source of truth**: <https://ruby-doc.org/core/doc/syntax/control_expressions_rdoc.html> and <https://ruby-doc.org/core/doc/syntax/exceptions_rdoc.html>. For **blocks and iterators**, see [blocks-procs-lambdas.md](blocks-procs-lambdas.md).

## Conditionals

### if / elsif / else

```ruby
if score >= 90
  "A"
elsif score >= 80
  "B"
else
  "F"
end
```

`if` is an expression — it returns a value. Use it inline:

```ruby
grade = if score >= 90 then "A" else "F" end

# Modifier form
puts "high" if score > 90
```

### unless

The opposite of `if`. Use `unless` only when there's no `else`:

```ruby
log_unless_silent unless quiet?
```

### case / when

`case` works with `===` (the case-equality operator):

```ruby
case obj
when String
  "string"
when Integer
  "integer"
when 1..10
  "small range"
when /^foo/
  "starts with foo"
else
  "unknown"
end
```

- `===` is overridden by each type: `Class#===` checks `is_a?`, `Range#===` checks `cover?`, `Regexp#===` checks `=~`
- No `break` needed — `case` doesn't fall through like C
- Use `case` with no value to chain boolean conditions (acts like `if/elsif`)

### Ternary

```ruby
score >= 60 ? "pass" : "fail"
```

## Loops and Iteration

### while / until

```ruby
while line = gets
  puts line
end

until queue.empty?
  process(queue.shift)
end
```

Both have modifier forms: `do_work while pending?`.

### for ... in

```ruby
for item in collection
  puts item
end
```

`for` calls `.each` internally and **does not create a new scope** — the loop variable leaks. Prefer `each` with a block:

```ruby
collection.each { |item| puts item }  # item is block-scoped
```

### loop

```ruby
loop do
  line = gets
  break if line.nil?
end
```

`loop` is a method (not a keyword) that takes a block. It has no condition — use `break` to exit.

### break / next / redo / retry

| Keyword | Behavior |
|---------|---------|
| `break [value]` | Exits the loop; `loop` returns the value |
| `next [value]` | Skips to next iteration; `map` uses the value |
| `redo` | Restarts the current iteration without re-checking condition |
| `retry` | Only valid inside `rescue` — re-runs the `begin` block from the top |

```ruby
result = [1, 2, 3, 4, 5].map do |n|
  next n if n.even?      # returns n for even
  n * 10                  # returns n*10 for odd
end
# => [10, 2, 30, 4, 50]
```

### catch / throw

Non-exception-based jump out of nested blocks:

```ruby
catch(:done) do
  items.each do |item|
    process(item)
    throw :done if item.last?
  end
end
```

`throw` with a second arg: `throw(:done, value)` — `catch` returns the value.

## Exceptions

### raise

```ruby
raise StandardError.new("something went wrong")
raise "simple message"              # raises RuntimeError
raise StandardError, "with message"
raise                              # re-raises $! (current exception)
```

`raise` creates and raises an exception object. Without an argument, re-raises the current exception (only valid inside `rescue`).

### begin / rescue / ensure / else

```ruby
begin
  risky_operation
rescue NetworkError => e
  log "network: #{e.message}"
rescue => e
  log "generic: #{e.message}"
else
  log "success"
ensure
  cleanup
end
```

| Clause | When executed |
|--------|---------------|
| `rescue Type => e` | When an exception of `Type` (or subclass) is raised |
| `rescue => e` | Shorthand for `rescue StandardError => e` |
| `else` | When no exception was raised (not executed if `rescue` ran) |
| `ensure` | Always (even if `rescue` re-raises or returns) |

**Default rescue class**: `StandardError`. `rescue` without a class catches `StandardError` and subclasses. It does **not** catch `SystemExit`, `Interrupt`, `NoMemoryError`, `ScriptError`, etc. — those are outside `StandardError`.

### rescue modifier

```ruby
value = risky_call rescue "default"
```

The rescue modifier catches `StandardError`. If `risky_call` raises, the expression evaluates to `"default"`.

### Exception Hierarchy

```
Exception
  ├── NoMemoryError
  ├── ScriptError
  │     ├── NotImplementedError
  │     └── SyntaxError
  ├── SecurityError
  ├── SignalException
  │     └── Interrupt
  ├── StandardError     ← default rescue catches this
  │     ├── ArgumentError
  │     ├── EncodingError
  │     ├── FiberError
  │     ├── IOError
  │     ├── IndexError
  │     │     └── StopIteration
  │     ├── NameError
  │     │     └── NoMethodError
  │     ├── RangeError
  │     │     └── FloatDomainError
  │     ├── RegexpError
  │     ├── RuntimeError
  │     ├── SystemCallError
  │     │     └── Errno::*
  │     ├── ThreadError
  │     ├── TypeError
  │     └── ZeroDivisionError
  ├── SystemExit
  └── SystemStackError
```

**Create custom exceptions** by subclassing `StandardError`:

```ruby
class MyError < StandardError; end
class ValidationError < StandardError
  attr_reader :field
  def initialize(field, message)
    @field = field
    super(message)
  end
end
```

### Global Variables

| Variable | Content |
|----------|---------|
| `$!` | The current exception (last raised) |
| `$@` | The backtrace array of `$!` |
| `$ERROR_INFO` | `$!` (from `English` library) |
| `$ERROR_POSITION` | `$@` (from `English` library) |

### retry

```ruby
attempts = 0
begin
  risky_call
rescue
  attempts += 1
  retry if attempts < 3
  raise
end
```

`retry` re-runs the entire `begin` block. Use with a counter to avoid infinite loops.

## Exception Best Practices

1. **Rescue specific classes** — `rescue => e` catches everything under `StandardError` and can mask bugs
2. **Don't rescue `Exception`** — it catches `SystemExit`, `Interrupt`, `NoMemoryError`; never rescue the root `Exception` class
3. **Always re-raise if you can't handle** — `raise` (bare) re-raises `$!`
4. **Use `ensure` for cleanup** — it runs regardless of how the block exits
5. **`else` is rarely needed** — it only runs when no exception occurred, and can mask logic; prefer explicit success path inside `begin`

## Edge Cases and Known Issues

- **`rescue` in a method body without `begin`** — a method body is implicitly a `begin...end` block, so you can write `rescue` directly:
  ```ruby
  def foo
    risky
  rescue
    "fallback"
  end
  ```
- **`else` doesn't run if `rescue` ran** — but it also doesn't run if `begin` exits via `return` or `break`
- **`ensure` return value is ignored** — except when `begin` doesn't return normally; if `ensure` has a `return`, it overrides the `begin` block's return value (rare, avoid)
- **`retry` re-runs `begin`** — the entire block, not just the failing line
- **`raise` with no current exception** — raises `RuntimeError` with no message (when not inside a `rescue`)
- **`for` loop variable leaks** — unlike `each`, the block variable persists after the loop

## Related

- [blocks-procs-lambdas.md](blocks-procs-lambdas.md) — `break`/`next` in blocks passed to methods
- [types.md](types.md) — `Comparable` uses `<=>` for `case` comparisons
- [object-model.md](object-model.md) — custom exception classes
