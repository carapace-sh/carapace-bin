# Core Types

Ruby's built-in types — String, Symbol, Integer, Float, Numeric, Array, Hash, Range, Regexp, MatchData, Proc, Method, and the special values `nil`, `true`, `false`.

> **Source of truth**: <https://ruby-doc.org/core/>. For **blocks/Procs/lambdas**, see [blocks-procs-lambdas.md](blocks-procs-lambdas.md).

## Truthiness

In Ruby, everything is truthy **except** `false` and `nil`:

```ruby
!!0        # => true   (0 is truthy in Ruby, unlike C/Python)
!!""       # => true   (empty string is truthy)
!![]       # => true   (empty array is truthy)
!!nil      # => false
!!false    # => false
```

## nil

`nil` is the "no value" sentinel. It's the only instance of `NilClass`:

```ruby
nil.nil?        # => true
nil.to_a        # => []  (convenient for splat)
nil.to_s        # => ""  (convenient for interpolation)
nil.inspect     # => "nil"
nil&.upcase     # => nil  (safe navigation)
```

`true` and `false` are instances of `TrueClass` / `FalseClass`. There's no `Boolean` superclass — use `obj == true` or `!!obj` if you need a boolean.

## Equality

Ruby has three main equality methods with different semantics:

| Method | Purpose | Strict |
|--------|---------|--------|
| `equal?` | Object identity (same `object_id`) | Strictest — rarely overridden |
| `==` | Value equality | Overridden per-type; usually case-insensitive for some |
| `eql?` | Strict value equality (for Hash keys) | Usually stricter than `==` |

```ruby
1 == 1.0       # => true   (== allows numeric coercion)
1.eql?(1.0)   # => false  (eql? does not)
1.hash == 1.0.hash  # => false (different hash for Integer vs Float)

# String equality
"foo" == "foo"      # => true
"foo".eql?("foo")   # => true
"foo".equal?("foo") # => false (different objects)
```

### <=> (Spaceship)

The `<=>` operator returns -1, 0, 1, or `nil` (incomparable). `Comparable` adds `<`, `<=`, `==`, `>=`, `>`, `between?` based on `<=>`:

```ruby
class Version
  include Comparable
  attr_reader :parts
  def initialize(str); @parts = str.split(".").map(&:to_i); end
  def <=>(other); parts <=> other.parts; end
end

Version.new("1.2.3") < Version.new("1.3")  # => true
```

## String

Strings are mutable byte sequences with an associated `Encoding` (default UTF-8 in source files):

```ruby
s = "hello"
s << " world"
s.upcase!  # in-place mutation
s.frozen?  # => false (unless from a literal in a frozen-string-literal file)
```

**Frozen string literals**: With the magic comment `# frozen_string_literal: true`, all string literals in the file are frozen (immutable):

```ruby
# frozen_string_literal: true
s = "hello"
s << " world"  # => RuntimeError: can't modify frozen String
```

Key methods: `+` (concat, returns new), `<<` (append, mutates), `[]` (slice/index), `gsub`, `sub`, `split`, `scan`, `strip`, `chomp`, `each_line`, `chars`, `bytes`, `encode`, `force_encoding`.

See [io-encoding.md](io-encoding.md) for encoding details.

## Symbol

Symbols are immutable, interned strings. They have a fixed `object_id` for the same name across the runtime:

```ruby
:foo.object_id == :foo.object_id  # => true
"foo".object_id == "foo".object_id  # => usually false (different objects)
```

- Use Symbols for **identifiers** (hash keys, method names, enum-like values)
- Use Strings for **data** (user input, file content, text processing)
- `Symbol#to_s` converts to String; `String#to_sym` converts to Symbol
- Symbols are **not** frozen Strings — they're a distinct type

```ruby
hash = { name: "Alice", age: 30 }  # keys are :name, :age (Symbols)
hash[:name]  # => "Alice"
```

## Integer and Float

Ruby has arbitrary-precision integers (no overflow):

```ruby
10 ** 100  # => a 101-digit Integer
(2 ** 64).class  # => Integer
```

Division:
- `7 / 2`  → `3`  (integer division)
- `7.0 / 2` → `3.5` (float division when either operand is Float)
- `7.fdiv(2)` → `3.5` (always returns Float)
- `7.quo(2)` → `(7/2)` as a Rational

Numeric hierarchy: `Numeric` → `Integer`, `Float`, `Rational`, `Complex`, `BigDecimal` (stdlib).

## Array

Arrays are ordered, integer-indexed, and heterogeneous:

