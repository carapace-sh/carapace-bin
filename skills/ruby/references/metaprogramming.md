# Metaprogramming

Ruby's metaprogramming facilities — dynamic method definition, `eval` and `Binding`, `send`/`method`, `method_missing`, `Refinements`, and `TracePoint`.

> **Source of truth**: <https://ruby-doc.org/core/> and <https://github.com/ruby/ruby/blob/master/doc/syntax.rdoc>. For the **object model** (class hierarchy, method lookup, ancestors), see [object-model.md](object-model.md).

## Dynamic Method Definition

### define_method

```ruby
class Foo
  [:admin, :user, :guest].each do |role|
    define_method("is_#{role}?") { self.role == role }
  end
end

Foo.new.is_admin?  # calls the dynamically defined method
```

`define_method` takes a block (closure) as the method body. The block captures its defining context, so it has access to `role` at call time.

### define_singleton_method

For class-level or per-object methods:

```ruby
str = "hello"
str.define_singleton_method(:shout) { upcase + "!" }
str.shout  # => "HELLO!"
```

## class_eval / instance_eval / module_eval

These methods change `self` and the constant lookup context:

| Method | `self` becomes | Constant lookup |
|--------|----------------|-----------------|
| `instance_eval { }` | The receiver | Single nesting level |
| `class_eval { }` / `module_eval { }` | The receiver class/module | Normal nesting |

```ruby
String.class_eval do
  def shout
    upcase + "!"
  end
end

"hello".shout  # => "HELLO!"
```

`class_eval` with a string evaluates in the class context:

```ruby
Integer.class_eval("def double; self * 2; end")
5.double  # => 10
```

## eval and Binding

`eval` executes a string of Ruby code:

```ruby
eval("x = 1 + 2")  # => 3
```

`Binding` captures the execution context (variables, `self`, method dispatch) at a point:

```ruby
def get_binding
  local_var = 42
  binding
end

b = get_binding
eval("local_var", b)  # => 42
```

`Binding#receiver` returns the object, `Binding#local_variables` returns the list of locals, `Binding#local_variable_get` / `Binding#local_variable_set` access them.

```ruby
TOPLEVEL_BINDING.eval("self")  # => main
```

## send and friends

Send a message to an object by name (even private methods):

```ruby
"hello".send(:upcase)          # => "HELLO"
"hello".public_send(:upcase)  # => "HELLO" (raises on private methods)
1.send(:+, 2)                 # => 3
```

Get a `Method` object (callable, can be converted to a Proc):

```ruby
method = "hello".method(:upcase)
method.call        # => "HELLO"
method.()          # => "HELLO" (shorthand call)
method.to_proc     # => a Proc
method.arity       # => 0
method.receiver    # => "hello"

# Convert symbol to proc with &:method(:upcase)
%w[hello world].map(&:upcase)  # => ["HELLO", "WORLD"]
```

## method_missing and Dynamic Dispatch

When method lookup fails, `method_missing` is called:

```ruby
class Proxy
  def initialize(target)
    @target = target
  end

  def method_missing(name, *args, &block)
    if @target.respond_to?(name)
      @target.send(name, *args, &block)
    else
      super
    end
  end

  def respond_to_missing?(name, include_private = false)
    @target.respond_to?(name, include_private) || super
  end
end
```

**Performance note**: `method_missing` bypasses inline method caching. For a fixed set of method names, prefer `define_method` at definition time.

## Hooks

Ruby calls special methods when the object model changes:

| Hook | When |
|------|------|
| `method_added(name)` | Instance method defined in this class |
| `singleton_method_added(name)` | Singleton method defined on this object |
| `method_undefined(name)` | `undef_method` called |
| `method_removed(name)` | `remove_method` called |
| `included(mod)` | Module is `include`d into a class |
| `prepended(mod)` | Module is `prepend`ed into a class |
| `extended(mod)` | Module is `extend`ed onto an object |
| `inherited(subclass)` | Class is subclassed |
| `const_missing(name)` | Constant not found in lookup |

```ruby
module Observable
  def self.included(base)
    base.extend(ClassMethods)
  end

  module ClassMethods
    def observe(*args); end
  end
end

class Model
  include Observable  # triggers Observable.included(self)
end
```

## Refinements

Refinements are lexically-scoped class modifications. They avoid the global impact of monkey-patching:

```ruby
module ShoutString
  refine String do
    def shout
      upcase + "!"
    end
  end
end

module MyApp
  using ShoutString  # activated for the rest of this module's scope

  "hello".shout  # => "HELLO!"
end

"hello".shout  # => NoMethodError (refinement not active here)
```

- `using` activates a refinement until the end of the current file/module/class scope
- Refinements are **not** inherited by subclasses of the class that uses them (in lexical scope only)
- `Module.used_refinements` returns the active refinements
- Refinements affect method lookup: refined methods are inserted **before** the class's own methods in the lookup chain

## TracePoint

`TracePoint` replaces `set_trace_func` for tracing execution events:

```ruby
TracePoint.new(:call, :return) do |tp|
  puts "#{tp.event} #{tp.defined_class}##{tp.method_id}"
end.enable do
  my_method()
end
# call  MyModule#my_method
# return MyModule#my_method
```

Events: `:line`, `:class`, `:end`, `:call`, `:return`, `:b_call`, `:b_return`, `:raise`, `:rescue`, `:c_call`, `:c_return`, `:thread_begin`, `:thread_end`, `:fiber_switch`.

`TracePoint` is preferred over `set_trace_func` because it allows per-event-type filtering and can be enabled only within a block.

## const_set / const_get

```ruby
module Config
  const_set(:VERSION, "1.0")
end
Config::VERSION        # => "1.0"
Config.const_get(:VERSION)  # => "1.0"
Config.const_defined?(:VERSION)  # => true
```

`const_get` with a string traverses paths: `Object.const_get("Config::VERSION")`.

## Edge Cases and Known Issues

- **`instance_eval` block changes `self`** but not `Method` resolution — methods defined in the block are singleton methods on the receiver, not on the enclosing class
- **Refinements are lexical** — they don't propagate through `send` or dynamic dispatch unless the call site is in the refined scope
- **`method_missing` + `super`** — always call `super` when you can't handle the message, to allow normal `NoMethodError`
- **`define_method` captures the block** — variables in the defining context are shared, not copied; later mutations are visible
- **`eval` with untrusted input is dangerous** — use `RubyVM::InstructionSequence.compile_option = {safe_intern: true}` and consider a sandbox; the old `$SAFE` mechanism was removed in Ruby 2.7+
- **`class_eval` with a block vs string** — block form has normal `self`/constant nesting; string form evaluates in the class's context (like reopening the class)
- **`Binding#eval`** is the mechanism behind `binding.irb` and `binding.pry` debug sessions

## Related

- [object-model.md](object-model.md) — class hierarchy, method lookup, ancestors
- [blocks-procs-lambdas.md](blocks-procs-lambdas.md) — `define_method` takes a block
- [types.md](types.md) — `Method`, `Proc`, `Symbol#to_proc`
