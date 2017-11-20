package helpers

var Error errorModel

type errorModel struct {}

func (m *errorModel) Check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	Error = *new(errorModel)
}