import requests
import json

BASE_URL = "http://localhost:8081"
resp = requests.post(f"{BASE_URL}/api/v1/auth/login", json={"username": "admin", "password": "password"})
token = resp.json()["token"]
headers = {"Authorization": f"Bearer {token}", "Content-Type": "application/json"}

resp = requests.get(f"{BASE_URL}/api/v1/students", headers=headers)
students = resp.json()["data"]
student = next((s for s in students if s["student_id"] == "student2"), None)
sid = student["id"]

resp = requests.post(f"{BASE_URL}/api/v1/students/student2/claim-l2", headers=headers, json={
    "fund_id": 1,
    "tranche_index": 0,
    "recipient": "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC",
    "dst_chain_id": 42161
})
print("Claim response:", resp.status_code, resp.text)
