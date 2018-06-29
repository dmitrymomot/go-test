package infrastructure

import "log"

// NewLogger is a factory function,
// returns a new Logger struct instance
func NewLogger(debug bool) *Logger {
	return &Logger{debug}
}

// Logger interactor structure
type Logger struct {
	debug bool
}

// Error log level
func (l *Logger) Error(err error) {
	log.Println(err)
}

// Warn log level
func (l *Logger) Warn(msg string) {
	log.Println(msg)
}

// Debug log level
func (l *Logger) Debug(data interface{}) {
	if l.debug {
		log.Println(data)
	}
}

// Info log level
func (l *Logger) Info(msg string) {
	log.Println(msg)
}
