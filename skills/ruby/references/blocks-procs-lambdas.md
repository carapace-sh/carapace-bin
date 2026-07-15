# Blocks, Procs, and Lambdas

Ruby's callable objects — blocks, `Proc`, `lambda`, `Method`, and the `&` operator. Covers argument semantics, `return` behavior, `yield`, and conversion between types.

> **Source of truth**: <https://ruby-doc.org/core/Proc.html>, <https://github.com/ruby/ruby/blob/master/doc/syntax.rdoc>. For **metaprogramming** with `define_method`, see [metaprogramming.md](metaprogramming.md).

## Blocks

A block is an unnamed chunk of code passed to a method. Every method can accept one implicit block:

```ruby
[1, 2, 3].each { |n| puts n }          # do/end block
[1, 2, 3].map { |n| n * 2 }             # braces block

[1, 2, 3].each do |n|
  puts n
end
```

Braces `{}` and `do...end` are semantically equivalent. Convention: braces for single-line, `do...end` for multi-line.

### yield

`yield` invokes the block:

```ruby
def twice
  yield
  yield
end

twice { puts "hi" }
# hi
# hi
```

`yield` takes arguments and returns the block's result:

```ruby
def with_result
  result = yield 10
  puts "block returned: #{result}"
end

with_result { |n| n * 2 }  # => "block returned: 20"
```

### block_given?

Check whether a block was passed:

```ruby
def maybe_yield
  if block_given?
    yield
  else
    "no block"
  end
end
```

### Block Arguments

Blocks have lenient arity by default — extra arguments are `nil`, extra values are silently ignored:

```ruby
proc { |a, b| [a, b] }.call(1)      # => [1, nil]
proc { |a, b| [a, b] }.call(1, 2, 3) # => [1, 2] (extra ignored)
```

## Proc

A `Proc` is a block wrapped as an object. Create with `proc` or `Proc.new`:

```ruby
p = Proc.new { |x| x * 2 }
p.call(5)   # => 10
p.(5)       # => 10  (shorthand)
p[5]        # => 10  (another shorthand)
p === 5     # => 10  (used in case/when)
```

### return in a Proc

`return` in a `Proc` returns from the **enclosing method** (not just the Proc):

```ruby
def foo
  p = Proc.new { return 1 }
  p.call
  puts "never reached"
end

foo  # => 1
```

This is because a Proc is a closure that shares the return context of where it was defined.

### break in a Proc

`break` in a Proc behaves differently depending on context:

- In a method with a block: breaks out of the block, returns `nil` to the method
- In a loop (`loop`): breaks the loop, returns the value
- At top-level: raises `LocalJumpError`

## lambda

Lambdas are stricter Procs — they check arity and `return` only returns from the lambda:

```ruby
lam = lambda { |a, b| a + b }
lam.call(1)       # => ArgumentError (wrong number of arguments)
lam.call(1, 2)    # => 3

# Lambda literal syntax (->)
add = ->(a, b) { a + b }
add.(1, 2)  # => 3
```

### return in a lambda

`return` in a lambda returns from the **lambda** only, not the enclosing method:

```ruby
def bar
  lam = lambda { return 1 }
  lam.call
  puts "reached"
end

bar  # => prints "reached"
```

## Proc vs Lambda

| Aspect | Proc | Lambda |
|--------|------|--------|
| Arity check | Lenient (extra args → `nil`) | Strict (raises `ArgumentError`) |
| `return` | Returns from enclosing method | Returns from the lambda |
| `self` | Closure `self` (where defined) | Same |
| `lambda?` | `false` | `true` |
| Default args | Supported | Supported |

```ruby
proc {}.lambda?     # => false
lambda {}.lambda?    # => true
```

## Method

`Method` objects are bound to a specific method on a specific receiver:

```ruby
method = "hello".method(:upcase)
method.call      # => "HELLO"
method.()        # => "HELLO"
method.receiver  # => "hello"
method.name      # => :upcase
method.arity     # => 0
method.to_proc   # => a Proc that calls the method

# Unbound (class-level, needs binding)
um = String.instance_method(:upcase)
um.bind("hello").call  # => "HELLO"
```

`Method` objects behave like lambdas — strict arity, `return` returns from the method.

## The & Operator

The `&` prefix converts between blocks and Procs:

```ruby
# Proc → block
p = Proc.new { |x| x * 2 }
[1, 2, 3].map(&p)  # => [2, 4, 6]

# block → Proc (capture in method)
def capture(&block)
  block  # returns a Proc
end

# Symbol → Proc (calls Symbol method on each element)
%w[hello world].map(&:upcase)  # => ["HELLO", "WORLD"]

# Method → Proc
m = "hello".method(:upcase)
["a", "b"].map(&m)  # => ["A", "B"] (calls m on each)
```

`&` calls `to_proc` on the object if it's not already a Proc. `Symbol#to_proc` returns `->(obj) { obj.send(self) }`.

## Capturing Blocks

```ruby
def render(&block)
  content = block.call
  "<div>#{content}</div>"
end

render { "Hello" }  # => "<div>Hello</div>"
```

`block.call` is equivalent to `yield`, but the block is a first-class Proc object.

## Argument Splat and Double Splat

```ruby
# Method definition with splat
def log(*args)
  args.inspect
end

# Double splat (keyword arguments)
def configure(**opts)
  opts
end

configure(name: "foo", level: 3)  # => { name: "foo", level: 3 }

# Splat in call
args = [1, 2, 3]
log(*args)  # spreads array into arguments

# Double splat in call
opts = { name: "foo" }
configure(**opts)  # spreads hash into keyword arguments
```

## curry

`Proc#curry` returns a curried Proc:

```ruby
add = lambda { |a, b| a + b }
add_one = add.curry[1]
add_one.(2)  # => 3

# Proc (lenient) also works
p = proc { |a, b, c| a + b + c }.curry
p.(1).(2).(3)  # => 6
```

## block vs Proc vs lambda vs Method

| Type | Created by | Strict arity? | `return` from | `self` |
|------|-----------|---------------|---------------|--------|
| Block | `{ ... }` or `do...end` | No (lenient) | Enclosing method | Definition scope |
| Proc | `Proc.new {}` or `proc {}` | No (lenient) | Enclosing method | Definition scope |
| Lambda | `lambda {}` or `->{}` | Yes | Lambda only | Definition scope |
| Method | `obj.method(:name)` | Yes | Method | Receiver |

## Edge Cases and Known Issues

- **`break` in a block passed to `loop`** — returns the value: `loop { break 42 }` returns `42`
- **`next` in a block** — skips to the next iteration, can take a value: `next 42`
- **`Proc#call` vs `Proc#()`** — `p.()` is syntactic sugar for `p.call()`, but slightly faster
- **`lambda` with `&`** — when a lambda is passed as a block with `&`, it keeps its strict arity
- **`define_method` takes a block (Proc semantics)** — the method body has lenient arity and `return` returns from the method (like Proc)
- **`method(:foo)` creates a new Method object each time** — it's not interned; `Method#to_proc` creates a new Proc each time too

## Related

- [metaprogramming.md](metaprogramming.md) — `define_method` (takes a block)
- [types.md](types.md) — `Symbol#to_proc`, `Method`
- [control-flow.md](control-flow.md) — `break`, `next`, `return`, `retry` in blocks
