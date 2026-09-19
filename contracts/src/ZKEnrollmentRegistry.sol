// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./MockZKVerifier.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

/// @title ZKEnrollmentRegistry
/// @notice Manages a Merkle tree root of enrolled student hashes.
///         Students can anonymously prove enrollment by submitting a ZK proof.
///         A verified proof mints a single-use "verified" flag per nullifier,
///         which the ScholarshipEscrowContract can check before releasing funds.
contract ZKEnrollmentRegistry is AccessControl {
    bytes32 public constant ADMIN_ROLE = DEFAULT_ADMIN_ROLE;
    bytes32 public constant REGISTRAR_ROLE = keccak256("REGISTRAR_ROLE");

    MockZKVerifier public verifier;

    /// @notice Current Merkle root of enrolled student commitment hashes.
    bytes32 public enrollmentMerkleRoot;

    /// @notice Mapping: nullifierHash => verified student address
    ///         Once verified, a nullifier grants a specific address one-time proof of enrollment.
    mapping(bytes32 => address) public verifiedNullifier;

    event MerkleRootUpdated(bytes32 indexed newRoot, address indexed updatedBy);
    event EnrollmentProofVerified(bytes32 indexed nullifierHash, address indexed student);

    error ZKRegistry__InvalidProof();
    error ZKRegistry__NullifierAlreadyUsed();

    constructor(address _admin, address _verifier) {
        _grantRole(DEFAULT_ADMIN_ROLE, _admin);
        _grantRole(REGISTRAR_ROLE, _admin);
        verifier = MockZKVerifier(_verifier);
    }

    /// @notice Registrar updates the Merkle root (e.g., when new students enroll each semester).
    function updateMerkleRoot(bytes32 newRoot) external onlyRole(REGISTRAR_ROLE) {
        enrollmentMerkleRoot = newRoot;
        emit MerkleRootUpdated(newRoot, msg.sender);
    }

    /// @notice Student submits a ZK proof to prove enrollment without revealing their identity.
    /// @param proof         The ZK-SNARK proof bytes.
    /// @param nullifierHash Unique per (student, session). Prevents reuse.
    /// @param claimedRoot   The Merkle root the student is proving against.
    function proveEnrollment(
        bytes calldata proof,
        bytes32 nullifierHash,
        bytes32 claimedRoot
    ) external {
        if (verifiedNullifier[nullifierHash] != address(0)) revert ZKRegistry__NullifierAlreadyUsed();

        bool valid = verifier.verifyProof(proof, nullifierHash, claimedRoot, enrollmentMerkleRoot);
        if (!valid) revert ZKRegistry__InvalidProof();

        verifiedNullifier[nullifierHash] = msg.sender;
        emit EnrollmentProofVerified(nullifierHash, msg.sender);
    }

    /// @notice Check if a given nullifier has been used to verify enrollment.
    /// @param nullifierHash  The nullifier to check.
    /// @return student       The address that verified using this nullifier (or address(0) if not verified).
    function isEnrollmentVerified(bytes32 nullifierHash) external view returns (address student) {
        return verifiedNullifier[nullifierHash];
    }
}
