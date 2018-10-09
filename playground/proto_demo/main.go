package main

import (
	"flag"
	"net/http"
	"path"
	"strings"

	"github.com/golang/glog"

	gw "proto_demo/demo"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	grpcPort = "50051"
)

var (
	getEndpoint  = flag.String("get", "localhost:"+grpcPort, "endpoint of YourService")
	postEndpoint = flag.String("post", "localhost:"+grpcPort, "endpoint of YourService")

	swaggerDir = flag.String("swagger_dir", "template", "path to the directory which contains swagger definitions")
)

func newGateway(ctx context.Context, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterHelloWorldHandlerFromEndpoint(ctx, mux, *getEndpoint, dialOpts)

	if err != nil {
		return nil, err
	}

	err = gw.RegisterHelloWorldHandlerFromEndpoint(ctx, mux, *postEndpoint, dialOpts)

	if err != nil {
		return nil, err
	}

	return mux, err
}

func serverSwagger(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
		glog.Errorf("Swagger JSON not Found:%s", r.URL.Path)
		http.NotFound(w, r)
		return
	}
	glog.Infof("Serving %s", r.URL.Path)
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(*swaggerDir, p)
	http.ServeFile(w, r, p)
}

func Run(address string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	gw, _ := newGateway(ctx, opts...)
	mux.Handle("/", gw)
	return http.ListenAndServe(address, mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()
	if err := Run(":8080"); err != nil {
		glog.Fatal(err)
	}
}
