package apis

import (
	"testing"

	mocks "../mock_apis"
	gomock "github.com/golang/mock/gomock"
)

func TestContainerUNIT(t *testing.T) {
	//This is a unit test because it tests a single unit IE NewContainerStruct in isolation from all other code
	//It tests a unit in isolation because all dependencies are stubbed (mocked) out
	//Contrast this test with SST below
	mockCtrl := gomock.NewController(t)
	//The following 2 lines instantiate the autogenetated stubs
	num := mocks.NewMockNumberAPI(mockCtrl)
	str := mocks.NewMockStringAPI(mockCtrl)
	//The following 2 lines provide runtime stubs for functions: GetANumber and GetAString
	//Without these stubs this test will fail!
	num.EXPECT().GetANumber().Return(24).Times(1)
	str.EXPECT().GetAString().Return("24th Street").Times(1)

	cont := NewContainerStruct(str, num)
	whoami := cont.WhoAmI()
	if whoami != "Number is '24'. String is '24th Street'" {
		t.Errorf("WhoAmI returned unexpected string '%s'. ", whoami)
	}
	//Amoungst other things, mockCtrl.Finish checks that each of the stub functions has been called the correct number of times
	///EG Try changing 'Time(1)' above to 'Time(2)' and test will fail when calling mockCtrl.Finish
	mockCtrl.Finish()

}
func TestContainerSST(t *testing.T) {
	//This is a subsystem test because it tests a number of components
	//at the same time (IE it tests NewNumberStruct, NewStringStruct, NewContainerStruct)
	num := NewNumberStruct(42)
	str := NewStringStruct("42nd Street")
	cont := NewContainerStruct(str, num)
	whoami := cont.WhoAmI()
	if whoami != "Number is '42'. String is '42nd Street'" {
		t.Errorf("WhoAmI returned unexpected string '%s'. ", whoami)
	}

}
