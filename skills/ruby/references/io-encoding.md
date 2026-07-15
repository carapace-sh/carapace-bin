# I/O and Encoding

Ruby's I/O system — `IO`, `File`, `StringIO`, `ARGF`, `STDIN`/`STDOUT`/`STDERR`, and the `Encoding` system including default external/internal encodings, transcoding, magic comments, and frozen string literals.

> **Source of truth**: <https://ruby-doc.org/core/IO.html>, <https://ruby-doc.org/core/Encoding.html>. For **String methods**, see [types.md](types.md).

## IO and File

### Reading

```ruby
# Entire file
content = File.read("path.txt")
lines = File.readlines("path.txt")  # Array of lines (with newlines)
File.foreach("path.txt") { |line| puts line }  # streaming, one line at a time

# Explicit IO
File.open("path.txt") do |f|
  f.each_line { |line| process(line) }
end
```

`File.open` with a block auto-closes the IO when the block exits. Without a block, returns an open File that you must close manually.

### Writing

```ruby
File.write("out.txt", "hello")          # overwrites
File.open("out.txt", "w") { |f| f.puts "hello" }
File.open("out.txt", "a") { |f| f.puts "appended" }  # append mode
```

### Modes

| Mode | Description |
|------|-------------|
| `r` | Read (default), start at beginning |
| `r+` | Read/write, start at beginning |
| `w` | Write (truncate), create if missing |
| `w+` | Read/write (truncate) |
| `a` | Append, create if missing |
| `a+` | Read/append |
| `b` | Binary mode (Windows: no CRLF translation) |
| `t` | Text mode (Windows: default) |
| `x` | Exclusive create (fail if exists) |

### Binary mode

On non-Windows, `b` has no effect. On Windows, `b` prevents CRLF <-> LF translation. Use `binmode` for I/O:

```ruby
File.open("data.bin", "wb") { |f| f.write("\x00\x01\x02") }
```

## ARGF and ARGV

`ARGV` is the array of command-line arguments. `ARGF` is a virtual stream that reads from files listed in `ARGV`, or from STDIN if no files are given:

```ruby
# Print all lines from files given as arguments, or from STDIN
ARGF.each_line { |line| puts line }

# Skip to next file
ARGF.argv  # current ARGV

# Get current filename
ARGF.filename
```

## STDIN, STDOUT, STDERR

```ruby
STDOUT.puts "normal output"
STDERR.puts "error output"

STDOUT.sync = true  # unbuffered output (flush after every write)

# Redirect
$stdout = File.open("log.txt", "w")
puts "goes to log.txt"
```

`$stdin`, `$stdout`, `$stderr` are the global aliases. `STDIN`/`STDOUT`/`STDERR` are constants (originals).

## StringIO

`StringIO` wraps a String in an IO-like interface (from `require 'stringio'`):

```ruby
require 'stringio'
io = StringIO.new("hello\nworld\n")
io.gets  # => "hello\n"
io.gets  # => "world\n"
```

Useful for testing code that expects an IO, without touching the filesystem.

## Encoding System

### String and Encoding

Every String has an associated `Encoding`:

```ruby
"hello".encoding            # => #<Encoding::UTF-8>
"hello".bytesize            # => 5 (bytes)
"hello".length              # => 5 (characters)
"héllo".bytesize            # => 6 (é is 2 bytes in UTF-8)
"héllo".length              # => 5 (characters)
```

### Default External and Internal

```ruby
Encoding.default_external  # => #<Encoding::UTF-8> (from locale)
Encoding.default_internal  # => nil (no transcoding by default)
```

- **default_external** — the encoding used when reading files, talking to STDIN/STDOUT, etc.
- **default_internal** — if set, IO automatically transcodes from external to internal when reading

```ruby
Encoding.default_external = Encoding::UTF_8
Encoding.default_internal = Encoding::UTF_8

# Or via command line:
# ruby -E utf-8:utf-8 script.rb
```

### Magic Comments

The first line (or second, after shebang) can specify encoding:

```ruby
# encoding: utf-8
# coding: utf-8      (alias)
# -*- coding: utf-8 -*-  (Emacs-style)
```

Default is UTF-8 since Ruby 2.0.

### frozen_string_literal magic comment

```ruby
# frozen_string_literal: true

s = "hello"
s << "!"  # => RuntimeError: can't modify frozen String
```

