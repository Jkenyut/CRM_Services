
## Run Project

0. Link API postman [postman](https://documenter.getpostman.com/view/16127230/2s93sW9ayD)
1. Dump Sql extension to your mysql database because program disable auto-migrate
2. Build enviroment ```.env```
```
CONNECT_DB = // your uri database
HOUR = //string in number to hour
ACCESS_TOKEN_JWT = // secret-jwt
```
3. Run to download go package 
``` 
   go mod download 
   go mod tidy
```
4. Run program
```azure
go run main.go
```

### PRIVILLAGE

```
CREATE USER 'superadmin'@'0.0.0.0' IDENTIFIED BY 'superadmin';
GRANT ALL PRIVILEGES ON crm_service.* TO 'superadmin'@'0.0.0.0';
SHOW GRANTS FOR 'superadmin'@'0.0.0.0';
```

### ERD DATABASES
![](./databases1.png)
