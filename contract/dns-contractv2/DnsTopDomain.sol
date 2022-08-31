//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./ERC721/ERC721.sol";
//import "./ERC721/utils/math/SafeMath.sol";
import "./Owner.sol";
import "./DNSLib.sol";
import "./DnsSubDomain.sol";
import "./ITopDomain.sol";
import "./IDnsPayment.sol";
import "./IDnsAccountant.sol";

contract DnsTopDomain is ERC721,owned,IDnsPayment,IDnsAccountant{
    using LibDnsPayment for LibDnsPayment.DnsPaymentList;
    LibDnsPayment.DnsPaymentList private paymentList;

    struct RootNameItem {
        bytes   rootName;
        bool    openToPublic;
        bool    isCustom;
        uint256 customPrice;
        uint256 expireTime;
        bytes32 rootHash;
        address subContractAddr;
        uint256 tokenId;
    }

    uint256 private rootTokenId;

    mapping(bytes32=>RootNameItem) public RootNameList;
    mapping(uint256=>bytes32) public RootNameIdIndex;
    mapping(address=>bytes32) public RootNameContractIndex;

    mapping(address=>mapping(address=>uint256)) AccountAmount;
    mapping(address=>uint256) RootAccountAmount;


    event EvMintRootName(string rootName,
                        bool openToPublic,
                        bool isCustom,
                        uint256 customPrice,
                        uint256 expireTime,
                        address subContractAddr,
                        uint256 tokenId,
                        address ownerUser);

    event EvChargeRootName(bytes32 rootNameHash,
                            uint256 expireTime,
                            uint256 tokenId,
                            address ownerUser);


    constructor(string name_, string symbol_) ERC721(name_,symbol_){
        paymentList.addEthConf();
        rootTokenId = 0;
    }

    function addPayment(address erc20Addr_,uint8 decimal_,string calldata name_,string calldata symbol_,uint8 feeType_,uint256 percent_,uint256 reserve_) external{
        paymentList.insert(erc20Addr_,decimal_,name_,symbol_,feeType_,percent_,reserve_);
    }

    function updatePayment (address erc20Addr_,uint8 decimals_,uint8 feeType_, uint256 percent_, uint256 resrvCost_) external {
        paymentList.update(erc20Addr_,decimals_,feeType_,percent_,resrvCost_);
    }

    function removePayment(address erc20Addr_) external{
        paymentList.remove(erc20Addr_);
    }

    function getPayment(address erc20Addr) public view returns  (uint8,uint8,uint256,uint256,string,string){
        return paymentList.get(erc20Addr);
    }

    function setPercentBase(uint256 base_) external{
        return paymentList.setPercentBase(base_);
    }

    function getPercentBase() public view returns (uint256){
        return paymentList.percentBase;
    }

    function switchPayment(address erc20Addr_, bool switch_) external {
        if(switch_){
            paymentList.enable(erc20Addr_);
        }else{
            paymentList.disable(erc20Addr_);
        }
    }

    function supportsInterface(bytes4 interfaceId) public view override(ERC165, IERC165,IERC721,IERC721Metadata) returns (bool) {
        return
        interfaceId == type(DnsTopDomain).interfaceId ||
        super.supportsInterface(interfaceId);
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

    function MintRootName(string rootName_,uint256 year_,bool openToPublic_, bool isCustom_, uint256 customPrice_) external {
        require(LibDnsRootDomain.verifyRoot(rootName_) == true,"not a valid root name");
        bytes32 rootDomainHash = LibDnsHash.Hash(rootName_);
        require(RootNameList[rootDomainHash].ExpireTime == 0,"root name have been mint");
        rootTokenId ++;
        super._safeMint(msg.sender,rootTokenId);

        DnsSubDomain subContract = new DnsSubDomain("DnsDao",rootName_,this);

        RootNameList[rootDomainHash]=RootNameItem(rootName_,
                                              openToPublic_,
                                              isCustom_,
                                              customPrice_,
                                              block.timestamp+365days,
                                              rootDomainHash,
                                              subContract,
                                              rootTokenId);
        RootNameIdIndex[rootTokenId]=subDomainHash;
        RootNameContractIndex[subContract] = subDomainHash;

        emit EvMintRootName(rootName_,
                            openToPublic_,
                            isCustom_,
                            customPrice_,
                            RootNameList[rootDomainHash].expireTime,
                            subContract,
                            rootTokenId,
                            msg.sender);
    }

    function ChargeRootName(bytes32 rootNameHash_, uint256 year_) external {
        require(RootNameList[rootNameHash_].expireTime>0,"root name have not registered");
        RootNameItem item = RootNameList[rootNameHash_];
        if (item.ExpireTime < block.timestamp){
            //root name Expired
            oriOwner = super.ownerOf(item.tokenId);
            if (msg.sender != oriOwner){
                //transfer
                super.transferFrom(oriOwner,msg.sender);
            }
            item.expireTime = block.timestamp + year_*365days;
        }else{
            item.expireTime = item.expireTime + year_*365days;
        }

        emit EvChargeRootName(rootNameHash_,
                                item.expireTime,
                                item.tokenId,
                                msg.sender);
    }

    function DnsNameTax(address erc20Addr_, address sender_) public{
        if (erc20Addr == address(0)){
            //transfer eth to this contract
//            sender_.transfer()
        }else{
            (uint8 d,bool ip,uint256 p ,uint256 r) = paymentList.get(erc20Addr);
            if(ip == false){
                //transfer erc20 to this contract use reserve cost
//                payable(this)
            }else{

            }

            //check owner is or not expired
        }
    }

    function Accounting(address erc20Addr_, address receiveAddr_,uint256 amount_,uint256 dnsAmount_) public{
//        uint256 amount = AccountAmount[receiveAddr_][erc20Addr_];
        AccountAmount[receiveAddr_][erc20Addr_] += amount_;
        RootAccountAmount[erc20Addr_] += dnsAmount_;
    }

    function () public payable{}

}