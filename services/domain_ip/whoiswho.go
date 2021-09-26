package domain_ip

import (
	"fmt"
	. "github.com/k4s/phantomgo"
	"github.com/microcosm-cc/bluemonday"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// getWhois makes request to viewdns.info/whois/ service
// and parses response to extract whois information.
func getWhois(target string) ([]string, error) {
	var dnsInfo []string


	baseURL := "https://viewdns.info/whois/?domain="
	targetURL := baseURL + target

	// Set up http client and make request.
	httpClient := &http.Client{}
	req, _ := http.NewRequest("GET", targetURL, nil)
	req.Header.Set("Content-Type", "text/html; charset=UTF-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0")
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	re := regexp.MustCompile(`<br.*?>(.*)<br>`)
	submatchall := re.FindAllStringSubmatch(string(body), -1)

	for _, element := range submatchall {
		s := strings.Split(element[1], "<br>")
		for _, r := range s {
			if strings.Contains(r, "For more information on Whois ") {
				break
			} else {
				dnsInfo = append(dnsInfo, r)
			}
		}
	}

	return dnsInfo, nil
}

// getWhoisHistoryFree makes request to osint.sh/whoishistory/ free service
// and parses response to extract whois history information.
func GetWhoisHistoryFree (target string) []string {
	var records []string

	baseURL := "https://osint.sh/whoishistory/"
	postBody := "domain=" + target

	// osint.sh requires JS so here we use PhantomJS web browser to make request.
	p := &Param{
		Method:       "POST",
		Url:          baseURL,
		Header:       http.Header{"User-Agent": []string{`Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:92.0) Gecko/20100101 Firefox/92.0`}},
		UsePhantomJS: true,
		PostBody:     postBody,
	}
	browser := NewPhantom()
	resp, err := browser.Download(p)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	re := regexp.MustCompile(`<div class="col-lg-8">[\s\S]*</div>`)
	submatchall := re.FindAllStringSubmatch(string(body), -1)

	// We use StripTagsPolicy to cut off all HTML tags.
	stripPolicy := bluemonday.StripTagsPolicy()

	// Here we beautify what we've got from whois history
	// and place it into one slice.
	replacer := strings.NewReplacer("Update Date", "", "Expiry Date", "", "Create Date", "",
		"Owner", "", "Address", "", "Email", "", "Phone", "", "Name Server", "",
		)
	for n, element := range submatchall {
		s := strings.Split(element[n], "</div>")
		for _, r := range s {
			counter := 0
			sanitizedStrings := stripPolicy.Sanitize(r)
			cs := replacer.Replace(sanitizedStrings)
			cs = strings.TrimSpace(cs)
			ss := strings.Split(cs, "\n")
			for _, str := range ss {
				str = strings.TrimSpace(str) // TODO: fix this for cases when there is empty data in required lines (Update Date or Expire Date or so on)
				if (str == "") || (strings.Contains(str, "Crafted")) {
					continue
				} else {
					switch counter {
					case 0:
						str = "RECORD DATING FROM " + str
					case 1:
						str = "Update Date: " + str
					case 2:
						str = "Create Date: " + str
					case 3:
						str = "Expiry Date: " + str
					case 4:
						str = "Registrant: " + str
					case 5:
						str = "Registrant Address: " + str
					case 6:
						str = "Registrant Email: " + str
					case 7:
						str = "Registrant Phone: " + str
					case 8:
						str = "Name Server: " + str
					}
					records = append(records, str)
					counter++
				}
			}
		}
	}

	// Here we split the records slice into sub-slices
	// where each sub-slice is a record from whois history.
	size := 9
	var j int
	for i := 0; i < len(records); i += size {
		j += size
		if j > len(records) {
			j = len(records)
		}
		// do what do you want to with the sub-slice, here just printing the sub-slices
		fmt.Println(records[i:j])
	}

	return records

}

// TODO: make getWhoisHistoryPaid for paid service whoxy.com API
