package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"time"
)

var (
	// 定义一个计数器，用于统计HTTP请求总数
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)

	// 定义一个直方图，用于统计HTTP请求的处理时间
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)
)

func init() {
	// 注册指标
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
}

// 中间件函数，用于记录请求的处理时间和请求总数
func prometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()

		// 记录请求总数和处理时间
		httpRequestsTotal.WithLabelValues(c.Request.URL.Path).Inc()
		httpRequestDuration.WithLabelValues(c.Request.URL.Path).Observe(duration)
	}
}

func main() {
	// 创建一个Gin引擎
	r := gin.Default()

	// 使用中间件
	r.Use(prometheusMiddleware())

	// 定义一个简单的处理函数
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, Prometheus!")
	})

	// 暴露指标端点
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	log.Println("Starting server on :8080")
	log.Fatal(r.Run(":8080"))
}
