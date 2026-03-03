package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// os and filesystem APIs are easiest to learn with a temp directory.
	root, err := os.MkdirTemp("", "go-fs-demo-*")
	if err != nil {
		fmt.Println("MkdirTemp error:", err)
		return
	}
	defer os.RemoveAll(root)

	path := filepath.Join(root, "example.txt")
	if err := os.WriteFile(path, []byte("gopher\nacademy\n"), 0o644); err != nil {
		fmt.Println("WriteFile error:", err)
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("ReadFile error:", err)
		return
	}

	fmt.Println("file contents:", strings.TrimSpace(string(data)))
	fmt.Println("file base:", filepath.Base(path))
	fmt.Println("file ext:", filepath.Ext(path))

	entries, err := os.ReadDir(root)
	if err != nil {
		fmt.Println("ReadDir error:", err)
		return
	}

	for _, entry := range entries {
		fmt.Println("dir entry:", entry.Name(), "isDir:", entry.IsDir())
	}

	// io/fs style traversal works against filesystem abstractions.
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		fmt.Println("walk:", filepath.Base(path), "dir:", d.IsDir())
		return nil
	})
	if err != nil {
		fmt.Println("WalkDir error:", err)
	}
}
