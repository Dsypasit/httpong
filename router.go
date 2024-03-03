package httpong

type Handler func(*Context) error

type Router struct {
	route []Route
}

type Route struct {
	Method   string
	Path     string
	Function Handler
}

func newRouter() Router {
	return Router{
		route: []Route{},
	}
}

func (r *Router) registerRoute(rg Route) error {
	r.route = append(r.route, rg)
	return nil
}
