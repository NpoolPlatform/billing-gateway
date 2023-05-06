// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/service-template/pkg/db/ent/pubsubmessage"
	"github.com/google/uuid"
)

// PubsubMessageCreate is the builder for creating a PubsubMessage entity.
type PubsubMessageCreate struct {
	config
	mutation *PubsubMessageMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (pmc *PubsubMessageCreate) SetCreatedAt(u uint32) *PubsubMessageCreate {
	pmc.mutation.SetCreatedAt(u)
	return pmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pmc *PubsubMessageCreate) SetNillableCreatedAt(u *uint32) *PubsubMessageCreate {
	if u != nil {
		pmc.SetCreatedAt(*u)
	}
	return pmc
}

// SetUpdatedAt sets the "updated_at" field.
func (pmc *PubsubMessageCreate) SetUpdatedAt(u uint32) *PubsubMessageCreate {
	pmc.mutation.SetUpdatedAt(u)
	return pmc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pmc *PubsubMessageCreate) SetNillableUpdatedAt(u *uint32) *PubsubMessageCreate {
	if u != nil {
		pmc.SetUpdatedAt(*u)
	}
	return pmc
}

// SetDeletedAt sets the "deleted_at" field.
func (pmc *PubsubMessageCreate) SetDeletedAt(u uint32) *PubsubMessageCreate {
	pmc.mutation.SetDeletedAt(u)
	return pmc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pmc *PubsubMessageCreate) SetNillableDeletedAt(u *uint32) *PubsubMessageCreate {
	if u != nil {
		pmc.SetDeletedAt(*u)
	}
	return pmc
}

// SetAutoID sets the "auto_id" field.
func (pmc *PubsubMessageCreate) SetAutoID(i int) *PubsubMessageCreate {
	pmc.mutation.SetAutoID(i)
	return pmc
}

// SetMessageID sets the "message_id" field.
func (pmc *PubsubMessageCreate) SetMessageID(s string) *PubsubMessageCreate {
	pmc.mutation.SetMessageID(s)
	return pmc
}

// SetNillableMessageID sets the "message_id" field if the given value is not nil.
func (pmc *PubsubMessageCreate) SetNillableMessageID(s *string) *PubsubMessageCreate {
	if s != nil {
		pmc.SetMessageID(*s)
	}
	return pmc
}

// SetState sets the "state" field.
func (pmc *PubsubMessageCreate) SetState(s string) *PubsubMessageCreate {
	pmc.mutation.SetState(s)
	return pmc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (pmc *PubsubMessageCreate) SetNillableState(s *string) *PubsubMessageCreate {
	if s != nil {
		pmc.SetState(*s)
	}
	return pmc
}

// SetRespToID sets the "resp_to_id" field.
func (pmc *PubsubMessageCreate) SetRespToID(u uuid.UUID) *PubsubMessageCreate {
	pmc.mutation.SetRespToID(u)
	return pmc
}

// SetNillableRespToID sets the "resp_to_id" field if the given value is not nil.
func (pmc *PubsubMessageCreate) SetNillableRespToID(u *uuid.UUID) *PubsubMessageCreate {
	if u != nil {
		pmc.SetRespToID(*u)
	}
	return pmc
}

// SetUndoID sets the "undo_id" field.
func (pmc *PubsubMessageCreate) SetUndoID(u uuid.UUID) *PubsubMessageCreate {
	pmc.mutation.SetUndoID(u)
	return pmc
}

// SetNillableUndoID sets the "undo_id" field if the given value is not nil.
func (pmc *PubsubMessageCreate) SetNillableUndoID(u *uuid.UUID) *PubsubMessageCreate {
	if u != nil {
		pmc.SetUndoID(*u)
	}
	return pmc
}

// SetArguments sets the "arguments" field.
func (pmc *PubsubMessageCreate) SetArguments(s string) *PubsubMessageCreate {
	pmc.mutation.SetArguments(s)
	return pmc
}

// SetNillableArguments sets the "arguments" field if the given value is not nil.
func (pmc *PubsubMessageCreate) SetNillableArguments(s *string) *PubsubMessageCreate {
	if s != nil {
		pmc.SetArguments(*s)
	}
	return pmc
}

// SetID sets the "id" field.
func (pmc *PubsubMessageCreate) SetID(u uuid.UUID) *PubsubMessageCreate {
	pmc.mutation.SetID(u)
	return pmc
}

// Mutation returns the PubsubMessageMutation object of the builder.
func (pmc *PubsubMessageCreate) Mutation() *PubsubMessageMutation {
	return pmc.mutation
}

// Save creates the PubsubMessage in the database.
func (pmc *PubsubMessageCreate) Save(ctx context.Context) (*PubsubMessage, error) {
	var (
		err  error
		node *PubsubMessage
	)
	if err := pmc.defaults(); err != nil {
		return nil, err
	}
	if len(pmc.hooks) == 0 {
		if err = pmc.check(); err != nil {
			return nil, err
		}
		node, err = pmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PubsubMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pmc.check(); err != nil {
				return nil, err
			}
			pmc.mutation = mutation
			if node, err = pmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pmc.hooks) - 1; i >= 0; i-- {
			if pmc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pmc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pmc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PubsubMessage)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PubsubMessageMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pmc *PubsubMessageCreate) SaveX(ctx context.Context) *PubsubMessage {
	v, err := pmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pmc *PubsubMessageCreate) Exec(ctx context.Context) error {
	_, err := pmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmc *PubsubMessageCreate) ExecX(ctx context.Context) {
	if err := pmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pmc *PubsubMessageCreate) defaults() error {
	if _, ok := pmc.mutation.CreatedAt(); !ok {
		if pubsubmessage.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized pubsubmessage.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := pubsubmessage.DefaultCreatedAt()
		pmc.mutation.SetCreatedAt(v)
	}
	if _, ok := pmc.mutation.UpdatedAt(); !ok {
		if pubsubmessage.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized pubsubmessage.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := pubsubmessage.DefaultUpdatedAt()
		pmc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pmc.mutation.DeletedAt(); !ok {
		if pubsubmessage.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized pubsubmessage.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := pubsubmessage.DefaultDeletedAt()
		pmc.mutation.SetDeletedAt(v)
	}
	if _, ok := pmc.mutation.MessageID(); !ok {
		v := pubsubmessage.DefaultMessageID
		pmc.mutation.SetMessageID(v)
	}
	if _, ok := pmc.mutation.State(); !ok {
		v := pubsubmessage.DefaultState
		pmc.mutation.SetState(v)
	}
	if _, ok := pmc.mutation.RespToID(); !ok {
		if pubsubmessage.DefaultRespToID == nil {
			return fmt.Errorf("ent: uninitialized pubsubmessage.DefaultRespToID (forgotten import ent/runtime?)")
		}
		v := pubsubmessage.DefaultRespToID()
		pmc.mutation.SetRespToID(v)
	}
	if _, ok := pmc.mutation.UndoID(); !ok {
		if pubsubmessage.DefaultUndoID == nil {
			return fmt.Errorf("ent: uninitialized pubsubmessage.DefaultUndoID (forgotten import ent/runtime?)")
		}
		v := pubsubmessage.DefaultUndoID()
		pmc.mutation.SetUndoID(v)
	}
	if _, ok := pmc.mutation.Arguments(); !ok {
		v := pubsubmessage.DefaultArguments
		pmc.mutation.SetArguments(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pmc *PubsubMessageCreate) check() error {
	if _, ok := pmc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "PubsubMessage.created_at"`)}
	}
	if _, ok := pmc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "PubsubMessage.updated_at"`)}
	}
	if _, ok := pmc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "PubsubMessage.deleted_at"`)}
	}
	if _, ok := pmc.mutation.AutoID(); !ok {
		return &ValidationError{Name: "auto_id", err: errors.New(`ent: missing required field "PubsubMessage.auto_id"`)}
	}
	return nil
}

func (pmc *PubsubMessageCreate) sqlSave(ctx context.Context) (*PubsubMessage, error) {
	_node, _spec := pmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (pmc *PubsubMessageCreate) createSpec() (*PubsubMessage, *sqlgraph.CreateSpec) {
	var (
		_node = &PubsubMessage{config: pmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: pubsubmessage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: pubsubmessage.FieldID,
			},
		}
	)
	_spec.OnConflict = pmc.conflict
	if id, ok := pmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pmc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pubsubmessage.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pmc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pubsubmessage.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := pmc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pubsubmessage.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := pmc.mutation.AutoID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pubsubmessage.FieldAutoID,
		})
		_node.AutoID = value
	}
	if value, ok := pmc.mutation.MessageID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pubsubmessage.FieldMessageID,
		})
		_node.MessageID = value
	}
	if value, ok := pmc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pubsubmessage.FieldState,
		})
		_node.State = value
	}
	if value, ok := pmc.mutation.RespToID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: pubsubmessage.FieldRespToID,
		})
		_node.RespToID = value
	}
	if value, ok := pmc.mutation.UndoID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: pubsubmessage.FieldUndoID,
		})
		_node.UndoID = value
	}
	if value, ok := pmc.mutation.Arguments(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pubsubmessage.FieldArguments,
		})
		_node.Arguments = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PubsubMessage.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PubsubMessageUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (pmc *PubsubMessageCreate) OnConflict(opts ...sql.ConflictOption) *PubsubMessageUpsertOne {
	pmc.conflict = opts
	return &PubsubMessageUpsertOne{
		create: pmc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PubsubMessage.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pmc *PubsubMessageCreate) OnConflictColumns(columns ...string) *PubsubMessageUpsertOne {
	pmc.conflict = append(pmc.conflict, sql.ConflictColumns(columns...))
	return &PubsubMessageUpsertOne{
		create: pmc,
	}
}

type (
	// PubsubMessageUpsertOne is the builder for "upsert"-ing
	//  one PubsubMessage node.
	PubsubMessageUpsertOne struct {
		create *PubsubMessageCreate
	}

	// PubsubMessageUpsert is the "OnConflict" setter.
	PubsubMessageUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *PubsubMessageUpsert) SetCreatedAt(v uint32) *PubsubMessageUpsert {
	u.Set(pubsubmessage.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PubsubMessageUpsert) UpdateCreatedAt() *PubsubMessageUpsert {
	u.SetExcluded(pubsubmessage.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *PubsubMessageUpsert) AddCreatedAt(v uint32) *PubsubMessageUpsert {
	u.Add(pubsubmessage.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PubsubMessageUpsert) SetUpdatedAt(v uint32) *PubsubMessageUpsert {
	u.Set(pubsubmessage.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PubsubMessageUpsert) UpdateUpdatedAt() *PubsubMessageUpsert {
	u.SetExcluded(pubsubmessage.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *PubsubMessageUpsert) AddUpdatedAt(v uint32) *PubsubMessageUpsert {
	u.Add(pubsubmessage.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PubsubMessageUpsert) SetDeletedAt(v uint32) *PubsubMessageUpsert {
	u.Set(pubsubmessage.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PubsubMessageUpsert) UpdateDeletedAt() *PubsubMessageUpsert {
	u.SetExcluded(pubsubmessage.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *PubsubMessageUpsert) AddDeletedAt(v uint32) *PubsubMessageUpsert {
	u.Add(pubsubmessage.FieldDeletedAt, v)
	return u
}

// SetAutoID sets the "auto_id" field.
func (u *PubsubMessageUpsert) SetAutoID(v int) *PubsubMessageUpsert {
	u.Set(pubsubmessage.FieldAutoID, v)
	return u
}

// UpdateAutoID sets the "auto_id" field to the value that was provided on create.
func (u *PubsubMessageUpsert) UpdateAutoID() *PubsubMessageUpsert {
	u.SetExcluded(pubsubmessage.FieldAutoID)
	return u
}

// AddAutoID adds v to the "auto_id" field.
func (u *PubsubMessageUpsert) AddAutoID(v int) *PubsubMessageUpsert {
	u.Add(pubsubmessage.FieldAutoID, v)
	return u
}

// SetMessageID sets the "message_id" field.
func (u *PubsubMessageUpsert) SetMessageID(v string) *PubsubMessageUpsert {
	u.Set(pubsubmessage.FieldMessageID, v)
	return u
}

// UpdateMessageID sets the "message_id" field to the value that was provided on create.
func (u *PubsubMessageUpsert) UpdateMessageID() *PubsubMessageUpsert {
	u.SetExcluded(pubsubmessage.FieldMessageID)
	return u
}

// ClearMessageID clears the value of the "message_id" field.
func (u *PubsubMessageUpsert) ClearMessageID() *PubsubMessageUpsert {
	u.SetNull(pubsubmessage.FieldMessageID)
	return u
}

// SetState sets the "state" field.
func (u *PubsubMessageUpsert) SetState(v string) *PubsubMessageUpsert {
	u.Set(pubsubmessage.FieldState, v)
	return u
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *PubsubMessageUpsert) UpdateState() *PubsubMessageUpsert {
	u.SetExcluded(pubsubmessage.FieldState)
	return u
}

// ClearState clears the value of the "state" field.
func (u *PubsubMessageUpsert) ClearState() *PubsubMessageUpsert {
	u.SetNull(pubsubmessage.FieldState)
	return u
}

// SetRespToID sets the "resp_to_id" field.
func (u *PubsubMessageUpsert) SetRespToID(v uuid.UUID) *PubsubMessageUpsert {
	u.Set(pubsubmessage.FieldRespToID, v)
	return u
}

// UpdateRespToID sets the "resp_to_id" field to the value that was provided on create.
func (u *PubsubMessageUpsert) UpdateRespToID() *PubsubMessageUpsert {
	u.SetExcluded(pubsubmessage.FieldRespToID)
	return u
}

// ClearRespToID clears the value of the "resp_to_id" field.
func (u *PubsubMessageUpsert) ClearRespToID() *PubsubMessageUpsert {
	u.SetNull(pubsubmessage.FieldRespToID)
	return u
}

// SetUndoID sets the "undo_id" field.
func (u *PubsubMessageUpsert) SetUndoID(v uuid.UUID) *PubsubMessageUpsert {
	u.Set(pubsubmessage.FieldUndoID, v)
	return u
}

// UpdateUndoID sets the "undo_id" field to the value that was provided on create.
func (u *PubsubMessageUpsert) UpdateUndoID() *PubsubMessageUpsert {
	u.SetExcluded(pubsubmessage.FieldUndoID)
	return u
}

// ClearUndoID clears the value of the "undo_id" field.
func (u *PubsubMessageUpsert) ClearUndoID() *PubsubMessageUpsert {
	u.SetNull(pubsubmessage.FieldUndoID)
	return u
}

// SetArguments sets the "arguments" field.
func (u *PubsubMessageUpsert) SetArguments(v string) *PubsubMessageUpsert {
	u.Set(pubsubmessage.FieldArguments, v)
	return u
}

// UpdateArguments sets the "arguments" field to the value that was provided on create.
func (u *PubsubMessageUpsert) UpdateArguments() *PubsubMessageUpsert {
	u.SetExcluded(pubsubmessage.FieldArguments)
	return u
}

// ClearArguments clears the value of the "arguments" field.
func (u *PubsubMessageUpsert) ClearArguments() *PubsubMessageUpsert {
	u.SetNull(pubsubmessage.FieldArguments)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.PubsubMessage.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(pubsubmessage.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *PubsubMessageUpsertOne) UpdateNewValues() *PubsubMessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(pubsubmessage.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.PubsubMessage.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *PubsubMessageUpsertOne) Ignore() *PubsubMessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PubsubMessageUpsertOne) DoNothing() *PubsubMessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PubsubMessageCreate.OnConflict
// documentation for more info.
func (u *PubsubMessageUpsertOne) Update(set func(*PubsubMessageUpsert)) *PubsubMessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PubsubMessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *PubsubMessageUpsertOne) SetCreatedAt(v uint32) *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *PubsubMessageUpsertOne) AddCreatedAt(v uint32) *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PubsubMessageUpsertOne) UpdateCreatedAt() *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PubsubMessageUpsertOne) SetUpdatedAt(v uint32) *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *PubsubMessageUpsertOne) AddUpdatedAt(v uint32) *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PubsubMessageUpsertOne) UpdateUpdatedAt() *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PubsubMessageUpsertOne) SetDeletedAt(v uint32) *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *PubsubMessageUpsertOne) AddDeletedAt(v uint32) *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PubsubMessageUpsertOne) UpdateDeletedAt() *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAutoID sets the "auto_id" field.
func (u *PubsubMessageUpsertOne) SetAutoID(v int) *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetAutoID(v)
	})
}

// AddAutoID adds v to the "auto_id" field.
func (u *PubsubMessageUpsertOne) AddAutoID(v int) *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.AddAutoID(v)
	})
}

// UpdateAutoID sets the "auto_id" field to the value that was provided on create.
func (u *PubsubMessageUpsertOne) UpdateAutoID() *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateAutoID()
	})
}

// SetMessageID sets the "message_id" field.
func (u *PubsubMessageUpsertOne) SetMessageID(v string) *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetMessageID(v)
	})
}

// UpdateMessageID sets the "message_id" field to the value that was provided on create.
func (u *PubsubMessageUpsertOne) UpdateMessageID() *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateMessageID()
	})
}

// ClearMessageID clears the value of the "message_id" field.
func (u *PubsubMessageUpsertOne) ClearMessageID() *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.ClearMessageID()
	})
}

// SetState sets the "state" field.
func (u *PubsubMessageUpsertOne) SetState(v string) *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *PubsubMessageUpsertOne) UpdateState() *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateState()
	})
}

// ClearState clears the value of the "state" field.
func (u *PubsubMessageUpsertOne) ClearState() *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.ClearState()
	})
}

// SetRespToID sets the "resp_to_id" field.
func (u *PubsubMessageUpsertOne) SetRespToID(v uuid.UUID) *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetRespToID(v)
	})
}

// UpdateRespToID sets the "resp_to_id" field to the value that was provided on create.
func (u *PubsubMessageUpsertOne) UpdateRespToID() *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateRespToID()
	})
}

// ClearRespToID clears the value of the "resp_to_id" field.
func (u *PubsubMessageUpsertOne) ClearRespToID() *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.ClearRespToID()
	})
}

// SetUndoID sets the "undo_id" field.
func (u *PubsubMessageUpsertOne) SetUndoID(v uuid.UUID) *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetUndoID(v)
	})
}

// UpdateUndoID sets the "undo_id" field to the value that was provided on create.
func (u *PubsubMessageUpsertOne) UpdateUndoID() *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateUndoID()
	})
}

// ClearUndoID clears the value of the "undo_id" field.
func (u *PubsubMessageUpsertOne) ClearUndoID() *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.ClearUndoID()
	})
}

// SetArguments sets the "arguments" field.
func (u *PubsubMessageUpsertOne) SetArguments(v string) *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetArguments(v)
	})
}

// UpdateArguments sets the "arguments" field to the value that was provided on create.
func (u *PubsubMessageUpsertOne) UpdateArguments() *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateArguments()
	})
}

// ClearArguments clears the value of the "arguments" field.
func (u *PubsubMessageUpsertOne) ClearArguments() *PubsubMessageUpsertOne {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.ClearArguments()
	})
}

// Exec executes the query.
func (u *PubsubMessageUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PubsubMessageCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PubsubMessageUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PubsubMessageUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: PubsubMessageUpsertOne.ID is not supported by MySQL driver. Use PubsubMessageUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PubsubMessageUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PubsubMessageCreateBulk is the builder for creating many PubsubMessage entities in bulk.
type PubsubMessageCreateBulk struct {
	config
	builders []*PubsubMessageCreate
	conflict []sql.ConflictOption
}

// Save creates the PubsubMessage entities in the database.
func (pmcb *PubsubMessageCreateBulk) Save(ctx context.Context) ([]*PubsubMessage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pmcb.builders))
	nodes := make([]*PubsubMessage, len(pmcb.builders))
	mutators := make([]Mutator, len(pmcb.builders))
	for i := range pmcb.builders {
		func(i int, root context.Context) {
			builder := pmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PubsubMessageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pmcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pmcb *PubsubMessageCreateBulk) SaveX(ctx context.Context) []*PubsubMessage {
	v, err := pmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pmcb *PubsubMessageCreateBulk) Exec(ctx context.Context) error {
	_, err := pmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmcb *PubsubMessageCreateBulk) ExecX(ctx context.Context) {
	if err := pmcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PubsubMessage.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PubsubMessageUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (pmcb *PubsubMessageCreateBulk) OnConflict(opts ...sql.ConflictOption) *PubsubMessageUpsertBulk {
	pmcb.conflict = opts
	return &PubsubMessageUpsertBulk{
		create: pmcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PubsubMessage.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (pmcb *PubsubMessageCreateBulk) OnConflictColumns(columns ...string) *PubsubMessageUpsertBulk {
	pmcb.conflict = append(pmcb.conflict, sql.ConflictColumns(columns...))
	return &PubsubMessageUpsertBulk{
		create: pmcb,
	}
}

// PubsubMessageUpsertBulk is the builder for "upsert"-ing
// a bulk of PubsubMessage nodes.
type PubsubMessageUpsertBulk struct {
	create *PubsubMessageCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.PubsubMessage.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(pubsubmessage.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *PubsubMessageUpsertBulk) UpdateNewValues() *PubsubMessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(pubsubmessage.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PubsubMessage.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *PubsubMessageUpsertBulk) Ignore() *PubsubMessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PubsubMessageUpsertBulk) DoNothing() *PubsubMessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PubsubMessageCreateBulk.OnConflict
// documentation for more info.
func (u *PubsubMessageUpsertBulk) Update(set func(*PubsubMessageUpsert)) *PubsubMessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PubsubMessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *PubsubMessageUpsertBulk) SetCreatedAt(v uint32) *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *PubsubMessageUpsertBulk) AddCreatedAt(v uint32) *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PubsubMessageUpsertBulk) UpdateCreatedAt() *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PubsubMessageUpsertBulk) SetUpdatedAt(v uint32) *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *PubsubMessageUpsertBulk) AddUpdatedAt(v uint32) *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PubsubMessageUpsertBulk) UpdateUpdatedAt() *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PubsubMessageUpsertBulk) SetDeletedAt(v uint32) *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *PubsubMessageUpsertBulk) AddDeletedAt(v uint32) *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PubsubMessageUpsertBulk) UpdateDeletedAt() *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAutoID sets the "auto_id" field.
func (u *PubsubMessageUpsertBulk) SetAutoID(v int) *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetAutoID(v)
	})
}

// AddAutoID adds v to the "auto_id" field.
func (u *PubsubMessageUpsertBulk) AddAutoID(v int) *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.AddAutoID(v)
	})
}

// UpdateAutoID sets the "auto_id" field to the value that was provided on create.
func (u *PubsubMessageUpsertBulk) UpdateAutoID() *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateAutoID()
	})
}

// SetMessageID sets the "message_id" field.
func (u *PubsubMessageUpsertBulk) SetMessageID(v string) *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetMessageID(v)
	})
}

// UpdateMessageID sets the "message_id" field to the value that was provided on create.
func (u *PubsubMessageUpsertBulk) UpdateMessageID() *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateMessageID()
	})
}

// ClearMessageID clears the value of the "message_id" field.
func (u *PubsubMessageUpsertBulk) ClearMessageID() *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.ClearMessageID()
	})
}

// SetState sets the "state" field.
func (u *PubsubMessageUpsertBulk) SetState(v string) *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *PubsubMessageUpsertBulk) UpdateState() *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateState()
	})
}

// ClearState clears the value of the "state" field.
func (u *PubsubMessageUpsertBulk) ClearState() *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.ClearState()
	})
}

// SetRespToID sets the "resp_to_id" field.
func (u *PubsubMessageUpsertBulk) SetRespToID(v uuid.UUID) *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetRespToID(v)
	})
}

// UpdateRespToID sets the "resp_to_id" field to the value that was provided on create.
func (u *PubsubMessageUpsertBulk) UpdateRespToID() *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateRespToID()
	})
}

// ClearRespToID clears the value of the "resp_to_id" field.
func (u *PubsubMessageUpsertBulk) ClearRespToID() *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.ClearRespToID()
	})
}

// SetUndoID sets the "undo_id" field.
func (u *PubsubMessageUpsertBulk) SetUndoID(v uuid.UUID) *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetUndoID(v)
	})
}

// UpdateUndoID sets the "undo_id" field to the value that was provided on create.
func (u *PubsubMessageUpsertBulk) UpdateUndoID() *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateUndoID()
	})
}

// ClearUndoID clears the value of the "undo_id" field.
func (u *PubsubMessageUpsertBulk) ClearUndoID() *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.ClearUndoID()
	})
}

// SetArguments sets the "arguments" field.
func (u *PubsubMessageUpsertBulk) SetArguments(v string) *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.SetArguments(v)
	})
}

// UpdateArguments sets the "arguments" field to the value that was provided on create.
func (u *PubsubMessageUpsertBulk) UpdateArguments() *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.UpdateArguments()
	})
}

// ClearArguments clears the value of the "arguments" field.
func (u *PubsubMessageUpsertBulk) ClearArguments() *PubsubMessageUpsertBulk {
	return u.Update(func(s *PubsubMessageUpsert) {
		s.ClearArguments()
	})
}

// Exec executes the query.
func (u *PubsubMessageUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PubsubMessageCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PubsubMessageCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PubsubMessageUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
