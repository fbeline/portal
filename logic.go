package main

import (
	"errors"
	"os/user"
	"strconv"
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

func Compute(directories map[string]dir, path string) {
	path = strings.TrimSpace(path)
	tokens := strings.Split(path, "/")
	name := tokens[len(tokens)-1]
	nd, ok := directories[path]
	if ok {
		nd.Score = nd.Score + 1
		directories[path] = nd
	} else {
		directories[path] = dir{path, name, 1}
	}
}

func PrettyList(directories map[string]dir) []string {
	var res []string
	for _, el := range directories {
		line := el.Path + " : " + strconv.Itoa(el.Score)
		res = append(res, line)
	}
	return res
}

func storagePath() string {
	usr, _ := user.Current()
	return usr.HomeDir + "/.local/share/portal.gob"
}

func isValidMatch(path string, target string, targetScore int, score int) bool {
	return path != "" &&
		strings.Contains(strings.ToLower(target), strings.ToLower(path)) &&
		targetScore > score
}
