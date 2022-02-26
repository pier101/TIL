//SPDX-License-Identifier: MIT

pragma solidity 0.8.1;

import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/access/Ownable.sol";
import "https://github.com/OpenZeppelin/openzeppelin-contracts/contracts/utils/math/SafeMath.sol";

contract Allowance is Ownable {

     using SafeMath for uint; // sol 8.0 버전이후 정수 유형 변수는 더 이상 오버플로할 수 없다
                              // 8 버전 이하에서만 사용 (여기서도 사용 안할꺼임)

    //이벤트 추가
    event AllowanceChange(address indexed _forWho, address indexed _byWho, uint _oldAmount, uint _newAmount);
    
    // 관리자 여부
    function isOwner() internal view returns(bool) {
        return owner() == msg.sender;
    }

    // 접근 가능자
    mapping(address => uint) public allowance;


     // @ 특정 주소에 특정 amount 접근 허용 (관리자권한)
    function setAllowance(address _who, uint _amount) public onlyOwner {
        emit AllowanceChange(_who, msg.sender, allowance[_who], _amount); //변화시키는 함수므로 이벤트 발생
        allowance[_who] = _amount;
    }

    // 관리자or 접근허용주소만 접근 가능 수정자
    modifier ownerOrAllowed(uint _amount){
        require(isOwner() || allowance[msg.sender] >= _amount, "You are not allowed!");
        _;
    }

    // 접근 허용 주소의 허용량 감소
    function reduceAllowance(address _who, uint _amount) internal ownerOrAllowed(_amount) {
        emit AllowanceChange(_who, msg.sender, allowance[_who], allowance[_who] - _amount); //변화시키는 함수므로 이벤트 발생
        allowance[_who] -= _amount;
    }
}