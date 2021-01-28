# Challenge

# Dependencias 

- Serverless Framework

    - Instalacion  [Serverless Framework](https://www.serverless.com/framework/docs/providers/aws/guide/installation/)

- AWS CLi 

   - Crear usuario aws y agregar credenciales [Serverless AWS](https://www.serverless.com/framework/docs/providers/aws/guide/credentials/)

# Configuración

- Cambiar el nombre del archivo conf.yml.example a conf.yml

- Asignar las variables     

    - account:  1234123421342133  - ID usuario en AWS
    - token:    "APP_USR-410832320990733" - Token mercadolibre
    - host:     "https://api.mercadolibre.com"

# Make y deploy de la Lambda 

-  Usa en la raíz del proyecto `make build` para generar el bin.
-  Despliega con `sls deploy` en la raíz 
-  Al desplegar en consola se muestra el endpoint del servicio


# Uso de la API "Punto 2"

- Items creados de pruebas "MLA905913105", "MLA906002266","MLA906002298", "MLA906003946","MLA906004194"

- Para crear o eliminar mas item ingresar a Mercadolibre [Mercadolibre](https://www.mercadolibre.com/jms/mla/lgz/msl/login/H4sIAAAAAAAEAzWOQQ7DIAwE_-JzFO4c-xHkEkNQoSDjiFZR_l4Tqcddj8c-IdeY3k6-jcACfVpOPgks0DJKqFxc2nRQslY9Cf0jTgQZCwlxB3tOUaTtQbo0VcIHKYOH7C7kOrS6T2kXq4ZdpHVrzBhjLcQet5rTk2n1tazIRjmmmLrqaX5w-64FAnZxwuhfYAPmTtcPRoWPuMQAAAA/user) con el usuario:  
    - "email":"test_user_41464011@testuser.com"
    - "password":"qatest5737"

- Servicio POST `https://9f0ybavyh3.execute-api.us-east-1.amazonaws.com/dev/coupon`

    ## Body ejemplo

    `{"item_ids": ["MLA905913105", "MLA906002266","MLA906002298", "MLA906003946","MLA906004194"],"amount": 500}`


# Pruebas y coverage 

    -  Usa en la raíz del proyecto `go test ./... -cover` 