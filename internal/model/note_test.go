package model

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNoteDocumentFromNote(t *testing.T) {
	sucessID := primitive.NewObjectID()
	errID := "123"

	successSource := &Note{
		ID:     sucessID.Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Data:   "text",
		Title:  "title",
		Meta:   "meta",
	}

	errSource := &Note{
		ID:     errID,
		UserID: primitive.NewObjectID().Hex(),
		Data:   "text",
		Title:  "title",
		Meta:   "meta",
	}

	successResult := &NoteDocument{
		ID:     sucessID,
		UserID: successSource.UserID,
		Title:  successSource.Title,
		Data:   successSource.Data,
		Meta:   successSource.Meta,
	}
	type args struct {
		note Note
	}
	tests := []struct {
		name    string
		args    args
		want    *NoteDocument
		wantErr bool
	}{
		{
			name: "1# success",
			args: args{
				note: *successSource,
			},
			want:    successResult,
			wantErr: false,
		},
		{
			name: "2# error",
			args: args{
				note: *errSource,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NoteDocumentFromNote(tt.args.note)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoteDocumentFromNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteDocumentFromNote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteFromNoteDocument(t *testing.T) {
	sucessID := primitive.NewObjectID()

	successSource := &NoteDocument{
		ID:     sucessID,
		UserID: primitive.NewObjectID().Hex(),
		Data:   "text",
		Title:  "title",
		Meta:   "meta",
	}

	successResult := &Note{
		ID:     sucessID.Hex(),
		UserID: successSource.UserID,
		Data:   successSource.Data,
		Title:  successSource.Title,
		Meta:   successSource.Meta,
	}
	type args struct {
		noteDocument NoteDocument
	}
	tests := []struct {
		name string
		args args
		want *Note
	}{
		{
			name: "1# success",
			args: args{
				noteDocument: *successSource,
			},
			want: successResult,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoteFromNoteDocument(tt.args.noteDocument); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteFromNoteDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}
