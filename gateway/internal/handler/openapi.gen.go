// Package handler provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package handler

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

// PaperCreate defines model for PaperCreate.
type PaperCreate struct {
	Id        string `json:"id"`
	Published string `json:"published"`
	Subject   string `json:"subject"`
	Title     string `json:"title"`
	Url       string `json:"url"`
}

// PaperGet defines model for PaperGet.
type PaperGet struct {
	CreatedAt string `json:"created_at"`
	Id        string `json:"id"`
	Published string `json:"published"`
	Subject   string `json:"subject"`
	Title     string `json:"title"`
	UpdatedAt string `json:"updated_at"`
	Url       string `json:"url"`
}

// PaperList defines model for PaperList.
type PaperList struct {
	Papers *[]PaperItem `json:"papers,omitempty"`
	Total  int          `json:"total"`
}

// PaperUpdate defines model for PaperUpdate.
type PaperUpdate struct {
	Published *string `json:"published,omitempty"`
	Subject   *string `json:"subject,omitempty"`
	Title     *string `json:"title,omitempty"`
	Url       *string `json:"url,omitempty"`
}

// ProblemDetail エラー表現として、 [Problem Details for HTTP APIs](https://tools.ietf.org/html/rfc7807) を用いる。
// API クライアントでのエラーハンドリングの実装時は、 `type` を利用する。
// リソース個別のエラー表現は、各API仕様に記載する。
type ProblemDetail struct {
	// Detail 個別の Problem を説明するヒューマンリーダブルな文章。
	// 通常、サーバーのエラー文を返す。
	Detail string `json:"detail"`

	// ErrorCode Title のスネークケース表現。Extension Member.
	ErrorCode string `json:"error_code"`

	// Instance エラー発生箇所を示す URI 表現 (リクエストを処理したエンドポイントのURI)
	Instance string `json:"instance"`

	// Status HTTP ステータスコード
	Status int `json:"status"`

	// Title Problem の `type` に対するヒューマンリーダブルな説明文
	Title string `json:"title"`

	// Type Problem の種別を一意に表現するURI
	Type string `json:"type"`
}

// SearchPaper defines model for SearchPaper.
type SearchPaper struct {
	Papers *[]struct {
		Title *string `json:"title,omitempty"`
		Url   *string `json:"url,omitempty"`
	} `json:"papers,omitempty"`
	Total int `json:"total"`
}

// PaperBase defines model for paperBase.
type PaperBase struct {
	Published string `json:"published"`
	Subject   string `json:"subject"`
	Url       string `json:"url"`
}

// PaperItem defines model for paperItem.
type PaperItem struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

// Author defines model for author.
type Author = string

// MaxResult defines model for max_result.
type MaxResult = int

// PaperId defines model for paper-id.
type PaperId = string

// Title defines model for title.
type Title = string

// InternalServerError エラー表現として、 [Problem Details for HTTP APIs](https://tools.ietf.org/html/rfc7807) を用いる。
// API クライアントでのエラーハンドリングの実装時は、 `type` を利用する。
// リソース個別のエラー表現は、各API仕様に記載する。
type InternalServerError = ProblemDetail

// NotFound エラー表現として、 [Problem Details for HTTP APIs](https://tools.ietf.org/html/rfc7807) を用いる。
// API クライアントでのエラーハンドリングの実装時は、 `type` を利用する。
// リソース個別のエラー表現は、各API仕様に記載する。
type NotFound = ProblemDetail

// SearchGetParams defines parameters for SearchGet.
type SearchGetParams struct {
	Title     Title      `form:"title" json:"title"`
	Author    *Author    `form:"author,omitempty" json:"author,omitempty"`
	MaxResult *MaxResult `form:"max_result,omitempty" json:"max_result,omitempty"`
}

// PaperPutJSONRequestBody defines body for PaperPut for application/json ContentType.
type PaperPutJSONRequestBody = PaperUpdate

