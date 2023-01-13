package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/increaseLevel", increaseLevel)
	http.HandleFunc("/resetLevel", resetLevel)
	http.HandleFunc("/level", getLevel)
	http.HandleFunc("/cargo", getCargo)
	http.HandleFunc("/updateCargo", updateCargo)
	http.HandleFunc("/resetCargo", resetCargo)
	http.HandleFunc("/turn", getTurn)
	http.HandleFunc("/increaseTurn", increaseTurn)
	http.HandleFunc("/resetTurn", resetTurn)
	http.HandleFunc("/targetHeading", getTarget)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

type Player struct {
	level  int
	cargo  int
	turn   int
	target int
}

var player = Player{level: 1, cargo: 10, turn: 1, target: 10}

func increaseLevel(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		player.level = player.level + 1
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func resetLevel(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		player.level = 1
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func getLevel(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		io.WriteString(w, strconv.Itoa(player.level))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func getCargo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		io.WriteString(w, strconv.Itoa(player.cargo))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func updateCargo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err.Error())
		}

		amount, err := strconv.Atoi(string(reqBody))
		if err != nil {
			fmt.Println(err.Error())
		}

		player.cargo = player.cargo + amount
		if player.cargo < 0 {
			player.cargo = 0
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func resetCargo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		player.cargo = 10
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func getTurn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		io.WriteString(w, strconv.Itoa(player.turn))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func increaseTurn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		player.turn = player.turn + 1
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func resetTurn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		player.turn = 1
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func getTarget(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		io.WriteString(w, strconv.Itoa(player.target))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func updateTarget(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// TODO: Put update target logic with rng
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func resetTarget(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		player.target = 10
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}
