# qBittorrent-EE-API
A example program for qBittorrent Enhanced Edition API

![example](https://i.imgur.com/n0YK76r.png)

## Example
I want to block client contains "Xunlei". Do the infinite loop in ```GET /sync/maindata``` for the torrent hash

### Response
```
[{
  ....................
	"hash": "abcdefghijklmnopqrstuvwxyz12345678900000",
  ....................
}]
```

Next, fetch ```GET /sync/torrent_peers?hash=:hash```
### Response
```
{
	"full_update": true,
	"peers": {
		"1.2.3.4:1234": {
			"client": "Xunlei 7.9.5.4480",
			"connection": "μTP",
			"country": "China",
			"country_code": "cn",
			"dl_speed": 0,
			"downloaded": 0,
			"files": "",
			"flags": "U X E P",
			"flags_desc": "interested(peer) and unchoked(local), peer from PEX, encrypted traffic, μTP",
			"ip": "1.2.3.4:1234",
			"port": 1234,
			"progress": 0.00000001,
			"relevance": 0,
			"up_speed": 47175,
			"uploaded": 33603584
		}
	},
	"rid": 1,
	"show_flags": true
}
```

## Block Peer
If the peer's client contains "Xunlei", just send a request to ```Post /command/tempblockPeer```
Default ban time: 1hour

### Parameters
| Name | Required? | Type | Description |
| ---- | --------- | ---- | ----------- |
| ```ip``` | Required | String | IP address that you want to block |

### Example Request
```
curl -X POST -F 'ip=1.2.3.4' http://hostname:port/command/tempblockPeer
```
