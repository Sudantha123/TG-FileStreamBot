package main

import (
 "fmt"
 "log"
 "net/http"
 "os"
)

func main() {
 host := getenv("HOST", "0.0.0.0")
 port := getenv("PORT", "8080")
 mux := http.NewServeMux()
 mux.HandleFunc("/healthz", health)
 mux.HandleFunc("/dav", dav)
 mux.HandleFunc("/web", web)
 addr := fmt.Sprintf("%s:%s", host, port)
 log.Printf("TelegramDAV listening on %s", addr)
 log.Fatal(http.ListenAndServe(addr, mux))
}

func getenv(k, fallback string) string { if v := os.Getenv(k); v != "" { return v }; return fallback }
func health(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK); _, _ = w.Write([]byte("ok")) }
func dav(w http.ResponseWriter, r *http.Request) { w.Header().Set("Content-Type", "text/plain; charset=utf-8"); _, _ = w.Write([]byte("TelegramDAV WebDAV core endpoint")) }
func web(w http.ResponseWriter, r *http.Request) { w.Header().Set("Content-Type", "text/html; charset=utf-8"); _, _ = w.Write([]byte("<!doctype html><html><head><meta name=\"viewport\" content=\"width=device-width,initial-scale=1\"><title>TelegramDAV</title><style>body{font-family:system-ui;margin:0;padding:28px;background:#0b1020;color:#eef2ff}main{max-width:960px;margin:auto}.card{margin-top:20px;padding:20px;background:#151c32;border:1px solid #2b3656;border-radius:18px}.grid{display:grid;grid-template-columns:repeat(auto-fit,minmax(180px,1fr));gap:12px}.stat{padding:16px;background:#0f172a;border-radius:14px}.muted{color:#94a3b8}.btn{display:inline-block;padding:10px 14px;background:#2563eb;border-radius:10px;color:white;text-decoration:none}</style></head><body><main><h1>TelegramDAV</h1><p class=\"muted\">Telegram-backed WebDAV storage gateway</p><a class=\"btn\" href=\"/dav\">Open DAV</a><section class=\"card\"><h2>Pipeline</h2><div class=\"grid\"><div class=\"stat\">Telegram<br><span class=\"muted\">Storage backend</span></div><div class=\"stat\">HTTP Range<br><span class=\"muted\">Fast seeking</span></div><div class=\"stat\">Bounded cache<br><span class=\"muted\">Low VPS usage</span></div><div class=\"stat\">WebDAV<br><span class=\"muted\">Client compatible</span></div></div></section></main></body></html>")) }
