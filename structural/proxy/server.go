package structural

type Server interface {
	handleRequest(string, string) (int, string)
}