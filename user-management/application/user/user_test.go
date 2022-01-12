package user

// cmd go test -coverprofile=coverage.out

import (
	"reflect"
	"testing"

	usr "mradulrathore/onboarding-assignments/user-management/domain/user"
)

func TestInsertUserDetails(t *testing.T) {

	users = append(users, usr.User{

		Name:    "Anshul",
		Age:     20,
		Address: "Indore,M.P.",
		RollNo:  43,
	})

	users = append(users, usr.User{
		Name:    "Rahul",
		Age:     20,
		Address: "Indore,M.P.",
		RollNo:  41,
	})

	usersInput := usr.User{
		Name:    "Mradul",
		Age:     20,
		Address: "Indore,M.P.",
		RollNo:  42,
	}

	expectedUserDetails := []usr.User{
		{

			Name:    "Anshul",
			Age:     20,
			Address: "Indore,M.P.",
			RollNo:  43,
		},
		{
			Name:    "Mradul",
			Age:     20,
			Address: "Indore,M.P.",
			RollNo:  42,
		},
		{
			Name:    "Rahul",
			Age:     20,
			Address: "Indore,M.P.",
			RollNo:  41,
		},
	}

	Insert(usersInput)

	if !(reflect.DeepEqual(expectedUserDetails, expectedUserDetails)) {
		t.Errorf("Got: %v, Expected: %v", usersInput, expectedUserDetails)
	}

}

// func TestValidateUserDetails(t *testing.T) {

// 	usersDetailsInput := []usr.User{
// 		{
// 			FullName: "Mradul",
// 			Age:      20,
// 			Address:  "Indore,M.P.",
// 			RollNo:   43,
// 		},
// 		// negative age
// 		{
// 			FullName: "Rahul",
// 			Age:      -20,
// 			Address:  "Indore,M.P.",
// 			RollNo:   43,
// 		},
// 		//blank fullname
// 		{
// 			FullName: "",
// 			Age:      20,
// 			Address:  "Indore,M.P.",
// 			RollNo:   43,
// 		},
// 		//blank address
// 		{
// 			FullName: "Mradul",
// 			Age:      20,
// 			Address:  "",
// 			RollNo:   43,
// 		},
// 		//rollno not provided
// 		{
// 			FullName: "Mradul",
// 			Age:      20,
// 			Address:  "Indore",
// 		},
// 		//address not provided
// 		{
// 			FullName: "Mradul",
// 			Age:      20,
// 			RollNo:   43,
// 		},
// 	}

// 	expected := []bool{true, false, false, false, false, false}

// 	for index, user := range usersDetailsInput {
// 		if ok, err := ValidateUserDetails(user); ok != expected[index] {
// 			if !expected[index] {
// 				t.Errorf("exception is occuring: %q", err)
// 			} else {
// 				t.Errorf("exception is not occuring: %q", err)
// 			}

// 		}
// 	}
// }
