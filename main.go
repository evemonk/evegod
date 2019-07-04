package main

import "fmt"
import "net/http"
import "encoding/json"
import "io/ioutil"

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

	resp, err := http.Get("https://esi.evetech.net/v4/characters/1337512245/?datasource=tranquility")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)

		body, err := ioutil.ReadAll(resp.Body)

		resp.Body.Close()

		if err != nil {
			panic(err)
		}

		res := Character{}

		fmt.Println(json.Unmarshal([]byte(body), &res))

		fmt.Printf("%s", body)
	}
}
