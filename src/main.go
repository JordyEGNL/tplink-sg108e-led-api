package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type RequestData struct {
	IP       string `json:"ip"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseData struct {
	Status  string `json:"status"`
	Action  string `json:"action,omitempty"`
	Message string `json:"message,omitempty"`
}

func loginAndControlLED(ip, username, password, action string) ResponseData {
	loginURL := "http://" + ip + "/logon.cgi"
	ledURL := "http://" + ip + "/led_on_set.cgi?rd_led=" + action + "&led_cfg=Apply"

	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	loginData := url.Values{
		"logon":    {"Login"},
		"username": {username},
		"password": {password},
	}

	loginReq, err := http.NewRequest("POST", loginURL, strings.NewReader(loginData.Encode()))
	if err != nil {
		return ResponseData{Status: "error", Message: err.Error()}
	}
	loginReq.Header.Set("Referer", "http://"+ip+"/Logout.htm")
	loginReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	loginResp, err := client.Do(loginReq)
	if err != nil {
		return ResponseData{Status: "error", Message: err.Error()}
	}
	defer loginResp.Body.Close()

	if loginResp.StatusCode != http.StatusOK {
		return ResponseData{Status: "error", Message: "Login failed"}
	}

	ledReq, err := http.NewRequest("GET", ledURL, nil)
	if err != nil {
		return ResponseData{Status: "error", Message: err.Error()}
	}
	ledReq.Header.Set("Referer", "http://"+ip+"/")

	ledResp, err := client.Do(ledReq)
	if err != nil {
		return ResponseData{Status: "error", Message: err.Error()}
	}
	defer ledResp.Body.Close()

	if ledResp.StatusCode != http.StatusOK {
		return ResponseData{Status: "error", Message: "LED control failed"}
	}

	return ResponseData{Status: "success", Action: action}
}

func ledControlHandler(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data RequestData
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
			return
		}

		ip := data.IP
		if ip == "" {
			ip = os.Getenv("TP_LINK_IP")
		}
		username := data.Username
		if username == "" {
			username = os.Getenv("TP_LINK_USERNAME")
		}
		password := data.Password
		if password == "" {
			password = os.Getenv("TP_LINK_PASSWORD")
		}

		if ip == "" || username == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Missing required parameters"})
			return
		}

		result := loginAndControlLED(ip, username, password, action)
		c.JSON(http.StatusOK, result)
	}
}

func main() {
	fmt.Println("Starting server...")
	r := gin.Default()

	r.POST("/led_on", ledControlHandler("1"))
	r.POST("/led_off", ledControlHandler("0"))

	fmt.Println("Listening on 0.0.0.0:5000")
	r.Run(":5000")
}
