# Architecture and Internals

Pandoc's internal architecture: the AST, reader-writer pipeline, PandocMonad typeclass, module structure, and key supporting packages.

> **Source of truth**: <https://pandoc.org/using-the-pandoc-api.html>, <https://hackage.haskell.org/package/pandoc-types>, and the [source code](https://github.com/jgm/pandoc).

## Overview

Pandoc is a Haskell library and CLI tool for converting between document markup formats. Its architecture is based on a **reader-writer pipeline** with a universal **abstract syntax tree (AST)** as the intermediate representation.

```
input text
  → reader (format-specific parser)
    → Pandoc AST (Meta + [Block])
      → filters (JSON / Lua / citeproc, applied sequentially)
        → internal transforms
          → writer (format-specific renderer)
            → template wrapping (standalone mode)
              → output text
```

This design gives M+N conversion support: only M readers + N writers are needed to convert between any pair of M input and N output formats.

## The Pandoc AST

The AST is the central data structure that every reader produces and every writer consumes. It is defined in the separate **`pandoc-types`** package (`Text.Pandoc.Definition`), allowing other packages to use the AST without depending on all of pandoc.

### Top-Level Type

```haskell
data Pandoc = Pandoc Meta [Block]
```

A `Pandoc` document is metadata plus a list of block-level elements.

### Metadata

```haskell
newtype Meta = Meta { unMeta :: Map Text MetaValue }

data MetaValue
  = MetaMap (Map Text MetaValue)
  | MetaList [MetaValue]
  | MetaBool Bool
  | MetaString Text
  | MetaInlines [Inline]
  | MetaBlocks [Block]
```

### Block-Level Elements

```haskell
data Block
  = Plain [Inline]
  | Para [Inline]
  | LineBlock [[Inline]]
  | CodeBlock Attr Text
  | RawBlock Format Text
  | BlockQuote [Block]
  | OrderedList ListAttributes [[Block]]
  | BulletList [[Block]]
  | DefinitionList [([Inline], [[Block]])]
  | Header Int Attr [Inline]
  | HorizontalRule
  | Table Attr Caption [ColSpec] TableHead [TableBody] TableFoot
  | Figure Attr Caption [Block]
  | Div Attr [Block]
```

### Inline-Level Elements

```haskell
data Inline
  = Str Text
  | Emph [Inline]
  | Underline [Inline]
  | Strong [Inline]
  | Strikeout [Inline]
  | Superscript [Inline]
  | Subscript [Inline]
  | SmallCaps [Inline]
  | Quoted QuoteType [Inline]
  | Cite [Citation] [Inline]
  | Code Attr Text
  | Space
  | SoftBreak
  | LineBreak
  | Math MathType Text
  | RawInline Format Text
  | Link Attr [Inline] Target
  | Image Attr [Inline] Target
  | Note [Block]
  | Span Attr [Inline]
```

### Supporting Types

| Type | Definition | Purpose |
|------|------------|---------|
| `Attr` | `(Text, [Text], [(Text, Text)])` | Identifier, classes, key-value pairs |
| `Format` | `Format Text` | Format name (e.g., "html", "latex") |
| `Target` | `(Text, Text)` | (URL, title) for links/images |
| `ListAttributes` | `(Int, ListNumberStyle, ListNumberDelim)` | Ordered list start number, style, delimiter |
| `QuoteType` | `SingleQuote \| DoubleQuote` | Quotation type |
| `MathType` | `DisplayMath \| InlineMath` | Math display mode |
| `Citation` | Record with `citationId`, `citationPrefix`, `citationSuffix`, `citationMode`, etc. | Citation data |

### AST Traversal (Walk Typeclass)

The `Walkable` typeclass in `Text.Pandoc.Walk` (part of `pandoc-types`) provides generic traversal:

```haskell
walk :: (a -> a) -> b -> b          -- non-monadic, bottom-up transform
walkM :: (Monad m) => (a -> m a) -> b -> m b  -- monadic version
query :: Monoid c => (a -> c) -> b -> c       -- information extraction
```

This allows filters and transforms to modify specific element types without manually recursing through the entire tree.

### AST Construction (Builder)

`Text.Pandoc.Builder` provides monoid-based construction:

