package hello

// Here, we're exercising the use of interfaces.
// Both Hello1 and Hello2 classes implement the interface
// HelloWriter. We can use one or another based on any variable.

type HelloWriter interface {
	WriteHello() string
}

type Hello1 struct {
	defaultMessage string
}

type Hello2 struct {
	defaultMessage string
}

func (h Hello1) WriteHello() string {
	return "This is Hello 1"
}

func (h Hello2) WriteHello() string {
	return "This is Hello 2"
}
