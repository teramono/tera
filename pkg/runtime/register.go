package runtime

import (
	"github.com/teramono/capabilities/pkg/globals"
	"github.com/teramono/capabilities/pkg/lib"
	"rogchap.com/v8go"
)

func (rt *Runtime) registerCoreGlobals() error {
	if err := rt.registerFunc("require", require); err != nil {
		return err
	}

	return nil // TODO: ???
}

func (rt *Runtime) registerCoreLib() error {
	if err := rt.registerAsyncFunc("open", open); err != nil {
		return err
	}

	return nil // TODO: ???
}

func (rt *Runtime) registerFunc(opName string, templateFn func(*v8go.FunctionCallbackInfo) *v8go.Value) error {
	// Create a global object template.
	global, err := v8go.NewObjectTemplate(rt.isolate)
	if err != nil {
		return err
	}

	requireFn, _ := v8go.NewFunctionTemplate(rt.isolate, templateFn)
	if err != nil {
		return err
	}

	// Set global object and make readonly.
	return global.Set(opName, requireFn, v8go.ReadOnly)
}

func (rt *Runtime) registerAsyncFunc(opName string, templateFn func(*v8go.FunctionCallbackInfo) *v8go.Value) error {
	global, err := v8go.NewObjectTemplate(rt.isolate)
	if err != nil {
		return err
	}

	requireFn, _ := v8go.NewFunctionTemplate(rt.isolate, templateFn)
	if err != nil {
		return err
	}

	return global.Set(opName, requireFn, v8go.ReadOnly)
}

func require(info *v8go.FunctionCallbackInfo) *v8go.Value {
	args := info.Args()
	args_len := len(args)
	if args_len < 1 {
		return &v8go.Value{} // TODO: Actually throw some error here.
	}

	// Argument 1
	arg1 := args[0]
	if !arg1.IsString() {
		return &v8go.Value{} // TODO: Actually throw some error here.
	}

	filename := arg1.String()

	globals.Require(filename)

	return &v8go.Value{} // TODO:
}

func open(info *v8go.FunctionCallbackInfo) *v8go.Value {
	path := ""
	fileopts := new(lib.FileOpts)

	lib.Open(path, fileopts)

	return &v8go.Value{} // TODO:
}
