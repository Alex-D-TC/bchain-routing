pragma solidity ^0.4.24;

import "./zeppelin-contracts/token/ERC20/MintableToken.sol";
import "./zeppelin-contracts/token/ERC20/DetailedERC20.sol";


contract ICOToken is DetailedERC20, MintableToken {

    constructor(string _name, string _symbol, uint8 _decimals) DetailedERC20(_name, _symbol, _decimals)
    {
    }

}
