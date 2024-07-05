package core

type Service interface {
	Init() error
	Dispose() error
}
