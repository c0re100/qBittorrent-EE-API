# qBittorrent-EE-API
A example program for qBittorrent Enhanced Edition API

![example](https://i.imgur.com/n0YK76r.png)

## Example
I want to block client contains "Xunlei". Do the infinite loop in `GET /api/v2/sync/maindata` for the torrent hash

##### Response
```
{
  "torrents": {
    "abcdefghijklmnopqrstuvwxyz12345678900000": {
      ....
    }
  }
}
```

Next, `GET /api/v2/sync/torrentPeers?hash=abcdefghijklmnopqrstuvwxyz12345678900000`
##### Response
```
{
  "full_update": true,
  "peers": {
    "1.2.3.4:1234": {
      "client": "7.9.5.4480",
      "connection": "BT",
      "country": "China",
      "country_code": "cn",
      "dl_speed": 0,
      "downloaded": 0,
      "files": "",
      "flags": "U X E P",
      "flags_desc": "interested(peer) and unchoked(local), peer from PEX, encrypted traffic, μTP",
      "ip": "1.2.3.4",
      "peer_id": "-XL0010-",
      "port": 1234,
      "progress": 0,
      "relevance": 0,
      "up_speed": 0,
      "uploaded": 0
    }
  },
  "rid": 1,
  "show_flags": true
}
```

## Block Peer
If the peer's client contains "Xunlei", send a request to `POST /api/v2/transfer/tempBlockPeer`
Default ban time: 1hour

### Parameters
| Name | Required? | Type | Description |
| ---- | --------- | ---- | ----------- |
| ```ip``` | Required | String | IP address that you want to block |

##### Example Request
```
curl -X POST -F 'ip=1.2.3.4' http://hostname:port/api/v2/transfer/tempBlockPeer
```
