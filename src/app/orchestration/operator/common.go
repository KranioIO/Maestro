package operator

// Operator defines how an task will Proceeds
type Operator interface {
	Execute() error
	GetLogs() []string
}
