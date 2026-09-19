import re

with open("backend/internal/service/scholarship_service.go", "r") as f:
    text = f.read()

# Fix line 345
text = text.replace("""	opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
	if err != nil {
		return fmt.Errorf("transact opts: %w", err)
	}
	
		opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
	if err != nil {
		return err
	}
	rootBytes""", """	opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
	if err != nil {
		return fmt.Errorf("transact opts: %w", err)
	}
	
	rootBytes""", 1)

# Fix Gasless logic
# Find SimulateGaslessClaim
gasless = text[text.find("func (s *ScholarshipService) SimulateGaslessClaim"):]
# Inside gasless, find the Publish block
pub_block = gasless[gasless.find("rootBytes, _ := s.escrow.ScholarshipEscrowContractCaller.TrancheRoots"):gasless.find("callData, err := escrowAbi.Pack")]

new_pub_block = """
	opts, confirm, rollback, err := s.txMgr.TransactOpts(ctx)
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
	}
"""
text = text.replace(pub_block, new_pub_block)

with open("backend/internal/service/scholarship_service.go", "w") as f:
    f.write(text)

