package main

import (
	"course-chart/config"
	"course-chart/models"
	"course-chart/routes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGETModule(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.Connect()

	t.Run("returns a module", func(t *testing.T) {
		router := routes.GetRoutes()
		request, _ := http.NewRequest("GET", fmt.Sprintf("/modules/%d", 1), nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, 200, response.Code)

		body, _ := ioutil.ReadAll(response.Body)
		nestedModule := MarshaledModule{}
		err := json.Unmarshal([]byte(body), &nestedModule)
		if err != nil {
			t.Errorf("Error marshaling JSON response\nError: %v", err)
		}

		var module models.Module
		db := config.Conn
		db.Preload("ModuleActivities.Activity").First(&module, 1)

		if reflect.DeepEqual(nestedModule.Data, models.Module{}) {
			t.Errorf("response does not contain an id property")
		}

		assertResponseValue(t, nestedModule.Message, "Module found", "Response message")
		assertResponseValue(t, nestedModule.Data.ID, module.ID, "Id")
		assertResponseValue(t, nestedModule.Data.Name, module.Name, "Name")
		assertResponseValue(t, nestedModule.Data.Number, module.Number, "Number")
		assertResponseValue(t, nestedModule.Data.CourseId, module.CourseId, "CourseId")
		firstResponseModActivity := nestedModule.Data.ModuleActivities[0]
		firstDBModActivity := module.ModuleActivities[0]
		assertResponseValue(t, firstResponseModActivity.Input, firstDBModActivity.Input, "Module Activity Input")
		assertResponseValue(t, firstResponseModActivity.Notes, firstDBModActivity.Notes, "Module Activity Notes")
		firstResponseActivity := firstResponseModActivity.Activity
		firstDBActivity := firstDBModActivity.Activity
		assertResponseValue(t, firstResponseActivity.ID, firstDBActivity.ID, "Activity Id")
		assertResponseValue(t, firstResponseActivity.Description, firstDBActivity.Description, "Activity Description")
		assertResponseValue(t, firstResponseActivity.Metric, firstDBActivity.Metric, "Activity Metric")
		assertResponseValue(t, firstResponseActivity.Multiplier, firstDBActivity.Multiplier, "Activity Multiplier")
	})
}

type MarshaledModule struct {
	Data    models.Module
	Message string
	Status  int
}

// func assertResponseValue(t *testing.T, got, want interface{}, field string) {
// 	if got != want {
// 		t.Errorf("got %v want %v for field %q", got, want, field)
// 	}
// }
