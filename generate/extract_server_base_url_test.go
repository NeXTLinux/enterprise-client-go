package main

import (
	"github.com/nextlinux/enterprise-client-go/pkg/enterprise"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExtractServerBaseURL(t *testing.T) {
	tests := []struct {
		name    string
		api     string
		want    string
		wantErr require.ErrorAssertionFunc
	}{
		{
			name: "enterprise api",
			api:  enterprise.Document(),
			want: "/enterprise",
		},
		{
			name:    "bad api",
			api:     "made-up",
			wantErr: require.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				tt.wantErr = require.NoError
			}
			got, err := extractServerBaseURL(tt.api)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
