package contract

type DiffMap map[string]*PackageChange

// PackageChange contains detailed information about a package difference.
type PackageChange struct {
	Package   PkgWrapper
	Operation Operation

	PreviousVersion PkgVersion // Only available for updated packages ! Empty (zero value) otherwise.
}
