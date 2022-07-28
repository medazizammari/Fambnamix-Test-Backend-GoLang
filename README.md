
# Fabnamix Test - Backend - GoLang



## Run Locally

Clone the project

```bash
  git clone https://github.com/medazizammari/Fambnamix-Test-Backend-GoLang
```

Go to the project directory


 #### 1-Make sure that MySQL server is installed
 #### 2- Create crud database
 #### 3- Create Table called person


 ```bash
  CREATE TABLE person (
                id int primary key auto_increment,
                firstname varchar(255), lastname varchar(255), date_of_birth date, address varchar(255), photo_id varchar(255)
)

```

Run server

```bash
  go run main.go
```

