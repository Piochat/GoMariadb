# GoMariadb

Prueba de conexión de Go con Mariadb. (El presente proyecto es experimental y tiene soluciones orientadas a lo que denomino "algoritmia simple" no se busca eficiencia ni uso real, recalco es un experimento para prácticar)
**CRUD** Básico tipo REST.

El paquete *middlew* es más un paquete de control que un Middleware como tal. Esto último es por conveniencia.

**Rutas**

```
localhost:8484/animeInsert 
localhost:8484/animeRead 
localhost:8484/getAnime?id= 
```

## Estructura de los datos

```json
{
  "serie": {
    "id_serie": 9,
    "name_serie": "Fairy Tail",
    "episodes_serie": 175,
    "status_serie": "Finished",
    "type_serie": "TV",
    "studio_id": 2
  },
  "studio": {
    "id_studio": 2,
    "name_studio": "A-1 Pictures"
  }
}
```

Estructura *Anime*, compuesta por la estructura *Serie* y la estructura *Studio*

**Modificaciones**

Las modificaciones son realizadas por separado, *Studio* y *Serie*.

```json
{
    "name_studio" : "Studio Deen"
}
```

---

```json
{
	"name_serie" : "Inuyasha",
	"episodes_serie" : 167,
	"status_serie" : "Finished",
	"type_serie" : "TV",
	"studio_id" : 8
}
```

