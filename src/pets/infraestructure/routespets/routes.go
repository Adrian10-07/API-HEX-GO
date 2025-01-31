package routes

import (
	"net/http"
	"API-HEX-GO/src/pets/infraestructure/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/pets", controllers.CreatePetHandler)            // Crear mascota
	http.HandleFunc("/view-pets", controllers.GetPetHandler)          // Obtener todas las mascotas
	http.HandleFunc("/delete-pets/", controllers.DeletePetHandler)    // Eliminar mascota
	http.HandleFunc("/update-pets/", controllers.UpdatePetHandler)    // Actualizar mascota
}
