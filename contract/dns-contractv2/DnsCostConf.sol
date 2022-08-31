//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "./ERC721/utils/math/SafeMath.sol";
import "./ERC721/utils/Address.sol";
import "./Owner.sol";
import "./ERC20/ERC20.sol";
import "./DNSLib.sol";
import "./IDnsPayment.sol";

contract DnsCostConfContract is owned{
    using SafeMath for uint256;
    using Address for address;

    address private Erc20Addr;
    IDnsPayment private dnsPayment;

    struct SubLevel{
        uint256 CharCnt;
        uint256 Cost;
        uint256 Percent;
    }

    struct SubLevelIndex{
        uint256 index;
        uint256 used;
    }

    SubLevel[] public SubLevelItems;

    mapping(uint256=>SubLevelIndex) SubLevelMap;

    constructor (address erc20Addr_,address paymentContract_) {
        dnsPayment = IDnsPayment(paymentContract_);
        (uint8 d,,,,,)=dnsPayment.getPayment(erc20Addr_);
        require(d>0,"payment not found");
        erc20Addr_ = erc20Addr_;
    }

    function UpdateLevel(uint256 charCnt_,uint256 cost_,uint256 percent_) external OnlyOwner {
        require(charCnt_>0,"char count not correct");
        if(SubLevelMap[charCnt_].used>0){
            SubLevelItems[charCnt_].Cost = cost_;
            SubLevelItems[charCnt_].Percent = percent_;
        }else{
            SubLevelMap[charCnt_] = SubLevelIndex(SubLevelItems.push(SubLevel(charCnt_,cost_,percent_)),1);
            for (uint256 i=0;i<SubLevelItems.length;i++){
                if (SubLevelItems[i].CharCnt >= charCnt_){
                    for(uint256 j=SubLevelItems.length-1;j>i;i++){
                        SubLevelItems[j] = SubLevelItems[j-1];
                        SubLevelMap[SubLevelItems[j].CharCnt] = j;
                    }
                    SubLevelItems[i] = SubLevel(charCnt_,cost_,percent_);
                    SubLevelMap[SubLevelItems[i].CharCnt].index = i;
                    break;
                }
            }

        }
    }

    function DelLevel(uint256 charCnt_) external OnlyOwner{
        require(charCnt_>0,"char count not correct");
        require(SubLevelMap[charCnt_].used>0,"level not existed");

        uint256 index = SubLevelMap[charCnt_].index;

        for(uint i=index;i<SubLevelItems.length-1;i++){
            SubLevelItems[i] = SubLevelItems[i+1];
            SubLevelMap[SubLevelItems[i].CharCnt].index = i;
        }
        SubLevelItems.pop();
        delete SubLevelMap[charCnt_];
    }

    function _max(uint256 x, uint256 y) internal pure returns(uint256){
        if (x > y){
            return (x);
        }
        return (y);
    }

    function _cost(uint length_) internal pure returns (uint256,uint256){
        (uint8 d,,uint256 percent,uint256 reserve,,,)=dnsPayment.getPayment(Erc20Addr);
        require(d>0,"fatal error, not found payment");

        uint256 index=type(uint256).max;
        if(SubLevelMap[length_].used>0){
            index = SubLevelMap[length_].index;
        }else{
            for(uint i=SubLevelItems.length-1;i>=0;i++){
                if (length_ > SubLevelItems[i].CharCnt){
                    index = i;
                    break;
                }
            }
        }

        if(index < type(uint256).max){
            return (_max(SubLevelItems[index].Cost,reserve),_max(percent,SubLevelItems[index].Percent));
        }

        return (reserve,percent);
    }

    function Cost(string calldata subName_) external returns (uint256,uint256) {
        require(LibDnsSubDomain.verifySub(subName_),"not a correct sub Domain");
        (string sub,string rootName,) = LibDnsSubDomain.splitSubName(subName_);

        return _cost(sub.length);
    }

}