//SPDX-License-Identifier: GPL-3.0
 
pragma solidity >=0.5.0 <0.9.0;
 
contract A{
    int public x = 10;
    int y = 20; // 선언 안 할 시 기본값 = internal
    
    // public function
    function get_y() public view returns(int){
    return y;
    }
    
    
    // private function, 이 계약 내에서만 호출할 수 있음(파생된 계약 내에서 아님)
    function f1() private view returns(int){
    return x;
    }
    
    function f2() public view returns(int){
    int a;
    a = f1(); // f1을 호출할 수 있음 
    return a;
    }
    
    // 이 계약과 파생 계약 내에서만 호출할 수 있음.
    function f3() internal view returns(int){
    return x;
    }
    
    // 외부 계약 및 애플리케이션에서만 호출할 수 있음(공개보다 효율적임)
    function f4() external view returns(int){
    return x;
    }
    
    
    function f5() public view returns(int){
    int b;
    // b = f4(); // ERROR => f4() is external
    b = f3(); // OK => f3() is internal
    b = f1(); // OK => f1() is private
    return b;
    }
}

// B는 A에서 파생됩니다.
contract B is A{
    int public xx = f3(); // internal 함수를 호출할 수 있음  
    // int public yy = f1(); // ERROR => f1()은 private이며 파생된 계약에서 호출할 수 없음.
}

contract C{
    A public contract_a = new A(); // C는 A를 배포.  
    int public xx = contract_a.f4(); // OK => f4() is external
    // int public y = contract_a.f1(); // ERROR => f1() is private
    // int public yy = contract_a.f3(); // ERROR => f3() is internal
}
