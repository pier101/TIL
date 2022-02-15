//SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.5.0 <0.9.0;

contract Deposit{

    address public owner;
    
    constructor(){
        owner = msg.sender;    
    }

    receive() external payable {}

    fallback() external payable {}

    function getBalance() public view returns(uint) {
        return address(this).balance;
    }

    function transferEther(address payable _to, uint _amount) public returns(bool) {
        require(owner == msg.sender, "you are not owner");

        if (_amount <= getBalance()) {
            _to.transfer(_amount);
            return true;
        } else {
            return false;
        }
    }
}
