# CLI Options

Complete reference for pandoc's command-line interface: general options, reader options, writer options, citation/math rendering flags, wrapper script options, and exit codes.

> **Source of truth**: <https://pandoc.org/MANUAL.html>. For format-specific details, see [formats.md](formats.md). For defaults file usage, see [defaults.md](defaults.md).

## Synopsis

```
pandoc [options] [input-file…]
```

If no input files are given, pandoc reads from stdin. If no `-o`/`--output` is given, pandoc writes to stdout. Multiple input files are concatenated before parsing (use `--file-scope` to parse each individually).

## General Options

| Option | Description |
|--------|-------------|
| `-f`, `-r`, `--from`, `--read` `FORMAT` | Input format. Append `+EXTENSION`/`-EXTENSION` to enable/disable extensions. |
| `-t`, `-w`, `--to`, `--write` `FORMAT` | Output format. |
| `-o`, `--output` `FILE` | Output file. |
| `--data-dir` `DIR` | User data directory (default: `~/.local/share/pandoc` on Unix, `%APPDATA%\pandoc` on Windows). |
| `-d`, `--defaults` `FILE` | YAML/JSON defaults file. See [defaults.md](defaults.md). |
| `--bash-completion` | Generate bash completion script. |
| `--sandbox` | Sandbox mode: limits file system and network access. |
| `--verbose` | Verbose debugging output. |
| `--quiet` | Suppress warnings. |
| `--fail-if-warnings` | Exit with error code 3 on warnings. |
| `--log` `FILE` | Write log messages as JSON lines to FILE. |
| `--list-input-formats` | Print supported input formats. |
| `--list-output-formats` | Print supported output formats. |
| `--list-extensions` `[FORMAT]` | List extensions for a format (`+` = enabled by default, `-` = disabled). |
| `--list-highlight-languages` | List supported syntax highlighting languages. |
| `--list-highlight-styles` | List highlight styles. |
| `-v`, `--version` | Print version. |
| `-h`, `--help` | Show usage. |

### Specifying Formats

If `--from` or `--to` is omitted, pandoc guesses from file extensions. For example, `pandoc input.md -o output.pdf` implies `--from markdown` and `--to pdf`.

### Character Encoding

Input and output are UTF-8. Pipe through `iconv` if working with other encodings.

### Reading from the Web

Absolute URIs can be given as input files; pandoc fetches via HTTP. Set custom headers with `--request-header`.

## Reader Options

| Option | Description |
|--------|-------------|
| `--shift-heading-level-by` `N` | Shift heading levels by N (positive or negative). `-1` promotes top-level headers to document title. |
| `--base-header-level` `N` | *(Deprecated)* Use `--shift-heading-level-by` instead. |
| `--indented-code-classes` `CLS` | Default classes for indented code blocks (space-separated). |
| `--default-image-extension` `EXT` | Default extension for image paths without one. |
| `--file-scope` | Parse each file individually before combining (instead of concatenating first). |
| `-F`, `--filter` `PROGRAM` | External JSON filter. See [filters.md](filters.md). |
| `-L`, `--lua-filter` `SCRIPT` | Lua filter script. See [filters.md](filters.md). |
| `-M`, `--metadata` `KEY[=VAL]` | Set metadata field. Boolean if no value. |
| `--metadata-file` `FILE` | Read metadata from YAML/JSON file. |
| `-p`, `--preserve-tabs` | Preserve tabs instead of converting to spaces. |
| `--tab-stop` `N` | Spaces per tab (default 4). |
| `--track-changes` `MODE` | Track changes handling: `accept` (default), `reject`, or `all`. |
| `--extract-media` `DIR` | Extract images/media to directory. |
| `--abbreviations` `FILE` | Custom abbreviations file. |
| `--typst-input` `KEY=VAL` | Set parameter for typst parser (`sys.inputs`). |
| `--trace` | Print diagnostic parser tracing output. |

## General Writer Options

| Option | Description |
|--------|-------------|
| `-s`, `--standalone` | Produce standalone document (with header/footer via template). |
| `--template` `FILE` | Custom template file. See [templates.md](templates.md). |
| `-V`, `--variable` `KEY[=VAL]` | Set template variable. Boolean if no value. |
| `--variable-json` `JSON` | Set template variable from JSON value. |
| `-D`, `--print-default-template` `FORMAT` | Print default template for FORMAT. |
| `--print-default-data-file` `FILE` | Print default data file. |
| `--eol` `MODE` | Line endings: `crlf`, `lf`, `native` (default). |
| `--dpi` `N` | Dots per inch (default 96). |
| `--wrap` `MODE` | Text wrapping: `auto` (default), `none`, `preserve`. |
| `--columns` `N` | Line length in characters (default 72). |
| `--toc`, `--table-of-contents` | Include table of contents. |
| `--toc-depth` `N` | Depth of TOC (default 3). |
| `--lof`, `--list-of-figures` | Include list of figures. |
| `--lot`, `--list-of-tables` | Include list of tables. |
| `--strip-comments` | Strip HTML comments. |
| `--syntax-highlighting` `SPEC` | Highlighting: `default`, `none`, `idiomatic`, `STYLE`, or `FILE`. |
| `--no-highlight` | *(Deprecated)* Use `--syntax-highlighting=none`. |
| `--highlight-style` | *(Deprecated)* Use `--syntax-highlighting`. |
| `--print-highlight-style` | Print JSON version of a highlight style. |
| `--syntax-definition` `FILE` | Load KDE XML syntax definition. |
| `-H`, `--include-in-header` `FILE` | Include file in header. |
| `-B`, `--include-before-body` `FILE` | Include file before body. |
| `-A`, `--include-after-body` `FILE` | Include file after body. |
| `--resource-path` `PATH` | Search paths for images/resources (colon-separated). |
| `--request-header` `HEADER` | Set HTTP request header. |
| `--no-check-certificate` | Disable certificate verification. |

