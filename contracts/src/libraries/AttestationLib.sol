pragma solidity ^0.8.20;
library AttestationLib {
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)");
    bytes32 public constant ATTESTATION_TYPEHASH = keccak256(
        "TrancheAttestation(uint256 fundId,bytes32 studentHash,uint256 trancheIndex,address recipient,uint256 nonce)"
    );
    struct TrancheAttestation {
        uint256 fundId;
        bytes32 studentHash;
        uint256 trancheIndex;
        address recipient;
        uint256 nonce; 
    }
    function domainSeparator(address contractAddress) internal view returns (bytes32) {
        return keccak256(
            abi.encode(
                DOMAIN_TYPEHASH,
                keccak256("UniTreasury ScholarshipEscrow"),
                keccak256("1"),
                block.chainid,
                contractAddress
            )
        );
    }
    function hashAttestation(TrancheAttestation memory att) internal pure returns (bytes32) {
        return keccak256(
            abi.encode(
                ATTESTATION_TYPEHASH,
                att.fundId,
                att.studentHash,
                att.trancheIndex,
                att.recipient,
                att.nonce
            )
        );
    }
    function recoverSigner(
        TrancheAttestation memory att,
        bytes32 domSep,
        bytes calldata signature
    ) internal pure returns (address) {
        require(signature.length == 65, "AttestationLib: invalid sig length");
        bytes32 digest = keccak256(abi.encodePacked("\x19\x01", domSep, hashAttestation(att)));
        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            r := calldataload(signature.offset)
            s := calldataload(add(signature.offset, 32))
            v := byte(0, calldataload(add(signature.offset, 64)))
        }
        if (v < 27) v += 27;
        return ecrecover(digest, v, r, s);
    }
    function computeNonce(
        uint256 fundId,
        bytes32 studentHash,
        uint256 trancheIndex
    ) internal pure returns (uint256) {
        return uint256(keccak256(abi.encode(fundId, studentHash, trancheIndex)));
    }
}
