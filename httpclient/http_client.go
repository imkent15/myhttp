package httpclient

import (
	"crypto/md5"
	"fmt"
	"myhttp/interfaces"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

var httpClient interfaces.IHTTPClient

func init() {
	httpClient = &http.Client{}
}

func ParallelFetch(parallelRequests int, addresses []string) {
	var wg sync.WaitGroup

	addressChannel := make(chan string, len(addresses))

	for i := 0; i < parallelRequests; i++ {
		wg.Add(1)
		go fetchAddress(addressChannel, &wg)
	}

	for _, address := range addresses {
		addressChannel <- address
	}

	close(addressChannel)
	wg.Wait()
}

func fetchAddress(addressChannel <-chan string, wg *sync.WaitGroup) {

	for address := range addressChannel {

		addressToRequest, err := validateAndAppendScheme(address)

		if err != nil {
			fmt.Printf("error : [%v]\n", err)
			continue
		}

		md5Checksum, err := getResponseMd5(addressToRequest)

		if err != nil {
			fmt.Printf("error : [%v]", err)
			continue
		} else {
			fmt.Printf("%s %x\n", addressToRequest, md5Checksum)
		}
	}

	wg.Done()

}

func validateAndAppendScheme(address string) (fixedAddress string, err error) {

	fixedAddress = address
	if !strings.HasPrefix(address, "http") {
		fixedAddress = "http://" + address
	}

	uri, err := url.ParseRequestURI(fixedAddress)

	if err != nil || uri.Host == "" {
		err = fmt.Errorf("invalid address [%s] error [%v]", address, err)
		return
	}

	return
}

func getResponseMd5(requestAddress string) (md5Checksum [16]byte, err error) {

	httpRequest, err := http.NewRequest(http.MethodGet, requestAddress, nil)
	if err != nil {
		err = fmt.Errorf("http request failed with error [%v]", err)
	}

	response, err := httpClient.Do(httpRequest)
	if err != nil {
		err = fmt.Errorf("http request failed with error [%v]", err)
		return
	}

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("http request failed with status code [%d]", response.StatusCode)
		return
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		err = fmt.Errorf("request failed with error [%v]", err)
	}

	md5Checksum = md5.Sum(responseBody)

	return
}
