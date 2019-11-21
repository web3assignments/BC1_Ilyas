pragma solidity >=0.4.0 <0.7.0;

/**
 * A storage of Product's.
 */
contract ProductRepository {
    /**
     * A data structure of a Product entry within the store.
     */
    struct Product {
        uint id;
        string name;
        uint price;
    }

    /**
     * The current value of the store, which is the aggregation
     * of every Product's price.
     */
    uint storeValue;

    /**
     * The product storage.
     */
    mapping (uint => Product) products;

    /**
     * Inserts a Product with the given details into the ledger. The given
     * price of the Product will increase the total store value.
     */
    function insert(uint id, string memory name, uint price) public {
        products[id] = Product(id, name, price);
        storeValue += price;
    }

    /**
     * Fetches the current store value.
     */
    function getStoreValue() public view returns (uint) {
        return storeValue;
    }

    /**
     * Removes a Product entry from the ledger. This operation may reduce
     * the total value of the store.
     */
    function remove(uint id) public {
        storeValue -= products[id].price;
        delete products[id];
    }
}