// PapersPostJSONRequestBody defines body for PapersPost for application/json ContentType.
type PapersPostJSONRequestBody = PaperCreate

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// ヘルスチェック
	// (GET /health)
	HealthGet(w http.ResponseWriter, r *http.Request)
	// 論文削除
	// (DELETE /paper/{paper-id})
	PaperDelete(w http.ResponseWriter, r *http.Request, paperId PaperId)
	// 論文詳細取得
	// (GET /paper/{paper-id})
	PaperGet(w http.ResponseWriter, r *http.Request, paperId PaperId)
	// 論文情報更新
	// (PUT /paper/{paper-id})
	PaperPut(w http.ResponseWriter, r *http.Request, paperId PaperId)
	// 論文一覧取得
	// (GET /papers)
	PapersGet(w http.ResponseWriter, r *http.Request)
	// 論文登録
	// (POST /papers)
	PapersPost(w http.ResponseWriter, r *http.Request)
	// 論文検索
	// (GET /search)
	SearchGet(w http.ResponseWriter, r *http.Request, params SearchGetParams)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// ヘルスチェック
// (GET /health)
func (_ Unimplemented) HealthGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// 論文削除
// (DELETE /paper/{paper-id})
func (_ Unimplemented) PaperDelete(w http.ResponseWriter, r *http.Request, paperId PaperId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// 論文詳細取得
// (GET /paper/{paper-id})
func (_ Unimplemented) PaperGet(w http.ResponseWriter, r *http.Request, paperId PaperId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// 論文情報更新
// (PUT /paper/{paper-id})
func (_ Unimplemented) PaperPut(w http.ResponseWriter, r *http.Request, paperId PaperId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// 論文一覧取得
// (GET /papers)
func (_ Unimplemented) PapersGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// 論文登録
// (POST /papers)
func (_ Unimplemented) PapersPost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// 論文検索
// (GET /search)
func (_ Unimplemented) SearchGet(w http.ResponseWriter, r *http.Request, params SearchGetParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// HealthGet operation middleware
func (siw *ServerInterfaceWrapper) HealthGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.HealthGet(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PaperDelete operation middleware
func (siw *ServerInterfaceWrapper) PaperDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "paper-id" -------------
	var paperId PaperId

	err = runtime.BindStyledParameterWithLocation("simple", false, "paper-id", runtime.ParamLocationPath, chi.URLParam(r, "paper-id"), &paperId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "paper-id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PaperDelete(w, r, paperId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PaperGet operation middleware
func (siw *ServerInterfaceWrapper) PaperGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "paper-id" -------------
	var paperId PaperId

	err = runtime.BindStyledParameterWithLocation("simple", false, "paper-id", runtime.ParamLocationPath, chi.URLParam(r, "paper-id"), &paperId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "paper-id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PaperGet(w, r, paperId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PaperPut operation middleware
func (siw *ServerInterfaceWrapper) PaperPut(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "paper-id" -------------
	var paperId PaperId

	err = runtime.BindStyledParameterWithLocation("simple", false, "paper-id", runtime.ParamLocationPath, chi.URLParam(r, "paper-id"), &paperId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "paper-id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PaperPut(w, r, paperId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PapersGet operation middleware
func (siw *ServerInterfaceWrapper) PapersGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PapersGet(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PapersPost operation middleware
func (siw *ServerInterfaceWrapper) PapersPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PapersPost(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// SearchGet operation middleware
func (siw *ServerInterfaceWrapper) SearchGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchGetParams

	// ------------- Required query parameter "title" -------------

	if paramValue := r.URL.Query().Get("title"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "title"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "title", r.URL.Query(), &params.Title)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "title", Err: err})
		return
	}

	// ------------- Optional query parameter "author" -------------

	err = runtime.BindQueryParameter("form", true, false, "author", r.URL.Query(), &params.Author)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "author", Err: err})
		return
	}

	// ------------- Optional query parameter "max_result" -------------

	err = runtime.BindQueryParameter("form", true, false, "max_result", r.URL.Query(), &params.MaxResult)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "max_result", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SearchGet(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/health", wrapper.HealthGet)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/paper/{paper-id}", wrapper.PaperDelete)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/paper/{paper-id}", wrapper.PaperGet)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/paper/{paper-id}", wrapper.PaperPut)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/papers", wrapper.PapersGet)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/papers", wrapper.PapersPost)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/search", wrapper.SearchGet)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"4eJE8b8BAAD//1wfd/PQIQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
