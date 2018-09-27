package main

func main() {

	// Start server
	e, err := createApp(nil)

	if err != nil {
		panic(err)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
