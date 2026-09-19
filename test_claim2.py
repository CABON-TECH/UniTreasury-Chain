import requests
import json
import time

BASE_URL = "http://localhost:8081"

# 1. Login as system admin to get token
resp = requests.post(f"{BASE_URL}/api/v1/auth/login", json={
    "username": "admin",
    "password": "password"
})
token = resp.json()["token"]
headers = {"Authorization": f"Bearer {token}", "Content-Type": "application/json"}

# 2. Create a demo fund
requests.post(f"{BASE_URL}/api/v1/demo/fund", json={"amount": 100000, "role": "admin"})

# 3. Create a student
resp = requests.post(f"{BASE_URL}/api/v1/students", headers=headers, json={
    "student_id": "student2", 
    "name": "Bob Jones", 
    "major": "Computer Science", 
    "wallet_address": "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"
})

resp = requests.get(f"{BASE_URL}/api/v1/students", headers=headers)
print("Students response:", resp.json())
