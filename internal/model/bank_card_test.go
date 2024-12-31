package model

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestBankCardDocumentFromBankCard(t *testing.T) {
	sucessID := primitive.NewObjectID()
	errID := "123"

	successSource := &BankCard{
		ID:     sucessID.Hex(),
		UserID: primitive.NewObjectID().Hex(),
		Data:   123,
		Title:  "title",
		Meta:   "meta",
	}

	errSource := &BankCard{
		ID:     errID,
		UserID: primitive.NewObjectID().Hex(),
		Data:   123,
		Title:  "title",
		Meta:   "meta",
	}

	successResult := &BankCardDocument{
		ID:     sucessID,
		UserID: successSource.UserID,
		Title:  successSource.Title,
		Data:   successSource.Data,
		Meta:   successSource.Meta,
	}

	type args struct {
		card BankCard
	}
	tests := []struct {
		name    string
		args    args
		want    *BankCardDocument
		wantErr bool
	}{
		{
			name: "1# success",
			args: args{
				card: *successSource,
			},
			want:    successResult,
			wantErr: false,
		},
		{
			name: "2# error",
			args: args{
				card: *errSource,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BankCardDocumentFromBankCard(tt.args.card)
			if (err != nil) != tt.wantErr {
				t.Errorf("BankCardDocumentFromBankCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BankCardDocumentFromBankCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBankCardFromBankCardDocument(t *testing.T) {
	sucessID := primitive.NewObjectID()

	successSource := &BankCardDocument{
		ID:     sucessID,
		UserID: primitive.NewObjectID().Hex(),
		Data:   123,
		Title:  "title",
		Meta:   "meta",
	}

	successResult := &BankCard{
		ID:     sucessID.Hex(),
		UserID: successSource.UserID,
		Data:   successSource.Data,
		Title:  successSource.Title,
		Meta:   successSource.Meta,
	}

	type args struct {
		cardDoc BankCardDocument
	}
	tests := []struct {
		name string
		args args
		want *BankCard
	}{
		{
			name: "1# success",
			args: args{
				cardDoc: *successSource,
			},
			want: successResult,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BankCardFromBankCardDocument(tt.args.cardDoc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BankCardFromBankCardDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}
