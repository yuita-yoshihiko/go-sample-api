package testutils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func AssertResponse(t *testing.T, got, want any) {
	t.Helper()
	diff := cmp.Diff(got, want)
	if diff != "" {
		t.Errorf("Returned value is not want.\nDiff: \n%s", diff)
	}
}

func AssertResponseWithOption(t *testing.T, got, want any, opt cmp.Option) {
	t.Helper()
	if diff := cmp.Diff(got, want, opt); diff != "" {
		t.Errorf("Returned value is not want.\nDiff: \n%s", diff)
	}
}

func DefaultIgnoreFieldsOptions[T any](target T) cmp.Options {
	return cmp.Options{
		cmpopts.IgnoreFields(
			target,
			"ID",
			"CreatedAt",
			"UpdatedAt",
		),
	}
}
