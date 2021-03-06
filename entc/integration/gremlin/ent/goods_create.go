// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/gremlin"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebook/ent/entc/integration/gremlin/ent/goods"
)

// GoodsCreate is the builder for creating a Goods entity.
type GoodsCreate struct {
	config
	mutation *GoodsMutation
	hooks    []Hook
}

// Mutation returns the GoodsMutation object of the builder.
func (gc *GoodsCreate) Mutation() *GoodsMutation {
	return gc.mutation
}

// Save creates the Goods in the database.
func (gc *GoodsCreate) Save(ctx context.Context) (*Goods, error) {
	var (
		err  error
		node *Goods
	)
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			node, err = gc.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GoodsCreate) SaveX(ctx context.Context) *Goods {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (gc *GoodsCreate) check() error {
	return nil
}

func (gc *GoodsCreate) gremlinSave(ctx context.Context) (*Goods, error) {
	res := &gremlin.Response{}
	query, bindings := gc.gremlin().Query()
	if err := gc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	_go := &Goods{config: gc.config}
	if err := _go.FromResponse(res); err != nil {
		return nil, err
	}
	return _go, nil
}

func (gc *GoodsCreate) gremlin() *dsl.Traversal {
	v := g.AddV(goods.Label)
	return v.ValueMap(true)
}

// GoodsCreateBulk is the builder for creating many Goods entities in bulk.
type GoodsCreateBulk struct {
	config
	builders []*GoodsCreate
}
