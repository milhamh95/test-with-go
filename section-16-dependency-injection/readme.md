# Summary of DI

- DI is not a framework
- Di is a design pattern
- Usually use DI in Go with:
    - Interfaces (most common)
    - Variables (more restrictive, but we saw this with the `logFn` example)

- DI helps us:
    - Remove global state
    - Write implementation agnostice code
    - Write more testable code

- DI is useful for testing but that isn't its sole purpose
    - Domain driven design and other impl-agnostic design patterns are enabled mostly through DI

- DI isn't free. It is usually more code.
    - Expecting all deps for each function is less end sure friendly and can seem like a lot of code

    - Using structs to make providing deps for a set of methods will reduce code, but does so at the cost of clarity.

    - Providing defaule values with DI can take a bit more work, but reduce the mental load on a user of your code

- Despite the extra work, DI almost always results in a better codebase for large projects
    - DI makes it easier for many devs to navigate a codebase without understanding the inner workings of it all
    - Implement agnostic code means you can ignore inner workings of parts you aren't concerned with

- For smaller projects, it may be overkill
    - probably understand the inner workings of all the code so you will have a better grasp

- you can always add DI later and gradually