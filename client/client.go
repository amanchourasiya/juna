package client

import (
	"context"
	"log"

	"github.com/amanchourasiya/juna/pkg/message"
	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

func RunClient() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial %v\n", err)
	}
	client := message.NewUserServiceClient(conn)
	//SaveNote(client, 23, "Finish grpc notes module")
	GetAllNotes(client)
}

func SaveNote(client message.UserServiceClient, priority int32, text string) {
	note := message.Note{Text: text, Priority: priority}
	req := message.SaveNoteRequest{Note: &note}
	resp, err := client.SaveNote(context.Background(), &req)
	if err != nil {
		log.Printf("Error saving note, Got err: %v\n", err)
	}
	log.Printf("Succesfully stored note , got response %v\n", resp.GetIsOk())

}

func GetAllNotes(client message.UserServiceClient) {
	notes, err := client.GetAllNotes(context.Background(), &message.GetAllNotesRequest{})
	if err != nil {
		log.Printf("Error getting all notes err: %v\n", err)
	}
	for _, note := range notes.Notes {
		log.Printf("%d %s\n", note.GetPriority(), note.GetText())
	}
}

func GetLatestNote() {

}
