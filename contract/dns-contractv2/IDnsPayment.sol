//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


interface IDnsPayment {
    function getPayment(address erc20Addr) public view returns (uint8,uint8,uint256,uint256,string,string);
    function getPercentBase() public view returns (uint256);
}