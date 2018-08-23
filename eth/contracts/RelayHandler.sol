pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

import "./SimpleToken.sol";

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

    struct RelayRequest {
        bool honored;
        Relay relay;
    }

    event RelayHonored(address user, uint relay, uint val);
    event RelayPaymentRequested(address user, uint relay);
    
    mapping(address => RelayRequest[]) relays;
    SimpleToken token;

    address owner;

    constructor(SimpleToken _token) public {
        token = _token;
        owner = msg.sender;
    }

    function submitRelay(
        uint128 _sentBytes,
        bytes memory _sentBytesHash, 
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
            
        RelayRequest memory request = RelayRequest({
            honored: false,
            relay: relay
        });

        address addr = addressFromBytes(relay.senderPublicKey);
        return relays[addr].push(request);
    }

    function getRelay(address _addr, uint _id) public view returns (
        uint128 sentBytes,
        bytes memory sentBytesSignature,
        bytes memory senderPublicKey,
        
        uint128[][3] memory ids,
        bytes[][2] memory keys,
        bytes[][2] memory signatures) {

        require(_id < relays[_addr].length, "Relay with the given id does not exist");

        // Messy return of relay data
        Relay storage relay = relays[_addr][_id].relay;

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
    }

    function honorRelay(address _userAddr, uint _relayId, uint _totalVal) public {
        require(_relayId < relays[_userAddr].length, "The relay must exist");
        require(!relays[_userAddr][_relayId].honored, "The relay request must not be handled beforehand");
    
        // We are highly optimistic people :>
        relays[_userAddr][_relayId].honored = true;

        // Claim the funds
        token.claimAllowance(_userAddr, _totalVal);
        
        // Send them to the relevant parties
        ProofOfRelay storage por = relays[_userAddr][_relayId].relay.por;
        
        // Split the funds (evenly for now)
        uint256 valChunk = _totalVal / por.porPubkey.length;
        uint256 valMod = _totalVal % por.porPubkey.length;

        for(uint i = 1; i < por.porPubkey.length; ++i) {
            address to = addressFromBytes(por.porPubkey[i]);
            uint256 toSend = valChunk;
            if(i <= valMod) {
                toSend += 1;
            }

            token.sendTo(to, toSend);
        }

        emit RelayHonored(_userAddr, _relayId, _totalVal);
    }

    function switchToken(SimpleToken _token) public {
        require(owner == msg.sender, "Only the owner can call this");
        token = _token;
    }

    function addressFromBytes(bytes memory _key) private pure returns (address) {

        require(_key.length >= 20, "The key must be of at least 20 bytes");

        uint160 result = 0;

        uint i = _key.length - 1;
        for(uint iterations = 0; iterations < 20; ++iterations) {
            bytes20 b = _key[i];
            result = result + uint160(b);
            result = result << 8;
            --i;
        }
        
        return address(result);
    }
}
