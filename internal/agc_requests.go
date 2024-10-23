// AppGallery Connect Requests
package internal

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

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
			// print out the cookies
			for _, cookie := range cookies {
				fmt.Printf("Cookie: %s = %s\n", cookie.Name, cookie.Value)
			}
			break
		}
		time.Sleep(1000 * time.Millisecond) // hang
	}

	return cookies, nil
}

func login() {
	// login_url := "https://developer.huawei.com/consumer/cn/service/josp/agc/index.html"
	login_url := "https://id1.cloud.huawei.com/CAS/portal/loginAuth.html"

	cookies, err := openRodBrowser(login_url)
	if err != nil {
		log.Fatalf("Error while opening browser: %v", err)
	}

	for _, cookie := range cookies {
		fmt.Printf("Cookie: %s = %s\n", cookie.Name, cookie.Value)
	}
}
