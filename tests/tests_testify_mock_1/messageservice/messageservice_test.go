package messageservice_test

import (
	"fmt"
	"testing"
	"tests_testify_mock_1/messageservice"
	"tests_testify_mock_1/messageservice/mocks"

	//"tests_testify_mock_1/messageservice/mocks"

	"github.com/stretchr/testify/mock"
)

// smsServiceMock
type SmsServiceMock struct {
	mock.Mock
}

// Our mocked smsService method
func (m *SmsServiceMock) SendChargeNotification(value int) error {
	fmt.Println("Mocked charge notification function")
	fmt.Printf("Value passed in: %d\n", value)
	// this records that the method was called and passes in the value
	// it was called with
	args := m.Called(value)
	// it then returns whatever we tell it to return
	// in this case true to simulate an SMS Service Notification
	// sent out
	return args.Error(0)
}

// TestChargeCustomer is where the magic happens
// here we create our SMSService mock
func TestChargeCustomer(t *testing.T) {
	//smsService := new(SmsServiceMock)
	smsService := &SmsServiceMock{}

	// we then define what should be returned from SendChargeNotification
	// when we pass in the value 100 to it. In this case, we want to return
	// true as it was successful in sending a notification
	smsService.On("SendChargeNotification", 100).Return(nil)

	// next we want to define the service we wish to test
	myService := messageservice.MyService{smsService}
	// and call said method
	myService.ChargeCustomer(100)

	// at the end, we verify that our myService.ChargeCustomer
	// method called our mocked SendChargeNotification method
	smsService.AssertExpectations(t)
}

func TestChargeCustomerMochery(t *testing.T) {
	//smsService := new(SmsServiceMock)
	smsService := &mocks.IMessageService{}

	// we then define what should be returned from SendChargeNotification
	// when we pass in the value 100 to it. In this case, we want to return
	// true as it was successful in sending a notification
	smsService.On("SendChargeNotification", 100).Return(nil)

	// next we want to define the service we wish to test
	myService := messageservice.MyService{smsService}
	// and call said method
	myService.ChargeCustomer(100)

	// at the end, we verify that our myService.ChargeCustomer
	// method called our mocked SendChargeNotification method
	smsService.AssertExpectations(t)
}
