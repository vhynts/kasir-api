package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
	"time"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items)
}

func (s *TransactionService) GetTodayReport() (*models.ReportToday, error) {
	now := time.Now()

	startTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	endTime := startTime.AddDate(0, 0, 1)

	return s.repo.GetReportByRange(startTime, endTime)
}

func (s *TransactionService) GetReport(startDateStr, endDateStr string) (*models.ReportToday, error) {
	layout := "2006-01-02"

	startTime, err := time.Parse(layout, startDateStr)
	if err != nil {
		return nil, err
	}

	endTime, err := time.Parse(layout, endDateStr)
	if err != nil {
		return nil, err
	}

	endTime = endTime.AddDate(0, 0, 1)

	return s.repo.GetReportByRange(startTime, endTime)
}
