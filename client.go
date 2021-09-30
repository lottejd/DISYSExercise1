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

//Listen to the terminal, and when input comes in, read it
func mainClient() {
	var temp string

	for {
		fmt.Scanln(&temp)
		readInput(temp)
	}
}

//Read input to assert command and parameters of the input
func readInput(in string) {
	var split []string = strings.Split(in, ",")

	switch split[0] {
	case "get":
		if split[1] == "all" {
			getAllCourses()
		} else {
			getCourseById(split[2])
		}

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

	case "delete":
		deleteCourse(split[2])

	case "update":
		var id string = split[2]
		workload, err := strconv.ParseInt(split[3], 0, 32)
		if err != nil {
			log.Fatalln(err)
		}
		rating, err := strconv.ParseInt(split[4], 0, 64)

		if err != nil {
			log.Fatalln(err)
		}

		putCourse(id, workload, rating)

	default:
		fmt.Println("Could not understand input. See the readme.md file for info on commands")
	}

}

func getAllCourses() {
	fmt.Println("Running get all")

	resp, err := http.Get("http://localhost:8080/courses")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to Course struct
	var c []Course
	json.Unmarshal(bodyBytes, &c)
	fmt.Printf("Response: %+v\n", c)

}

func getCourseById(_id string) {
	fmt.Println("Running get " + _id)

	resp, err := http.Get("http://localhost:8080/courses/" + _id)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to Course struct
	var c Course
	json.Unmarshal(bodyBytes, &c)
	fmt.Printf("Response: %+v\n", c)

}

func newCourse(_id string, _workload int64, _rating int64) {
	fmt.Println("Running new " + _id)

	newc := Course{_id, _workload, _rating}
	jsonReq, err := json.Marshal(newc)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post("http://localhost:8080/courses", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to Course struct
	var newcourse Course
	json.Unmarshal(bodyBytes, &newcourse)
	fmt.Printf("Response: %+v\n", newcourse)
}

func deleteCourse(_id string) {
	fmt.Println("Running delete " + _id)
	newc := Course{"4", 10, 80}
	jsonReq, err := json.Marshal(newc)
	if err != nil {
		log.Fatalln(err)
	}
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/courses/"+_id, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	fmt.Println("Deleted course " + _id)
}

func putCourse(_id string, _workload int64, _rating int64) {
	fmt.Println("Running put " + _id)
	newc := Course{_id, _workload, _rating}
	jsonReq, err := json.Marshal(newc)
	if err != nil {
		log.Fatalln(err)
	}
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/courses/"+_id+"/"+fmt.Sprint(_workload)+"/"+fmt.Sprint(_rating), bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to Course struct
	var newcourse Course
	json.Unmarshal(bodyBytes, &newcourse)
	fmt.Printf("Response:\n%+v\n", newcourse)
}