```haskell
type Inlines = Many Inline
type Blocks = Many Block

-- Monoid instances allow easy composition
body :: Blocks
body = para (str "Hello " <> emph (str "world"))
```

## PandocMonad Typeclass

All readers, writers, and transformations are parameterized over the `PandocMonad` typeclass, which abstracts over side effects:

```haskell
class (Functor m, Applicative m, Monad m) => PandocMonad m where
    getVerbosity :: m Verbosity
    setVerbosity :: Verbosity -> m ()
    getLog :: m [LogMessage]
    report :: LogMessage -> m ()
    fetchItem :: Text -> m (ByteString, Maybe MimeType)
    setResourcePath :: [FilePath] -> m ()
    -- ... and more
```

### Two Principal Instances

| Instance | Effects | Runner |
|----------|---------|--------|
| `PandocIO` | Real I/O (file reads, network) | `runIO :: PandocIO a -> IO a` |
| `PandocPure` | Pure/sandboxed, no side effects | `runPure :: PandocPure a -> a` |

`PandocPure` is useful for web applications and sandboxed environments where malicious behavior must be prevented. It supports no file system or network access.

## Reader-Writer Pipeline

The full pipeline is orchestrated by `convertWithOpts` in `Text.Pandoc.App`:

1. **Input handling** — Read input from files, stdin, or HTTP URIs
2. **Reader dispatch** — `getReader` in `Text.Pandoc.Readers` maps the format string to the reader function
3. **Parsing** — The reader parses input text into a `Pandoc` AST
4. **Metadata processing** — Metadata from CLI flags (`--metadata`), metadata files, and defaults files is merged into the document's `Meta` block
5. **Filter application** — Filters (JSON, Lua, citeproc) are applied sequentially via `applyFilters`
6. **Internal transforms** — Built-in transforms: `headerShift`, `eastAsianLineBreakFilter`, `filterToBeamer`, `flatten`, etc.
7. **Writer dispatch** — `getWriter` in `Text.Pandoc.Writers` maps the format string to the writer function
8. **Rendering** — The writer converts the AST into the target format text
9. **Template wrapping** — If `--standalone`, the output is wrapped in a template (see [templates.md](templates.md))
10. **PDF generation** — For PDF output, the intermediate format is compiled by an external engine (see [pdf.md](pdf.md))

## Source Code Organization

Pandoc is a multi-package project:

```
pandoc/
├── src/Text/Pandoc/              # Main library (pandoc package)
│   ├── App/                      # CLI orchestration
│   │   ├── Opt.hs                # Opt data type (CLI options)
│   │   ├── OutputSettings.hs     # Writer configuration
│   │   └── CommandLineOptions.hs # CLI parsing
│   ├── Readers/                  # Input format readers
│   │   ├── Markdown/             # Markdown reader (sub-modules)
│   │   ├── LaTeX.hs
│   │   ├── HTML.hs
│   │   ├── Docx.hs
│   │   ├── Org.hs
│   │   └── ... (~20+ readers)
│   ├── Writers/                  # Output format writers
│   │   ├── HTML.hs
│   │   ├── LaTeX.hs
│   │   ├── Markdown/             # Markdown writer (sub-modules)
│   │   ├── Docx.hs
│   │   ├── EPUB.hs
│   │   └── ... (~30+ writers)
│   ├── Class/                    # PandocMonad, CommonState, IO
│   ├── Filter/                   # Filter system
│   ├── Templates.hs              # Template system
│   ├── Options.hs                # ReaderOptions, WriterOptions
│   ├── Extensions.hs             # Extension system
│   ├── Logging.hs                # LogMessage, Verbosity
│   ├── Error.hs                  # PandocError
│   ├── PDF.hs                    # PDF generation
│   ├── Parsing.hs                # Common parsing infrastructure
│   └── MediaBag.hs               # Media resource management
├── pandoc-cli/                   # CLI binary (separate package)
├── pandoc-lua-engine/            # Lua filter engine (separate package)
├── pandoc-server/                # HTTP server (separate package)
├── wasm/                         # WebAssembly build
├── data/templates/               # Default templates per format
├── test/                         # Test suite
└── doc/                          # Documentation
```

## Key Modules

