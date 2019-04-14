package model

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	defaultDir = "/data/files"
)

type UserFile struct {
	Name, Type string
	User       User
	Content    []byte
	Path       string
}

func (u User) GetUserPath() string {
	var filePath string
	filePath = strings.Join([]string{defaultDir, u.Role.Name, u.Name, string(u.ID)}, "/")
	return filePath
}

type UserError struct {
	UserId   uint64
	UserRole UserRole
	errors   []string
}

func (e UserError) Error() string {
	return strings.Join(e.errors, "\n")
}

func (f UserFile) save() (string, error) {
	if f.User.Role.Save != true {
		e := UserError{
			f.User.ID,
			f.User.Role,
			[]string{"not enough rights to save files"},
		}
		return "error", e
	}
	filename := f.Name + f.Type
	if f.Path == "null" {
		f.Path = f.User.GetUserPath()
		file, err := os.Create(strings.Join([]string{f.Path, filename}, "/"))
		if err != nil {
			return "", err
		}
		n, err := file.Write(f.Content)
		if err != nil {
			return fmt.Sprintf("error %d: %s", n, err.Error()), err
		}
		return file.Name(), nil
	}
	err := ioutil.WriteFile(filename, f.Content, os.ModePerm)
	if err != nil {
		return "", err
	}
	return filename, nil
}
