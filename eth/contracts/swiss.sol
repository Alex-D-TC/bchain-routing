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

    struct PORIds {

        uint128[] porId;
        uint128[] porNextID;
        uint128[] porPrevID;
    }

    struct PORKeys {

        bytes[] porPubKeyN;
        uint[] porPubKeyE;

        bytes[] porPrevPubKeyN;
        uint[] porPrevPubKeyE;
    }

    struct PORSignatures {

        bytes[] porSignature;
        bytes[] porPrevSignature;
    }

    struct Relay {

        uint128 sentBytes;
        bytes sentBytesSignature;

        bytes senderPublicKey;

        ProofOfRelay por;
    }

    event RelayHonored(address, uint);
    
    mapping(address => Relay[]) pendingToHonor;
    mapping(address => uint) nextToHonor;

    constructor() public {
    
    }


    function submitRelay(
        uint128 sentBytes, 
        bytes memory sentBytesSignature, 
        bytes memory senderPublicKey, 
        uint128[][3] memory ids,
        bytes[][2] memory keysN,
        uint[][2] memory keysE,
        bytes[][2] memory signatures) public {
        
        ProofOfRelay memory por = ProofOfRelay({
            porId: ids[0],
            porNextID: ids[1],
            porPrevID: ids[2],

            porPubKeyN: keysN[0],
            porPubKeyE: keysE[0],

            porPrevPubKeyN: keysN[1],
            porPrevPubKeyE: keysE[1],
            
            porSignature: signatures[0],
            porPrevSignature: signatures[1]
            });
        
        
        Relay memory relay = Relay(
            sentBytes, 
            sentBytesSignature, 
            senderPublicKey, 
            por);
            
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
