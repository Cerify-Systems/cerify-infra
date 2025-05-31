package services

import (
	"github.com/cerify/cerify-infra/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Services struct {
	Contract    *ContractService
	Verification *VerificationService
	Analysis    *AnalysisService
}

func New(cfg *config.Config, db *mongo.Client, logger *zap.Logger) (*Services, error) {
	contractService := NewContractService(cfg, db, logger)
	verificationService := NewVerificationService(cfg, db, logger)
	analysisService := NewAnalysisService(cfg, db, logger)

	return &Services{
		Contract:    contractService,
		Verification: verificationService,
		Analysis:    analysisService,
	}, nil
}

type ContractService struct {
	cfg    *config.Config
	db     *mongo.Client
	logger *zap.Logger
}

func NewContractService(cfg *config.Config, db *mongo.Client, logger *zap.Logger) *ContractService {
	return &ContractService{
		cfg:    cfg,
		db:     db,
		logger: logger,
	}
}

type VerificationService struct {
	cfg    *config.Config
	db     *mongo.Client
	logger *zap.Logger
}

func NewVerificationService(cfg *config.Config, db *mongo.Client, logger *zap.Logger) *VerificationService {
	return &VerificationService{
		cfg:    cfg,
		db:     db,
		logger: logger,
	}
}

type AnalysisService struct {
	cfg    *config.Config
	db     *mongo.Client
	logger *zap.Logger
}

func NewAnalysisService(cfg *config.Config, db *mongo.Client, logger *zap.Logger) *AnalysisService {
	return &AnalysisService{
		cfg:    cfg,
		db:     db,
		logger: logger,
	}
} 