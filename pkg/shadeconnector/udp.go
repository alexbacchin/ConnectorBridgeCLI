package shadeconnector

import (
	"bytes"
	"crypto/aes"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/andreburgaud/crypt2go/ecb"
)

type Config struct {
	Port   string
	Host   string
	ApiKey string
	Token  string
}

type GetDeviceListAckMessage struct {
	MsgType         string `json:"msgType"`
	Mac             string `json:"mac"`
	DeviceType      string `json:"deviceType"`
	FwVersion       string `json:"fwVersion"`
	ProtocolVersion string `json:"ProtocolVersion"`
	Token           string `json:"token"`
	Data            []struct {
		Mac        string `json:"mac"`
		DeviceType string `json:"deviceType"`
	} `json:"data"`
}
type GetDeviceListMessage struct {
	MsgType string `json:"msgType"`
	MsgID   string `json:"msgID"`
}

type WriteDeviceMessage struct {
	MsgType     string `json:"msgType"`
	Mac         string `json:"mac"`
	DeviceType  string `json:"deviceType"`
	AccessToken string `json:"AccessToken"`
	MsgID       string `json:"msgID"`
	Data        map[string]interface{}
}

type ReadDeviceMessage struct {
	MsgType    string `json:"msgType"`
	Mac        string `json:"mac"`
	DeviceType string `json:"deviceType"`
	MsgID      string `json:"msgID"`
}

type DeviceStatusMessage struct {
	MsgType    string `json:"msgType"`
	Mac        string `json:"mac"`
	DeviceType string `json:"deviceType"`
	Data       struct {
		Type            int `json:"type"`
		Operation       int `json:"operation"`
		CurrentPosition int `json:"currentPosition"`
		CurrentAngle    int `json:"currentAngle"`
		CurrentState    int `json:"currentState"`
		VoltageMode     int `json:"voltageMode"`
		BatteryLevel    int `json:"batteryLevel"`
		WirelessMode    int `json:"wirelessMode"`
		RSSI            int `json:"RSSI"`
	} `json:"data"`
}

var cfg Config

func Init(host string, port string, apikey string) {
	cfg.Host = host
	cfg.Port = port
	cfg.ApiKey = apikey
}

func sendMessage(payload []byte, host string, port string) ([]byte, error) {
	udpServer, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		println("ResolveUDPAddr failed:", err.Error())
		return nil, err
	}
	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		println("Connection failed:", err.Error())
		return nil, err
	}

	_, err = conn.Write(payload)
	if err != nil {
		println("Write data failed:", err.Error())
		return nil, err
	}

	// buffer to get data
	received := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(time.Second * 30))
	_, err = conn.Read(received)
	if err != nil {
		println("Read data failed:", err.Error())
		return nil, err
	}

	//close the connection
	defer conn.Close()
	return received, nil
}

func GetDevices() (*GetDeviceListAckMessage, error) {
	gdl := GetDeviceListMessage{"GetDeviceList", makeTimestamp()}
	payload, err := json.Marshal(gdl)
	if err != nil {
		return nil, err
	}
	received, err := sendMessage([]byte(payload), cfg.Host, cfg.Port)

	received_trimmed := bytes.Trim(received, "\x00")
	var gdla GetDeviceListAckMessage
	err = json.Unmarshal(received_trimmed, &gdla)
	if err != nil {
		return nil, err
	}
	cfg.Token = gdla.Token
	return &gdla, nil
}

func ReadDevice(mac string, deviceType string) (*DeviceStatusMessage, error) {
	if deviceType == "" {
		deviceType = "10000000"
	}
	var readDeviceMessage ReadDeviceMessage
	readDeviceMessage.MsgType = "ReadDevice"
	readDeviceMessage.Mac = mac
	readDeviceMessage.DeviceType = deviceType
	readDeviceMessage.MsgID = makeTimestamp()

	payload, err := json.Marshal(&readDeviceMessage)
	if err != nil {
		return nil, err
	}
	received, err := sendMessage([]byte(payload), cfg.Host, cfg.Port)
	if err != nil {
		return nil, err
	}
	received_trimmed := bytes.Trim(received, "\x00")
	var deviceStatus DeviceStatusMessage
	err = json.Unmarshal(received_trimmed, &deviceStatus)
	if err != nil {
		return nil, err
	}
	return &deviceStatus, nil

}

func WriteDeviceOperation(operation int, mac string, deviceType string) (*DeviceStatusMessage, error) {
	if deviceType == "" {
		deviceType = "10000000"
	}
	var writeDeviceMessage WriteDeviceMessage
	writeDeviceMessage.DeviceType = deviceType
	writeDeviceMessage.Mac = mac
	writeDeviceMessage.Data = make(map[string]interface{}, 1)
	writeDeviceMessage.Data["operation"] = operation

	wdam, err := writeDevice(&writeDeviceMessage)
	if err != nil {
		return nil, err
	}
	return wdam, nil

}

func WriteDeviceTargetPosition(targetPosition int, mac string, deviceType string) (*DeviceStatusMessage, error) {
	if deviceType == "" {
		deviceType = "10000000"
	}
	var writeDeviceMessage WriteDeviceMessage
	writeDeviceMessage.DeviceType = deviceType
	writeDeviceMessage.Mac = mac
	writeDeviceMessage.Data = make(map[string]interface{}, 1)
	writeDeviceMessage.Data["targetPosition"] = targetPosition

	wdam, err := writeDevice(&writeDeviceMessage)
	if err != nil {
		return nil, err
	}
	return wdam, nil

}

func WriteDeviceTargetAngle(targetAngle int, mac string, deviceType string) (*DeviceStatusMessage, error) {
	if deviceType == "" {
		deviceType = "10000000"
	}
	var writeDeviceMessage WriteDeviceMessage
	writeDeviceMessage.DeviceType = deviceType
	writeDeviceMessage.Mac = mac
	writeDeviceMessage.Data = make(map[string]interface{}, 1)
	writeDeviceMessage.Data["targetAngle"] = targetAngle

	wdam, err := writeDevice(&writeDeviceMessage)
	if err != nil {
		return nil, err
	}
	return wdam, nil

}

func writeDevice(writeDeviceMessage *WriteDeviceMessage) (*DeviceStatusMessage, error) {
	writeDeviceMessage.MsgType = "WriteDevice"
	writeDeviceMessage.MsgID = makeTimestamp()
	writeDeviceMessage.AccessToken = calculateAccessToken()

	payload, err := json.Marshal(&writeDeviceMessage)
	if err != nil {
		return nil, err
	}

	received, err := sendMessage([]byte(payload), cfg.Host, cfg.Port)
	if err != nil {
		return nil, err
	}
	var wdam DeviceStatusMessage
	err = json.Unmarshal(bytes.Trim(received, "\x00"), &wdam)
	if err != nil {
		return nil, err
	}
	return &wdam, nil

}

func makeTimestamp() string {
	return strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
}

func calculateAccessToken() string {
	plaintext := []byte(cfg.Token)
	block, err := aes.NewCipher([]byte(cfg.ApiKey))
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBEncrypter(block)
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)
	return fmt.Sprintf("%X", ciphertext)
}
