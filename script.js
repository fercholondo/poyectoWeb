const botonBuscarEmpleado = document.querySelector(".search button");
const Id = document.querySelector(".search input");

botonBuscarEmpleado.addEventListener("click", async () => {
    try {
    console.log(Id.value)
    
    // Realiza la solicitud a la API principal
    const response = await axios.get(
      `https://localhost:8080/empleados/${Id}`
          );
    

    console.log(response.data.nombres);
    console.log(response.data.apellidos);
    console.log(response.data.identificacion);
  } catch (error) {
     // Manejo de errores
     console.error("Error al buscar el Id", error);   
     window.alert("Error al buscar el Id") 
  }
});
