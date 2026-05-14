package closer

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.uber.org/zap"
)

type closeFn struct {
	name string
	fn   func(context.Context) error
}

type closer struct {
	mu    sync.Mutex
	once  sync.Once
	funcs []closeFn
}

var globalCloser = &closer{}

func Add(name string, fn func(context.Context) error) {
	globalCloser.add(name, fn)
}

func CloseAll(ctx context.Context, logger *zap.Logger) error {
	return globalCloser.closeAll(ctx, logger)
}

func (c *closer) add(name string, fn func(context.Context) error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.funcs = append(c.funcs, closeFn{name: name, fn: fn})
}

func (c *closer) closeAll(ctx context.Context, logger *zap.Logger) error {
	var result error

	c.once.Do(func() {
		c.mu.Lock()
		funcs := c.funcs
		c.funcs = nil
		c.mu.Unlock()

		if len(funcs) == 0 {
			return
		}

		logger.Info("начинаем плавное завершение", zap.Any("count", len(funcs)))

		var errs []error

		for i := len(funcs) - 1; i >= 0; i-- {
			f := funcs[i]

			start := time.Now()
			logger.Info("закрываем ресурс", zap.Any("name", f.name))

			if err := f.fn(ctx); err != nil {
				logger.Error("ошибка при закрытии ресурса",
					zap.Any("name", f.name),
					zap.Error(err),
					zap.Any("duration", time.Since(start)),
				)

				errs = append(errs, err)
			} else {
				logger.Info(
					"ресурс закрыт",
					zap.Any("name", f.name),
					zap.Any("duration", time.Since(start)),
				)
			}
		}

		logger.Info("все ресурсы закрыты", zap.Any("count", len(funcs)))

		result = errors.Join(errs...)
	})

	return result
}
