pragma solidity ^0.4.24;

import "./SimpleToken.sol";

contract SwissCoin is SimpleToken {

    constructor(uint _maxSupply) SimpleToken("Swiss", "SWS", _maxSupply) public
    {
    }
}
