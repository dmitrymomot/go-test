package webservices

// Logger interface
type Logger interface {
	Error(error)
	Warn(string)
	Debug(interface{})
	Info(string)
}
