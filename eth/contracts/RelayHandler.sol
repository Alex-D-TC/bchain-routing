pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

contract RelayHandler {

    struct Relay {

        uint128 sentBytes;
        bytes sentBytesHash;
        bytes sentBytesSignature;
        bytes senderPublicKey;

        ProofOfRelay por;
    }

    struct ProofOfRelay {

        uint128[] porID;
        uint128[] porPrevID;
        uint128[] porNextID;
        
        bytes[] porPubkey;
        bytes[] porPrevPubkey;
        
        bytes[] porSignature;
        bytes[] porPrevSignature;
        
        bytes[] porRawHash;
    }

    event RelayHonored(address, uint);
    event RelayPaymentReqeusted(address, uint);
    
    mapping(address => Relay[]) relays;
    mapping(address => uint) nextToHonor;

    constructor() public {
    }

    function submitRelay(
        uint128 _sentBytes,
        bytes _sentBytesHash, 
        bytes memory _sentBytesSignature, 
        bytes memory _senderPublicKey, 
        uint128[][3] memory _ids,
        bytes[][2] memory _keys,
        bytes[][2] memory _signatures,
        bytes[] memory _porRawHash) public returns(uint) {
        
        ProofOfRelay memory por = ProofOfRelay({
            porID: _ids[0],
            porNextID: _ids[1],
            porPrevID: _ids[2],
            porPubkey: _keys[0],
            porPrevPubkey: _keys[1],
            porSignature: _signatures[0],
            porPrevSignature: _signatures[1],
            porRawHash: _porRawHash
        });
        
        Relay memory relay = Relay({
            sentBytes: _sentBytes,
            sentBytesHash: _sentBytesHash,
            sentBytesSignature: _sentBytesSignature,
            senderPublicKey: _senderPublicKey,
            por: por
        });
            
        address addr = addressFromBytes(relay.senderPublicKey);
        return relays[addr].push(relay);
    }

    function getRelay(address _addr, uint _id) public view returns (
        uint128 sentBytes,
        bytes sentBytesSignature,
        bytes senderPublicKey,
        
        uint128[][3] ids,
        bytes[][2] keys,
        bytes[][2] signatures) {

        require(_id < relays[_addr].length, "Relay with the given id does not exist");

        // Messy return of relay data
        Relay storage relay = relays[_addr][_id];

        sentBytes = relay.sentBytes;
        sentBytesSignature = relay.sentBytesSignature;
        senderPublicKey = relay.senderPublicKey;
                
        ids[0] = relay.por.porID;
        ids[1] = relay.por.porNextID;
        ids[2] = relay.por.porPrevID;

        keys[0] = relay.por.porPubkey;
        keys[1] = relay.por.porPrevPubkey;
                
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
