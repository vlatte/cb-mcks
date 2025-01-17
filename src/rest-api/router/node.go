package router

import (
	"net/http"
	"time"

	"github.com/cloud-barista/cb-mcks/src/core/model"
	"github.com/cloud-barista/cb-mcks/src/core/service"
	"github.com/cloud-barista/cb-mcks/src/utils/app"

	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
)

// ListNode godoc
// @Tags Node
// @Summary List all Nodes in specified Cluster
// @Description List all Nodes in specified Cluster
// @ID ListNode
// @Accept json
// @Produce json
// @Param	namespace	path	string	true  "Namespace ID"
// @Param	cluster	path	string	true  "Cluster Name"
// @Success 200 {object} model.NodeList
// @Failure 400 {object} app.Status
// @Router /ns/{namespace}/clusters/{cluster}/nodes [get]
func ListNode(c echo.Context) error {
	if err := app.Validate(c, []string{"cluster"}); err != nil {
		logger.Error(err)
		return app.SendMessage(c, http.StatusBadRequest, err.Error())
	}

	nodeList, err := service.ListNode(c.Param("namespace"), c.Param("cluster"))
	if err != nil {
		logger.Error(err)
		return app.SendMessage(c, http.StatusBadRequest, err.Error())
	}

	return app.Send(c, http.StatusOK, nodeList)
}

// GetNode godoc
// @Tags Node
// @Summary Get Node in specified Cluster
// @Description Get Node in specified Cluster
// @ID GetNode
// @Accept json
// @Produce json
// @Param	namespace	path	string	true  "Namespace ID"
// @Param	cluster	path	string	true  "Cluster Name"
// @Param	node	path	string	true  "Node Name"
// @Success 200 {object} model.Node
// @Failure 400 {object} app.Status
// @Failure 404 {object} app.Status
// @Router /ns/{namespace}/clusters/{cluster}/nodes/{node} [get]
func GetNode(c echo.Context) error {
	if err := app.Validate(c, []string{"cluster", "node"}); err != nil {
		logger.Error(err)
		return app.SendMessage(c, http.StatusBadRequest, err.Error())
	}

	node, err := service.GetNode(c.Param("namespace"), c.Param("cluster"), c.Param("node"))
	if err != nil {
		logger.Infof("not found a node (namespace=%s, cluster=%s, node=%s, cause=%s)", c.Param("namespace"), c.Param("cluster"), c.Param("node"), err)
		return app.SendMessage(c, http.StatusNotFound, err.Error())
	}

	return app.Send(c, http.StatusOK, node)
}

// AddNode godoc
// @Tags Node
// @Summary Add Node in specified Cluster
// @Description Add Node in specified Cluster
// @ID AddNode
// @Accept json
// @Produce json
// @Param	namespace	path	string	true  "Namespace ID"
// @Param	cluster	path	string	true  "Cluster Name"
// @Param nodeReq body model.NodeReq true "Request Body to add node"
// @Success 200 {object} model.Node
// @Failure 400 {object} app.Status
// @Failure 500 {object} app.Status
// @Router /ns/{namespace}/clusters/{cluster}/nodes [post]
func AddNode(c echo.Context) error {
	start := time.Now()
	if err := app.Validate(c, []string{"cluster"}); err != nil {
		logger.Error(err)
		return app.SendMessage(c, http.StatusBadRequest, err.Error())
	}

	nodeReq := &model.NodeReq{}
	if err := c.Bind(nodeReq); err != nil {
		logger.Error(err)
		return app.SendMessage(c, http.StatusBadRequest, err.Error())
	}

	err := app.NodeReqValidate(*nodeReq)
	if err != nil {
		logger.Error(err)
		return app.SendMessage(c, http.StatusBadRequest, err.Error())
	}

	node, err := service.AddNode(c.Param("namespace"), c.Param("cluster"), nodeReq)
	if err != nil {
		logger.Error(err)
		return app.SendMessage(c, http.StatusInternalServerError, err.Error())
	}

	duration := time.Since(start)
	logger.Info(" duration := ", duration)
	return app.Send(c, http.StatusOK, node)
}

// RemoveNode godoc
// @Tags Node
// @Summary Remove Node in specified Cluster
// @Description Remove Node in specified Cluster
// @ID RemoveNode
// @Accept json
// @Produce json
// @Param	namespace	path	string	true  "Namespace ID"
// @Param	cluster	path	string	true  "Cluster Name"
// @Param	node	path	string	true  "Node Name"
// @Success 200 {object} model.Status
// @Failure 400 {object} app.Status
// @Failure 500 {object} app.Status
// @Router /ns/{namespace}/clusters/{cluster}/nodes/{node} [delete]
func RemoveNode(c echo.Context) error {
	if err := app.Validate(c, []string{"cluster", "node"}); err != nil {
		logger.Error(err)
		return app.SendMessage(c, http.StatusBadRequest, err.Error())
	}

	status, err := service.RemoveNode(c.Param("namespace"), c.Param("cluster"), c.Param("node"))
	if err != nil {
		logger.Error(err)
		return app.SendMessage(c, http.StatusInternalServerError, err.Error())
	}

	return app.Send(c, http.StatusOK, status)
}
