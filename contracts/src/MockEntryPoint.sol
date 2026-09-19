// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

struct UserOperation {
    address sender;
    uint256 nonce;
    bytes initCode;
    bytes callData;
    uint256 callGasLimit;
    uint256 verificationGasLimit;
    uint256 preVerificationGas;
    uint256 maxFeePerGas;
    uint256 maxPriorityFeePerGas;
    bytes paymasterAndData;
    bytes signature;
}

contract MockEntryPoint {
    // In a real EntryPoint, this validates the paymaster and signature, then calls the sender contract
    // We will simulate handling a UserOp by just executing the callData on the target contract directly for our mock test.
    
    function handleOps(UserOperation[] calldata ops, address payable beneficiary) external {
        for (uint256 i = 0; i < ops.length; i++) {
            UserOperation calldata op = ops[i];
            
            // Assuming the sender in our mock is the Escrow contract, and callData is the function to call
            // Wait, normally sender is the Account contract which then calls the Escrow.
            // For this mock to work as a "Paymaster" simulation, we'll just extract the target from the op
            // and execute it. We'll assume sender is the target for the mock.
            (bool success, ) = op.sender.call{gas: op.callGasLimit}(op.callData);
            require(success, "MockEntryPoint: op failed");
        }
    }
}
