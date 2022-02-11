package infra

//codesを使うことでAWSError とかGCPErrorとか個別の構造は作らない方針

//Codesを使うなら下記は使用しない

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
