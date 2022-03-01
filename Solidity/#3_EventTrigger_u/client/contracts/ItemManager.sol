//SPDX-License-Identifier: MIT


pragma solidity ^0.8.11;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./Item.sol";

contract ItemManager is Ownable{
    
    // 공급 과정 상태
    enum SupplyChainState{Created, Paid, Delivered}

    // 아이템 구조체
    struct S_Item {
        Item item;
        // uint itemPrice; 구조변경 아이템네임 지우고 > Item contract추가
        string itemName;
        SupplyChainState state;
    }

    // 아이템 목록
    mapping(uint=>S_Item) public items;
    uint itemIndex;

    event SupplyChainStep(uint _itemIndex, uint _step, address _address);


    // @ 아이템 생성
    function createItem(string memory _itemName, uint _itemPrice) public onlyOwner {
        Item _item = new Item(this,  _itemPrice, itemIndex); // Item contract 활용 새 아이템 생성
        items[itemIndex].item = _item;
        items[itemIndex].itemName = _itemName;
        items[itemIndex].state = SupplyChainState.Created;
        emit SupplyChainStep(itemIndex, uint(items[itemIndex].state), address(_item));
        itemIndex++;
    }

    // 상태 변경( > 결제)
    function triggerPayment(uint _index) public payable {
        Item _item = items[_index].item;
        require(_item.priceInWei() <= msg.value, "Not fully paid");
        require(items[_index].state == SupplyChainState.Created, "Item is further in the supply chain");
        items[_index].state = SupplyChainState.Paid;
        emit SupplyChainStep(_index, uint(items[_index].state), address(_item));
   }

    // 상태 변경( > 배송)
    function triggerDelivery(uint _index) public onlyOwner {
        require(items[_index].state == SupplyChainState.Paid, "Item is further in the supply chain");
        items[_index].state = SupplyChainState.Delivered;
        emit SupplyChainStep(_index, uint(items[_index].state), address(items[_index].item));
    }
}
