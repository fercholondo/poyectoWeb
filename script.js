const botonBuscarEmpleado = document.querySelector(".search button");
const Id = document.querySelector(".search input");

botonBuscarEmpleado.addEventListener("click", async () => {
    //try {
    console.log(Id.value)
    
    // Realiza la solicitud a la API principal
    /*const response = await axios.get(
      `http://localhost:8080/empleados/${Id.value}`
          );*/
          response={"id":2,
          "identificacion":1017123682,
          "nombres":"Carolina",
          "apellidos":"Gomez",
          "fecha_ingreso":"2010-02-07T00:00:00Z",
          "cargo":"Gerente",
          "es_vinculado":true,
          "salario":"5'000.000"}
    
    document.getElementsByClassName
    console.log (response.nombres)
    console.log (response.apellidos)
    console.log (response.fecha_ingreso)
    console.log (response.cargo)
    console.log (response.es_vinculado)
    console.log (response.salario)
    const padre =document.getElementsByClassName("datos")
    const indentification=padre.getElementById("identificacion")
    const name=padre.getElementById("nombres")
    const lastName=padre.getElementById("apellidos")
    const dateEntry=padre.getElementById("fecha_ingreso")
    const ocupation=padre.getElementById("cargo")
    const vinculated=padre.getElementById("es_vinculado")
    const salary=padre.getElementById("salario")

    indentification.innerHTML=identificacion
    console.log(identificacion)
    //console.log(response.data.nombres);
    //console.log(response.data.apellidos);
    //console.log(response.data.identificacion);
 /* } catch (error) {
     // Manejo de errores
     console.error("Error al buscar el Id", error);   
     window.alert("Error al buscar el Id") 
  }*/
});
