package repository

import (
	"fmt"
	"os"
	"testing"

	usr "github.com/mradulrathore/user-management/service/user"
)

func TestLoad(t *testing.T) {
	repo1 := NewRepo()
	defer repo1.Close()

	dataEmptyFilePath := "user_data_empty_test.json"
	if err := repo1.Load(dataEmptyFilePath); err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "load data", err, nil)
	}
	defer os.Remove(repo1.file.Name())

	repo2 := NewRepo()
	defer repo2.Close()

	dataFilePath := "user_data_test.json"
	if err := repo2.Load(dataFilePath); err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "load data", err, nil)
	}
}

func TestAdd(t *testing.T) {
	repo := NewRepo()
	defer repo.Close()

	dataEmptyFilePath := "user_data_empty_test.json"
	if err := repo.Load(dataEmptyFilePath); err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "load data", err, nil)
	}
	defer os.Remove(dataEmptyFilePath)

	user, err := usr.New("Mradul", 21, "Indore", 43, []string{"A", "B", "C", "D"})
	if err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "new user", err, nil)
	}

	userAlreadyExist, err := usr.New("Rahul", 24, "Indore", 43, []string{"A", "B", "C", "D", "E"})
	if err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "new user", err, nil)
	}

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
	if err := repo.Load(dataEmptyFilePath); err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "load data", err, nil)
	}
	defer os.Remove(dataEmptyFilePath)

	tests := []struct {
		scenario string
		field    string
		order    int
		err      error
	}{{
		scenario: "get all users details in ascending order of name",
		field:    "name",
		order:    1,
		err:      nil,
	}, {
		scenario: "get all users details in descending order of name",
		field:    "name",
		order:    2,
		err:      nil,
	}, {
		scenario: "get all users details in ascending order of age",
		field:    "age",
		order:    1,
		err:      nil,
	}, {
		scenario: "get all users details in descending order of age",
		field:    "age",
		order:    2,
		err:      nil,
	}, {
		scenario: "get all users details in ascending order of address",
		field:    "address",
		order:    1,
		err:      nil,
	}, {
		scenario: "get all users details in descending order of address",
		field:    "address",
		order:    2,
		err:      nil,
	}, {
		scenario: "get all users details in ascending order of rollno",
		field:    "rollno",
		order:    1,
		err:      nil,
	}, {
		scenario: "get all users details in descending order of rollno",
		field:    "rollno",
		order:    2,
		err:      nil,
	}}

	for _, tc := range tests {
		users, err := repo.GetAll(tc.field, tc.order)
		if err != tc.err {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
		if err := checkSortingOrder(tc.field, tc.order, users); err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

func checkSortingOrder(field string, order int, users []usr.User) error {
	if order == 1 {
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
	} else if order == 2 {
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
	if err := repo.Load(dataEmptyFilePath); err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "load data", err, nil)
	}
	defer os.Remove(dataEmptyFilePath)

	user, err := usr.New("Mradul", 21, "Indore", 43, []string{"A", "B", "C", "D"})
	if err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "new user", err, nil)
	}

	if err := repo.Add(user); err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "addd user", err, nil)
	}

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
		err := repo.DeleteByRollNo(tc.rollNo)
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
	if err := repo.Load(dataEmptyFilePath); err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "load data", err, nil)
	}
	defer os.Remove(repo.file.Name())

	user, err := usr.New("Mradul", 21, "Indore", 43, []string{"A", "B", "C", "D"})
	if err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "new user", err, nil)
	}

	if err := repo.Add(user); err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "add user", err, nil)
	}

	users, err := repo.GetAll("name", 1)
	if err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "get all user", err, nil)
	}

	if err := repo.Save(users); err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "save data", err, nil)
	}

}
