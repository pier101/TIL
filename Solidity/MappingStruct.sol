//SPDX-License-Identifier: MIT

pragma solidity ^0.8.4;

contract MappingsStructExample {

    struct Payment {
        uint amount;
        uint timestamp;
    }

    struct Balance {
        uint totalBalance;
        uint numPayments;
        mapping(uint => Payment) payments;
    }

    mapping(address => Balance) public balanceReceived;


    function getBalance() public view returns(uint) {
        return address(this).balance;
    }

    function sendMoney() public payable {
        
        balanceReceived[msg.sender].totalBalance += msg.value; // 토탈 밸런스에 보낸 양만큼 추가
        // 지불 구조체 생성
        Payment memory payment = Payment(msg.value, block.timestamp); 
        // 밸런스 구조체 안에 생성한 지불 구조체 할당
        balanceReceived[msg.sender].payments[balanceReceived[msg.sender].numPayments] = payment; 
        balanceReceived[msg.sender].numPayments++;
    }

    // 부분 인출
    function withdrawMoney(address payable _to, uint _amount) public {
        require(_amount <= balanceReceived[msg.sender].totalBalance, "not enough funds");
        balanceReceived[msg.sender].totalBalance -= _amount;
        _to.transfer(_amount);
    }

    // 모두 인출
    function withdrawAllMoney(address payable _to) public {
        uint balanceToSend = balanceReceived[msg.sender].totalBalance;
        balanceReceived[msg.sender].totalBalance = 0;
        _to.transfer(balanceToSend);
    }
}