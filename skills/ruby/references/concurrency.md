# Concurrency

Ruby's concurrency model — Threads, the GVL (Global VM Lock / GIL), Fibers, Ractors, Mutex, Queue, and the Fiber scheduler.

> **Source of truth**: <https://ruby-doc.org/core/Thread.html>, <https://github.com/ruby/ruby/blob/master/doc/fiber.md>, <https://github.com/ruby/ruby/blob/master/doc/ractor.md>. For **GC and memory**, see [gc-memory.md](gc-memory.md).

## The GVL (Global VM Lock)

MRI has a **Global VM Lock** (also called GIL — Global Interpreter Lock). At any moment, only one Thread can execute Ruby bytecode.

Implications:
- **CPU-bound parallelism is impossible** with pure Ruby threads in MRI
- **I/O operations release the GVL** — `read`, `write`, `sleep`, `select` release the lock, allowing other threads to run
- **C extensions can release the GVL** — via `rb_thread_call_without_gvl()` (e.g., heavy computation in native code)
- **JRuby and TruffleRuby have no GVL** — they achieve true parallelism

```ruby
threads = 4.times.map do
  Thread.new { fib(40) }  # CPU-bound: runs serially due to GVL
end
threads.each(&:join)  # total time ≈ 4 × single-thread time
```

```ruby
threads = 4.times.map do
  Thread.new { sleep 1 }  # I/O-bound: releases GVL, runs in parallel
end
threads.each(&:join)  # total time ≈ 1 second
```

## Thread

```ruby
t = Thread.new do
  puts "running"
end

t.status    # => false (completed) / nil (killed) / "run" / "sleep" / Integer
t.alive?    # => true if still running
t.join      # wait for completion
t.value     # wait and return the block's return value
t.priority  # => 0 (default)
t.priority = 2
```

### Thread.current and Thread.main

```ruby
Thread.current            # the current thread
Thread.main               # the main thread
Thread.list               # all active threads
ThreadGroup               # thread groups (mostly deprecated)
```

### Thread Variables

Each thread has its own storage:

```ruby
Thread.current[:request_id] = SecureRandom.uuid
Thread.current[:request_id]  # => "..." (thread-local)
```

Thread-local variables are isolated — child threads inherit a **copy** of the parent's variables at creation time, but changes in parent or child are independent.

### abort_on_exception

```ruby
Thread.abort_on_exception = true  # global default
# or per thread:
t = Thread.new { ... }
t.abort_on_exception = true
```

When `true`, an unrescued exception in the thread prints and exits the process. When `false` (default), the exception is stored and re-raised on `#join` / `#value`.

## Mutex

Protect shared mutable state:

```ruby
counter = 0
mutex = Mutex.new

threads = 10.times.map do
  Thread.new do
    mutex.synchronize { counter += 1 }
  end
end
threads.each(&:join)

counter  # => 10
```

| Method | Description |
|--------|-------------|
| `mutex.lock` | Acquire (blocks) |
| `mutex.unlock` | Release |
| `mutex.synchronize { }` | Acquire, run block, release (even on exception) |
| `mutex.try_lock` | Non-blocking acquire (returns `true`/`false`) |
| `mutex.owned?` | Does the current thread hold this lock? |

**Important**: Ruby's Mutex is **not reentrant** — calling `lock` or `synchronize` on a Mutex already held by the current thread will **deadlock**. For reentrant locking, use `Monitor`/`MonitorMixin` (see [stdlib.md](stdlib.md)).

## ConditionVariable

Wait for a condition, signal other threads:

```ruby
mutex = Mutex.new
cv = ConditionVariable.new
queue = []

producer = Thread.new do
  loop do
    mutex.synchronize do
      queue << produce_item
      cv.signal  # wake one waiter
    end
  end
end

consumer = Thread.new do
  loop do
    mutex.synchronize do
      cv.wait(mutex) while queue.empty?  # releases mutex while waiting
      item = queue.shift
    end
  end
end
```

`cv.wait(mutex)` atomically releases the mutex and sleeps; `signal` wakes one waiter; `broadcast` wakes all.

## Queue and SizedQueue

Thread-safe queues (from `require 'thread'`, built-in since Ruby 3.0):

```ruby
q = Queue.new          # unbounded
q.push(item)
q.pop                 # blocks if empty
q.empty?
q.size

sq = SizedQueue.new(5)  # bounded
sq.push(item)          # blocks if full
sq.pop                 # blocks if empty
```

`Queue` and `SizedQueue` are the idiomatic building blocks for producer/consumer patterns.

### Close

```ruby
q.close              # no more pushes
q.push(item)         # raises ClosedQueueError
q.pop                # returns nil if empty and closed (doesn't block forever)
```

`Queue#close` enables graceful shutdown: producers close the queue, consumers drain remaining items then get `nil`.

