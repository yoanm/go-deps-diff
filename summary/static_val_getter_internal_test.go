package summary

import (
	"testing"
)

const _testUnmanagedSection markdownSection = "ARGH"

func Test_getSectionHeaderFor_unmanaged(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Log("function thrown a panic as expected.")
		}
	}()

	_ = getSectionHeaderFor(_testUnmanagedSection)

	t.Fatal("The code did not panic")
}

func Test_getSectionDescriptionFor_unmanaged(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Log("function thrown a panic as expected.")
		}
	}()

	_ = getSectionDescriptionFor(_testUnmanagedSection)

	t.Fatal("The code did not panic")
}

const _testUnmanagedCategory markdownCategory = "ARGH"

func Test_getCategoryHeaderFor_unmanaged(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Log("function thrown a panic as expected.")
		}
	}()

	_ = getCategoryHeaderFor(_testUnmanagedCategory)

	t.Fatal("The code did not panic")
}
