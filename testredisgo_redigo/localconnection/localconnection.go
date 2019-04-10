package main

import (
	"encoding/json"
	"fmt"
	"log"

	//"transaction"
	"testredisgo_redigo/localconnection/transaction"

	"github.com/gomodule/redigo/redis"
)

func main() {

	c, err := redis.Dial("tcp", "192.168.228.150:6379")

	if err != nil {
		fmt.Println(err)
	}

	defer c.Close()

	//ping(c)
	//set(c)
	//get(c)
	//setStruct(c)
	//getStruct(c)

	//hashSetFlatStruc(c)
	//hashSetFlatMap(c)
	//hashSetGetAll(c)

	//hashSetJSONStruc(c)
	//hashSetJSONGetAll(c)
	setTransaction(c)
}

func ping(c redis.Conn) error {
	s, err := redis.String(c.Do("PING"))
	if err != nil {
		return err
	}

	fmt.Printf("PING Response = %s\n", s)

	return nil
}

// set executes the redis SET command
func set(c redis.Conn) error {
	_, err := c.Do("SET", "Favorite Movie", "Repo Man")
	if err != nil {
		return err
	}
	_, err = c.Do("SET", "Release Year", 1984)
	if err != nil {
		return err
	}
	return nil
}

// get executes the redis GET command
func get(c redis.Conn) error {

	// Simple GET example with String helper
	key := "Favorite Movie"
	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		return (err)
	}
	fmt.Printf("%s = %s\n", key, s)

	// Simple GET example with Int helper
	key = "Release Year"
	i, err := redis.Int(c.Do("GET", key))
	if err != nil {
		return (err)
	}
	fmt.Printf("%s = %d\n", key, i)

	// Example where GET returns no results
	key = "Nonexistent Key"
	s, err = redis.String(c.Do("GET", key))
	if err == redis.ErrNil {
		fmt.Printf("%s does not exist\n", key)
	} else if err != nil {
		return err
	} else {
		fmt.Printf("%s = %s\n", key, s)
	}

	return nil
}

// User is a simple user struct for this example
type User struct {
	Username  string `json:"username"`
	MobileID  int    `json:"mobile_id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func setStruct(c redis.Conn) error {

	const objectPrefix string = "user:"

	usr := User{
		Username:  "otto",
		MobileID:  1234567890,
		Email:     "ottoM@repoman.com",
		FirstName: "Otto",
		LastName:  "Maddox",
	}

	// serialize User object to JSON
	json, err := json.Marshal(usr)
	if err != nil {
		return err
	}

	// SET object
	_, err = c.Do("SET", objectPrefix+usr.Username, json)
	if err != nil {
		return err
	}

	return nil
}

func getStruct(c redis.Conn) error {

	const objectPrefix string = "user:"

	username := "otto"
	s, err := redis.String(c.Do("GET", objectPrefix+username))
	if err == redis.ErrNil {
		fmt.Println("User does not exist")
	} else if err != nil {
		return err
	}

	usr := User{}
	err = json.Unmarshal([]byte(s), &usr)

	fmt.Printf("%+v\n", usr)

	return nil
}

type movie struct {
	Title  string `redis:"title"`
	Author string `redis:"author"`
	Body   string `redis:"body"`
}

func hashSetFlatStruc(c redis.Conn) error {

	var p1 struct {
		Title  string `redis:"title"`
		Author string `redis:"author"`
		Body   string `redis:"body"`
	}

	p1.Title = "Example"
	p1.Author = "Gary"
	p1.Body = "Hello"

	if _, err := c.Do("HMSET", redis.Args{}.Add("id1").AddFlat(&p1)...); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func hashSetFlatMap(c redis.Conn) error {

	m := map[string]string{
		"title":  "Example2",
		"author": "Steve",
		"body":   "Map",
	}

	if _, err := c.Do("HMSET", redis.Args{}.Add("id2").AddFlat(m)...); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func hashSetFlatGetAll(c redis.Conn) error {

	var p2 struct {
		Title  string `redis:"title"`
		Author string `redis:"author"`
		Body   string `redis:"body"`
	}

	for _, id := range []string{"id1", "id2"} {

		v, err := redis.Values(c.Do("HGETALL", id))
		if err != nil {
			fmt.Println(err)
			return err
		}

		if err := redis.ScanStruct(v, &p2); err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Printf("%+v\n", p2)
	}

	return nil
}

func hashSetJSONStruc(c redis.Conn) error {

	usr := movie{
		Title:  "Example",
		Author: "Gary",
		Body:   "Hello",
	}

	// serialize User object to JSON
	json, err := json.Marshal(&usr)
	if err != nil {
		return err
	}

	m := map[string]string{usr.Title: string(json)}

	if _, err := c.Do("HMSET", redis.Args{}.Add("hashSetJSONStruc").AddFlat(m)...); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func hashSetJSONGetAll(c redis.Conn) error {

	mov := movie{}

	for _, id := range []string{"hashSetJSONStruc"} {

		v, err := c.Do("HGETALL", id)
		if err != nil {
			fmt.Println(err)
			return err
		}

		m, err := redis.StringMap(v, err)
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(m["Example"]), &mov)

		fmt.Printf("%+v\n", mov)
	}

	return nil
}

func setTransaction(c redis.Conn) {

	t := transaction.NewTransaction(c)

	todas := func(c redis.Conn) {
		c.Send("INCR", "Counter")
		c.Send("DECR", "Counter")
		c.Send("INCR", "Counter")
		c.Send("DECR", "Counter")
		c.Send("INCR", "Counter")
		c.Send("DECR", "Counter")
		c.Send("INCR", "Counter")
		c.Send("DECR", "Counter")
	}

	t.Do(todas)

	suces := func(reply interface{}) {
		log.Println("Success!")
		log.Println(reply)
	}

	t.OnSuccess(suces)

	fail := func(err error) {
		log.Println("Oh no, transaction failed, alert user.")
		log.Println(err)
	}

	t.OnFail(fail)

	// NewTransaction(c).Do(func(c redis.Conn) {
	// 	c.Send("INCR", "Counter")
	// 	c.Send("DECR", "Counter")
	// 	c.Send("INCR", "Counter")
	// 	c.Send("DECR", "Counter")
	// 	c.Send("INCR", "Counter")
	// 	c.Send("DECR", "Counter")
	// 	c.Send("INCR", "Counter")
	// 	c.Send("DECR", "Counter")
	// }).OnSuccess(func(reply interface{}) {
	// 	log.Println("Success!")
	// 	log.Println(reply)
	// }).OnFail(func(err error) {
	// 	log.Println("Oh no, transaction failed, alert user.")
	// 	log.Println(err)
	// })

}
