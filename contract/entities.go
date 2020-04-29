package contract

// Request is a collection of two integer parameters to be passed to the arithmetic methods
type Request struct {
	X, Y int
}

// Response is a struct with a field that will contain the result of an arithmetic operation
type Response struct {
	Result int
}
