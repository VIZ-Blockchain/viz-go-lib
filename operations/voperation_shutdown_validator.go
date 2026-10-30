package operations

//ShutdownValidatorOperation represents shutdown_validator operation data.
type ShutdownValidatorOperation struct {
	Owner string `json:"owner"`
}

//Type function that defines the type of operation ShutdownValidatorOperation.
func (op *ShutdownValidatorOperation) Type() OpType {
	return TypeShutdownValidator
}

//Data returns the operation data ShutdownValidatorOperation.
func (op *ShutdownValidatorOperation) Data() interface{} {
	return op
}
