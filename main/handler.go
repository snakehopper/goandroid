// handler
package main

import (
	"bufio"
	"fmt"
	"github.com/snakehopper/goandroid"
	"log"
	"os"
	"path"
)

func handleInsertStringXml(workspace string) {
	var xmlPath = path.Join(workspace, "res", "values", "strings.xml")
	var err error
	_, err = os.Stat(xmlPath)
	if err != nil {
		log.Fatal(err)
	}
	var xmlZhtwPath = path.Join(workspace, "res", "values-zh-rTW", "strings.xml")
	_, err = os.Stat(xmlZhtwPath)
	if err != nil {
		log.Fatal(err)
	}

	// ask user input
	key := askInput("\n<string name=")
	valDefault := askInput("default value:")
	valZhtw := askInput("正體中文:")

	if key == "" || valDefault == "" || valZhtw == "" {
		os.Exit(1)
	}

	//process default string
	res, err := goandroid.NewResourceXml(xmlPath)
	if err != nil {
		log.Fatal(err)
	}
	res = res.AddString(key, valDefault)
	err = res.ExportResourceXml(xmlPath)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\n%s\n\t<string name=%q>%s</string>\n", xmlPath, key, valDefault)
	}

	//process zh-rTW string
	res, err = goandroid.NewResourceXml(xmlZhtwPath)
	if err != nil {
		log.Fatal(err)
	}
	res = res.AddString(key, valZhtw)
	err = res.ExportResourceXml(xmlZhtwPath)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s\n\t<string name=%q>%s</string>\n\n", xmlZhtwPath, key, valZhtw)
	}
}

func askInput(hint string) string {
	fmt.Print(hint)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return ""
}
