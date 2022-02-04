// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.5.0 <0.9.0;

contract Property{
    
    // # 상태변수 ex) price,location
    int public price=100; // 기본값 = 0
    string constant public location = "London";
    // price=100;  solidity에선 상태변수의 기본값 변경이 허용되지 않는다. 

    // # 지역변수 ex) x
    function f1() public pure returns(int){
        int x = 5;
        x = x * 2;

        // string memory s1 = "abc";
        return x;
    }
}