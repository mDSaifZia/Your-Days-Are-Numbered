package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/increaseLevel", increaseLevel)
	http.HandleFunc("/resetLevel", resetLevel)
	http.HandleFunc("/level", getLevel)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

type Player struct {
	level int
}

var player = Player{level: 1}

func increaseLevel(w http.ResponseWriter, r *http.Request) {
	player.level = player.level + 1
}

func resetLevel(w http.ResponseWriter, r *http.Request) {
	player.level = 1
}

func getLevel(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, strconv.Itoa(player.level))
}
