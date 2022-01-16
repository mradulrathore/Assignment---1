package view

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type user struct {
	name         string
	age          int
	address      string
	rollNo       int
	courseNum    int
	coursesEnrol []string
	userChoice   string
}

func TestInit(t *testing.T) {

	testAddUser := user{
		name:         "Mradul",
		age:          21,
		address:      "Indore",
		rollNo:       43,
		courseNum:    5,
		coursesEnrol: []string{"A", "C", "D", "E", "F"},
		userChoice:   "n",
	}

	tests := []struct {
		scenario string
		req      *os.File
		err      error
	}{
		{
			scenario: "add user",
			req:      setInput("1", testAddUser),
			err:      nil,
		}}

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

func setInput(userChoice string, user user) *os.File {
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
