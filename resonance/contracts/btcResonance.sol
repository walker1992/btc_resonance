pragma solidity ^0.5.0;


contract btcResonance {
    address payable public PoolAddr;
    uint FunndationAmount;

    struct Resonance {
        address payable receiver;
        uint amount;
        bool resonanced;
        uint signcount;
    }

    mapping(address => uint) private _resonanceNods;
    mapping(bytes32 => Resonance) public resonanceTrades;

    event LogResonance( address indexed receiver,uint amount);

    modifier ResonanceSent(){
      require(msg.value > 0,"msg.value must be >0");
      _;
    }


    modifier IsSelf(){
        require(msg.sender ==PoolAddr,"nont resonancePool address!");
        _;
    }

    modifier IsRegistNode(address _nodeAddr){
        require(_resonanceNods[_nodeAddr] != 0, "Nodeaddress is not regestResonanceNode!");
        _;
    }
    modifier resonanceExists(bytes32 _resonanceId){
        require(resonanceTrades[_resonanceId].receiver != address(0), "resonanceId is not exist!");
        _;
    }

    constructor(uint amount) payable public {
        PoolAddr = msg.sender;
        FunndationAmount = amount;
    }


//   function initialFunds() payable public {

//   }

  function invitationRest() payable public {

  }

  function resonanceTrade(address payable _receiver,uint _amount) public {
      _receiver.transfer(_amount);
      emit LogResonance(_receiver,_amount);
  }

//节点确认共振，如何保证每个节点只approve一次
  function approveResonance(address payable _receiver,uint _amount,bytes32 _resonanceId)
  IsRegistNode(msg.sender)
  public
  returns (uint)
  {
      if (resonanceTrades[_resonanceId].receiver == address(0)){
          resonanceTrades[_resonanceId] = Resonance(
              _receiver,
              _amount,
              false,
              1
          );
          return resonanceTrades[_resonanceId].signcount;
      }else{
          resonanceTrades[_resonanceId].signcount = resonanceTrades[_resonanceId].signcount +1;
          if (resonanceTrades[_resonanceId].signcount >=3 && resonanceTrades[_resonanceId].resonanced == false){
              resonanceTrades[_resonanceId].resonanced =true;
              resonanceTrades[_resonanceId].receiver.transfer(resonanceTrades[_resonanceId].amount);
              emit LogResonance(_receiver,_amount);
          }
      }

  }

  //注册共振多中心化节点，需要抵押3C币
  function registResonanceNode() payable public{
     require(msg.value > 10000,"the deposit is not enough");
     _resonanceNods[msg.sender] = msg.value;
  }


}