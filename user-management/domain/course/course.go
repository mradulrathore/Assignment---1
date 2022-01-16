package course

import (
	"encoding/json"
	"log"

	"github.com/mradulrathore/user-management/domain/course/enum"
)

type Course struct {
	Enrol []enum.Course `json:"enrol"`
}

func New(courseEnrol []string) (Course, error) {
	var course Course
	var err error
	for _, c := range courseEnrol {
		var courseEnum enum.Course
		courseEnum, err = enum.CourseString(c)
		if err != nil {
			log.Println(err)
			return Course{}, err
		}
		course.Enrol = append(course.Enrol, courseEnum)
	}

	return course, nil
}

func (course Course) String() []string {
	var courses []string
	for _, c := range course.Enrol {
		courses = append(courses, c.String())
	}

	return courses
}

func (course *Course) EncodeCourse() ([]byte, error) {
	courseB, err := json.Marshal(course)
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}

	return courseB, nil
}

func DecodeCourse(courseB []byte) (Course, error) {
	var course Course
	if err := json.Unmarshal(courseB, &course); err != nil {
		log.Println(err)
		return course, err
	}

	return course, nil
}
