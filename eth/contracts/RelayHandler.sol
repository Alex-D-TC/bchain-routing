pragma solidity ^0.4.24;

import "./SimpleToken.sol";

contract RelayHandler {

    struct Relay {

        address sender;
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

    event RelayHonored(address indexed honorer, uint relay, uint val);
    event RelayPaymentRequested(address indexed sender, uint relay);
    
    mapping(address => RelayRequest[]) relays;
    mapping(address => uint) nextToHonor;

    SimpleToken token;

    address owner;

    constructor(SimpleToken _token) public {
        token = _token;
        owner = msg.sender;
    }

    function submitRelay(
        uint128 _sender,
        uint128 _sentBytes,
        bytes memory _sentBytesHash, 
        bytes memory _sentBytesSignature, 
        bytes memory _senderPublicKey, 
        bytes memory _ipfsRelayHash,
        address[] memory _relayers) public returns(uint) {
        
        Relay memory relay = Relay({
            sender: _sender,
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

        uint relayId = relays[_sender].push(request) - 1;
        emit RelayPaymentRequested(_sender, relayId);
        
        return relayId;
    }

    function getRelay(address _addr, uint _id) public view returns (
        bool honored,
        address sender,
        uint128 sentBytes,
        bytes memory sentBytesSignature,
        bytes memory senderPublicKey,
        bytes memory sentBytesHash,
        bytes memory ipfsRelayHash) {

        require(_id < relays[_addr].length, "Relay with the given id does not exist");

        // Messy return of relay data
        RelayRequest storage relay = relays[_addr][_id];

        honored = relay.honored;
        sender = relay.relay.sender;

        sentBytes = relay.relay.sentBytes;
        sentBytesSignature = relay.relay.sentBytesSignature;
        senderPublicKey = relay.relay.senderPublicKey;
        sentBytesHash = relay.relay.sentBytesHash;

        ipfsRelayHash = relay.relay.ipfsRelayHash;                
    }

    function honorRelay(uint _totalVal) public {
    
        uint nextRelay = nextToHonor[msg.sender];

        // Get the next possible relay candidate
        for (; nextRelay < relays[msg.sender].length; nextRelay++) {
            if(relays[msg.sender][nextRelay].relay.sentBytes == _totalVal) {
                break;
            }
        }

        require(relays[msg.sender][nextRelay].relay.sentBytes == _totalVal);

        // We are highly optimistic people :>
        relays[msg.sender][nextRelay].honored = true;

        // Claim the funds
        token.claimAllowance(msg.sender, relays[msg.sender][nextRelay].relay.sentBytes);
        
        // Send them to the relevant parties
        address[] storage relayers = relays[msg.sender][nextRelay].relay.relayers;

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
        for(; next < relays[msg.sender].length && relays[msg.sender][next].honored; next++){}

        nextToHonor[msg.sender] = next;

        emit RelayHonored(msg.sender, nextRelay, _totalVal);
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
