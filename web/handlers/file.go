package handlers

import (
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
