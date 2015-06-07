package storage

type ISandboxStorage interface {
	LoadUserNameById(id int64) string
}
