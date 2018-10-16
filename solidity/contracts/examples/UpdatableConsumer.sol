pragma solidity ^0.4.24;

import "./Consumer.sol";

contract UpdatableConsumer is Consumer {

  constructor(bytes32 _specId, address _ens) public {
    specId = _specId;
    newChainlinkWithENS(_ens);
  }

  function updateOracle() public {
    updateOracleWithENS();
  }

  function getChainlinkToken() public view returns (address) {
    return chainlinkToken();
  }
  
  function getOracle() public view returns (address) {
    return oracleAddress();
  }

}
