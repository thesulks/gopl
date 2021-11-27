```
❯ go run fetch.go http://golang.org
GET https://go.dev/
2021/11/27 20:36:51 http://golang.org done
failed to read response body: context canceled
```

### TODO

- Prevent response from being cancelled