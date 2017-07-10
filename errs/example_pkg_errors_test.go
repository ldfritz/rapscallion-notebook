package errs_test

import (
	"fmt"
	"github.com/pkg/errors"
)

func raiseAnError() error {
	return errors.Wrap(errors.New("All is not well"), "error raiser")
}

func anotherLayer() error {
	return errors.Wrap(raiseAnError(), "anotherLayer()")
}

func withoutAnError() error {
	return nil
}

func anotherLayerWithoutAnError() error {
	return errors.Wrap(withoutAnError(), "anotherLayerWithoutAnError")
}

func Example_pkgErrors() {
	fmt.Println(raiseAnError())
	// Output: error raiser: All is not well
}

func Example_pkgErrorsDeeper() {
	fmt.Println(anotherLayer())
	// Output: anotherLayer(): error raiser: All is not well
}

func Example_pkgErrorsDeeperWithCause() {
	fmt.Println(errors.Cause(anotherLayer()))
	// Output: All is not well
}

func Example_pkgErrorsDepthWithoutError() {
	fmt.Println(anotherLayerWithoutAnError())
	// Output: <nil>
}
