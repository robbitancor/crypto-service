package server

type Client interface {
	GetRequest(uri string) ([]byte, error)
}
