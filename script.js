const botonBuscarEmpleado = document.querySelector(".search button");
const Id = document.querySelector(".search input");


const botonCrearEmpleado = document.querySelector(".crear button");
const botonEliminarEmpleado = document.querySelector(".eliminar button");
const botonModificarEmpleado = document.querySelector(".modificar button");

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
   console.log(nombres.value);
    
    // leer datos de la BD
 
   var ident = document.getElementById("identificacion");
   ident.value=response.data["identificacion"];
   var nomb = document.getElementById("nombres");
   nomb.value=response.data["nombres"];
   var apelli = document.getElementById("apellidos");
   apelli.value=response.data["apellidos"];
   var fecha_in= document.getElementById("fecha_ingreso");
   fecha_in.value=response.data["fecha_ingreso"];
   var es_vin = document.getElementById("es_vinculado");
   es_vin.value=response.data["es_vinculado"];
   var car = document.getElementById("cargo");
   car.value=response.data["cargo"];
   var salar = document.getElementById("salario");
   salar.value=response.data["salario"];


  } catch (error) {
     // Manejo de errores
     console.error("Error al buscar el Id", error);   
     window.alert("este Id no exite; pudo haberse eliminado... intente con otro Id") 
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
  


  
  window.alert("se creo nuevo empleado") 
  

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
        window.alert("se elimino empleadocon ID: "+Id.value) 
  
  } catch (error) {
   // Manejo de errores
   console.error("Error al eliminar empleado", error);   
   window.alert("Error  al eliminar empleado") 
}
});

botonModificarEmpleado.addEventListener("click", async () => {
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

    
    // Realiza la solicitud a la API principal
   const response = await axios.patch(
      `http://localhost:8080/empleados/${Id.value}`,empleado
          );
    
  
    //console.log("cualquier cosa")
    
    window.alert("se modificó empleado con Id: "+Id.value) 
    //console.log(typeof(ident),typeof(nomb),typeof(apelli),typeof(fecha_in),typeof(car),typeof(es_vin),typeof(salar)
   // )
  
  } catch (error) {
     // Manejo de errores
     console.error("Error al modificar empleado", error);   
     window.alert("Error  al modificar empleado") 
  }
  });