package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"strings"
	"flag"
	"errors"
)

func main() {

	postFilePointer := flag.String("file", "first-post.txt", "a string")
	postDirPointer := flag.String("dir", ".", "a string")
	outputDirPointer := flag.String("output", "output", "a string")
	postDir := *postDirPointer
	outputDir := *outputDirPointer
	

	postSlice := []map[string]string{}

	if(postFilePointer != nil) {
		postSlice = append(postSlice, map[string]string{
			"path":*postFilePointer,
			"name":*postFilePointer,
		})
	}
	if(postFilePointer != nil) {
		filesArray, err := ioutil.ReadDir(postDir)
		if err != nil {
			panic(err)
		}
		for _, file := range filesArray {
			name := string(file.Name())
			if strings.Contains(name, ".txt") {
				postSlice = append(postSlice, map[string]string{
					"name":name,
					"path":postDir + "/" + name,
				})
			}
		}
	}

	// Read Template File
	templateDir := "template.tmpl"
	templateFile, err := os.Open(templateDir)
	if err != nil {
        log.Fatal(err)
    }
	defer templateFile.Close()
	template, err := ioutil.ReadAll(templateFile)
	templateString := string(template)

	for _, post := range postSlice {
		postString := readPostFile(post["path"])
		htmlString := strings.Replace(templateString, "{{ content }}", postString, -1)
		// fmt.Println(htmlString)
		save(htmlString, post["name"], outputDir)
	}
}

func readPostFile(postPath string) string {
	// Read Post Files
	postFile, err := os.Open(postPath)
	if err != nil {
		log.Fatal(err)
	}
	defer postFile.Close()
	post, err := ioutil.ReadAll(postFile)
	postString := string(post)
	return postString
}

func save(htmlString string, postDir string, outputDir string) {
	if _, err := os.Stat(outputDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(outputDir, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	htmlDir := outputDir + "/" + strings.Replace(postDir, "txt", "html", -1)
	fmt.Println(htmlDir)
	htmlFile, _ := os.Create(htmlDir)
	defer htmlFile.Close()
	htmlFile.WriteString(htmlString)
}
