//SPDX-License-Identifier: GPL-3.0
 
pragma solidity >=0.5.0 <0.9.0;

contract GlobalVariables{

    // 타임스탬프로서의 현재 시간(1970년 1월 1일부터의 초)
    uint public this_moment = block.timestamp; // `now`는 더 이상 사용되지 않으며 block.timestamp의 별칭임)  
 
    // 현재 블록 번호
    uint public block_number = block.number;
    
    // 현재 블록 난이도
    uint public difficulty = block.difficulty;
    
    // 현재 블록 가스 한도
    uint public gaslimit = block.gaslimit;

    address public owner;
    uint public sentValue;

    constructor(){
        // msg.sender는 계약과 상호 작용하는 주소(이 경우 배포).
        owner = msg.sender;
    }

    function changeOwner() public{
        owner = msg.sender;
    }

    function  sendEther() public payable{ // 트랜잭션과 함께 ETH를 받으려면 지불할 수 있어야 함
    // msg.value = 이 트랜잭션에서 보낸 wei의 값(함수 호출 시).
        sentValue = msg.value;
    }

    // 계약 잔액 반환
    function getBalance() public view returns(uint){
        return address(this).balance;
    }

    function howMuchGas() public view returns(uint) {
        uint start = gasleft();
        uint j = 1;
        for(uint i = 1; i < 20; i++) {
            j *= i;
        }

        uint end = gasleft();
        return start - end;
    }

}