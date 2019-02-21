# Type of test

## Unit Testing

Testing very small things, like a function , in isolation

Example:
```go

func Magic(a, b int) int {
    return (a + b) * (a + b)
}

func TestMagic(t *testing.T) {
    got := Magic(1,2)
    if got != 9 {
        t.Errorf("Magic() = %v, want %v", got, 9)
    }
}
```

## Integration tests

Test how 2+ systems work togeteher

Example:
```go
type UserStore struct{
    db *sql.DB
}

func (us *UserStore) Create(user *User) error {

}

func TestUserStore_Create(t *testing.T) {

}
```

Unit tests = using mock, maybe work well in isolation, but do not know if they work well together
Integration tests = not using mock, will verify expectations how system should work are correct

Example : using payment API
```go
customer := api.GetCustomer(email)
charge = api.CreateCharge(customer, 100)
order := createOrder(...)
shipment := createShipment(...)
```

We might test our code with mocks and assume api will work