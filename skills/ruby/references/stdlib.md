# Standard Library

Ruby's standard library (stdlib) — the modules and classes bundled with the interpreter but requiring an explicit `require`. Covers commonly used libraries for data, I/O, networking, utilities, and language features.

> **Source of truth**: <https://ruby-doc.org/stdlib/>. For **core classes** (String, Array, etc.), see [types.md](types.md). For **I/O and encoding**, see [io-encoding.md](io-encoding.md).

## Overview

The standard library is split into:
- **Core** — available without `require` (String, Array, Integer, etc.)
- **Stdlib** — bundled but requires `require 'name'` (json, csv, set, etc.)

Since Ruby 3.0+, many stdlib gems are bundled as **default gems** and can be updated independently via `gem install`.

## Data and Serialization

### json

```ruby
require 'json'
JSON.parse('{"a":1}')          # => {"a"=>1}
JSON.generate({a: 1})          # => '{"a":1}'
JSON.dump(obj)                 # alias for generate
{a: 1}.to_json                 # => '{"a":1}'  (requires 'json')
JSON.parse(str, symbolize_names: true)  # => {a: 1}
```

Options: `symbolize_names`, `create_additions` (deserialize custom types), `allow_nan`, `max_nesting`.

### yaml / psych

```ruby
require 'yaml'
YAML.load_file("config.yml")  # => Hash
YAML.dump(obj)                # => string
YAML.load(str)                # => parsed object
```

`Psych` is the underlying engine. `YAML.safe_load` restricts types to avoid arbitrary object instantiation (security).

### csv

```ruby
require 'csv'
CSV.read("data.csv")              # => [[...], ...]  (2D array)
CSV.read("data.csv", headers: true)  # => CSV::Table
CSV.foreach("data.csv") { |row| ... }
CSV.write("out.csv", data)
```

### set

```ruby
require 'set'
s = Set.new([1, 2, 3])
s.add(4)
s.include?(2)  # => true
s & Set.new([2, 3])  # intersection
s | Set.new([5, 6])  # union
s.to_a  # => [1, 2, 3, 4]  (insertion order, since Ruby 3.2)
```

`Set` is backed by `Hash`. Since Ruby 3.2, `Set` preserves insertion order.

## Filesystem and Path

### fileutils

```ruby
require 'fileutils'
FileUtils.cp("a.txt", "b.txt")
FileUtils.mv("a.txt", "dir/")
FileUtils.rm("a.txt")
FileUtils.mkdir_p("a/b/c")  # recursive, like mkdir -p
FileUtils.chdir("/tmp") { ... }
FileUtils.touch("file.txt")
```

All methods have `verbose: true` and `noop: true` options.

### pathname

```ruby
require 'pathname'
p = Pathname.new("/usr/local/bin/ruby")
p.dirname    # => #<Pathname:/usr/local/bin>
p.basename   # => #<Pathname:ruby>
p.extname    # => ""
p.parent     # => #<Pathname:/usr/local/bin>
p.children   # => [#<Pathname:...>, ...] (directory entries)
Pathname.glob("*.rb")
```

### tempfile

```ruby
require 'tempfile'
Tempfile.create("foo") do |f|
  f.write("data")
end  # file deleted on block exit

# Without block (must close/unlink manually)
f = Tempfile.new("foo")
f.write("data")
f.close
f.unlink
```

`Tempfile.create` (Ruby 3.0+) is preferred over `Tempfile.new` — it auto-cleans via `ensure` even on exceptions.

### tmpdir

```ruby
require 'tmpdir'
Dir.mktmpdir { |dir| ... }  # creates a temp directory, yields path, removes on exit
```

## Utilities

### monitor

```ruby
require 'monitor'
class Counter
  include MonitorMixin
  def initialize; @n = 0; super; end
  def increment; synchronize { @n += 1 }; end
end
```

`MonitorMixin` provides a reentrant mutex (unlike `Mutex`, which can deadlock on nested `synchronize`).

### singleton

```ruby
require 'singleton'
class Logger
  include Singleton
  def log(msg); puts msg; end
end
Logger.instance.log("hello")  # always the same instance
```

### forwardable

```ruby
require 'forwardable'
class Queue
  extend Forwardable
  def_delegators :@arr, :push, :pop, :length
  def initialize; @arr = []; end
end
```

### delegate

```ruby
require 'delegate'
class StringDecorator < SimpleDelegator
  def shout; upcase + "!"; end
end
sd = StringDecorator.new("hello")
sd.shout  # => "HELLO!"
sd.length # => 5  (delegated to the String)
```

### observer

```ruby
require 'observer'
class Thermometer
  include Observable
  def temp_changed(new_temp)
    changed
    notify_observers(new_temp)
  end
end
```

### logger

```ruby
require 'logger'
logger = Logger.new(STDOUT)
logger = Logger.new("app.log", "daily")  # rotate daily
logger.level = Logger::INFO
logger.info("started")
logger.error("failed: #{e.message}")
logger.formatter = proc { |s, t, p, msg| "[#{t}] #{s}: #{msg}\n" }
```

### optparse

