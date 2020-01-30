const CASINO_ADDRESS = "0x8b173E32DbE8DA7746ef1a61FDb214b5Fa6F7DE4";
const casinoContract = new web3.eth.Contract(getBankBalanceABI, CASINO_ADDRESS);

web3 = new Web3(Web3.givenProvider); // provider from metamask
web3.extend({ // web3.eth.requestAccounts() isn't available (yet)
    methods: [{
        name: 'eth_requestAccounts',
        call: 'eth_requestAccounts',
        params: 0
    }]
});

async function foo() {
    log(`Version of web3.js: ${web3.version}`);
    var result=await web3.eth_requestAccounts().catch(x=>log(x.message));
    log(`Value from eth_requestAccounts: ${JSON.stringify(result)}`);
    var acts=await web3.eth.getAccounts().catch(log);
    log(`Here are the accounts: ${JSON.stringify(acts)}`);

    ethereum.on('accountsChanged', newacts => { 
        acts=newacts;
        log(`We have new accounts: ${JSON.stringify(acts)}`);
    })
}

const getBankBalanceABI = [{
    "constant": true,
    "inputs": [],
    "name": "getBankBalance",
    "outputs": [
        {
        "internalType": "uint256",
        "name": "ret",
        "type": "uint256"
        }
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
}];

async function f() {
    var result = await casinoContract.methods.getBankBalance().call();
    log(`There is ${result} Wei in the bank`);
}

function log(logstr) {
    document.getElementById("log").innerHTML +=logstr+"\n";
}

foo();