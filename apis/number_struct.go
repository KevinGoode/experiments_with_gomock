package apis

//NumberStruct is ...
type NumberStruct struct {
	value int
}

//GetANumber is ...
func (numb *NumberStruct) GetANumber() int {
	return numb.value
}

//SetANumber is ...
func (numb *NumberStruct) SetANumber(num int) {
	numb.value = num
}

// NewNumberStruct is ...
func NewNumberStruct(number int) NumberAPI {
	num := NumberStruct{}
	num.value = number
	return &num
}
