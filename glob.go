package main

import (
	"log"
	"path/filepath"
)

type GlobManager struct {
	Cache map[string][]string
}

func NewGlobManager() GlobManager {
	return GlobManager{
		Cache: map[string][]string{},
	}
}

func (gm *GlobManager) Get(pattern string) []string {
	cached, ok := gm.Cache[pattern]
	if ok {
		return cached
	}
	paths, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatal(err)
	}
	gm.Cache[pattern] = paths
	return paths
}