## Fiber

Fibers are cooperative coroutines — they yield control back to the caller without OS involvement:

```ruby
f = Fiber.new do
  puts "started"
  Fiber.yield 1
  puts "resumed"
  Fiber.yield 2
  puts "done"
  3
end

f.resume  # => prints "started", returns 1
f.resume  # => prints "resumed", returns 2
f.resume  # => prints "done", returns 3
f.resume  # => FiberError (can't resume dead fiber)
```

- `Fiber.yield` pauses the fiber and returns control to `resume`
- `resume` returns the value passed to `yield` (or the block's return on final resume)
- Fibers are **cooperative** — they only switch at `yield`/`resume`, so no data races within Ruby code
- Fibers are **not** threads — they share the GVL with the thread that created them

### Fiber with arguments

```ruby
f = Fiber.new do |first|
  puts "first: #{first}"
  second = Fiber.yield
  puts "second: #{second}"
end

f.resume("A")    # prints "first: A"
f.resume("B")    # prints "second: B"
```

The first `resume` passes its argument as the block's first argument. Subsequent `resume` arguments become `Fiber.yield`'s return value.

### Non-blocking Fibers and the Scheduler

Ruby 3.0+ introduced the **Fiber scheduler** — a mechanism for non-blocking I/O in fibers:

```ruby
require 'fiber'

scheduler = MyScheduler.new  # implement #io_wait, #io_read, #io_write, etc.
Fiber.set_scheduler(scheduler)

Fiber.schedule do
  # blocking I/O calls (sleep, read, write) are auto-awaited
  # by the scheduler without blocking the thread
end
```

When a fiber performs a blocking operation (like `sleep` or `IO#read`), the scheduler can switch to another ready fiber. This enables async I/O concurrency without threads.

The built-in `Timeout`, `Socket`, and `IO` classes integrate with the scheduler when one is set. Third-party gems like `async` provide scheduler implementations.

## Ractor

Ractors (Ruby Actors, Ruby 3.0+) are **parallel, isolated** workers that run without sharing mutable state:

```ruby
r = Ractor.new do
  msg = Ractor.receive  # receive a message
  Ractor.yield msg * 2  # send a result
end

r.send(21)
r.take  # => 42
```

- Each Ractor has its own GVL — they run truly in parallel
- Ractors cannot access each other's objects directly
- Communication is via message passing (`send`/`receive`, `yield`/`take`)
- Messages must be **shareable**: frozen objects, integers, symbols, floats, and specially-constructed data
- Ractors have their own object space — no shared global variables (except constants)

```ruby
data = [1, 2, 3].freeze  # frozen, so shareable

ractors = 4.times.map do |i|
  Ractor.new(data, i) do |data, i|
    data.sum + i  # computation in parallel
  end
end

ractors.map(&:take)  # => [6, 7, 8, 9]
```

### Ractor Limitations

- Cannot access non-shareable objects from outside (e.g. an unfrozen Array, a class instance variable)
- Global variables (`$xxx`) are not shared (except a few like `$stdin`)
- Constants are shared but must be frozen or marked shareable
- `Ractor.make_shareable(obj)` deep-freezes an object to make it shareable
- `Ractor.new`'s block can't access the caller's local variables (must be passed as args)

## When to Use What

| Need | Use |
|------|-----|
| I/O concurrency (network, file) | Threads (GVL releases on I/O) or Fibers + scheduler |
| CPU parallelism (pure Ruby) | Ractors |
| CPU parallelism (C extension) | Threads (extension releases GVL) |
| Cooperative multitasking | Fibers |
| Producer/consumer queue | `Queue` / `SizedQueue` |
| Protect shared state | `Mutex` + `ConditionVariable` |

## Edge Cases and Known Issues

- **GVL is per-thread** — threads on different Ractors each have their own GVL, allowing parallelism
- **Thread-local variables are copy-on-write** — child threads get a snapshot of parent's variables at creation
- **`Mutex` is not reentrant** — calling `synchronize` within `synchronize` on the same Mutex deadlocks; use `Monitor` for reentrant locking
- **`Fiber` cannot switch threads** — a fiber is bound to the thread that created it
- **Ractor message passing is copy-on-write** for non-shareable objects, which can be expensive
- **`Queue#pop` blocks forever** if the queue is empty and not closed — always call `close` on shutdown
- **`Thread.pass`** yields the GVL to other threads — useful for fairness but rarely needed
- **`Thread#report_on_exception`** (default `true` since Ruby 3.0) prints a warning if a thread dies with an unrescued exception

## Related

- [gc-memory.md](gc-memory.md) — GC and threads, finalizers and threads
- [io-encoding.md](io-encoding.md) — `IO.select`, non-blocking I/O
- [stdlib.md](stdlib.md) — `Monitor`, `Singleton`, `concurrent-ruby` integration
