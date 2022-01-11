package course

import (
	"errors"
	"strconv"
	"strings"

	cours "mradulrathore/onboarding-assignments/assignment2/domain/course"
)

// return true if duplicate course exist
func checkDuplicateCourse(courses []string) (exists bool, err error) {

	courseFrequency := make(map[string]int)

	for _, course := range courses {

		_, exist := courseFrequency[course]

		if exist {
			return true, errors.New("duplicate course")
		} else {
			courseFrequency[course] = 1
		}
	}
	return false, nil
}

func ValidateCourse(coursesEnrol string) (ok bool, err error) {
	courses := strings.Split(coursesEnrol, ",")

	if len(courses) < cours.NumberOfCousesRequired {
		errorMessage := "please choose atleast " + strconv.Itoa(cours.NumberOfCousesRequired) + " courses"
		return false, errors.New(errorMessage)
	}

	duplicateExist, err := checkDuplicateCourse(courses)
	if duplicateExist {
		return false, err
	}

	for _, course := range courses {
		if _, exists := cours.CourseAvailable[course]; !exists {
			return false, errors.New("select course from available courses")
		}
	}

	return true, nil
}
