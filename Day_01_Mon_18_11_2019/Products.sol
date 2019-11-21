pragma solidity >=0.4.0 <0.7.0;

contract ProductRepository {
  	struct Product {
      uint id;
      string name;
      uint price;
    }

  	uint storeValue;
    mapping (uint => Product) products;

    function insert(uint id, string memory name, uint price) public {
      products[id] = Product(id, name, price);
      storeValue += price;
    }

  	function getStoreValue() public view returns (uint) {
    	return storeValue;
    }

  	function remove(uint id) public {
    	storeValue -= products[id].price;
      delete products[id];
    }
}