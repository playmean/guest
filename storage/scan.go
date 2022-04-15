package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-billy/v5"
	"golang.org/x/exp/slices"
)

type ScanOptions struct {
	Ignore             []string
	Suffix             string
	IncludeDirectories bool
	// TODO fix a weird thing (pointer)
	Depth *int
}

func ScanDirectory(srcDir string, srcFs billy.Filesystem, opts *ScanOptions) ([]string, error) {
	paths := make([]string, 0)

	entries, err := srcFs.ReadDir(srcDir)
	if err != nil {
		return nil, err
	}

	if opts.Depth == nil {
		opts.Depth = new(int)

		*opts.Depth = -1
	}

	for _, entry := range entries {
		path := filepath.Join(srcDir, entry.Name())

		if len(opts.Ignore) > 0 && slices.Contains(opts.Ignore, path) {
			continue
		}

		fileInfo, err := srcFs.Stat(path)
		if err != nil {
			return nil, err
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if opts.IncludeDirectories {
				if len(opts.Suffix) > 0 && strings.HasSuffix(path, opts.Suffix) {
					paths = append(paths, path)
				}
			}

			if *opts.Depth > 0 {
				*opts.Depth--
			}

			if *opts.Depth == 0 {
				continue
			}

			children, err := ScanDirectory(path, srcFs, opts)
			if err != nil {
				return nil, err
			}

			paths = append(paths, children...)
		default:
			if len(opts.Suffix) > 0 && !strings.HasSuffix(path, opts.Suffix) {
				continue
			}

			paths = append(paths, path)
		}
	}

	return paths, nil
}

func ScanParent(srcDir string, srcFs billy.Filesystem, suffix string) (string, error) {
	notFoundErr := fmt.Errorf("file not found")
	depth := 1

	paths, err := ScanDirectory(srcDir, srcFs, &ScanOptions{
		Suffix:             suffix,
		IncludeDirectories: true,
		Depth:              &depth,
	})
	if err != nil {
		return "", err
	}

	if len(paths) == 0 {
		if srcDir == "" {
			return "", notFoundErr
		}

		prevDir := strings.TrimSuffix(srcDir, "/"+filepath.Base(srcDir))

		if prevDir == srcDir {
			return "", notFoundErr
		}

		return ScanParent(prevDir, srcFs, suffix)
	}

	return paths[0], nil
}
