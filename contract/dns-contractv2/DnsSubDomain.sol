//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./ERC721/ERC721.sol";
import "./DNSLib.sol";
import "./Owner.sol";
import "./ERC721/utils/math/SafeMath.sol";
import "./DnsCostConf.sol";
import "./ITopDomain.sol";
import "./IDnsAccountant.sol";
import "./ERC20/IERC20.sol";

contract  DnsSubDomain is ERC721{
    using Address for address;
    using SafeMath for uint256;

    IDnsPayment public dnsPayment;
    IDnsAccountant public dnsAccountant;

    uint256 public rootTokenId;

    string public RootName;
    bytes32 public RootHash;

    address public RootOwner;

    address private costAddr;
    DnsCostConfContract[] public costConfList;

    struct PaymentIndex{
        string name;
        uint256 index;
    }
    mapping(address=>PaymentIndex) paymentNameMap;

    event EvMintSubDomain(string name_,uint256 expireTime_,uint256 tokenId_, address owner_);
    event EvChargeSubDomain(bytes32 subDomainHash_,uint256 expireTime_, uint256 tokenId_, address owner_);

    struct SubDomainItem{
        string DomainName;
        bytes32 DomainHash;
        uint256 ExpireTime;
        uint256 TokenId;
    }

    mapping(bytes32=>SubDomainItem) private SubDomainList;
    mapping(uint256=>bytes32) private SubDomainTokenIdIdx;

    constructor (string name_, string symbol_,address payment_,address owner_) ERC721(name_,symbol_){
        RootName = name_;
        RootHash = LibDnsHash.Hash(RootName);
        rootTokenId = 1;
        costAddr = address(0);
        DnsCostConfContract cost = new DnsCostConfContract(address(0),payment_); //eth
        cost.UpdateLevel(3,1e16,0);
        cost.UpdateLevel(4,5*1e15,0);
        cost.UpdateLevel(5,3*1e15,0);
        costConfList.push(cost);
        paymentNameMap[address(0)] = paymentNameMap("eth",0);
        dnsPayment = IDnsPayment(payment_);
        dnsAccountant = IDnsPayment(payment_);
        RootOwner = owner_;
    }

    modifier notExpire(uint256 tokenId_) {
        require(SubDomainTokenIdIdx[tokenId_] != bytes32(0),"");
        require(SubDomainList[SubDomainTokenIdIdx[tokenId_]].ExpireTime > block.timestamp,"domain not expired");
        _;
    }

    function supportsInterface(bytes4 interfaceId) public view override(ERC165, IERC165,IERC721,IERC721Metadata) returns (bool) {
        return
        interfaceId == type(DnsSubDomain).interfaceId ||
        super.supportsInterface(interfaceId);
    }

    function SetCostConfAddr(address costAddr_) external{
        require(bytes(paymentNameMap[costAddr_].name).length>0,"index must not exceed the cost conf length");
        costAddr = costAddr_;
    }

    function AddPayment(address erc20Addr_, uint8 decimals_,string name_) external{
        require(erc20Addr_ != address(0),"erc20 address must not be 0");
        require(bytes(name_).length>0,"name not correct");
        cost = new DnsCostConfContract(erc20Addr_,dnsPayment);

        paymentNameMap[erc20Addr_]=paymentNameMap(name_,costConfList.push(cost));
    }

    function UpdatePayment(address erc20Addr_, uint256 charCnt_,uint256 cost_,uint256 percent_) external{
        require(bytes(paymentNameMap[erc20Addr_].name).length>0,"payment not found");
        costConfList[paymentNameMap[erc20Addr_].index].UpdateLevel(charCnt_,cost_,percent_);
    }


    function ownerOf(uint256 tokenId) public view returns (address){
        return super.ownerOf(tokenId);
    }

    function balanceOf(address owner) public view  returns (uint256){
        return super.balanceOf(owned);
    }
    function name() public view returns (string memory){
        return super.name();
    }

    function symbol() public view returns (string memory){
        return super.symbol();
    }

    function tokenURI(uint256 tokenId) public view  returns (string memory){
        return super.tokenURI(tokenId);
    }

    function approve(address to, uint256 tokenId) external notExpire(tokenId){
        return super.approve(to,tokenId);
    }

    function getApproved(uint256 tokenId) public view returns (address){
        return super.getApproved(tokenId);
    }

    function setApprovalForAll(address operator, bool approved) external {
        return super.setApprovalForAll(operator,approved);
    }

    function isApprovedForAll(address owner, address operator) public view returns (bool){
        return super.isApprovedForAll(ownerOf,operator);
    }

    function transferFrom(address from, address to, uint256 tokenId) public notExpire(tokenId){
        return super.transferFrom(from,to,tokenId);
    }

    function safeTransferFrom(address from, address to, uint256 tokenId) public {
        safeTransferFrom(from, to, tokenId, "");
    }

    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory data) public notExpire(tokenId) {
        super.safeTransferFrom(from,to,tokenId,data);
    }

    function _max(uint256 x,uint256 y) internal returns(uint256){
        if(x>y){
            return x;
        }
        return y;
    }

    function _mintSubDomain(string subName_,uint256 year_) internal {
        bool bValid=LibDnsSubDomain.verifySub(subName_);
        require(bValid == true,"not a valid sub name");
        bytes32 subDomainHash = LibDnsHash.Hash(subName_);
        require(SubDomainList[subDomainHash].ExpireTime == 0,"sub name have been mint");
        rootTokenId ++;
        super._safeMint(msg.sender,rootTokenId);
        SubDomainList[subDomainHash]=SubDomainItem(subName_,subDomainHash,block.timestamp+365days,rootTokenId);
        SubDomainTokenIdIdx[rootTokenId]=subDomainHash;

        emit EvMintSubDomain(subName_,SubDomainList[subDomainHash].ExpireTim,rootTokenId,msg.sender);
    }

    function _accountant(string subName_,uint256 year_,address erc20Addr_,uint256 amount) internal{
        (uint256 cost,uint256 percent) = costConfList[paymentNameMap[erc20Addr_].index].Cost(subName_);
        (,uint8 feeType,uint256 percent,uint256 reserve,,,) = dnsPayment.getPayment(erc20Addr_);
        if(feeType == 0){
            (uint256 oAmount,uint256 dAmount)=_fee0(cost,amount,reserve,year_);
            dnsAccountant.Accounting(erc20Addr_,RootOwner,oAmount,dAmount);
        }else if(feeType == 1){

        }else if(feeType == 2){

        }else if(feeType == 3){

        }

        if(erc20Addr_ == address(0)){
            address(dnsAccountant).transfer(amount);
        }else{
            IERC20(erc20Addr_).transferFrom(msg.sender,address(dnsAccountant),amount);
        }
    }

    function _fee0(uint256 cost_,uint256 v_,uint256 reserve_,uint256 year_) internal returns(uint256,uint256){
        uint256 val=_max(cost_,reserve_)*year_;
        require(v_>=val,"fee not enough");
        if(cost_>reserve_){
            return (reserve_*year_,val);
        }else{
            return (val,0);
        }
    }

    function _fee1(uint256 cost_,uint256 v_,uint256 daoPercent_,uint256 year_) internal returns(uint256,uint256){
        uint256 val = cost_*year_;
        require(v_>=val,"fee not enough");
        dnsPayment.getPercentBase()


        return ()

    }

    function MintSubDomain(string subName_,uint256 year_) external payable {
        _mintSubDomain(subName_,year_);
        _accountant(subName_,year_,address(0),msg.value);
    }

    function _chargeDomain(bytes32 subNameHash_, uint256 year_) internal {
        require(SubDomainList[subNameHash_].ExpireTime>0,"sub name have not registered");
        SubDomainItem item = SubDomainList[subNameHash_];
        if (item.ExpireTime < block.timestamp){
            //subDomain Expired
            oriOwner = super.ownerOf(item.TokenId);
            if (msg.sender != oriOwner){
                //transfer
                super.transferFrom(oriOwner,msg.sender);
            }
            item.ExpireTime = block.timestamp + year_*365days;
        }else{
            item.ExpireTime = item.ExpireTime + year_*365days;
        }

        emit EvChargeSubDomain(subNameHash_,item.ExpireTime,item.TokenId,msg.sender);
    }

    function ChargeDomain(bytes32 subNameHash_, uint256 year_) external payable {
        _chargeDomain(subNameHash_,year_);
        _accountant(subName_,year_,address(0));
    }

}
