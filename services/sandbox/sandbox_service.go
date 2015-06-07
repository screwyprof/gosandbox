package sandbox

import (
	"github.com/screwyprof/gosandbox/storage/sandbox"
)

type ISandboxService interface {
	LoadUserNameById(id int64) string
}

type sandbox struct {
	Storage storage.ISandboxStorage
}

func NewInstance(storage storage.ISandboxStorage) ISandboxService {
	return &sandbox{storage}
}

func (s sandbox) LoadUserNameById(id int64) string {
	return s.Storage.LoadUserNameById(id)
}
