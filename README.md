
# ConnectorBridgeCLI

 
This ConnectorBridgeCLI allows for both directly controlling blinds that support wifi-connection and controlling Uni- and Bi-direction blinds that connect to a 433MHz WiFi bridge. The following bridges are reported to work with this integration: 

- CM-20 Motion Blinds bridge
- CMD-01 Motion Blinds mini-bridge
- DD7002B Connector bridge
- D1554 Connector mini-bridge
- DD7002B Brel-Home box
- D1554 Brel Home USB plug
  

This ConnectorBridgeCLI integration allows you to control the following brands:

- AMP Motorization
- Bliss Automation - Alta Window Fashions
- Bloc Blinds
- Brel Home
- 3 Day Blinds
- Diaz
- Dooya
- Gaviota
- Havana Shade
- Hurrican Shutters Wholesale
- Inspired Shades
- iSmartWindow
- Martec
- Motion Blinds
- Raven Rock MFG
- ScreenAway
- Smart Blinds
- Smart Home
- Uprise Smart Shades
- No Bull Blinds

## Configuration

The default communication is via Multicast, but it might not work on your network.  You can override with the Bridge IP address.

1. Find your bridge IP address. You might need to look at your router 
2. The bridge port is set to 32100. Unlikely the default value needs to be changed
3. Retrieve the API key from the mobile app. Settings -> About -> Tap 5 times anywhere on the screen
4. Devices IDs are usually the same as set on the remote control. 
5. Download the executable from [Releases](https://github.com/alexbacchin/ConnectorBridgeCLI/releases) page

### Variables
Flag | Environment Variable | Description  |
|--|--|--|
| --host |CONNECTOR_BRIDGE_HOST  | The hostname of IP address of the Connector Bridge -0 |
| --port |CONNECTOR_BRIDGE_PORT  | The port for the Connector Bridge connection. Default 32100 |
| --apikey | CONNECTOR_BRIDGE_APIKEY | The ApiKey to authenticate with Connector Bridge |

```bash 
sconnector-cli <command> <device_id> --host=192.168.0.189 --apikey=<Bridge API key> --port=32100
```
```bash 
export CONNECTOR_BRIDGE_HOST=192.168.0.189
export CONNECTOR_BRIDGE_APIKEY=<Bridge API key> 
sconnector-cli <command> <device_id>
```
## Commands
The following commands were added on this:

### Open, Close and Stop
Send the command to the device
```
sconnector-cli open 1 --host=192.168.0.189 --apikey=<Bridge API key> 
sconnector-cli stop 1 --host=192.168.0.189 --apikey=<Bridge API key> 
sconnector-cli close 1 --host=192.168.0.189 --apikey=<Bridge API key> 
```
### Set Position
Send the command to set the position  the device. Position must be between 0 and 100
```bash 
sconnector-cli set-position 1 50 --host=192.168.0.189 --apikey=<Bridge API key> 
```
### Server
Start a API Server to receive commands via HTTP, this is useful to integrate with other devices that do not support command line.
This command has extra settings:
| Flag | Environment Variable | Description  |
|--|--|--|
| --server-port | CONNECTOR_API_SERVER_PORT | The listening port for the web server. Default 8080 |
| --server-apikey | CONNECTOR_API_SERVER_APIKEY | The API Key to be used when requesting to the web server (Header: X-API-Key) |

Running the server
```bash 
export CONNECTOR_BRIDGE_HOST=192.168.0.189
export CONNECTOR_BRIDGE_APIKEY=<Bridge API key> 
sconnector-cli server --server-apikey=ControlMe!
```
To stop exit the server CTRL+C

#### Server Paths
The commands are available with simple GET
/open/<device_id>
/close/<device_id>
/stop/<device_id>
/position/<device_id>/<position>

Sending a request to the server:
```bash
curl -H "X-API-Key: ControlMe!" http://localhost:8080/open/1
```
### Docker
Docker images https://hub.docker.com/r/alexbacchin/sconnector-cli

