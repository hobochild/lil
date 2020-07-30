package main

import (
	"flag"
	"log"
  "fmt"

	"github.com/FooSoft/goldsmith"
	"github.com/FooSoft/goldsmith-components/devserver"
	"github.com/FooSoft/goldsmith-components/filters/condition"
	"github.com/FooSoft/goldsmith-components/plugins/frontmatter"
	"github.com/FooSoft/goldsmith-components/plugins/markdown"
	"github.com/FooSoft/goldsmith-components/plugins/minify"
	"github.com/FooSoft/goldsmith-components/plugins/syntax"
  "github.com/hobochild/fig/layout"
  "github.com/FooSoft/goldsmith-components/plugins/livejs"
)

type builder struct {
	dist bool
}



func (b *builder) Build(srcDir, dstDir, cacheDir string) {
  l := layout.New()
  l.DefaultLayout("__basic__")

	errs := goldsmith.
		Begin(srcDir).                     // read files from srcDir
		Chain(frontmatter.New()).          // extract frontmatter and store it as metadata
		Chain(markdown.New()).             // convert *.md files to *.html files
		Chain(l).               // apply *.gohtml templates to *.html files
		Chain(syntax.New()).               // apply *.gohtml templates to *.html files
		FilterPush(condition.New(b.dist)). // push a dist-only conditional filter onto the stack
		Chain(minify.New()).               // minify *.html, *.css, *.js, etc. files
		FilterPop().                       // pop off the last filter pushed onto the stack
    Chain(livejs.New()).
		End(dstDir)                        // write files to dstDir

	for _, err := range errs {
		log.Print(err)
	}
}

func main() {
	port := flag.Int("port", 8080, "server port")
	dev := flag.Bool("dev", false, "run dev server")
	flag.Parse()
  dist := !*dev
  b := builder{dist}

  if *dev {
    fmt.Printf("Running dev server on port %d", *port)
	  devserver.DevServe(&b, *port, "content", "build", "cache")
  } else {
    b.Build("content", "build", "cache")
  }
}
