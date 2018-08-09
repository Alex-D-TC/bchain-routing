#/bin/bash

startPort=$1
endPort=$2

while [ "$startPort" -le "$endPort" ]; do
    $GOPATH/bin/bchain-routing -port $(($startPort)) -bootstrap-ip=192.168.31.113 -bootstrap-port=1025 -join=true & 
    let startPort++
done
