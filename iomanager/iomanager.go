package iomanager

// IOManager - interface with contract methods to be implemented
type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data any) error
}
