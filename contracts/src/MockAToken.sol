pragma solidity ^0.8.20;
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
contract MockAToken is ERC20, Ownable {
    address public underlyingAsset;
    constructor(
        string memory name,
        string memory symbol,
        address _underlyingAsset
    ) ERC20(name, symbol) Ownable(msg.sender) {
        underlyingAsset = _underlyingAsset;
    }
    function mint(address account, uint256 amount) external onlyOwner {
        _mint(account, amount);
    }
    function burn(address account, uint256 amount) external onlyOwner {
        _burn(account, amount);
    }
}
