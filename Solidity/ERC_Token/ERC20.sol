//SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.6.0 <0.9.0;


/*
    - ERC20은 지갑과 같은 응용 프로그램에서 사용 하는 표준 ""인터페이스"".
    - 겉으론 Ether를 저장하는 동일한 지갑을 사용하여 토큰을 구매, 판매 또는 전송하지만 
      실제로는 contract와 상호 작용한다.
    - 이더리움 네트워크가 ETH를 소유한 사람을 추적하는 것과 같은 방식으로 Contract는 토큰 소유권을 추적한다.
    - ERC20을 완전히 준수하는 토큰과 부분적으로 준수하는 토큰이 있다.
    - 완벽하게 호환되는 ERC20 토큰은 6개의 기능과 2개의 이벤트를 구현해야 한다(아래 사용함).
*/

//완벽하게 호환되는 ERC20 토큰의 예
interface ERC20Interface {
    function totalSupply() external view returns (uint);
    function balanceOf(address tokenOwner) external view returns(uint balance);
    function transfer(address to, uint tokens) external returns(bool success);

    function allowance(address tokenOwner, address spender) external view returns(uint remaining);
    function approve(address spender, uint tokens) external returns(bool success);
    function transferFrom(address from, address to, uint tokens) external returns(bool success);

    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}

contract Cryptos is ERC20Interface{
    string public name = "Wook";
    string public symbol = "KDW";
    uint public decimals = 0; //18
    uint public override totalSupply;

    address public founder;
    // 주소별 토큰수 
    mapping(address => uint) public balances;
    
    mapping(address => mapping(address => uint)) allowed;

    constructor(){
        totalSupply = 1000000;
        founder = msg.sender;
        balances[founder] = totalSupply;
    }

    function balanceOf(address tokenOwner) public view override returns(uint balance){
        return balances[tokenOwner];
    }
    function transfer(address to, uint tokens) public override returns(bool success){
        require(balances[msg.sender] >= tokens);

        balances[to] += tokens;
        balances[msg.sender] -= tokens;
        emit Transfer(msg.sender, to, tokens);
        
        return true;
    }

    // 인출 가능량 조회
    function allowance(address tokenOwner, address spender) view public override returns(uint){
        return allowed[tokenOwner][spender];
    }
    
    // 인출 허용량 설정
    function approve(address spender, uint tokens) public override returns (bool success){
        require(balances[msg.sender] >= tokens);
        require(tokens > 0);
        
        allowed[msg.sender][spender] = tokens;
        
        emit Approval(msg.sender, spender, tokens);
        return true;
    }
    
    // 한도 내 인출하기
    function transferFrom(address from, address to, uint tokens) public override returns (bool success){
         require(allowed[from][to] >= tokens);
         require(balances[from] >= tokens);
         
         balances[from] -= tokens;
         balances[to] += tokens;
         allowed[from][to] -= tokens;
         
         emit Transfer(from, to, tokens);
         
         return true;
     }
}
 
