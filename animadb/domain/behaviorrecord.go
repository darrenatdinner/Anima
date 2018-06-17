package domain

// BehaviorRecord - Detailed Information About The Animal's Behavior
type BehaviorRecord struct {
	AuditData

	ReactionStatement string `json:"reactionStatement,omitempty"`
	ReactionTowards   string `json:"reactionTowards,omitempty"`
	ReactionNotes     string `json:"reactionNotes,omitempty"`
}
