package main

import (
	"flag"
	"net/http"
	"path"
	"strings"

	"github.com/golang/glog"

	"google.golang.org/grpc"
	
	gw "demo/demo_proto"
	
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
)

const (
	grpcPort = "10000"
)

var (
	getEndPoint = flag.String("get", "localhost:"+grpcPort, "endPoint of your service")
	postPoint   = flag.String("post", "localhost:"+grpcPort, "endpoint of you service")

	swaggerDir = flag.String("swagger_dir", "demo_proto", "含有swagger json文件的路径")
)

func newGateway(ctx context.Context, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterHelloHandlerFromEndpoint(ctx, mux, *getEndPoint, dialOpts)

	if err != nil {
		return nil, err
	}

	err = gw.RegisterHelloHandlerFromEndpoint(ctx, mux, *postPoint, dialOpts)
	if err != nil {
		return nil, err
	}
	return mux, err
}

func Run(address string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	mux := http.NewServeMux()

	mux.HandleFunc("/swagger/", serveSwagger)
	// serverSwaggerUI(mux)
	gw, err := newGateway(ctx, opts...)
	if err != nil {
		return err
	}

	mux.Handle("/", gw)
	return http.ListenAndServe(address, mux)
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
		glog.Errorf("Swagger JSON not found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	glog.Infof("Serving %s", r.URL.Path)
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(*swaggerDir, p)
	http.ServeFile(w, r, p)
}

// func serverSwaggerUI(mux *http.ServeMux) {
// 	fileServer := http.FileServer(&assetfs.AssetFS{
// 		Asset: swagger.Asset,
// 		AssetDir: swagger.AssetDir,
// 		Prefix: "third_party/swagger-ui",
// 	})

// 	prefix :="/swagger-ui/"
// 	mux.Handle(prefix,http.StripPrefix(prefix,fileServer))
// }

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := Run(":8080"); err != nil {
		glog.Fatal(err)
	}
}
