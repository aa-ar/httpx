package errors

type Error interface {
	Status() int
	Error() string
}
