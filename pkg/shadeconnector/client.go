package shadeconnector

import (
	"errors"
	"fmt"
	"time"
)

type OperationEnum int

const (
	Close  OperationEnum = 0
	Open   OperationEnum = 1
	Stop   OperationEnum = 2
	Status OperationEnum = 3
)

type DeviceStatus struct {
	Mac             string `json:"mac"`
	Operation       int    `json:"operation"`
	CurrentPosition int    `json:"currentPosition"`
	CurrentAngle    int    `json:"currentAngle"`
	CurrentState    int    `json:"currentState"`
	BatteryLevel    int    `json:"batteryLevel"`
	WirelessMode    int    `json:"wirelessMode"`
	RSSI            int    `json:"RSSI"`
}

func createDeviceStatus(dsm *DeviceStatusMessage) *DeviceStatus {
	var deviceStatus DeviceStatus
	deviceStatus.BatteryLevel = dsm.Data.BatteryLevel
	deviceStatus.CurrentAngle = dsm.Data.CurrentAngle
	deviceStatus.CurrentPosition = dsm.Data.CurrentPosition
	deviceStatus.Mac = dsm.Mac
	deviceStatus.Operation = dsm.Data.Operation
	deviceStatus.WirelessMode = dsm.Data.WirelessMode
	deviceStatus.RSSI = dsm.Data.RSSI
	return &deviceStatus
}

func Operation(device_id int, operation int) (*DeviceStatus, error) {
	listDevices, err := GetDevices()
	var message *DeviceStatusMessage
	if err != nil {
		error_message := fmt.Sprintf("Cannot list devices: %s", err.Error())
		return nil, errors.New(error_message)
	}
	if device_id == 0 {
		for i := 0; i < len(listDevices.Data); i++ {
			message, err = WriteDeviceOperation(operation, listDevices.Data[i].Mac, "")
			if err != nil {
				return nil, errors.New(fmt.Sprintf("Cannot write to device:%d, error: %s", i, err.Error()))
			}
			time.Sleep(10 * time.Millisecond)
		}
	}

	if device_id > 0 && device_id >= len(listDevices.Data) {
		return nil, errors.New("device id does not exist")
	}
	message, err = WriteDeviceOperation(operation, listDevices.Data[device_id].Mac, "")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot write to device:%d, error: %s", device_id, err.Error()))
	}
	return createDeviceStatus(message), nil
}

func SetPosition(device_id int, position int) (*DeviceStatus, error) {
	listDevices, err := GetDevices()
	var message *DeviceStatusMessage
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot list devices: %s", err.Error()))
	}
	if device_id == 0 {
		for i := 0; i < len(listDevices.Data); i++ {
			WriteDeviceTargetPosition(position, listDevices.Data[i].Mac, "")
			time.Sleep(10 * time.Millisecond)
		}
	}

	if device_id > 0 && device_id >= len(listDevices.Data) {
		return nil, errors.New("device id does not exist")
	}
	message, err = WriteDeviceTargetPosition(position, listDevices.Data[device_id].Mac, "")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot write to device:%d, error: %s", device_id, err.Error()))
	}

	return createDeviceStatus(message), nil
}

func QueryStatus(device_id int) (*DeviceStatus, error) {
	listDevices, err := GetDevices()
	var message *DeviceStatusMessage
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot list devices: %s", err.Error()))
	}
	message, err = ReadDevice(listDevices.Data[device_id].Mac, "")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot write to device:%d, error: %s", device_id, err.Error()))
	}
	return createDeviceStatus(message), nil
}
