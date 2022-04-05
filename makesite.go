package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"strings"
)

func main() {
	// Read Post File
	postDir := "first-post.txt"
	postFile, err := os.Open(postDir)
	if err != nil {
        log.Fatal(err)
    }
	defer postFile.Close()
	post, err := ioutil.ReadAll(postFile)
	post_string := string(post)

	// Read Template File
	templateDir := "template.tmpl"
	templateFile, err := os.Open(templateDir)
	if err != nil {
        log.Fatal(err)
    }
	defer templateFile.Close()
	template, err := ioutil.ReadAll(templateFile)
	template_string := string(template)


	html_string := strings.Replace(template_string, "{{ content }}", post_string, -1)
	fmt.Println(html_string)

	htmlDir := strings.Replace(postDir, "txt", "html", -1)
	htmlFile, err := os.Create(htmlDir)
	defer htmlFile.Close()
	htmlFile.WriteString(html_string)
}
