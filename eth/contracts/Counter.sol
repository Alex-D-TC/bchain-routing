pragma solidity ^0.4.24;

contract Counter {

    uint public value;
    constructor() public {
        
    }
    
    function add(uint _value) external returns (uint) {
        value += _value;
        return value;
    }
    
}
