package main

type appData struct {
    nextLocationAreaUrl *string
    prevLocationAreaUrl *string
}

func main() {
    appd := appData{}
    startRepl(&appd)
}

