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
	Root      string
	ClipBoard []byte
}

func NewSaveManager(root string) SaveManager {
	return SaveManager{
		Root:      root,
		ClipBoard: []byte{},
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

func (m *SaveManager) Copy() {
	data := g.Character.StatusMap[g.StatusMenu.Status]
	bs, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	m.ClipBoard = bs
	fmt.Println("Copied!!")
}

func (m *SaveManager) Paste() {
	var data Status
	bs := m.ClipBoard
	if len(bs) == 0 {
		return
	}
	if err := json.Unmarshal(bs, &data); err != nil {
		log.Fatal(err)
	}
	g.Character.StatusMap[g.StatusMenu.Status] = data
	fmt.Println("Pasted!!")
}
