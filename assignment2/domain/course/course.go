package course

import (
	"encoding/json"
	"log"
)

const (
	TotalCourses           = 6
	NumberOfCousesRequired = 4
)

var CourseAvailable = map[string]bool{"A": true, "B": true, "C": true, "D": true, "E": true, "F": true}

type Course struct {
	Name string `json:"name"`
}

func (course *Course) EncodeCourse() (courseB []byte, err error) {

	courseB, err = json.Marshal(course)

	if err != nil {
		log.Println(err)
		return courseB, err
	}

	return courseB, nil
}

func DecodeCourse(courseB []byte) (course Course, err error) {

	if err := json.Unmarshal(courseB, &course); err != nil {
		log.Println(err)
		return course, err
	}

	return course, nil

}
