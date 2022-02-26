// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.1;

contract SendMoneyEx {
    uint public balancedReceived;
    uint public withdrawTimestamp;

        struct Payment {
        uint amount;
        uint timestamp;
    }

    struct Balance {
        uint totalBalance;
        uint numPayments;
        mapping(uint=>Payment) payments;
    }

    mapping(address => Balance) public balanceReceived;

    function sendMoney() public payable {
        balanceReceived[msg.sender].totalBalance += msg.value;
    }

    function receiveMoney() public payable {
        balancedReceived += msg.value;
        withdrawTimestamp = block.timestamp + 1 minutes;
    }

    function getBalance() public view returns(uint){
        return address(this).balance;
    }

    function withdrawMoney() public {
        address payable to = payable(msg.sender);
        to.transfer(getBalance());
    }

    function withdrawAllMoney(address payable _to, uint _sendValue) public {
        require(balanceReceived[msg.sender].totalBalance >= _sendValue,"gaga");
        balanceReceived[msg.sender].totalBalance -= _sendValue;
        _to.transfer(_sendValue);
    }




    function withdrawMoneyTo(address payable _addr) public {
        require(withdrawTimestamp  < block.timestamp );
        _addr.transfer(getBalance());
       
    }
}