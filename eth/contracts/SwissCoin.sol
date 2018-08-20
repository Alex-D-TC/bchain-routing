pragma solidity ^0.4.24;

import "./dep-contracts/ICOToken.sol";

contract SwissCoin is ICOToken {

    constructor(uint8 _decimals, uint _max_supply) ICOToken("Swiss", "SWS", _decimals, _max_supply)
    {
    }
}
