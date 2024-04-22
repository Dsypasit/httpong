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

func (r *Router) FindRoute(req Req) *Route {
	for _, route := range r.route {
		if route.Path == req.Path && route.Method == req.Method {
			return &route
		}
	}
	return nil
}

func (r *Router) registerRoute(rg Route) error {
	r.route = append(r.route, rg)
	return nil
}
