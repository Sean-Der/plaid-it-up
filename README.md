#Data Structures

## Customer
```
{
  id:   int
  name: string
}
```

## Account
```
{
  id:          int
  balance:     int
  customer_id: int
}
```

## Transfer
```
{
  id:             int
  amount:         int
  src_account_id: int
  dst_account_id: int

}
```

##API
There exists an endpoint for each structure at a path that matches the name of the target.

Each API provides a GET and a POST

A GET either fetches all resources, or you can provide an argument to fetch a single id. Accounts also have the src_account_id, dst_account_id paramater
allowing a user to do finer searching

A POST is used to create any object, you post the data structure. The id field may be omitted, if included it will just be ignored. Note that posting an
empty/missing field will just create the structure with the default value. ('' for string, 0 for int)

##Running
`go get github.com/sean-der/plaid-it-up`
`cd $GOPATH/sean-der/plaid-it-up`
`go run main.go`

The service runs on port 4321 by default, and at this time can only be changed by editing main.go (but could easily be an argument)
