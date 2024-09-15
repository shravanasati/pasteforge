package pastes

import (
	"context"
	"time"
)

func (h *Handler) DeleteExpiredPastes(ticker *time.Ticker, done chan struct{}) {
	for {
		select {
		case <-ticker.C:
			err := h.pasteStore.DeleteExpiredPastes(context.Background())
			if err != nil {
				h.logger.Error("failed to delete expired pastes", "err", err)
			} else {
				h.logger.Info("successfully deleted expired pastes")
			}
		case <-done:
			h.logger.Info("DeleteExpiredPastes goroutine shutting down")
			return
		default:
			time.Sleep(time.Second)
		}
	}
}
