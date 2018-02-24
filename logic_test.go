package main

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func genDirectories() map[string]dir {
	directories := make(map[string]dir)
	Compute(directories, "/home/user/Downloads")
	Compute(directories, "/home/user/Documents")
	Compute(directories, "/home/user/Documents")
	return directories
}

func TestCompute(t *testing.T) {
	directories := genDirectories()
	var expected map[string]dir
	expected = make(map[string]dir)
	expected["/home/user/Downloads"] = dir{"/home/user/Downloads", "Downloads", 1}
	expected["/home/user/Documents"] = dir{"/home/user/Documents", "Documents", 2}

	assert.Equal(t, expected, directories)
}

func TestMatchSanity(t *testing.T) {
	directories := genDirectories()
	result, _ := Match(directories, "Doc")
	assert.Equal(t, "/home/user/Documents", result)
}

func TestMatchCase(t *testing.T) {
	directories := genDirectories()
	result, _ := Match(directories, "doc")
	assert.Equal(t, "/home/user/Documents", result)
}

func TestEmptyMatch(t *testing.T) {
	directories := genDirectories()
	_, err := Match(directories, "")
	assert.Error(t, err, "directory not found!")
}

func TestPrettyList(t *testing.T) {
	directories := genDirectories()
	res := PrettyList(directories)
	sort.Strings(res)
	expected := []string{"/home/user/Downloads : 1", "/home/user/Documents : 2"}
	sort.Strings(expected)
	assert.Equal(t, expected, res)
}
