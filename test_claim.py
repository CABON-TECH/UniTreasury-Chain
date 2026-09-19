import requests
import json
import time

BASE_URL = "http://localhost:8081"

resp = requests.post(f"{BASE_URL}/api/v1/auth/login", json={
    "username": "admin",
    "password": "password"
})
if resp.status_code != 200:
    print("Login failed:", resp.text)
    exit(1)
token = resp.json()["token"]
headers = {"Authorization": f"Bearer {token}", "Content-Type": "application/json"}

print("Logged in as admin, creating demo fund...")
requests.post(f"{BASE_URL}/api/v1/demo/fund", json={"amount": 100000, "role": "admin"})

print("Creating student...")
resp = requests.post(f"{BASE_URL}/api/v1/students", headers=headers, json={
    "student_id": "student2", 
    "name": "Bob Jones", 
    "program": "Computer Science",
    "year": 1,
    "wallet_address": "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"
})

resp = requests.get(f"{BASE_URL}/api/v1/students", headers=headers)
students = resp.json()["data"]
student = next((s for s in students if s["student_id"] == "student2"), None)
if not student:
    print("Student not found!")
    exit(1)
sid = student["id"]

print("Verifying KYC...")
requests.post(f"{BASE_URL}/api/v1/students/{sid}/kyc", headers=headers)

print("Adding credits...")
requests.post(f"{BASE_URL}/api/v1/students/{sid}/credits", headers=headers, json={
    "credits_to_add": 15, "new_gpa": 3.8
})

time.sleep(2)

print("Triggering L2 claim...")
# Claim logic expects the caller to be a student, but I used auth.RoleAdmin in the router so we can test with admin token
resp = requests.post(f"{BASE_URL}/api/v1/students/student2/claim-l2", headers=headers, json={
    "fund_id": 1,
    "tranche_index": 0,
    "recipient": "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC",
    "dst_chain_id": 42161
})
print("Claim response:", resp.status_code, resp.text)
