package filter

import (
	"errors"
	"reflect"
	"sync"
)

// Simple error collection
// Source https://github.com/bastianrob/go-experiences/tree/master/filter
var (
	ErrInvalidSourceKind = errors.New("source value is not an array or slice")
	ErrFilterFuncNil     = errors.New("filter function cannot be nil")
	ErrFilterNotFunc     = errors.New("filter argument must be a function")
)

// Simple an array/slice will guarantee same input order of results
// Source https://github.com/bastianrob/go-experiences/tree/master/filter
func Simple(source any, filterFunc any) (any, error) {
	srcV := reflect.ValueOf(source)
	kind := srcV.Kind()

	if kind != reflect.Slice && kind != reflect.Array {
		return nil, ErrInvalidSourceKind
	}

	if filterFunc == nil {
		return nil, ErrFilterFuncNil
	}

	fv := reflect.ValueOf(filterFunc)
	if fv.Kind() != reflect.Func {
		return nil, ErrFilterNotFunc
	}

	T := reflect.TypeOf(source).Elem()
	sliceOfT := reflect.MakeSlice(reflect.SliceOf(T), 0, 0)
	ptrToSliceOfT := reflect.New(sliceOfT.Type())
	ptrToElementOfSliceT := ptrToSliceOfT.Elem()

	for i := 0; i < srcV.Len(); i++ {
		entry := srcV.Index(i)
		valid := fv.
			Call([]reflect.Value{entry})[0].
			Interface().(bool)

		if valid {
			appendResult := reflect.Append(ptrToElementOfSliceT, entry)
			ptrToElementOfSliceT.Set(appendResult)
		}
	}

	return ptrToElementOfSliceT.Interface(), nil
}

// Parallel an array using go routine
// This function will not guarantee order of results
// Source https://github.com/bastianrob/go-experiences/tree/master/filter
func Parallel(source any, filterFunc any) (any, error) {
	srcV := reflect.ValueOf(source)
	kind := srcV.Kind()

	if kind != reflect.Slice && kind != reflect.Array {
		return nil, ErrInvalidSourceKind
	}

	if filterFunc == nil {
		return nil, ErrFilterFuncNil
	}

	fv := reflect.ValueOf(filterFunc)

	if fv.Kind() != reflect.Func {
		return nil, ErrFilterNotFunc
	}

	T := reflect.TypeOf(source).Elem()
	sliceOfT := reflect.MakeSlice(reflect.SliceOf(T), 0, 0)
	ptrToSliceOfT := reflect.New(sliceOfT.Type())
	ptrToElementOfSliceT := ptrToSliceOfT.Elem()

	wg := &sync.WaitGroup{}
	wg.Add(srcV.Len())

	queue := make(chan *reflect.Value, 3)

	go func() {
		for entry := range queue {
			if entry != nil {
				appendResult := reflect.Append(ptrToElementOfSliceT, *entry)
				ptrToElementOfSliceT.Set(appendResult)
			}
			wg.Done()
		}
	}()

	for i := 0; i < srcV.Len(); i++ {
		go func(idx int, entry reflect.Value) {
			valid := fv.
				Call([]reflect.Value{entry})[0].
				Interface().(bool)
			if valid {
				queue <- &entry
			} else {
				queue <- nil
			}
		}(i, srcV.Index(i))
	}

	wg.Wait()
	close(queue)

	return ptrToElementOfSliceT.Interface(), nil
}

// Contains check if element exist in a slice
func Contains(source any, item any) (bool, error) {
	result, err := Simple(source, func(entry any) bool {
		return entry == item
	})

	if err != nil {
		return false, err
	}

	srcV := reflect.ValueOf(result)

	return srcV.Len() > 0, nil
}

// ContainsParallel check if element exist in a slice using go routines
func ContainsParallel(source any, item any) (bool, error) {
	result, err := Parallel(source, func(entry any) bool {
		return entry == item
	})

	if err != nil {
		return false, err
	}

	srcV := reflect.ValueOf(result)

	return srcV.Len() > 0, nil
}
