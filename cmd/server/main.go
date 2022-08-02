package server

import "fmt"

// struct which contains pointer
// database connections
type App struct {
}

//setting up application
func (app *App) Run() error {
	fmt.Println("Setting up our app")
	return nil
}

func main() {
	fmt.Println("Go rest api")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our rest api")
		fmt.Println(err)
	}
}
