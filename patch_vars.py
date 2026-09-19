with open("contracts/src/TreasuryContract.sol", "r") as f:
    text = f.read()

text = text.replace("IERC20 public immutable usdcToken;", "IERC20 public usdcToken;\n    address public aavePool;\n    address public aUsdcToken;")

with open("contracts/src/TreasuryContract.sol", "w") as f:
    f.write(text)
