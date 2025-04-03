package fcrouter

import (
	"context"
	"regexp"
	"strings"

	"github.com/aliyun/fc-runtime-go-sdk/events"
)

type GlobalHandler = func(context.Context, events.HTTPTriggerEvent) (*events.HTTPTriggerResponse, error)
type RouteHandler = func(context.Context, *events.HTTPTriggerEvent, map[string]string) *events.HTTPTriggerResponse

type Route struct {
	pathPattern    string
	regex          *regexp.Regexp
	paramNames     []string
	methodHandlers map[string]RouteHandler
}

type ErrorHandlers struct {
	InternalServerError func(context.Context, *events.HTTPTriggerEvent, any) *events.HTTPTriggerResponse
	MethodNotAllowed    func(context.Context, *events.HTTPTriggerEvent, []string) *events.HTTPTriggerResponse
	NotFound            func(context.Context, *events.HTTPTriggerEvent) *events.HTTPTriggerResponse
	NotHTTPTrigger      func(context.Context, *events.HTTPTriggerEvent) *events.HTTPTriggerResponse
}

type Router struct {
	routes   []*Route
	handlers ErrorHandlers
}

func NewRouter(handlers ErrorHandlers) *Router {
	return &Router{handlers: handlers}
}

func (r *Router) AddRoute(pathPattern string, methodHandlers map[string]RouteHandler) {
	parts := strings.Split(pathPattern, "/")
	regexStr := "^"
	paramNames := []string{}
	for _, part := range parts {
		if part == "" {
			continue
		}
		if strings.HasPrefix(part, ":") {
			paramName := part[1:]
			paramNames = append(paramNames, paramName)
			regexStr += `/(?P<` + paramName + `>[^/]+)`
		} else {
			regexStr += `\/` + regexp.QuoteMeta(part)
		}
	}
	regexStr += "$"

	newRoute := &Route{
		pathPattern:    pathPattern,
		regex:          regexp.MustCompile(regexStr),
		paramNames:     paramNames,
		methodHandlers: methodHandlers,
	}

	r.routes = append(r.routes, newRoute)
}

func (r *Router) FindRoute(path string) (*Route, map[string]string) {
	for _, route := range r.routes {
		matches := route.regex.FindStringSubmatch(path)
		if matches == nil {
			continue
		}

		params := make(map[string]string)
		for i, name := range route.paramNames {
			params[name] = matches[i+1]
		}

		return route, params
	}

	return nil, nil
}

func (r *Router) GetHandler() GlobalHandler {
	return func(ctx context.Context, event events.HTTPTriggerEvent) (response *events.HTTPTriggerResponse, err error) {
		defer func() {
			if err := recover(); err != nil {
				response = r.handlers.InternalServerError(ctx, &event, err)
			}
		}()

		http := &event.TriggerContext.Http
		if http.Method == nil || http.Path == nil || event.Body == nil {
			return r.handlers.NotHTTPTrigger(ctx, &event), nil
		}

		route, params := r.FindRoute(*event.TriggerContext.Http.Path)
		if route == nil {
			return r.handlers.NotFound(ctx, &event), nil
		}

		routeHandler, exists := route.methodHandlers[*event.TriggerContext.Http.Method]
		if exists {
			return routeHandler(ctx, &event, params), nil
		}

		allowedMethods := make([]string, 0, len(route.methodHandlers))
		for method := range route.methodHandlers {
			allowedMethods = append(allowedMethods, method)
		}
		return r.handlers.MethodNotAllowed(ctx, &event, allowedMethods), nil
	}
}
