package swiss

import (
	"github.com/alex-d-tc/bchain-routing/util"
)

func IPFSStoreRelayFile(relay IPFSRelay) (string, error) {
	json, err := util.JSONEncode(relay)
	if err != nil {
		return "", err
	}

	path, err := util.WriteToTemp(json)
	if err != nil {
		return "", err
	}

	ipfsFile, err := util.IPFSAddFile(path)
	if err != nil {
		return "", err
	}

	return ipfsFile.Hash, err
}

func IPFSReadRelayFile(addr string) ([]RelayBlock, error) {

	rawJSON, err := util.IPFSReadFile(addr)
	if err != nil {
		return nil, err
	}

	var result []RelayBlock
	err = util.JSONDecode(rawJSON, &result)
	return result, err
}
