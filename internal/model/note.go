package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Note - structure of note
type Note struct {
	ID     string `json:"id,omitempty" bson:"id,omitempty"`
	UserID string `json:"user_id,omitempty" bson:"user_id"`
	Title  string `json:"title" bson:"title"`
	Data   string `json:"data" bson:"data"`
	Meta   string `json:"meta" bson:"meta"`
}

// NoteDocument - structure of note document
type NoteDocument struct {
	ID     primitive.ObjectID `bson:"_id"`
	UserID string             `bson:"user_id"`
	Title  string             `bson:"title"`
	Data   string             `bson:"data"`
	Meta   string             `bson:"meta"`
}

// NoteDocumentFromNote - create new note document from note
func NoteDocumentFromNote(note Note) (*NoteDocument, error) {
	id, err := primitive.ObjectIDFromHex(note.ID)

	if err != nil {
		return nil, err
	}

	return &NoteDocument{
		ID:     id,
		UserID: note.UserID,
		Title:  note.Title,
		Data:   note.Data,
		Meta:   note.Meta,
	}, nil
}

// NoteFromNoteDocument - create new note from note document
func NoteFromNoteDocument(noteDocument NoteDocument) *Note {
	return &Note{
		ID:     noteDocument.ID.Hex(),
		UserID: noteDocument.UserID,
		Title:  noteDocument.Title,
		Data:   noteDocument.Data,
		Meta:   noteDocument.Meta,
	}
}
