import requests
import json
import time
import secrets

BASE_URL = "http://localhost:8081/api/v1"
LOGIN_URL = "http://localhost:8081/api/v1/auth/login"

def login(username, password):
    resp = requests.post(LOGIN_URL, json={"username": username, "password": password})
    resp.raise_for_status()
    return resp.json()["token"]

def main():
    print("Testing ZK Identity Feature...")
    
    # 1. Login as Admin
    print("Logging in as admin...")
    admin_token = login("admin", "password")
    admin_headers = {"Authorization": f"Bearer {admin_token}"}
    
    # 2. Login as Student
    print("Logging in as student1...")
    student_token = login("student1", "password")
    student_headers = {"Authorization": f"Bearer {student_token}"}
    
    # 3. Update Merkle Root (Admin)
    dummy_root = "112233445566778899001122334455667788990011223344556677889900aabb"
    print("Updating Merkle Root...")
    resp = requests.post(
        f"{BASE_URL}/zk/update-root",
        json={"merkle_root": dummy_root},
        headers=admin_headers
    )
    if resp.status_code != 200:
        print(f"Update Root failed: {resp.status_code} {resp.text}")
        return
    print(f"Root updated: {resp.json().get('tx_hash', 'success')}")
    
    # Wait for block to mine
    time.sleep(2)
    
    # 4. Fetch Current Root
    print("Fetching current root...")
    resp = requests.get(f"{BASE_URL}/zk/root", headers=student_headers)
    if resp.status_code != 200:
        print(f"Get Root failed: {resp.status_code} {resp.text}")
        return
    
    current_root = resp.json().get("merkle_root", "")
    print(f"Current Root on-chain: {current_root}")
    if current_root != dummy_root:
        print("Error: Root mismatch!")
        # Note: keeping going anyway to test the flow
    
    # 5. Submit Proof of Enrollment (Student)
    dummy_proof = "01020304" # Just needs to be non-empty and start with non-zero
    dummy_nullifier = secrets.token_hex(32)
    
    print("Submitting ZK Enrollment Proof...")
    resp = requests.post(
        f"{BASE_URL}/zk/prove-enrollment",
        json={
            "proof": dummy_proof,
            "nullifier_hash": dummy_nullifier,
            "claimed_root": current_root
        },
        headers=student_headers
    )
    if resp.status_code != 200:
        print(f"Prove Enrollment failed: {resp.status_code} {resp.text}")
        return
    print(f"Enrollment Proved: {resp.json().get('tx_hash', 'success')}")
    
    # Wait for block to mine
    time.sleep(2)
    
    # 6. Check Nullifier
    print("Checking Nullifier...")
    resp = requests.get(f"{BASE_URL}/zk/nullifier/{dummy_nullifier}", headers=student_headers)
    if resp.status_code != 200:
        print(f"Check Nullifier failed: {resp.status_code} {resp.text}")
        return
    
    data = resp.json()
    if data.get("used"):
        print(f"Success! Nullifier {dummy_nullifier} is marked as used by {data.get('student', 'unknown')}")
    else:
        print("Error: Nullifier is not marked as used!")

if __name__ == "__main__":
    main()
