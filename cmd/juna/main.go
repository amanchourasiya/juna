package main

import (
	"flag"
	"log"

	"github.com/amanchourasiya/juna/client"
	"github.com/amanchourasiya/juna/server"
)

func main() {
	mode := flag.String("mode", "client", "Mode for running application (server/client)")
	flag.Parse()

	if *mode == "server" {
		log.Println("Starting application in server mode")
		server.StartServer()
	} else if *mode == "client" {
		log.Println("Starting application in client mode")
		client.RunClient()
	} else {
		log.Fatal("Wrong mode selected . Please select from (server/client)")
	}
}

/*func test_heap(){
	node := notes.Note{
		Id: uuid.New(),


	}
}
*/
