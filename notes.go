package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var noteFile = filepath.Join("notes-data", "notes.json")

type note struct {
	Title string
	Body  string
}

func fetchNotes() []note {
	var notes []note
	file, err1 := os.Open(noteFile)
	if err1 != nil {
		//create a file
		if err := os.Mkdir("notes-data", 0777); err != nil {
			log.Fatal(err)
		}
		if _, err := os.Create(noteFile); err != nil {
			log.Fatal(err)
		}
		return notes
	}
	data, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		return notes
	}
	if err := json.Unmarshal(data, &notes); err != nil {
		return notes
	}
	return notes
}

func saveNotes(n []note) {
	bs, err := json.Marshal(n)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(noteFile, bs, 0666)
}

func addNote(t, b string) *note {
	notes := fetchNotes()
	newNote := note{
		Title: t,
		Body:  b,
	}
	// check for duplicate
	var dupNote []note
	for _, note := range notes {
		if note.Title == t {
			dupNote = append(dupNote, note)
		}
	}
	if len(dupNote) == 0 {
		notes = append(notes, newNote)
		saveNotes(notes)
	} else {
		log.Fatal("Note already exists")
	}
	return &newNote
}

func (n *note) log() {
	fmt.Println("---")
	msg := "Title: " + n.Title + "\nBody: " + n.Body + "\n"
	fmt.Fprintf(os.Stdout, "%s", msg)
}

func getNote(t string) *note {
	notes := fetchNotes()
	var rnote []note // a slice containing note to be read
	for _, note := range notes {
		if note.Title == t {
			rnote = append(rnote, note)
		}
	}
	if len(rnote) == 0 {
		log.Fatal("Note doesn't exist")
	}
	return &rnote[0]
}

func removeNote(b bool, t string) bool {
	notes := fetchNotes()
	var fnotes []note
	if b {
		fnotes = []note{}
		saveNotes(fnotes)
	} else {
		for _, note := range notes {
			if t != note.Title {
				fnotes = append(fnotes, note)
			}
		}
		saveNotes(fnotes)
	}
	return len(notes) != len(fnotes)
}

func getAll() []note {
	return fetchNotes()
}
