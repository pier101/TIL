//SPDX-License-Identifier: MIT

pragma solidity 0.8.1;

import "./Allowance.sol";

contract SharedWallet is Allowance {
    
    // 이벤트 추가 (누군가가 자금을 입금하거나 인출할 때 공유 지갑에 이벤트가 나타날 목적으로 사용)
    event MoneySent(address indexed _beneficiary, uint _amount);
    event MoneyReceived(address indexed _from, uint _amount);

    // @ 전송 기능 (관리자/접근허용자 권한)
    function withdrawMoney(address payable _to, uint _amount) public ownerOrAllowed(_amount) {
        require(_amount <= address(this).balance, "Contract doesn't own enough money");
        
        if(!isOwner()){ //관리자가 아닌경우 
            reduceAllowance(msg.sender, _amount);
        }
        emit MoneySent(_to,_amount); //이벤트 발동
        _to.transfer(_amount);
    }

    receive() external payable {
        emit MoneyReceived(msg.sender, msg.value); // 이벤트 발동
    }

    // Onable.sol의 renounceOwnership() : 소유권 포기 기능이 있다.
    // 해당 기능을 동작하지 못하게 하기 위해 override하고 재설정하여 컨트랙트 기능을 없애지 못하도록 만듦. 
    function renounceOwnership() public override onlyOwner {
        revert("can't renounceOwnership here"); //not possible with this smart contract
    }

}