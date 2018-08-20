pragma solidity ^0.4.24;

import "./zeppelin-contracts/token/ERC20/BurnableToken.sol";
import "./zeppelin-contracts/token/ERC20/CappedToken.sol";
import "./zeppelin-contracts/token/ERC20/DetailedERC20.sol";


contract ICOToken is DetailedERC20, CappedToken, BurnableToken {

    constructor(string _name, string _symbol, uint8 _decimals, uint _max_supply) CappedToken(_max_supply * (10 ** uint256(_decimals))) DetailedERC20(_name, _symbol, _decimals)
    {
    }

}
