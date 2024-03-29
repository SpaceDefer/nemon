import importlib.util
import sys

name = 'winapps'
if name in sys.modules:
    pass
elif (spec := importlib.util.find_spec(name)) is not None:
    module = importlib.util.module_from_spec(spec)
    sys.modules[name] = module
    spec.loader.exec_module(module)
else:
    sys.exit(f"can't find the {name!r} module")

winapps = sys.modules[name]

for app in winapps.list_installed():
    print(app.name)