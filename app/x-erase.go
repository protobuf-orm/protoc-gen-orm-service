package app

import (
	"github.com/protobuf-orm/protoc-gen-orm-service/internal/ast"
)

func (w *fileWork) xRpcErase() ast.Rpc {
	return w.defineRpc(
		ast.Comment("Erase deletes a "+w.entity.Name()),
		ast.Rpc{
			Name:         "Erase",
			RequestType:  w.xMsgRef().Name,
			ResponseType: w.xMsgErase().Name,
		},
	)
}

// xMsgErase defines the response of the Erase RPC.
//
// # Why it is not Empty, which it was
//
// Because erasing what is not there **succeeds**, and has to: a caller
// cancelling something that may already be gone should not have to tell a race
// from a mistake, and callers were written against that. So the RPC cannot say
// "this call did it" by failing.
//
// Which left nothing that could say it at all, and once-only is a property
// built on exactly that answer. Anything single-use -- a nonce, a
// second-factor attempt, a magic link -- is spent by erasing it, and without
// this every concurrent presenter of one handle is told the same thing the
// winner is told. The server did the right thing at the row and then said
// nothing about it.
//
// It is not a version check and not a lock. It reports what the statement did:
// one row, or none.
//
// # Wire compatibility
//
// `Empty` has no fields, so a client generated before this decodes the new
// response as an `Empty` with one unknown field and carries on. What changes is
// the generated signature, which is a recompile rather than a redeploy.
func (w *fileWork) xMsgErase() ast.Message {
	return w.defineMsg("EraseResponse", func(m *ast.Message) {
		m.Body = []ast.MessageBody{
			ast.Comment("Erased is whether this call is the one that erased the row.\n" +
				"\n" +
				"False for a row that was already gone, was never there, or is out\n" +
				"of this caller's reach -- which are one answer on purpose, and the\n" +
				"reason the RPC does not fail instead."),
			ast.MessageField{
				Type:   "bool",
				Name:   "erased",
				Number: 1,
			},
		}
	})
}
