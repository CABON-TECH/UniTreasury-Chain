package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("cmd/api/main.go")
	s := string(b)
	
	// Add userRepo
	s = strings.Replace(s, "proposalRepo := postgres.NewProposalRepo(pool)", "proposalRepo := postgres.NewProposalRepo(pool)\n\tuserRepo := postgres.NewUserRepo(pool)", 1)
	
	// Add authSvc
	s = strings.Replace(s, "jwtMgr := auth.NewJWTManager(cfg.JWTSecret, cfg.JWTExpiryHours)", "jwtMgr := auth.NewJWTManager(cfg.JWTSecret, cfg.JWTExpiryHours)\n\tauthSvc := service.NewAuthService(userRepo, jwtMgr, log)\n\t_ = authSvc.BootstrapAdmin(ctx, \"admin\")", 1)

	ioutil.WriteFile("cmd/api/main.go", []byte(s), 0644)
}
