package server

import (
	"context"

	"github.com/amanchourasiya/juna/pkg/message"
	"github.com/amanchourasiya/juna/pkg/notes"
)

type UserService struct {
}

func (user *UserService) SaveNote(ctx context.Context, req *message.SaveNoteRequest) (*message.SaveNoteResponse, error) {
	//log.Println("calling SaveNote function")
	notes.CreateNote(float32(req.Note.GetPriority()), req.Note.GetText())

	return &message.SaveNoteResponse{IsOk: true}, nil
}

func (user *UserService) GetAllNotes(ctx context.Context, req *message.GetAllNotesRequest) (*message.GetAllNotesResponse, error) {
	msgnotes := []*message.Note{}
	allnotes := notes.GetAllNotes()
	for _, note := range allnotes {
		msgnote := &message.Note{
			Priority: int32(note.GetKey()),
			Text:     note.GetText(),
		}
		msgnotes = append(msgnotes, msgnote)
	}
	resp := message.GetAllNotesResponse{Notes: msgnotes}
	return &resp, nil
}

func (user *UserService) GetLatestNote(ctx context.Context, req *message.GetLatestNoteRequest) (*message.GetLatestNoteResponse, error) {
	return nil, nil
}
