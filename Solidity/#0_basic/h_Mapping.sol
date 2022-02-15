//SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.5.0 <0.9.0;
contract Auction{
    
    // # 매핑 유형의 변수 선언
    // key는 address type이고 value는 uint type.
    mapping(address => uint) public bids;
    
    // # 매핑 변수 초기화
    // 여기서 key는 함수를 호출하는 계정의 주소.
    // value는 함수를 호출할 때 wei가 보낸 값
    function bid() payable public{
        bids[msg.sender] = msg.value;
    }
}