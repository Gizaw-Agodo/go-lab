package runner

import (
	"context"
	"fmt"
	"time"
)

func Worker(ctx context.Context, task Task) {

	for {
		select {

		case <-ctx.Done():
			fmt.Printf(
				"Task %d stopped: %v\n",
				task.ID,
				ctx.Err(),
			)
			return

		default:
			fmt.Printf(
				"Task %d (%s) is processing...\n",
				task.ID,
				task.Name,
			)

			time.Sleep(time.Second)
		}
	}
}