pragma solidity ^0.8.20;
interface ILayerZeroEndpoint {
    function send(uint16 _dstChainId, bytes calldata _dstAddress, bytes calldata _payload, address payable _refundAddress, address _zroPaymentAddress, bytes calldata _adapterParams) external payable;
}
