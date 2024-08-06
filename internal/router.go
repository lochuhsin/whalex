package internal

import "sync"

var uRoute *upstreamRoute
var uOnce sync.Once = sync.Once{}

type upstreamRoute struct {
	route sync.Map
}

func newUpstreamRouter() *upstreamRoute {
	return &upstreamRoute{
		route: sync.Map{},
	}
}

func (u *upstreamRoute) Get(path string) (string, bool) {
	p, ok := u.route.Load(path)
	if !ok {
		return "", ok
	}
	return p.(string), ok
}

func (u *upstreamRoute) Set(p1 string, p2 string) {
	u.route.Store(p1, p2)
}

func GetUpstreamRouter() *upstreamRoute {
	return uRoute
}

func InitUpstreamRoute() {
	uOnce.Do(func() {
		router := newUpstreamRouter()
		router.Set("a", "abcde")
		/**
		 * add path map here
		 */
		uRoute = router
	})
}
