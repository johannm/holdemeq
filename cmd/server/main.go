package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
	"github.com/johannm/holdemeq/pkg/eval"
)

type payload struct {
    Hand1 string
    Hand2 string
    Board string
}

type result struct {
    Hand1 string
    Hand2 string
    Board string
    N int
    Win int
    Lose int
    Draw int
    Equity float64
}

func handleReq(w http.ResponseWriter, req *http.Request) {
    if req.Method != "POST" {
        http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    if req.URL.Path != "/equity" {
        http.Error(w, "404 not found", http.StatusNotFound)
        return
    }

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        http.Error(w, "400 bad request", http.StatusBadRequest)
        return
    }
    var d payload
    err = json.Unmarshal(body, &d)
    if err != nil {
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
    }

    hand1 := eval.ParseStr(d.Hand1)
	hand2 := eval.ParseStr(d.Hand2)
	var board []eval.Card
	if len(d.Board) > 0 {
		board = eval.ParseStr(d.Board)
	}
	n := 100000
	win, lose, draw := eval.CalculateHoldemEquity(hand1, hand2, board, n)

	resp := &result{
		Hand1: toStr(hand1), 
		Hand2: toStr(hand2),
		Board: toStr(board),
		N: n,
		Win: win,
		Lose: lose,
		Draw: draw, 
		Equity: (float64(win)+0.5*float64(draw))/float64(n),
	}
	respStr, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
	_, _ = fmt.Fprintln(w, string(respStr))
}

func main() {
    http.HandleFunc("/", handleReq)

    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func toStr(hand []eval.Card) string {
	var s string
	for _, c := range hand {
		s += c.ToStr()
	}
	return s
}