package worker

import (
	"context"
	"log"

	"github.com/user/quantum-server/internal/dto"
	"github.com/user/quantum-server/internal/repository"
	"github.com/user/quantum-server/internal/service"
)

type Worker interface {
	Start(ctx context.Context)
}

type calculationWorker struct {
	mortgageService service.MortgageService
	calcRepo        repository.MortgageCalculationRepository
	taskChan        <-chan dto.MortgageTask
}

func NewCalculationWorker(mortgageService service.MortgageService, calcRepo repository.MortgageCalculationRepository, taskChan <-chan dto.MortgageTask) Worker {
	return &calculationWorker{
		mortgageService: mortgageService,
		calcRepo:        calcRepo,
		taskChan:        taskChan,
	}
}

func (w *calculationWorker) Start(ctx context.Context) {
	log.Println("Calculation worker started")
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case t, ok := <-w.taskChan:
				if !ok {
					return
				}
				w.processTask(ctx, t)
			}
		}
	}()
}

func (w *calculationWorker) processTask(ctx context.Context, t dto.MortgageTask) {
	log.Printf("Processing task for calculation ID: %d", t.CalcID)

	err := w.mortgageService.ProcessCalculation(ctx, t.CalcID, t.Request)
	if err != nil {
		log.Printf("Error processing calculation %d: %v", t.CalcID, err)
	}
}
