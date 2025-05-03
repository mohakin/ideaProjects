package main

import (
	"example.com/REST_API/db"
	"example.com/REST_API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")

	// localhost :8080
}

// What is a REST API? Think of it as a website that is not serving HTML code but instead data -
// - that can then be used by any client that needs that data.
// ------------------------Project Description-----------------------------
// GET /events --> Get a list of available events          <----------  X -
// GET /events/<id> --> Get a list of available events     <----------  X -
// POST /events --> Create a new bookable event            <----------  X -
// PUT /events/<id> --> Update an event          <--------------------  X -
// DELETE /events/<id> --> Delete an event       <--------------------  X -
// POST /signup --> Create a new user            <--------------------  X -
// POST /login --> Authenticate user             <--------------------   -
// POST /events/<id>/register --> Register user for event   <---------   -
// DELETE /events/<id>/register --> Cancel registration     <---------   -
// -----------------------------------------------------------------------
