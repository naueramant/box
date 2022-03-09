package resource

import "io"

type Repository interface {
	CreateFile(path string, data io.Reader) (err error)
	ReadFile(path string) (data io.Reader, err error)
	MoveFile(sourcePath string, targetPath string) (err error)
	DeleteFile(path string) (err error)
	CreateDir(path string) (err error)
	ListDir(path string) (resources []Resource, err error)
	MoveDir(sourcePath string, targetPath string) (err error)
	DeleteDir(path string) (err error)
}
