package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type Person struct {
	Name  string
	Phone string
}

func main() {

	nuoche("122", 344)
	//session, err := mgo.Dial("mongodb://localhost")
	//if err != nil {
	//	panic(err)
	//}
	//defer session.Close()
	//
	//// Optional. Switch the session to a monotonic behavior.
	//session.SetMode(mgo.Monotonic, true)
	//c := session.DB("test").C("people")
	//err = c.Insert(&Person{"Ale", "+55 53 8116 9999"},
	//	&Person{"Cla", "+55 53 8402 8888"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//result := Person{}
	//err = c.Find(bson.M{"name": "Ale"}).One(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Phone:", result.Phone)
}

func nuoche(userId string, juli int16) {

	if juli > 3000 {

		log.Fatal("距离大于3000米")
		return
	}

	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("nuoche").C("cishu")

	result := Nuochecisghu{}
	err = c.Find(bson.M{"userId": userId, "Time": time.Now().Format("2006-01-02")}).One(&result)

	if err != nil {
		fmt.Println("次数:", result.Cheshu)
		if result.Cheshu == 0 {
			err = c.Insert(&Nuochecisghu{userId, result.Cheshu + 1, time.Now().Format("2006-01-02")})
			if err != nil {
				log.Fatal(err)
			}
		} else if result.Cheshu <= 3 {
			err = c.Update(bson.M{"userId": userId, "Time": time.Now().Format("2006-01-02")}, bson.M{"Cheshu": result.Cheshu + 1})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

type Nuochecisghu struct {
	UserId string
	Cheshu int
	Time   string
}
