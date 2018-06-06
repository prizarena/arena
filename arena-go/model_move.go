package arena

type Move struct {
	*MoveEntity
}

type MoveEntity struct {
	Bid    int    `datastore:",noindex,omitempty"`
	Target string `datastore:",noindex,omitempty"`
}
