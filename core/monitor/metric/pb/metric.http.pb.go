// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: metric.proto

package pb

import (
	context "context"
	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
	http1 "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// MetricServiceHandler is the server API for MetricService service.
type MetricServiceHandler interface {
	// POST /query
	QueryWithInfluxFormat(context.Context, *QueryWithInfluxFormatRequest) (*QueryWithInfluxFormatResponse, error)
	// GET /query
	SearchWithInfluxFormat(context.Context, *QueryWithInfluxFormatRequest) (*QueryWithInfluxFormatResponse, error)
	// POST /api/monitor/metric-query
	QueryWithTableFormat(context.Context, *QueryWithTableFormatRequest) (*QueryWithTableFormatResponse, error)
	// GET /api/monitor/metric-query
	SearchWithTableFormat(context.Context, *QueryWithTableFormatRequest) (*QueryWithTableFormatResponse, error)
	// POST /api/monitor/metric-general-query
	GeneralQuery(context.Context, *GeneralQueryRequest) (*GeneralQueryResponse, error)
	// GET /api/monitor/metric-general-query
	GeneralSearch(context.Context, *GeneralQueryRequest) (*GeneralQueryResponse, error)
}

// RegisterMetricServiceHandler register MetricServiceHandler to http.Router.
func RegisterMetricServiceHandler(r http.Router, srv MetricServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		return func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
	}

	add_QueryWithInfluxFormat := func(method, path string, fn func(context.Context, *QueryWithInfluxFormatRequest) (*QueryWithInfluxFormatResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*QueryWithInfluxFormatRequest))
		}
		var QueryWithInfluxFormat_info transport.ServiceInfo
		if h.Interceptor != nil {
			QueryWithInfluxFormat_info = transport.NewServiceInfo("erda.core.monitor.metric.MetricService", "QueryWithInfluxFormat", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in QueryWithInfluxFormatRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["q"]; len(vals) > 0 {
					in.Statement = vals[0]
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, QueryWithInfluxFormat_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_SearchWithInfluxFormat := func(method, path string, fn func(context.Context, *QueryWithInfluxFormatRequest) (*QueryWithInfluxFormatResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*QueryWithInfluxFormatRequest))
		}
		var SearchWithInfluxFormat_info transport.ServiceInfo
		if h.Interceptor != nil {
			SearchWithInfluxFormat_info = transport.NewServiceInfo("erda.core.monitor.metric.MetricService", "SearchWithInfluxFormat", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in QueryWithInfluxFormatRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["q"]; len(vals) > 0 {
					in.Statement = vals[0]
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, SearchWithInfluxFormat_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_QueryWithTableFormat := func(method, path string, fn func(context.Context, *QueryWithTableFormatRequest) (*QueryWithTableFormatResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*QueryWithTableFormatRequest))
		}
		var QueryWithTableFormat_info transport.ServiceInfo
		if h.Interceptor != nil {
			QueryWithTableFormat_info = transport.NewServiceInfo("erda.core.monitor.metric.MetricService", "QueryWithTableFormat", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in QueryWithTableFormatRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["q"]; len(vals) > 0 {
					in.Statement = vals[0]
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, QueryWithTableFormat_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_SearchWithTableFormat := func(method, path string, fn func(context.Context, *QueryWithTableFormatRequest) (*QueryWithTableFormatResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*QueryWithTableFormatRequest))
		}
		var SearchWithTableFormat_info transport.ServiceInfo
		if h.Interceptor != nil {
			SearchWithTableFormat_info = transport.NewServiceInfo("erda.core.monitor.metric.MetricService", "SearchWithTableFormat", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in QueryWithTableFormatRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["q"]; len(vals) > 0 {
					in.Statement = vals[0]
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, SearchWithTableFormat_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GeneralQuery := func(method, path string, fn func(context.Context, *GeneralQueryRequest) (*GeneralQueryResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GeneralQueryRequest))
		}
		var GeneralQuery_info transport.ServiceInfo
		if h.Interceptor != nil {
			GeneralQuery_info = transport.NewServiceInfo("erda.core.monitor.metric.MetricService", "GeneralQuery", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in GeneralQueryRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["q"]; len(vals) > 0 {
					in.Statement = vals[0]
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GeneralQuery_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GeneralSearch := func(method, path string, fn func(context.Context, *GeneralQueryRequest) (*GeneralQueryResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GeneralQueryRequest))
		}
		var GeneralSearch_info transport.ServiceInfo
		if h.Interceptor != nil {
			GeneralSearch_info = transport.NewServiceInfo("erda.core.monitor.metric.MetricService", "GeneralSearch", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				var in GeneralQueryRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["q"]; len(vals) > 0 {
					in.Statement = vals[0]
				}
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GeneralSearch_info)
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_QueryWithInfluxFormat("POST", "/query", srv.QueryWithInfluxFormat)
	add_SearchWithInfluxFormat("GET", "/query", srv.SearchWithInfluxFormat)
	add_QueryWithTableFormat("POST", "/api/monitor/metric-query", srv.QueryWithTableFormat)
	add_SearchWithTableFormat("GET", "/api/monitor/metric-query", srv.SearchWithTableFormat)
	add_GeneralQuery("POST", "/api/monitor/metric-general-query", srv.GeneralQuery)
	add_GeneralSearch("GET", "/api/monitor/metric-general-query", srv.GeneralSearch)
}
