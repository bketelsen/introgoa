package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("CongoStorageGroup", func() {
	Description("This is the global storage group")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the Postgres relational store")
		Model("User", func() {
			BuildsFrom(UserPayload)
			RendersTo(User)
			Description("User Model")
			HasMany("Reviews", "Review")
			HasMany("Proposals", "Proposal")
			Field("id", gorma.PKInteger, func() {
				Description("This is the User Model PK field")
			})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		Model("Proposal", func() {
			BuildsFrom(ProposalPayload)
			RendersTo(Proposal)
			Description("Proposal Model")
			BelongsTo("User")
			HasMany("Reviews", "Review")
			Field("id", gorma.PKInteger, func() {
				Description("This is the Payload Model PK field")
			})
			Field("title", func() {
				Alias("proposal_title")
			})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		Model("Review", func() {
			BuildsFrom(ReviewPayload)
			RendersTo(Review)
			Description("Review Model")
			BelongsTo("User")
			BelongsTo("Proposal")
			Field("id", gorma.PKInteger, func() {
				Description("This is the Review Model PK field")
			})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})
	})
})
