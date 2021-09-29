package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
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

	command, e := os.ReadFile("./allCommand.txt")

	if e != nil {
		log.Fatal(e)
	}

	out, err := exec.Command("PowerShell", string(command)).Output()
	fmt.Println(out)
	printResult(string(out))

	if err != nil {
		log.Fatal(err)
	}
}

func getCourseById(_id string) {
	fmt.Println("running get " + _id)
	command, e := os.ReadFile("./idCommand.txt")

	if e != nil {
		log.Fatal(e)
	}

	out, err := exec.Command("PowerShell", string(command)).Output()
	printResult(string(out))

	if err != nil {
		log.Fatal(err)
	}

}

func newCourse(_id string, _workload int64, _rating int64) {
	fmt.Println("running new " + _id)
	command, e := os.ReadFile("./newCommand.txt")

	if e != nil {
		log.Fatal(e)
	}

	out, err := exec.Command("PowerShell", string(command)).Output()
	printResult(string(out))

	if err != nil {
		log.Fatal(err)
	}
}

func printResult(in string) {
	index := strings.Index(in, "\"id\"")
	result := string(in[index : index+100])
	result = strings.TrimSpace(result)
	result = strings.ReplaceAll(result, " ", "")

	fmt.Println("\nInfo for course:")
	fmt.Println(result)
}

//get,course,1
//new,course,4,10,90
//get,all
//(delete,course,1)
//(update,course,3,10,90)
