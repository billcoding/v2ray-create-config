#!/bin/bash

vPORT=`getport`
echo "vPORT : ${vPORT}"
for LINE in `cat /root/v2ray-create-config/remote_nodes.txt`
do
    HOST=$(echo $LINE | awk -F , '{print $1}' | cat);
    PORT=$(echo $LINE | awk -F , '{print $2}' | cat);
    PASSWD=$(echo $LINE | awk -F , '{print $3}' | cat);
    echo "update node : ${HOST}";
    sshpass -p $PASSWD ssh root@${HOST} -p $PORT "bash /root/docker-compose.sh $vPORT && compose -f /root/docker-compose.yml up --build -d > /dev/null 2>&1";
done