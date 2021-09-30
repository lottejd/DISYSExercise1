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

	case "delete":
		deleteCourse(split[2])
		break

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
		break
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
	if err != nil {
		log.Fatalln(err)
	}
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

func deleteCourse(_id string) {
	fmt.Println("Performing Http Delete...")
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
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("Body: " + bodyString)
}

func putCourse(_id string, _workload int64, _rating int64) {
	fmt.Println("Performing Http Put...")
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

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("Body: " + bodyString)

	// Convert response body to Todo struct
	var newcourse Course
	json.Unmarshal(bodyBytes, &newcourse)
	fmt.Printf("API Response as struct:\n%+v\n", newcourse)
}

//get,course,1
//new,course,4,10,90
//get,all,courses
//delete,course,1
//update,course,2,10,90
