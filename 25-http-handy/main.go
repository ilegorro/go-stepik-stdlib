package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Handy предоставляет удобный интерфейс
// для выполнения HTTP-запросов
type Handy struct {
	uri     string
	client  *http.Client
	headers map[string]string
	params  map[string]string
	form    map[string]string
	json    any
}

// NewHandy создает новый экземпляр Handy
func NewHandy() *Handy {
	return &Handy{
		"",
		http.DefaultClient,
		make(map[string]string),
		make(map[string]string),
		make(map[string]string),
		nil,
	}
}

// URL устанавливает URL, на который пойдет запрос
func (h *Handy) URL(uri string) *Handy {
	h.uri = uri
	return h
}

// Client устанавливает HTTP-клиента
// вместо умолчательного http.DefaultClient
func (h *Handy) Client(client *http.Client) *Handy {
	h.client = client
	return h
}

// Header устанавливает значение заголовка
func (h *Handy) Header(key, value string) *Handy {
	h.headers[key] = value
	return h
}

// Param устанавливает значение URL-параметра
func (h *Handy) Param(key, value string) *Handy {
	h.params[key] = value
	return h
}

// Form устанавливает данные, которые будут закодированы
// как application/x-www-form-urlencoded и отправлены в теле запроса
// с соответствующим content-type
func (h *Handy) Form(form map[string]string) *Handy {
	h.json = nil
	h.form = form
	return h
}

// JSON устанавливает данные, которые будут закодированы
// как application/json и отправлены в теле запроса
// с соответствующим content-type
func (h *Handy) JSON(v any) *Handy {
	h.form = nil
	h.json = v
	return h
}

// Get выполняет GET-запрос с настроенными ранее параметрами
func (h *Handy) Get() *HandyResponse {
	handyResp := &HandyResponse{}
	req, err := http.NewRequest(http.MethodGet, h.uri, nil)
	if err != nil {
		handyResp.RespErr = err
		return handyResp
	}

	p := url.Values{}
	for k, v := range h.params {
		p.Set(k, v)
	}
	req.URL.RawQuery = p.Encode()

	for k, v := range h.headers {
		req.Header.Set(k, v)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		handyResp.RespErr = err
		return handyResp
	}
	defer resp.Body.Close()

	handyResp.StatusCode = resp.StatusCode
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		handyResp.RespErr = err
		return handyResp
	}
	handyResp.RespBody = body

	return handyResp
}

// Post выполняет POST-запрос с настроенными ранее параметрами
func (h *Handy) Post() *HandyResponse {
	var resp *http.Response
	var req *http.Request
	var err error
	var b []byte

	handyResp := &HandyResponse{}
	if h.form != nil {
		f := url.Values{}
		for k, v := range h.form {
			f.Set(k, v)
		}
		rd := bytes.NewReader([]byte(f.Encode()))
		req, err = http.NewRequest(http.MethodPost, h.uri, rd)
		if err != nil {
			handyResp.RespErr = err
			return handyResp
		}

		for k, v := range h.headers {
			req.Header.Set(k, v)
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		b, err = json.Marshal(h.json)
		if err != nil {
			handyResp.RespErr = err
			return handyResp
		}
		req, err = http.NewRequest(http.MethodPost, h.uri, bytes.NewReader(b))
		if err != nil {
			handyResp.RespErr = err
			return handyResp
		}

		for k, v := range h.headers {
			req.Header.Set(k, v)
		}
		req.Header.Add("Content-Type", "application/json")
	}

	p := url.Values{}
	for k, v := range h.params {
		p.Set(k, v)
	}
	req.URL.RawQuery = p.Encode()

	resp, err = h.client.Do(req)
	if err != nil {
		handyResp.RespErr = err
		return handyResp
	}
	defer resp.Body.Close()

	handyResp.StatusCode = resp.StatusCode
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		handyResp.RespErr = err
		return handyResp
	}
	handyResp.RespBody = body

	return handyResp
}

// HandyResponse представляет ответ на HTTP-запрос
type HandyResponse struct {
	StatusCode int
	RespErr    error
	RespBody   []byte
}

// OK возвращает true, если во время выполнения запроса
// не произошло ошибок, а код HTTP-статуса ответа равен 200
func (r *HandyResponse) OK() bool {
	return r.StatusCode == 200 && r.RespErr == nil
}

// Bytes возвращает тело ответа как срез байт
func (r *HandyResponse) Bytes() []byte {
	return r.RespBody
}

// String возвращает тело ответа как строку
func (r *HandyResponse) String() string {
	return string(r.RespBody)
}

// JSON декодирует тело ответа из JSON и сохраняет
// результат по адресу, на который указывает v
func (r *HandyResponse) JSON(v any) {
	// работает аналогично json.Unmarshal()
	// если при декодировании произошла ошибка,
	// она должна быть доступна через r.Err()
	err := json.Unmarshal(r.RespBody, &v)
	if err != nil {
		r.RespErr = err
	}
}

// Err возвращает ошибку, которая возникла при выполнении запроса
// или обработке ответа
func (r *HandyResponse) Err() error {
	return r.RespErr
}

func main() {
	{
		// примеры запросов

		// GET-запрос с параметрами
		NewHandy().URL("https://httpbingo.org/get").Param("id", "42").Get()

		// HTTP-заголовки
		NewHandy().
			URL("https://httpbingo.org/get").
			Header("Accept", "text/html").
			Header("Authorization", "Bearer 1234567890").
			Get()

		// POST формы
		params := map[string]string{
			"brand":    "lg",
			"category": "tv",
		}
		NewHandy().URL("https://httpbingo.org/post").Form(params).Post()

		// POST JSON-документа
		NewHandy().URL("https://httpbingo.org/post").JSON(params).Post()
	}

	{
		// пример обработки ответа

		// отправляем GET-запрос с параметрами
		resp := NewHandy().URL("https://httpbingo.org/get").Param("id", "42").Get()
		if !resp.OK() {
			panic(resp.String())
		}

		// декодируем ответ в JSON
		var data map[string]any
		resp.JSON(&data)

		fmt.Println(data["url"])
		// "https://httpbingo.org/get"
		fmt.Println(data["args"])
		// map[id:[42]]
	}
}
