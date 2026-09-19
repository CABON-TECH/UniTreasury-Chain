with open("contracts/src/TreasuryContract.sol", "r") as f:
    text = f.read()

text = text.replace("address _usdcToken\n    ) {", "address _usdcToken,\n        address _aavePool,\n        address _aUsdcToken\n    ) {")

with open("contracts/src/TreasuryContract.sol", "w") as f:
    f.write(text)
