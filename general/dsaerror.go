package dsaerror

type DsaError struct {
	Err string
}

// Error implements error
func (e *DsaError) Error() string {
	return e.Err
}
