package notes

import (
	"fmt"

	"github.com/amanchourasiya/juna/pkg/utils/heap"
	"github.com/google/uuid"
)

var (
	heapstore string
)

// Note will hold the notes
type Note struct {
	id       uuid.UUID
	text     string
	priority float32
}

// GetKey returns the priority for the given note that can be used for comparisons
func (note *Note) GetKey() float32 {
	return note.priority
}

// Compare function compares two notes w.r.t. their priority
func (note *Note) Compare(note2 heap.Comparable) int {
	if note.GetKey() < note2.GetKey() {
		return -1
	} else if note.GetKey() > note2.GetKey() {
		return 1
	} else {
		return 0
	}
}

func (note *Note) ToString() string {
	ret := fmt.Sprintf("%v %f %s\n", note.id, note.priority, note.text)
	return ret
}

func (note *Note) GetText() string {
	return note.text
}

func CreateNote(priority float32, text string) {
	note := &Note{
		priority: priority,
		text:     text,
		id:       uuid.New(),
	}

	heap.Insert(note)
	saveNotes(heap.GetAllElements())
	//heap.TraverseHeap()

}

func GetAllNotes() []*Note {
	return readNotes()
}
