import requests
import json
import time

BASE_URL = "http://localhost:8081"

resp = requests.post(f"{BASE_URL}/api/v1/auth/login", json={"username": "admin", "password": "password"})
token = resp.json()["token"]
headers = {"Authorization": f"Bearer {token}", "Content-Type": "application/json"}

resp = requests.post(f"{BASE_URL}/api/v1/students", headers=headers, json={
    "student_id": "student2", 
    "name": "Bob Jones", 
    "major": "Computer Science", 
    "wallet_address": "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"
})
print("Create student response:", resp.text)
