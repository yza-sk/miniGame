package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/zeromicro/go-zero/rest"
	"gopkg.in/yaml.v3"

	"miniGame/internal/config"
	"miniGame/internal/handler"
	"miniGame/internal/svc"
)

func main() {
	cfg, err := loadConfig("etc/config.yaml")
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	// If DSN not present in config (YAML mapping issues), fall back to sensible default
	if cfg.Database.DSN == "" {
		cfg.Database.DSN = "root:password@tcp(mysql:3306)/minigame?parseTime=true&loc=UTC"
		log.Printf("Database DSN not found in config, using default DSN: %s", cfg.Database.DSN)
	} else {
		log.Printf("Database DSN from config: %s", cfg.Database.DSN)
	}
	ctx, err := svc.NewServiceContext(cfg)
	if err != nil {
		log.Fatalf("service context error: %v", err)
	}

	// 输出 RestConf 便于排查监听地址/端口
	log.Printf("rest conf: %+v", cfg.RestConf)
	log.Printf("cors allowed origins: %+v", cfg.Cors.AllowedOrigins)

	// 如果 Cors 配置未载入，设置默认允许列表（包含开发和容器端口）
	if len(cfg.Cors.AllowedOrigins) == 0 {
		cfg.Cors.AllowedOrigins = []string{"http://localhost:5173", "http://localhost:8080"}
		log.Printf("cors allowed origins not found in config, using defaults: %+v", cfg.Cors.AllowedOrigins)
	}

	// 如果配置未正确映射 RestConf，确保使用默认监听地址和端口
	if cfg.RestConf.Host == "" {
		cfg.RestConf.Host = "0.0.0.0"
	}
	if cfg.RestConf.Port == 0 {
		cfg.RestConf.Port = 8888
	}

	server := rest.MustNewServer(cfg.RestConf)
	defer server.Stop()

	// CORS middleware (go-zero rest.Middleware signature)
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			allowed := false
			if origin != "" {
				allowed = allowOrigin(origin, cfg.Cors.AllowedOrigins)
			}
			// debug log for CORS
			log.Printf("CORS check: origin=%q allowed=%v path=%s method=%s", origin, allowed, r.URL.Path, r.Method)
			if allowed {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
				w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
				w.Header().Set("Vary", "Origin")
			}
			if r.Method == http.MethodOptions {
				// For preflight, ensure headers are present and return 204
				if allowed {
					w.WriteHeader(http.StatusNoContent)
					return
				}
				// If not allowed, return 403
				http.Error(w, "CORS origin not allowed", http.StatusForbidden)
				return
			}
			start := time.Now()
			next(w, r)
			_ = start
		}
	})

	server.AddRoutes([]rest.Route{
		{Method: http.MethodPost, Path: "/api/records/submit", Handler: handler.SubmitHandler(ctx)},
		{Method: http.MethodGet, Path: "/api/records/rank", Handler: handler.RankHandler(ctx)},
		{Method: http.MethodGet, Path: "/api/records/recent", Handler: handler.RecentHandler(ctx)},
	})

	// Ensure preflight OPTIONS are explicitly handled so framework won't return 405 without CORS headers
	setOptions := func(path string) rest.Route {
		return rest.Route{
			Method: http.MethodOptions,
			Path:   path,
			Handler: func(w http.ResponseWriter, r *http.Request) {
				origin := r.Header.Get("Origin")
				if origin != "" && allowOrigin(origin, cfg.Cors.AllowedOrigins) {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
					w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
					w.Header().Set("Vary", "Origin")
					w.WriteHeader(http.StatusNoContent)
					return
				}
				http.Error(w, "CORS origin not allowed", http.StatusForbidden)
			},
		}
	}
	server.AddRoutes([]rest.Route{
		setOptions("/api/records/submit"),
		setOptions("/api/records/rank"),
		setOptions("/api/records/recent"),
	})

	server.Start()
}

func allowOrigin(o string, list []string) bool {
	for _, v := range list {
		if v == o {
			return true
		}
	}
	return false
}

func loadConfig(path string) (config.Config, error) {
	var c config.Config
	b, err := os.ReadFile(path)
	if err != nil {
		return c, err
	}
	if err := yaml.Unmarshal(b, &c); err != nil {
		return c, err
	}
	return c, nil
}
