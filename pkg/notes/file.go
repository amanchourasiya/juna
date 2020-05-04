package notes

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/amanchourasiya/juna/pkg/utils/heap"
	"github.com/google/uuid"
)

var (
	notesFile = "./notes.txt"
)

func saveNotes(notes []heap.Comparable) {
	f, err := os.OpenFile(notesFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open file:%s to save notes for error:%v\n", notesFile, err)
	}
	for _, v := range notes {
		f.Write([]byte(v.ToString()))
	}
}

func readNotes() []*Note {
	ret := []*Note{}
	f, err := os.Open(notesFile)
	if err != nil {
		log.Fatalf("Failed to open notes file.Got error: %v\n", err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), "")
		id, err := uuid.Parse(words[0])
		if err != nil {
			log.Fatalf("Failed to convert string to uuid %v\n", err)
		}
		priority, _ := strconv.ParseFloat(words[1], 32)
		note := &Note{
			id:       id,
			priority: float32(priority),
			text:     words[2],
		}
		ret = append(ret, note)
	}
	return ret

}
