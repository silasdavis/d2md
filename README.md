# d2md

This wrapper uses goldmark and forks goldmark-d2 so that the scale and other render options can be passed through

Usage:

Create a markdown file containg codeblocks with langauge info string `d2` containing valid d2 syntax

Then run, to produce a rendered HTML with the d2 diagrams inline as SVG

```
d2md doc.md > doc.html
```
