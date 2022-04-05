package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"strings"
	"flag"
)

func main() {

	postDirPointer := flag.String("file", "first-post.txt", "a string")
	postDir := *postDirPointer

	// Read Post File
	postFile, err := os.Open(postDir)
	if err != nil {
        log.Fatal(err)
    }
	defer postFile.Close()
	post, err := ioutil.ReadAll(postFile)
	postString := string(post)

	// Read Template File
	templateDir := "template.tmpl"
	templateFile, err := os.Open(templateDir)
	if err != nil {
        log.Fatal(err)
    }
	defer templateFile.Close()
	template, err := ioutil.ReadAll(templateFile)
	templateString := string(template)


	htmlString := strings.Replace(templateString, "{{ content }}", postString, -1)
	fmt.Println(htmlString)
	save(htmlString, postDir)
}

func save(htmlString string, postDir string) {
	htmlDir := strings.Replace(postDir, "txt", "html", -1)
	htmlFile, _ := os.Create(htmlDir)
	defer htmlFile.Close()
	htmlFile.WriteString(htmlString)
}
