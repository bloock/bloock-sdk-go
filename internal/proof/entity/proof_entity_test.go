package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	t.Run("Given a valid Proof, should return true", func(t *testing.T) {
		p := NewProof([]string{"02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5"},
		[]string{"bb6986853646d083929d1d92638f3d4741a3b7149bd2b63c6bfedd32e3c684d3",
				"0616067c793ac533815ae2d48d785d339e0330ce5bb5345b5e6217dd9d1dbeab",
				"68b8f6b25cc700e64ed3e3d33f2f246e24801f93d29786589fbbab3b11f5bcee"},
				"0004000600060005", "bfdf7000")

		r := IsValid(p)
		assert.True(t, r)
	})

	t.Run("Given a minimalist test, should return true", func(t *testing.T) {
		p := NewProof([]string{"02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5"},
		[]string{},
		"0004", "bf")

		r := IsValid(p)
		assert.True(t, r)
	})

	t.Run("Given invalid hex leaves, should return false", func(t *testing.T) {
		p := NewProof([]string{"02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aeeg"},
		[]string{"bb6986853646d083929d1d92638f3d4741a3b7149bd2b63c6bfedd32e3c684d3",
			"0616067c793ac533815ae2d48d785d339e0330ce5bb5345b5e6217dd9d1dbeab",
			"68b8f6b25cc700e64ed3e3d33f2f246e24801f93d29786589fbbab3b11f5bcee"},
			"0004000600060005", "bfdf7000")

		r := IsValid(p)
		assert.False(t, r)
	})

	t.Run("Given invalid hex nodes, should return false", func(t *testing.T) {
		p := NewProof([]string{"02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5"},
			[]string{"bb6986853646d083929d1d92638f3d4741a3b7149bd2b63c6bfedd32e3c684d3",
				"0616067c793ac533815ae2d48d785d339e0330ce5bb5345b5e6217dd9d1dbeag",
				"68b8f6b25cc700e64ed3e3d33f2f246e24801f93d29786589fbbab3b11f5bcee"},
			"0004000600060005", "bfdf7000")

		r := IsValid(p)
		assert.False(t, r)
	})

	t.Run("Given an invalid short bitmap, should return false", func(t *testing.T) {
		p := NewProof([]string{"02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5"},
			[]string{"bb6986853646d083929d1d92638f3d4741a3b7149bd2b63c6bfedd32e3c684d3",
				"0616067c793ac533815ae2d48d785d339e0330ce5bb5345b5e6217dd9d1dbeab",
				"68b8f6b25cc700e64ed3e3d33f2f246e24801f93d29786589fbbab3b11f5bcee",
				"bb6986853646d083929d1d92638f3d4741a3b7149bd2b63c6bfedd32e3c684d3",
				"0616067c793ac533815ae2d48d785d339e0330ce5bb5345b5e6217dd9d1dbeab",
				"68b8f6b25cc700e64ed3e3d33f2f246e24801f93d29786589fbbab3b11f5bcee",
				"bb6986853646d083929d1d92638f3d4741a3b7149bd2b63c6bfedd32e3c684d3",
				"0616067c793ac533815ae2d48d785d339e0330ce5bb5345b5e6217dd9d1dbeab",
				"68b8f6b25cc700e64ed3e3d33f2f246e24801f93d29786589fbbab3b11f5bcee"},
			"0004000600060005000600060005000600060005", "bf")

		r := IsValid(p)
		assert.False(t, r)
	})

	t.Run("Given a invalid short depth, should return false", func(t *testing.T) {
		p := NewProof([]string{"02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5"},
		[]string{"bb6986853646d083929d1d92638f3d4741a3b7149bd2b63c6bfedd32e3c684d3",
				"0616067c793ac533815ae2d48d785d339e0330ce5bb5345b5e6217dd9d1dbeab",
				"68b8f6b25cc700e64ed3e3d33f2f246e24801f93d29786589fbbab3b11f5bcee"},
			"000400060006000", "bfdf7000")

		r := IsValid(p)
		assert.False(t, r)
	})

	t.Run("Given a invalid long depth, should return false", func(t *testing.T) {
		p := NewProof([]string{"02aae7e86eb50f61a62083a320475d9d60cbd52749dbf08fa942b1b97f50aee5"},
			[]string{"bb6986853646d083929d1d92638f3d4741a3b7149bd2b63c6bfedd32e3c684d3",
				"0616067c793ac533815ae2d48d785d339e0330ce5bb5345b5e6217dd9d1dbeab",
				"68b8f6b25cc700e64ed3e3d33f2f246e24801f93d29786589fbbab3b11f5bcee"},
			"0004000600060", "bfdf7000")

		r := IsValid(p)
		assert.False(t, r)
	})


}