// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package render_test

import (
	"testing"

	"github.com/derailed/k9s/internal/client"
	"github.com/derailed/k9s/internal/model1"
	"github.com/derailed/k9s/internal/render"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReferenceRender(t *testing.T) {
	o := render.ReferenceRes{
		Namespace: "ns1",
		Name:      "blee",
		GVR:       client.SecGVR.String(),
	}

	var (
		ref = render.Reference{}
		r   model1.Row
	)
	require.NoError(t, ref.Render(o, "fred", &r))
	assert.Equal(t, "ns1/blee", r.ID)
	assert.Equal(t, model1.Fields{
		"ns1",
		"blee",
		client.SecGVR.String(),
	}, r.Fields)
}
