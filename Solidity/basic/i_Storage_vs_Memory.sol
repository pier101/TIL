//SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.5.0 <0.9.0;

contract A {
    string[] public cities = ['seoul','suwon'];

    function f_memory() public{
        string[] memory s1 = cities;
        s1[0] = 'andong';
        //s1이라는 복사본에서만  값 변경됨
    }

    function f_storage() public{
        string[] storage s1 = cities;
        s1[0] = 'andong';
        //내부에서 S1만 변경했음에도 불구하고 함수가 상태 변수를 변경함.
    }
}