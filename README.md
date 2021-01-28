# cupon

    # Dependencias 

    - Serverless Framework

        Intalacion  [Serverless Framework](https://www.serverless.com/framework/docs/providers/aws/guide/installation/)

    - AWS CLi 

        Crear usuario aws y agregar credenciales [Serverless AWS] https://www.serverless.com/framework/docs/providers/aws/guide/credentials/

    # Confiuracion

    - Cambiar el nombre del archivo conf.yml.example a conf.yml

    - Asignar las variables     

        account:  1234123421342133  - ID usuario en AWS
        token:    "APP_USR-410832320990733" - Token mercadolibre
        host:     "https://api.mercadolibre.com"

    # Make y deploy de la Lambda 

    -  Usa en la raiz del proyecto `make build` genera el bin.
    -  Despliega `sls deploy`
    -  Al desplegar en consola se muestra el endpoint del servicio 



    





