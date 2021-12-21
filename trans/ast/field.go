package ast

import (
	"fmt"
	"strings"

	"github.com/dragonfax/java_converter/input/parser"
	"github.com/dragonfax/java_converter/tool"
	"github.com/dragonfax/java_converter/trans/node"
)

type FieldList []*Field

func (fl FieldList) String() string {
	return strings.Join(node.NodeListToStringList(fl), ",")
}

func (fl FieldList) Children() []node.Node {
	return node.ListOfNodesToNodeList(fl)
}

type Field struct {
	*BaseClassScope
	*VariableDeclNode

	Public    bool
	Transient bool
	Static    bool
}

func NewField(vardecl *VariableDeclNode) *Field {
	return &Field{BaseClassScope: NewClassScope(), VariableDeclNode: vardecl}
}

func (f *Field) Children() []node.Node {
	return nil
}

func NewFields(ctx *parser.FieldDeclarationContext) FieldList {
	members := make([]*Field, 0)

	typ := NewTypeNode(ctx.TypeType())

	for _, varDec := range ctx.VariableDeclarators().AllVariableDeclarator() {
		varDecCtx := varDec

		name := varDecCtx.VariableDeclaratorId().GetText()

		var init node.Node
		if varDecCtx.VariableInitializer() != nil {
			initCtx := varDecCtx.VariableInitializer()
			if initCtx.Expression() != nil {
				init = ExpressionProcessor(initCtx.Expression())
			} else if initCtx.ArrayInitializer() != nil {
				init = NewArrayLiteral(initCtx.ArrayInitializer())
			}
		}

		node := NewVariableDecl(typ, name, init)
		members = append(members, NewField(node))
	}

	return members
}

func (f *Field) Declaration() string {
	return fmt.Sprintf("%s %s", f.Name, f.Type)
}

func (f *Field) HasInitializer() bool {
	return !tool.IsNilInterface(f.Expression)
}

func (f *Field) Initializer() string {
	return fmt.Sprintf("%s = %s", f.Name, f.Expression)
}

func (f *Field) SetPublic(public bool) {
	f.Public = public
}

func (f *Field) SetStatic(static bool) {
	f.Static = static
}

func (f *Field) SetTransient(transient bool) {
	f.Transient = transient
}
