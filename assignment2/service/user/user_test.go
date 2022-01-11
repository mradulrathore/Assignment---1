package user

// cmd go test -coverprofile=coverage.out

import (
	"reflect"
	"testing"

	usr "mradulrathore/onboarding-assignments/assignment2/domain/user"
)

func TestInsertUserDetails(t *testing.T) {

	usersDetails = append(usersDetails, usr.User{

		FullName: "Anshul",
		Age:      20,
		Address:  "Indore,M.P.",
		RollNo:   43,
	})

	usersDetails = append(usersDetails, usr.User{
		FullName: "Rahul",
		Age:      20,
		Address:  "Indore,M.P.",
		RollNo:   41,
	})

	usersDetailsInput := usr.User{
		FullName: "Mradul",
		Age:      20,
		Address:  "Indore,M.P.",
		RollNo:   42,
	}

	expectedUserDetails := []usr.User{
		{

			FullName: "Anshul",
			Age:      20,
			Address:  "Indore,M.P.",
			RollNo:   43,
		},
		{
			FullName: "Mradul",
			Age:      20,
			Address:  "Indore,M.P.",
			RollNo:   42,
		},
		{
			FullName: "Rahul",
			Age:      20,
			Address:  "Indore,M.P.",
			RollNo:   41,
		},
	}

	err := InsertUserDetails(usersDetailsInput)
	if err != nil {
		t.Errorf("exception is  occuring: %q", err)
	}
	if !(reflect.DeepEqual(expectedUserDetails, usersDetails)) {
		t.Errorf("Got: %v, Expected: %v", usersDetails, expectedUserDetails)
	}

}

func TestValidateUserDetails(t *testing.T) {

	usersDetailsInput := []usr.User{
		{
			FullName: "Mradul",
			Age:      20,
			Address:  "Indore,M.P.",
			RollNo:   43,
		},
		// negative age
		{
			FullName: "Rahul",
			Age:      -20,
			Address:  "Indore,M.P.",
			RollNo:   43,
		},
		//blank fullname
		{
			FullName: "",
			Age:      20,
			Address:  "Indore,M.P.",
			RollNo:   43,
		},
		//blank address
		{
			FullName: "Mradul",
			Age:      20,
			Address:  "",
			RollNo:   43,
		},
		//rollno not provided
		{
			FullName: "Mradul",
			Age:      20,
			Address:  "Indore",
		},
		//address not provided
		{
			FullName: "Mradul",
			Age:      20,
			RollNo:   43,
		},
	}

	expected := []bool{true, false, false, false, false, false}

	for index, user := range usersDetailsInput {
		if ok, err := ValidateUserDetails(user); ok != expected[index] {
			if !expected[index] {
				t.Errorf("exception is occuring: %q", err)
			} else {
				t.Errorf("exception is not occuring: %q", err)
			}

		}
	}
}
