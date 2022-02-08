package infra

type InfraError struct {
	Message      string
	OrignalError error
}

func (i *InfraError) Internal() bool {
	return true
}

func (i *InfraError) Error() string {
	return i.Message + ":" + i.OrignalError.Error()
}
