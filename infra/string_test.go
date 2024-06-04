package infra

import (
	"testing"
)

func Test_SnakeCase(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "normal",
			args: "PabBBicNotify",
			want: "pab_b_bic_notify",
		},
		{
			name: "empty",
			args: "",
			want: "",
		},
		{
			name: "start with number",
			args: "2PabBicNotify",
			want: "x_2pab_bic_notify",
		},
		{
			name: "contain number",
			args: "PabB2BicNotify",
			want: "pab_b2bic_notify",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeCase(tt.args); got != tt.want {
				t.Errorf("SnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Kebab(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "normal",
			args: "PabBBicNotify",
			want: "pab-b-bic-notify",
		},
		{
			name: "empty",
			args: "",
			want: "",
		},
		{
			name: "start with number",
			args: "2PabBicNotify",
			want: "x-2pab-bic-notify",
		},
		{
			name: "contain number",
			args: "PabB2BicNotify",
			want: "pab-b2bic-notify",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Kebab(tt.args); got != tt.want {
				t.Errorf("Kebab() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SmallCamelCase(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "normal",
			args: "pab_b_bic_notify",
			want: "pabBBicNotify",
		},
		{
			name: "empty",
			args: "",
			want: "",
		},
		{
			name: "start with number",
			args: "x_2pab_b2bic_notify",
			want: "x_2PabB2BicNotify",
		},
		{
			name: "contain number",
			args: "pab_b2bic_notify",
			want: "pabB2BicNotify",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallCamelCase(tt.args); got != tt.want {
				t.Errorf("SmallCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PascalCase(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "normal",
			args: "pab_b_bic_notify",
			want: "PabBBicNotify",
		},
		{
			name: "empty",
			args: "",
			want: "",
		},
		{
			name: "start with number",
			args: "x_2pab_b2bic_notify",
			want: "X_2PabB2BicNotify",
		},
		{
			name: "start with delimiter",
			args: "_2pab_b2bic_notify",
			want: "X2PabB2BicNotify",
		},
		{
			name: "contain number",
			args: "pab_b2bic_notify",
			want: "PabB2BicNotify",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PascalCase(tt.args); got != tt.want {
				t.Errorf("PascalCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
