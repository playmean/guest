package core

import (
	"fmt"

	"github.com/playmean/guest/modules"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type Core struct {
	Data *modules.ModuleData

	L *lua.LState

	modules map[string]modules.Module
}

func NewCore(data *modules.ModuleData) *Core {
	c := new(Core)
	c.Data = data

	c.modules = make(map[string]modules.Module)

	return c
}

func (c *Core) Init() error {
	if c.L != nil {
		err := c.Destroy()
		if err != nil {
			return err
		}
	}

	c.L = lua.NewState(lua.Options{
		SkipOpenLibs: true,
	})

	c.L.SetGlobal("import", luar.New(c.L, c.luaImport))

	for _, m := range modules.Internal {
		err := c.LoadModule(m)
		if err != nil {
			c.L.Close()

			return err
		}
	}

	return nil
}

func (c *Core) Destroy() error {
	if c.L == nil {
		return fmt.Errorf("core state not initialized")
	}

	c.L.Close()

	return nil
}

func (c *Core) LoadModule(module modules.Module) error {
	pkgName := module.GetName()

	if _, ok := c.modules[pkgName]; ok {
		return fmt.Errorf("module %s already exists", pkgName)
	}

	c.modules[pkgName] = module

	return nil
}

func (c *Core) Execute(source string) error {
	if err := c.L.DoString(source); err != nil {
		return err
	}

	return nil
}

func (c *Core) luaImport(pkg string) *lua.LTable {
	if module, ok := c.modules[pkg]; ok {
		return c.luaCreateImportTable(module.GetExports(c.Data))
	}

	return nil
}

func (c *Core) luaCreateImportTable(exports map[string]interface{}) *lua.LTable {
	pkg := c.L.NewTable()

	for k, v := range exports {
		c.L.SetField(pkg, k, luar.New(c.L, v))
	}

	return pkg
}
