# GC and Memory

Ruby's garbage collection, `ObjectSpace`, object representation, and memory management. Covers the generational GC (RGenGC), write barriers, weak references, finalizers, and internal object layout.

> **Source of truth**: <https://ruby-doc.org/core/GC.html>, <https://ruby-doc.org/core/ObjectSpace.html>, and the `objspace` stdlib docs. For **concurrency** interactions with GC, see [concurrency.md](concurrency.md).

## GC Module

The `GC` module provides control and introspection over the garbage collector:

```ruby
GC.start                    # force a full GC cycle
GC.disable                  # disable GC (use with caution)
GC.enable                   # re-enable
GC.stress = true            # GC after every allocation (debugging)
GC.stress = false           # normal mode
GC.auto_compact = true      # auto-compact heap (reduce fragmentation)
```

### GC.stat

```ruby
GC.stat
# => {
#   :count=>42,                # major GC count
#   :minor_gc_count=>100,      # minor GC count
#   :major_gc_count=>42,       # major GC count (same as :count)
#   :heap_allocated_pages=>...,
#   :heap_sorted_length=>...,
#   :heap_allocatable_slots=>...,
#   :heap_available_slots=>...,
#   :heap_live_slots=>...,     # live objects
#   :heap_free_slots=>...,     # free slots
#   :heap_final_slots=>...,    # objects awaiting finalization
#   :heap_marked_slots=>...,   # marked (reachable)
#   :total_allocated_objects=>...,  # lifetime allocations
#   :total_freed_objects=>...,
#   :malloc_increase_bytes=>...,
#   :malloc_increase_bytes_limit=>...,
#   :gc_time_ms=>...           # total time in GC (Ruby 3.3+)
# }
```

`GC.stat(:count)` fetches a single stat. `GC.stat_heap` returns per-heap stats.

### GC Config (Environment Variables)

| Variable | Effect |
|----------|--------|
| `RUBY_GC_HEAP_INIT_SLOTS` | Initial number of object slots |
| `RUBY_GC_HEAP_GROWTH_FACTOR` | Heap growth multiplier |
| `RUBY_GC_HEAP_GROWTH_MAX_SLOTS` | Max slots to add per growth |
| `RUBY_GC_HEAP_OLDOBJECT_LIMIT_FACTOR` | Old object threshold for major GC |
| `RUBY_GC_HEAP_MIN_SLOTS` | Minimum heap slots |
| `RUBY_GC_MALLOC_LIMIT_MIN` | Minimum malloc increase limit |
| `RUBY_GC_MALLOC_LIMIT_MAX` | Maximum malloc increase limit |
| `RUBY_GC_MALLOC_LIMIT_GROWTH_FACTOR` | Growth factor for malloc limit |

These are typically set via `RUBYOPT` or environment, not code.

## Generational GC (RGenGC)

MRI uses **RGenGC** — a generational, incremental, mark-and-sweep collector:

- **Minor GC** — scans only young objects (fast, frequent). Young objects that survive a minor GC become "old".
- **Major GC** — scans all objects (slow, less frequent). Triggered when the old object pool grows, or explicitly via `GC.start`.

### Write Barriers

RGenGC relies on **write barriers** to track cross-generational references. When an old object gains a reference to a young object, the write barrier marks the old object as "shady" (needs re-scanning).

Most built-in types have write barriers. User-defined classes get them automatically (through `attr_accessor`, `attr_writer`, and direct instance variable writes), **except** for objects modified via C extensions that bypass the barrier — these become "shady" and force major GCs.

## Object Representation

### VALUE and Tagged Pointers

In MRI, every Ruby object is represented by a `VALUE` (a C type). Small values use **tagged pointers**:

| Type | Tag | Stored in VALUE directly? |
|------|-----|---------------------------|
| `Integer` (small) | ...001 (bit 0 set) | Yes (immediate) |
| `Symbol` | ...01110 (0x0e) | Yes (immediate) |
| `nil` | 0x08 | Yes (immediate, special VALUE) |
| `true` | 0x14 | Yes (immediate) |
| `false` | 0x00 | Yes (immediate) |
| `Float` (some) | — | No, allocated (except flonum optimization) |

"Immediate" values have no separate allocation — the `VALUE` **is** the value.

### RVALUE

Non-immediate objects are allocated on the heap as **RVALUE** structs. Each RVALUE is a fixed-size slot (40 bytes on 64-bit). If an object needs more, it stores a pointer to a malloc'd buffer.

- `RString` — inline for short strings (< 24 bytes), else malloc'd buffer
- `RArray` — inline for ≤ 3 elements, else malloc'd buffer
- `RHash` — inline table pointer + size

### Object Shapes (Ruby 3.0+)

