package main

import (
	"encoding/gob"
	"os"
)

type Disk struct {
	directories map[string]dir
}

type dir struct {
	Path  string
	Name  string
	Score int
}

func NewStorage() *Disk {
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

func (d Disk) List() []string {
	var res []string
	for _, el := range d.directories {
		res = append(res, el.Path)
	}
	return res
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