```ruby
a = [1, "two", :three]
a.first   # => 1
a.last    # => :three
a[-1]     # => :three (negative index from end)
a[1, 2]   # => ["two", :three] (start, length)
a[0..1]   # => [1, "two"] (range)
```

Key methods: `push`/`<<`, `pop`, `shift`, `unshift`, `map`, `select`/`filter`, `reject`, `reduce`/`inject`, `each`, `each_with_index`, `zip`, `flatten`, `compact`, `uniq`, `sort`, `sort_by`, `group_by`, `partition`, `tally`, `min`, `max`, `sum`, `first(n)`, `sample`, `rotate`.

Negative indices and ranges are idiomatic. `Array.new(size) { |i| ... }` initializes with a block.

## Hash

Hashes are ordered (insertion order, since Ruby 1.9):

```ruby
h = { name: "Alice", age: 30 }
h[:name]      # => "Alice"
h.fetch(:missing, "default")  # => "default"
h.transform_values { |v| v.to_s }  # => { name: "Alice", age: "30" }
```

Key methods: `[]`, `fetch`, `store`, `delete`, `keys`, `values`, `each`, `map`, `transform_values`, `transform_keys`, `merge`, `dig`, `to_a`, `invert`, `group_by`, `tally`, `filter_map`.

Hash keys use `eql?` and `hash` for equality. `Symbol` and small `Integer` are common keys.

## Range

```ruby
(1..5).to_a      # => [1, 2, 3, 4, 5]
(1...5).to_a     # => [1, 2, 3, 4]  (exclusive)
('a'..'e').to_a  # => ["a", "b", "c", "d", "e"]
(1..).to_a       # => error (endless Range, can't convert to Array)
```

- `..` is inclusive, `...` is exclusive
- `(1..)` is an endless Range (Ruby 2.6+)
- `(..5)` is a beginless Range
- `(1..10).step(2)` → enumerator stepping by 2
- `(1..10).cover?(5)` → true (uses `<=>`)
- `(1..10).include?(5)` → alias for `cover?` (or member check for non-comparable)

## Regexp and MatchData

```ruby
/(\w+)@(\w+)/.match("user@example.com")
# => #<MatchData "user@example.com" 1:"user" 2:"example.com">

md = "user@example.com"[/(\w+)@(\w+)/, 1]  # => "user"
```

- `/.../` literal; `%r{...}` for strings with slashes
- `?`, `*`, `+`, `{n,m}`, `^`, `$`, `\b`, `\d`, `\w`, `\s` (Perl-style)
- Named captures: `/(?<name>\w+)/` → `md[:name]`
- `Regexp#match` returns `MatchData` or `nil`
- `String#=~` returns index of match or `nil`; `String#match?` is a boolean (faster, no MatchData allocation)
- `String#scan` returns all matches as an array
- `String#gsub` with block or hash substitution

## Proc and Method

`Proc` and `Method` are callable objects — see [blocks-procs-lambdas.md](blocks-procs-lambdas.md) for full details.

## freeze and Immutability

```ruby
s = "hello"
s.freeze
s << "!"  # => RuntimeError: can't modify frozen String

Object.new.frozen?  # => false
Symbol.all_symbols.first.frozen?  # => true (Symbols are always frozen)
Integer.ancestors.first.frozen?   # => true (Integer instances are frozen)
```

`freeze` makes an object's state immutable. It's shallow — a frozen Array's elements can still be mutated unless they're frozen too. `dup` and `clone` create copies; `clone` preserves the frozen state.

## Edge Cases and Known Issues

- **`0` is truthy** — unlike C/Python, `0` is not falsy in Ruby
- **`nil.to_s` returns `""`** — convenient but can mask bugs; use `&.` to make absence explicit
- **Float precision** — `0.1 + 0.2 == 0.3` is `false`; use `Rational` or `BigDecimal` for exact decimal
- **String key vs Symbol key** — `{ "foo" => 1 }[:foo]` returns `nil`; String and Symbol keys are distinct
- **`==` on String** returns true for same content, but they're different objects (see `equal?`)
- **Array equality** uses `==` element-wise; `equal?` checks identity
- **Frozen Integer literals** — small integers are pre-allocated and frozen; `1.frozen?` returns `true` in Ruby 3.0+
- **`Regexp` with `=~` returns integer index** — can be truthy `0`, which is truthy in Ruby

## Related

- [blocks-procs-lambdas.md](blocks-procs-lambdas.md) — `Proc`, `Method`, `&block`
- [io-encoding.md](io-encoding.md) — String encoding details
- [object-model.md](object-model.md) — type hierarchy and `Comparable`
