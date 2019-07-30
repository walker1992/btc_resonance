pragma solidity ^0.5.0;

contract Resonance {
    address private owner;
    mapping( address => bool) private managers;
    uint private funds;

    uint constant MIN_SIGNATURES = 3;

    modifier onlyOwner{
    require(owner == msg.sender);
        _;
    }

    modifier onlyManager{
        require( managers[msg.sender] == true);
        _;
    }

    constructor() payable public {
        owner = msg.sender;
        funds = msg.value;
    }

    function addManager(address _manager) public onlyOwner{
        managers[_manager] = true;
    }

    function removeManager(address _manager)public onlyOwner{
        managers[_manager] = false;
    }

    function Ismanger(address _manager)public view returns(bool){
      return managers[_manager];
    }

    event LogResonance(bytes32 resonanceId,address resonancer,uint amount);
    //锚定节点执行
    function resonanceTrade(bytes32 _resonanceId, address payable _receiver,  uint _amount, uint[] memory v, bytes32[] memory r, bytes32[] memory s)
    public
    onlyManager
    returns(bool)
    {
        require (_receiver != address(0x0));
        require(address(this).balance >= _amount);
        require (v.length == r.length);
        require (v.length == s.length);
        require(verifySignAndCount(keccak256(abi.encodePacked(_resonanceId, _receiver, _amount)),  v, r, s) >=3,"the signCount less then 3"); //签名数量

        _receiver.transfer(_amount);
        emit LogResonance(_resonanceId,_receiver,_amount);
    }

    function verifySignAndCount(bytes32 hash,  uint[] memory v, bytes32[] memory r, bytes32[] memory s) private view returns (uint8 count) {
        for (uint i = 0; i < v.length; i++ ){
            address temp = ecrecover(hash, uint8(v[i]), r[i], s[i]);
            if (managers[temp]){
                count++;
            }
        }
    }


    function poolBalance() public view returns(uint){
        return address(this).balance;
    }


    function invitationRest() payable public {

    }

    function withdraw(address payable _addr)public onlyOwner {
        _addr.transfer(address(this).balance);
    }
}