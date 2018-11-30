package apis

//StringStruct is ...
type StringStruct struct {
	value string
}

//GetAString is ...
func (str *StringStruct) GetAString() string {
	return str.value
}

//SetAString is ...
func (str *StringStruct) SetAString(instr string) {
	str.value = instr
}

// NewStringStruct is ...
func NewStringStruct(instring string) StringAPI {
	str := StringStruct{}
	str.value = instring
	return &str
}
