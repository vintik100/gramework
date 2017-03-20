package gramework

// ToTLSHandler returns handler that redirects user to HTTP scheme
func (app *App) ToTLSHandler() func(*Context) {
	return func(ctx *Context) {
		ctx.ToTLS()
	}
}

// HTTP returns HTTP-only router
func (app *App) HTTP() *Router {
	return app.defaultRouter.HTTP()
}

// HTTPS returns HTTPS-only router
func (app *App) HTTPS() *Router {
	return app.defaultRouter.HTTPS()
}

// MethodNotAllowed sets MethodNotAllowed handler
func (app *App) MethodNotAllowed(c func(ctx *Context)) *App {
	app.defaultRouter.router.MethodNotAllowed = c
	return app
}
