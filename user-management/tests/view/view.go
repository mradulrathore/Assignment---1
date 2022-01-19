package view

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/mradulrathore/user-management/view"
)

type userTest struct {
	name         string
	age          int
	address      string
	rollNo       int
	courseNum    int
	coursesEnrol []string
	userChoice   string
}

type displayTest struct {
	field      string
	order      int
	userChoice string
}

type deleteTest struct {
	rollNo     int
	userChoice string
}

func TestInit(t *testing.T) {
	testAddUser := userTest{
		name:         "Mradul",
		age:          21,
		address:      "Indore",
		rollNo:       43,
		courseNum:    5,
		coursesEnrol: []string{"A", "C", "D", "E", "F"},
		userChoice:   "n",
	}

	testAddUserInvalidCourse := userTest{
		name:         "Mradul",
		age:          21,
		address:      "Indore",
		rollNo:       43,
		courseNum:    5,
		coursesEnrol: []string{"A", "C", "A", "E", "F"},
		userChoice:   "n",
	}

	testAddDuplicateUser := testAddUser

	testDisplayByName := displayTest{
		field:      "name",
		order:      1,
		userChoice: "n",
	}

	testDisplayByRollNo := displayTest{
		field:      "rollno",
		order:      1,
		userChoice: "n",
	}

	testDisplayByAge := displayTest{
		field:      "age",
		order:      1,
		userChoice: "n",
	}

	testDisplayByAddress := displayTest{
		field:      "address",
		order:      1,
		userChoice: "n",
	}

	testDelete := deleteTest{
		rollNo:     43,
		userChoice: "n",
	}

	tests := []struct {
		scenario string
		req      *os.File
		err      error
	}{
		{
			scenario: "add user",
			req:      setInputAdd("1", testAddUser),
			err:      nil,
		}, {
			scenario: "add user which already exists",
			req:      setInputAddDuplicateUser("1", testAddUser, testAddDuplicateUser),
			err:      nil,
		}, {
			scenario: "add user with invalid course",
			req:      setInputAddInvalidCourse("1", testAddUserInvalidCourse),
			err:      nil,
		}, {
			scenario: "display user sorted by name",
			req:      setInputDisplay("2", testDisplayByName),
			err:      nil,
		}, {
			scenario: "display user sorted by rollno",
			req:      setInputDisplay("2", testDisplayByRollNo),
			err:      nil,
		}, {
			scenario: "display user sorted by age",
			req:      setInputDisplay("2", testDisplayByAge),
			err:      nil,
		}, {
			scenario: "display user sorted by address",
			req:      setInputDisplay("2", testDisplayByAddress),
			err:      nil,
		}, {
			scenario: "delete user by rollno",
			req:      setInputDelete("3", testDelete),
			err:      nil,
		}, {
			scenario: "save user",
			req:      setInputSave("4"),
			err:      nil,
		}, {
			scenario: "save user",
			req:      setInputInvalideConfirmChoice("5"),
			err:      nil,
		},
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	for _, tc := range tests {
		os.Stdin = tc.req
		err := view.Init()
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

		os.Remove(tc.req.Name())
	}
}

func setInputAdd(userChoice string, user userTest) *os.File {
	content := fmt.Sprintf("%s\n%s\n%d\n%s\n%d\n%d", userChoice, user.name, user.age, user.address, user.rollNo, user.courseNum)
	for i := 0; i < user.courseNum; i++ {
		content = fmt.Sprintf("%s\n%s", content, user.coursesEnrol[i])
	}
	content = fmt.Sprintf("%s\n%s\n%s\n", content, "5", user.userChoice)

	contentB := []byte(content)

	tmpfile, err := ioutil.TempFile("", "temp")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Write(contentB); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	return tmpfile
}

func setInputAddDuplicateUser(userChoice string, user userTest, user1 userTest) *os.File {
	content := fmt.Sprintf("%s\n%s\n%d\n%s\n%d\n%d", userChoice, user.name, user.age, user.address, user.rollNo, user.courseNum)
	for i := 0; i < user.courseNum; i++ {
		content = fmt.Sprintf("%s\n%s", content, user.coursesEnrol[i])
	}

	content = fmt.Sprintf("%s\n%s\n%s\n%d\n%s\n%d\n%d", content, userChoice, user1.name, user1.age, user1.address, user1.rollNo, user1.courseNum)
	for i := 0; i < user.courseNum; i++ {
		content = fmt.Sprintf("%s\n%s", content, user1.coursesEnrol[i])
	}
	content = fmt.Sprintf("%s\n%s\n%s\n", content, "5", user.userChoice)
	contentB := []byte(content)

	tmpfile, err := ioutil.TempFile("", "temp")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Write(contentB); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	return tmpfile
}

func setInputAddInvalidCourse(userChoice string, user userTest) *os.File {
	content := fmt.Sprintf("%s\n%s\n%d\n%s\n%d\n%d", userChoice, user.name, user.age, user.address, user.rollNo, user.courseNum)
	for i := 0; i < user.courseNum; i++ {
		content = fmt.Sprintf("%s\n%s", content, user.coursesEnrol[i])
	}
	content = fmt.Sprintf("%s\n%s\n%s\n", content, "5", user.userChoice)
	contentB := []byte(content)

	tmpfile, err := ioutil.TempFile("", "temp")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Write(contentB); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	return tmpfile
}

func setInputDisplay(userChoice string, display displayTest) *os.File {
	content := fmt.Sprintf("%s\n%s\n%d", userChoice, display.field, display.order)
	content = fmt.Sprintf("%s\n%s\n%s\n", content, "5", display.userChoice)

	contentB := []byte(content)

	tmpfile, err := ioutil.TempFile("", "temp")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Write(contentB); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	return tmpfile
}

func setInputDelete(userChoice string, delete deleteTest) *os.File {
	content := fmt.Sprintf("%s\n%d", userChoice, delete.rollNo)
	content = fmt.Sprintf("%s\n%s\n%s\n", content, "5", delete.userChoice)

	contentB := []byte(content)

	tmpfile, err := ioutil.TempFile("", "temp")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Write(contentB); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	return tmpfile
}

func setInputSave(userChoice string) *os.File {
	content := fmt.Sprintf("%s\n%s\n%s\n", userChoice, "5", "n")

	contentB := []byte(content)

	tmpfile, err := ioutil.TempFile("", "temp")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Write(contentB); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	return tmpfile
}

func setInputInvalideConfirmChoice(userChoice string) *os.File {
	content := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", userChoice, "5", "no", "5", "n")

	contentB := []byte(content)

	tmpfile, err := ioutil.TempFile("", "temp")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Write(contentB); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	return tmpfile
}
