package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Course struct {
	ID       string `json:"id"`
	Workload int64  `json:"workload"`
	Rating   int64  `json:"rating"`
}

var courses = []Course{
	{ID: "0", Workload: 10, Rating: 80},
	{ID: "1", Workload: 10, Rating: 90},
	{ID: "2", Workload: 20, Rating: 75},
}

func mainServer() {
	router := gin.Default()
	router.GET("/courses", getCourses)
	router.GET("/courses/:id", getCourseByID)
	router.POST("/courses", postCourse)
	router.DELETE("/courses/:id", delCourse)
	router.PUT("/courses/:id/:workload/:rating", updateCourse)

	router.Run("localhost:8080")
}

func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func postCourse(c *gin.Context) {
	var newCourse Course

	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

func getCourseByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range courses {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
}

func updateCourse(c *gin.Context) {
	idstring := c.Params.ByName("id")
	id, err := strconv.ParseInt(idstring, 0, 64)
	workload, err := strconv.ParseInt(c.Params.ByName("workload"), 0, 32)
	if err != nil {
		log.Fatalln(err)
	}
	rating, err := strconv.ParseInt(c.Params.ByName("rating"), 0, 32)
	if err != nil {
		log.Fatalln(err)
	}
	updatedcourse := Course{idstring, workload, rating}

	fmt.Println("Deleting course...")

	for _, a := range courses {
		ID, err := strconv.ParseInt(a.ID, 0, 64)
		if err == nil && ID == id {
			if int64(len(courses)) == id+1 {
				courses = courses[:id]
			} else if id == 0 {
				courses = courses[id+1:]
			} else {
				courses = append(courses[:id], courses[id+1:]...)
			}
		}
	}
	fmt.Println("Appending updated course...")
	courses = append(courses, updatedcourse)
	c.Bind(&updatedcourse)
	c.IndentedJSON(http.StatusCreated, updatedcourse)

	sort.Slice(courses, func(i, j int) bool {
		return courses[i].ID < courses[j].ID
	})
}

func delCourse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	if err != nil {
		log.Fatalln(err)
	}

	for _, a := range courses {
		ID, err := strconv.ParseInt(a.ID, 0, 64)
		if err == nil && ID == id {
			if int64(len(courses)) == id+1 {
				courses = courses[:id]
			} else if id == 0 {
				courses = courses[id+1:]
			} else {
				courses = append(courses[:id], courses[id+1:]...)
			}
			return
		}
	}
}

//Command for posting
/* curl http://localhost:8080/courses ^
   --include ^
   --header "Content-Type: application/json" ^
   --request "POST" ^
   --data "{\"id\": \"4\", \"Workload\": 10, \"Rating\": 60}" */

//Command for getting
/* curl http://localhost:8080/courses ^
--header "Content-Type: application/json" ^
--request "GET" */

//Command for get by id=2
// curl http://localhost:8080/courses/2
