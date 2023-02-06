package usecase

import (
	"github.com/CSalih/go-clean-architecture/internal/users/domain/model"
	"github.com/CSalih/go-clean-architecture/internal/users/infrastrucure/presenter"
	"reflect"
	"testing"
)

func Test_addUserInteractor_Handle(t *testing.T) {
	type fields struct {
		addNewUserGateway         AddNewUserGateway
		doesUsernameExistsGateway DoesUsernameExistsGateway
	}
	type args struct {
		command   AddUserCommand
		presenter presenter.Presenter
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
				addNewUserGateway: mockAddNewUserGateway{
					Username: "tester",
					Status:   "NEW",
				},
				doesUsernameExistsGateway: mockDoesUsernameExistsGateway{
					Exists: false,
				},
			},
			args: args{
				command: AddUserCommand{
					Username: "tester",
				},
				presenter: mockPresenter{
					Want: model.User{
						Username: "tester",
						Status:   "NEW",
					},
					Test:    t,
					WantErr: false,
				},
			},
		},
		{
			name: "should not create a user when username already exists",
			fields: fields{
				addNewUserGateway: mockAddNewUserGateway{
					Username: "tester",
					Status:   "NEW",
				},
				doesUsernameExistsGateway: mockDoesUsernameExistsGateway{
					Exists: true,
				},
			},
			args: args{
				command: AddUserCommand{
					Username: "tester",
				},
				presenter: mockPresenter{
					Want:    model.User{},
					Test:    t,
					WantErr: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := addUserInteractor{
				addNewUserGateway:         tt.fields.addNewUserGateway,
				doesUsernameExistsGateway: tt.fields.doesUsernameExistsGateway,
			}
			r.Handle(tt.args.command, tt.args.presenter)
		})
	}
}

type mockAddNewUserGateway struct {
	Username string
	Status   string
}

type mockDoesUsernameExistsGateway struct {
	Exists bool
}

func (r mockAddNewUserGateway) AddNewUser(_ AddUserCommand) (model.User, error) {
	return model.User{
		Username: r.Username,
		Status:   r.Status,
	}, nil
}

func (r mockDoesUsernameExistsGateway) Exist(_ UsernameExistsQuery) (bool, error) {
	return r.Exists, nil
}

type mockPresenter struct {
	Want    interface{}
	WantErr bool
	Test    *testing.T
}

func (p mockPresenter) OnSuccess(result interface{}) {
	if !reflect.DeepEqual(result, p.Want) {
		p.Test.Errorf("Handle() got = %v, want %v", result, p.Want)
	}
}
func (p mockPresenter) OnError(err error) {
	if (err != nil) != p.WantErr {
		p.Test.Errorf("Handle() error = %v, wantErr %v", err, p.WantErr)
		return
	}
}
