# Input and Output Formats

Pandoc's format system: supported readers and writers, the extension mechanism, format dispatch, and Markdown variants.

> **Source of truth**: <https://pandoc.org/MANUAL.html>. For the internal reader/writer dispatch architecture, see [architecture.md](architecture.md). For Markdown syntax details, see [markdown.md](markdown.md).

## Design Principle

Pandoc uses a reader-writer architecture: each input format has a **reader** that parses text into the Pandoc AST, and each output format has a **writer** that renders the AST into target text. This gives M+N conversion support (M readers + N writers) rather than M*N pairwise converters.

```
input format --reader--> Pandoc AST --writer--> output format
```

Conversions from pandoc's Markdown to all formats aspire to be perfect. Conversions from more expressive formats (e.g., LaTeX, docx) can be lossy because the AST is a "least common denominator" representation.

## Input Formats (Readers)

| Format | Description |
|--------|-------------|
| `ascii` | *(output only)* |
| `asciidoc` | AsciiDoc |
| `bibtex` | BibTeX |
| `biblatex` | BibLaTeX |
| `bits` | Alias for `jats` |
| `commonmark` | CommonMark |
| `commonmark_x` | CommonMark with extensions |
| `creole` | Creole 1.0 |
| `csljson` | CSL JSON |
| `csv` | Comma-separated values |
| `tsv` | Tab-separated values |
| `djot` | Djot markup |
| `docbook` | DocBook 4 |
| `docx` | Microsoft Word DOCX |
| `dokuwiki` | DokuWiki markup |
| `endnotexml` | EndNote XML |
| `epub` | EPUB |
| `fb2` | FictionBook2 |
| `gfm` | GitHub-Flavored Markdown |
| `haddock` | Haddock markup |
| `html` | HTML |
| `ipynb` | Jupyter notebook |
| `jats` | JATS XML |
| `jira` | Jira/Confluence wiki markup |
| `json` | Pandoc JSON (native AST serialization) |
| `latex` | LaTeX |
| `markdown` | Pandoc Markdown |
| `markdown_mmd` | MultiMarkdown |
| `markdown_phpextra` | PHP Markdown Extra |
| `markdown_strict` | Original Markdown |
| `mediawiki` | MediaWiki markup |
| `man` | Man page (roff) |
| `mdoc` | mdoc (roff) |
| `muse` | Emacs Muse |
| `native` | Pandoc native (AST) |
| `odt` | OpenDocument Text |
| `opml` | OPML |
| `org` | Emacs Org mode |
| `pod` | Perl POD |
| `pptx` | PowerPoint PPTX |
| `ris` | RIS bibliography |
| `rtf` | Rich Text Format |
| `rst` | reStructuredText |
| `t2t` | txt2tags |
| `textile` | Textile |
| `tikiwiki` | TikiWiki markup |
| `twiki` | TWiki markup |
| `typst` | Typst |
| `vimwiki` | Vimwiki |
| `xlsx` | Excel XLSX |
| `xml` | Generic XML |

Custom Lua readers can be placed in the `custom` subdirectory of the data directory.

## Output Formats (Writers)

| Format | Description |
|--------|-------------|
| `ansi` | ANSI-formatted terminal output |
| `asciidoc` | AsciiDoc |
| `asciidoc_legacy` | Legacy AsciiDoc |
| `asciidoctor` | *(Deprecated)* Use `asciidoc` |
| `bbcode` | BBCode |
| `bbcode_fluxbb` | FluxBB BBCode |
| `bbcode_phpbb` | phpBB BBCode |
| `bbcode_steam` | Steam BBCode |
| `bbcode_hubzilla` | Hubzilla BBCode |
| `bbcode_xenforo` | XenForo BBCode |
| `beamer` | LaTeX Beamer slides |
| `bibtex` | BibTeX |
| `biblatex` | BibLaTeX |
| `chunkedhtml` | Chunked HTML (zip archive) |
| `commonmark` | CommonMark |
| `commonmark_x` | CommonMark with extensions |
| `context` | ConTeXt |
| `csljson` | CSL JSON |
| `djot` | Djot markup |
| `docbook` | DocBook 4 (alias: `docbook4`) |
| `docbook5` | DocBook 5 |
| `docx` | Microsoft Word DOCX |
| `dokuwiki` | DokuWiki markup |
| `epub` | EPUB v3 (alias: `epub3`) |
| `epub2` | EPUB v2 |
| `fb2` | FictionBook2 |
| `gfm` | GitHub-Flavored Markdown |
| `haddock` | Haddock markup |
| `html` | HTML5 (alias: `html5`) |
| `html4` | HTML4 |
| `icml` | InDesign ICML |
| `ipynb` | Jupyter notebook |
| `jats` | JATS XML (alias: `jats_archiving`) |
| `jats_articleauthoring` | JATS articleauthoring |
| `jats_publishing` | JATS publishing |
| `jira` | Jira/Confluence wiki markup |
| `json` | Pandoc JSON |
| `latex` | LaTeX |
| `man` | Man page (roff) |
| `markdown` | Pandoc Markdown |
| `markdown_mmd` | MultiMarkdown |
| `markdown_phpextra` | PHP Markdown Extra |
| `markdown_strict` | Original Markdown |
| `markua` | Markua |
| `mediawiki` | MediaWiki markup |
| `ms` | roff ms |
| `muse` | Emacs Muse |
| `native` | Pandoc native (AST) |
| `odt` | OpenDocument Text |
| `opml` | OPML |
| `opendocument` | OpenDocument |
| `org` | Emacs Org mode |
| `pdf` | PDF (via external engine, see [pdf.md](pdf.md)) |
| `plain` | Plain text |
| `pptx` | PowerPoint PPTX |
| `rst` | reStructuredText |
| `rtf` | Rich Text Format |
| `texinfo` | Texinfo |
| `textile` | Textile |
| `slideous` | Slideous slides |
| `slidy` | Slidy slides |
| `dzslides` | DZSlides |
| `revealjs` | reveal.js slides |
| `s5` | S5 slides |
| `tei` | TEI XML |
| `typst` | Typst |
| `vimdoc` | Vim help file |
| `xml` | Generic XML |
| `xwiki` | XWiki markup |
| `zimwiki` | Zim wiki markup |

