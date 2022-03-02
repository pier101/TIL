import React, { useState, useEffect } from "react";
import ItemManager from "./contracts/ItemManager.json";
import Item from "./contracts/Item.json";
import getWeb3 from "./getWeb3";
import "./App.css";

const App = () => {

  const [isLoading, setIsLoading] = useState(false)
  const [accounts, setAccounts] = useState()
  const [itemManagerContract, setItemManagerContract] = useState()
  const [itemContract, setItemContract] = useState()
  const [itemName, setItemName] = useState();
  const [cost, setCost] = useState();
  

  const init = async () => {
    try {

      const web3 = await getWeb3();
      const account = await web3.eth.getAccounts(); 
      const networkId = await web3.eth.net.getId();
      // 계약 인스턴스를 가져옵니다.
      const _ItemManagerContract = await new web3.eth.Contract(
        ItemManager.abi,
        ItemManager.networks[networkId] && ItemManager.networks[networkId].address,
      );
      const _ItemContract = await new web3.eth.Contract(
        Item.abi,
        Item.networks[networkId] && Item.networks[networkId].address,
      );

      const listenToPaymentEvent = () => {
        console.log("이벤트 밖")
          itemManagerContract.events.SupplyChainStep().on("data",async function(e) {
          console.log("이벤트 안")
          if(e.returnValues._step == 1) {
            let item = await itemManagerContract.methods.items(e.returnValues._itemIndex).call();
            console.log(item);
            alert("Item " + item.itemName + " was paid, deliver it now!");
          };
          console.log(e);
        });
      }
      setAccounts(account);
      setIsLoading(true);
      setItemManagerContract(_ItemManagerContract);
      setItemContract(_ItemContract);
      listenToPaymentEvent();
      console.log("유즈이펙트");

    } catch (error) {
      alert(`Failed to load web3, accounts, or contract. Check console for details.`);
      console.error(error);
    }
  }

  useEffect(async () => {
    init();
  }, [])

  const listenToPaymentEvent = () => {
    console.log("이벤트 밖")
      itemManagerContract.events.SupplyChainStep().on("data",async function(e) {
      console.log("이벤트 안")
      if(e.returnValues._step == 1) {
        let item = await itemManagerContract.methods.items(e.returnValues._itemIndex).call();
        console.log(item);
        alert("Item " + item.itemName + " was paid, deliver it now!");
      };
      console.log(e);
    });
  }

  const handleSubmit = async () => {

    console.log(itemName, cost, itemManagerContract);
    let result = await itemManagerContract.methods.createItem(itemName, cost).send({ from: accounts[0] });
    console.log(result);
    alert("Send "+cost+" Wei to "+result.events.SupplyChainStep.returnValues._address);
    console.log(itemManagerContract)
    // await itemManagerContract.once.('SupplyChainStep',(e)=>console.log(e));
  };

  const handleInputChange = (event) => {
    const target = event.target;
    const value = target.type === 'checkbox' ? target.checked : target.value;
    const name = target.name;

    setItemName(name)
    setCost(value);
  }


  if (!isLoading) {
    return <div>Loading Web3, accounts, and contract...</div>;
  }
  return (
    <div className="App">
      <h1>Simply Payment/Supply Chain Example!</h1>
      <h2>Items</h2>

      <h2>Add Element</h2>
      Cost: <input type="text" name="cost" value={cost} onChange={handleInputChange} />
      Item Name: <input type="text" name="itemName" value={itemName} onChange={handleInputChange} />
      <button type="button" onClick={handleSubmit}>Create new Item</button>
    </div>
  );

}

export default App
