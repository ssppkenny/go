package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		panic("wrong number of arguments")
	}
	ext := argsWithoutProg[0]

	files := checkExt(ext)
	for _, s := range files {
		fmt.Println(s)
		target := strings.TrimSuffix(s, filepath.Ext(s)) + ".mp3"
		cmd := exec.Command("ffmpeg", "-i", s, "-ab", "320k", target)
		cmdOutput := &bytes.Buffer{}
		cmd.Stdout = cmdOutput
		err := cmd.Run()
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
		fmt.Print(string(cmdOutput.Bytes()))

	}

}

func checkExt(ext string) []string {
	pathS, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var files []string
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})
	return files
}
