package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/users/domain/model"
	"reflect"
	"testing"
)

func Test_addUserInteractor_Handle(t *testing.T) {
	type fields struct {
		gateway AddNewUserGateway
	}
	type args struct {
		command AddUserCommand
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.User
		wantErr bool
	}{
		{
			name: "should create a user when username not exists",
			fields: fields{
				gateway: mockAddNewUserGateway{
					Username: "tester",
					Status:   "NEW",
					Exists:   false,
				},
			},
			args: args{
				command: AddUserCommand{
					Username: "tester",
				},
			},
			want: model.User{
				Username: "tester",
				Status:   "NEW",
			},
			wantErr: false,
		},
		{
			name: "should not create a user when username already exists",
			fields: fields{
				gateway: mockAddNewUserGateway{
					Username: "tester",
					Status:   "NEW",
					Exists:   true,
				},
			},
			args: args{
				command: AddUserCommand{
					Username: "tester",
				},
			},
			want:    model.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := addUserInteractor{
				gateway: tt.fields.gateway,
			}
			got, err := r.Handle(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handle() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type mockAddNewUserGateway struct {
	Username string
	Status   string
	Exists   bool
}

func (r mockAddNewUserGateway) AddNewUser(_ AddUserCommand) (model.User, error) {
	return model.User{
		Username: r.Username,
		Status:   r.Status,
	}, nil
}

func (r mockAddNewUserGateway) Exist(_ UsernameExistsQuery) (bool, error) {
	return r.Exists, nil
}
