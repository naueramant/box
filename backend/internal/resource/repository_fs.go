package resource

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	goerr "errors"

	"github.com/pkg/errors"
)

type RepositoryFilesystem struct {
	root string
}

func NewRepositoryFS(root string) Repository {
	return &RepositoryFilesystem{
		root: root,
	}
}

func (r *RepositoryFilesystem) CreateFile(path string, data io.Reader) error {
	ok, err := r.exists(path)
	if err != nil {
		return errors.Wrap(err, "Failed to check if path exists")
	}
	if ok {
		return errors.Wrap(err, "Path already exists")
	}

	out, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "Failed to create file")
	}

	defer out.Close()
	_, err = io.Copy(out, data)

	return errors.Wrap(err, "Failed to write file")
}

func (r *RepositoryFilesystem) ReadFile(path string) (io.Reader, error) {
	ok, err := r.isFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to check if path is a file")
	}
	if !ok {
		return nil, errors.Wrap(err, "Path is not a regular file")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read file")
	}

	return file, nil
}

func (r *RepositoryFilesystem) MoveFile(source string, target string) error {
	ok, err := r.isFile(source)
	if err != nil {
		return errors.Wrap(err, "Failed to check if source is a file")
	}
	if !ok {
		return errors.Wrap(err, "Source is not a regular file")
	}

	ok, err = r.exists(target)
	if err != nil {
		return errors.Wrap(err, "Failed to check if target exists")
	}
	if ok {
		return errors.Wrap(err, "Target already exists")
	}

	return errors.Wrap(
		os.Rename(source, target),
		"Failed to move file",
	)
}

func (r *RepositoryFilesystem) DeleteFile(path string) error {
	ok, err := r.isFile(path)
	if err != nil {
		return errors.Wrap(err, "Failed to check if path is a file")
	}
	if !ok {
		return errors.Wrap(err, "Path is not a regular file")
	}

	return errors.Wrap(
		os.Remove(path),
		"Failed to delete file",
	)
}

func (r *RepositoryFilesystem) CreateDir(path string) error {
	ok, err := r.exists(path)
	if err != nil {
		return errors.Wrap(err, "Failed to check if path already exists")
	}
	if ok {
		return errors.Wrap(err, "Path already exists")
	}

	return errors.Wrap(
		os.Mkdir(path, os.ModePerm),
		"Failed to create directory",
	)
}

func (r *RepositoryFilesystem) ListDir(path string) ([]Resource, error) {
	ok, err := r.isDir(path)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to check if path is a directory")
	}
	if !ok {
		return nil, errors.Wrap(err, "Path is not a directory")
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to list directory")
	}

	var resources []Resource
	for _, f := range files {
		if f.IsDir() || f.Mode().IsRegular() {
			resources = append(resources, mapFileInfoToResource(path, f))
		}
	}

	return resources, nil
}

func (r *RepositoryFilesystem) MoveDir(source string, target string) error {
	ok, err := r.isDir(source)
	if err != nil {
		return errors.Wrap(err, "Failed to check if source is a directory")
	}
	if !ok {
		return errors.Wrap(err, "Source is not a directory")
	}

	ok, err = r.exists(target)
	if err != nil {
		return errors.Wrap(err, "Failed to check if target exists")
	}
	if ok {
		return errors.Wrap(err, "Target already exists")
	}

	return errors.Wrap(
		os.Rename(source, target),
		"Failed to move directory",
	)
}

func (r *RepositoryFilesystem) DeleteDir(path string) error {
	ok, err := r.isDir(path)
	if err != nil {
		return errors.Wrap(err, "Failed to check if path is a directory")
	}
	if !ok {
		return errors.Wrap(err, "Path is not a directory")
	}

	return errors.Wrap(
		os.RemoveAll(path),
		"Failed to delete directory",
	)
}

func mapFileInfoToResource(dirPath string, fi os.FileInfo) Resource {
	return Resource{
		IsDir:     fi.IsDir(),
		Name:      fi.Name(),
		Size:      int(fi.Size()),
		Modified:  fi.ModTime(),
		Path:      path.Join(dirPath, fi.Name()),
		Extension: filepath.Ext(fi.Name()),
	}
}

func (r *RepositoryFilesystem) isDir(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return stat.IsDir(), nil
}

func (r *RepositoryFilesystem) isFile(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return stat.Mode().IsRegular(), nil
}

func (r *RepositoryFilesystem) exists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if goerr.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
