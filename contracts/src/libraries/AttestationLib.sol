// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title AttestationLib
/// @notice EIP-712 attestation verification helper for scholarship tranche releases.
///         The attestor (off-chain backend service) signs a struct committing to a specific
///         (fundId, studentHash, trancheIndex, recipient) tuple. The contract verifies the
///         signature and records the nonce to prevent replay.
///
/// @dev Trust model: this is a SINGLE trusted attestor — a centralization trade-off explicitly
///      accepted in the system's threat model. See: docs/threat-model.md.
library AttestationLib {
    // ── EIP-712 Type Hashes ───────────────────────────────────────────────────────

    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)");

    bytes32 public constant ATTESTATION_TYPEHASH = keccak256(
        "TrancheAttestation(uint256 fundId,bytes32 studentHash,uint256 trancheIndex,address recipient,uint256 nonce)"
    );

    // ── Structs ───────────────────────────────────────────────────────────────────

    struct TrancheAttestation {
        uint256 fundId;
        bytes32 studentHash;
        uint256 trancheIndex;
        address recipient;
        uint256 nonce; // keccak256(fundId, studentHash, trancheIndex) cast to uint256
    }

    // ── Functions ─────────────────────────────────────────────────────────────────

    /// @notice Compute the EIP-712 domain separator for a given contract
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

    /// @notice Hash the attestation struct (EIP-712 structured data hash)
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

    /// @notice Recover signer from an EIP-712 signature
    /// @param att       The attestation struct
    /// @param domSep    Domain separator (computed per-contract)
    /// @param signature 65-byte (r, s, v) signature
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

    /// @notice Compute the canonical nonce for a (fundId, studentHash, trancheIndex) tuple
    function computeNonce(
        uint256 fundId,
        bytes32 studentHash,
        uint256 trancheIndex
    ) internal pure returns (uint256) {
        return uint256(keccak256(abi.encode(fundId, studentHash, trancheIndex)));
    }
}
