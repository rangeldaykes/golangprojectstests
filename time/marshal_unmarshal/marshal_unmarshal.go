package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type JsonIso8601 time.Time

const JsonIso8601Layout = "2006-01-02T15:04:05"

func (j *JsonIso8601) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	var t time.Time
	if s != "null" {
		t, err = time.Parse(JsonIso8601Layout, s)
		if err != nil {
			return
		}
	}

	*j = JsonIso8601(t)
	return nil
}

func (j JsonIso8601) MarshalJSON() ([]byte, error) {
	if time.Time(j).IsZero() {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", j.Format(JsonIso8601Layout))), nil
}

func (j JsonIso8601) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

func (j JsonIso8601) String() string {
	t := time.Time(j)
	return t.Format(JsonIso8601Layout)
}

type User struct {
	ID        int
	Descricao string
	Birthday  JsonIso8601
	//Birthday2 JsonIso8601
}

//var jsonUser string = `{ "ID": 456, "Descricao": "Zé lelélão", "Birthday": "2019-04-25T05:55:00", "Birthday2": null}`
var jsonUser = `{ "ID": 456, "Descricao": "Zé lelélão", "Birthday": "2019-04-25T05:55:23"}`

//var jsonUser = `{ "ID": 456, "Descricao": "Zé lelélão"}`

func printUser(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "User is:")
	//user := User{ID: 123, Descricao: "Zé lelélão", Birthday: JsonIso8601(time.Now())}
	//user := User{ID: 123, Descricao: "Zé lelélão"}

	user := User{}
	err := json.Unmarshal([]byte(jsonUser), &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	fmt.Println(user.Birthday)
	//user.Birthday = JsonIso8601(time.Now())
	//fmt.Println(user.Birthday)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func main() {
	http.HandleFunc("/", printUser)

	http.ListenAndServe(":3000", nil)
}
