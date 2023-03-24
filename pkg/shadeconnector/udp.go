package shadeconnector

import (
	"bytes"
	"crypto/aes"
	"encoding/json"
	"fmt"
	"net"
	"os"
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

func sendMessage(payload []byte, host string, port string) []byte {
	udpServer, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		println("ResolveUDPAddr failed:", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		println("Connection failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write(payload)
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}

	// buffer to get data
	received := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(time.Second * 30))
	_, err = conn.Read(received)
	if err != nil {
		println("Read data failed:", err.Error())
		os.Exit(1)
	}

	//close the connection
	defer conn.Close()
	return received
}

func GetDevices() GetDeviceListAckMessage {
	gdl := GetDeviceListMessage{"GetDeviceList", makeTimestamp()}
	payload, err := json.Marshal(gdl)
	if err != nil {
		panic(err)
	}
	received := bytes.Trim(sendMessage([]byte(payload), cfg.Host, cfg.Port), "\x00")
	var gdla GetDeviceListAckMessage
	err = json.Unmarshal(received, &gdla)
	if err != nil {
		panic(err)
	}
	cfg.Token = gdla.Token
	return gdla
}

func ReadDevice(mac string, deviceType string) DeviceStatusMessage {
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
		panic(err)
	}
	received := bytes.Trim(sendMessage([]byte(payload), cfg.Host, cfg.Port), "\x00")
	var deviceStatus DeviceStatusMessage
	err = json.Unmarshal(received, &deviceStatus)
	if err != nil {
		panic(err)
	}
	return deviceStatus

}

func WriteDeviceOperation(operation int, mac string, deviceType string) DeviceStatusMessage {
	if deviceType == "" {
		deviceType = "10000000"
	}
	var writeDeviceMessage WriteDeviceMessage
	writeDeviceMessage.DeviceType = deviceType
	writeDeviceMessage.Mac = mac
	writeDeviceMessage.Data = make(map[string]interface{}, 1)
	writeDeviceMessage.Data["operation"] = operation

	return writeDevice(&writeDeviceMessage)

}

func WriteDeviceTargetPosition(targetPosition int, mac string, deviceType string) DeviceStatusMessage {
	if deviceType == "" {
		deviceType = "10000000"
	}
	var writeDeviceMessage WriteDeviceMessage
	writeDeviceMessage.DeviceType = deviceType
	writeDeviceMessage.Mac = mac
	writeDeviceMessage.Data = make(map[string]interface{}, 1)
	writeDeviceMessage.Data["targetPosition"] = targetPosition

	return writeDevice(&writeDeviceMessage)

}

func WriteDeviceTargetAngle(targetAngle int, mac string, deviceType string) DeviceStatusMessage {
	if deviceType == "" {
		deviceType = "10000000"
	}
	var writeDeviceMessage WriteDeviceMessage
	writeDeviceMessage.DeviceType = deviceType
	writeDeviceMessage.Mac = mac
	writeDeviceMessage.Data = make(map[string]interface{}, 1)
	writeDeviceMessage.Data["targetAngle"] = targetAngle

	return writeDevice(&writeDeviceMessage)

}

func writeDevice(writeDeviceMessage *WriteDeviceMessage) DeviceStatusMessage {
	writeDeviceMessage.MsgType = "WriteDevice"
	writeDeviceMessage.MsgID = makeTimestamp()
	writeDeviceMessage.AccessToken = calculateAccessToken()

	payload, err := json.Marshal(&writeDeviceMessage)
	if err != nil {
		panic(err)
	}
	received := bytes.Trim(sendMessage([]byte(payload), cfg.Host, cfg.Port), "\x00")
	var wdam DeviceStatusMessage
	err = json.Unmarshal(received, &wdam)
	if err != nil {
		panic(err)
	}
	return wdam

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
