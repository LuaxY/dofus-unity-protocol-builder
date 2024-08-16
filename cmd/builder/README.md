# Protocol Builder

## Extraction Proto Descriptors

- Create a new directory `output` in the root of the game.  
- Install last [BepInEx-Unity.IL2CPP](https://builds.bepinex.dev/projects/bepinex_be)   
- Install last [CinematicUnityExplorer (CUE)](https://github.com/originalnicodr/CinematicUnityExplorer) ([instructions](https://framedsc.com/GeneralGuides/cinematic-unity-explorer.htm))  
- Run this C# code in the game

```cs
var assemblies = AppDomain.CurrentDomain.GetAssemblies();

foreach (var ass in assemblies)
{
    if (ass.FullName == null || !ass.FullName.StartsWith("Ankama.Dofus.Protocol"))
        continue;
    
    var types = ass.GetTypes();
    foreach (var t in types)
    {
        if (!t.Name.EndsWith("Reflection"))
            continue;
        
        // once we have the type we get the static Field "Descriptor"
        var descriptorProperty = t.GetProperty("Descriptor");

        if (descriptorProperty == null)
            continue;

        var descriptor = descriptorProperty.GetValue(null);
        var descriptorType = descriptor.GetType();
        
        var protoProperty = descriptorType.GetProperty("Proto");
        var proto = protoProperty.GetValue(descriptor);
        
        // we invoke the method "ToString()"
        var toStringMethod = proto.GetType().GetMethod("ToString");
        var res = (string)toStringMethod.Invoke(proto, Array.Empty<object>());

        Il2CppSystem.IO.File.WriteAllText("./ouput/" + t.Name + ".json", res);
        
        Console.WriteLine(descriptor);
    } 
}
```

## Convert Proto Descriptors to Proto Files

- Run go program

`go run ./cmd/builder/main.go -input_dir=./path/to/game/output -output_dir=./proto`

## Thank You

- Aristochat: for BipInEx and CUE solution
- Fallen: for the extraction solution idea
- [Alpa](https://github.com/AlpaGit): for the C# extraction code