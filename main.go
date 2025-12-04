package main

import "gofr.dev/pkg/gofr"

// main demonstrates common GoFr features such as routing,
// middleware, query params, path params, JSON responses,
// error handling, and body parsing.
func main() {
	app := gofr.New()

	// -------------------------
	// Global Middleware Example
	// -------------------------
	app.Use(func(next gofr.HandlerFunc) gofr.HandlerFunc {
		return func(ctx *gofr.Context) (any, error) {
			ctx.Logger.Infof("Request: %s %s", ctx.Request.Method, ctx.Request.URL.Path)
			return next(ctx)
		}
	})

	// -------------------------
	// Simple GET Route
	// -------------------------
	app.GET("/greet", func(ctx *gofr.Context) (any, error) {
		return map[string]string{"message": "Hello World!"}, nil
	})

	// -------------------------
	// Path Parameter Example
	// GET /greet/anvesh
	// -------------------------
	app.GET("/greet/{name}", func(ctx *gofr.Context) (any, error) {
		name := ctx.PathParam("name")
		return map[string]string{"message": "Hello " + name}, nil
	})

	// -------------------------
	// Query Parameter Example
	// GET /welcome?name=anvesh
	// -------------------------
	app.GET("/welcome", func(ctx *gofr.Context) (any, error) {
		name := ctx.QueryParam("name")
		if name == "" {
			name = "Guest"
		}
		return map[string]string{"message": "Welcome " + name}, nil
	})

	// -------------------------
	// POST Body Example
	// -------------------------
	app.POST("/user", func(ctx *gofr.Context) (any, error) {
		var user struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		if err := ctx.Bind(&user); err != nil {
			return nil, err
		}

		return map[string]any{
			"message": "User created successfully",
			"user":    user,
		}, nil
	})

	// -------------------------
	// Health Check
	// -------------------------
	app.GET("/health", func(ctx *gofr.Context) (any, error) {
		return map[string]string{"status": "ok"}, nil
	})

	// -------------------------
	// Example Error Route
	// -------------------------
	app.GET("/error", func(ctx *gofr.Context) (any, error) {
		return nil, gofr.NewError("example error occurred")
	})

	// Start server on localhost:8000
	app.Run()
}
