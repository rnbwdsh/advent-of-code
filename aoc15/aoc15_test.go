package aoc15

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"
)

func readDay(day int) string {
	// make http request with cookie from ~/.config/aocd/token
	client := &http.Client{}
	url := "https://adventofcode.com/2015/day/" + strconv.Itoa(day) + "/input"
	req, _ := http.NewRequest("GET", url, nil)
	return readBodyWithCookie(req, client)
}

func readBodyWithCookie(req *http.Request, client *http.Client) string {
	// open token file and read the token
	token, err := os.ReadFile("/home/m/.config/aocd/token")
	if err != nil {
		panic(err)
	}
	tokenStr := string(token[:len(token)-1])

	req.AddCookie(&http.Cookie{Name: "session", Value: tokenStr})
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)[:len(body)-1]
}

var fun = []func(string) int{
	level1a, level1b, level2a, level2b, level3a, level3b, level4a, level4b, level5a, level5b,
	level6a, level6b, level7a, level7b, level8a, level8b, level9a, level9b, level10a, level10b,
	level11a, level11b, level12a, level12b, level13a, level13b, level14a, level14b, level15a, level15b,
	level16a, level16b, level17a, level17b, level18a, level18b, level19a, level19b, level20a, level20b,
	level21a, level21b, level22a, level22b, level23a, level23b, level24a, level24b, level25a, level25b,
}

func TestFirst(t *testing.T) {
	for day := 1; day <= len(fun)/2; day++ {
		data := readDay(day)
		idx := (day - 1) * 2

		// submit a
		var resp = fun[idx](data)
		println(day, "a", resp)

		// submit b
		resp = fun[idx+1](data)
		println(day, "b", resp)
	}
}
