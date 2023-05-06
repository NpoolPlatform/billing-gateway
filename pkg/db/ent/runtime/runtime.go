// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/service-template/pkg/db/ent/detail"
	"github.com/NpoolPlatform/service-template/pkg/db/ent/pubsubmessage"
	"github.com/NpoolPlatform/service-template/pkg/db/ent/schema"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	detailMixin := schema.Detail{}.Mixin()
	detail.Policy = privacy.NewPolicies(detailMixin[0], schema.Detail{})
	detail.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := detail.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	detailMixinFields0 := detailMixin[0].Fields()
	_ = detailMixinFields0
	detailFields := schema.Detail{}.Fields()
	_ = detailFields
	// detailDescCreatedAt is the schema descriptor for created_at field.
	detailDescCreatedAt := detailMixinFields0[0].Descriptor()
	// detail.DefaultCreatedAt holds the default value on creation for the created_at field.
	detail.DefaultCreatedAt = detailDescCreatedAt.Default.(func() uint32)
	// detailDescUpdatedAt is the schema descriptor for updated_at field.
	detailDescUpdatedAt := detailMixinFields0[1].Descriptor()
	// detail.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	detail.DefaultUpdatedAt = detailDescUpdatedAt.Default.(func() uint32)
	// detail.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	detail.UpdateDefaultUpdatedAt = detailDescUpdatedAt.UpdateDefault.(func() uint32)
	// detailDescDeletedAt is the schema descriptor for deleted_at field.
	detailDescDeletedAt := detailMixinFields0[2].Descriptor()
	// detail.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	detail.DefaultDeletedAt = detailDescDeletedAt.Default.(func() uint32)
	// detailDescSampleCol is the schema descriptor for sample_col field.
	detailDescSampleCol := detailFields[1].Descriptor()
	// detail.DefaultSampleCol holds the default value on creation for the sample_col field.
	detail.DefaultSampleCol = detailDescSampleCol.Default.(string)
	// detailDescID is the schema descriptor for id field.
	detailDescID := detailFields[0].Descriptor()
	// detail.DefaultID holds the default value on creation for the id field.
	detail.DefaultID = detailDescID.Default.(func() uuid.UUID)
	pubsubmessageMixin := schema.PubsubMessage{}.Mixin()
	pubsubmessage.Policy = privacy.NewPolicies(pubsubmessageMixin[0], schema.PubsubMessage{})
	pubsubmessage.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := pubsubmessage.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	pubsubmessageMixinFields0 := pubsubmessageMixin[0].Fields()
	_ = pubsubmessageMixinFields0
	pubsubmessageFields := schema.PubsubMessage{}.Fields()
	_ = pubsubmessageFields
	// pubsubmessageDescCreatedAt is the schema descriptor for created_at field.
	pubsubmessageDescCreatedAt := pubsubmessageMixinFields0[0].Descriptor()
	// pubsubmessage.DefaultCreatedAt holds the default value on creation for the created_at field.
	pubsubmessage.DefaultCreatedAt = pubsubmessageDescCreatedAt.Default.(func() uint32)
	// pubsubmessageDescUpdatedAt is the schema descriptor for updated_at field.
	pubsubmessageDescUpdatedAt := pubsubmessageMixinFields0[1].Descriptor()
	// pubsubmessage.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	pubsubmessage.DefaultUpdatedAt = pubsubmessageDescUpdatedAt.Default.(func() uint32)
	// pubsubmessage.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	pubsubmessage.UpdateDefaultUpdatedAt = pubsubmessageDescUpdatedAt.UpdateDefault.(func() uint32)
	// pubsubmessageDescDeletedAt is the schema descriptor for deleted_at field.
	pubsubmessageDescDeletedAt := pubsubmessageMixinFields0[2].Descriptor()
	// pubsubmessage.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	pubsubmessage.DefaultDeletedAt = pubsubmessageDescDeletedAt.Default.(func() uint32)
	// pubsubmessageDescMessageID is the schema descriptor for message_id field.
	pubsubmessageDescMessageID := pubsubmessageFields[1].Descriptor()
	// pubsubmessage.DefaultMessageID holds the default value on creation for the message_id field.
	pubsubmessage.DefaultMessageID = pubsubmessageDescMessageID.Default.(string)
	// pubsubmessageDescState is the schema descriptor for state field.
	pubsubmessageDescState := pubsubmessageFields[2].Descriptor()
	// pubsubmessage.DefaultState holds the default value on creation for the state field.
	pubsubmessage.DefaultState = pubsubmessageDescState.Default.(string)
	// pubsubmessageDescRespToID is the schema descriptor for resp_to_id field.
	pubsubmessageDescRespToID := pubsubmessageFields[3].Descriptor()
	// pubsubmessage.DefaultRespToID holds the default value on creation for the resp_to_id field.
	pubsubmessage.DefaultRespToID = pubsubmessageDescRespToID.Default.(func() uuid.UUID)
	// pubsubmessageDescUndoID is the schema descriptor for undo_id field.
	pubsubmessageDescUndoID := pubsubmessageFields[4].Descriptor()
	// pubsubmessage.DefaultUndoID holds the default value on creation for the undo_id field.
	pubsubmessage.DefaultUndoID = pubsubmessageDescUndoID.Default.(func() uuid.UUID)
	// pubsubmessageDescArguments is the schema descriptor for arguments field.
	pubsubmessageDescArguments := pubsubmessageFields[5].Descriptor()
	// pubsubmessage.DefaultArguments holds the default value on creation for the arguments field.
	pubsubmessage.DefaultArguments = pubsubmessageDescArguments.Default.(string)
}

const (
	Version = "v0.11.2" // Version of ent codegen.
)
