package model

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBinaryDocumentFromBinary(t *testing.T) {
	sucessID := primitive.NewObjectID()
	errID := "123"

	successSource := &Binary{
		ID:     sucessID.Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Data:   []byte("123"),
		Title:  "title",
		Meta:   "meta",
	}

	errSource := &Binary{
		ID:     errID,
		UserID: primitive.NewObjectID().Hex(),
		Data:   []byte("123"),
		Title:  "title",
		Meta:   "meta",
	}

	successResult := &BinaryDocument{
		ID:     sucessID,
		UserID: successSource.UserID,
		Title:  successSource.Title,
		Data:   successSource.Data,
		Meta:   successSource.Meta,
	}
	type args struct {
		binary Binary
	}
	tests := []struct {
		name    string
		args    args
		want    *BinaryDocument
		wantErr bool
	}{
		{
			name: "1# success",
			args: args{
				binary: *successSource,
			},
			want:    successResult,
			wantErr: false,
		},
		{
			name: "2# error",
			args: args{
				binary: *errSource,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinaryDocumentFromBinary(tt.args.binary)
			if (err != nil) != tt.wantErr {
				t.Errorf("BinaryDocumentFromBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryDocumentFromBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryFromBinaryDocument(t *testing.T) {
	sucessID := primitive.NewObjectID()

	successSource := &BinaryDocument{
		ID:     sucessID,
		UserID: primitive.NewObjectID().Hex(),
		Data:   []byte("123"),
		Title:  "title",
		Meta:   "meta",
	}

	successResult := &Binary{
		ID:     sucessID.Hex(),
		UserID: successSource.UserID,
		Data:   successSource.Data,
		Title:  successSource.Title,
		Meta:   successSource.Meta,
	}
	type args struct {
		binaryDoc BinaryDocument
	}
	tests := []struct {
		name string
		args args
		want *Binary
	}{
		{
			name: "1# success",
			args: args{
				binaryDoc: *successSource,
			},
			want: successResult,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinaryFromBinaryDocument(tt.args.binaryDoc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryFromBinaryDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}
