pragma solidity ^0.8.20;
import {ScholarshipEscrowContract} from "./ScholarshipEscrowContract.sol";
contract ScholarshipEscrowContractV2 is ScholarshipEscrowContract {
    string public version;
    function setVersion(string memory _version) external onlyRole(ADMIN_ROLE) {
        version = _version;
    }
}
