package main

func main() {
	app := App()

	app.Listen(":5000")
}
