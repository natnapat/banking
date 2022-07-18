# banking
2
​
3
Banking provide APIs for frontend to do following things:
4
1. Create and manage bank accounts, which are composed of owner's name, balance, and currency.
5
2. Authenticate user using PASETO or JWT token.
6
3. Record all Transactions. So every time some money is transfered, an account entry record will be created.
7
4. Perfrom a money transfer between 2 accounts. This should happen within a transaction, so that either both 
8
accounts' balance are updated or not to prevent db deadlock or race condition.
9
​
10
This project is containerized using docker and kubernetes. It's deployed on AWS.
11
​
12
I followed along with this [course](https://www.youtube.com/watch?v=rx6CPDK_5mU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=1) to create this project.
13
​
14
To practice following topics: <br/>
15
Go, Postgresql, github action, paseto and jwt, docker, kubernetes, AWS(IAM, ECR, RDS, EKS), postman
16
​