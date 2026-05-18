package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TaskAPI interface {
	AddTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	GetTaskByID(c *gin.Context)
	GetTaskList(c *gin.Context)
	GetTaskListByCategory(c *gin.Context)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskRepo service.TaskService) *taskAPI {
	return &taskAPI{taskRepo}
}

func (t *taskAPI) AddTask(c *gin.Context) {
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := t.taskService.Store(&newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add task success"})
}

func (t *taskAPI) UpdateTask(c *gin.Context) {
	
	session, err := c.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{Error: "Unauthorized"})
		return
	}
	fmt.Println("Session Token:", session)

   IntID, err := strconv.Atoi(c.Param("id"))
   if err != nil {
	   c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
	   return
   }

   
   var updatedTask model.Task
   if err := c.ShouldBindJSON(&updatedTask); err != nil {
	   c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	   return
   }

  
   updatedTask.ID = IntID

   
   err = t.taskService.Update(IntID, &updatedTask)
   if err != nil {
	   c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
	   return
   }

  
   c.JSON(http.StatusOK, model.SuccessResponse{Message: "update task success"})
	
}

func (t *taskAPI) DeleteTask(c *gin.Context) {
	
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	
	err = t.taskService.Delete(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "delete task success"})
	
}

func (t *taskAPI) GetTaskByID(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	task, err := t.taskService.GetByID(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t *taskAPI) GetTaskList(c *gin.Context) {
	tasks, err := t.taskService.GetList()
	if err != nil {
	
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	
	c.JSON(http.StatusOK, tasks)
	
}

func (t *taskAPI) GetTaskListByCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid category ID"})
		return
	}

	tasks, err := t.taskService.GetTaskCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	
	if len(tasks) == 0 {
		c.JSON(http.StatusOK, []model.Task{})
		return
	}

	c.JSON(http.StatusOK, tasks)
	
}
