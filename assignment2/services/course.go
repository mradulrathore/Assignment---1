package services

import (
	"errors"
	"strconv"
	"strings"

	models "mradulrathore/onboarding-assignments/assignment2/models"
)

// return true if duplicate course exist
func checkDuplicateCourse(courses []string) (exists bool, err error) {

	courseFrequency := make(map[string]int)

	for _, course := range courses {
		// check if the item/element exist in the duplicate_frequency map

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

	if len(courses) < models.NumberOfCousesRequired {
		errorMessage := "please choose atleast " + strconv.Itoa(models.NumberOfCousesRequired) + " courses"
		return false, errors.New(errorMessage)
	}

	duplicateExist, err := checkDuplicateCourse(courses)
	if duplicateExist {
		return false, err
	}

	for _, course := range courses {
		if _, exists := models.CourseAvailable[course]; !exists {
			return false, errors.New("select course from available courses")
		}
	}

	return true, nil
}
