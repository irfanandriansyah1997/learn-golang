package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) string() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) addEntry(text string) {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
}

func (j *Journal) removeEntry(index int) {
	j.entries = append(j.entries[:index], j.entries[index+1:]...)
	entryCount = len(j.entries)
}

// breaks SRP (single responsibility)
// Since journaling should only support adding and removing journal entries
// we can develop a separate function for saving to a file.

// func (j *Journal) Save(filename string) {
// 	_ = ioutil.WriteFile(filename,
// 		[]byte(j.String()), 0644)
// }

func checkFileIsExists(filename string) bool {
	_, err := os.Stat(filename)

	return !errors.Is(err, os.ErrNotExist)
}

func createFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()
}

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) saveToFile(j *Journal, filename string) {
	if isFileExists := checkFileIsExists(filename); !isFileExists {
		createFile(filename)
	}

	err := os.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	j := Journal{}
	j.addEntry("I Cried Today")
	j.addEntry("I Ate Bugs")
	j.addEntry("Sample Crud")
	fmt.Println(j.string())

	p := Persistence{lineSeparator: "\n"}
	p.saveToFile(&j, "journal.txt")

}
