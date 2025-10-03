package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// middleware, helper - file

func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Lesson on Api")
	r := mux.NewRouter()
	// seeding
	courses = append(courses, Course{
		CourseId: "2", CourseName: "ReactJS", CoursePrice: 299,
		Author: &Author{Fullname: "Hites", Website: "learn.com"}},
	)
	courses = append(courses, Course{
		CourseId: "3", CourseName: "Golang", CoursePrice: 399,
		Author: &Author{Fullname: "Terence", Website: "learn.com"}},
	)
	//Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getSingleCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":4010", r))
}

// controllers -file
// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCode online</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get single course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

/*
Notes on Encode and Decode
Request: client -> server ; therefore we decode JSON to Go structs
Response: server -> client ; therefore we encode Go structs into JSON
So:
- Decoder = reads JSON and turns them to Go Values
- Encoder = takes Go values and turns them to JSON
-> Encoding is changing data from Go to JSON

So both json.NewDecoder and json.NewEncoder need to know where to read/write JSON
- json.NewDecoder(r.body) -> reads JSON from request body (client -> server)
- json.NewEncoder(w) -> writes JSON to the response writer (server -> client)
*/

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode(("No data inside the JSON"))
		return
	}

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return // This return statement is redudant since once the encoder response is returned, it automatically exits
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// Todo: send response when is not found
}

/* Notes on ... in the line courses = append(courses[:index], courses[index+1:]...)

First this line slices a bigger sloce into two parts, everthing before index and everthing after index
- If the line was courses = append(courses[:index], courses[index +1:]) -> without ...
	-we'd get a nested slice; slice B courses[index+1 : ] inside of slice A courses[:index]

... is called the variadic argument unpacking operator

- it basically takes a slice and unpacks its elements as individual arguments
- In the code example:
	it takes all elements before index
	appends all the elements after index one by one (not as a nested list)

example
numbers := []int{1, 2, 3}
fmt.Println(sum(numbers...)) // works ✅-> the numbers are unpacked individually


*/

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	// loop, id , remove(index, index +1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// send message of deletion
			break
		}
	}

}
