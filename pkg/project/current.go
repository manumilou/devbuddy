package project

import (
	"fmt"
	"path/filepath"

	"github.com/pior/dad/pkg/manifest"
)

var ManifestFilename = "dev.yml"

func FindCurrent(path string) (*Project, error) {
	for i := 0; i < 10; i++ {
		man, err := manifest.Load(filepath.Join(path, ManifestFilename))
		if err == nil {
			return &Project{Path: path, Manifest: man}, nil
		}

		path = filepath.Dir(path)
		if path == "/" {
			break
		}
	}

	return nil, fmt.Errorf("project not found (no manifest)")
}