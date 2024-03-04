[![Test and coverage](https://github.com/ThreeDP/rinha-de-backend-2024/actions/workflows/Teste.yml/badge.svg)](https://github.com/ThreeDP/rinha-de-backend-2024/actions/workflows/Teste.yml)
[![codecov](https://codecov.io/github/ThreeDP/rinha-de-backend-2024/graph/badge.svg?token=ATN3IPF3VK)](https://codecov.io/github/ThreeDP/rinha-de-backend-2024)
# Rinha de Backend 2024
![this is fine](https://media.licdn.com/dms/image/C4E12AQGEy0nHsryvOg/article-cover_image-shrink_720_1280/0/1561242730098?e=2147483647&v=beta&t=Zvez_D5lTF6HeyMW9_KSObThHV1sRQY62sQo-QXyuRk)
O desafio para a rinha de backend 2024 / Q1 é sobre controle de concorrência em servições de crédito e débitos bancários.

## Tecnologias
Tecnologias usadas para desenvolver o desafio da rinha.

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
> Para construir a api 

![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
> Para manipular containers

![Nginx](https://img.shields.io/badge/nginx-%23009639.svg?style=for-the-badge&logo=nginx&logoColor=white)
> Como load balance

![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
> Para armazenar informações de usuários e transações 

![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white)
> Para armazenar informações gerais

## Informações sobre autor
**Nome:** Davy Paulino\
**Repo:** [Rinha de Backend](https://github.com/ThreeDP/rinha-de-backend-2024)

[![LinkedIn](https://img.shields.io/badge/linkedin-%230077B5.svg?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/davypaulinodsd/)
[![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/ThreeDP)
[![LeetCode](https://img.shields.io/badge/LeetCode-000000?style=for-the-badge&logo=LeetCode&logoColor=#d16c06)](https://leetcode.com/user7709da/)

## Resultados

## Build Project

### Create .env
> Create a `.env` file and insert the following environment variables:

```env
    POSTGRES_DB=<db_name>
    POSTGRES_USER=<db_username>
    POSTGRES_PASSWORD=<db_password>
    ALLOW_EMPTY_PASSWORD=<allow_empty_pass_on_redis[true / false]>
```

### Up Containers
Here are the commands for managing containers:

**Build and Start Containers**
```sh
make
```

**Build Containers**
```
make build
```

**Start Containers**
```sh
make up
```

**Stop Containers**
```sh
make down
```

**Cleaning Images and Volumes**
```sh
make clean
```

**Terminating Containers Process**
```sh
make fclean
```

Feel free to use these commands to set up and manage the project environment easily.
