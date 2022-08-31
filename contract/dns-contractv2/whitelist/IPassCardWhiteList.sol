//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../Owner.sol";


interface IPassCardWhiteList {
//    function IsAWhiteMember(address user) external pure returns(bool);
    function LevelAddr(address user_) external view returns(uint8);
}


contract SimpleWhiteList is IPassCardWhiteList,owned{

    mapping(address=>uint8) whiteAddrList;
    address public Operator;
    uint8 public defaultLevel;

    event EvAddWhiteAddr(address,address,uint8);
    event EvUpdateWhiteAddr(address,address,uint8);

    modifier onlyOperator {
        require(msg.sender == Operator);
        _;
    }

    constructor(){
        Operator = msg.sender;
        defaultLevel = 3;
    }

    function AddOperator(address op) external onlyOwner{
        Operator = op;
    }

    function setDefaultLevel(uint8 cnt_) external onlyOperator{
        defaultLevel = cnt_;
    }

    function addWhiteAddr(address user_) internal {
        whiteAddrList[user_] = defaultLevel;
    }

    function AddWhiteAddress(address user_, uint8 level_) external onlyOperator{
        require(whiteAddrList[user_]==0,"user is existed");
        whiteAddrList[user_] = level_;

        emit EvAddWhiteAddr(msg.sender,user_,level_);
    }

    function UpdateWhiteAddress(address user_, uint8 level_) external onlyOperator{
        require(whiteAddrList[user_]>0,"user not added");
        whiteAddrList[user_] = level_;
        emit EvUpdateWhiteAddr(msg.sender,user_,level_);
    }

    function LevelAddr(address user_) external override view returns(uint8){
        return whiteAddrList[user_];
    }

    function AddMultiWhiteAddr(address[] calldata addrs) external onlyOperator{
        for(uint256 i=0;i<addrs.length;i++){
            addWhiteAddr(addrs[i]);
        }
    }

}