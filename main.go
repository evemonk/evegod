package main

import "os"
import "fmt"
import "strconv"
import "net/http"
import "encoding/json"
import "io/ioutil"
import "database/sql"
import _ "github.com/lib/pq"

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

const eve_etags_table = "eve_etags"
const eve_characters_table = "eve_characters"

const characters_table = "characters"

func main() {
	fmt.Println("EveGO Daemon")

	os.Setenv("DATABASE_URL", "postgres://localhost/evemonk_development?sslmode=disable")
	connStr := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var db_eve_character_id int = 1337512245

	var eve_character_url string = fmt.Sprintf("https://esi.evetech.net/v4/characters/%s/?datasource=tranquility", strconv.Itoa(db_eve_character_id))

	fmt.Println(eve_character_url)

	rows, err := db.Query(`SELECT * FROM eve_etags WHERE url = $1`, eve_character_url)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var id int
	var url string
	var etag string
	var created_at string
	var updated_at string

	for rows.Next() {
		err := rows.Scan(&id, &url, &etag, &created_at, &updated_at)
		if err != nil {
			// log.Fatal(err)
			panic(err)
		}
		fmt.Println()
		fmt.Println(id)
		fmt.Println(url)
		fmt.Println(etag)
		fmt.Println()
	}

	// fmt.Println(rows)

	if db != nil {
		fmt.Println("hello")
	}

	resp, err := http.Get(eve_character_url)

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
