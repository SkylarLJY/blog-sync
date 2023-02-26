package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func syncPost(src string, dest string, blogFormat bool) {
	srcContent, err := os.ReadFile(src + ".md")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opeing src file: %s\n", err)
		os.Exit(1)
	}

	// create dest folder if does not exit
	if _, err := os.Stat(dest); errors.Is(err, os.ErrNotExist) {
		if err = os.Mkdir(dest, os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "error creating dest folder: %s\n", err)
			os.Exit(1)
		}
	}

	destFile, err := os.OpenFile(filepath.Join(dest, "index.md"), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening dest file: %s\n", err)
		os.Exit(1)
	}
	defer destFile.Close()

	if blogFormat {
		title := filepath.Base(src)
		time := time.Now().Format(time.RFC3339)

		description := string(srcContent)
		if len(srcContent) > 250 {
			description = description[:250]
		}
		descriptionStr := strings.Replace(description, "\n\n", "", -1) + "..."

		// TODO: let user define what header to add
		metadata := fmt.Sprintf(`---
title: %q
date: %q
description: %q
---
`, title, time, descriptionStr)

		_, err = destFile.Write([]byte(metadata))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	if _, err = destFile.Write(srcContent); err != nil {
		fmt.Fprintf(os.Stderr, "error writing to dest file: %s", err)
		os.Exit(1)
	}

}

func main() {
	file := flag.String("file", "", "name of the blog post to reformat")
	destFileName := flag.String("destName", "", "name of the destination file folder")
	format := flag.Bool("blogFormat", true, "reformat for blog post or not")
	flag.Parse()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if *file == "" {
		fmt.Fprintln(os.Stderr, "please specify the file to reformate")
		os.Exit(0)
	}
	src := os.Getenv("BLOG_SYNC_SRC_DIR")
	dest := os.Getenv("BLOG_SYNC_DEST_DIR")

	if *destFileName == "" {
		*destFileName = *file
	}

	syncPost(filepath.Join(src, *file), filepath.Join(dest, *destFileName), *format)
}
