package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/go-git/go-billy/v5"
	"golang.org/x/exp/slices"
)

type CopyOptions struct {
	Ignore []string
}

func CopyDirectory(scrDir, destDir string, srcFs, destFs billy.Filesystem, opts *CopyOptions) error {
	entries, err := srcFs.ReadDir(scrDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(destDir, entry.Name())

		if slices.Contains(opts.Ignore, sourcePath) {
			continue
		}

		fileInfo, err := srcFs.Stat(sourcePath)
		if err != nil {
			return err
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := CreateDirIfNotExists(destPath, 0755, destFs); err != nil {
				return err
			}
			if err := CopyDirectory(sourcePath, destPath, srcFs, destFs, opts); err != nil {
				return err
			}
		case os.ModeSymlink:
			if err := CopySymLink(sourcePath, destPath, srcFs, destFs); err != nil {
				return err
			}
		default:
			if err := CopyFile(sourcePath, destPath, srcFs, destFs); err != nil {
				return err
			}
		}
	}

	return nil
}

func CopyFile(srcFile, dstFile string, srcFs, destFs billy.Filesystem) error {
	out, err := destFs.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := srcFs.Open(srcFile)
	if err != nil {
		return err
	}

	defer in.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

func Exists(filePath string, destFs billy.Filesystem) bool {
	if _, err := destFs.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateDirIfNotExists(dir string, perm os.FileMode, destFs billy.Filesystem) error {
	if Exists(dir, destFs) {
		return nil
	}

	if err := destFs.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

func CopySymLink(source, dest string, srcFs, destFs billy.Filesystem) error {
	link, err := srcFs.Readlink(source)

	if err != nil {
		return err
	}

	return destFs.Symlink(link, dest)
}
