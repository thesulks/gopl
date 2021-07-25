```
❯ ./pipeline -pipe=2000000 
npipe: 2000000  904.19952ms
~/go/src/gopl/ch9/ex4 master*
❯ ./pipeline -pipe=3000000
npipe: 3000000  1.390734452s
~/go/src/gopl/ch9/ex4 master* 7s
❯ ./pipeline -pipe=5000000
npipe: 5000000  10.251441244s
~/go/src/gopl/ch9/ex4 master* 20s
❯ ./pipeline -pipe=10000000
npipe: 10000000  33.157361124s
❯ ./pipeline -pipe=100000000
... killed
```
