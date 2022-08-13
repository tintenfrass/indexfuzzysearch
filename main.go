package main

import (
	"fmt"
	"indexfuzzysearch/config"
	"indexfuzzysearch/search"
	"indexfuzzysearch/ui"
)

func main() {
	fmt.Println("Index Fuzzy Search v0.9")

	search.ImportFromGit()
	config.Load()
	ui.BuildAndRun()
}
