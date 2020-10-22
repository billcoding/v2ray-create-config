{
  "tag": "proxy",
  "protocol": "vmess",
  "settings": {
    "vnext": [
      {
        "address": "%s",
        "port": %d,
        "users": [
          {
            "id": "%s",
            "alterId": 32,
            "email": "t@t.tt",
            "security": "auto"
          }
        ]
      }
    ]
  },
  "streamSettings": {
    "network": "ws",
    "wsSettings": {
      "connectionReuse": true,
      "path": "%s"
    }
  },
  "mux": {
    "enabled": true
  }
}