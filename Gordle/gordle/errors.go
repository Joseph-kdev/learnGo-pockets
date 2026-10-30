package gordle

// corpusError defines a sentinel error
type CorpusError string

// Error is the implementation of the error interface by corpusError
func (e CorpusError) Error() string {
	return string(e)
}