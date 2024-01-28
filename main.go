package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"opa-auth/models"

	"github.com/gin-gonic/gin"
	"github.com/open-policy-agent/opa/rego"
)

var data models.BundleData

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadBundleData() {
	dat, err := os.ReadFile("./bundle/data.json")
	check(err)

	err = json.Unmarshal(dat, &data)
	check(err)
}

func getBundleData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data)
}

func getBundlePolicy(c *gin.Context) {
	c.FileAttachment("./bundle/policy.rego", "policy.rego")
}

func eval(c *gin.Context) {
	var input map[string]interface{}

	err := c.BindJSON(&input)
	check(err)

	opaFiles := []string{
		"./bundle/policy.rego",
		"./bundle/data.json",
	}

	queryStr := "data.auth.allow"

	r := rego.New(
		rego.Query(queryStr),
		rego.Load(opaFiles, nil),
	)

	ctx := context.Background()

	query, err := r.PrepareForEval(ctx)
	check(err)

	rs, err := query.Eval(ctx, rego.EvalInput(input))
	check(err)

	c.IndentedJSON(http.StatusOK, rs[0].Expressions[0].String())
}

func main() {
	loadBundleData()

	router := gin.Default()
	router.GET("/bundle/data", getBundleData)
	router.GET("/bundle/policy", getBundlePolicy)
	router.POST("/opa/eval", eval)

	router.Run("localhost:8080")
}
