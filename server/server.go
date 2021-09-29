package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type course struct {
	ID       string `json:"id"`
	Workload int    `json:"workload"`
	Rating   int    `json:"rating"`
}

var courses = []course{
	{ID: "1", Workload: 10, Rating: 80},
	{ID: "2", Workload: 10, Rating: 90},
	{ID: "3", Workload: 20, Rating: 75},
}

func main() {
	router := gin.Default()
	router.GET("/courses", getCourses)
	router.GET("/courses/:id", getCourseByID)
	router.POST("/courses", postCourse)

	router.Run("localhost:8080")
}

func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func postCourse(c *gin.Context) {
	var newCourse course

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
	id := c.Params.ByName("id")
	var course course

	for _, a := range courses {
		if a.ID == id {
			course = a
			return
		}
	}
	c.Bind(&course)

	courses = append(courses, course)
	c.IndentedJSON(http.StatusCreated, course)

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/users/1
}

/* func deleteCourse(c *gin.Context) {
	id := c.Params.ByName("id")
	var course course

	for _, a := range courses {
		if a.ID == id {
			course = a
			return
		}
	}

	//delete the course here

	c.JSON(200, gin.H{"success": "Course #" + id + " deleted"})

	// curl -i -X DELETE http://localhost:8080/api/v1/users/1
} */

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
