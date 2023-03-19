package external_command

import (
	"context"
	"log"
	"strconv"
	"sync"

	"github.com/IgorAleksandroff/agent-status/internal/entity"
	"github.com/IgorAleksandroff/agent-status/internal/usecase/external_command/commands"
)

const countOfWorkers = 3

type Factory interface {
	GetCommandFromType(commandType entity.CommandType, params map[string]string) (commands.Base, commands.Executor, error)
}

type Worker struct {
	queue   chan entity.Event
	factory Factory
}

func NewWorker(q chan entity.Event, f Factory) *Worker {
	return &Worker{
		queue:   q,
		factory: f,
	}
}

func (w *Worker) Run(ctx context.Context) {
	wg := &sync.WaitGroup{}

	wg.Add(countOfWorkers)
	for i := 0; i < countOfWorkers; i++ {
		go func() {
			w.run(ctx)
			wg.Done()
		}()
	}
	wg.Wait()
}

func (w *Worker) run(ctx context.Context) {
	log.Println("worker: start")

	for {
		event, opened := <-w.queue
		if !opened {
			log.Println("worker: stop")
			return
		}

		p := *event.Params

		c, e, err := w.factory.GetCommandFromType(event.Type, p)
		if err != nil {
			log.Printf("failed to procces event: %+v", event)
			continue
		}

		if !e.ValidityCheck(ctx, c) {
			continue
		}

		err = e.Execute(ctx, c)
		if err == nil {
			continue
		}
		log.Printf("failed to execute command: %+v", c)

		if e.Retry(ctx, c) {
			counterParam := p["counter"]
			counter, err := strconv.Atoi(counterParam)
			if err != nil {
				log.Printf("failed to parse counter: %+v", counterParam)
				continue
			}

			p["counter"] = strconv.Itoa(counter - 1)
			event.Params = &p

			w.queue <- event
		}
	}
}

func (w *Worker) Shutdown() {
	close(w.queue)
}
