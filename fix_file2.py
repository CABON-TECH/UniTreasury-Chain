import re

with open("backend/internal/service/scholarship_service.go", "r") as f:
    text = f.read()

# Replace the one after callData in Gasless
text = text.replace("""	opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
	if err != nil {
		return err
	}

	op :=""", """	opts, confirm, rollback, err = s.txMgr.TransactOpts(ctx)
	if err != nil {
		return err
	}

	op :=""")

with open("backend/internal/service/scholarship_service.go", "w") as f:
    f.write(text)
