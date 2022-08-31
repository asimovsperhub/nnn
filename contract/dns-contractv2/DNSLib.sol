//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "./ERC721/utils/math/SafeMath.sol";

uint8 constant MAX_DOMAIN_SEG_LEN   = 64;

function validChar(bytes1 c) internal pure returns (bool) {
    return
    (c >= 0x30 && c <= 0x39) ||
    (c >= 0x61 && c <= 0x7a) ||
    c == 0x2d ||
    c == 0x5f;
}

library LibDnsRootDomain{
    function verifyRoot(bytes memory name) public pure returns (bool) {

        if (name.length == 0 || name.length >= MAX_DOMAIN_SEG_LEN) {
            return false;
        }

        for (uint8 i = 0; i < name.length; i++) {

            if (!validChar(name[i])) {
                return false;
            }
        }

        return true;
    }
}


library LibDnsSubDomain{

    using SafeMath for uint256;

    uint16 constant MAX_SUB_DOMAIN_LEN = 192;

    bytes1 constant DOT_CHAR_VAL = 0x2e;


    function splitSubName(string memory subName_) public pure  returns (string,string,bool){
        uint256 idx = 0;

        for(uint256 i=subName_.length-1;i>=0;i++){
            if(subName_[i] == DOT_CHAR_VAL){
                return (subName_[:i],subName_[i+1:],true);
            }
        }
        return ("","",false);
    }



    function verifySub(bytes memory name) public pure returns (bool) {

        if (name.length >= MAX_SUB_DOMAIN_LEN){
            return false;
        }

        if (name[0] == DOT_CHAR_VAL || name[name.length-1] == DOT_CHAR_VAL){
            return false;
        }

        uint256 segementLength = 0;

        for (uint256 i = 0; i < name.length; i++) {

            bytes1 char = name[i];
            if (!validChar(char) && char != DOT_CHAR_VAL){
                return false;
            }

            if (char == DOT_CHAR_VAL){
                if (segementLength == 0){
                    return false;
                }

                segementLength = 0;
                continue;
            }

            segementLength += 1;
            if (segementLength >= MAX_DOMAIN_SEG_LEN){
                return false;
            }
        }

        return true;
    }
}

library LibDnsHash {

    function Hash(bytes memory name)
    internal
    pure
    returns (bytes32) {
        return keccak256(abi.encodePacked(name));
    }
}



library LibDnsPayment{
    //use reserveCost feeType = 0;
    //use percent feetype = 1;
    //reserve+(cost-reserve)*percent feetype=2;
    //if (cost-reserve)*percent < reserve then fee = reserve;
    //else fee = (cost-reserve)*percent feetype=3
    struct DnsPaymentItem {
        address ERC20Addr;
        uint8 decimals;
        string name;
        string symbol;
        uint8 feeType;
        uint256 percent;
        uint256 reserveCost;
        bool disabled;
    }

    struct DnsPaymentList {
        uint256 percentBase;
        address[] arrErc20Addr;
        mapping(address=>DnsPaymentItem) items;
    }

    event evAddEthConf(uint8 feeType,uint256 ethPercent,uint256 ethReserveCost);
    event evAddPayment(uint8 feeType, address erc20Addr,uint8 decimal,string name,string symbol, uint256 percent,uint256 reserveCost);
    event evUpdatePayment(uint8 feeType, address erc20Addr,uint256 decimal,uint256 percent,uint256 reserveCost);
    event evRemovePayment(address erc20Addr);
    event evDisablePayment(address erc20Addr);
    event evEnablePayment(address erc20Addr);

    function addEthConf(DnsPaymentList storage self) public{
        self.arrErc20Addr.push(address(0));
        if (self.percentBase == 0){
            self.percentBase = 1e8;
        }
        // percent = 1% ; reserveCost = 0.01eth
        self.items[address(0)] = DnsPaymentItem(address(0),1e18,"ethereum","eth",false,1e6,1e16);
        emit evAddEthConf(0,1e6,1e16);
    }

    function setPercentBase(DnsPaymentList storage self, uint256 percentBase_) public{
        self.percentBase = percentBase_;
    }

    function insert(DnsPaymentList storage self, address erc20Addr_, uint8 decimal_,string calldata name_, string calldata symbol_,uint8 feeType_,uint256 percent_, uint256 resrvCost_) public {
        require(erc20Addr_ != address(0),"erc20 address is not correct");
        require(self.items[erc20Addr_].decimals>0,"duplicate payment type");

        self.arrErc20Addr.push(erc20Addr_);
        self.items[erc20Addr_] = DnsPaymentItem(erc20Addr_,decimal_,name_,symbol_,feeType_,percent_,resrvCost_);
        emit evAddPayment(feeType_,erc20Addr_,decimal_,name_,percent_,resrvCost_);
    }

    function update(DnsPaymentList storage self, address erc20Addr_,uint8 decimals_,uint8 feeType_, uint256 percent_, uint256 resrvCost_) public {
        require(self.items[erc20Addr_].decimals>0,"no payment in dns payment list");

        self.items[erc20Addr_].decimals = decimals_;
        self.items[erc20Addr_].feeType = feeType_;
        self.items[erc20Addr_].percent = percent_;
        self.items[erc20Addr_].reserveCost = resrvCost_;
        emit evUpdatePayment(feeType_,erc20Addr_,decimals_,percent_,resrvCost_);
    }

    function remove(DnsPaymentList storage self,address erc20Addr_) public {
        require(self.items[erc20Addr_].decimals>0,"no payment in dns payment list");

        for (uint i=0;i<self.arrErc20Addr.length;i++){
            if (erc20Addr_ == self.arrErc20Addr[i]){
                self.arrErc20Addr[i] = self.arrErc20Addr[self.arrErc20Addr.length-1];
                self.arrErc20Addr.pop();
                break;
            }
        }

        delete self.items[erc20Addr_];
        emit evRemovePayment(erc20Addr_);
    }

    function disable(DnsPaymentList storage self, address erc20Addr_) public {
        require(self.items[erc20Addr_].decimals>0,"no payment in dns payment list");

        self.items[erc20Addr_].disabled = true;

        emit evDisablePayment(erc20Addr_);
    }

    function enable(DnsPaymentList storage self, address erc20Addr_) public{
        require(self.items[erc20Addr_].decimals>0,"no payment in dns payment list");

        self.items[erc20Addr_].disabled = false;

        emit evEnablePayment(erc20Addr_);
    }

    function get(DnsPaymentList storage self, address erc20Addr_) public view returns (uint8,uint8,uint256,uint256,string,string){
        require(self.items[erc20Addr_].decimals>0,"no payment in dns payment list");

        return (self.items[erc20Addr_].decimals,
                self.items[erc20Addr_].feeType,
                self.items[erc20Addr_].percent,
                self.items[erc20Addr_].reserveCost,
                self.items[erc20Addr_].name,
                self.items[erc20Addr_].symbol);
    }

}
