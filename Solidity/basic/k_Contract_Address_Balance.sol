//SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.6.0 <0.9.0;

contract Deposit{
    
    address public owner;
    
    constructor(){
        owner = msg.sender;    
    }

    // Receive() 또는 fallback()은 contract가 ETH를 수신하는 데 필수.
    // ETH를 컨트랙트의 주소로 전송

    // ETH를 컨트랙트 주소로 보낼 때 실행되는 receive() 함수 선언
    // Solidity 0.6에서 도입되었으며 계약은 수신 기능을 하나만 가질 수 있다.
    // function 키워드와 인수 없이 선언
    receive() external payable{
    }

    // msg.data가 비어 있지 않거나
    // 일치하는 다른 함수가 없을 때
    fallback() external payable{
    }

    // 계약 잔액 반환
    function getBalance() public view returns(uint){
        // 명시적으로 주소로 변환할 수 있는 현재 계약입니다.  
        return address(this).balance;
    }

    // 이 함수를 호출하는 trasaction의 계약에서도 ether를 보내고 받을 수 있다.
    function sendEther() public payable{
        uint x;
        x++;
    }
    // 컨트랙트에서 다른 주소(수신자)로 이더 전송
    function transferEther(address payable recipient, uint amount) public returns(bool){
        // 계약 소유자만 계약에서 다른 주소로 이더를 보낼 수 있는지 확인
        require(owner == msg.sender, "transaction failed"); //"트랜잭션 실패 당신은 contract소유자가 아닙니다!"
        
        if(amount <= getBalance()){
            // 계약서에서 받는 사람 주소로 wei 금액 전송
             // 이 함수를 호출할 수 있는 모든 사람은 계약 자금에 액세스할 수 있다. => 위에 require문으로 권한제한 둠으로 보안 강화시킴
            recipient.transfer(amount);
            return true;
        } else{
            return false;
        }
    }
}
