package application

import (
	"fmt"
	"os"
	"testing"

	"github.com/mradulrathore/user-management/service/repository"
	usr "github.com/mradulrathore/user-management/service/user"
)

func TestAdd(t *testing.T) {
	repo := repository.NewRepo()
	defer repo.Close()

	dataEmptyFilePath := "../service/repository/user_data_test.json"
	if err := repo.Load(dataEmptyFilePath); err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "add user", err, nil)
	}

	app := New(repo)

	user, err := usr.New("Mradul", 21, "Indore", 43, []string{"A", "B", "C", "D"})
	if err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "new user", err, nil)
	}

	if err := app.Add(user); err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "add user", err, nil)

	}

}

func TestGetAll(t *testing.T) {
	repo := repository.NewRepo()
	defer repo.Close()

	dataEmptyFilePath := "user_data_empty_test.json"
	if err := repo.Load(dataEmptyFilePath); err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "get all users", err, nil)
	}
	defer os.Remove(dataEmptyFilePath)

	app := New(repo)

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
		users, err := app.GetAll(tc.field, tc.order)
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
	repo := repository.NewRepo()
	defer repo.Close()

	dataEmptyFilePath := "user_data_empty_test.json"
	if err := repo.Load(dataEmptyFilePath); err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "delete by rollno", err, nil)
	}
	defer os.Remove(dataEmptyFilePath)

	app := New(repo)

	user, err := usr.New("Mradul", 21, "Indore", 43, []string{"A", "B", "C", "D"})
	if err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "new user", err, nil)
	}

	if err := app.Add(user); err != nil {
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
		err := app.DeleteByRollNo(tc.rollNo)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}

func TestSave(t *testing.T) {
	repo := repository.NewRepo()
	defer repo.Close()

	dataEmptyFilePath := "../service/repository/user_data_test.json"
	if err := repo.Load(dataEmptyFilePath); err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "save user", err, nil)
	}

	app := New(repo)

	if err := app.Save(); err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "save user", err, nil)

	}
}

func TestConfirmSave(t *testing.T) {
	repo := repository.NewRepo()
	defer repo.Close()

	dataEmptyFilePath := "../service/repository/user_data_test.json"
	if err := repo.Load(dataEmptyFilePath); err != nil {
		t.Errorf("Scenario: %s \n got: %v, expected: %v", "confirm save user", err, nil)
	}

	app := New(repo)

	tests := []struct {
		scenario   string
		userChoice string
		err        error
	}{{
		scenario:   "user choice(y)",
		userChoice: "y",
		err:        nil,
	}, {
		scenario:   "user choice(n)",
		userChoice: "n",
		err:        nil,
	}}

	for _, tc := range tests {
		err := app.ConfirmSave(tc.userChoice)
		if tc.err != nil && err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if tc.err == nil && err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}
	}
}
