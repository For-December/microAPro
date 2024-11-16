package img

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/visual"
	"microAPro/constant/config"
	"microAPro/utils/logger"
)

type resType struct {
	Code int `json:"code"`
	Data struct {
		XHCAnylora1106StrengthClip  float64 `json:"300XHC_anylora_1106_strength_clip"`
		XHCAnylora1106StrengthModel float64 `json:"300XHC_anylora_1106_strength_model"`
		Field3                      int     `json:"346868647"`
		AlgorithmBaseResp           struct {
			StatusCode    int    `json:"status_code"`
			StatusMessage string `json:"status_message"`
		} `json:"algorithm_base_resp"`
		AnimeoutlineV416StrengthClip  float64       `json:"animeoutlineV4_16_strength_clip"`
		AnimeoutlineV416StrengthModel float64       `json:"animeoutlineV4_16_strength_model"`
		ApplyIdLayer                  string        `json:"apply_id_layer"`
		BinaryDataBase64              []interface{} `json:"binary_data_base64"`
		ClipSkip                      int           `json:"clip_skip"`
		CnMode                        int           `json:"cn_mode"`
		ComfyuiCost                   int           `json:"comfyui_cost"`
		ControlnetWeight              int           `json:"controlnet_weight"`
		DdimSteps                     int           `json:"ddim_steps"`
		I2TTagText                    string        `json:"i2t_tag_text"`
		IdWeight                      float64       `json:"id_weight"`
		ImageUrls                     []string      `json:"image_urls"`
		LogoInfo                      struct {
			AddLogo         bool   `json:"add_logo"`
			Language        int    `json:"language"`
			LogoTextContent string `json:"logo_text_content"`
			Position        int    `json:"position"`
		} `json:"logo_info"`
		LongResolution int `json:"long_resolution"`
		LoraMap        struct {
			XHCAnylora1106 struct {
				StrengthClip  float64 `json:"strength_clip"`
				StrengthModel float64 `json:"strength_model"`
			} `json:"300XHC_anylora_1106"`
			AnimeoutlineV416 struct {
				StrengthClip  float64 `json:"strength_clip"`
				StrengthModel float64 `json:"strength_model"`
			} `json:"animeoutlineV4_16"`
		} `json:"lora_map"`
		Prompt         string   `json:"prompt"`
		ReturnUrl      bool     `json:"return_url"`
		SamplerName    string   `json:"sampler_name"`
		Scale          int      `json:"scale"`
		Scheduler      string   `json:"scheduler"`
		Seed           int      `json:"seed"`
		Strength       float64  `json:"strength"`
		SubPrompts     []string `json:"sub_prompts"`
		TaggerSettings struct {
			Switch      bool     `json:"switch"`
			TaggerTypes []string `json:"tagger_types"`
		} `json:"tagger_settings"`
	} `json:"data"`
	Message     string `json:"message"`
	RequestId   string `json:"request_id"`
	Status      int    `json:"status"`
	TimeElapsed string `json:"time_elapsed"`
}

var BDImg2ImgInChannel = make(chan string, 100)
var BDImg2ImgOutChannel = make(chan string, 100)

func init() {
	go func() {
		for {
			select {
			case url := <-BDImg2ImgInChannel:
				BDImg2ImgOutChannel <- convertImg2Img(url)
			}
		}
	}()
}
func convertImg2Img(inputUrl string) string {
	visual.DefaultInstance.Client.SetAccessKey(config.EnvCfg.DouBaoAccessKey)
	visual.DefaultInstance.Client.SetSecretKey(config.EnvCfg.DouBaoSecretKey)
	visual.DefaultInstance.SetRegion("cn-north-1")
	//visual.DefaultInstance.SetHost("host")

	//请求Body(查看接口文档请求参数-请求示例，将请求参数内容复制到此)
	reqBody := map[string]interface{}{
		"req_key": "img2img_makoto_style_usage",
		"image_urls": []string{
			inputUrl,
		},
		"return_url": true,
		"logo_info": map[string]interface{}{
			"add_logo":          true,
			"position":          0,
			"language":          0,
			"logo_text_content": "这里是明水印内容",
		},
	}

	resp, status, err := visual.DefaultInstance.CVProcess(reqBody)
	fmt.Println(status, err)
	b, _ := json.Marshal(resp)

	res := resType{}

	if err = json.Unmarshal(b, &res); err != nil {
		logger.Error(err)
		return ""
	}
	if len(res.Data.ImageUrls) != 0 {
		return res.Data.ImageUrls[0]
	}
	//fmt.Println(string(b))
	return ""
}