This makes **all string literals** in the file frozen. To get a mutable copy, use `+`:

```ruby
s = "hello" + ""  # mutable copy
s << "!"  # OK
```

`String.new` also returns an unfrozen string regardless of the magic comment.

### Transcoding

```ruby
# Convert a String's encoding
"hello".encode("UTF-8")           # returns a new String
"hello".encode!("UTF-8")          # mutates (if compatible)

# Force encoding without conversion (re-interpret bytes)
"\xe9".force_encoding("UTF-8")    # treat bytes as UTF-8

# With options
invalid_utf8.encode("UTF-8", invalid: :replace, replace: "?")
```

- `encode` converts bytes from one encoding to another (may raise on invalid)
- `force_encoding` just changes the label without touching bytes (fast, dangerous)
- `String#bytes` returns raw bytes regardless of encoding

### IO Encoding

When opening a file, specify encodings:

```ruby
File.open("data.txt", "r:utf-8")           # external encoding
File.open("data.txt", "r:utf-8:utf-8")     # external:internal (transcode on read)
File.open("data.txt", "r:UTF-32")          # read UTF-32
```

Or via `set_encoding`:

```ruby
io.set_encoding("UTF-8", "UTF-8")  # external, internal
io.set_encoding_by_bom              # detect BOM
```

### BOM (Byte Order Mark)

UTF-16/32 files may have a BOM. Open with `bom|UTF-8` to auto-detect:

```ruby
File.open("file.txt", "r:bom|utf-8")
```

## File and Pathname

`Pathname` (stdlib) provides a higher-level path API:

```ruby
require 'pathname'
Pathname.new("/usr/local/bin").parent      # => #<Pathname:/usr/local>
Pathname.new("foo/bar").join("baz")        # => #<Pathname:foo/bar/baz>
Pathname.glob("*.rb")                       # => [#<Pathname:foo.rb>, ...]
```

`File` has class methods for filesystem operations: `File.exist?`, `File.size`, `File.mtime`, `File.extname`, `File.basename`, `File.dirname`, `File.join`, `File.expand_path`.

For filesystem operations (copy, mkdir, chmod), see `FileUtils` in [stdlib.md](stdlib.md).

## IO.select and Non-blocking I/O

```ruby
io = IO.pipe  # => [read_end, write_end]
read, = IO.select([io[0]], nil, nil, 5)  # 5-second timeout
if read
  data = io[0].gets_nonblock
end
```

- `IO.select` waits for readability/writability/exceptions on arrays of IOs
- `read_nonblock` / `write_nonblock` raise `IO::WaitReadable` / `IO::WaitWritable` when the operation would block
- `IO::pipe` creates a pair of connected IO objects

## flock (File Locking)

```ruby
File.open("file.lock") do |f|
  f.flock(File::LOCK_EX)  # exclusive lock
  # do work
end  # lock released on close
```

| Lock | Meaning |
|------|---------|
| `LOCK_SH` | Shared (read) lock |
| `LOCK_EX` | Exclusive (write) lock |
| `LOCK_UN` | Unlock |
| `LOCK_NB` | Non-blocking (combined with SH/EX via `\|`) |

## Edge Cases and Known Issues

- **`\r\n` on Windows** — text mode translates newlines; use binary mode (`b`) for binary data
- **Encoding mismatch raises `Encoding::CompatibilityError`** — e.g. concatenating a UTF-8 string with an ASCII-8BIT string
- **`force_encoding` is dangerous** — it doesn't convert bytes, just relabels; use `encode` for real conversion
- **`String#length` counts characters, not bytes** — for byte count use `bytesize`
- **`File.read` loads the entire file** — for large files, use `File.foreach` or `File.open` with `each_line`
- **`ARGF` in `-n`/`-p` mode** — Ruby's `-n` flag wraps the script in `while gets ... end`, with `ARGV`/`ARGF` as the source
- **`STDOUT.sync = true` is global** — affects all default output; consider `f.sync = true` per-IO

## Related

- [types.md](types.md) — String methods and encoding-related behavior
- [stdlib.md](stdlib.md) — `FileUtils`, `Pathname`, `StringIO`, `Tempfile`
- [environment.md](environment.md) — `RUBYOPT`, encoding flags
