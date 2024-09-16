package httphandler

import (
	"net/http"

	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/httphandler/controller"
	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	httpServer      *http.Server
	router          *gin.Engine
	eventController *controller.EventController
}

func NewHttpHandler(c *controller.EventController) *HttpHandler {
	r := gin.Default()
	srv := &http.Server{
		Addr:    "0.0.0.0" + ":" + "5515",
		Handler: r,
	}
	handler := &HttpHandler{
		router:          r,
		httpServer:      srv,
		eventController: c,
	}

	handler.RegisterRoutes()

	return handler
}

func (h *HttpHandler) RegisterRoutes() {
	h.router.GET("/:aggregateId/events", h.eventController.GetEventsForAggregate)
	h.router.POST("/:aggregateId/event", h.eventController.AddEventToAggregate)
}

func (h *HttpHandler) Start() error {
	return h.httpServer.ListenAndServe()
}
