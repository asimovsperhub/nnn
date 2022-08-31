//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IDnsAccountant {
    function Accounting(address erc20Addr_, address receiveAddr_,uint256 amount_,uint256 dnsAmount_) public;
}