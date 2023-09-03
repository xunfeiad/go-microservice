package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
)

// 注册服务
func RegisterService(r Registration) error {
	serviceUpdateURL, err := url.Parse(r.ServiceUpdateURL)
	if err != nil {
		return err
	}
	http.Handle(serviceUpdateURL.Path, &serviceUpdateHandler{})
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err = enc.Encode(r)
	if err != nil {
		log.Println(err)
		return err
	}
	resp, err := http.Post(ServerURL, "application/json", buf)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service. Registry service responded with code: %v", resp.StatusCode)
	}
	return nil
}

// 前置服务启动

type serviceUpdateHandler struct{}

func (suh serviceUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	dec := json.NewDecoder(r.Body)
	var p patch
	err := dec.Decode(&p)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("Updated received %v\n", p)
	prov.Update(p)
}

// 取消服务
func ShutdownService(url string) error {
	req, err := http.NewRequest(http.MethodDelete, ServerURL, bytes.NewBuffer([]byte(url)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "text/plain")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service. Registry service responded with code: %v", res.StatusCode)
	}
	return nil
}

type providers struct {
	serivices map[ServiceName][]string
	mutex     *sync.RWMutex
}

var prov = providers{
	serivices: make(map[ServiceName][]string),
	mutex:     new(sync.RWMutex),
}

func (p *providers) Update(pat patch) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, patchEntry := range pat.Added {
		if _, ok := p.serivices[patchEntry.Name]; !ok {
			p.serivices[patchEntry.Name] = make([]string, 0)
		}
		p.serivices[patchEntry.Name] = append(p.serivices[patchEntry.Name], patchEntry.URL)
	}
	for _, patchEntry := range pat.Removed {
		if providerURLs, ok := p.serivices[patchEntry.Name]; ok {
			for i := range providerURLs {
				if providerURLs[i] == patchEntry.URL {
					p.serivices[patchEntry.Name] = append(providerURLs[:i], providerURLs[i+1:]...)
				}
			}

		}
	}
}

func (p providers) get(name ServiceName) (string, error) {
	providers, ok := p.serivices[name]
	if !ok {
		return "", fmt.Errorf("no providers avaliable for service %v", name)
	}
	idx := int(rand.Float32() * float32(len(providers)))
	return providers[idx], nil
}

func GetProvider(name ServiceName) (string, error) {
	return prov.get(name)
}
