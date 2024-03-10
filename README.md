
## Run Project

0. Link API postman [postman](https://documenter.getpostman.com/view/16127230/2s93sW9ayD)
1. Dump Sql extension to your mysql database because program disable auto-migrate
2. Run to download go package
``` azure
   go mod download
   go mod tidy
```

3. Build enviroment ```Make config```
You can command in CLI ``make config``
```
listen: 0.0.0.0
port: 8081
database:
    crm:
        enable: true
        driver: ""
        host: 127.0.0.1
        port: 3306
        username: root
        password:
        database: numeric_crm_services
        options: ""
        connection: ""
        autoreconnect: false
        startinterval: 2
        maxerror: 5
        customPool: false
        maxConn: 5
        maxIdle: 5
        lifeTime: 5
        useMock: false
        multiStatements: false
    timeout: 30000
jwt:
    access: random
    refresh: random
    expiredAccess: 30
    expiredRefresh: 24
keyAES: '!a@b3b$n$j6(KQM1'
```

4. Run program
```
    make run
```

5. API URI
```

 POST   /v1/actor/register        --> crm_service/modules/actor.RequestHandlerActorStruct.CreateActor-fm (10 handlers)
 GET    /v1/actor/:id             --> crm_service/modules/actor.RequestHandlerActorStruct.GetActorById-fm (10 handlers)
 GET    /v1/actor                 --> crm_service/modules/actor.RequestHandlerActorStruct.GetAllActor-fm (10 handlers)
 PUT    /v1/actor/:id             --> crm_service/modules/actor.RequestHandlerActorStruct.UpdateActorById-fm (10 handlers)
 DELETE /v1/actor/:id             --> crm_service/modules/actor.RequestHandlerActorStruct.DeleteActorById-fm (10 handlers)
 GET    /v1/actor/:id/activate    --> crm_service/modules/actor.RequestHandlerActorStruct.ActivateActorById-fm (10 handlers)
 GET    /v1/actor/:id/deactivate  --> crm_service/modules/actor.RequestHandlerActorStruct.DeactivateActorById-fm (10 handlers)
 POST   /v1/actor/login           --> crm_service/modules/actor.RequestHandlerActorStruct.LoginActor-fm (9 handlers)
 GET    /v1/actor/logout          --> crm_service/modules/actor.RequestHandlerActorStruct.LogoutActor-fm (9 handlers)
 POST   /v1/customer/register     --> crm_service/modules/customer.RequestHandlerCustomerStruct.CreateCustomer-fm (10 handlers)
 GET    /v1/customer/:id          --> crm_service/modules/customer.RequestHandlerCustomerStruct.GetCustomerById-fm (11 handlers)
 GET    /v1/customer              --> crm_service/modules/customer.RequestHandlerCustomerStruct.GetAllCustomer-fm (11 handlers)
 PUT    /v1/customer/:id          --> crm_service/modules/customer.RequestHandlerCustomerStruct.UpdateCustomerById-fm (10 handlers)
 DELETE /v1/customer/:id          --> crm_service/modules/customer.RequestHandlerCustomerStruct.DeleteCustomerById-fm (10 handlers)
```
### PRIVILLAGE

```
CREATE USER 'superadmin'@'0.0.0.0' IDENTIFIED BY 'superadmin';
GRANT ALL PRIVILEGES ON crm_service.* TO 'superadmin'@'0.0.0.0';
SHOW GRANTS FOR 'superadmin'@'0.0.0.0';
```

### ERD DATABASES

![database](/databasex.png)

### TECH
![Tech](/tech.PNG)


## IMPROVEMENT FUTURE
1. Tracing with jaeger


