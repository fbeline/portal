package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSave(t *testing.T) {
	ds := Disk{make(map[string]dir)}
	ds.Save("/home/user/documents0")
	ds.Save("/home/user/documents")
	ds.Save("/home/user/documents")

	var expected map[string]dir
	expected = make(map[string]dir)
	expected["/home/user/documents0"] = dir{"/home/user/documents0", "documents0", 0}
	expected["/home/user/documents"] = dir{"/home/user/documents", "documents", 1}

	assert.Equal(t, expected, ds.directories)
}

func TestMatch(t *testing.T) {
	ds := Disk{make(map[string]dir)}
	ds.Save("/home/user2/documents")
	ds.Save("/home/user/documents")
	ds.Save("/home/user/documents")

	result, _ := ds.Match("documents")

	assert.Equal(t, "/home/user/documents", result)

	nomatch, _ := ds.Match("")
	assert.Equal(t, "", nomatch)
}
