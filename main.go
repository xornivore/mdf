package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/atsushinee/go-markdown-generator/doc"
)

var (
	linkPrefix = kingpin.Flag("link-prefix", "Link prefix").String()
	columns    = kingpin.Flag("column", "List of columns").Strings()
	path       = kingpin.Arg("path", "Path with files to list").String()
)

func main() {

	kingpin.Parse()

	md := doc.NewMarkDown()
	files, err := ioutil.ReadDir(*path)

	if err != nil {
		fmt.Println(err)
		return

	}

	if len(*columns) == 0 {
		fmt.Fprintf(os.Stderr, "Expecting --columns to be provided")
		os.Exit(1)
	}

	var links []string
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		links = append(links, fmt.Sprintf("[%s](%s/%s)", f.Name(), *linkPrefix, f.Name()))
	}

	t := doc.NewTable(len(links), len(*columns))

	for i, c := range *columns {
		t.SetTitle(i, c)
	}

	for i := range links {
		t.SetContent(i, 0, links[i])
	}
	md.WriteTable(t)

	fmt.Print(md.String())
}
