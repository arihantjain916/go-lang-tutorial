package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	runServer()
}

// Model for DB
type Courses struct {
	CourseId string  `json:"course_id"`
	Name     string  `json:"course_name"`
	Price    int     `json:"course_price"`
	Platform string  `json:"course_platform"`
	Author   *Author `json:"author"`
}

type Author struct {
	Name string `json:"author_name"`
	Age  int    `json:"author_age"`
}

// fake DB
var courses []Courses

// middleware, helpers
func (c *Courses) IsEmpty() bool {
	return c.Name == ""
}

// controllers
func runServer() {
	router := mux.NewRouter()

	// GET request
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/courses/{id}", getSpecificCourse).Methods("GET")

	// POST request
	router.HandleFunc("/addCourses", createCourse).Methods("POST")

	// PUT request
	router.HandleFunc("/updateCourses/{id}", updateCourse).Methods("PUT")

	// DELETE request
	router.HandleFunc("/deleteCourses/{id}", deleteCourse).Methods("DELETE")

	fmt.Println("Server is running on port https://localhost:8080")

	err := http.ListenAndServe(":8080", router)

	printErr(err)
}

func serveHome(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Write([]byte("<h1>Welcome to my API</h1>"))
}

func getAllCourses(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getSpecificCourse(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range courses {
		if item.CourseId == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given ID")
}

func createCourse(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Courses

	err := json.NewDecoder(r.Body).Decode(&course)

	printErr(err)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Course Name is not empty")
		return
	}

	// generate unique ID for course
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

func updateCourse(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range courses {
		if item.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Courses
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given ID")
}

func deleteCourse(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range courses {
		if item.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			json.NewEncoder(w).Encode("Courses deleted Successfully...")
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given ID")
}

func printErr(err error) {
	if err != nil {
		panic(err)
	}
}
