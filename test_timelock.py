import requests
import time
import sys

BASE_URL = "http://localhost:8081/api/v1"
LOGIN_URL = "http://localhost:8081/api/v1/auth/login"
RPC_URL = "http://localhost:8545"

def login(username, password):
    resp = requests.post(LOGIN_URL, json={"username": username, "password": password})
    resp.raise_for_status()
    return resp.json()["token"]

def advance_time(seconds):
    print(f"Advancing EVM time by {seconds} seconds...")
    payload = {
        "jsonrpc": "2.0",
        "method": "evm_increaseTime",
        "params": [seconds],
        "id": 1
    }
    resp = requests.post(RPC_URL, json=payload)
    resp.raise_for_status()
    
    # Mine a block to ensure the time change takes effect
    payload = {
        "jsonrpc": "2.0",
        "method": "evm_mine",
        "params": [],
        "id": 2
    }
    resp = requests.post(RPC_URL, json=payload)
    resp.raise_for_status()
    print("Time advanced.")

def main():
    print("Testing Feature 5: Multi-Signature Timelock...")
    
    # 1. Login as Admin
    print("Logging in as admin...")
    admin_token = login("admin", "password")
    admin_headers = {"Authorization": f"Bearer {admin_token}"}
    
    # 2. Set Timelock Delay
    delay_seconds = 48 * 3600  # 48 hours
    print(f"Setting timelock delay to {delay_seconds} seconds...")
    resp = requests.post(
        f"{BASE_URL}/treasury/timelock",
        json={"delay": delay_seconds},
        headers=admin_headers
    )
    if resp.status_code != 200:
        print(f"Failed to set timelock delay: {resp.status_code} {resp.text}")
        sys.exit(1)
    
    time.sleep(2) # wait for tx to be mined
    
    # 3. Create a Withdrawal Proposal
    print("Creating withdrawal proposal...")
    resp = requests.post(
        f"{BASE_URL}/treasury/proposals",
        json={
            "recipient": "0x1111222233334444555566667777888899990000",
            "amount": 10,
            "purpose": "Timelock test"
        },
        headers=admin_headers
    )
    if resp.status_code not in (200, 201):
        print(f"Failed to create proposal: {resp.status_code} {resp.text}")
        sys.exit(1)
    
    proposal_id = resp.json()["id"]
    print(f"Proposal created with ID: {proposal_id}")
    time.sleep(2)
    
    # 4. Approve the Proposal
    print("Approving proposal...")
    resp = requests.post(
        f"{BASE_URL}/treasury/proposals/{proposal_id}/approve",
        headers=admin_headers
    )
    if resp.status_code != 200:
        print(f"Failed to approve proposal: {resp.status_code} {resp.text}")
        sys.exit(1)
    
    time.sleep(2)
    
    # 5. Try to execute immediately (Should Fail)
    print("Attempting to execute proposal immediately (should fail)...")
    resp = requests.post(
        f"{BASE_URL}/treasury/proposals/{proposal_id}/execute",
        headers=admin_headers
    )
    if resp.status_code == 200:
        print("ERROR: Execution succeeded but it should have failed due to timelock!")
        sys.exit(1)
    else:
        print(f"Execution failed as expected: {resp.text}")
        
    # 6. Advance Time
    advance_time(delay_seconds + 10)
    time.sleep(2)
    
    # 7. Try to execute again (Should Succeed)
    print("Attempting to execute proposal after timelock expired...")
    resp = requests.post(
        f"{BASE_URL}/treasury/proposals/{proposal_id}/execute",
        headers=admin_headers
    )
    if resp.status_code != 200:
        print(f"Failed to execute proposal: {resp.status_code} {resp.text}")
        sys.exit(1)
        
    print("Executed successfully! Timelock works as expected.")

if __name__ == "__main__":
    main()
