package gogpt

import (
  "bytes"
  "context"
  "encoding/json"
  "net/http"
  "net/url"
  "strconv"
)

type OverviewRequest struct {
  Prompt           string  `json:"prompt"`
  Model            string  `json:"model"`
  Temperature      float64 `json:"temperature"`
  MaxTokens        int     `json:"max_tokens"`
  TopP             float64 `json:"top_p"`
  FrequencyPenalty float64 `json:"frequency_penalty"`
  PresencePenalty  float64 `json:"presence_penalty"`
  SessionID        string  `json:"session_id"`
}

//
func (c *Client) CreateOverview(
  ctx context.Context,
  request OverviewRequest,
) (response CompletionResponse, err error) {
  var reqBytes []byte
  reqBytes, err = json.Marshal(request)
  if err != nil {
    return
  }

  urlSuffix := "/overview"
  req, err := http.NewRequest("POST", c.assistantURL(urlSuffix), bytes.NewBuffer(reqBytes))
  if err != nil {
    return
  }

  req = req.WithContext(ctx)
  err = c.sendRequest(req, &response)
  return
}

//func (c *Client) CreateOverview(ctx context.Context, request OverviewRequest) (response CompletionResponse, err error) {
//  urlSuffix := "/overview"
//  apiURL := c.assistantURL(urlSuffix)
//  requestURL := buildURL(apiURL, request)
//  req, err := http.NewRequest("GET", requestURL, nil)
//  if err != nil {
//    return
//  }
//
//  req = req.WithContext(ctx)
//  err = c.sendRequest(req, &response)

//
//// Send the request
//resp, err := http.Get(requestURL)
//if err != nil {
//  log.Fatal(err)
//}
//
//// Read the response
//defer resp.Body.Close()
//body, err := ioutil.ReadAll(resp.Body)
//if err != nil {
//  return
//}
//// Print the response
//fmt.Println(string(body))
//  return
//}

/*
params := make(map[string]string)
  params["prompt"] = request.Prompt
  params["model"] = request.Model
  params["temperature"] = strconv.FormatFloat(request.Temperature, 'f', -1, 64)
  params["max_tokens"] = strconv.FormatInt(int64(request.MaxTokens), 10)
  params["top_p"] = strconv.FormatFloat(request.TopP, 'f', -1, 64)
  params["frequency_penalty"] = strconv.FormatFloat(request.FrequencyPenalty, 'f', -1, 64)
  params["presence_penalty"] = strconv.FormatFloat(request.PresencePenalty, 'f', -1, 64)
  params["session_id"] = request.SessionID
*/
func buildURL(apiURL string, request OverviewRequest) (urlStr string) {
  queryParams := url.Values{}
  queryParams.Add("prompt", request.Prompt)
  queryParams.Add("model", request.Model)
  queryParams.Add("temperature", strconv.FormatFloat(request.Temperature, 'f', -1, 64))
  queryParams.Add("max_tokens", strconv.FormatInt(int64(request.MaxTokens), 10))
  queryParams.Add("top_p", strconv.FormatFloat(request.TopP, 'f', -1, 64))
  queryParams.Add("frequency_penalty", strconv.FormatFloat(request.FrequencyPenalty, 'f', -1, 64))
  queryParams.Add("presence_penalty", strconv.FormatFloat(request.PresencePenalty, 'f', -1, 64))
  queryParams.Add("session_id", request.SessionID)
  // 构建 URL
  u, err := url.Parse(apiURL)
  if err != nil {
    // 处理错误
  }
  u.RawQuery = queryParams.Encode() // 将查询参数添加到 URL 中
  urlStr = u.String()               // 获取完整的 URL 字符串
  return
}
