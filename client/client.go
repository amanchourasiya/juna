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
	SaveNote(client, 21, "Finish grpc notes module")
	//SaveNote(client, 2, "Finish AWS course")
	//SaveNote(client, 20, "Finish SDR module")

	//GetLatestNote(client)
	//GetLatestNote(client)
	//GetAllNotes(client)
	//ClearAllNotes(client)
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

func GetLatestNote(client message.UserServiceClient) {
	note, err := client.GetLatestNote(context.Background(), &message.GetLatestNoteRequest{})
	if err != nil {
		log.Printf("Failed to get latest note %v\n", err)
	}
	log.Printf("Latest Note received %d %s\n", note.GetNote().GetPriority(), note.Note.GetText())
}

func ClearAllNotes(client message.UserServiceClient) {
	_, err := client.ClearAllNotes(context.Background(), &message.ClearAllNotesRequest{})
	if err != nil {
		log.Printf("Failed to clear all notes %v\n", err)
	}
	log.Println("All notes successfully cleared")

}
