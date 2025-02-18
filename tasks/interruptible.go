package tasks

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/DanLavine/goasync/v2"
)

type interruptible struct {
	signals []os.Signal
	subTask goasync.Task
}

//	PARAMETERS:
//	- signals: OS signals to capture and initiate the reload operation
//	- task: task process that is wrapped to handle interruptions
//
// Interruptible Tasks are used to run any tasks that can be reloaded at runtime.
// When this task recieves an interrupt signal it will cancel the child Task process
// and call Cleanup. Then this will call Initialize and Execute again for the child Task.
func Interruptible(signals []os.Signal, task goasync.Task) *interruptible {
	return &interruptible{
		signals: signals,
		subTask: task,
	}
}

// passthrough init and cleanup
func (i *interruptible) Initialize(ctx context.Context) error { return i.subTask.Initialize(ctx) }
func (i *interruptible) Cleanup(ctx context.Context) error    { return i.subTask.Cleanup(ctx) }

func (i *interruptible) Execute(ctx context.Context) error {
	for {
		taskCtx, _ := signal.NotifyContext(ctx, i.signals...)
		if err := i.subTask.Execute(taskCtx); err != nil {
			// if there is an error, always return it
			return err
		}

		select {
		case <-ctx.Done():
			// we are shutting down the application and don't have an error
			return nil
		default:
			select {
			case <-taskCtx.Done():
				// This is the case where the Interupt signal triggered the context to close and not the parent.
				// So we need to restart our task
				if err := i.subTask.Cleanup(ctx); err != nil {
					return err
				}
				if err := i.subTask.Initialize(ctx); err != nil {
					return err
				}
			default:
				return fmt.Errorf("Interruptible task unexpectedly stopped")
			}
		}
	}
}
