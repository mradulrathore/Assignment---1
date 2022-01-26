package repository

import (
	"fmt"
	"os"
	"testing"

	usr "github.com/mradulrathore/user-management/service/user"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	repo1 := NewRepo()
	defer repo1.Close()

	dataEmptyFilePath := "user_data_empty_test.json"
	err := repo1.Load(dataEmptyFilePath)
	require.Nil(t, err)

	defer os.Remove(repo1.file.Name())

	repo2 := NewRepo()
	defer repo2.Close()

	dataFilePath := "user_data_test.json"
	err = repo2.Load(dataFilePath)
	require.Nil(t, err)
}

func TestAdd(t *testing.T) {
	repo := NewRepo()
	defer repo.Close()

	dataEmptyFilePath := "user_data_empty_test.json"
	err := repo.Load(dataEmptyFilePath)
	require.Nil(t, err)

	defer os.Remove(dataEmptyFilePath)

	user, err := usr.New("Mradul", 21, "Indore", 43, []string{"A", "B", "C", "D"})
	require.Nil(t, err)

	userAlreadyExist, err := usr.New("Rahul", 24, "Indore", 43, []string{"A", "B", "C", "D", "E"})
	require.Nil(t, err)

	tests := []struct {
		scenario string
		req      usr.User
		err      error
	}{{
		scenario: "add user with proper input",
		req:      user,
		err:      nil,
	}, {
		scenario: "add user which already exist",
		req:      userAlreadyExist,
		err:      fmt.Errorf("user already exists"),
	}}

	for _, tc := range tests {
		err := repo.Add(tc.req)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}

}

func TestGetAll(t *testing.T) {
	repo := NewRepo()
	defer repo.Close()

	dataEmptyFilePath := "user_data_empty_test.json"
	err := repo.Load(dataEmptyFilePath)
	require.Nil(t, err)

	defer os.Remove(dataEmptyFilePath)

	tests := []struct {
		scenario string
		field    string
		ASCOrder bool
		err      error
	}{{
		scenario: "get all users details in ascending order of name",
		field:    "name",
		ASCOrder: true,
		err:      nil,
	}, {
		scenario: "get all users details in descending order of name",
		field:    "name",
		ASCOrder: false,
		err:      nil,
	}, {
		scenario: "get all users details in ascending order of age",
		field:    "age",
		ASCOrder: true,
		err:      nil,
	}, {
		scenario: "get all users details in descending order of age",
		field:    "age",
		ASCOrder: false,
		err:      nil,
	}, {
		scenario: "get all users details in ascending order of address",
		field:    "address",
		ASCOrder: true,
		err:      nil,
	}, {
		scenario: "get all users details in descending order of address",
		field:    "address",
		ASCOrder: false,
		err:      nil,
	}, {
		scenario: "get all users details in ascending order of rollno",
		field:    "rollno",
		ASCOrder: true,
		err:      nil,
	}, {
		scenario: "get all users details in descending order of rollno",
		field:    "rollno",
		ASCOrder: false,
		err:      nil,
	}}

	for _, tc := range tests {
		users, err := repo.List(tc.field, tc.ASCOrder)
		if err != tc.err {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
		if err := checkSortingOrder(tc.field, tc.ASCOrder, users); err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

func checkSortingOrder(field string, ASCOrder bool, users []usr.User) error {
	if ASCOrder {
		//ascending
		switch field {
		case "name ":
			for i := 1; i < len(users); i++ {
				if users[i].Name < users[i-1].Name {
					return fmt.Errorf("not in ascending order of name")
				}
			}
		case "age":
			for i := 1; i < len(users); i++ {
				if users[i].Age < users[i-1].Age {
					return fmt.Errorf("not in ascending order of age")
				}
			}
		case "rollno":
			for i := 1; i < len(users); i++ {
				if users[i].RollNo < users[i-1].RollNo {
					return fmt.Errorf("not in ascending order of rollno")
				}
			}
		case "address":
			for i := 1; i < len(users); i++ {
				if users[i].Address < users[i-1].Address {
					return fmt.Errorf("not in ascending order of address")
				}
			}
		}
	} else {
		//descending
		switch field {
		case "name ":
			for i := 1; i < len(users); i++ {
				if users[i].Name > users[i-1].Name {
					return fmt.Errorf("not in descending order of name")
				}
			}
		case "age":
			for i := 1; i < len(users); i++ {
				if users[i].Age > users[i-1].Age {
					return fmt.Errorf("not in descending order of age")
				}
			}
		case "rollno":
			for i := 1; i < len(users); i++ {
				if users[i].RollNo > users[i-1].RollNo {
					return fmt.Errorf("not in descending order of rollno")
				}
			}
		case "address":
			for i := 1; i < len(users); i++ {
				if users[i].Address > users[i-1].Address {
					return fmt.Errorf("not in descending order of address")
				}
			}
		}
	}
	return nil
}

func TestDeleteByRollNo(t *testing.T) {
	repo := NewRepo()
	defer repo.Close()

	dataEmptyFilePath := "user_data_empty_test.json"
	err := repo.Load(dataEmptyFilePath)
	require.Nil(t, err)

	defer os.Remove(dataEmptyFilePath)

	user, err := usr.New("Mradul", 21, "Indore", 43, []string{"A", "B", "C", "D"})
	require.Nil(t, err)

	err = repo.Add(user)
	require.Nil(t, err)

	tests := []struct {
		scenario string
		rollNo   int
		err      error
	}{{
		scenario: "delete user details when user exist",
		rollNo:   43,
		err:      nil,
	}, {
		scenario: "delete user details when user doesn't exist",
		rollNo:   42,
		err:      fmt.Errorf("user does not exist"),
	}}

	for _, tc := range tests {
		err := repo.Delete(tc.rollNo)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

func TestSave(t *testing.T) {
	repo := NewRepo()
	defer repo.Close()

	dataEmptyFilePath := "user_data_empty_test.json"
	err := repo.Load(dataEmptyFilePath)
	require.Nil(t, err)

	defer os.Remove(repo.file.Name())

	user, err := usr.New("Mradul", 21, "Indore", 43, []string{"A", "B", "C", "D"})
	require.Nil(t, err)

	err = repo.Add(user)
	require.Nil(t, err)

	users, err := repo.List("name", true)
	require.Nil(t, err)

	err = repo.Save(users)
	require.Nil(t, err)
}
