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

//Hard-coded instances
var courses = []Course{
	{ID: "0", Workload: 10, Rating: 80},
	{ID: "1", Workload: 10, Rating: 90},
	{ID: "2", Workload: 20, Rating: 75},
}

//The server runs using gin's router. Here possible commands are defined
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

	sort.Slice(courses, func(i, j int) bool {
		return courses[i].ID < courses[j].ID
	})

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
	id := c.Params.ByName("id")
	workload, err := strconv.ParseInt(c.Params.ByName("workload"), 0, 32)
	if err != nil {
		log.Fatalln(err)
	}
	rating, err := strconv.ParseInt(c.Params.ByName("rating"), 0, 32)
	if err != nil {
		log.Fatalln(err)
	}

	newC := Course{id, workload, rating}
	oldCourses := courses

	courses = courses[0:0]

	for _, a := range oldCourses {
		if err == nil && a.ID == id {
			fmt.Println(len(courses))
			for _, c := range oldCourses {
				if c.ID != a.ID {
					courses = append(courses, c)
				}
			}

		}

	}

	courses = append(courses, newC)

	sort.Slice(courses, func(i, j int) bool {
		return courses[i].ID < courses[j].ID
	})

	c.IndentedJSON(http.StatusOK, newC)
}

func delCourse(c *gin.Context) {
	id := c.Params.ByName("id")

	oldCourses := courses

	courses = courses[0:0]

	for _, a := range oldCourses {
		ID := a.ID
		if ID == id {
			fmt.Println(len(courses))
			for _, c := range oldCourses {
				if c.ID != a.ID {
					courses = append(courses, c)
				}
			}
			c.IndentedJSON(http.StatusOK, a)
		}
	}

	sort.Slice(courses, func(i, j int) bool {
		return courses[i].ID < courses[j].ID
	})

}
