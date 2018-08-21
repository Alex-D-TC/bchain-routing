solName=$1
abigen --sol ${GOPATH}/src/github.com/alex-d-tc/bchain-routing/eth/contracts/${solName}.sol --pkg eth --out ${GOPATH}/src/github.com/alex-d-tc/bchain-routing/eth/build-go/${solName}.go

