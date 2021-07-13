#!/bin/bash

PORT=`getport`
UUID="90f3c0e2-6fc9-466a-a6a6-77b7f80e5eee"
WS="/ws"
eval "cat << EOF
$(< /root/v2ray-create-config/nodes.tpl)
EOF
" > /root/v2ray-create-config/nodes.txt