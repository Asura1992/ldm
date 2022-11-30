package srv

import (
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"go.uber.org/zap"
	"ldm/common/config"
	"ldm/utils/swagger"
	"net/http"
	"path"
	"strings"
)

//注册swagger
func initSwagger() {
	cfg := config.GlobalConfig
	if !cfg.Swagger.Enabled {
		return
	}
	mux := http.NewServeMux()
	mux.Handle("/", gateWayMux)
	mux.HandleFunc("/swagger/", swaggerFile)
	swaggerUI(mux)
	err := http.ListenAndServe(cfg.Swagger.SwaggerAddr, mux)
	if err != nil {
		zap.S().Error("failed to initSwagger:", err.Error())
	}
}

/**
swaggerFile: 提供对swagger.json文件的访问支持
*/
func swaggerFile(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		http.NotFound(w, r)
		return
	}
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	name := path.Join("common/swagger", p)
	http.ServeFile(w, r, name)
}

/**
serveSwaggerUI: 提供UI支持
*/
func swaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "common/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
