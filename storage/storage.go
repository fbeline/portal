package storage

import (
	"encoding/gob"
	"errors"
	"os"
	"os/user"
	"strings"
)

type Disk struct {
	directories map[string]dir
}

type dir struct {
	Path  string
	Name  string
	Score int
}

func Create() *Disk {
	d := Disk{}
	err := d.load()
	if err != nil {
		d.directories = make(map[string]dir)
	}
	return &d
}

func (d *Disk) Add(path string) {
	d.Save(path)
	d.persist()
}

func (d Disk) Match(path string) (string, error) {
	result := dir{Score: -1}
	for k := range d.directories {
		if isValidMatch(path, k, d.directories[k].Score, result.Score) {
			result = d.directories[k]
		}
	}

	if result.Path == "" {
		return result.Path, errors.New("directory not found!")
	} else {
		return result.Path, nil
	}
}

func (d Disk) List() []string {
	var res []string
	for _, el := range d.directories {
		res = append(res, el.Path)
	}
	return res
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

func (d *Disk) load() error {
	var directories map[string]dir
	file, err := os.Open(storagePath())
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(&directories)
		d.directories = directories
	}
	file.Close()
	return err
}

func (d Disk) persist() error {
	file, err := os.Create(storagePath())
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(d.directories)
	}
	file.Close()
	return err
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
