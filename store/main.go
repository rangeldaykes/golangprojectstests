package main

import (
	"log"
	"net/http"
	"store/api/routers"
	"store/infrastucture/data/infraredis/infraredigo"
)

func main() {
	//t1 := time.Now()
	//time.Sleep(1500 * time.Millisecond)

	//t2 := time.Now()
	//diff := t2.Sub(t1)
	//fmt.Println(diff)
	//fmt.Println(int64(diff / time.Millisecond))

	//fmt.Println(float64(diff) / float64(time.Millisecond))

	//out := time.Time{}.Add(diff)
	//fmt.Println(out.Format("15:04:05:"))

	//log.Fatal(http.ListenAndServe(":9003", routers.ChiRouterGetInstance().InitRouter()))

	http.ListenAndServe(":9003", routers.ChiRouterGetInstance().InitRouter())
}

func testRedis() {
	resp, err := infraredigo.PersistenceRedigo{}.Ping()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Redis ping response = %s", resp)

}
