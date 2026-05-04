# summary

## Functions

### func [Generate](./summary.go#L17)

`func Generate(mrkList SectionsMap) string`

### func [GenerateForChanges](./main.go#L14)

`func GenerateForChanges(changes shared.DiffMap) string`

## Types

### type [CategoriesMap](./types.go#L48)

`type CategoriesMap map[MarkdownCategory]SubCategoriesMap`

### type [ItemsMap](./types.go#L46)

`type ItemsMap map[MarkdownItem]PkgList`

### type [MarkdownCategory](./types.go#L15)

`type MarkdownCategory string`

#### Constants

```golang
const (
    ProdUsageCategory    MarkdownCategory = "PROD_USAGE"
    DevOnlyUsageCategory MarkdownCategory = "DEV_ONLY_USAGE"
)
```

### type [MarkdownItem](./types.go#L29)

`type MarkdownItem string`

#### Constants

```golang
const (
    UnknownUpdateItem        MarkdownItem = "UNKNOWN_UPDATE"
    SemverMajorUpgradeItem   MarkdownItem = "SEMVER_MAJOR_UPGRADE"
    SemverMinorUpgradeItem   MarkdownItem = "SEMVER_MINOR_UPGRADE"
    SemverPatchUpgradeItem   MarkdownItem = "SEMVER_PATCH_UPGRADE"
    SemverMajorDowngradeItem MarkdownItem = "SEMVER_MAJOR_DOWNGRADE"
    SemverMinorDowngradeItem MarkdownItem = "SEMVER_MINOR_DOWNGRADE"
    SemverPatchDowngradeItem MarkdownItem = "SEMVER_PATCH_DOWNGRADE"
    RemovalItem              MarkdownItem = "REMOVAL"
    AdditionItem             MarkdownItem = "ADDITION"
    SameItem                 MarkdownItem = "SAME"
)
```

### type [MarkdownSection](./types.go#L5)

`type MarkdownSection string`

#### Constants

```golang
const (
    CautionSection   MarkdownSection = "CAUTION"
    WarningSection   MarkdownSection = "WARNING"
    ImportantSection MarkdownSection = "IMPORTANT"
    TipSection       MarkdownSection = "TIP"
    NoteSection      MarkdownSection = "NOTE"
)
```

### type [MarkdownSubCategory](./types.go#L22)

`type MarkdownSubCategory string`

#### Constants

```golang
const (
    RequirementSubCategory MarkdownSubCategory = "REQUIREMENT"
    TransitiveSubCategory  MarkdownSubCategory = "TRANSITIVE"
)
```

### type [PkgList](./types.go#L45)

`type PkgList []*shared.PackageChange`

### type [SectionsMap](./types.go#L49)

`type SectionsMap map[MarkdownSection]CategoriesMap`

### type [SubCategoriesMap](./types.go#L47)

`type SubCategoriesMap map[MarkdownSubCategory]ItemsMap`

## Sub Packages

* [markdown](./markdown)

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
