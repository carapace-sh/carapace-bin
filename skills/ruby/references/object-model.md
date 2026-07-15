# Object Model

Ruby's object model — classes, modules, method lookup, eigenclasses, constant resolution, and the relationship between `BasicObject`, `Object`, `Kernel`, and `Module`.

> **Source of truth**: <https://ruby-doc.org/core/> and <https://github.com/ruby/ruby/blob/master/doc/syntax.rdoc>. For **metaprogramming** APIs (define_method, eval, etc.), see [metaprogramming.md](metaprogramming.md).

## Class Hierarchy

Every value in Ruby is an object. The root of the class hierarchy is `BasicObject`:

```
BasicObject
  └── Object (includes Kernel)
        └── ... all user classes ...
              └── String, Integer, Array, Hash, ...
```

| Class | Role |
|-------|------|
| `BasicObject` | Minimal root — almost no methods. Useful for delegate/proxy objects that must not shadow any method. |
| `Object` | The default root for all classes. Includes `Kernel`, giving every object `puts`, `require`, `raise`, etc. |
| `Kernel` | A module mixed into `Object`. Provides system-level methods like `puts`, `gets`, `require`, `exit`, `sleep`, `rand`, `proc`, `lambda`. |
| `Module` | Container for methods and constants. Superclass of `Class`. |
| `Class` | A subclass of `Module` that can be instantiated. |

```ruby
class Foo; end
Foo.class                   # => Class
Foo.class.superclass         # => Module
Foo.class.superclass.superclass # => Object
Foo.superclass              # => Object
Foo.ancestors               # => [Foo, Object, Kernel, BasicObject]
```

## Method Lookup

When Ruby sends a message to an object, it walks the **method resolution chain** (ancestors):

1. The object's **eigenclass** (singleton class)
2. The object's class
3. Each included/prepended module in reverse inclusion order
4. Superclass and its modules
5. ... up to `BasicObject`

```ruby
module M1; def greet; "M1"; end; end
module M2; def greet; "M2"; end; end

class C
  prepend M1
  include M2
  def greet; "C"; end
end

C.new.greet  # => "M1" (prepend inserts M1 before C in ancestors)
C.ancestors  # => [M1, C, M2, Object, Kernel, BasicObject]
```

### include vs prepend vs extend

| Operation | Inserts where in ancestors | Affects |
|-----------|-----------------------------|---------|
| `include M` | After the class, before its superclass | Instance methods |
| `prepend M` | Before the class itself | Instance methods (M overrides C) |
| `extend M` | Into the eigenclass | Singleton (class) methods |

```ruby
module Trackable
  def track
    super  # calls the original method (works with prepend)
    puts "tracked"
  end
end

class User
  prepend Trackable
  def track; puts "original"; end
end

User.new.track
# => "original"
# => "tracked"
```

`super` in a prepended module calls the **original** method. `super` in an included module calls the next method in the ancestor chain above the module.

## Eigenclass (Singleton Class)

Every object has an **eigenclass** (also called the singleton class) — a hidden class that holds methods defined on that one object:

```ruby
str = "hello"
def str.shout
  upcase + "!"
end
str.shout        # => "HELLO!"
"world".shout    # => NoMethodError

str.singleton_class.instance_methods(false)  # => [:shout]
```

For class objects, this is how class methods work:

```ruby
class Foo
  def self.bar  # defines on Foo's eigenclass
    "bar"
  end
end
Foo.bar  # => "bar"
```

`self` inside `class << self ... end` is the eigenclass itself.

## Constant Resolution

Constants are identified by a leading uppercase letter. They are resolved via the **constant lookup** path, which differs from method lookup:

1. **Lexical scope** — constants defined in the enclosing module/class of the code
2. **Inheritance chain** — `Module.nesting` plus ancestors
3. `Object` — top-level constants

```ruby
module A
  CONST = "A"

  class B
    CONST = "B"

    def self.which
      CONST  # => "B" (lexical lookup wins)
    end
  end
end

A::B.which      # => "B"
A::B::CONST     # => "B"
A::CONST        # => "A"
```

`::Foo` (leading `::`) bypasses the lookup and starts at top-level `Object`.

### autoload

```ruby
module DB
  autoload :User, "db/user"
  autoload :Post, "db/post"
end

DB::User  # requires "db/user" on first access, then returns the constant
```

## Method Missing and Friends

When method lookup fails, Ruby invokes several "magic" methods in sequence:

| Method | When called |
|--------|------------|
| `method_missing` | No matching method found |
| `respond_to_missing?` | `respond_to? :foo` check for dynamic methods |
| `method_added` (class) | When a new instance method is defined |
| `singleton_method_added` (class) | When a singleton method is defined |

```ruby
class Dynamic
  def method_missing(name, *args)
    if name.to_s.start_with?("fetch_")
      "fetched #{name}"
    else
      super
    end
  end

  def respond_to_missing?(name, include_private = false)
    name.to_s.start_with?("fetch_") || super
  end
end

d = Dynamic.new
d.fetch_users  # => "fetched fetch_users"
d.respond_to?(:fetch_users)  # => true
```

Always override `respond_to_missing?` when overriding `method_missing`.

## ObjectSpace

`ObjectSpace` provides access to all live objects in the runtime:

```ruby
ObjectSpace.each_object(String) { |s| puts s if s.frozen? }

# Count instances of a class
ObjectSpace.count_objects  # returns hash of type counts
ObjectSpace.count_objects[:T_STRING]  # number of live strings

# Find by object_id
obj = "hello"
ObjectSpace._id2ref(obj.object_id)  # => "hello"

# Define a finalizer
ObjectSpace.define_finalizer(obj) { puts "finalized" }
```

Finalizers run **after** the object is GC'd, so they cannot reference the object directly — capture only the object_id.

## Edge Cases and Known Issues

- **`include` is idempotent** — including the same module twice has no effect
- **`prepend` + `super`** — the prepended module can call `super` to reach the original method
- **`method_missing` is slow** — it bypasses inline caching; prefer `define_method` for known method names
- **`ObjectSpace.each_object` can miss objects** — objects that have been finalized or are being finalized may not be enumerated
- **`autoload` is not thread-safe** in MRI < 3.0 — concurrent autoload of the same constant could require the file twice. Fixed in Ruby 3.0+ with a per-constant lock.
- **`Kernel` is included in `Object`** — so `Kernel.private_instance_methods` includes `puts`, `require`, etc., callable without `self.`

## Related

- [metaprogramming.md](metaprogramming.md) — dynamic method definition, `eval`, `Binding`, `TracePoint`
- [types.md](types.md) — the built-in types (String, Integer, Array, etc.)
- [gc-memory.md](gc-memory.md) — `ObjectSpace` internals, finalizers, GC
