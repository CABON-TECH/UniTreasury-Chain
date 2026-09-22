import requests
import json

BASE_URL = "http://localhost:8081"
resp = requests.post(f"{BASE_URL}/api/v1/auth/login", json={"username": "admin", "password": "password"})
token = resp.json()["token"]
headers = {"Authorization": f"Bearer {token}", "Content-Type": "application/json"}

resp = requests.post(f"{BASE_URL}/api/v1/students/student2/claim-l2", headers=headers, json={
    "fund_id": 1,
    "tranche_index": 0,
    "recipient": "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
    "dst_chain_id": 42161
})
print("Claim response:", resp.status_code, resp.text)
