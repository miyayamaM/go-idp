package main

func main() {
	e, err := InitializeRouter()
	if err != nil {
		panic(err)
	}
	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
