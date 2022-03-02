// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RabdomPick {

    event pickResult(address indexed _pickAddr, uint indexed _pickCount, uint _pickItem);
    event bonusPickResult(address indexed _pickAddr, uint _pickItem);

    address private owner;

    constructor  (){
        owner = msg.sender;
    }

    receive()external payable {}
    
    uint[10] public itemList = [1,2,3,4,5,6,7,8,9,10];
    uint[] public itemlistarr;
    uint private turn;
    mapping(address=>uint[]) public itemsPerAddress; //계정당 아이템
    mapping(address=>uint) private pickCountPerAddress; //계정당 뽑기 횟수
    mapping(address=>bool) private bonusTurn;



    function randomPick () public payable {
        require(msg.value == 2*10**16);
        uint random = itemList[randomItem()];
        itemsPerAddress[msg.sender].push(itemList[randomItem()]);

        emit pickResult(msg.sender, pickCountPerAddress[msg.sender], random);
        turn++;
        if(turn % 5 == 0){
            bonusTurn[msg.sender] = true;
        }
    }

    function bonusPick() public {
        require(bonusTurn[msg.sender]==true,"you are not bonus turn");
        uint bonus = itemList[randomItem()];
        itemsPerAddress[msg.sender].push(itemList[randomItem()]);
        emit bonusPickResult(msg.sender,bonus);
        bonusTurn[msg.sender] = false;
    }

    //난수 생성
    function randomItem() public view returns(uint){
        return  uint(keccak256(abi.encodePacked(block.difficulty, block.timestamp, pickCountPerAddress[msg.sender]))) % 10 + 1;
    }

    function getItemsPerAddress() public view returns(uint[] memory){
        return itemsPerAddress[msg.sender];
    }

}
