package main

type app struct {
	config *AppConfigBarang
}

func (a app) run() {
	NewBarangDelevery(a.config).create()
}

func newApp() app {
	config := NewConfig()
	return app{config}
}

func main() {
	newApp().run()
}
