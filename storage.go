package main

import (
	"encoding/gob"
	"os"
)

type Storage struct {
	directories map[string]dir
}

type dir struct {
	Path  string
	Name  string
	Score int
}

func NewStorage() *Storage {
	d := Storage{}
	err := d.load()
	if err != nil {
		d.directories = make(map[string]dir)
	}
	return &d
}

func (s *Storage) Add(path string) {
	Compute(s.directories, path)
	s.persist()
}

func (s Storage) List() []string {
	var res []string
	for _, el := range s.directories {
		res = append(res, el.Path)
	}
	return res
}

func (s *Storage) load() error {
	var directories map[string]dir
	file, err := os.Open(storagePath())
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(&directories)
		s.directories = directories
	}
	file.Close()
	return err
}

func (s Storage) persist() error {
	file, err := os.Create(storagePath())
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(s.directories)
	}
	file.Close()
	return err
}
