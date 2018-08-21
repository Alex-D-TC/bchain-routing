pragma solidity ^0.4.24;

contract SimpleToken {

    address public owner;

    string public name;
    string public symbol;
    uint public maxSupply;
    uint public supply;

    mapping(address => uint) balances;

    event Transfer(address indexed from, address indexed to, uint256 value);

    constructor(string _name, string _symbol, uint _maxSupply) public {
        name = _name;
        symbol = _symbol;
        maxSupply = _maxSupply;
        owner = msg.sender;
    }

    modifier isOwner(address addr) {
        require(addr == owner);
        _;
    }

    function totalSupply() public view returns (uint256) {
        return supply;
    }

    function balanceOf(address who) public view returns (uint256) {
        return balances[who];
    }


    function sendTo(address to, uint256 value) public {
        require(balances[msg.sender] >= value, "Cannot send more coins than I own");
        
        balances[msg.sender] -= value;
        balances[to] += value;

        emit Transfer(msg.sender, to, value);
    }

    function mint(address to, uint256 value) public isOwner(msg.sender) {
        require(value + supply <= maxSupply, "Cannot mint more than the maximum allowed");
        balances[to] += value;
        supply += value;
    }

    function changeOwner(address newOwner) public isOwner(msg.sender) {
        owner = newOwner;
    }

}
