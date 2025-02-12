package jwt

import "testing"

func TestCreateToken(t *testing.T) {
	
	tests := []struct {
		name    string
		id uint
		username string
		secretKey string
		wantErr bool
	}{
		{
			name:    "Success token creation",
			id:      1,
			username: "xprasetio",
			secretKey: "secrect_key",
			wantErr: false,
		},
		{
			name:    "Empty secret key",
			id:      1,
			username: "xprasetio",
			secretKey: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateToken(tt.id, tt.username, tt.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == "" {
				t.Errorf("CreateToken() returned empty token string")
			}
		})
	}
}
