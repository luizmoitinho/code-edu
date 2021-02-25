## Get Started
1. Clone the project in  **home/src/github.com/** and open it in the terminal
 

2. Up docker
```shell
docker-compose up -d
```
Check if it was successful 
```shell
docker ps
```

3. install dependencies
```shell
go mod init
```
```shell
go mod tidy
```

4. Run the file server.go with and test user persistence:
```shell 
go run framework/cmd/server/server.go 
```

Response:
```shell
&{{11fe48f3-f4e7-4acc-9647-5875500b19ca 2021-02-25 00:35:12.930188684 -0300 -03 m=+0.484368712 2021-02-25 00:35:12.93021402 -0300 -03 m=+0.484393987 0001-01-01 00:00:00 +0000 UTC} 
Luiz Carlos Moitinho luizcarlos_costam@hotmail.com $2a$10$4bGuP.MpGCiIh8lfLZgmJ.kwTNl08V2zta2X0jOZq3hqjJ5CXXbxO d6af0152-c753-4532-8944-091317411d66}

```
