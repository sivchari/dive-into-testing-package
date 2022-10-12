package testing

import "go/build"

func pmain() {
	// START OMIT
	// Build main package.
	pmain = &Package{
		PackagePublic: PackagePublic{
			Name:       "main",
			Dir:        p.Dir,
			GoFiles:    []string{"_testmain.go"},
			ImportPath: p.ImportPath + ".test",
			Root:       p.Root,
			Imports:    str.StringList(TestMainDeps),
			Module:     p.Module,
		},
		Internal: PackageInternal{
			Build:          &build.Package{Name: "main"},
			BuildInfo:      p.Internal.BuildInfo,
			Asmflags:       p.Internal.Asmflags,
			Gcflags:        p.Internal.Gcflags,
			Ldflags:        p.Internal.Ldflags,
			Gccgoflags:     p.Internal.Gccgoflags,
			OrigImportPath: p.Internal.OrigImportPath,
		},
	}
	// END OMIT
}
