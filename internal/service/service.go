package service

import (
	"compress/gzip"
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/IgorAleksandroff/agent-status/internal/api"
	"github.com/go-chi/chi"
	"google.golang.org/grpc"

	"github.com/IgorAleksandroff/agent-status/internal/config"
	"github.com/IgorAleksandroff/agent-status/internal/usecase"
)

const (
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 5 * time.Second
	defaultShutdownTimeout = 3 * time.Second
)

type Server interface {
	Run(ctx context.Context)
}

type server struct {
	serverHTTP   *http.Server
	serverGRPC   *grpc.Server
	gRPCListener net.Listener
}

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func New(cfg config.ServerConfig, auth usecase.Authorization) (*server, error) {
	// init HTTP server
	r := chi.NewRouter()

	r.Use(gzipUnzip)
	r.Use(gzipHandle)

	restHandler := api.New(auth)

	restHandler.Register(r, http.MethodPost, "/api/user/register", restHandler.HandleUserRegister)
	restHandler.Register(r, http.MethodPost, "/api/user/login", restHandler.HandleUserLogin)

	r.Group(func(r chi.Router) {
		r.Use(restHandler.UserIdentity)
	})

	// init GRPC server

	return &server{
		serverHTTP: &http.Server{
			Handler:      r,
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
			Addr:         cfg.Host,
		},
	}, nil
}

func (s server) Run(ctx context.Context) {
	notifyHTTP := make(chan error, 1)
	go func() {
		notifyHTTP <- s.serverHTTP.ListenAndServe()
		close(notifyHTTP)
	}()

	notifyGRPC := make(chan error, 1)
	go func() {
		notifyGRPC <- s.serverGRPC.Serve(s.gRPCListener)
		close(notifyGRPC)
	}()

	select {
	case <-ctx.Done():
		log.Println("server interrupted by", ctx.Err())

		s.serverGRPC.GracefulStop()
		s.shutdownHTTP()
	case err := <-notifyHTTP:
		log.Printf("HTTP server stopped: %s", err)

		s.serverGRPC.GracefulStop()
	case err := <-notifyGRPC:
		log.Printf("gRPC server stopped: %s", err)

		s.shutdownHTTP()
	}
}

func (s server) shutdownHTTP() {
	ctxShutdown, cancel := context.WithTimeout(context.Background(), defaultShutdownTimeout)
	defer cancel()

	err := s.serverHTTP.Shutdown(ctxShutdown)
	if err != nil {
		log.Printf("error shutdown http server: %s", err)
	}
}

func (w gzipWriter) Write(b []byte) (int, error) {
	// w.Writer будет отвечать за gzip-сжатие, поэтому пишем в него
	return w.Writer.Write(b)
}

func gzipHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// проверяем, что клиент поддерживает gzip-сжатие
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// если gzip не поддерживается, передаём управление
			// дальше без изменений
			next.ServeHTTP(w, r)
			return
		}

		// создаём gzip.Writer поверх текущего w
		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		// передаём обработчику страницы переменную типа gzipWriter для вывода данных
		next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
	})
}

func gzipUnzip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// проверяем, что клиент поддерживает gzip-сжатие
		if r.Header.Get(`Content-Encoding`) != `gzip` {
			// если не сжато методом gzip, передаём управление
			// дальше без изменений
			next.ServeHTTP(w, r)
			return
		}

		// создаём *gzip.Reader, который будет читать тело запроса
		// и распаковывать его
		gz, err := gzip.NewReader(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// не забывайте потом закрыть *gzip.Reader
		defer gz.Close()

		// меняем Body на тип gzip.Reader для распаковки данных
		r.Body = gz

		next.ServeHTTP(w, r)
	})
}
