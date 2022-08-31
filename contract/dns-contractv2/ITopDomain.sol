

interface ITopDomain {
    function DnsNameTax(address erc20Addr, address sender ) public ;
    function PaymentIsSupport(address erc20Addr) returns (bool);
}