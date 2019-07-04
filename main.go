package main

import "fmt"
import "net/http"

type Character struct {
	alliance_id     int64
	ancestry_id     int64
	birthday        string
	bloodline_id    int64
	corporation_id  int64
	description     string
	gender          string
	name            string
	race_id         int64
	security_status float32
}

func main() {
	fmt.Println("Hello go")

	resp, err := http.Get("http://example.com/")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
