package service

// // cmd go test -coverprofile=coverage.out

// import (
// 	"reflect"
// 	"testing"

// 	usr "github.com/mradulrathore/user-management/domain/user"
// )

// func TestInsert(t *testing.T) {

// 	users = append(users, usr.User{

// 		Name:    "Anshul",
// 		Age:     20,
// 		Address: "Indore,M.P.",
// 		RollNo:  43,
// 	})

// 	users = append(users, usr.User{
// 		Name:    "Rahul",
// 		Age:     20,
// 		Address: "Indore,M.P.",
// 		RollNo:  41,
// 	})

// 	usersInput := usr.User{
// 		Name:    "Mradul",
// 		Age:     20,
// 		Address: "Indore,M.P.",
// 		RollNo:  42,
// 	}

// 	expectedUserDetails := []usr.User{
// 		{

// 			Name:    "Anshul",
// 			Age:     20,
// 			Address: "Indore,M.P.",
// 			RollNo:  43,
// 		},
// 		{
// 			Name:    "Mradul",
// 			Age:     20,
// 			Address: "Indore,M.P.",
// 			RollNo:  42,
// 		},
// 		{
// 			Name:    "Rahul",
// 			Age:     20,
// 			Address: "Indore,M.P.",
// 			RollNo:  41,
// 		},
// 	}

// 	Insert(usersInput)

// 	if !(reflect.DeepEqual(expectedUserDetails, expectedUserDetails)) {
// 		t.Errorf("Got: %v, Expected: %v", usersInput, expectedUserDetails)
// 	}

// 	usersDetailsInput := []usr.User{
// 		{
// 			Name:    "Mradul",
// 			Age:     20,
// 			Address: "Indore,M.P.",
// 			RollNo:  43,
// 		},
// 		// negative age
// 		{
// 			Name:    "Rahul",
// 			Age:     -20,
// 			Address: "Indore,M.P.",
// 			RollNo:  43,
// 		},
// 		//blank Name
// 		{
// 			Name:    "",
// 			Age:     20,
// 			Address: "Indore,M.P.",
// 			RollNo:  43,
// 		},
// 		//blank address
// 		{
// 			Name:    "Mradul",
// 			Age:     20,
// 			Address: "",
// 			RollNo:  43,
// 		},
// 		//rollno not provided
// 		{
// 			Name:    "Mradul",
// 			Age:     20,
// 			Address: "Indore",
// 		},
// 		//address not provided
// 		{
// 			Name:   "Mradul",
// 			Age:    20,
// 			RollNo: 43,
// 		},
// 	}

// 	for _, user := range usersDetailsInput {
// 		Insert(user)
// 	}

// }
