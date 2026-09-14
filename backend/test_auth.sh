curl -v -X POST http://localhost:8081/api/v1/auth/token \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "role=student&student_id=CS/001/2021"
