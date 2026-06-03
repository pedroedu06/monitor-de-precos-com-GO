package worker

import (
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/model"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/repository"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/scrapper"
)

const maxConcurrentScrapes = 10

type Worker struct {
	productRepo      *repository.ProdutoRepository
	priceHistoryRepo *repository.PriceHistoryRepository
	scrapper         *scrapper.Scrapper
	interval         time.Duration
}

func NewWorker(
	productRepo *repository.ProdutoRepository,
	priceHistoryRepo *repository.PriceHistoryRepository,
	scrapper *scrapper.Scrapper,
) *Worker {
	return &Worker{
		productRepo:      productRepo,
		priceHistoryRepo: priceHistoryRepo,
		scrapper:         scrapper,
		interval:         1 * time.Hour, // TODO: load from env
	}
}

func (w *Worker) StartWorker() {
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	w.VerifyAll()

	for range ticker.C {
		w.VerifyAll()
	}
}

func (w *Worker) VerifyAll() {
	products, err := w.productRepo.ListAllActive()
	if err != nil {
		log.Printf("worker: failed to list active products: %v", err)
		return
	}

	sem := make(chan struct{}, maxConcurrentScrapes)
	var wg sync.WaitGroup

	for _, p := range products {
		wg.Add(1)
		sem <- struct{}{} // blocks here when 10 are already running
		go func(product model.ProdutoDomain) {
			defer wg.Done()
			defer func() { <-sem }()
			w.verifyProduct(product)
		}(p)
	}

	wg.Wait()
}

func (w *Worker) verifyProduct(p model.ProdutoDomain) {
	currentPrice, err := w.scrapper.PriceColector(p.URL)
	if err != nil {
		log.Printf("worker: scrape failed for product %s: %v", p.ID, err)
		return
	}

	latest, err := w.priceHistoryRepo.GetLatestPrice(p.ID)
	switch {
	case err == sql.ErrNoRows:
		
	case err != nil:
		log.Printf("worker: failed to fetch latest price for product %s: %v", p.ID, err)
		return
	case latest.Price == currentPrice:
		return
	}

	if _, err := w.priceHistoryRepo.InsertPrice(p.ID, currentPrice); err != nil {
		log.Printf("worker: failed to insert price for product %s: %v", p.ID, err)
		return
	}
}