## Writer-Specific Options

| Option | Description |
|--------|-------------|
| `--self-contained` | *(Deprecated)* Use `--embed-resources --standalone`. |
| `--embed-resources` | Embed resources via data: URIs (HTML only). |
| `--link-images` | Link instead of embed images (ODT only). |
| `--html-q-tags` | Use `<q>` tags for quotes in HTML. |
| `--ascii` | Use ASCII only in output. |
| `--reference-links` | Use reference-style links (Markdown). |
| `--reference-location` `LOC` | Position of notes: `block`, `section`, `document`. |
| `--figure-caption-position` `POS` | `above` or `below`. |
| `--table-caption-position` `POS` | `above` or `below`. |
| `--markdown-headings` `STYLE` | `setext` or `atx`. |
| `--list-tables` | Render as list tables (RST). |
| `--top-level-division` `DIV` | `default`, `section`, `chapter`, or `part`. |
| `-N`, `--number-sections` | Number section headings. |
| `--number-offset` `N…` | Offsets for section numbers. |
| `--listings` | *(Deprecated)* Use `listings` package for LaTeX code. |
| `-i`, `--incremental` | Incremental list display in slides. |
| `--slide-level` `N` | Heading level for slide breaks. |
| `--section-divs` | Wrap sections in `<section>` tags (HTML). |
| `--email-obfuscation` `MODE` | `none`, `javascript`, or `references`. |
| `--id-prefix` `PREFIX` | Prefix for identifiers. |
| `-T`, `--title-prefix` `PREFIX` | Title prefix for HTML header. |
| `-c`, `--css` `URL` | Link to CSS stylesheet. |
| `--reference-doc` `FILE` | Style reference file (docx/ODT/pptx). |
| `--split-level` `N` | Heading level for splitting EPUB/chunked HTML. |
| `--epub-chapter-level` `N` | *(Deprecated)* Use `--split-level`. |
| `--epub-cover-image` `FILE` | Cover image for EPUB. |
| `--epub-title-page` `BOOL` | Include title page (true/false). |
| `--epub-metadata` `FILE` | Dublin Core metadata XML file. |
| `--epub-embed-font` `FILE` | Embed font in EPUB. |
| `--epub-subdirectory` `DIR` | Subdirectory in OCF container. |
| `--ipynb-output` `MODE` | `all`, `none`, or `best`. |
| `--pdf-engine` `PROGRAM` | PDF engine. See [pdf.md](pdf.md). |
| `--pdf-engine-opt` `OPT` | Command-line argument to PDF engine. |

## Citation Rendering Options

| Option | Description |
|--------|-------------|
| `-C`, `--citeproc` | Process citations and add bibliography. See [citations.md](citations.md). |
| `--bibliography` `FILE` | Bibliography file. |
| `--csl` `FILE` | CSL citation style file. |
| `--citation-abbreviations` `FILE` | Journal abbreviations JSON file. |
| `--natbib` | Use natbib for LaTeX citations. |
| `--biblatex` | Use biblatex for LaTeX citations. |

## Math Rendering Options (HTML)

| Option | Description |
|--------|-------------|
| `--mathjax` | Use MathJax. |
| `--mathml` | Convert to MathML. |
| `--webtex` `[URL]` | Use external service for formula images. |
| `--katex` | Use KaTeX. |
| `--gladtex` | Use GladTeX. |

## Wrapper Script Options

| Option | Description |
|--------|-------------|
| `--dump-args` | Print command-line argument info to stdout. |
| `--ignore-args` | Ignore command-line arguments. |

## Exit Codes

| Code | Error |
|------|-------|
| 0 | Success |
| 1 | PandocIOError |
| 3 | PandocFailOnWarningError |
| 4 | PandocAppError |
| 5 | PandocTemplateError |
| 6 | PandocOptionError |
| 21 | PandocUnknownReaderError |
| 22 | PandocUnknownWriterError |
| 23 | PandocUnsupportedExtensionError |
| 24 | PandocCiteprocError |
| 25 | PandocBibliographyError |
| 31 | PandocEpubSubdirectoryError |
| 43 | PandocPDFError |
| 44 | PandocXMLError |
| 47 | PandocPDFProgramNotFoundError |
| 61 | PandocHttpError |
| 62 | PandocShouldNeverHappenError |
| 63 | PandocSomeError |
| 64 | PandocParseError |
| 66 | PandocMakePDFError |
| 67 | PandocSyntaxMapError |
| 83 | PandocFilterError |
| 84 | PandocLuaError |
| 89 | PandocNoScriptingEngine |
| 91 | PandocMacroLoop |
| 92 | PandocUTF8DecodingError |
| 93 | PandocIpynbDecodingError |
| 94 | PandocUnsupportedCharsetError |
| 95 | PandocInputNotTextError |
| 97 | PandocCouldNotFindDataFileError |
| 98 | PandocCouldNotFindMetadataFileError |
| 99 | PandocResourceNotFound |

## Related References

- [formats.md](formats.md) — Input/output format lists and extension system
- [defaults.md](defaults.md) — Defaults files for persistent option sets
- [templates.md](templates.md) — Template system and variables
- [filters.md](filters.md) — Filter types and pipeline
- [pdf.md](pdf.md) — PDF generation specifics
