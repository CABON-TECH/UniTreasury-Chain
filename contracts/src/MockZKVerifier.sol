pragma solidity ^0.8.22;
contract MockZKVerifier {
    mapping(bytes32 => bool) public nullifierUsed;
    event ProofVerified(bytes32 indexed nullifierHash, address indexed prover);
    error ZKVerifier__InvalidProof();
    error ZKVerifier__NullifierAlreadyUsed(bytes32 nullifierHash);
    function verifyProof(
        bytes calldata proof,
        bytes32 nullifierHash,
        bytes32 merkleRoot,
        bytes32 expectedRoot
    ) external returns (bool valid) {
        if (merkleRoot != expectedRoot) revert ZKVerifier__InvalidProof();
        if (proof.length == 0 || proof[0] == 0x00) revert ZKVerifier__InvalidProof();
        if (nullifierUsed[nullifierHash]) revert ZKVerifier__NullifierAlreadyUsed(nullifierHash);
        nullifierUsed[nullifierHash] = true;
        emit ProofVerified(nullifierHash, msg.sender);
        return true;
    }
}