Custom Lua writers can be placed in the `custom` subdirectory of the data directory.

## Extension System

Extensions modify how a reader or writer interprets a format. They are enabled/disabled by appending `+EXTENSION` or `-EXTENSION` to the format name:

```bash
pandoc -f markdown+smart -t html
pandoc -f markdown-smart -t latex
pandoc -f markdown+pipe_tables+yaml_metadata_block -t html
```

List available extensions for a format:

```bash
pandoc --list-extensions=markdown
```

Output shows `+name` for extensions enabled by default, `-name` for disabled.

### Key Extension Categories

| Category | Representative Extensions |
|----------|--------------------------|
| Typography | `smart` (curly quotes, em-dashes, ellipses) |
| Headings/IDs | `auto_identifiers`, `gfm_auto_identifiers`, `ascii_identifiers` |
| Math input | `tex_math_dollars`, `tex_math_gfm`, `tex_math_single_backslash`, `tex_math_double_backslash` |
| Raw passthrough | `raw_html`, `raw_tex`, `native_divs`, `native_spans` |
| Literate Haskell | `literate_haskell` (`+lhs`) |
| Lists | `task_lists`, `fancy_lists`, `startnum`, `example_lists`, `definition_lists` |
| Tables | `pipe_tables`, `grid_tables`, `simple_tables`, `multiline_tables`, `table_captions` |
| Links | `autolink_bare_uris`, `shortcut_reference_links`, `spaced_reference_links`, `link_attributes` |
| Footnotes | `footnotes`, `inline_notes` |
| Code blocks | `fenced_code_blocks`, `backtick_code_blocks`, `fenced_code_attributes`, `inline_code_attributes` |
| Divs/spans | `fenced_divs`, `bracketed_spans`, `raw_attribute` |
| Line breaks | `hard_line_breaks`, `ignore_line_breaks`, `east_asian_line_breaks`, `escaped_line_breaks` |
| Other | `emoji`, `alerts`, `gutenberg`, `rebase_relative_paths`, `mark`, `sourcepos`, `short_subsuperscripts`, `wikilinks_title_after_pipe`, `wikilinks_title_before_pipe` |

### Format-Specific Extensions

Some extensions only apply to specific formats:

| Extension | Format | Effect |
|-----------|--------|--------|
| `empty_paragraphs` | ODT, ODF | Allow empty paragraphs |
| `native_numbering` | ODT, ODF, docx | Native figure/table numbering |
| `xrefs_name` | ODT, ODF | Cross-references by name |
| `xrefs_number` | ODT, ODF | Cross-references by number |
| `styles` | docx | Preserve custom styles |
| `citations` | typst, org, docx | Citation parsing |
| `element_citations` | JATS | `<element-citation>` elements |
| `ntb` | ConTeXt | Natural Tables |
| `tagging` | ConTeXt | Tagged PDF markup |
| `raw_markdown` | ipynb | Raw Markdown blocks |
| `amuse` | muse | Text::Amuse extensions |
| `fancy_lists` (org) | org | Alphabetical list markers |
| `smart_quotes` (org) | org | Curly quotes |
| `special_strings` (org) | org | Dashes/ellipses |

## Markdown Variants

Pandoc supports several Markdown variants, each with a different set of default extensions:

| Format | Description |
|--------|-------------|
| `markdown` | Pandoc Markdown (most extensions enabled) |
| `markdown_strict` | Original Markdown (Gruber spec) |
| `markdown_phpextra` | PHP Markdown Extra |
| `markdown_mmd` | MultiMarkdown |
| `commonmark` | CommonMark spec (few extensions) |
| `commonmark_x` | CommonMark with extra extensions |
| `gfm` | GitHub-Flavored Markdown |
| `markua` | Markua |

Use `--list-extensions=FORMAT` to see which extensions each variant enables by default.

## Format Dispatch

Internally, format dispatch is handled by `Text.Pandoc.Readers` (`getReader`) and `Text.Pandoc.Writers` (`getWriter`). These functions map format name strings (including extension modifiers) to the appropriate reader/writer function. For details on the internal dispatch, see [architecture.md](architecture.md).

## Related References

- [cli-options.md](cli-options.md) — `--from`, `--to`, `--list-*` flags
- [markdown.md](markdown.md) — Pandoc Markdown syntax reference
- [architecture.md](architecture.md) — Reader/writer modules and dispatch
