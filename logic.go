package main

import (
	"errors"
	"os/user"
	"strings"
)

func Match(directories map[string]dir, path string) (string, error) {
	result := dir{Score: -1}
	for k := range directories {
		if isValidMatch(path, k, directories[k].Score, result.Score) {
			result = directories[k]
		}
	}

	if result.Path == "" {
		return result.Path, errors.New("directory not found!")
	} else {
		return result.Path, nil
	}
}

func (d *Disk) Save(path string) {
	path = strings.TrimSpace(path)
	tokens := strings.Split(path, "/")
	name := tokens[len(tokens)-1]
	nd, ok := d.directories[path]
	if ok {
		nd.Score = nd.Score + 1
		d.directories[path] = nd
	} else {
		d.directories[path] = dir{path, name, 0}
	}
}
func storagePath() string {
	usr, _ := user.Current()
	return usr.HomeDir + "/.local/share/portal.gob"
}

func isValidMatch(path string, target string, targetScore int, score int) bool {
	return path != "" &&
		strings.Contains(target, path) &&
		targetScore > score
}
