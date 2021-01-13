package main


func main()  {

	//Run server
	mux := Routes()
	server := NewServer(mux)
	server.Run()

}
