package main

import(
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)
type Prox struct {
	// port the the proxy will run on
	port	string

	// base target url of the reverse proxy
	tBase	string

	// api endpoint that will be pulled from the request object and added to the baseTarget string
	api	string

	// target url of the reverse proxy
	target	*url.URL

	// the instance of the reverse proxy that will make the request
	proxy	*httputil.ReverseProxy
}


func main() {
	p := new(Prox)
	p.GetPort()
	p.GetTargetBase()
	// p.proxy := httputil.NewSingleHostReverseProxy()

	http.HandleFunc("/", p.ProxHandler)
	err := http.ListenAndServe(p.port, nil)
	errCheck(err)

}

func (p *Prox)ProxHandler() func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		p.GetApi(req)
		p.BuildTarget()
		*p.proxy = httputil.NewSingleHostReverseProxy(p.target)

		p.proxy.ServeHTTP(w, req)
		// p.proxy.ServeTLS(w, req)

	}
}

// func getPort() (string) {
// 	fmt.Println("Enter the localhost port on which HotProx should listen :  ")
// 	var port = os.Args[0]
// 	return port
// }

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// checks to make sure that user inputs a valid target base
func regexCheck(s string) (string, error){

	if s == s {

		return s, nil
	} else {
		return nil, "Something is wrong with the regex"
	}
}

func (p *Prox)GetPort() {
	var err error
	fmt.Println("Enter the localhost port on which HotProx should listen :  ")
	arg := os.Args[0]
	p.port, err = regexCheck(arg)
	errCheck(err)
}

func (p *Prox)GetTargetBase() {
	var err error
	fmt.Println("Example: https://www.google.com")
	fmt.Println("Enter the base url (leaving out any api endpoints or end slashes")
	fmt.Println("where your request should be forwarded :  ")
	arg := os.Args[0]
	p.tBase, err = regexCheck(arg)
	errCheck(err)
}

func (p *Prox)GetApi(req *http.Request) {
	reqUrl := io.WriteString(*req.URL)
	// api := bytes.Replace([]byte(reqUrl), []byte(p.tBase), []byte(""))
	// should assign this api variable to the *p.api pointer if this Replace works correctly
	p.api = bytes.Replace([]byte(reqUrl), []byte(p.tBase), []byte(""))

	// incase you need to map querystring variables in the future
	// queryString := *req.URL.Query()
}

func (p *Prox)BuildTarget() {
	var err error
	// target := fmt.Sprintf("%x%x", p.tBase, p.api)
	target := fmt.Sprint(p.tBase, p.api)
	*p.target, err = url.Parse(target)
	errCheck(err)
}



// func (req *http.Request)RunProxy(writer http.ResponseWriter) {
//
// 	req.URL = http.ProxyURL(*target)
//
//
// }

// func main() {
//
// 	// Sets which port to listen for on the localhost
// 	port := []byte{":"}
// 	port = append(port, []byte(os.Args[0]))
//
// 	target = &http.Request.URL(os.Args[1])
// 	// targetUrl := "localhost:2482/api/GroceryList"
//
// //	port := string(":") + os.Args[0]
//
// 	prox := &http.Server{
// 		Addr:	port,
// 		Handler:	ProxHandler,
// 		ReadTimeout:	10 * time.Second,
// 		WriteTimeout:	10 * time.Second,
// 		MaxHeaderBytes:	1 << 20,
// 	}
// 	log.Fatal(s.ListenAndServeTLS()
//
// }
