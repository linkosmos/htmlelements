package htmlelements

import "golang.org/x/net/html/atom"

// VoidElements - HTML elements that are void of has self closing tag
// http://www.w3.org/TR/html5/syntax.html#void-elements
var VoidElements = AtomSlice{
	atom.Area, atom.Base, atom.Br, atom.Col,
	atom.Embed, atom.Hr, atom.Img, atom.Input,
	atom.Keygen, atom.Link, atom.Meta, atom.Param,
	atom.Source, atom.Track, atom.Wbr,
}

// RawTextElements - can have text
// http://www.w3.org/TR/html5/syntax.html#raw-text-elements<F37>
var RawTextElements = AtomSlice{
	atom.Script, atom.Style,
}

// EscapableRawTextElements - can have text and character references, but the text must not contain an ambiguous ampersand.
// http://www.w3.org/TR/html5/syntax.html#escapable-raw-text-elements
var EscapableRawTextElements = AtomSlice{
	atom.Textarea, atom.Title,
}

// ForeignElements - whose start tag is marked as self-closing
// can't have any contents (since, again, as there's no end tag,
// no content can be put between the start tag and the end tag).
// Foreign elements whose start tag is not marked as self-closing
// can have text, character references, CDATA sections, other elements,
// and comments, but the text must not contain the character '<' (U+003C)
// or an ambiguous ampersand.
// http://www.w3.org/TR/html5/syntax.html#foreign-elements
var ForeignElements = AtomSlice{
	atom.Math, atom.Svg,
}

// NormalElements - All other allowed HTML elements are normal elements.
// http://www.w3.org/TR/html5/syntax.html#normal-elements
// generated;
var NormalElements = AtomSlice{
	atom.A,
	atom.Abbr,
	atom.Address,
	atom.Article,
	atom.Aside,
	atom.Audio,
	atom.B,
	atom.Bdi,
	atom.Bdo,
	atom.Blockquote,
	atom.Body,
	atom.Button,
	atom.Canvas,
	atom.Caption,
	atom.Cite,
	atom.Code,
	atom.Colgroup,
	atom.Command,
	atom.Data,
	atom.Datalist,
	atom.Dd,
	atom.Del,
	atom.Details,
	atom.Dfn,
	atom.Dialog,
	atom.Div,
	atom.Dl,
	atom.Dt,
	atom.Em,
	atom.Fieldset,
	atom.Figcaption,
	atom.Figure,
	atom.Footer,
	atom.Form,
	atom.H1,
	atom.H2,
	atom.H3,
	atom.H4,
	atom.H5,
	atom.H6,
	atom.Head,
	atom.Header,
	atom.Hgroup,
	atom.Html,
	atom.I,
	atom.Iframe,
	atom.Ins,
	atom.Kbd,
	atom.Label,
	atom.Legend,
	atom.Li,
	atom.Map,
	atom.Mark,
	atom.Menu,
	// atom.Menuitem,
	atom.Meter,
	atom.Nav,
	atom.Noscript,
	atom.Object,
	atom.Ol,
	atom.Optgroup,
	atom.Option,
	atom.Output,
	atom.P,
	atom.Pre,
	atom.Progress,
	atom.Q,
	atom.Rp,
	atom.Rt,
	atom.Ruby,
	atom.S,
	atom.Samp,
	atom.Script,
	atom.Section,
	atom.Select,
	atom.Small,
	atom.Span,
	atom.Strong,
	atom.Style,
	atom.Sub,
	atom.Summary,
	atom.Sup,
	atom.Table,
	atom.Tbody,
	atom.Td,
	// atom.Template,
	atom.Textarea,
	atom.Tfoot,
	atom.Th,
	atom.Thead,
	atom.Time,
	atom.Title,
	atom.Tr,
	atom.U,
	atom.Ul,
	atom.Var,
	atom.Video,
}
