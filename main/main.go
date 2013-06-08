// main.go
package main

import (
	"flag"
	"fmt"
)

var workspace string

func init() {
	const (
		defaultWorkspace = "."
		usageWorkspace   = "android project root path, expected AndroidManifest.xml, res, src"
	)
	flag.StringVar(&workspace, "workspace", defaultWorkspace, usageWorkspace)
	flag.StringVar(&workspace, "w", defaultWorkspace, usageWorkspace)
}

const hintAction = "======================================\n" +
	"Available action mode:\n" +
	"\t[1] append a word to both values & values-zh-rTW strings.xml\n" +
	"\nWhich mode:"

func main() {
	flag.Parse()
	var action string
	if len(flag.Args()) > 1 {
		action = flag.Args()[1]
	} else {
		action = askInput(hintAction)
	}

	switch action {
	case "1":
		for retry := "y"; retry == "y"; retry = askInput("\nMore input? (y/n) [n]: ") {
			handleInsertStringXml(workspace)
			fmt.Println("Done.")
		}
	default:
		fmt.Println("unknown mode", action)
	}

}
