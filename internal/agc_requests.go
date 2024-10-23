// AppGallery Connect Requests
package internal

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

type User struct {
	Name   string
	Email  string
	UserID string
	Client resty.Client
}

// https://developer.huawei.com/consumer/cn/service/josp/agc/index.html#/harmonyOSDevPlatform/9249519184596237889
func openBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin": // macOS
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Start() // start the command
}

func openRodBrowser(url string) ([]*proto.NetworkCookie, error) {
	u := launcher.New().
		Headless(false).
		MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(url)
	page.MustWaitLoad()
	initialURL := page.MustInfo().URL
	var cookies []*proto.NetworkCookie
	for {
		currentURL := page.MustInfo().URL
		if currentURL != initialURL {
			currentCookies, err := page.Cookies([]string{url})
			if err != nil {
				return nil, fmt.Errorf("Error while getting cookies: %v", err)
			}
			cookies = currentCookies
			break
		}
		time.Sleep(1000 * time.Millisecond) // hang
	}

	return cookies, nil
}

// login and return the user cookies
func (u *User) Login() ([]*proto.NetworkCookie, error) {
	// loginURL := "https://developer.huawei.com/consumer/cn/service/josp/agc/index.html"
	loginURL := "https://id1.cloud.huawei.com/CAS/portal/loginAuth.html"

	cookies, err := openRodBrowser(loginURL)
	if err != nil {
		log.Fatalf("Error while opening browser: %v", err)
		return nil, err
	}
	// print the cookies
	fmt.Println("Cookies from return value:")
	for _, cookie := range cookies {
		fmt.Printf("Cookie: %s = %s\n", cookie.Name, cookie.Value)
	}
	return cookies, nil
}

func InitBrowser() (*User, error) {
	user := &User{}
	// user login
	rodCookies, err := user.Login()
	if err != nil {
		fmt.Printf("Error during login: %v\n", err)
		return user, err
	}
	user.Client = *resty.New()
	// assign the cookies to the client
	for _, rodCookie := range rodCookies {
		cookie := &http.Cookie{
			Name:     rodCookie.Name,
			Value:    rodCookie.Value,
			Path:     rodCookie.Path,
			Domain:   rodCookie.Domain,
			Expires:  time.Unix(int64(rodCookie.Expires), 0).UTC(),
			HttpOnly: rodCookie.HTTPOnly,
			Secure:   rodCookie.Secure,
		}
		user.Client.SetCookie(cookie)
	}
	user.Name = "Me"
	user.Email = ""
	user.UserID = ""
	user.listCert()
	user.getUserInfo()
	return user, nil
}

func (u *User) listCert() error {
	// GET: https://agc-drcn.developer.huawei.com/agc/edge/cps/harmony-cert-manage/v1/cert/list
	url := "https://agc-drcn.developer.huawei.com/agc/edge/cps/harmony-cert-manage/v1/cert/list"
	resp, err := u.Client.R().Get(url)
	if err != nil {
		return fmt.Errorf("Error while listing certs: %v", err)
	}
	fmt.Println("Response Status Code: ", resp.StatusCode())
	fmt.Println("Response Body: ", resp.String())
	return nil
}

func (u *User) getUserInfo() error {
	url := "https://agc-drcn.developer.huawei.com/agc/edge/apios/invokeService/AGCHomePageOrchestration/getUserInfo"
	resp, err := u.Client.R().Post(url)
	if err != nil {
		return fmt.Errorf("Error while getting user info: %v", err)
	}
	fmt.Println("Response Status Code: ", resp.StatusCode())
	fmt.Println("Response Body: ", resp.String())
	return nil
}
