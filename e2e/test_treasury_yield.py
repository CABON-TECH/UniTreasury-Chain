import requests
import time
BASE_URL = "http://localhost:8081"
resp = requests.post(f"{BASE_URL}/api/v1/auth/login", json={
    "username": "admin",
    "password": "password"
})
if resp.status_code not in [200, 201]:
    print("Login failed:", resp.status_code, resp.text)
    exit(1)
token = resp.json().get("token")
headers = {"Authorization": f"Bearer {token}"}
resp = requests.post(f"{BASE_URL}/api/v1/treasury/proposals", json={
    "amount": 1000,
    "recipient": "0x111122223333444455556666777788889999aAaa",
    "purpose": "Test Yield Unstaking"
}, headers=headers)
if resp.status_code not in [200, 201]:
    print("Propose failed:", resp.status_code, resp.text)
    exit(1)
proposal_id = resp.json().get("id")
print("Proposed:", proposal_id)
time.sleep(3)
resp = requests.post(f"{BASE_URL}/api/v1/treasury/proposals/{proposal_id}/approve", headers=headers)
if resp.status_code not in [200, 201]:
    print("Approve failed:", resp.status_code, resp.text)
    exit(1)
print("Approved")
time.sleep(3)
resp = requests.post(f"{BASE_URL}/api/v1/treasury/proposals/{proposal_id}/execute", headers=headers)
if resp.status_code not in [200, 201]:
    print("Execute failed:", resp.status_code, resp.text)
    exit(1)
print("Executed Successfully! Yield unstaking works.")
