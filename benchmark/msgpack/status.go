package msgpack

//go:generate msgp
type Status struct {
	Age   int
	Name  string
	Money float64
	Areas []string
	QinFu map[string]int
}