MRI uses **object shapes** to optimize instance variable access. Each unique set of IV names maps to a shape. Object creation is O(1) — the shape is assigned based on the class and IV layout.

- Two objects of the same class with the same IVs have the **same shape**
- Adding/removing an IV transitions the shape
- Shape changes are tracked per-class; transitions are cached
- This speeds up method dispatch for `attr_accessor` methods by allowing inline caches to key on shape

## ObjectSpace

```ruby
ObjectSpace.each_object(String) { |s| ... }   # iterate live strings
ObjectSpace.each_object(Integer) { |i| ... }  # iterate small ints (immediates not yielded)
ObjectSpace.count_objects                      # => { T_STRING => 1234, ... }
ObjectSpace.count_objects[:T_STRING]           # => 1234
ObjectSpace._id2ref(obj.object_id)             # fetch object by ID
```

`count_objects` returns a hash keyed by internal type:

| Key | Type |
|-----|------|
| `:TOTAL` | Total live objects |
| `:FREE` | Free slots |
| `:T_OBJECT` | Generic objects |
| `:T_CLASS` | Classes and modules |
| `:T_STRING` | Strings |
| `:T_ARRAY` | Arrays |
| `:T_HASH` | Hashes |
| `:T_DATA` | C-extension data (wrapped structs) |
| `:T_FLOAT` | Floats |
| `:T_IVAR` | Instance variables container (internal) |
| `:T_SYMBOL` | Symbols |
| `:T_NODE` | AST nodes (not yet compiled) |
| `:T_MUTEX` / `:T_THREAD` / `:T_FIBER` | Concurrency primitives |

### Finalizers

```ruby
obj = Object.new
ObjectSpace.define_finalizer(obj) { puts "finalized!" }
obj = nil
GC.start  # triggers finalizer
```

**Important**: finalizer blocks run **after** the object is gone. They cannot reference the object. Capture only primitive values (like `object_id`) or frozen data:

```ruby
ObjectSpace.define_finalizer(obj) { |id| log("finalized #{id}") }
```

Finalizers run in a dedicated thread (since Ruby 3.2+). They can raise, and Ruby prints the error but continues.

## WeakRef and WeakMap

```ruby
require 'weakref'
ref = WeakRef.new(some_big_object)
ref.__getobj__        # the strong reference (raises if GC'd)
ref.weakref_alive?   # true if the object still exists

# WeakMap: keys or values can be GC'd
wm = ObjectSpace::WeakMap.new
wm[obj] = metadata
```

`ObjectSpace::WeakMap` holds weak references — if a key (or value, depending on implementation) is GC'd, the entry is removed.

`WeakRef` wraps an object in a weak reference. `__getobj__` retrieves it, but raises `WeakRef::RefError` if the object has been collected.

## Memory Profiling

```ruby
require 'objspace'
ObjectSpace.memsize_of(obj)          # bytes used by this object
ObjectSpace.memsize_of_all(String)   # total bytes used by all strings
ObjectSpace.trace_object_allocations # start tracking allocation sites
ObjectSpace.allocation_sourcefile(obj)
ObjectSpace.allocation_sourceline(obj)
```

For production profiling, use `memory_profiler` or `stackprof` gems (see [cli.md](cli.md)).

## Heap Layout

```
Ruby heap:
  Heap pages (each 16 KB, containing ~408 RVALUE slots)
    ├── RVALUE slot (40 bytes)
    ├── RVALUE slot
    └── ...
```

- Objects are allocated from heap pages
- When a page is full, a new page is allocated (growth)
- Freed slots become free-list entries
- `GC.auto_compact` (Ruby 3.0+) moves live objects to reduce fragmentation

## Edge Cases and Known Issues

- **`GC.disable` is dangerous** — memory grows unbounded; use `GC.start` manually if you disable
- **Finalizers run later** — they run on the next GC after the object is unreachable, and may not run at all if the process exits
- **`ObjectSpace.each_object` is slow** — it scans the entire heap; avoid in hot paths
- **Symbols are never GC'd** (pre-2.2) — dynamic symbols (`to_sym` on a String) are GC'd since Ruby 2.2; static symbols (literals) are not
- **Shady objects force major GC** — a single shady object can increase GC frequency for the whole heap
- **`memsize_of` is approximate** — it reports the RVALUE + inline buffer, but shared buffers (like shared strings) may double-count
- **`GC.implicit_compaction_disabled?`** — compaction may be disabled on some platforms due to C extension issues with moving objects

## Related

- [object-model.md](object-model.md) — `ObjectSpace.each_object` for class instances
- [concurrency.md](concurrency.md) — threads, fibers, and GC interaction
- [stdlib.md](stdlib.md) — `objspace`, `weakref`, `objspace/memory` profiling tools