```ruby
require 'optparse'
options = {}
OptionParser.new do |opts|
  opts.banner = "Usage: foo [options]"
  opts.on("-v", "--verbose", "Verbose mode") { |v| options[:verbose] = v }
  opts.on("-n NAME", "--name NAME", "Name") { |v| options[:name] = v }
  opts.on("-p", "--port PORT", Integer, "Port") { |v| options[:port] = v }
end.parse!(ARGV)
```

### pp / prettyprint

```ruby
require 'pp'
pp({ a: 1, b: { c: 2 } })  # pretty-prints with indentation
```

`pp` is available without `require` since Ruby 2.5, but `require 'pp'` gives the full feature set.

### benchmark

```ruby
require 'benchmark'
puts Benchmark.measure { 1000.times { "x" * 100 } }
# => 0.000123 0.000012 0.000135 ( 0.000150)
```

### digest

```ruby
require 'digest'
Digest::SHA256.hexdigest("hello")
Digest::MD5.hexdigest("hello")
Digest::SHA256.file("data.txt").hexdigest
```

### securerandom

```ruby
require 'securerandom'
SecureRandom.uuid           # => "..."
SecureRandom.hex(10)       # => "..."
SecureRandom.random_number(100)  # => 42
SecureRandom.urlsafe_base64
```

## Time and Date

### time

```ruby
require 'time'
Time.parse("2023-01-01 12:00:00")
Time.now.iso8601
Time.now.rfc2822
```

### date

```ruby
require 'date'
Date.today
Date.new(2023, 1, 1)
Date.parse("2023-01-01")
(Date.today..Date.today + 7).each { |d| ... }
```

## Math

### matrix

```ruby
require 'matrix'
m = Matrix[[1, 2], [3, 4]]
m * m
m.transpose
m.determinant
```

### tsort

```ruby
require 'tsort'
class Graph
  include TSort
  def initialize; @g = Hash.new { |h, k| h[k] = [] }; end
  def add_edge(a, b); @g[a] << b; end
  def tsort_each_node(&b); @g.each_key(&b); end
  def tsort_each_child(n, &b); @g[n].each(&b); end
end
```

Topological sort for dependency resolution.

## Networking

### socket

```ruby
require 'socket'
server = TCPServer.new("0.0.0.0", 3000)
client = server.accept
client.puts "hello"
client.close
```

### net/http

```ruby
require 'net/http'
require 'uri'
uri = URI("https://example.com/api")
response = Net::HTTP.get_response(uri)
Net::HTTP.post(uri, "data", "Content-Type" => "application/json")
```

### uri

```ruby
require 'uri'
uri = URI("https://user:pass@example.com:8080/path?q=1#frag")
uri.scheme, uri.host, uri.port, uri.path
URI.encode_www_form(q: "hello world")
```

### openssl

```ruby
require 'openssl'
digest = OpenSSL::Digest::SHA256.new
digest.update("hello")
digest.hexdigest
```

Provides SSL/TLS, crypto, and X.509 support.

## Text and Templates

### erb

See [cli.md](cli.md) for `erb` CLI and usage.

### cgi

```ruby
require 'cgi'
CGI.escape("hello world")  # => "hello+world"
CGI.unescape("hello+world")
CGI.escapeHTML("<script>")  # => "&lt;script&gt;"
```

## Debugging and Introspection

### objspace

```ruby
require 'objspace'
ObjectSpace.memsize_of(obj)
ObjectSpace.memsize_of_all(String)
ObjectSpace.trace_object_allocations_start
ObjectSpace.allocation_sourcefile(obj)
ObjectSpace.allocation_sourceline(obj)
```

### debug_inspector

```ruby
require 'debug_inspector'
DebugInspector.open { |i| i.frame_locals(0) }  # low-level frame access
```

### io/console

```ruby
require 'io/console'
IO.console.getpass  # password input (no echo)
IO.console.raw { ... }  # raw terminal mode
```

## Edge Cases and Known Issues

- **`json` and `Symbol`** — JSON has no Symbol type; use `symbolize_names: true` to parse keys as Symbols
- **`YAML.safe_load` is the default** since Ruby 3.0+ for `YAML.load` — but `YAML.unsafe_load` exists for backward compat
- **`Set` ordering** — pre-3.2 `Set#to_a` had undefined order; 3.2+ preserves insertion order
- **`Tempfile` may not be cleaned up on crash** — `Tempfile.create` with a block is safer
- **`optparse` mutates `ARGV` by default** — use `parse` (without `!`) to leave `ARGV` intact
- **`net/http` default timeout is large** — set `read_timeout`, `open_timeout` explicitly
- **`socket` + `Timeout.timeout`** — can leave sockets in a bad state; prefer non-blocking I/O or `IO.select`
- **`openssl` version mismatch** — `openssl` gem version must match the OpenSSL library installed

## Related

- [io-encoding.md](io-encoding.md) — `File`, `IO`, encoding
- [concurrency.md](concurrency.md) — `Monitor`, `Mutex`, `Queue`
- [gc-memory.md](gc-memory.md) — `objspace`, `ObjectSpace`
- [cli.md](cli.md) — `irb`, `rake`, `rdoc`, profiling tools
