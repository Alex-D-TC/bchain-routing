pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

contract SwissContract {

    struct RelayBlock {
        uint128 id;
        uint128 nextID;
        uint128 prevID;

        bytes pubKey;
        bytes prevPubKey;

        bytes signature;
        bytes prevSignature;
    }

    struct Relay {
        
        uint128 sentBytes;
        bytes sentBytesSignature;

        bytes senderPublicKey;
        bytes senderPrivateKey;

        RelayBlock[] por;
    }

    constructor() public {

    }

    function submitRelay(Relay relay) public {

    }

    function addressFromBytes(bytes key) private pure returns (address) {

        require(key.length >= 20, "The key must be of at least 20 bytes");

        uint160 result = 0;

        for(uint i = 0; i < 20; ++i) {
            result = result + uint160(key[i]);
            result = result << 8;
        }
        
        return address(result);
    }
    
}
