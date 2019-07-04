package main

import "fmt"
import "net/http"
import "encoding/json"
import "io/ioutil"

type EveCharacter struct {
	AllianceID     int64   `json:"alliance_id"`
	AncestryID     int64   `json:"ancestry_id"`
	Birthday       string  `json:"birthday"`
	BloodlineID    int64   `json:"bloodline_id"`
	CorporationID  int64   `json:"corporation_id"`
	Description    string  `json:"description"`
	Gender         string  `json:"gender"`
	Name           string  `json:"name"`
	RaceID         int64   `json:"race_id"`
	SecurityStatus float64 `json:"security_status"`
}

const character_url_pattern = "https://esi.evetech.net/v4/characters/%s/"

func main() {
	fmt.Println("EveGO Daemon")

	resp, err := http.Get("https://esi.evetech.net/v4/characters/1337512245/?datasource=tranquility")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	eve_character := new(EveCharacter)

	json.Unmarshal([]byte(body), eve_character)

	fmt.Printf("AllianceID: %d\n", eve_character.AllianceID)
	fmt.Printf("AncestryID: %d\n", eve_character.AncestryID)
	fmt.Printf("Birthday: %s\n", eve_character.Birthday)
	fmt.Printf("BloodlineID: %d\n", eve_character.BloodlineID)
	fmt.Printf("CorporationID: %d\n", eve_character.CorporationID)
	fmt.Printf("Description: \"%s\"\n", eve_character.Description)
	fmt.Printf("Gender: %s\n", eve_character.Gender)
	fmt.Printf("Name: %s\n", eve_character.Name)
	fmt.Printf("RaceID: %d\n", eve_character.RaceID)
	fmt.Printf("SecurityStatus: %f\n", eve_character.SecurityStatus)

	fmt.Printf("%s\n", body)

	fmt.Println()
}
