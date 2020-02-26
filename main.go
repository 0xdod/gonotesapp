package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	titleUsage   = "Title of note"
	defaultTitle = "Title"
	bodyUsage    = "Body of note"
	defaultBody  = "Body"
)

var title string
var body string
var del bool

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Unexpected inputs")
		os.Exit(1)
	}
	// creating subcommands
	add := flag.NewFlagSet("add", flag.ExitOnError)
	add.StringVar(&title, "title", defaultTitle, titleUsage)
	add.StringVar(&title, "t", defaultTitle, titleUsage)
	add.StringVar(&body, "body", defaultTitle, titleUsage)
	add.StringVar(&body, "b", defaultBody, bodyUsage)
	list := flag.NewFlagSet("list", flag.ExitOnError)
	remove := flag.NewFlagSet("remove", flag.ExitOnError)
	remove.StringVar(&title, "title", defaultTitle, titleUsage)
	remove.StringVar(&title, "t", defaultTitle, titleUsage)
	remove.BoolVar(&del, "all", false, "Deletes all notes")
	remove.BoolVar(&del, "a", false, "Deletes all notes")
	read := flag.NewFlagSet("read", flag.ExitOnError)
	read.StringVar(&title, "title", defaultTitle, titleUsage)
	read.StringVar(&title, "t", defaultTitle, titleUsage)

	switch os.Args[1] {
	case "add":
		add.Parse(os.Args[2:])
		n := addNote(title, body)
		fmt.Println("Adding note")
		n.log()
	case "list":
		list.Parse(os.Args[2:])
		allNotes := getAll()
		if len(allNotes) == 0 {
			log.Fatal("No note saved!!")
		}
		fmt.Println("Listing " + strconv.Itoa(len(allNotes)) + " note(s)")
		for _, n := range allNotes {
			n.log()
		}

	case "remove":
		fmt.Println("Removing note...")
		remove.Parse(os.Args[2:])
		rmvd := removeNote(del, title)
		if rmvd {
			if del {
				fmt.Println("All notes removed")
			} else {
				fmt.Println(title + "Removed\n")
			}
		} else {
			fmt.Fprintf(os.Stderr, "%s", "Error: Note doesn't exist\n")
		}
	case "read":
		read.Parse(os.Args[2:])
		r := getNote(title)
		fmt.Println("Reading " + r.Title)
		r.log()
	default:
		fmt.Println("Wrong input")
		os.Exit(1)
	}
}
