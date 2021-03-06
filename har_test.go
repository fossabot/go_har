package har

import (
	"testing"
)

func TestMain(t *testing.T) {
	har, _ := parseHar("examples/FireFox.har")
	if har.Version != "1.2" {
		t.Errorf("Invalid version. Wanted version 1.2 got %s", har.Version)
	}
	if har.Browser.Name != "Firefox" && har.Browser.Name == har.Creator.Name {
		t.Errorf("Invalid browser name. Wanted FireFox got %s", har.Browser.Name)
	}
	if har.Browser.Version != "79.0" && har.Browser.Version == har.Creator.Version {
		t.Errorf("Invalid browser version. Wanted FireFox got %s", har.Browser.Version)
	}
}

func TestEntry(t *testing.T) {
	har, _ := parseHar("examples/FireFox.har")
	entry := har.Entries[0]
	if entry.IP != "104.27.153.17" {
		t.Errorf("Expected IP of '104.27.153.17'. Got %s", entry.IP)
	}
	if entry.PageID != "page_1" {
		t.Errorf("Expected PageRef 'page_1'. Got %s", entry.PageID)
	}
	if entry.Port != "443" {
		t.Errorf("Expected Port 443. Got %s", entry.Port)
	}
	if entry.Secure != "secure" {
		t.Errorf("Expected secure status of 'secure'. Got %s", entry.Secure)
	}
	if entry.StartedTime != "2020-08-22T10:04:37.230-04:00" {
		t.Errorf("Expected a StartedTime of '2020-08-22T10:04:37.230-04:00'. Got %s", entry.StartedTime)
	}
	if entry.Time != 51 {
		t.Errorf("Expected entry Time of 51. Got %d", entry.Time)
	}
}

func TestTiming(t *testing.T) {
	har, _ := parseHar("examples/FireFox.har")
	timings := har.Entries[0].Timing
	if timings.Blocked != 1 {
		t.Errorf("Expected Timing blocked of 1. Got %d", timings.Blocked)
	}
	if timings.DNS != 0 {
		t.Errorf("Expected Timing DNS of 0. Got %d", timings.DNS)
	}
	if timings.Connect != 0 {
		t.Errorf("Expected Timing connect of 0. Got %d", timings.Connect)
	}
	if timings.SSL != 0 {
		t.Errorf("Expected Timing SSL of 0. Got %d", timings.SSL)
	}
	if timings.Send != 0 {
		t.Errorf("Expected Timing Send of 0. Got %d", timings.Send)
	}
	if timings.Wait != 50 {
		t.Errorf("Expected Timing Wait of 1. Got %d", timings.Wait)
	}
	if timings.Receive != 0 {
		t.Errorf("Expected Timing Receive of 1. Got %d", timings.Receive)
	}

}

func TestRequest(t *testing.T) {
	har, _ := parseHar("examples/FireFox.har")
	request := har.Entries[0].Request

	if request.URL != "https://www.jwhite.network/" {
		t.Errorf("Invalid URL. Wanted 'https://www.jwhite.network/' got %v", request.URL)
	}
	if len(request.Cookies) != 2 {
		t.Errorf("There should be 2 request cookies. Got %d", len(request.Cookies))
	}
	if len(request.Headers) != 12 {
		t.Errorf("There should be 12 request headers. Got %d", len(request.Headers))
	}
	if request.Method != "GET" {
		t.Errorf("The request method should be 'GET'. Got %s", request.Method)
	}
	if request.BodySize != 0 {
		t.Errorf("Expected a request body size of 0. Got %d", request.BodySize)
	}
	if request.HTTPVersion != "HTTP/2" {
		t.Errorf("Expected a httpVersion of 'HTTP/2. Got %s", request.HTTPVersion)
	}
	if request.HeaderSize != 581 {
		t.Errorf("Expected Request Headersize of '581'. Got %d", request.HeaderSize)
	}
	if len(request.QueryString) != 0 {
		t.Errorf("Expected Querystring length of 0. Got %d", len(request.QueryString))
	}
}

func TestResponse(t *testing.T) {
	har, _ := parseHar("examples/FireFox.har")
	response := har.Entries[0].Response

	if response.Status != 200 {
		t.Errorf("Expected Response 200. Got %d", response.Status)
	}
	if response.StatusText != "OK" {
		t.Errorf("Expected Response Text 'OK'. Got %s", response.StatusText)
	}
	if response.HTTPVersion != "HTTP/2" {
		t.Errorf("Expected a httpVersion of 'HTTP/2. Got %s", response.HTTPVersion)
	}
	if len(response.Cookies) != 0 {
		t.Errorf("There should be 0 response cookies. Got %d", len(response.Cookies))
	}
	if response.HeaderSize != 2118 {
		t.Errorf("Expected response Headersize of '2118'. Got %d", response.HeaderSize)
	}
	if response.BodySize != 7411 {
		t.Errorf("Expected a response body size of 0. Got %d", response.BodySize)
	}
	if response.RedirectURL != "" {
		t.Errorf("Expected no response redirectURL. Got %s", response.RedirectURL)
	}
	if response.Content.MIMEType != "text/html; charset=utf-8" {
		t.Errorf("Expect response Content MIMEType of 'text/html; charset=utf-8'. Got %s", response.Content.MIMEType)
	}
	if len(response.Content.Text) != 18501 {
		t.Errorf("Expect response Text length of 18501. Got %d", len(response.Content.Text))
	}
	if response.Content.Size != 18504 {
		t.Errorf("Expect response size of 18504. Got %d", response.Content.Size)
	}
}

func TestFailure(t *testing.T) {
	_, err := parseHar("FireFox.har")
	if err.Error() != "open FireFox.har: no such file or directory" {
		t.Errorf("Wanted unabled to open error got %s", err)
	}
}
