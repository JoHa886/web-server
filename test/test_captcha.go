package main

import (
	"image/png"
	"net/http"

	"github.com/afocus/captcha"
)

func main1() {
	cap := captcha.New()
	cap.SetFont("comic.ttf")
	// 创建验证码 4个字符 captcha.NUM 字符模式数字类型
	// 返回验证码图像对象以及验证码字符串 后期可以对字符串进行对比 判断验证

	http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request) {
		img, str := cap.Create(6, captcha.ALL)
		png.Encode(w, img)
		println(str)
	})
	http.ListenAndServe(":8085", nil)
}
