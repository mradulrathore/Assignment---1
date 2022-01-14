package repository

import (
	"os"
	"testing"
)

//TODO
func TestOpen(t *testing.T) {

	tests := []struct {
		scenario string
		req      *os.File
		err      error
	}{{
		scenario: "open file",
		err:      nil,
	}}

	_ = tests
}

//TODO
func TestSave(t *testing.T) {

	tests := []struct {
		scenario string
		req      *os.File
		err      error
	}{{
		scenario: "save file",
		err:      nil,
	}}

	_ = tests
}

//TODO
func TestRetrieve(t *testing.T) {

	tests := []struct {
		scenario string
		req      *os.File
		err      error
	}{{
		scenario: "retrieve file",
		err:      nil,
	}}

	_ = tests
}
