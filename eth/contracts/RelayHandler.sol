pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

import "./SimpleToken.sol";

contract RelayHandler {

    struct Relay {

        uint128 sentBytes;
        bytes sentBytesHash;
        bytes sentBytesSignature;
        bytes senderPublicKey;

        bytes ipfsRelayHash;
        address[] relayers;
    }

    struct RelayRequest {
        bool honored;
        Relay relay;
    }

    event RelayHonored(address indexed user, uint relay, uint val);
    event RelayPaymentRequested(address indexed user, uint relay);
    
    mapping(address => RelayRequest[]) relays;
    mapping(address => uint) nextToHonor;

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
        bytes memory _ipfsRelayHash,
        address[] memory _relayers) public returns(uint) {
        
        Relay memory relay = Relay({
            sentBytes: _sentBytes,
            sentBytesHash: _sentBytesHash,
            sentBytesSignature: _sentBytesSignature,
            senderPublicKey: _senderPublicKey,
            ipfsRelayHash: _ipfsRelayHash,
            relayers: _relayers
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
        bytes memory sentBytesHash,
        bytes memory ipfsRelayHash) {

        require(_id < relays[_addr].length, "Relay with the given id does not exist");

        // Messy return of relay data
        Relay storage relay = relays[_addr][_id].relay;

        sentBytes = relay.sentBytes;
        sentBytesSignature = relay.sentBytesSignature;
        senderPublicKey = relay.senderPublicKey;

        ipfsRelayHash = relay.ipfsRelayHash;                
    }

    function honorRelay(address _userAddr, uint _totalVal) public {
    
        uint nextRelay = nextToHonor[_userAddr];

        // Get the next possible relay candidate
        for (; nextRelay < relays[_userAddr].length; nextRelay++) {
            if(relays[_userAddr][i].relay.sentBytes == _totalVal) {
                break;
            }
        }

        require(relays[_userAddr][nextRelay].relay.sentBytes == _totalVal);

        // We are highly optimistic people :>
        relays[_userAddr][nextRelay].honored = true;

        // Claim the funds
        token.claimAllowance(_userAddr, relays[_userAddr][nextRelay].relay.sentBytes);
        
        // Send them to the relevant parties
        address[] storage relayers = relays[_userAddr][nextRelay].relay.relayers;

        // Split the funds (evenly for now)
        uint256 valChunk = _totalVal / relayers.length;
        uint256 valMod = _totalVal % relayers.length;

        for(uint i = 1; i < relayers.length; ++i) {
            address to = relayers[i];
            uint256 toSend = valChunk;
            if(i <= valMod) {
                toSend += 1;
            }

            token.sendTo(to, toSend);
        }

        uint next = nextRelay + 1;
        for(; next < relays[_userAddr].length && relays[_userAddr][next].honored; next++){}

        nextToHonor[_userAddr] = next;

        emit RelayHonored(_userAddr, nextRelay, _totalVal);
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
