package view

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
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

	testDisplay := displayTest{
		field:      "name",
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
			scenario: "display user",
			req:      setInputDisplay("2", testDisplay),
			err:      nil,
		}, {
			scenario: "delete user by rollno",
			req:      setInputDelete("3", testDelete),
			err:      nil,
		},
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	for _, tc := range tests {
		os.Stdin = tc.req
		err := Init()
		if err != nil && tc.err == nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		} else if err == nil && tc.err != nil {
			t.Errorf("Scenario: %s \n got: %v, expected: %v", tc.scenario, err, tc.err)
		}

		if err := tc.req.Close(); err != nil {
			log.Fatal(err)
		}
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

	defer os.Remove(tmpfile.Name())

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

	defer os.Remove(tmpfile.Name())

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

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(contentB); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	return tmpfile
}
