import re

with open("backend/internal/service/scholarship_service.go", "r") as f:
    text = f.read()

# I will just replace the logic in SimulateGaslessClaim to handle the lock correctly
old_block = """	opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
	if err != nil {
		return err
	}
	
	rootBytes, _ := s.escrow.ScholarshipEscrowContractCaller.TrancheRoots(nil, onChainFundId, tIndex)
	var emptyRoot [32]byte
	if rootBytes == emptyRoot {
	    var treeRoot32 [32]byte
	    copy(treeRoot32[:], tree[len(tree)-1][0])
	    _, err := s.escrow.ScholarshipEscrowContractTransactor.PublishTrancheRoot(opts, onChainFundId, tIndex, treeRoot32)
	    if err != nil {
	        rollback()
	        return err
	    }
	    confirm()
	    
	    opts, confirm, rollback, err = s.txMgr.TransactOpts(ctx)
	    if err != nil {
	        return err
	    }
	}"""

new_block = """	rootBytes, _ := s.escrow.ScholarshipEscrowContractCaller.TrancheRoots(nil, onChainFundId, tIndex)
	var emptyRoot [32]byte
	if rootBytes == emptyRoot {
	    opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
	    if err != nil {
	        return err
	    }
	    var treeRoot32 [32]byte
	    copy(treeRoot32[:], tree[len(tree)-1][0])
	    _, err = s.escrow.ScholarshipEscrowContractTransactor.PublishTrancheRoot(opts, onChainFundId, tIndex, treeRoot32)
	    if err != nil {
	        rollback()
	        return err
	    }
	    confirm()
	}"""

text = text.replace(old_block, new_block)

# And make sure op := bindings.UserOperation... has a TransactOpts before it
text = text.replace("""	callData, err := escrowAbi.Pack("claimTranche", onChainFundId, studentHash, tIndex, recipient, claimAmount, proof32)
	if err != nil {
		return err
	}

	opts, confirm, rollback, err = s.txMgr.TransactOpts(ctx)""", """	callData, err := escrowAbi.Pack("claimTranche", onChainFundId, studentHash, tIndex, recipient, claimAmount, proof32)
	if err != nil {
		return err
	}

	opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)""")

with open("backend/internal/service/scholarship_service.go", "w") as f:
    f.write(text)
