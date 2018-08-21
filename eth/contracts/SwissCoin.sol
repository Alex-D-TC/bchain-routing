pragma solidity ^0.4.24;

import "./dep-contracts/ICOToken.sol";

contract SwissCoin is ICOToken {

    constructor(uint8 _decimals) ICOToken("Swiss", "SWS", _decimals) public
    {
    }
}
