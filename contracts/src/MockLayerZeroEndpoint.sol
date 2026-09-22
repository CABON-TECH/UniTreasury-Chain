pragma solidity ^0.8.20;
contract MockLayerZeroEndpoint {
    event MessageSent(uint16 dstChainId, bytes dstAddress, uint256 amount);
    function send(uint16 _dstChainId, bytes calldata _dstAddress, bytes calldata _payload, address payable _refundAddress, address _zroPaymentAddress, bytes calldata _adapterParams) external payable {
        uint256 amount = abi.decode(_payload, (uint256));
        emit MessageSent(_dstChainId, _dstAddress, amount);
    }
}
