//go:build mage

package main

import (
	"context"

	"github.com/magefile/mage/mg"
	//mage:import
	"github.com/dosquad/mage"
)

var Default = TestLocal

func TestLocal(ctx context.Context) error {
	mg.CtxDeps(ctx, mage.Test)
	return nil
}
