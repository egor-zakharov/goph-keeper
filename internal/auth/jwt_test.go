package auth

import "testing"

func TestBuildJWTString(t *testing.T) {
	type args struct {
		userID    string
		sessionID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				userID:    "1",
				sessionID: "1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := BuildJWTString(tt.args.userID, tt.args.sessionID)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildJWTString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetUserID(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				tokenString: "1",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserID(tt.args.tokenString); got != tt.want {
				t.Errorf("GetUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSessionID(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				tokenString: "1",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSessionID(tt.args.tokenString); got != tt.want {
				t.Errorf("GetSessionID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidToken(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				tokenString: "1",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsValidToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValidToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsValidToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
