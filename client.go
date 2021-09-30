package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func mainClient() {
	var temp string

	for {
		fmt.Scanln(&temp)
		fmt.Println(temp)
		readInput(temp)
	}
}

func readInput(in string) {
	var split []string = strings.Split(in, ",")

	switch split[0] {
	case "get":
		if split[1] == "all" {
			getAllCourses()
		} else {
			getCourseById(split[2])
		}
		break

	case "new":
		var id string = split[2]
		workload, err := strconv.ParseInt(split[3], 0, 32)
		rating, er := strconv.ParseInt(split[4], 0, 64)

		if err == nil && er == nil {
			newCourse(id, workload, rating)
		} else {
			log.Fatal(err)
			log.Fatal(er)
		}
		break

		/* case "delete":
			//deleteCourse(split[2])
			break

		case "update":
			var id string = split[2]
			workload, err := strconv.ParseInt(split[3], 2, 32)
			rating, err := strconv.ParseInt(split[4], 2, 32)

			//putCourse(id,workload,rating)
			break */
	}

}

func getAllCourses() {

	fmt.Println("running get all")

	resp, err := http.Get("http://localhost:8080/courses")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var c []Course
	json.Unmarshal(bodyBytes, &c)
	fmt.Printf("API Response as struct %+v\n", c)

}

func getCourseById(_id string) {
	fmt.Println("running get " + _id)

	resp, err := http.Get("http://localhost:8080/courses/" + _id)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var c Course
	json.Unmarshal(bodyBytes, &c)
	fmt.Printf("API Response as struct %+v\n", c)

}

func newCourse(_id string, _workload int64, _rating int64) {
	fmt.Println("running new " + _id)

	fmt.Println("2. Performing Http Post...")
	newc := Course{_id, _workload, _rating}
	jsonReq, err := json.Marshal(newc)
	resp, err := http.Post("http://localhost:8080/courses", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert response body to Todo struct
	var newcourse Course
	json.Unmarshal(bodyBytes, &newcourse)
	fmt.Printf("%+v\n", newcourse)
}

/* func printResult(in string) {
	index := strings.Index(in, "\"id\"")
	result := string(in[index : index+100])
	result = strings.TrimSpace(result)
	result = strings.ReplaceAll(result, " ", "")

	fmt.Println("\nInfo for course:")
	fmt.Println(result)
} */

//get,course,1
//new,course,4,10,90
//get,all
//(delete,course,1)
//(update,course,3,10,90)
