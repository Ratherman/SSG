### install Golang

```
wget -c https://go.dev/dl/go1.18.2.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
export PATH=$PATH:/usr/local/go/bin
source ~/.profile

```


### run server

```
go run main.go
```