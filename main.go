// main.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
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
	"[1] append a word to both values & values-zh-rTW strings.xml\n" +
	"Which mode:"

func handleInsertStringXml(workspace string) {
	_, err := os.Stat(path.Join(workspace, "res", "values", "strings.xml"))
	if err != nil {
		log.Fatal(err)
	}
}

func askAction() string {
	fmt.Println(hintAction)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return ""
}
func main() {
	flag.Parse()
	var action string
	if len(flag.Args()) > 1 {
		action = flag.Args()[1]
	} else {
		action = askAction()
	}

	switch action {
	case "1":
		handleInsertStringXml(workspace)
	default:
		fmt.Println("unknown mode", action)
	}

}