| Module | Package | Purpose |
|--------|---------|---------|
| `Text.Pandoc` | pandoc | Main public API, re-exports |
| `Text.Pandoc.Definition` | pandoc-types | AST types (`Pandoc`, `Block`, `Inline`, `Meta`, `Attr`) |
| `Text.Pandoc.Builder` | pandoc-types | Monoid-based AST construction (`Inlines`, `Blocks`) |
| `Text.Pandoc.Walk` | pandoc-types | Generic AST traversal (`walk`, `walkM`, `query`) |
| `Text.Pandoc.Class` | pandoc | `PandocMonad` typeclass, `PandocIO`, `PandocPure`, `runIO`, `runPure` |
| `Text.Pandoc.Options` | pandoc | `ReaderOptions`, `WriterOptions`, `Extensions` |
| `Text.Pandoc.Readers` | pandoc | Reader dispatch (`getReader`) |
| `Text.Pandoc.Writers` | pandoc | Writer dispatch (`getWriter`) |
| `Text.Pandoc.Parsing` | pandoc | Common parsing infrastructure (ParsecT, ParserState) |
| `Text.Pandoc.Writers.Shared` | pandoc | Shared writer utilities (tables, template context) |
| `Text.Pandoc.Templates` | pandoc | Template retrieval, compilation, rendering |
| `Text.Pandoc.Filter` | pandoc | Filter orchestration (`applyFilters`) |
| `Text.Pandoc.App` | pandoc | CLI orchestration (`convertWithOpts`) |
| `Text.Pandoc.Error` | pandoc | `PandocError` sum type |
| `Text.Pandoc.Logging` | pandoc | `LogMessage`, `Verbosity` |
| `Text.Pandoc.Extensions` | pandoc | Extension system (`Ext_smart`, `Ext_citations`, etc.) |
| `Text.Pandoc.PDF` | pandoc | PDF generation via external engines |
| `Text.Pandoc.MediaBag` | pandoc | Media resource management |

## Supporting Packages

| Package | Purpose |
|---------|---------|
| `pandoc-types` | AST definition, Builder, Walk — used by external packages |
| `citeproc` | Citation processing (CSL) |
| `texmath` | Math formula rendering (LaTeX, MathML, OMML, etc.) |
| `doctemplates` | Template engine used by `Text.Pandoc.Templates` |
| `skylighting` | Syntax highlighting |
| `pandoc-lua-engine` | Embedded Lua interpreter for filters |

## Using Pandoc as a Library

```haskell
import Text.Pandoc

-- Convert markdown to HTML
convertToHtml :: Text -> IO Text
convertToHtml input = do
  result <- runIO $ do
    doc <- readMarkdown def { readerExtensions = pandocExtensions } input
    writeHtml5String def doc
  case result of
    Right html -> return html
    Left err   -> error (show err)

-- Pure conversion (no I/O)
convertPure :: Text -> Text
convertPure input =
  case runPure $ do
    doc <- readMarkdown def { readerExtensions = pandocExtensions } input
    writeHtml5String def doc of
    Right html -> html
    Left err   -> error (show err)
```

## Pandoc Server

Pandoc can run as an HTTP server with a JSON API:

```bash
pandoc-server              # Start server
pandoc server              # Alternative invocation
pandoc-server.cgi          # CGI mode
```

The server exposes most conversion functionality via JSON requests.

## Security Model

1. **Filters can perform arbitrary file operations** — JSON filters are external programs; Lua filters have `io`/`os` access
2. **`--sandbox`** limits file system and network access for Lua filters
3. **`PandocPure` monad** provides full isolation in library mode (no I/O)
4. **Include directives** in LaTeX/Org/RST/Typst can expose file contents
5. **Embedded images** in output can expose local files
6. **HTML `iframe` elements** can include local files
7. **Parsers can exhibit pathological performance** — use `commonmark` reader for untrusted input, apply timeouts
8. **HTML output is not sanitized** — sanitize before serving to prevent XSS

## Related References

- [formats.md](formats.md) — Reader/writer modules and format dispatch
- [filters.md](filters.md) — Filter pipeline, `applyFilters`, Walk typeclass usage
- [templates.md](templates.md) — Template system, `metaToContext`, `doctemplates`
- [cli-options.md](cli-options.md) — `Opt` data type and CLI parsing
- [pdf.md](pdf.md) — `Text.Pandoc.PDF` and PDF engine integration
