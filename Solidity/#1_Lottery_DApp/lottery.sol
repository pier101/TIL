// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.5.0 <0.9.0;

contract Lotterty {

    address payable[] public players;
    address public manager;

    constructor(){
        manager = msg.sender;
    }
    receive() external payable {
        // 접미사가 없는 경우 wei로 간주 
        // ex)100000000000000000 (0.1 eth임)
        // # 0.1 ether만 전송가능하게 처리
        require(msg.value == 0.1 ether, "Only 0.1eth can be transferred!");
        //  payable(msg.sender) : 지불 가능한 주소로 변경
        players.push(payable(msg.sender));
    }

    function getBalance() public view returns(uint) {
        // # 관리자만 잔액 확인하게 처리
        require(msg.sender == manager, "you are not manager!");
        return address(this).balance;
    }
}

/*
    1. 상태 변수 정의
    2. 생성자 정의 
    3. contract가 ether 수신하는 기능 (receive())
    4. 계약의 잔액을 wei로 반환하는 함수 만들기 (getBalance())
    --- 중간 test ---
    5. 데이터 유효성 검사 & ethereum에서 다양한 조건 test

 */