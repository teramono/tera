package runtime

import (
	"rogchap.com/v8go"
)

type Script struct {
	Filename string
	Content  []byte
}

type Runtime struct {
	isolate *v8go.Isolate
}

func NewRuntime() (Runtime, error) {
	isolate, err := v8go.NewIsolate()
	if err != nil {
		return Runtime{}, err
	}

	return Runtime{isolate: isolate}, nil
}

func NewRuntimeWithCoreGlobals() (Runtime, error) {
	isolate, err := v8go.NewIsolate()
	if err != nil {
		return Runtime{}, err
	}

	rt := Runtime{isolate: isolate}

	if err := rt.registerCoreGlobals(); err != nil {
		return Runtime{}, err
	}

	return rt, nil
}

func (rt *Runtime) RunScript(script Script) (*v8go.Value, error) {
	ctx, err := v8go.NewContext(rt.isolate)
	if err != nil {
		return &v8go.Value{}, err
	}

	return ctx.RunScript(string(script.Content), script.Filename)
}

func (rt *Runtime) RunScriptWithCoreLib(script Script) (*v8go.Value, error) {
	ctx, err := v8go.NewContext(rt.isolate)
	if err != nil {
		return &v8go.Value{}, err
	}

	// SEC: Templates only exist for a context lifetime, so this is good. The core lib won't leak into other contexts.
	// https://v8docs.nodesource.com/node-0.8/d8/d83/classv8_1_1_function_template.html
	if err := rt.registerCoreLib(); err != nil {
		return &v8go.Value{}, err
	}

	return ctx.RunScript(string(script.Content), script.Filename)
}
