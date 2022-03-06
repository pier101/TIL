//SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.6.0 <0.9.0;

import "./ERC20.sol";

contract CryptosICO is Cryptos{
    address public admin;
    address payable public deposit;
    uint tokenPrice = 0.001 ether;  
    uint public hardCap = 300 ether;
    uint public raisedAmount; // wei 단위
    uint public saleStart = block.timestamp;
    uint public saleEnd = block.timestamp + 604800; //one week
    
    uint public tokenTradeStart = saleEnd + 604800; //판매 종료후 일주일 이내 거래 가능
    uint public maxInvestment = 5 ether;
    uint public minInvestment = 0.1 ether;
    
    enum State { beforeStart, running, afterEnd, halted} // ICO states 
    State public icoState;
    
    constructor(address payable _deposit){
        deposit = _deposit; 
        admin = msg.sender; 
        icoState = State.beforeStart;
    }
 
    
    modifier onlyAdmin(){
        require(msg.sender == admin);
        _;
    }
    
    
    // iCO 정지
    function halt() public onlyAdmin{
        icoState = State.halted;
    }
    
    // ICO 재개
    function resume() public onlyAdmin{
        icoState = State.running;
    }
    
    
    function changeDepositAddress(address payable newDeposit) public onlyAdmin{
        deposit = newDeposit;
    }
    
    
    function getCurrentState() public view returns(State){
        if(icoState == State.halted){
            return State.halted;
        }else if(block.timestamp < saleStart){
            return State.beforeStart;
        }else if(block.timestamp >= saleStart && block.timestamp <= saleEnd){
            return State.running;
        }else{
            return State.afterEnd;
        }
    }
 
 
    event Invest(address investor, uint value, uint tokens);
    
    
    
    function invest() payable public returns(bool){ 
        icoState = getCurrentState();
        require(icoState == State.running);
        require(msg.value >= minInvestment && msg.value <= maxInvestment);
        
        raisedAmount += msg.value;
        require(raisedAmount <= hardCap);
        
        uint tokens = msg.value / tokenPrice;
 
        // founder의 토큰 잔액에서 Inverstor의 토큰 잔액에 토큰 추가
        balances[msg.sender] += tokens;
        balances[founder] -= tokens; 
        deposit.transfer(msg.value); // ICO로 전송된 값을 입금 주소로 전송 

        
        emit Invest(msg.sender, msg.value, tokens);
        
        return true;
    }
   
   
   // 누군가가 ETH를 컨트랙트의 주소로 보낼 때(직접 보낼 경우) 자동으로 호출된다.
   receive () payable external{
        invest();
    }
  
    
    // 안 팔린 토큰 소각
    function burn() public returns(bool){
        icoState = getCurrentState();
        require(icoState == State.afterEnd);
        balances[founder] = 0;
        return true;
        
    }
    
    
    function transfer(address to, uint tokens) public override returns (bool success){
        require(block.timestamp > tokenTradeStart); // 토큰은 tokenTradeStart 이후에만 양도 가능합니다. 
        
        // 기본 계약의 전달 함수 호출
        super.transfer(to, tokens);  // same as Cryptos.transfer(to, tokens);
        return true;
    }
    
    
    function transferFrom(address from, address to, uint tokens) public override returns (bool success){
        require(block.timestamp > tokenTradeStart); 
       
        Cryptos.transferFrom(from, to, tokens);
        return true;
     
    }
}