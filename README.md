# Go Chain

Go Chain is a micro library that can be used to remove some of the common error handling boilerplate often used in GoLang.

It's intended use is to just be copy pasted in to your own project rather than importing it.

## Usage

Tired of having to write the following boilerplate after each function call?

    Var, Err := MyFunction()
    if Err != nil {
        ...
        return Err
    }

    Var2, Err := NextFunction(Var)
    if Err != nil {
        ...
        return Err
    }

    ...

Using this library allows you to easily chain function calls protecting the next function call in the event of an error being returned along the chain, minimizing the error checking in your code down to just having to do it once.

    func MyStarterFunc() (int, error) {
        // ...
    }
    func NextFunc(i int) (bool, error) {
        // ...
    }
    func NextFunc2(b bool) (float, error) {
        // ...
    }

    Var, Err := Result(Then(Then(Run(
		MyStarterFunc),
		NextFunc),
		NextFunc2))

	if Err != nil {
		...
        return Error
	}
    // Use Var

The library provides a few versions of the starter Run function, to be used when initiating a chain.
- `Run(func)` takes a 0 argument function to initate a chain
- `Run1(func, arg1)` takes a 1 argument function, and the initial argument to use when calling it, to initate a chain
- `Run2(func, arg1, arg2)` takes a 2 argument function, and the initial 2 arguments to use when calling it, to initate a chain

If your starting function is more complex than these allow, you will want to manually create a Chain variable and populate it yourself, then passing it in to `Then` methods
    func MyStarterFunc(...) (int, error) {
        // ...
    }

    Value, Error := MyStarterFunc(...)
	chain := &Chain[int] {
		Value: Value,
		Error: Error,
	}

    Var, Err := Result(Then(Then(chain,
		NextFunc),
		NextFunc2))
