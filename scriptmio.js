const botonBuscarEmpleado = document.querySelector(".search button");
const Id = document.querySelector(".search input");
const Identificacion = document.querySelector(".identificacion");
const nombres = document.querySelector(".nombres");
const apellidos = document.querySelector(".apellidos");
const fecha_ingreso = document.querySelector(".fecha_ingreso");
const cargo = document.querySelector(".cargo");
const vinculado = document.querySelector(".es_vinculado");
const salario = document.querySelector(".salario");
const botonCrearEmpleado = document.querySelector(".crear button");
const botonEliminarEmpleado = document.querySelector(".eliminar button");
const botonModificarEmpleado = document.querySelector(".modificar button");

//const identificacionCol1= Stringdocument.querySelector("#identificacion placeholder");

botonBuscarEmpleado.addEventListener("click", async () => {
    try {
    console.log(Id.value)
    
    // Realiza la solicitud a la API principal
    const response = await axios.get(
      `http://localhost:8080/empleados/${Id.value}`
          );
    
    console.log(response.data.nombres);
    console.log(response.data.apellidos);
    console.log(response.data.identificacion);
    // leer datos de la BD
   Identificacion.innerHTML = response.data["identificacion"];
   nombres.innerHTML = response.data["nombres"];
   apellidos.innerHTML = response.data["apellidos"];
   fecha_ingreso.innerHTML = response.data["fecha_ingreso"];
   vinculado.innerHTML = response.data["es_vinculado"];
   cargo.innerHTML = response.data["cargo"];
   salario.innerHTML = response.data["salario"];
   //Number(document.querySelector("#identificacion" ).value)= response.data["identificacion"];
   //console.log(identificacionCol1.value+"este es"+response.data["identificacion"])
   //console.log(identificacionCol1)
   //document.getElementById('identificacion')[0].placeholder = "This is my new text";


  } catch (error) {
     // Manejo de errores
     console.error("Error al buscar el Id", error);   
     window.alert("Error al buscar empleado con Id:" + Id.value) 
  }
});
botonCrearEmpleado.addEventListener("click", async () => {
  try {
 const ident=Number(document.querySelector("#identificacion" ).value)
 const nomb=document.querySelector("#nombres" ).value
const  apelli=document.querySelector("#apellidos" ).value
 const fecha_in=document.querySelector("#fecha_ingreso" ).value
 const car=document.querySelector("#cargo" ).value
 const es_vin=Boolean(document.querySelector("#es_vinculado" ).value)
 const salar=String(document.querySelector("#salario" ).value)

 
 const empleado={identificacion: ident,
 nombres:`${nomb}`,
 apellidos:`${apelli}`,
 fecha_ingreso:`${fecha_in}`,
 cargo:`${car}`,
 es_vinculado:es_vin,
 salario:`${salar}`}
 console.log (empleado)
  
  // Realiza la solicitud a la API principal
 const response = await axios.post(
    `http://localhost:8080/empleados`,empleado
        );
        window.alert("se creó empleado") 
  

  console.log(typeof(ident),typeof(nomb),typeof(apelli),typeof(fecha_in),typeof(car),typeof(es_vin),typeof(salar)
  )



 

} catch (error) {
   // Manejo de errores
   console.error("Error al crear empleado", error);   
   window.alert("Error  al crear empleado") 
}
});
botonEliminarEmpleado.addEventListener("click", async () => {
  try {
    console.log(Id.value)
  
  
  // Realiza la solicitud a la API principal
 const response = await axios.delete(
    `http://localhost:8080/empleados/${Id.value}`
        );
        console.log("si paso")
        window.alert("se eliminó empleado con Id: " +Id.value) 
  
  } catch (error) {
   // Manejo de errores
   console.error("Error al eliminar empleado", error);   
   window.alert("Error  al eliminar empleado") 
}
});

botonModificarEmpleado.addEventListener("click", async () => {
  try {
/*const ident=Number(document.querySelector("#identificacion" ).value)
const nomb=document.querySelector("#nombres" ).value
const  apelli=document.querySelector("#apellidos" ).value
 const fecha_in=document.querySelector("#fecha_ingreso" ).value
 const car=document.querySelector("#cargo" ).value
 const es_vin=Boolean(document.querySelector("#es_vinculado" ).value)
 const salar=String(document.querySelector("#salario" ).value)*/
const ident=1234
const nomb="fercho"
const apelli="carro"
const fecha_in="12-12-12"
const car="gerente"
const es_vin=true
const salar="454545"
 
 const empleado={identificacion: ident,
 nombres:`${nomb}`,
 apellidos:`${apelli}`,
 fecha_ingreso:`${fecha_in}`,
 cargo:`${car}`,
 es_vinculado:es_vin,
 salario:`${salar}`}
 console.log (empleado)
    console.log(Id.value)
  
  
  // Realiza la solicitud a la API principal
 const response = await axios.patch(
    `http://localhost:8080/empleados/${Id.value},`,empleado
        );
        console.log("si paso")
        window.alert("se modificó empleado con Id: " +Id.value) 
  
  } catch (error) {
   // Manejo de errores
   console.error("Error al modificar  empleado", error);   
   window.alert("Error  al modificar empleado") 
}
});