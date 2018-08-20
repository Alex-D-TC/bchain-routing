pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

contract RelayHandler {

    struct ProofOfRelay {

        uint128[] porId;
        uint128[] porNextID;
        uint128[] porPrevID;

        bytes[] porPubKeyN;
        uint[] porPubKeyE;

        bytes[] porPrevPubKeyN;
        uint[] porPrevPubKeyE;
        
        bytes[] porSignature;
        bytes[] porPrevSignature;
    }

    struct Relay {

        uint128 sentBytes;
        bytes sentBytesSignature;

        bytes senderPublicKey;

        ProofOfRelay por;
    }

    event RelayHonored(address, Relay, uint);
    
    mapping(address => Relay[]) pendingToHonor;
    mapping(address => uint) nextToHonor;

    constructor() public {
    
    }


    function submitRelay(
		uint128 sentBytes, 
		bytes memory sentBytesSignature, 
		bytes memory senderPublicKey, 
		ProofOfRelay memory por) public {
        
        Relay memory relay = Relay(sentBytes, sentBytesSignature, senderPublicKey, por);
		
		address addr = addressFromBytes(relay.senderPublicKey);
        pendingToHonor[addr].push(relay);
    }

    function addressFromBytes(bytes memory key) private pure returns (address) {

        require(key.length >= 20, "The key must be of at least 20 bytes");

        uint160 result = 0;

        uint i = key.length - 1;
        for(uint iterations = 0; iterations < 20; ++iterations) {
            bytes20 b = key[i];
            result = result + uint160(b);
            result = result << 8;
            --i;
        }
        
        return address(result);
    }
    
}
