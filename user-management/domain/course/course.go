package course

import (
	"encoding/json"
	"log"
	"mradulrathore/onboarding-assignments/user-management/domain/course/enum"
)

type Course struct {
	Enrol []enum.Course `json:"enrol"`
}

func New(courseEnrol []string) (course Course, err error) {
	for _, c := range courseEnrol {
		var courseEnum enum.Course
		courseEnum, err = enum.CourseString(c)
		if err != nil {
			log.Println(err)
			return
		}
		course.Enrol = append(course.Enrol, courseEnum)
	}

	return
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

func (course Course) String() []string {

	var courses []string
	for _, c := range course.Enrol {
		courses = append(courses, c.String())
	}
	return courses
}
