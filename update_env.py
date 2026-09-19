import json

with open("contracts/broadcast/DeployAll.s.sol/31337/run-latest.json") as f:
    data = json.load(f)

env_vars = {}
for tx in data.get("transactions", []):
    contract_name = tx.get("contractName")
    addr = tx.get("contractAddress")
    if contract_name == "MockUSDC": env_vars["MOCK_USDC_ADDRESS"] = addr
    elif contract_name == "TreasuryContract": env_vars["TREASURY_CONTRACT_ADDRESS"] = addr
    elif contract_name == "FeeRegistryContract": env_vars["FEE_REGISTRY_CONTRACT_ADDRESS"] = addr
    elif contract_name == "ERC1967Proxy": env_vars["SCHOLARSHIP_ESCROW_CONTRACT_ADDRESS"] = addr
    elif contract_name == "MockLayerZeroEndpoint": env_vars["MOCK_LZ_ENDPOINT_ADDRESS"] = addr
    elif contract_name == "MockEntryPoint": env_vars["ENTRYPOINT_ADDRESS"] = addr
    elif contract_name == "UniPaymaster": env_vars["PAYMASTER_ADDRESS"] = addr

with open("backend/.env", "r") as f:
    lines = f.readlines()

with open("backend/.env", "w") as f:
    for line in lines:
        if "=" in line:
            key = line.split("=")[0]
            if key in env_vars and env_vars[key] is not None:
                f.write(f"{key}={env_vars[key]}\n")
                continue
        f.write(line)
