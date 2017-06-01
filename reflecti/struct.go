package reflecti

type Bohler struct {
	Name  string `json:"name"`
	age   int
	Phone string `json:"phone"`
	Addr  string `json:"addr"`
}

func (b Bohler) GetAge() int {
	return b.age
}
