// Code generated by github.com/vmware-tanzu/graph-framework-for-microservices/src/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/vmware-tanzu/graph-framework-for-microservices/src/gqlgen/plugin/federation/fedruntime"
)

var (
	ErrUnknownType  = errors.New("unknown type")
	ErrTypeNotFound = errors.New("type not found")
)

func (ec *executionContext) __resolve__service(ctx context.Context) (fedruntime.Service, error) {
	if ec.DisableIntrospection {
		return fedruntime.Service{}, errors.New("federated introspection disabled")
	}

	var sdl []string

	for _, src := range sources {
		if src.BuiltIn {
			continue
		}
		sdl = append(sdl, src.Input)
	}

	return fedruntime.Service{
		SDL: strings.Join(sdl, "\n"),
	}, nil
}

func (ec *executionContext) __resolve_entities(ctx context.Context, representations []map[string]interface{}) []fedruntime.Entity {
	list := make([]fedruntime.Entity, len(representations))

	repsMap := map[string]struct {
		i []int
		r []map[string]interface{}
	}{}

	// We group entities by typename so that we can parallelize their resolution.
	// This is particularly helpful when there are entity groups in multi mode.
	buildRepresentationGroups := func(reps []map[string]interface{}) {
		for i, rep := range reps {
			typeName, ok := rep["__typename"].(string)
			if !ok {
				// If there is no __typename, we just skip the representation;
				// we just won't be resolving these unknown types.
				ec.Error(ctx, errors.New("__typename must be an existing string"))
				continue
			}

			_r := repsMap[typeName]
			_r.i = append(_r.i, i)
			_r.r = append(_r.r, rep)
			repsMap[typeName] = _r
		}
	}

	isMulti := func(typeName string) bool {
		switch typeName {
		default:
			return false
		}
	}

	resolveEntity := func(ctx context.Context, typeName string, rep map[string]interface{}, idx []int, i int) (err error) {
		// we need to do our own panic handling, because we may be called in a
		// goroutine, where the usual panic handling can't catch us
		defer func() {
			if r := recover(); r != nil {
				err = ec.Recover(ctx, r)
			}
		}()

		switch typeName {
		case "ExternalExtension":
			resolverName, err := entityResolverNameForExternalExtension(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "ExternalExtension": %w`, err)
			}
			switch resolverName {

			case "findExternalExtensionByUpc":
				id0, err := ec.unmarshalNString2string(ctx, rep["upc"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findExternalExtensionByUpc(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindExternalExtensionByUpc(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "ExternalExtension": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Hello":
			resolverName, err := entityResolverNameForHello(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Hello": %w`, err)
			}
			switch resolverName {

			case "findHelloByName":
				id0, err := ec.unmarshalNString2string(ctx, rep["name"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findHelloByName(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindHelloByName(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Hello": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "NestedKey":
			resolverName, err := entityResolverNameForNestedKey(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "NestedKey": %w`, err)
			}
			switch resolverName {

			case "findNestedKeyByIDAndHelloName":
				id0, err := ec.unmarshalNString2string(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findNestedKeyByIDAndHelloName(): %w`, err)
				}
				id1, err := ec.unmarshalNString2string(ctx, rep["hello"].(map[string]interface{})["name"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 1 for findNestedKeyByIDAndHelloName(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindNestedKeyByIDAndHelloName(ctx, id0, id1)
				if err != nil {
					return fmt.Errorf(`resolving Entity "NestedKey": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "VeryNestedKey":
			resolverName, err := entityResolverNameForVeryNestedKey(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "VeryNestedKey": %w`, err)
			}
			switch resolverName {

			case "findVeryNestedKeyByIDAndHelloNameAndWorldFooAndWorldBarAndMoreWorldFoo":
				id0, err := ec.unmarshalNString2string(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findVeryNestedKeyByIDAndHelloNameAndWorldFooAndWorldBarAndMoreWorldFoo(): %w`, err)
				}
				id1, err := ec.unmarshalNString2string(ctx, rep["hello"].(map[string]interface{})["name"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 1 for findVeryNestedKeyByIDAndHelloNameAndWorldFooAndWorldBarAndMoreWorldFoo(): %w`, err)
				}
				id2, err := ec.unmarshalNString2string(ctx, rep["world"].(map[string]interface{})["foo"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 2 for findVeryNestedKeyByIDAndHelloNameAndWorldFooAndWorldBarAndMoreWorldFoo(): %w`, err)
				}
				id3, err := ec.unmarshalNInt2int(ctx, rep["world"].(map[string]interface{})["bar"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 3 for findVeryNestedKeyByIDAndHelloNameAndWorldFooAndWorldBarAndMoreWorldFoo(): %w`, err)
				}
				id4, err := ec.unmarshalNString2string(ctx, rep["more"].(map[string]interface{})["world"].(map[string]interface{})["foo"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 4 for findVeryNestedKeyByIDAndHelloNameAndWorldFooAndWorldBarAndMoreWorldFoo(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindVeryNestedKeyByIDAndHelloNameAndWorldFooAndWorldBarAndMoreWorldFoo(ctx, id0, id1, id2, id3, id4)
				if err != nil {
					return fmt.Errorf(`resolving Entity "VeryNestedKey": %w`, err)
				}

				entity.ID, err = ec.unmarshalNString2string(ctx, rep["id"])
				if err != nil {
					return err
				}
				entity.Hello.Secondary, err = ec.unmarshalNString2string(ctx, rep["hello"].(map[string]interface{})["secondary"])
				if err != nil {
					return err
				}
				list[idx[i]] = entity
				return nil
			}
		case "World":
			resolverName, err := entityResolverNameForWorld(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "World": %w`, err)
			}
			switch resolverName {

			case "findWorldByFoo":
				id0, err := ec.unmarshalNString2string(ctx, rep["foo"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findWorldByFoo(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindWorldByFoo(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "World": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			case "findWorldByBar":
				id0, err := ec.unmarshalNInt2int(ctx, rep["bar"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findWorldByBar(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindWorldByBar(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "World": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}

		}
		return fmt.Errorf("%w: %s", ErrUnknownType, typeName)
	}

	resolveManyEntities := func(ctx context.Context, typeName string, reps []map[string]interface{}, idx []int) (err error) {
		// we need to do our own panic handling, because we may be called in a
		// goroutine, where the usual panic handling can't catch us
		defer func() {
			if r := recover(); r != nil {
				err = ec.Recover(ctx, r)
			}
		}()

		switch typeName {

		default:
			return errors.New("unknown type: " + typeName)
		}
	}

	resolveEntityGroup := func(typeName string, reps []map[string]interface{}, idx []int) {
		if isMulti(typeName) {
			err := resolveManyEntities(ctx, typeName, reps, idx)
			if err != nil {
				ec.Error(ctx, err)
			}
		} else {
			// if there are multiple entities to resolve, parallelize (similar to
			// graphql.FieldSet.Dispatch)
			var e sync.WaitGroup
			e.Add(len(reps))
			for i, rep := range reps {
				i, rep := i, rep
				go func(i int, rep map[string]interface{}) {
					err := resolveEntity(ctx, typeName, rep, idx, i)
					if err != nil {
						ec.Error(ctx, err)
					}
					e.Done()
				}(i, rep)
			}
			e.Wait()
		}
	}
	buildRepresentationGroups(representations)

	switch len(repsMap) {
	case 0:
		return list
	case 1:
		for typeName, reps := range repsMap {
			resolveEntityGroup(typeName, reps.r, reps.i)
		}
		return list
	default:
		var g sync.WaitGroup
		g.Add(len(repsMap))
		for typeName, reps := range repsMap {
			go func(typeName string, reps []map[string]interface{}, idx []int) {
				resolveEntityGroup(typeName, reps, idx)
				g.Done()
			}(typeName, reps.r, reps.i)
		}
		g.Wait()
		return list
	}
}

func entityResolverNameForExternalExtension(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["upc"]; !ok {
			break
		}
		return "findExternalExtensionByUpc", nil
	}
	return "", fmt.Errorf("%w for ExternalExtension", ErrTypeNotFound)
}

func entityResolverNameForHello(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["name"]; !ok {
			break
		}
		return "findHelloByName", nil
	}
	return "", fmt.Errorf("%w for Hello", ErrTypeNotFound)
}

func entityResolverNameForNestedKey(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		m = rep
		if val, ok = m["hello"]; !ok {
			break
		}
		if m, ok = val.(map[string]interface{}); !ok {
			break
		}
		if _, ok = m["name"]; !ok {
			break
		}
		return "findNestedKeyByIDAndHelloName", nil
	}
	return "", fmt.Errorf("%w for NestedKey", ErrTypeNotFound)
}

func entityResolverNameForVeryNestedKey(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		m = rep
		if val, ok = m["hello"]; !ok {
			break
		}
		if m, ok = val.(map[string]interface{}); !ok {
			break
		}
		if _, ok = m["name"]; !ok {
			break
		}
		m = rep
		if val, ok = m["world"]; !ok {
			break
		}
		if m, ok = val.(map[string]interface{}); !ok {
			break
		}
		if _, ok = m["foo"]; !ok {
			break
		}
		m = rep
		if val, ok = m["world"]; !ok {
			break
		}
		if m, ok = val.(map[string]interface{}); !ok {
			break
		}
		if _, ok = m["bar"]; !ok {
			break
		}
		m = rep
		if val, ok = m["more"]; !ok {
			break
		}
		if m, ok = val.(map[string]interface{}); !ok {
			break
		}
		if val, ok = m["world"]; !ok {
			break
		}
		if m, ok = val.(map[string]interface{}); !ok {
			break
		}
		if _, ok = m["foo"]; !ok {
			break
		}
		return "findVeryNestedKeyByIDAndHelloNameAndWorldFooAndWorldBarAndMoreWorldFoo", nil
	}
	return "", fmt.Errorf("%w for VeryNestedKey", ErrTypeNotFound)
}

func entityResolverNameForWorld(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["foo"]; !ok {
			break
		}
		return "findWorldByFoo", nil
	}
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["bar"]; !ok {
			break
		}
		return "findWorldByBar", nil
	}
	return "", fmt.Errorf("%w for World", ErrTypeNotFound)
}
