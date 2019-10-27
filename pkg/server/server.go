// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package server

import (
	"fmt"
	"net"
	"net/http"

	"sewadong-service/pkg/api"
	"sewadong-service/pkg/config"
	"sewadong-service/pkg/service"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	"github.com/sirupsen/logrus"
)

// Server handles starting REST API server
type Server struct {
	config    *config.Config
	listener  net.Listener
	container *restful.Container
}

// New new server instance
func New(config *config.Config, service *service.Service) *Server {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		logrus.Error("unable to create HTTP listener: ", err)
		return nil
	}

	server := &Server{
		config:    config,
		listener:  listener,
		container: restful.NewContainer(),
	}

	rootService := new(restful.WebService)
	server.container.Add(api.AddRouteUser(service, config.BasePath))
	server.container.Add(api.AddRouteRole(service, config.BasePath))
	server.container.Add(api.AddRouteProduct(service, config.BasePath))
	server.container.Add(api.AddRouteCategory(service, config.BasePath))
	server.registerSwaggerUI()

	cors := server.addCORS()
	server.container.Filter(cors.Filter)
	server.container.Add(rootService)

	return server
}

// Stop stops HTTP listener (REST API server)
func (server *Server) Stop() {
	logrus.Debug("stopping server")
	err := server.listener.Close()
	if err != nil {
		logrus.Error("unable to close HTTP listener: ", err)
	}
}

// Serve start REST API server
func (server *Server) Serve() {
	logrus.Info("Staring server...")
	err := http.Serve(server.listener, server.container)
	if err != nil {
		logrus.Error("unable to serve HTTP: ", err)
	}
}

func (server *Server) registerSwaggerUI() {
	conf := restfulspec.Config{
		WebServices: server.container.RegisteredWebServices(),
		APIPath:     server.config.BasePath + "/api.json",
		PostBuildSwaggerObjectHandler: func(swo *spec.Swagger) {
			swo.Info = &spec.Info{
				InfoProps: spec.InfoProps{
					Title:       "sewadong-service",
					Description: "sewadong service",
					Contact: &spec.ContactInfo{
						Name: "Sewadong",
					},
				},
			}
			swo.Tags = []spec.Tag{
				{TagProps: spec.TagProps{
					Name:        "sewadong-service",
					Description: ""}}}
		},
	}

	server.container.ServeMux.Handle(server.config.BasePath+"/apidocs/",
		http.StripPrefix(server.config.BasePath+"/apidocs/",
			http.FileServer(http.Dir("/contents/swagger-ui"))))
	server.container.Add(restfulspec.NewOpenAPIService(conf))
}

func (server *Server) addCORS() restful.CrossOriginResourceSharing {
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-Forwarded-For"},
		AllowedHeaders: []string{"Content-Type"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		Container:      server.container,
	}

	return cors
}

// StartHTTPTestServer starts the HTTP test server
func (server *Server) StartHTTPTestServer(responseRecorder http.ResponseWriter, req *http.Request) {
	server.container.ServeHTTP(responseRecorder, req)
}
