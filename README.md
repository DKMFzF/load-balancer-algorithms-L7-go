# L7 balancers on Go 

L7 balancing algorithms that have been implemented:
- Round Robin
- Weighted Round Robin
- Least Connections

### Version
```bash
go version go1.24.6 linux/amd64  
```

### How to start?

```bash
git clone https://github.com/DKMFzF/load-balancer-algorithms-L7-go.git

cd load-balancer-algorithms-L7-go 

go mod download

make run-servers

# choose which algorithm you want to use
make run-rr
make run-wrr
make run-lb
```

