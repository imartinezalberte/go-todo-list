package main

import (
	"net/http"
	"path/filepath"
)

const indexFileHTML = "index.html"

type CustomSystem struct {
	fs http.FileSystem
}

func (c CustomSystem) Open(path string) (http.File, error) {
	f, err := c.fs.Open(path)
	if err != nil {
		return nil, err
	}

	os, err := f.Stat()
	if os.IsDir() {
		index := filepath.Join(path, indexFileHTML)
		if _, err = c.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, err
			}
			return nil, err
		}
	}

	return f, nil
}
