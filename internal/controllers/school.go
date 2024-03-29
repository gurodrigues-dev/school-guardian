package controllers

import (
	"gin/types"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ct *controller) CreateSchool(c *gin.Context) {

	var input types.School

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	err := ct.service.CreateSchool(c, &input)

	if err != nil {
		log.Printf("error to create school: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an error occurred when creating school"})
		return
	}

	c.JSON(http.StatusCreated, input)

}

func (ct *controller) ReadSchool(c *gin.Context) {

	cnpj, err := ct.service.ParserJwtSchool(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, "cnpj of cookie don't found")
		return
	}

	log.Print("consulting page of read school --> ", cnpj)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Printf("error to parse string: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	school, err := ct.service.ReadSchool(c, &id)

	if err != nil {
		log.Printf("error while found school: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "school don't found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"school": school})

}

func (ct *controller) ReadAllSchools(c *gin.Context) {

	cnpj, err := ct.service.ParserJwtSchool(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, "cnpj of cookie don't found")
		return
	}

	log.Print("consulting page -->", cnpj)

	schools, err := ct.service.ReadAllSchools(c)

	if err != nil {
		log.Printf("error while found schools: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "schools don't found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"schools": schools})

}

func (ct *controller) UpdateSchool(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "updated w successfully"})

}

func (ct *controller) DeleteSchool(c *gin.Context) {

	cnpjInterface, err := ct.service.ParserJwtSchool(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cnpj of cookie don't found"})
		return
	}

	cnpj, err := ct.service.InterfaceToString(cnpjInterface)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the value isn't string"})
		return
	}

	err = ct.service.DeleteSchool(c, cnpj)

	if err != nil {
		log.Printf("error whiling deleted school: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to deleted school"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)

	c.JSON(http.StatusOK, gin.H{"message": "school deleted w successfully"})

}

func (ct *controller) AuthSchool(c *gin.Context) {

	var input types.School

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"erro": "invalid body content"})
		return
	}

	school, err := ct.service.AuthSchool(c, &input)

	if err != nil {
		log.Printf("wrong email or password: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong email or password"})
		return
	}

	jwt, err := ct.service.CreateTokenJWTSchool(c, school)

	if err != nil {
		log.Printf("error to create jwt token: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to create jwt token"})
		return
	}

	c.SetCookie("token", jwt, 3600, "/", c.Request.Host, false, true)

	c.JSON(http.StatusAccepted, gin.H{
		"school": school,
		"token":  jwt,
	})

}

func (ct *controller) InviteNewDriver(c *gin.Context) {

}
