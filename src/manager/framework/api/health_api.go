package api

import (
	"net/http"

	"github.com/Dataman-Cloud/swan/src/manager/apiserver"
	"github.com/Dataman-Cloud/swan/src/manager/apiserver/metrics"
	"github.com/Dataman-Cloud/swan/src/manager/framework/scheduler"

	"github.com/emicklei/go-restful"
)

type HealthyService struct {
	Scheduler *scheduler.Scheduler
	apiserver.ApiRegister
}

func NewAndInstallHealthyService(apiServer *apiserver.ApiServer, eng *scheduler.Scheduler) *HealthyService {
	healthyService := &HealthyService{
		Scheduler: eng,
	}
	apiserver.Install(apiServer, healthyService)
	return healthyService
}

func (api *HealthyService) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		ApiVersion(API_PREFIX).
		Path("/ping").
		Doc("ping API").
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").To(metrics.InstrumentRouteFunc("GET", "Ping", api.Ping)).
		Doc("Ping").
		Operation("ping").
		Returns(200, "OK", ""))

	container.Add(ws)
}

func (api *HealthyService) Ping(request *restful.Request, response *restful.Response) {
	pong := "pong"
	response.WriteHeaderAndEntity(http.StatusOK, pong)
}
