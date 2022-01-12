package course

import (
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

func (course Course) String() []string {

	var courses []string
	for _, c := range course.Enrol {
		courses = append(courses, c.String())
	}
	return courses
}
