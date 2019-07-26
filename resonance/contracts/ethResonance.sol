pragma solidity ^0.5.0;

contract MultisigResonance {
    address private owner;
    mapping( address => bool) private managers;
    uint private funds;




    uint constant MIN_SIGNATURES = 2;
    uint private transactionIdx;

    struct Resonance {
        address payable receiver;
        uint amount;
        uint8 signatureCount;
        mapping( address => uint8) signatures;
    }

    mapping (bytes32 => Resonance) private resonances;

    event LogResonanceCreated(address creater,bytes32 resonanceId,address receiver,uint amount);
    event LogResonanced(bytes32 resonanceId,address receiver,uint amount);

    modifier isOwner{
        require(owner == msg.sender);
        _;
    }

    modifier isManager{
        require( managers[msg.sender] == true);
        _;
    }

    modifier resonanceExists(bytes32 _resonanceId){
       require( haveResonance(_resonanceId),"resonanceId does not exist");
       _;
    }


    constructor() public {
        owner = msg.sender;
    }

    function ()  external payable{
    }

    function addManager(address _manager) public isOwner{
        managers[_manager] = true;
    }

    function removeManager(address _manager)public isOwner{
        managers[_manager] = false;
    }

    function createResonance(bytes32 _resonanceId, address payable _receiver, uint _amount)
      public
      isManager
    {
        require(address(this).balance >= _amount);
        Resonance memory resonance;
        resonance.receiver = _receiver;
        resonance.amount = _amount;
        resonance.signatureCount = 0;

        resonances[_resonanceId] = resonance;
        emit LogResonanceCreated(msg.sender,_resonanceId,_receiver,_amount);
    }


    function sigResonance(bytes32 _resonanceId)
      public
      isManager
      resonanceExists(_resonanceId)
    {
        Resonance storage resonance = resonances[_resonanceId];
        require(resonance.receiver !=address(0));
        require(msg.sender != resonance.receiver);
        require(resonance.signatures[msg.sender] != 1 );

        resonance.signatures[msg.sender] =1;
        resonance.signatureCount++;

        if (resonance.signatureCount >=MIN_SIGNATURES){
            require(address(this).balance >= resonance.amount);
            resonance.receiver.transfer(resonance.amount);

            emit LogResonanced(_resonanceId,resonance.receiver,resonance.amount);
            delete resonances[_resonanceId];
        }
    }

    function invitationRest() payable public {

    }

    function haveResonance(bytes32 _resonanceId)
      public
      view
      returns(bool exists)
    {
        exists = (resonances[_resonanceId].receiver != address(0));
    }

    function PoolBalance() public isManager view returns(uint){
        return address(this).balance;
    }
}