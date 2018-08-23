package swiss

import (
	"github.com/alex-d-tc/bchain-routing/util"
)

func IPFSStoreRelayFile(msg *Message) (string, error) {
	json, err := util.JSONEncode(msg)
	if err != nil {
		return "", err
	}

	path, err := util.WriteToTemp(json)
	if err != nil {
		return "", err
	}

	return util.IPFSAddFile(path)
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
