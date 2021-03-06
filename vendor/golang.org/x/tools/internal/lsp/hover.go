// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsp

import (
	"context"
	"fmt"

	"golang.org/x/tools/internal/lsp/protocol"
	"golang.org/x/tools/internal/lsp/source"
	"golang.org/x/tools/internal/span"
)

func (s *Server) hover(ctx context.Context, params *protocol.TextDocumentPositionParams) (*protocol.Hover, error) {
	uri := span.NewURI(params.TextDocument.URI)
	view := s.findView(ctx, uri)
	f, m, err := newColumnMap(ctx, view, uri)
	if err != nil {
		return nil, err
	}
	spn, err := m.PointSpan(params.Position)
	if err != nil {
		return nil, err
	}
	identRange, err := spn.Range(m.Converter)
	if err != nil {
		return nil, err
	}
	ident, err := source.Identifier(ctx, view, f, identRange.Start)
	if err != nil {
		return nil, err
	}
	decl, doc, err := ident.Hover(ctx, nil, s.enhancedHover)
	if err != nil {
		return nil, err
	}
	identSpan, err := ident.Range.Span()
	if err != nil {
		return nil, err
	}
	rng, err := m.Range(identSpan)
	if err != nil {
		return nil, err
	}
	return &protocol.Hover{
		Contents: markupContent(decl, doc, s.preferredContentFormat),
		Range:    &rng,
	}, nil
}

func markupContent(decl, doc string, kind protocol.MarkupKind) protocol.MarkupContent {
	result := protocol.MarkupContent{
		Kind: kind,
	}
	switch kind {
	case protocol.PlainText:
		result.Value = decl
	case protocol.Markdown:
		result.Value = "```go\n" + decl + "\n```"
	}
	if doc != "" {
		result.Value = fmt.Sprintf("%s\n%s", doc, result.Value)
	}
	return result
}
