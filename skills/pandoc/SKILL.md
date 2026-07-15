---
name: pandoc
description: >
  Use when working with pandoc — the universal document converter. Covers CLI
  options, input/output formats, extensions, defaults files, templates, filters
  (JSON, Lua, citeproc), citations/CSL, PDF generation, slide shows, EPUBs,
  chunked HTML, Jupyter notebooks, pandoc Markdown syntax, and the Haskell
  AST/reader-writer architecture. Triggers on: "pandoc", "pandoc Markdown",
  "pandoc template", "pandoc filter", "pandoc lua filter", "pandoc defaults",
  "pandoc citeproc", "pandoc PDF", "pandoc epub", "pandoc slide", "pandoc AST",
  "PandocMonad", "pandoc-types", "Text.Pandoc", "--from", "--to", "--standalone",
  "--citeproc", "--lua-filter", "--defaults", "--template", "--pdf-engine".
user-invocable: true
---

# Pandoc In-Depth Reference

Comprehensive reference for [pandoc](https://pandoc.org/), the universal document converter. Covers CLI usage, format support, extensions, defaults files, templates, filters, citations, PDF generation, specialized output formats, pandoc Markdown, and the internal Haskell architecture.

## Data Flow

```
input file(s)
  → reader (format-specific parser)
    → Pandoc AST (Meta + [Block])
      → filters (JSON / Lua / citeproc, applied sequentially)
        → internal transforms (headerShift, line breaks, etc.)
          → writer (format-specific renderer)
            → template wrapping (standalone mode)
              → output file
```

For PDF output, the writer produces an intermediate format (LaTeX, ConTeXt, HTML, or Typst) and an external engine renders the PDF.

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| CLI options, --from, --to, --output, --standalone, --data-dir, --sandbox, --verbose, --quiet, --fail-if-warnings, --log, --list-input-formats, --list-output-formats, --list-extensions, --list-highlight-languages, --version, --help, exit codes, reader options, writer options, --shift-heading-level-by, --file-scope, --metadata, --metadata-file, --tab-stop, --track-changes, --extract-media, --indented-code-classes, --default-image-extension, --wrap, --columns, --toc, --table-of-contents, --toc-depth, --include-in-header, --include-before-body, --include-after-body, --resource-path, --eol, --dpi, --top-level-division, --number-sections, --reference-doc, --strip-comments | [references/cli-options.md](references/cli-options.md) |
| input formats, output formats, readers, writers, extensions, --list-extensions, smart, auto_identifiers, raw_html, raw_tex, tex_math_dollars, fenced_code_attributes, pipe_tables, grid_tables, task_lists, citations, literate_haskell, markdown variants, commonmark, gfm, format dispatch, getReader, getWriter | [references/formats.md](references/formats.md) |
| defaults file, -d, --defaults, YAML defaults, JSON defaults, environment interpolation, ${HOME}, ${USERDATA}, ${.}, defaults directory, option mapping, filter shorthand, metadata-file shorthand | [references/defaults.md](references/defaults.md) |
| template, --template, default template, -D, --print-default-template, template syntax, $variable$, $if$, $for$, $endif$, $endfor$, $elseif$, partials, pipes, template variables, metadata variables, LaTeX variables, HTML variables, Beamer variables, ConTeXt variables, Typst variables, context construction, metaToContext | [references/templates.md](references/templates.md) |
| filter, --filter, JSON filter, --lua-filter, Lua filter, --citeproc, filter pipeline, applyFilters, pandoc-lua-engine, pandoc-lua, Lua interpreter, PANDOC_VERSION, PANDOC_STATE, Walkable, walk, walkM, query, filter ordering, security | [references/filters.md](references/filters.md) |
| citation, @key, citeproc, CSL, --csl, --bibliography, --citation-abbreviations, --natbib, --biblatex, nocite, references div, bibliography placement, link-citations, BibLaTeX, BibTeX, CSL JSON, CSL YAML, RIS, Chicago author-date | [references/citations.md](references/citations.md) |
| PDF, --pdf-engine, pdflatex, lualatex, xelatex, tectonic, latexmk, wkhtmltopdf, weasyprint, context, groff, typst, prince, pagedjs-cli, --pdf-engine-opt, LaTeX packages, PDF/A, PDF/UA, PDF/X, pdfstandard, SOURCE_DATE_EPOCH, reproducible builds, accessible PDF, tagged PDF | [references/pdf.md](references/pdf.md) |
| slide show, revealjs, beamer, pptx, slidy, dzslides, slideous, s5, slide level, incremental, --slide-level, -i, pauses, speaker notes, columns, fragile frame, EPUB, epub:type, Dublin Core, cover image, epub fonts, chunked HTML, Jupyter notebook, ipynb, vimdoc, custom reader, custom writer, Lua reader, Lua writer | [references/specialized-output.md](references/specialized-output.md) |
| pandoc Markdown, Markdown syntax, headings, setext, ATX, block quote, verbatim, code block, fenced code, line block, bullet list, ordered list, definition list, task list, example list, horizontal rule, table, pipe table, grid table, simple table, multiline table, metadata block, YAML metadata, backslash escape, emphasis, strong, strikeout, superscript, subscript, verbatim inline, math, raw HTML, raw TeX, LaTeX macros, links, images, footnotes, inline notes, citation syntax, Div, Span, fenced div, bracketed span, alerts, emoji, wikilinks, non-default extensions | [references/markdown.md](references/markdown.md) |
| architecture, AST, Pandoc, Meta, Block, Inline, Attr, pandoc-types, Text.Pandoc.Definition, Text.Pandoc.Builder, Text.Pandoc.Walk, PandocMonad, PandocIO, PandocPure, runIO, runPure, reader-writer pipeline, convertWithOpts, Text.Pandoc.App, source code organization, module structure, pandoc-cli, pandoc-server, pandoc-types package, citeproc, texmath, doctemplates, skylighting | [references/architecture.md](references/architecture.md) |

## Quick Guide

- **How do I convert a file?** → [references/cli-options.md](references/cli-options.md)
- **What formats does pandoc support?** → [references/formats.md](references/formats.md)
- **How do I use a defaults file?** → [references/defaults.md](references/defaults.md)
- **How do I write or customize a template?** → [references/templates.md](references/templates.md)
- **How do I write a Lua or JSON filter?** → [references/filters.md](references/filters.md)
- **How do citations and bibliographies work?** → [references/citations.md](references/citations.md)
- **How do I generate a PDF?** → [references/pdf.md](references/pdf.md)
- **How do I create slides, EPUBs, or Jupyter notebooks?** → [references/specialized-output.md](references/specialized-output.md)
- **What is pandoc's Markdown syntax?** → [references/markdown.md](references/markdown.md)
- **How does pandoc work internally (AST, readers, writers)?** → [references/architecture.md](references/architecture.md)
- **How do I enable/disable Markdown extensions?** → [references/formats.md](references/formats.md) and [references/markdown.md](references/markdown.md)
- **What exit codes does pandoc use?** → [references/cli-options.md](references/cli-options.md)
- **How do I run pandoc as a Lua interpreter?** → [references/filters.md](references/filters.md)
- **How do I make reproducible or accessible PDFs?** → [references/pdf.md](references/pdf.md)

## Cross-Project References

- For shell completion integration with pandoc (completer specs, bridge actions), see the **carapace** skill.
- For carapace-bin's pandoc completer implementation, see `completers/common/pandoc_completer/`.
