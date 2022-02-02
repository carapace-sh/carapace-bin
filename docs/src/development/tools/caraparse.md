# caraparse

[caraparse](https://github.com/rsteube/carapace-bin/tree/master/cmd/caraparse) is a parser for gnu help pages. It can handle variations to a certain degree but often it is best to prepare the text to fit the regex it uses:

```
-s              shorthand flag without argument
-s arg          shorthand flag with argument
--long          longhand flag without argument
--long arg      longhand flag with argument
-s, --long      short and longhand flag without argument
-s, --long arg  short and longhand flag with argument
```

Single space between flag and a word (`arg`) determines it accepts an argument and two or more for the description.

[![asciicast](https://asciinema.org/a/357895.svg)](https://asciinema.org/a/357895)
