package discount

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidatePercent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		percent int
		wantErr error
	}{
		{name: "zero", percent: 0},
		{name: "valid", percent: 10},
		{name: "maximum", percent: 100},
		{name: "below zero", percent: -1, wantErr: ErrInvalidPercent},
		{name: "above maximum", percent: 101, wantErr: ErrInvalidPercent},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := ValidatePercent(tt.percent)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestApply(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		price   int
		percent int
		want    int
		wantErr error
	}{
		{name: "regular discount", price: 200, percent: 15, want: 170},
		{name: "no discount", price: 200, percent: 0, want: 200},
		{name: "full discount", price: 200, percent: 100},
		{name: "zero price", price: 0, percent: 15},
		{name: "integer rounding", price: 99, percent: 10, want: 89},
		{name: "negative price", price: -1, percent: 10, wantErr: ErrNegativePrice},
		{name: "negative percent", price: 200, percent: -1, wantErr: ErrInvalidPercent},
		{name: "percent above maximum", price: 200, percent: 101, wantErr: ErrInvalidPercent},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Apply(tt.price, tt.percent)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestFinalPrice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		items   []Item
		want    int
		wantErr error
	}{
		{name: "empty list", items: []Item{}},
		{name: "nil list"},
		{
			name: "multiple items",
			items: []Item{
				{Price: 200, DiscountPercent: 10},
				{Price: 150, DiscountPercent: 20},
				{Price: 50, DiscountPercent: 0},
			},
			want: 350,
		},
		{
			name: "invalid price after valid item",
			items: []Item{
				{Price: 100, DiscountPercent: 10},
				{Price: -1, DiscountPercent: 0},
			},
			wantErr: ErrNegativePrice,
		},
		{
			name: "invalid discount after valid item",
			items: []Item{
				{Price: 100, DiscountPercent: 10},
				{Price: 200, DiscountPercent: 101},
			},
			wantErr: ErrInvalidPercent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := FinalPrice(tt.items)
			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}
