import json

with open("contracts/broadcast/DeployAll.s.sol/31337/run-latest.json") as f:
    data = json.load(f)

for log in data["logs"]:
    print(log)
