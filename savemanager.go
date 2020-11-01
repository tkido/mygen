package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type SaveManager struct {
	Root string
}

func NewSaveManager(root string) SaveManager {
	return SaveManager{
		Root: root,
	}
}

func (m *SaveManager) FileName(id int) string {
	return filepath.Join(".", m.Root, fmt.Sprintf("%04d.json", id))
}

func (m *SaveManager) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (m *SaveManager) Load(id int) *Character {
	var char Character
	file := m.FileName(id)
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(bs, &char); err != nil {
		log.Fatal(err)
	}
	return &char
}

func (m *SaveManager) Save() {
	bs, err := json.Marshal(g.Character)
	if err != nil {
		log.Fatal(err)
	}
	file := m.FileName(g.Character.Id)
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err = f.Write(bs); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Saved!!")
}
