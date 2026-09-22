pragma solidity ^0.8.22;
import "./MockZKVerifier.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
contract ZKEnrollmentRegistry is AccessControl {
    bytes32 public constant ADMIN_ROLE = DEFAULT_ADMIN_ROLE;
    bytes32 public constant REGISTRAR_ROLE = keccak256("REGISTRAR_ROLE");
    MockZKVerifier public verifier;
    bytes32 public enrollmentMerkleRoot;
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
    function updateMerkleRoot(bytes32 newRoot) external onlyRole(REGISTRAR_ROLE) {
        enrollmentMerkleRoot = newRoot;
        emit MerkleRootUpdated(newRoot, msg.sender);
    }
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
    function isEnrollmentVerified(bytes32 nullifierHash) external view returns (address student) {
        return verifiedNullifier[nullifierHash];
    }
}
