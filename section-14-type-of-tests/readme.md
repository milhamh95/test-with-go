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

## End to end tests

Test entire application

Ex1: 
1. Start app
2. Open chrome
3. Navigate to your app
4. Login as a user would
5. And so on...

Great for simultaing real user scenarions.

Con : not always great at pointing at why those bugs occured or how to fix them.
You can figure it out, but they don't tell you quickly and clearly.

Joke ex (both e2e & integration): Two doors both work in isolation, but put the two together and suddenly
the doorknobs block each other from opening!