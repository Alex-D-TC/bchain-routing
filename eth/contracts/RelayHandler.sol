pragma solidity ^0.4.24;

import "./SimpleToken.sol";

contract RelayHandler {

    struct Relay {
        address sender;
        bytes senderPubkeyRaw;

        uint128 senderID;
        uint128 receiverID;
        uint128 sentBytes;

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
        address _sender,
        bytes memory _senderPubkeyRaw,
        uint128 _senderID,
        uint128 _receiverID,
        uint128 _sentBytes,
        bytes memory _ipfsRelayHash,
        address[] memory _relayers) public returns(uint) {
        
        Relay memory relay = Relay({
            sender: _sender,
            sentBytes: _sentBytes,
            senderPubkeyRaw: _senderPubkeyRaw,
            senderID: _senderID,
            receiverID: _receiverID,
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
        bytes memory senderPubkeyRaw,
        uint128 senderID,
        uint128 receiverID,
        uint128 sentBytes,
        bytes memory ipfsRelayHash,
        address[] memory relayers) {

        require(_id < relays[_addr].length, "Relay with the given id does not exist");

        // Messy return of relay data
        RelayRequest storage relay = relays[_addr][_id];

        honored = relay.honored;
        sender = relay.relay.sender;

        sentBytes = relay.relay.sentBytes;
        senderPubkeyRaw = relay.relay.senderPubkeyRaw;
        senderID = relay.relay.senderID;
        receiverID = relay.relay.receiverID;

        ipfsRelayHash = relay.relay.ipfsRelayHash;
        relayers = relay.relay.relayers;
    }

    function honorRelay(uint _totalVal) public {

        uint nextRelay = nextToHonor[msg.sender];
        
        for(; nextRelay < relays[msg.sender].length; nextRelay++) {
            if (relays[msg.sender][nextRelay].relay.sentBytes == _totalVal && !relays[msg.sender][nextRelay].honored) {
                break;
            }
        }

        // All relays have been honored
        if(nextRelay >= relays[msg.sender].length) {
            return;
        }

        assert(relays[msg.sender][nextRelay].relay.sentBytes == _totalVal);
    
        RelayRequest storage relay = relays[msg.sender][nextRelay];
        
        // OPTIMISM :>. Also in order to prevent reentry attacks from draining all reserved coins in the contract
        relay.honored = true;

        // Get the coins from the SWS contract
        token.claimAllowance(msg.sender, _totalVal);

        // Send SWS to relayers
        // If there is only one relayer, refund
        if(relay.relay.relayers.length == 1) {
            token.sendTo(relay.relay.relayers[0], _totalVal);
        } else {

            uint relayersCount = relay.relay.relayers.length; 

            uint bucketSize = _totalVal / (relayersCount - 1);
            uint bucketRemainder = _totalVal % (relayersCount - 1);
        
            for(uint i = 1; i < relayersCount; ++i) {
                
                address relayer = relay.relay.relayers[i];
                uint valToSend = bucketSize;
                
                if(i <= bucketRemainder) {
                    valToSend += 1;
                }

                token.sendTo(relayer, valToSend);
            }
        }

        // Clientside performance optimisation
        // Select the first next relay which hasn't been honored
        uint next = nextRelay + 1;
        for(; next < relay.relay.relayers.length && relays[msg.sender][next].honored; ++next) {}

        nextToHonor[msg.sender] = next;

        emit RelayHonored(msg.sender, nextRelay, _totalVal);
    }

    function relayCount() public view returns (uint) {
        return relays[msg.sender].length;
    }

    function nextUnhonoredRelay() public view returns (uint) {
        return nextToHonor[msg.sender];
    }

    function switchToken(SimpleToken _token) public {
        require(owner == msg.sender, "Only the owner can call this");
        token = _token;
    }
}
