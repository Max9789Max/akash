package types

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"
)

type Route interface {
	Prefix() string
	Run(path string) ([]byte, error)
}

type Router interface {
	Accept(path string) bool
	Run(path string) ([]byte, error)
}

// how to handle trailing slashes?
// /account/{account-id}
// vs
// /accounts

func NewRouter(routes []Route) Router {
	return router{routes}
}

type router struct {
	routes []Route
}

func (r router) Accept(path string) bool {
	route, _ := r.resolve(path)
	return route != nil
}

func (r router) Run(path string) ([]byte, error) {
	route, suffix := r.resolve(path)
	if route == nil {
		return nil, fmt.Errorf("unknown route: %v", path)
	}
	return route.Run(suffix)
}

func (r router) resolve(path string) (Route, string) {

	for _, route := range r.routes {
		if strings.HasPrefix(path, route.Prefix()) {
			return route, strings.TrimPrefix(path, route.Prefix())
		}
	}
	return nil, ""
}

type route struct {
	prefix  string
	handler Handler
}

func NewRoute(prefix string, handler Handler) Route {
	return route{prefix: prefix, handler: handler}
}

func (r route) Prefix() string {
	return r.prefix
}

func (r route) Run(path string) ([]byte, error) {
	return r.handler.Run(path)
}

type Handler interface {
	Run(string) ([]byte, error)
}

type handler struct {
	parser     Parser
	collection bool
	adapter    Adapter
}

func (h handler) Run(path string) ([]byte, error) {
	id, err := h.parser.ParseText(path)
	if err != nil {
		return nil, err
	}

	if !h.collection && !id.Complete() {
		return nil, errors.New("single return value only")
	}

	var obj proto.Message

	if h.collection {
		obj, err = h.adapter.List(id)
	} else {
		obj, err = h.adapter.Get(id)
	}
	if err != nil {
		return nil, err
	}

	return proto.Marshal(obj)
}

type Parser interface {
	ParseText(string) (Key, error)
}

type Key interface {
	Complete() bool
}

type Adapter interface {
	List(Key) (proto.Message, error)
	Get(Key) (proto.Message, error)
}
