pragma solidity ^0.4.24;

import "./SimpleTokenImpl.sol";

contract SwissCoin is SimpleTokenImpl {

    constructor(uint _maxSupply) SimpleTokenImpl("Swiss", "SWS", _maxSupply) public
    {
    }
}
