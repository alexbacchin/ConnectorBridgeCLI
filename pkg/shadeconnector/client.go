package shadeconnector

import (
	"errors"
	"time"
)

type OperationEnum int

const (
	Close  OperationEnum = 0
	Open   OperationEnum = 1
	Stop   OperationEnum = 2
	Status OperationEnum = 3
)

func Operation(device_id int, operation int) error {
	listDevices := GetDevices()
	if device_id == 0 {
		for i := 0; i < len(listDevices.Data); i++ {
			WriteDeviceOperation(operation, listDevices.Data[i].Mac, "")
			time.Sleep(10 * time.Millisecond)
		}
	}

	if device_id > 0 && device_id >= len(listDevices.Data) {
		return errors.New("device id does not exist")
	}
	WriteDeviceOperation(operation, listDevices.Data[device_id].Mac, "")
	return nil
}

func SetPosition(device_id int, position int) error {
	listDevices := GetDevices()
	if device_id == 0 {
		for i := 0; i < len(listDevices.Data); i++ {
			WriteDeviceTargetPosition(position, listDevices.Data[i].Mac, "")
			time.Sleep(10 * time.Millisecond)
		}
	}

	if device_id > 0 && device_id >= len(listDevices.Data) {
		return errors.New("device id does not exist")
	}
	WriteDeviceTargetPosition(position, listDevices.Data[device_id].Mac, "")
	return nil
}
