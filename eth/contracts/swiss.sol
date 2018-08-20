pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

contract RelayHandler {

    struct Relay {

        uint128 sentBytes;
        bytes sentBytesSignature;

        bytes senderPublicKey;

        ProofOfRelay por;
    }

    struct ProofOfRelay {

        uint128[] porId;
        uint128[] porPrevID;
        uint128[] porNextID;
        
        bytes[] porPubKeyN;
        uint[] porPubKeyE;

        bytes[] porPrevPubKeyN;
        uint[] porPrevPubKeyE;
        
        bytes[] porSignature;
        bytes[] porPrevSignature;
    }

    event RelayHonored(address, uint);
    event RelayPaymentReqeusted(address, uint);
    
    mapping(address => Relay[]) relays;
    mapping(address => uint) nextToHonor;

    constructor() public {
    
    }

    function submitRelay(
        uint128 _sentBytes, 
        bytes memory _sentBytesSignature, 
        bytes memory _senderPublicKey, 
        uint128[][3] memory _ids,
        bytes[][2] memory _keysN,
        uint[][2] memory _keysE,
        bytes[][2] memory _signatures) public returns(uint) {
        
        ProofOfRelay memory por = ProofOfRelay({
            porId: _ids[0],
            porNextID: _ids[1],
            porPrevID: _ids[2],
            porPubKeyN: _keysN[0],
            porPubKeyE: _keysE[0],
            porPrevPubKeyN: _keysN[1],
            porPrevPubKeyE: _keysE[1],
            porSignature: _signatures[0],
            porPrevSignature: _signatures[1]
        });
        
        Relay memory relay = Relay(
            _sentBytes, 
            _sentBytesSignature, 
            _senderPublicKey, 
            por);
            
        address addr = addressFromBytes(relay.senderPublicKey);
        return relays[addr].push(relay);
    }

    function getRelay(address _addr, uint _id) public view returns (
        uint128 sentBytes,
        bytes sentBytesSignature,
        bytes senderPublicKey,
        
        uint128[][3] ids,
        bytes[][2] keysN,
        uint[][2] keysE,
        bytes[][2] signatures) {

        require(_id < relays[_addr].length, "Relay with the given id does not exist");

        // Messy return of relay data
        Relay storage relay = relays[_addr][_id];

        sentBytes = relay.sentBytes;
        sentBytesSignature = relay.sentBytesSignature;
        senderPublicKey = relay.senderPublicKey;
                
        ids[0] = relay.por.porId;
        ids[1] = relay.por.porNextID;
        ids[2] = relay.por.porPrevID;

        keysN[0] = relay.por.porPubKeyN;
        keysE[0] = relay.por.porPubKeyE;

        keysN[1] = relay.por.porPrevPubKeyN;
        keysE[1] = relay.por.porPrevPubKeyE;
                
        signatures[0] = relay.por.porSignature;
        signatures[1] = relay.por.porPrevSignature;
        
        return;
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
