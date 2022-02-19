package behavior_chain

import "context"

/*


 */

type Handler interface {
	HandleFunc(ctx context.Context, args ...interface{}) error
}

type Handler1 struct {
}

func (lh *Handler1) HandleFunc(ctx context.Context, args ...interface{}) error {
	var err error

	if err != nil {
		return err
	}
	return nil
}

type Handler2 struct {
}

func (ph *Handler2) HandleFunc(ctx context.Context, args ...interface{}) error {
	var err error

	if err != nil {
		return err
	}
	return nil
}

type HandlerChain struct {
	handlers []Handler
}

func (hc *HandlerChain) AddHandler(hs ...Handler) {
	if hc.handlers == nil {
		hc.handlers = []Handler{}
	}
	hc.handlers = append(hc.handlers, hs...)
}

func (hc *HandlerChain) Handler(ctx context.Context, args ...interface{}) error {
	for _, ha := range hc.handlers {
		if err := ha.HandleFunc(ctx, args); err != nil {
			return err
		}
	}
	return nil
}
