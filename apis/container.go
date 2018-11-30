package apis

import (
	"fmt"
)

//ContainerStruct is ...
type ContainerStruct struct {
	str StringAPI
	num NumberAPI
}

//WhoAmI prints contents to console
func (me ContainerStruct) WhoAmI() string {
	return fmt.Sprintf("Number is '%d'. String is '%s'", me.num.GetANumber(), me.str.GetAString())
}

// NewContainerStruct is ...
func NewContainerStruct(str StringAPI, num NumberAPI) ContainerStruct {
	cont := ContainerStruct{}
	cont.str = str
	cont.num = num
	return cont
}
