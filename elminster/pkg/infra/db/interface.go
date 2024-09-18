package db

type Connection interface {
	CreatePoolConnection() error
	CreateURL() string
}
