// Copyright (c) 2014 Alex Diaconu.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// Package webkitdomgtk provides WebKit2DOMGTK+ bindings for Go.
package webkitdomgtk

// #cgo pkg-config: webkit2gtk-3.0
//
// #include <stdint.h>
// #include <stdlib.h>
// #include <string.h>
//
// #include <webkitdom/webkitdom.h>
// #include "dom.go.h"
import "C"

import (
    "runtime"
    "errors"
    "unsafe"
    "github.com/conformal/gotk3/glib"
)

func init() {
    tm := []glib.TypeMarshaler{
        // Objects/Interfaces
        {glib.Type(C.webkit_dom_node_get_type()), marshalNode},
        {glib.Type(C.webkit_dom_object_get_type()), marshalObject},
        {glib.Type(C.webkit_dom_node_list_get_type()), marshalNodeList},
        {glib.Type(C.webkit_dom_document_get_type()), marshalDocument},
        {glib.Type(C.webkit_dom_attr_get_type()), marshalAttr},
        {glib.Type(C.webkit_dom_document_get_type()), marshalDocument},
        {glib.Type(C.webkit_dom_element_get_type()), marshalElement},
        {glib.Type(C.webkit_dom_html_element_get_type()), marshalHTMLElement},
        {glib.Type(C.webkit_dom_html_collection_get_type()), marshalHTMLCollection},
        {glib.Type(C.webkit_dom_style_sheet_get_type()), marshalStyleSheet},
        {glib.Type(C.webkit_dom_style_sheet_list_get_type()), marshal},
        {glib.Type(C.webkit_dom_character_data_get_type()), marshalCharacterData},
        {glib.Type(C.webkit_dom_processing_instruction_get_type()), marshalProcessingInstruction},
        {glib.Type(C.webkit_dom_text_get_type()), marshalText},
        {glib.Type(C.webkit_dom_file_list_get_type()), marshalFileList},
        {glib.Type(C.webkit_dom_file_get_type()), marshalFile},
        {glib.Type(C.webkit_dom_blob_get_type()), marshalBlob},
        {glib.Type(C.webkit_dom_media_list_get_type()), marshalMediaList},
        {glib.Type(C.webkit_dom_css_style_sheet_get_type()), marshalCSSStyleSheet},
        {glib.Type(C.webkit_dom_css_rule_get_type()), marshalCSSRule},
        {glib.Type(C.webkit_dom_css_rule_list_get_type()), marshalCSSRuleList},
        {glib.Type(C.webkit_dom_css_style_declaration_get_type()), marshalCSSStyleDeclaration},
        {glib.Type(C.webkit_dom_css_value_get_type()), marshalCSSValue},
        {glib.Type(C.webkit_dom_event_get_type()), marshalEvent},
        {glib.Type(C.webkit_dom_xpath_expression_get_type()), marshalXPathExpression},
        {glib.Type(C.webkit_dom_xpath_result_get_type()), marshalXPathResult},
        {glib.Type(C.webkit_dom_window_get_type()), marshalWindow},
        {glib.Type(C.webkit_dom_html_anchor_element_get_type()), marshalHTMLAnchorElement},
        {glib.Type(C.webkit_dom_html_applet_element_get_type()), marshalHTMLAppletElement},
        {glib.Type(C.webkit_dom_html_area_element_get_type()), marshalHTMLAreaElement},
        {glib.Type(C.webkit_dom_html_br_element_get_type()), marshalHTMLBRElement},
        {glib.Type(C.webkit_dom_html_base_element_get_type()), marshalHTMLBaseElement},
        {glib.Type(C.webkit_dom_html_title_element_get_type()), marshalHTMLTitleElement},
        {glib.Type(C.webkit_dom_html_text_area_element_get_type()), marshalHTMLTextAreaElement},
        {glib.Type(C.webkit_dom_html_head_element_get_type()), marshalHTMLHeadElement},
        {glib.Type(C.webkit_dom_html_link_element_get_type()), marshalHTMLLinkElement},
        {glib.Type(C.webkit_dom_html_body_element_get_type()), marshalHTMLBodyElement},
        {glib.Type(C.webkit_dom_html_button_element_get_type()), marshalHTMLButtonElement},
        {glib.Type(C.webkit_dom_html_div_element_get_type()), marshalHTMLDivElement},
        {glib.Type(C.webkit_dom_html_document_get_type()), marshalHTMLDocument},
        {glib.Type(C.webkit_dom_html_frame_element_get_type()), marshalHTMLFrameElement},
        {glib.Type(C.webkit_dom_html_frame_set_element_get_type()), marshalHTMLFrameSetElement},
        {glib.Type(C.webkit_dom_html_embed_element_get_type()), marshalHTMLEmbedElement},
        {glib.Type(C.webkit_dom_html_field_set_element_get_type()), marshalHTMLFieldSetElement},
        {glib.Type(C.webkit_dom_html_input_element_get_type()), marshalHTMLInputElement},
        {glib.Type(C.webkit_dom_html_label_element_get_type()), marshalHTMLLabelElement},
        {glib.Type(C.webkit_dom_html_image_element_get_type()), marshalHTMLImageElement},
        {glib.Type(C.webkit_dom_html_form_element_get_type()), marshalHTMLFormElement},
        {glib.Type(C.webkit_dom_html_meta_element_get_type()), marshalHTMLMetaElement},
        {glib.Type(C.webkit_dom_html_object_element_get_type()), marshalHTMLObjectElement},
        {glib.Type(C.webkit_dom_html_opt_group_element_get_type()), marshalHTMLOptGroupElement},
        {glib.Type(C.webkit_dom_html_option_element_get_type()), marshalHTMLOptionElement},
        {glib.Type(C.webkit_dom_html_options_collection_get_type()), marshalHTMLOptionsCollection},
        {glib.Type(C.webkit_dom_html_script_element_get_type()), marshalHTMLScriptElement},
        {glib.Type(C.webkit_dom_html_select_element_get_type()), marshalHTMLSelectElement},
    }
    glib.RegisterGValueMarshalers(tm)
}

type Object struct {
    *glib.Object
}

func marshalObject(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapObject(obj), nil
}

func wrapObject(obj *glib.Object) *Object {
    return &Object{obj}
}

func (o *Object) native() *C.WebKitDOMObject {
    if o == nil || o.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(o.GObject)
    return C.toWebKitDOMObject(p)
}

func newObject(p unsafe.Pointer) *Object {
    return wrapObject(newGlibObject(p))
}

type Node struct {
    Object
}

func marshalNode(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapNode(obj), nil
}

func wrapNode(obj *glib.Object) *Node {
    return &Node{Object{obj}}
}

func (n *Node) native() *C.WebKitDOMNode {
    if n == nil || n.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(n.GObject)
    return C.toWebKitDOMNode(p)
}

func newNode(p unsafe.Pointer) *Node {
    return wrapNode(newGlibObject(p))
}

// ReplaceChild wraps webkit_dom_node_replace_child().
func (n *Node) ReplaceChild(newChild, oldChild *Node) (*Node, error) {
    var err *C.GError = nil
    C.webkit_dom_node_replace_child(n.native(), newChild.native(), oldChild.native(), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return oldChild, nil
}


// RemoveChild wraps webkit_dom_node_remove_child().
func (n *Node) RemoveChild(oldChild *Node) (*Node, error) {
    var err *C.GError = nil
    C.webkit_dom_node_remove_child(n.native(), oldChild.native(), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return oldChild, nil
}

// AppendChild wraps webkit_dom_node_append_child().
//
// http://webkitgtk.org/reference/webkitdomgtk/2.4.6/WebKitDOMNode.html#webkit-dom-node-append-child
// http://www.w3.org/TR/DOM-Level-3-Core/core.html#ID-184E7107
func (n *Node) AppendChild(newChild *Node) (*Node, error) {
    var err *C.GError = nil
    C.webkit_dom_node_append_child(n.native(), newChild.native(), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newChild, nil
}

// HasChildNodes wraps webkit_dom_node_has_child_nodes().
func (n *Node) HasChildNodes() bool {
    return gobool(C.webkit_dom_node_has_child_nodes(n.native()))
}

func newGlibObject(p unsafe.Pointer) *glib.Object {
    obj := &glib.Object{GObject: glib.ToGObject(p)}
    obj.RefSink()
    runtime.SetFinalizer(obj, (*glib.Object).Unref)
    return obj
}

// GLib <-> Go boolean conversion
func gboolean(b bool) C.gboolean {
    if b {
        return C.gboolean(1)
    }
    return C.gboolean(0)
}

func gobool(b C.gboolean) bool {
    return b != 0
}

// Clone wraps webkit_dom_node_clone_node().
func (n *Node) Clone(deep bool) *Node {
    c := C.webkit_dom_node_clone_node(n.native(), gboolean(deep))
    if c == nil {
        return nil
    }
    return newNode(unsafe.Pointer(c))
}

// Normalize wraps webkit_dom_node_normalize().
func (n *Node) Normalize() {
    C.webkit_dom_node_normalize(n.native())
}

// IsSupported wraps webkit_dom_node_is_supported().
func (n *Node) IsSupported(feature, version string) bool {
    f := C.CString(feature)
    defer C.free(unsafe.Pointer(f))
    v := C.CString(version)
    defer C.free(unsafe.Pointer(v))
    b := C.webkit_dom_node_is_supported(n.native(), (*C.gchar)(f), (*C.gchar)(v))
    return gobool(b)
}

// TODO: IsSameNode

// IsEqual wraps webkit_dom_node_is_equal_node().
func (n *Node) IsEqual(other *Node) bool {
    b := C.webkit_dom_node_is_equal_node(n.native(), other.native())
    return gobool(b)
}

// LookupPrefix wraps webkit_dom_node_lookup_prefix().)
func (n *Node) LookupPrefix(namespaceURI string) string {
    uri := C.CString(namespaceURI)
    defer C.free(unsafe.Pointer(uri))
    s := C.webkit_dom_node_lookup_prefix(n.native(), (*C.gchar)(uri))
    return C.GoString((*C.char)(s))
}

// LookupNamespaceURI wraps webkit_dom_node_lookup_namespace_uri().)
func (n *Node) LookupNamespaceURI(prefix string) string {
    p := C.CString(prefix)
    defer C.free(unsafe.Pointer(p))
    s := C.webkit_dom_node_lookup_namespace_uri(n.native(), (*C.gchar)(p))
    return C.GoString((*C.char)(s))
}

// IsDefaultNamespace wraps webkit_dom_node_is_default_namespace().
func (n *Node) IsDefaultNamespace(namespaceURI string) bool {
    uri := C.CString(namespaceURI)
    defer C.free(unsafe.Pointer(uri))
    b := C.webkit_dom_node_is_default_namespace(n.native(), (*C.gchar)(uri))
    return gobool(b)
}

// CompareDocumentPosition wraps webkit_dom_node_compare_document_position().
func (n *Node) CompareDocumentPosition(other *Node) DocumentPosition {
    dp := C.webkit_dom_node_compare_document_position(n.native(), other.native())
    return DocumentPosition(dp)
}

// Contains wraps webkit_dom_node_contains().
func (n *Node) Contains(other *Node) bool {
    b := C.webkit_dom_node_contains(n.native(), other.native())
    return gobool(b)
}

// Name wraps webkit_dom_node_get_node_name().
func (n *Node) Name() string {
    s := C.webkit_dom_node_get_node_name(n.native())
    return C.GoString((*C.char)(s))
}

// Value wraps webkit_dom_node_get_node_value().
func (n *Node) Value() string {
    s := C.webkit_dom_node_get_node_value(n.native())
    return C.GoString((*C.char)(s))
}

// SetValue wraps webkit_dom_node_set_node_value().
//
// http://webkitgtk.org/reference/webkitdomgtk/2.4.6/WebKitDOMNode.html#webkit-dom-node-set-node-value
func (n *Node) SetValue(value string) error {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    var err *C.GError = nil
    C.webkit_dom_node_set_node_value(n.native(), (*C.gchar)(v), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

// Type wraps webkit_dom_node_get_node_type().
func (n *Node) Type() Type {
    t := C.webkit_dom_node_get_node_type(n.native())
    return Type(t)
}

// ParentNode wraps webkit_dom_node_get_parent_node().
func (n *Node) ParentNode() *Node {
    pn := C.webkit_dom_node_get_parent_node(n.native())
    if pn == nil {
        return nil
    }
    return newNode(unsafe.Pointer(pn))
}

// ChildNodes wraps webkit_dom_node_get_child_nodes().
func (n *Node) ChildNodes() *NodeList {
    cl := C.webkit_dom_node_get_child_nodes(n.native())
    return newNodeList(unsafe.Pointer(cl))
}

// FirstChild wraps webkit_dom_node_get_first_child().
func (n *Node) FirstChild() *Node {
    f := C.webkit_dom_node_get_first_child(n.native())
    if f == nil {
        return nil
    }
    return newNode(unsafe.Pointer(f))
}

// LastChild wraps webkit_dom_node_get_last_child().
func (n *Node) LastChild() *Node {
    l := C.webkit_dom_node_get_first_child(n.native())
    if l == nil {
        return nil
    }
    return newNode(unsafe.Pointer(l))
}

// PreviousSibling wraps webkit_dom_node_get_previous_sibling().
func (n *Node) PreviousSibling() *Node {
    ps := C.webkit_dom_node_get_previous_sibling(n.native())
    if ps == nil {
        return nil
    }
    return newNode(unsafe.Pointer(ps))
}

// NextSibling wraps webkit_dom_node_get_next_sibling().
func (n *Node) NextSibling() *Node {
    ns := C.webkit_dom_node_get_next_sibling(n.native())
    if ns == nil {
        return nil
    }
    return newNode(unsafe.Pointer(ns))
}

// NamespaceURI wraps webkit_dom_node_get_namespace_uri().
func (n *Node) NamespaceURI() string {
    uri := C.webkit_dom_node_get_namespace_uri(n.native())
    if uri == nil {
        return ""
    }
    return C.GoString((*C.char)(uri))
}

// Prefix wraps webkit_dom_node_get_prefix().
func (n *Node) Prefix() string {
    p := C.webkit_dom_node_get_prefix(n.native())
    if p == nil {
        return ""
    }
    return C.GoString((*C.char)(p))
}

// SetPrefix wraps webkit_dom_node_set_prefix().
func (n *Node) SetPrefix(value string) error {
    p := C.CString(value)
    defer C.free(unsafe.Pointer(p))
    var err *C.GError = nil
    C.webkit_dom_node_set_prefix(n.native(), (*C.gchar)(p), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

// LocalName wraps webkit_dom_node_get_local_name().
func (n *Node) LocalName() string {
    ln := C.webkit_dom_node_get_local_name(n.native())
    if ln == nil {
        return ""
    }
    return C.GoString((*C.char)(ln))
}

// BaseURI wraps webkit_dom_node_get_base_uri().
func (n *Node) BaseURI() string {
    uri := C.webkit_dom_node_get_base_uri(n.native())
    if uri == nil {
        return ""
    }
    return C.GoString((*C.char)(uri))
}

// TextContent wraps webkit_dom_node_get_text_content().
func (n *Node) TextContent() string {
    t := C.webkit_dom_node_get_text_content(n.native())
    if t == nil {
        return ""
    }
    return C.GoString((*C.char)(t))
}

// SetTextContent wraps webkit_dom_node_set_text_content().
func (n *Node) SetTextContent(content string) error {
    c := C.CString(content)
    defer C.free(unsafe.Pointer(c))
    var err *C.GError = nil
    C.webkit_dom_node_set_text_content(n.native(), (*C.gchar)(c), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (n *Node) ParentElement() *Element {
    e := C.webkit_dom_node_get_parent_element(n.native())
    return newElement(unsafe.Pointer(e))
}

/*
 * NodeList
 */

type NodeList struct {
    Object
}

func marshalNodeList(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapNodeList(obj), nil
}

func newNodeList(p unsafe.Pointer) *NodeList {
    return wrapNodeList(newGlibObject(p))
}

func wrapNodeList(obj *glib.Object) *NodeList {
    return &NodeList{Object{obj}}
}

func (l *NodeList) native() *C.WebKitDOMNodeList {
    if l == nil || l.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(l.GObject)
    return C.toWebKitDOMNodeList(p)
}

// Len wraps webkit_dom_node_list_get_length().
func (l *NodeList) Len() int {
    return int(C.webkit_dom_node_list_get_length(l.native()))
}

// Item wraps webkit_dom_node_list_item().
func (l *NodeList) Item(index int) *Node {
    n := C.webkit_dom_node_list_item(l.native(), (C.gulong)(index))
    return newNode(unsafe.Pointer(n))
}

type NodeMap struct {
    Object
}

func newNodeMap(p unsafe.Pointer) *NodeMap {
    return wrapNodeMap(obj)
}

func marshalNodeMap(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapNodeMap(obj), nil
}

func wrapNodeMap(obj *glib.Object) *NodeMap {
    return &NodeMap{Object{obj}}
}

func (m *NodeMap) native() *C.WebKitDOMNamedNodeMap {
    if m == nil || m.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(m.GObject)
    return C.toWebKitDOMNamedNodeMap(p)
}

func (m *NodeMap) toWebKitDOMNamedNodeMap() *C.WebKitDOMNamedNodeMap {
    return m.native();
}

func (m *NodeMap) Node(name string) *Node {
    s := C.CString(name)
    defer C.free(unsafe.Pointer(s))
    n := C.webkit_dom_named_node_map_get_named_item(m.native(), (*C.gchar)(s))
    return newNode(unsafe.Pointer(n))
}

func (m *NodeMap) Set(node *Node) *Node {
    n := C.webkit_dom_named_node_map_set_named_item(m.native(), node.native(), nil)
    if n != nil {
        return newNode(unsafe.Pointer(n))
    }
    return nil
}

func (m *nodeMap) Len() int {
    return int(C.webkit_dom_named_node_map_get_length(m.native()))
}

func (m *nodeMap) At(index int) *Node {
    n := C.webkit_dom_named_node_map_item(m.native(), (C.gulong)(index))
    if n != nil {
        return newNode(unsafe.Pointer(n))
    }
    return nil
}

/*
 * NodeIterator
 */

type NodeIterator struct {
    Object
}

func marshalNodeIterator(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapNodeIterator(obj), nil
}

func wrapNodeIterator(obj *glib.Object) *NodeIterator {
    return &NodeIterator{Object{obj}}
}

func (n *NodeIterator) native() *C.WebKitDOMNodeIterator {
    if n == nil || n.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(n.GObject)
    return C.toWebKitDOMNodeIterator(p)
}

func newNodeIterator(p unsafe.Pointer) *NodeIterator {
    return wrapNodeIterator(newGlibObject(p))
}

func (i *NodeIterator) NextNode() (*Node, error) {
    var err *C.GError
    n := C.webkit_dom_node_iterator_next_node(i.native(), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newNode(unsafe.Pointer(n)), nil
}

func (i *NodeIterator) PreviousNode() (*Node, error) {
    var err *C.GError
    n := C.webkit_dom_node_iterator_next_node(i.native(), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newNode(unsafe.Pointer(n)), nil
}

func (i *NodeIterator) Detach() {
    C.webkit_dom_node_iterator_detach(i.native())
}

func (i *NodeIterator) Root() *Node {
    n := C.webkit_dom_node_iterator_get_root(i.native())
    return newNode(unsafe.Pointer(n))
}

func (i *NodeIterator) WhatToShow() uint {
    return uint(C.webkit_dom_node_iterator_get_what_to_show(i.native()))
}

func (i *NodeIterator) Filter() *NodeFilter {
    f := C.webkit_dom_node_iterator_get_filter(i.native())
    return newNodeFilter(unsafe.Pointer(f))
}

func (i *NodeIterator) ExpandEntityReferences() bool {
    return gobool(C.webkit_dom_node_iterator_get_expand_entity_references(i.native()))
}

func (i *NodeIterator) ReferenceNode() *Node {
    n := C.webkit_dom_node_iterator_get_reference_node(i.native())
    return newNode(unsafe.Pointer(n))
}

func (i *NodeIterator) PointerBeforeReferenceNode() bool {
    return gobool(C.webkit_dom_node_iterator_get_pointer_before_reference_node(i.native()))
}

/*
 * NodeFilter
 */
type NodeFilter struct {
    // TODO
}

/*
 * Document
 */
type Document struct {
    Node
}

func marshalDocument(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapDocument(obj), nil
}

func wrapDocument(obj *glib.Object) *Document {
    return &Document{Node{Object{obj}}}
}

func (d *Document) native() *C.WebKitDOMDocument {
    if d == nil || d.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(d.GObject)
    return C.toWebKitDOMDocument(p)
}

func newDocument(p unsafe.Pointer) *Document {
    return wrapDocument(newGlibObject(p))
}

func (d *Document) CreateElement(tag string) error {
    t := C.CString(tag)
    defer C.free(unsafe.Pointer(t))
    var err *C.GError
    e := C.webkit_dom_document_create_element(d.native(), (*C.gchar)(t), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newElement(unsafe.Pointer(e)), nil
}

func (d *Document) CreateTextNode(data string) *Text {
    d := C.CString(data)
    defer C.free(unsafe.Pointer(d))
    t := C.webkit_dom_document_create_text_node(d.native(), (*C.gchar)(d))
    return newText(unsafe.Pointer(t))
}

// func (d *Document) CreateComment(data string) *Comment {
//     cd := C.CString(data)
//     defer C.free(unsafe.Pointer(cd))
//     c := C.webkit_dom_document_create_comment(d.native(), (*C.gchar)(cd))
//     return newComment(unsafe.Pointer(c))
// }

// func (d *Document) CreateCDataSection(data string) (*CDataSection, error) {
//     cd := C.CString(data)
//     defer C.free(unsafe.Pointer(cd))
//     var err *C.GError
//     cds := C.webkit_dom_document_create_cdata_section(d.native(), (*C.gchar)(cd), &err)
//     if err != nil {
//         defer C.g_error_free(err)
//         return nil, errors.New(C.GoString((*C.char)(err.message)))
//     }
//     return newCDataSection(unsafe.Pointer(cds)), nil
// }
//

func (d *Document) CreateProcessingInstruction(target, data string) (*ProcessingInstruction, error) {
    t := C.CString(target)
    defer C.free(unsafe.Pointer(t))
    pid := C.CString(data)
    defer C.free(unsafe.Pointer(pid))
    var err *C.GError
    pi := C.webkit_dom_document_create_processing_instruction(d.native(), (*C.gchar)(pid), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newProcessingInstruction(unsafe.Pointer(pi)), nil
}

func (d *Document) CreateAttribute(name string) (*Attr, error) {
    n := C.CString(name)
    defer C.free(unsafe.Pointer(n))
    var err *C.GError
    a := C.webkit_dom_document_create_attribute(d.native(), (*C.gchar)(n), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newAttr(unsafe.Pointer(a))
}

func (d *Document) CreateEntityReference(name string) error {
    n := C.CString(name)
    defer C.free(unsafe.Pointer(n))
    er := C.webkit_dom_document_create_entity_reference(d.native(), (*C.gchar)(n), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newEntityReference(unsafe.Pointer(er)), nil
}

func (d *Document) ElementsByTagName(tag string) *NodeList {
    t := C.CString(tag)
    defer C.free(unsafe.Pointer(t))
    l := C.webkit_dom_document_get_elements_by_tag_name(d.native(), (*C.gchar)(t))
    return newNodeList(unsafe.Pointer(l))
}

func (d *Document) ImportNode(node *Node, deep bool) (*Node, error) {
    var err *C.GError
    in := C.webkit_dom_document_import_node(d.native(), node.native(), gboolean(deep), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    // This should be a clone of the node we take as a parameter to this function.
    return newNode(unsafe.Pointer(in)), nil
}

// webkit_dom_document_create_element_ns
// webkit_dom_document_create_attribute_ns
// webkit_dom_document_get_elements_by_tag_name_ns

func (d *Document) ElementByID(id string) *Element {
    i := C.CString(id)
    defer C.free(unsafe.Pointer(i))
    e := C.webkit_dom_document_get_element_by_id(d.native(), (*C.gchar)(i))
    return newElement(unsafe.Pointer(e))
}

func (d *Document) AdoptNode(source *Node) (*Node, error) {
    var err *C.GError
    C.webkit_dom_document_adopt_node(d.native(), source.native(), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    // Apparently, this doesn't create a copy of the source node, so we can return that here.
    return source, nil
}

func (d *Document) CreateEvent(event string) (*Event, error) {
    es := C.CString(event)
    defer C.free(unsafe.Pointer(es))
    var err *C.GError
    e := C.webkit_dom_document_create_event(d.native(), (*C.gchar)(es), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newEvent(e), nil
}

func (d *Document) ElementsByName(name string) *NodeList {
    n := C.CString(name)
    defer C.free(unsafe.Pointer(n))
    nl := C.webkit_dom_document_get_elements_by_name(d.native(), (*C.gchar)(n))
    return newNodeList(unsafe.Pointer(nl))
}

func (d *Document) ElementFromPoint(x, y int) *Element {
    e := C.webkit_dom_document_element_from_point(d.native(), C.glong(x), C.glong(y))
    return newElement(unsafe.Pointer(e))
}

func (d *Document) CreateCSSStyleDeclaration() *CSSStyleDeclaration {
    csd := C.webkit_dom_document_create_css_style_declaration(d.native())
    return newCSSStyleDeclaration(unsafe.Pointer(csd))
}

func (d *Document) ElementsByClassName(tag string) *NodeList {
    t := C.CString(tag)
    defer C.free(unsafe.Pointer(t))
    nl := C.webkit_dom_document_get_elements_by_class_name(d.native(), (*C.gchar)(t))
    return newNodeList(unsafe.Pointer(nl))
}

func (d *Document) HasFocus() bool {
    return gobool(C.webkit_dom_document_has_focus(d.native()))
}

func (d *Document) QuerySelector(selectors string) (*Element, error) {
    s := C.CString(selectors)
    defer C.free(unsafe.Pointer(s))
    var err *C.GError
    e := C.webkit_dom_document_query_selector(d.native(), (*C.gchar)(s), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newElement(unsafe.Pointer(e)), nil
}

func (d *Document) QuerySelectorAll(selectors string) (*NodeList, error) {
    s := C.CString(selectors)
    defer C.free(unsafe.Pointer(s))
    nl := C.webkit_dom_document_query_selector_all(d.native(), (*C.gchar)(s), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newNodeList(unsafe.Pointer(nl)), nil
}

// TODO
// webkit_dom_document_get_doctype ()
// webkit_dom_document_get_implementation ()

func (d *Document) DocumentElement() *Element {
    e := C.webkit_dom_document_get_document_element(d.native())
    return newElement(unsafe.Pointer(e))
}

func (d *Document) URI() string {
    u := C.webkit_dom_document_get_document_uri(d.native())
    return C.GoString((*C.char)(u))
}

func (d *Document) SetURI(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_document_set_document_uri(d.native(), (*C.gchar)(v))
}

// TODO
// WebKitDOMDOMWindow *	webkit_dom_document_get_default_view ()
func (d *Document) StyleSheets() *StyleSheetList {
    ssl := C.webkit_dom_document_get_style_sheets(d.native())
    return newStyleSheetList(unsafe.Pointer(ssl))
}

func (d *Document) Title() string {
    t := C.webkit_dom_document_get_title(d.native())
    return C.GoString((*C.char)(t))
}

func (d *Document) SetTitle(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_document_set_title(d.native(), (*C.gchar)(v))
}

func (d *Document) Referrer() string {
    r := C.webkit_dom_document_get_referrer(d.native())
    return C.GoString((*C.char)(r))
}

func (d *Document) Domain() string {
    domain := C.webkit_dom_document_get_domain(d.native())
    return C.GoString((*C.char)(domain))
}

func (d *Document) URL() string {
    u := C.webkit_dom_document_get_url()
    return C.GoString((*C.char)(u))
}

func (d *Document) Cookie() (string, error) {
    var err *C.GError
    c := C.webkit_dom_document_get_cookie(d.native(), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return C.GoString((*C.char)(c)), nil
}

func (d *Document) SetCookie(value string) error {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_document_set_cookie(d.native(), (*C.gchar)(v), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (d *Document) Body() *HTMLElement {
    b := C.webkit_dom_document_get_body(d.native())
    return newHTMLElement(unsafe.Pointer(b))
}

func (d *Document) SetBody(value string) error {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    var err *C.GError
    C.webkit_dom_document_set_body(d.native(), (*C.gchar)(v), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (d *Document) Head() *HTMLHeadElement {
    h := C.webkit_dom_document_get_head(d.native())
    return newHTMLHeadElement(unsafe.Pointer(he))
}

func (d *Document) Images() *HTMLCollection {
    is := C.webkit_dom_document_get_images(d.native())
    return newHTMLCollection(unsafe.Pointer(is))
}

func (d *Document) Applets() *HTMLCollection {
    as := C.webkit_dom_document_get_applets(d.native())
    return newHTMLCollection(unsafe.Pointer(as))
}

func (d *Document) Links() *HTMLCollection {
    ls := C.webkit_dom_document_get_links(d.native())
    return newHTMLCollection(unsafe.Pointer(ls))
}

func (d *Document) Forms() *HTMLCollection {
    fs := C.webkit_dom_document_get_forms(d.native())
    return newHTMLCollection(unsafe.Pointer(fs))
}

func (d *Document) Anchors() *HTMLCollection {
    as := C.webkit_dom_document_get_anchors(d.native())
    return newHTMLCollection(unsafe.Pointer(as))
}

func (d *Document) LastModified() string {
    lm := C.webkit_dom_document_get_last_modified(d.native())
    return C.GoString((*C.char)(lm))
}

func (d *Document) Charset() string {
    cs := C.webkit_dom_document_get_charset(d.native())
    return C.GoString((*C.char)(cs))
}

func (d *Document) SetCharset(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_document_set_charset(d.native(), (*C.gchar)(v))
}

func (d *Document) DefaultCharset() string {
    dcs := C.webkit_dom_document_get_default_charset(d.native())
    return C.GoString((*C.char)(dcs))
}

func (d *Document) ReadyState() string {
    rs := C.webkit_dom_document_get_ready_state(d.native())
    return C.GoString((*C.char)(rs))
}

func (d *Document) PreferredStyleSheet() string {
    pss := C.webkit_dom_document_get_preferred_stylesheet_set(d.native())
    return C.GoString((*C.char)(pss))
}

func (d *Document) SelectedStyleSheet() string {
    s := C.webkit_dom_document_get_selected_stylesheet_set(d.native())
    return C.GoString((*C.char)(s))
}

func (d *Document) SetSelectedStyleSheet(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_document_set_selected_stylesheet_set(d.native(), (*C.gchar)(v))
}

func (d *Document) ActiveElement() *Element {
    e := C.webkit_dom_document_get_active_element(d.native())
    return newElement(unsafe.Pointer(e))
}

/*
 * Attr
 */

type Attr struct {
    Node
}

func marshalAttr(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapAttr(obj), nil
}

func wrapAttr(obj *glib.Object) *Attr {
    return &Attr{Node{Object{obj}}}
}

func (a *Attr) native() *C.WebKitDOMAttr {
    if a == nil || a.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(a.GObject)
    return C.toWebKitDOMAttr(p)
}

func newAttr(p unsafe.Pointer) *Attr {
    return wrapAttr(newGlibObject(p))
}

func (a *Attr) Name() string {
    s := C.webkit_dom_attr_get_name(a.native())
    return C.GoString((*C.char)(s))
}

func (a *Attr) Value() string {
    s := C.webkit_dom_attr_get_value(a.native())
    return C.GoString((*C.char)(s))
}

func (a *Attr) SetValue(value string) error {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    var err *C.GError
    C.webkit_dom_attr_set_value(a.native(), (*C.gchar)(v), &err)
    return err
}

func (a *Attr) Specified() bool {
    return gobool(C.webkit_dom_attr_get_specified(a.native()))
}

/*
 * Element
 */
type Element struct {
    Node
}

func marshalElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapElement(obj), nil
}

func wrapElement(obj *glib.Object) *Element {
    return &Element{Node{Object{obj}}}
}

func (e *Element) native() *C.WebKitDOMElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMElement(p)
}

func newElement(p unsafe.Pointer) *Element {
    return wrapElement(newGlibObject(p))
}

func (e *Element) Attribute(name string) string {
    n := C.CString(name)
    defer C.free(unsafe.Pointer(n))
    return C.GoString((*C.char)(C.webkit_dom_element_get_attribute(e.native(), (*C.gchar)(n))))
}

func (e *Element) SetAttribute(name, value string) error {
    n := C.CString(name)
    defer C.free(unsafe.Pointer(n))
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    var err *C.GError
    C.webkit_dom_element_set_attribute(e.native(), (*C.gchar)(n), (*C.gchar)(v), &err)
    return err
}

func (e *Element) RemoveAttribute(name string) {
    n := C.CString(name)
    defer C.free(unsafe.Pointer(n))
    C.webkit_dom_element_remove_attribute(e.native(), (*C.gchar)(n))
}

func (e *Element) AttributeNode(name string) *Attr {
    n := C.CString(name)
    defer C.free(unsafe.Pointer(n))
    ap := C.webkit_dom_element_get_attribute_node(e.native(), (*C.gchar)(n))
    return newAttr(unsafe.Pointer(ap))
}

func (e *Element) SetAttributeNode(attr *Attr) (*Attr, error) {
    var err *C.GError = nil
    oldAttr := C.webkit_dom_element_set_attribute_node(e.native(), attr.native(), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newAttr(unsafe.Pointer(oldAttr)), nil
}

func (e *Element) RemoveAttributeNode(attr *Attr) (*Attr, error) {
    var err *C.GError = nil
    nattr := C.webkit_dom_element_set_attribute_node(e.native(), attr.native(), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newAttr(unsafe.Pointer(nattr)), nil
}

func (e *Element) ElementsByTagName(tag string) *NodeList {
    t := C.CString(tag)
    defer C.free(unsafe.Pointer(t))
    nl := C.webkit_dom_element_get_elements_by_tag_name(e.native(), (*C.gchar)(t))
    return newNodeList(unsafe.Pointer(nl))
}

func (e *Element) HasAttributes() bool {
    return gobool(C.webkit_dom_element_has_attributes(e.native()))
}

// webkit_dom_element_get_attribute_ns ()
// webkit_dom_element_set_attribute_ns ()
// webkit_dom_element_remove_attribute_ns ()
// webkit_dom_element_get_elements_by_tag_name_ns ()
// webkit_dom_element_get_attribute_node_ns ()
// webkit_dom_element_set_attribute_node_ns ()

func (e *Element) HasAttribute(name string) bool {
    n := C.CString(name)
    defer C.free(unsafe.Pointer(n))
    return gobool(C.webkit_dom_element_has_attribute(e.native(), (*C.gchar)(n)))
}

// webkit_dom_element_has_attribute_ns ()

func (e *Element) Focus() {
    C.webkit_dom_element_focus(e.native())
}

func (e *Element) Blur() {
    C.webkit_dom_element_blur(e.native())
}

func (e *Element) ScrollIntoView(alignWithTop bool) {
    C.webkit_dom_element_scroll_into_view(e.native(), gboolean(alignWithTop))
}

func (e *Element) ScrollIntoViewIfNeed(centerIfNeeded bool) {
    C.webkit_dom_element_scroll_into_view(e.native(), gboolean(centerIfNeeded))
}

func (e *Element) ScrollByLines(lines int) {
    C.webkit_dom_element_scroll_by_lines(e.native(), C.glong(lines))
}

func (e *Element) ScrollByPages(pages int) {
    C.webkit_dom_element_scroll_by_pages(e.native(), C.glong(pages))
}

func (e *Element) ElementsByClassName(class string) *NodeList {
    c := C.CString(class)
    defer C.free(unsafe.Pointer(c))
    nl := C.webkit_dom_element_get_elements_by_class_name(e.native(), (*C.gchar)(c))
    return newNodeList(unsafe.Pointer(nl))
}

func (e *Element) QuerySelector(selectors string) (*NodeList, error) {
    s := C.CString(selectors)
    defer C.free(unsafe.Pointer(s))
    var err *C.GError
    nl := C.webkit_dom_element_query_selector(e.native(), (*C.gchar)(s), &err)
    if err != nil {
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newNodeList(unsafe.Pointer(nl)), nil
}

func (e *Element) QuerySelectorAll(selectors string) (*NodeList, error) {
    s := C.CString(selectors)
    defer C.free(unsafe.Pointer(s))
    var err *C.GError
    nl := C.webkit_dom_element_query_selector_all(e.native(), (*C.gchar)(s), &err)
    if err != nil {
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newNodeList(unsafe.Pointer(nl)), nil
}

func (e *Element) TagName() string {
    t := C.webkit_dom_element_get_tag_name(e.native())
    return C.GoString((*C.char)(t))
}

// func (e *Element) Attributes() NodeMap {
//     nm := C.webkit_dom_element_get_attributes(e.native())
//     return newNodeMap(unsafe.Pointer(nm))
// }

func (e *Element) Style() *CSSStyleDeclaration {
    csd := C.webkit_dom_element_get_style(e.native())
    return newCSSStyleDeclaration(unsafe.Pointer(csd))
}

func (e *Element) ID() string {
    s := C.webkit_dom_element_get_id(e.native())
    return C.GoString((*C.char)(s))
}

func (e *Element) SetID(id string) {
    i := C.CString(id)
    defer C.free(unsafe.Pointer(i))
    C.wekbit_dom_element_set_id(e.native, (*C.gchar)(i))
}

func (e *Element) OffsetLeft() float64 {
    return float64(C.webkit_dom_element_get_offset_left(e.native()))
}

func (e *Element) OffsetTop() float64 {
    return float64(C.webkit_dom_element_get_offset_top(e.native()))
}

func (e *Element) OffsetWidth() float64 {
    return float64(C.webkit_dom_element_get_offset_width(e.native()))
}

func (e *Element) OffsetHeight() float64 {
    return float64(C.webkit_dom_element_get_offset_height(e.native()))
}

func (e *Element) ClientLeft() float64 {
    return float64(C.webkit_dom_element_get_client_left(e.native()))
}

func (e *Element) ClientTop() float64 {
    return float64(C.webkit_dom_element_get_client_top(e.native()))
}

func (e *Element) ClientWidth() float64 {
    return float64(C.webkit_dom_element_get_client_width(e.native()))
}

func (e *Element) ClientHeight() float64 {
    return float64(C.webkit_dom_element_get_client_height(e.native()))
}

func (e *Element) ScrollLeft() int {
    return int(C.webkit_dom_element_get_scroll_left(e.native()))
}

func (e *Element) SetScrollLeft(value int) {
    C.webkit_dom_element_set_scroll_left(e.native(), C.glong(value))
}

func (e *Element) ScrollTop() int {
    return int(C.webkit_dom_element_get_scroll_top(e.native()))
}

func (e *Element) SetScrollTop(value int) {
    C.webkit_dom_element_set_scroll_top(e.native(), C.glong(value))
}

func (e *Element) ScrollWidth() int {
    return int(C.webkit_dom_element_get_scroll_width(e.native()))
}

func (e *Element) ScrollHeight() int {
    return int(C.webkit_dom_element_get_scroll_height(e.native()))
}

func (e *Element) OffsetParent() *Element {
    e := C.webkit_dom_element_get_offset_parent(e.native())
    return newElement(unsafe.Pointer(e))
}

func (e *Element) ClassName() string {
    v := C.webkit_dom_element_get_class_name(e.native())
    return C.GoString((*C.char)(v))
}

func (e *Element) SetClassName(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_element_set_class_name(e.native(), (*C.gchar)(v))
}

func (e *Element) FirstElementChild() *Element {
    e := C.webkit_dom_element_get_first_element_child(e.native())
    return newElement(unsafe.Pointer(e))
}

func (e *Element) LastElementChild() *Element {
    e := C.webkit_dom_element_get_last_element_child(e.native())
    return newElement(unsafe.Pointer(e))
}

func (e *Element) PreviousElementSibling() *Element {
    e := C.webkit_dom_element_get_previous_element_sibling(e.native())
    return newElement(unsafe.Pointer(e))
}

func (e *Element) NextElementSibling() *Element {
    e := C.webkit_dom_element_get_next_element_sibling(e.native())
    return newElement(unsafe.Pointer(e))
}

func (e *Element) ChildElementCount() uint {
    return uint(C.webkit_dom_element_get_child_element_count(e.native()))
}

/*
 * HTMLElement
 */

type HTMLElement struct {
    Element
}

func marshalHTMLElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLElement(obj), nil
}

func wrapHTMLElement(obj *glib.Object) *HTMLElement {
    return &HTMLElement{Element{Node{Object{obj}}}}
}

func (e *HTMLElement) native() *C.WebKitDOMHTMLElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLElement(p)
}

func newHTMLElement(p unsafe.Pointer) *HTMLElement {
    return wrapHTMLElement(newGlibObject(p))
}

func (e *HTMLElement) Click() {
    C.webkit_dom_html_element_click(e.native())
}

func (e *HTMLElement) Title() string {
    t := C.webkit_dom_html_element_get_title(e.native())
    return C.GoString((*C.char)(t))
}

func (e *HTMLElement) SetTitle(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_element_set_title(e.native(), (*C.gchar)(v))
}

func (e *HTMLElement) Lang() string {
    l := C.webkit_dom_html_element_get_lang(e.native())
    return C.GoString((*C.char)(l))
}

func (e *HTMLElement) SetLang(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_element_set_lang(e.native(), (*C.gchar)(v))
}

func (e *HTMLElement) Dir() string {
    d := C.webkit_dom_html_element_get_dir(e.native())
    return C.GoString((*C.char)(d))
}

func (e *HTMLElement) SetDir(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_element_set_dir(e.native(), (*C.gchar)(v))
}

func (e *HTMLElement) TabIndex() int {
    return int(C.webkit_dom_html_element_get_tab_index(e.native()))
}

func (e *HTMLElement) SetTabIndex(value int) {
    C.webkit_dom_html_element_set_tab_index(e.native(), C.glong(value))
}

func (e *HTMLElement) AccessKey() string {
    k := C.webkit_dom_html_element_get_access_key(e.native())
    return C.GoString((*C.char)(k))
}

func (e *HTMLElement) SetAccessKey(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_element_set_access_key(e.native(), (*C.gchar)(v))
}

func (e *HTMLElement) InnerHTML() string {
    h := C.webkit_dom_html_element_get_inner_html(e.native())
    return C.GoString((*C.char)(h))
}

func (e *HTMLElement) SetInnerHTML(value string) error {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    var err *C.GError
    C.webkit_dom_html_element_set_inner_html(e.native(), (*C.gchar)(v), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (e *HTMLElement) InnerText() string {
    t := C.webkit_dom_html_element_get_inner_text(e.native())
    return C.GoString((*C.char)(t))
}

func (e *HTMLElement) SetInnerText(value string) error {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_element_set_inner_text(e.native(), (*C.gchar)(v))
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (e *HTMLElement) OuterHTML() string {
    h := C.webkit_dom_html_element_get_outer_html(e.native())
    return C.GoString((*C.char)(h))
}

func (e *HTMLElement) SetOuterHTML(value string) error {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_element_set_outer_html(e.native(), (*C.gchar)(v), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (e *HTMLElement) OuterText() string {
    t := C.webkit_dom_html_element_get_outer_text(e.native())
    return C.GoString((*C.char)(t))
}

func (e *HTMLElement) SetOuterText(value string) error {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_element_set_outer_text(e.native(), (*C.gchar)(v), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (e *HTMLElement) Children() *HTMLCollection {
    c := C.webkit_dom_html_element_get_children(e.native())
    return newHTMLCollection(unsafe.Pointer(c))
}

func (e *HTMLElement) ContentEditable() string {
    c := C.webkit_dom_html_element_get_content_editable(e.native())
    return C.GoString((*C.char)(c))
}

func (e *HTMLElement) SetContentEditable(value string) error {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_element_set_content_editable(e.native(), (*C.gchar)(v), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (e *HTMLElement) IsContentEditable() bool {
    return gobool(C.webkit_dom_html_element_get_is_content_editable(e.native()))
}

/*
 * HTMLCollection
 */

type HTMLCollection struct {
    Object
}

func marshalHTMLCollection(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLCollection(obj), nil
}

func wrapHTMLCollection(obj *glib.Object) *HTMLCollection {
    return &HTMLCollection{Object{obj}}
}

func (c *HTMLCollection) native() *C.WebKitDOMHTMLCollection {
    if c == nil || c.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(c.GObject)
    return C.toWebKitDOMHTMLCollection(p)
}

func newHTMLCollection(p unsafe.Pointer) *HTMLCollection {
    return wrapHTMLCollection(newGlibObject(p))
}

func (c *HTMLCollection) Item(index uint) *Node {
    n := C.webkit_dom_html_collection_item(c.native(), C.gulong(index))
    return newNode(unsafe.Pointer(n))
}

func (c *HTMLCollection) NamedItem(name string) *Node {
    n := C.CString(name)
    defer C.free(unsafe.Pointer(n))
    ip := C.webkit_dom_html_collection_named_item(e.native(), (*C.gchar)(n))
    return newNode(unsafe.Pointer(ip))
}

func (c *HTMLCollection) Len() uint {
    return uint(C.webkit_dom_html_collection_get_length(e.native()))
}

/*
 * StyleSheet
 */

type StyleSheet struct {
    Object
}

func marshalStyleSheet(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapStyleSheet(obj), nil
}

func wrapStyleSheet(obj *glib.Object) *StyleSheet {
    return &StyleSheet{Object{obj}}
}

func (s *StyleSheet) native() *C.WebKitDOMStyleSheet {
    if s == nil || s.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(s.GObject)
    return C.toWebKitDOMStyleSheet(p)
}

func newStyleSheet(p unsafe.Pointer) *StyleSheet {
    return wrapStyleSheet(newGlibObject(p))
}

func (s *StyleSheet) ContentType() string {
    c := C.webkit_dom_style_sheet_get_content_type(s.native())
    return C.GoString((*C.char)(c))
}

func (s *StyleSheet) Disabled() bool {
    return gobool(C.webkit_dom_style_sheet_get_disabled(s.native()))
}

func (s *StyleSheet) SetDisabled(value bool) {
    C.webkit_dom_style_sheet_set_disabled(s.native(), gboolean(value))
}

func (s *StyleSheet) OwnerNode() *Node {
    node := C.webkit_dom_style_sheet_get_owner_node(s.native())
    return newNode(unsafe.Pointer(node))
}

func (s *StyleSheet) ParentStyleSheet() *StyleSheet {
    ps := C.webkit_dom_style_sheet_get_parent_style_sheet(s.native())
    return newStyleSheet(ps)
}

func (s *StyleSheet) Href() string {
    h := C.webkit_dom_style_sheet_get_href(s.native())
    return C.GoString((*C.char)(h))
}

func (s *StyleSheet) Title() string {
    t := C.webkit_dom_style_sheet_get_title(s.native())
    return C.GoString((*C.char)(t))
}

func (s *StyleSheet) Media() *MediaList {
    ml := C.webkit_dom_style_sheet_get_media(s.native())
    return newMediaList(unsafe.Pointer(ml))
}

type StyleSheetList struct {
    Object
}

func marshalStyleSheetList(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapStyleSheetList(obj), nil
}

func wrapStyleSheetList(obj *glib.Object) *StyleSheetList {
    return &StyleSheetList{Object{obj}}
}

func (l *StyleSheetList) native() *C.WebKitDOMStyleSheetList {
    if l == nil || l.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(l.GObject)
    return C.toWebKitDOMStyleSheetList(p)
}

func newStyleSheetList(p unsafe.Pointer) *StyleSheetList {
    return wrapStyleSheetList(newGlibObject(p))
}

func (l *StyleSheetList) Item(index uint) *StyleSheet {
    ss := C.webkit_dom_style_sheet_list_item(l.native(), C.gulong(index))
    return newStyleSheet(unsafe.Pointer(ss))
}

func (l *StyleSheetList) Len() uint {
    return uint(C.webkit_dom_style_sheet_list_get_length(l.native()))
}

/*
 * MediaList
 */

type MediaList struct {
    Object
}

func marshalMediaList(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapMediaList(obj), nil
}

func wrapMediaList(obj *glib.Object) *MediaList {
    return &MediaList{Obj{obj}}
}

func (l *MediaList) native() *C.WebKitDOMMediaList {
    if l == nil || l.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(l.GObject)
    return C.toWebKitDOMMediaList(p)
}

func newMediaList(p unsafe.Pointer) *MediaList {
    return wrapMediaList(newGlibObject(p))
}

func (l *MediaList) Item(index uint) string {
    i := C.webkit_dom_media_list_item(l.native(), C.gulong(index))
    return C.GoString((*C.char)(i))
}

func (l *MediaList) DeleteMedium(medium string) error {
    m := C.CString(medium)
    defer C.free(unsafe.Pointer(m))
    var err *C.GError
    C.webkit_dom_media_list_delete_medium(l.native(), (*C.gchar)(m), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (l *MediaList) AppendMedium(medium string) error {
    m := C.CString(medium)
    defer C.free(unsafe.Pointer(m))
    var err *C.GError
    C.webkit_dom_media_list_append_medium(l.native(), (*C.gchar)(m), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (l *MediaList) MediaText() string {
    t := C.webkit_dom_media_list_get_media_text(l.native())
    return C.GoString((*C.char)(t))
}

func (l *MediaList) SetMediaText(value string) error {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    var err *C.GError
    C.webkit_dom_media_list_set_media_text(l.native(), (*C.gchar)(v), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (l *MediaList) Len() uint {
    return uint(C.webkit_dom_media_list_get_length(l.native()))
}

/*
 * CSSStyleSheet
 */

type CSSStyleSheet struct {
    StyleSheet
}

func marshalCSSStyleSheet(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapCSSStyleSheet(obj), nil
}

func wrapCSSStyleSheet(obj *glib.Object) *CSSStyleSheet {
    return &CSSStyleSheet{StyleSheet{Object{obj}}}
}

func (s *CSSStyleSheet) native() *C.WebKitDOMCSSStyleSheet {
    if s == nil || s.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(s.GObject)
    return C.toWebKitDOMCSSStyleSheet(p)
}

func newCSSStyleSheet(p unsafe.Pointer) *CSSStyleSheet {
    return wrapCSSStyleSheet(newGlibObject(p))
}

func (s *CSSStyleSheet) InsertRule(rule string, index uint) error {
    r := C.CString(rule)
    defer C.free(unsafe.Pointer(r))
    var err *C.GError
    C.webkit_dom_css_style_sheet_insert_rule(s.native(), (*C.gchar)(r), C.gulong(index), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (s *CSSStyleSheet) DeleteRule(index uint) error {
    var err *C.GError
    C.webkit_dom_css_style_sheet_delete_rule(s.native(), C.gulong(index))
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (s *CSSStyleSheet) AddRule(selector, style string, index uint) error {
    se := C.CString(selector)
    defer C.free(unsafe.Pointer(se))
    st := C.CString(style)
    defer C.free(unsafe.Pointer(st))
    var err *C.GError
    C.webkit_dom_css_style_sheet_add_rule(s.native(), (*C.gchar)(se), (*C.gchar)(st), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (s *CSSStyleSheet) RemoveRule(index uint) error {
    var err *C.GError
    C.webkit_dom_css_style_sheet_remove_rule(s.native(), C.gulong(index))
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (s *CSSStyleSheet) OwnerRule() *CSSRule {
    r := C.webkit_dom_css_style_sheet_get_owner_rule(s.native())
    return newCSSRule(unsafe.Pointer(r))
}

// webkit_dom_css_style_sheet_get_css_rules?
func (s *CSSStyleSheet) Rules() *CSSRuleList {
    rs := C.webkit_dom_css_style_sheet_get_rules(s.native())
    return newCSSRuleList(unsafe.Pointer(rs))
}

/*
 * CSSRule
 */

type CSSRule struct {
    Object
}

func marshalCSSRule(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapCSSRule(obj), nil
}

func wrapCSSRule(obj *glib.Object) *CSSRule {
    return &CSSRule{Object{obj}}
}

func (r *CSSRule) native() *C.WebKitDOMCSSRule {
    if r == nil || r.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(r.GObject)
    return C.toWebKitDOMCSSRule(p)
}

func newCSSRule(p unsafe.Pointer) *CSSRule {
    return wrapCSSRule(newGlibObject(p))
}

// TODO: Fix this.
type CSSRuleType int

const (
    CSS_RULE_UNKNOWN_RULE   CSSRuleType = C.WEBKIT_DOM_CSS_RULE_UNKNOWN_RULE
    CSS_RULE_STYLE_RULE     CSSRuleType = C.WEBKIT_DOM_CSS_RULE_STYLE_RULE
    CSS_RULE_CHARSET_RULE   CSSRuleType = C.WEBKIT_DOM_CSS_RULE_CHARSET_RULE
    CSS_RULE_IMPORT_RULE    CSSRuleType = C.WEBKIT_DOM_CSS_RULE_IMPORT_RULE
    CSS_RULE_MEDIA_RULE     CSSRuleType = C.WEBKIT_DOM_CSS_RULE_MEDIA_RULE
    CSS_RULE_FONT_FACE_RULE CSSRuleType = C.WEBKIT_DOM_CSS_RULE_FONT_FACE_RULE
    CSS_RULE_PAGE_RULE      CSSRuleType = C.WEBKIT_DOM_CSS_RULE_PAGE_RULE
)

func marshalCSSRuleType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return CSSRuleType(c), nil
}

func (r *CSSRule) RuleType() CSSRuleType {
    r := C.webkit_dom_css_rule_get_rule_type(r.native())
    return CSSRuleType(r)
}

func (r *CSSRule) Text() string {
    t := C.webkit_dom_css_rule_get_css_text(r.native())
    return C.GoString((*C.char)(t))
}

func (r *CSSRule) SetText(value string) error {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    var err *C.GError
    C.webkit_dom_css_rule_set_css_text(e.native(), (*C.gchar)(v), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
}

func (r *CSSRule) ParentStyleSheet() *CSSStyleSheet {
    s := C.webkit_dom_css_rule_get_parent_style_sheet(r.native())
    return newCSSStyleSheet(unsafe.Pointer(s))
}

func (r *CSSRule) ParentRule() *CSSRule {
    r := C.webkit_dom_css_rule_get_parent_rule(e.native())
    return newCSSRule(unsafe.Pointer(r))
}

/*
 * CSSRuleList
 */

type CSSRuleList struct {
    Object
}

func marshalCSSRuleList(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapCSSRuleList(obj), nil
}

func wrapCSSRuleList(obj *glib.Object) *CSSRuleList {
    return &CSSRuleList{Object{obj}}
}

func (l *CSSRuleList) native() *C.WebKitDOMCSSRuleList {
    if l == nil || l.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(l.GObject)
    return C.toWebKitDOMCSSRuleList(p)
}

func newCSSRuleList(p unsafe.Pointer) *CSSRuleList {
    return wrapCSSRuleList(newGlibObject(p))
}

func (l *CSSRuleList) Item(index uint) *CSSRule {
    r := C.webkit_dom_css_rule_list_item(l.native(), C.gulong(index))
    return newCSSRule(unsafe.Pointer(l))
}

func (l *CSSRuleList) Len() uint {
    return uint(C.webkit_dom_css_rule_list_get_length(l.native()))
}

/*
 * CSSStyleDeclaration
 */

type CSSStyleDeclaration struct {
    Object
}

func marshalCSSStyleDeclaration(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapCSSStyleDeclaration(obj), nil
}

func wrapCSSStyleDeclaration(obj *glib.Object) *CSSStyleDeclaration {
    return &CSSStyleDeclaration{Object{obj}}
}

func (d *CSSStyleDeclaration) native() *C.WebKitDOMCSSStyleDeclaration {
    if d == nil || d.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(d.GObject)
    return C.toWebKitDOMCSSStyleDeclaration(p)
}

func newCSSStyleDeclaration(p unsafe.Pointer) *CSSStyleDeclaration {
    return wrapCSSStyleDeclaration(newGlibObject(p))
}

func (d *CSSStyleDeclaration) PropertyValue(property string) string {
    p := C.CString(property)
    defer C.free(unsafe.Pointer(p))
    v := C.webkit_dom_css_style_declaration_get_property_value(d.native(), (*C.gchar)(p))
    return C.GoString((*C.char)(v))
}

func (d *CSSStyleDeclaration) RemoveProperty(property string) (string, error) {
    p := C.CString(property)
    defer C.free(unsafe.Pointer(p))
    var err *C.GError
    v := C.webkit_dom_css_style_declaration_remove_property(d.native(), (*C.gchar)(p), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return C.GoString((*C.char)(v)), nil
}

func (d *CSSStyleDeclaration) PropertyPriority(property string) string {
    p := C.CString(property)
    defer C.free(unsafe.Pointer(p))
    v := C.webkit_dom_css_style_declaration_get_property_priority(d.native(), (*C.gchar)(p))
    return C.GoString((*C.char)(v))
}

func (d *CSSStyleDeclaration) SetPriority(property, value, priority string) error {
    prop := C.CString(property)
    defer C.free(unsafe.Pointer(prop))
    val := C.CString(value)
    defer C.free(unsafe.Pointer(val))
    prio := C.CString(priority)
    defer C.free(unsafe.Pointer(prio))
    var err *C.GError
    C.webkit_dom_css_style_declaration_set_property(d.native(), (*C.gchar)(prop), (*C.gchar)(val), (*C.gchar)(prio), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (d *CSSStyleDeclaration) Item(index uint) string {
    i := C.webkit_dom_css_style_declaration_item(d.native(), C.gulong(index))
    return C.GoString((*C.char)(i))
}

func (d *CSSStyleDeclaration) PropertyShorthand(property string) string {
    p := C.CString(property)
    defer C.free(unsafe.Pointer(p))
    s := C.webkit_dom_css_style_declaration_get_property_shorthand(d.native(), (*C.gchar)(p))
    return C.GoString((*C.char)(s))
}

func (d *CSSStyleDeclaration) IsPropertyImplicit(property string) bool {
    p := C.CString(property)
    defer C.free(unsafe.Pointer(p))
    b := C.webkit_dom_css_style_declaration_is_property_implicit(d.native(), (*C.gchar)(p))
    return gobool(b)
}

func (d *CSSStyleDeclaration) CSSText() string {
    t := C.webkit_dom_css_style_declaration_get_css_text(d.native())
    return C.GoString((*C.char)(t))
}

func (d *CSSStyleDeclaration) SetCSSText(value string) error {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    var err *C.GError
    C.webkit_dom_css_style_declaration_set_css_text(d.native(), (*C.gchar)(v), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (d *CSSStyleDeclaration) Len() uint {
    return uint(C.webkit_dom_css_style_declaration_get_length(d.native()))
}

func (d *CSSStyleDeclaration) ParentRule() *CSSRule {
    r := C.webkit_dom_css_style_declaration_get_parent_rule(d.native())
    return newCSSRule(unsafe.Pointer(r))
}

/*
 * CSSValue
 */

 // TODO: constants?
type CSSValue struct {
    Object
}

func marshalCSSValue(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapCSSValue(obj), nil
}

func wrapCSSValue(obj *glib.Object) *CSSValue {
    return &CSSValue{Object{obj}}
}

func (c *CSSValue) native() *C.WebKitDOMCSSValue {
    if c == nil || c.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(c.GObject)
    return C.toWebKitDOMCSSValue(p)
}

func newCSSValue(p unsafe.Pointer) *CSSValue {
    return wrapCSSValue(newGlibObject(p))
}

func (c *CSSValue) CSSText() string {
    t := C.webkit_dom_css_value_get_css_text(c.native())
    return C.GoString((*C.char)(t))
}

func (c *CSSValue) SetCSSText(value string) error {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    var err *C.GError
    C.webkit_dom_css_value_set_css_text(c.native(), (*C.gchar)(v), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

// TODO: constants?
func (c *CSSValue) CSSValueType() uint16 {
    return uint16(C.webkit_dom_css_value_get_css_value_type(c.native()))
}


/*
 * Event
 */

// TODO: constants?
type Event struct {
    Object
}

func marshalEvent(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapEvent(obj), nil
}

func wrapEvent(obj *glib.Object) *Event {
    return &Event{Object{obj}}
}

func (e *Event) native() *C.WebKitDOMEvent {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMEvent(p)
}

func newEvent(p unsafe.Pointer) *Event {
    return wrapEvent(newGlibObject(p))
}

func (e *Event) StopPropagation() {
    C.webkit_dom_event_stop_propagation(e.native())
}

func (e *Event) PreventDefault() {
    C.webkit_dom_event_prevent_default(e.native())
}

func (e *Event) InitEvent(eventType string, canBubble, cancelable bool) {
    et := C.CString(eventType)
    defer C.free(unsafe.Pointer(et))
    C.webkit_dom_event_init_event(e.native(), (*C.gchar)(et), gboolean(canBubble), gboolean(cancelable))
}

func (e *Event) Type() string {
    t := C.webkit_dom_event_get_event_type(e.natiev())
    return C.GoString((*C.char)(t))
}

// TODO
// func (e *Event) Target() EventTarget {
//     et := C.webkit_dom_event_get_target(e.native())
//     return newEventTarget(unsafe.Pointer(et))
// }
// func (e *Event) CurrentTarget() EventTarget {
//     et := C.webkit_dom_event_get_current_target(e.native())
//     return newEventTarget(unsafe.Pointer(et))
// }

func (e *Event) Phase() uint16 {
    return uint16(C.webkit_dom_event_get_event_phase(e.native()))
}

func (e *Event) Bubbles() bool {
    return gobool(C.webkit_dom_event_get_bubbles(e.native()))
}

func (e *Event) Cancelable() bool {
    return gobool(C.webkit_dom_event_get_cancelable(e.native()))
}

func (e *Event) Timestamp() uint32 {
    return uint32(C.webkit_dom_event_get_time_stamp(e.native()))
}

// TODO
// func (e *Event) SrcElement() EventTarget {
//     src := C.webkit_dom_event_get_src_element(e.native())
//     return newEventTarget(unsafe.Pointer(src))
// }

func (e *Event) ReturnValue() bool {
    return gobool(C.webkit_dom_event_get_return_value(e.native()))
}

func (e *Event) SetReturnValue(value bool) {
    C.webkit_dom_event_set_return_value(e.native(), gboolean(value))
}

// TODO: find a better name for this one.
func (e *Event) CancelBubble() bool {
    return gobool(C.webkit_dom_event_get_cancel_bubble(e.native()))
}

func (e *Event) SetCancelBubble(value bool) {
    C.webkit_dom_event_set_cancel_bubble(e.native(), gboolean(value))
}

/*
 * CharacterData
 */

type CharacterData struct {
    Node
}

func marshalCharacterData(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapCharacterData(obj), nil
}

func wrapCharacterData(obj *glib.Object) *CharacterData {
    return &CharacterData{Node{Object{{obj}}}}
}

func (c *CharacterData) native() *C.WebKitDOMCharacterData {
    if c == nil || c.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(c.GObject)
    return C.toWebKitDOMCharacterData(p)
}

func newCharacterData(p unsafe.Pointer) *CharacterData {
    return wrapCharacterData(newGlibObject(p))
}

func (d *CharacterData) SubstringData(offset, length uint) (string, error) {
    var err *C.GError
    s := C.webkit_dom_character_data_substring_data(d.native(), C.gulong(offset), C.gulong(length), &err)
    if err != nil {
        defer C.g_error_free(err)
        return "", errors.New(C.GoString((*C.char)(err.message)))
    }
    return C.GoString((*C.char)(s)), nil
}

func (d *CharacterData) AppendData(data string) error {
    ad := C.CString(data)
    defer C.free(unsafe.Pointer(ad))
    var err *C.GError
    C.webkit_dom_character_data_append_data(d.native(), (*C.gchar)(ad), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (d *CharacterData) InsertData(offset uint, data string) error {
    id := C.CString(data)
    defer C.free(unsafe.Pointer(id))
    var err *C.GError
    C.webkit_dom_character_data_insert_data(d.native(), (*C.gchar)(id), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (d *CharacterData) DeleteData(offset, length uint) error {
    var err *C.GError
    C.webkit_dom_character_data_delete_data(d.native(), C.gulong(offset), C.gulong(length), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (d *CharacterData) ReplaceData(offset, length uint, data string) error {
    rd := C.CString(data)
    defer C.free(unsafe.Pointer(rd))
    var err *C.GError
    C.webkit_dom_character_data_replace_data(d.native(), C.gulong(offset), C.gulong(length), (*C.gchar)(rd), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (d *CharacterData) Data() string {
    s := C.webkit_dom_character_data_get_data(d.native())
    return C.GoString((*C.char)(s))
}

func (d *CharacterData) SetData(value string) error {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    var err *C.GError
    C.webkit_dom_character_data_set_data(d.native(), (*C.gchar)(v), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

func (d *CharacterData) Len() uint {
    return uint(C.webkit_dom_character_data_get_length(d.native()))
}

/*
 * Text
 */

type Text struct {
    CharacterData
}

func marshalText(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapText(obj), nil
}

func wrapText(obj *glib.Object) *Text {
    return &Text{CharacterData{Node{Object{obj}}}}
}

func (t *Text) native() *C.WebKitDOMText {
    if t == nil || t.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(t.GObject)
    return C.toWebKitDOMText(p)
}

func newText(p unsafe.Pointer) *Text {
    return wrapText(newGlibObject(p))
}

func (t *Text) SplitText(offset uint) (*Text, error) {
    var err *C.GError
    st := C.webkit_dom_text_split_text(t.native(), C.gulong(offset), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newText(unsafe.Pointer(st)), nil
}

func (t *Text) ReplaceWholeText(content string) (*Text, error) {
    c := C.CString(content)
    defer C.free(unsafe.Pointer(c))
    var err *C.GError
    wt := C.webkit_dom_text_replace_whole_text(t.native(), (*C.gchar)(c), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newText(unsafe.Pointer(wt)), nil
}

func (t *Text) WholeText() string {
    wt := C.webkit_dom_text_get_whole_text(t.native())
    return C.GoString((*C.char)(wt))
}

// TODO: WebKitDOMCDATASection, WebKitDOMComment

/*
 * ProcessingInstruction
 */

type ProcessingInstruction struct {
    CharacterData
}

func marshalProcessingInstruction(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapProcessingInstruction(obj), nil
}

func wrapProcessingInstruction(obj *glib.Object) *ProcessingInstruction {
    return &ProcessingInstruction{CharacterData{Node{Object{obj}}}}
}

func (pi *ProcessingInstruction) native() *C.WebKitDOMProcessingInstruction {
    if pi == nil || pi.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(pi.GObject)
    return C.toWebKitDOMProcessingInstruction(p)
}

func newProcessingInstruction(p unsafe.Pointer) *ProcessingInstruction {
    return wrapProcessingInstruction(newGlibObject(p))
}

func (p *ProcessingInstruction) Target() string {
    t := C.webkit_dom_processing_instruction_get_target(p.native())
    return C.GoString((*C.char)(t))
}

func (p *ProcessingInstruction) Sheet() *StyleSheet {
    s := C.webkit_dom_processing_instruction_get_sheet(p.native())
    return newStyleSheet(unsafe.Pointer(s))
}

/*
 * FileList
 */

type FileList struct {
    Object
}

func marshalFileList(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapFileList(obj), nil
}

func wrapFileList(obj *glib.Object) *FileList {
    return &FileList{Object{obj}}
}

func (l *FileList) native() *C.WebKitDOMFileList {
    if l == nil || l.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(l.GObject)
    return C.toWebKitDOMFileList(p)
}

func newFileList(p unsafe.Pointer) *FileList {
    return wrapFileList(newGlibObject(p));
}

// Item is a wrapper around webkit_dom_file_list_item():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMFileList.html#webkit-dom-file-list-item
func (l *FileList) Item(index uint) *File {
    p := C.webkit_dom_file_list_item(l.native())
    return newFile(unsafe.Pointer(p))
}

// Len is a wrapper around webkit_dom_file_list_get_length():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMFileList.html#webkit-dom-file-list-get-length
func (l *FileList) Len() uint {
    return uint(C.webkit_dom_file_list_get_length(l.native()))
}

/*
 * File
 */

type File struct {
    Blob
}

func marshalFile(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapFile(obj), nil
}

func wrapFile(obj *glib.Object) *File {
    return &File{Blob{Object{obj}}}
}

func (f *File) native() *C.WebKitDOMFile {
    if f == nil || f.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(f.GObject)
    return C.toWebKitDOMFile(p)
}

func newFile(p unsafe.Pointer) *File {
    return wrapFile(newGlibObject(p));
}

// Name is a wrapper around webkit_dom_file_get_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMFile.html#webkit-dom-file-get-name
func (f *File) Name() string {
    cstr := C.webkit_dom_file_get_name(e.native())
    return C.GoString((*C.char)(cstr))
}

/*
 * Blob
 */

type Blob struct {
    Object
}

func marshalBlob(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapBlob(obj), nil
}

func wrapBlob(obj *glib.Object) *Blob {
    return &Blob{obj}
}

func (b *Blob) native() *C.WebKitDOMBlob {
    if b == nil || b.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(b.GObject)
    return C.toWebKitDOMBlob(p)
}

func newBlob(p unsafe.Pointer) *Blob {
    return wrapBlob(newGlibObject(p));
}

// Size is a wrapper around webkit_dom_blob_get_size():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMBlob.html#webkit-dom-blob-get-size
func (b *Blob) Size() uint64 {
    return uint64(C.webkit_dom_blob_get_size(b.native()))
}

/*
 * XPathExpression
 */

type XPathExpression struct {
    Object
}

func marshalXPathExpression(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapXPathExpression(obj), nil
}

func wrapXPathExpression(obj *glib.Object) *XPathExpression {
    return &XPathExpression{Object{obj}}
}

func (e *XPathExpression) native() *C.WebKitDOMXPathExpression {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMXPathExpression(p)
}

func newXPathExpression(p unsafe.Pointer) *XPathExpression {
    return wrapXPathExpression(newGlibObject(p));
}

// Evaluate is a wrapper around webkit_dom_xpath_expression_evaluate():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMXPathExpression.html#webkit-dom-xpath-expression-evaluate
func (e *XPathExpression) Evaluate(context *Node, kind int, inResult *XPathResult) (*XPathResult, error) {
    var err *C.GError
    p := C.webkit_dom_xpath_expression_evaluate(e.native(), context.native(), C.ushort(kind), inResult.native(), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newXPathResult(unsafe.Pointer(p)), nil
}

/*
 * XPathResult
 */

type XPathResult struct {
    Object
}

func marshalXPathResult(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapXPathResult(obj), nil
}

func wrapXPathResult(obj *glib.Object) *XPathResult {
    return &XPathResult{Object{obj}}
}

func (r *XPathResult) native() *C.WebKitDOMXPathResult {
    if r == nil || r.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(r.GObject)
    return C.toWebKitDOMXPathResult(p)
}

func newXPathResult(p unsafe.Pointer) *XPathResult {
    return wrapXPathResult(newGlibObject(p));
}

// IterateNext is a wrapper around webkit_dom_xpath_result_iterate_next():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMXPathResult.html#webkit-dom-xpath-result-iterate-next
func (r *XPathResult) IterateNext() (*Node, error) {
    var err *C.GError
    p := C.webkit_dom_xpath_result_iterate_next(e.native(), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newNode(unsafe.Pointer(p)), nil
}

// SnapshotItem is a wrapper around webkit_dom_xpath_result_snapshot_item():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMXPathResult.html#webkit-dom-xpath-result-snapshot-item
func (r *XPathResult) SnapshotItem(index uint) (*Node, error) {
    var err *C.GError
    p := C.webkit_dom_xpath_result_snapshot_item(e.native(), &err)
    if err != nil {
        defer C.g_error_free(err)
        return nil, errors.New(C.GoString((*C.char)(err.message)))
    }
    return newNode(unsafe.Pointer(p)), nil
}

/*
 * Window
 */

type Window struct {
    Object
}

func marshalWindow(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapWindow(obj), nil
}

func wrapWindow(obj *glib.Object) *Window {
    return &Window{Object{obj}}
}

func (w *Window) native() *C.WebKitDOMDOMWindow {
    if w == nil || w.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(w.GObject)
    return C.toWebKitDOMDOMWindow(p)
}

func newWindow(p unsafe.Pointer) *Window {
    return wrapWindow(newGlibObject(p));
}

// All HTML elements follow. Only part of them are implemented. They are either 
// fully implemented or not at all.

/*
 * HTMLAnchorElement
 */

type HTMLAnchorElement struct {
    HTMLElement
}

func marshalHTMLAnchorElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLAnchorElement(obj), nil
}

func wrapHTMLAnchorElement(obj *glib.Object) *HTMLAnchorElement {
    return &HTMLAnchorElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLAnchorElement) native() *C.WebKitDOMHTMLAnchorElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLAnchorElement(p)
}

func newHTMLAnchorElement(p unsafe.Pointer) *HTMLAnchorElement {
    return wrapHTMLAnchorElement(newGlibObject(p));
}

func (e *HTMLAnchorElement) Charset() string {
    c := C.webkit_dom_html_anchor_element_get_charset(e.native())
    return C.GoString((*C.char)(c))
}

func (e *HTMLAnchorElement) SetCharset(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_anchor_element_set_charset(e.native(), (*C.gchar)(v))
}

func (e *HTMLAnchorElement) Coords() string {
    c := C.webkit_dom_html_anchor_element_get_coords(e.native())
    return C.GoString((*C.char)(c))
}

func (e *HTMLAnchorElement) SetCoords(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_anchor_element_set_coords(e.native(), (*C.gchar)(v))
}

func (e *HTMLAnchorElement) Href() string {
    h := C.webkit_dom_html_anchor_element_get_href(e.native())
    return C.GoString((*C.char)(h))
}

func (e *HTMLAnchorElement) SetHref(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_anchor_element_set_href(e.native(), (*C.gchar)(v))
}

func (e *HTMLAnchorElement) Hreflang() string {
    hl := C.webkit_dom_html_anchor_element_get_hreflang(e.native())
    return C.GoString((*C.char)(hl))
}

func (e *HTMLAnchorElement) SetHreflang(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_anchor_element_set_hreflang(e.native(), (*C.gchar)(v))
}

func (e *HTMLAnchorElement) Name() string {
    n := C.webkit_dom_html_anchor_element_get_name(e.native())
    return C.GoString((*C.char)(n))
}

func (e *HTMLAnchorElement) SetName(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_anchor_element_set_name(e.native(), (*C.gchar)(v))
}

func (e *HTMLAnchorElement) Rel() string {
    r := C.webkit_dom_html_anchor_element_get_rel(e.native())
    return C.GoString((*C.char)(r))
}

func (e *HTMLAnchorElement) SetRel(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_anchor_element_set_rel(e.native(), (*C.gchar)(v))
}

func (e *HTMLAnchorElement) Rev() string {
    r := C.webkit_dom_html_anchor_element_get_rev(e.native())
    return C.GoString((*C.char)(r))
}

func (e *HTMLAnchorElement) SetRev(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_anchor_element_set_rev(e.native(), (*C.gchar)(v))
}

func (e *HTMLAnchorElement) Shape() string {
    s := C.webkit_dom_html_anchor_element_get_shape(e.native())
    return C.GoString((*C.char)(s))
}

func (e *HTMLAnchorElement) SetShape(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_anchor_element_set_shape(e.native(), (*C.gchar)(v))
}

func (e *HTMLAnchorElement) Target() string {
    t := C.webkit_dom_html_anchor_element_get_target(e.native())
    return C.GoString((*C.char)(t))
}

func (e *HTMLAnchorElement) SetTarget(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_anchor_element_set_target(e.native(), (*C.gchar)(v))
}

func (e *HTMLAnchorElement) TypeAttr() string {
    a := C.webkit_dom_html_anchor_element_get_type_attr(e.native())
    return C.GoString((*C.char)(a))
}

func (e *HTMLAnchorElement) SetTypeAttr(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_anchor_element_set_type_attr(e.native(), (*C.gchar)(v))
}

func (e *HTMLAnchorElement) Hash() string {
    h := C.webkit_dom_html_anchor_element_get_hash(e.native())
    return C.GoString((*C.char)(h))
}

func (e *HTMLAnchorElement) SetHash(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_anchor_element_set_hash(e.native(), (*C.gchar)(v))
}

func (e *HTMLAnchorElement) Host() string {
    h := C.webkit_dom_html_anchor_element_get_host(e.native())
    return C.GoString((*C.char)(h))
}

func (e *HTMLAnchorElement) SetHost(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_anchor_element_set_host(e.native(), (*C.gchar)(v))
}

func (e *HTMLAnchorElement) Hostname() strin {
    h := C.webkit_dom_html_anchor_element_get_hostname(e.native())
    return C.GoString((*C.char)(h))
}

func (e *HTMLAnchorElement) SetHostname(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_anchor_element_set_hostname(e.native(), (*C.gchar)(v))
}

func (e *HTMLAnchorElement) Pathname() string {
    p := C.webkit_dom_html_anchor_element_get_pathname(e.native())
    return C.GoString((*C.char)(p))
}

func (e *HTMLAnchorElement) SetPathname(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_anchor_element_set_pathname(e.native(), (*C.gchar)(v))
}

func (e *HTMLAnchorElement) Port() string {
    p := C.webkit_dom_html_anchor_element_get_port(e.native())
    return C.GoString((*C.char)(p))
}

func (e *HTMLAnchorElement) SetPort(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_anchor_element_set_port(e.native())
}

func (e *HTMLAnchorElement) Protocol() string {
    p := C.webkit_dom_html_anchor_element_get_protocol(e.native())
    return C.GoString((*C.char)(p))
}

func (e *HTMLAnchorElement) SetProtocol(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_anchor_element_set_protocol(e.native(), (*C.gchar)(v))
}

func (e *HTMLAnchorElement) Search() string {
    s := C.webkit_dom_html_anchor_element_get_search(e.native())
    return C.GoString((*C.char)(s))
}

func (e *HTMLAnchorElement) SetSearch(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_anchor_element_set_search(e.native(), (*C.gchar)(v))
}

func (e *HTMLAnchorElement) Text() string {
    t := C.webkit_dom_html_anchor_element_get_text(e.native())
    return C.GoString((*C.char)(t))
}

/*
 * HTMLAppletElement
 */

type HTMLAppletElement struct {
    HTMLElement
}

func marshalHTMLAppletElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLAppletElement(obj), nil
}

func wrapHTMLAppletElement(obj *glib.Object) *HTMLAppletElement {
    return &HTMLAppletElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLAppletElement) native() *C.WebKitDOMHTMLAppletElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLAppletElement(p)
}

func newHTMLAppletElement(p unsafe.Pointer) *HTMLAppletElement {
    return wrapHTMLAppletElement(newGlibObject(p));
}

func (e *HTMLAppletElement) Align() string {
    c := C.webkit_dom_html_applet_element_get_align(e.native())
    return C.GoString((*C.char)(c))
}

func (e *HTMLAppletElement) SetAlign(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_applet_element_set_align(e.native(), (*C.gchar)(v))
}

func (e *HTMLAppletElement) Alt() string {
    c := C.webkit_dom_html_applet_element_get_alt(e.native())
    return C.GoString((*C.char)(c))
}

func (e *HTMLAppletElement) SetAlt(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_applet_element_set_alt(e.native(), (*C.gchar)(v))
}

func (e *HTMLAppletElement) Archive() string {
    c := C.webkit_dom_html_applet_element_get_archive(e.native())
    return C.GoString((*C.char)(c))
}

func (e *HTMLAppletElement) SetArchive(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_applet_element_set_archive(e.native(), (*C.gchar)(v))
}

func (e *HTMLAppletElement) Code() string {
    c := C.webkit_dom_html_applet_element_get_code(e.native())
    return C.GoString((*C.char)(c))
}

func (e *HTMLAppletElement) SetCode(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_applet_element_set_code(e.native(), (*C.gchar)(v))
}

func (e *HTMLAppletElement) CodeBase() string {
    c := C.webkit_dom_html_applet_element_get_code_base(e.native())
    return C.GoString((*C.char)(c))
}

func (e *HTMLAppletElement) SetCodeBase(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_applet_element_set_code_base(e.native(), (*C.gchar)(v))
}

func (e *HTMLAppletElement) Height() string {
    c := C.webkit_dom_html_applet_element_get_height(e.native())
    return C.GoString((*C.char)(c))
}

func (e *HTMLAppletElement) SetHeight(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_applet_element_set_height(e.native(), (*C.gchar)(v))
}

func (e *HTMLAppletElement) HSpace() int {
    return int(C.webkit_dom_html_applet_element_get_hspace(e.native()))
}

func (e *HTMLAppletElement) SetHSpace(value int) {
    C.webkit_dom_html_applet_element_set_hspace(e.native(), C.glong(value))
}

func (e *HTMLAppletElement) Name() string {
    c := C.webkit_dom_html_applet_element_get_name(e.native())
    return C.GoString((*C.char)(c))
}

func (e *HTMLAppletElement) SetName(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_applet_element_set_name(e.native(), (*C.gchar)(v))
}

func (e *HTMLAppletElement) Object() string {
    c := C.webkit_dom_html_applet_element_get_object(e.native())
    return C.GoString((*C.char)(c))
}

func (e *HTMLAppletElement) SetObject(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_applet_element_set_object(e.native(), (*C.gchar)(v))
}

func (e *HTMLAppletElement) VSpace() int {
    return int(C.webkit_dom_html_applet_element_get_vspace(e.native()))
}

func (e *HTMLAppletElement) SetVSpace(value int) {
    C.webkit_dom_html_applet_element_set_vspace(e.native(), C.glong(value))
}

func (e *HTMLAppletElement) Width() string {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_applet_element_get_width(e.native(), (*C.gchar)(v))
}

func (e *HTMLAppletElement) SetWidth(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_applet_element_set_width(e.native(), (*C.gchar)(v))
}

/*
 * HTMLAreaElement
 */

type HTMLAreaElement struct {
	HTMLElement
}

func marshalHTMLAreaElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLAreaElement(obj), nil
}

func wrapHTMLAreaElement(obj *glib.Object) *HTMLAreaElement {
    return &HTMLAreaElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLAreaElement) native() *C.WebKitDOMHTMLAreaElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLAreaElement(p)
}

func newHTMLAreaElement(p unsafe.Pointer) *HTMLAreaElement {
    return wrapHTMLAreaElement(newGlibObject(p));
}

// Alt is a wrapper around webkit_dom_html_area_element_get_alt():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-get-alt
func (e *HTMLAreaElement) Alt() string {
    c := C.webkit_dom_html_area_element_get_alt(e.native())
    return C.GoString((*C.char)(c))
}

// SetAlt is a wrapper around webkit_dom_html_area_element_set_alt():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-set-alt
func (e *HTMLAreaElement) SetAlt(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_area_element_set_alt(e.native(), (*C.gchar)(v))
}

// Coords is a wrapper around webkit_dom_html_area_element_get_coords():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-get-coords
func (e *HTMLAreaElement) Coords() string {
    c := C.webkit_dom_html_area_element_get_coords(e.native())
    return C.GoString((*C.char)(c))
}

// SetCoords is a wrapper around webkit_dom_html_area_element_set_coords():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-set-coords
func (e *HTMLAreaElement) SetCoords(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_area_element_set_coords(e.native(), (*C.gchar)(v))
}

// Href is a wrapper around webkit_dom_html_area_element_get_href():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-get-href
func (e *HTMLAreaElement) Href() string {
    h := C.webkit_dom_html_area_element_get_href(e.native())
    return C.GoString((*C.char)(h))
}

// SetHref is a wrapper around webkit_dom_html_area_element_set_href():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-set-href
func (e *HTMLAreaElement) SetHref(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_area_element_set_href(e.native(), (*C.gchar)(v))
}

// NoHref is a wrapper around webkit_dom_html_area_element_get_no_href():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-get-no-href
func (e *HTMLAreaElement) NoHref() bool {
    return gobool(C.webkit_dom_html_area_element_get_no_href(e.native()))
}

// SetNoHref is a wrapper around webkit_dom_html_area_element_set_no_href():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-set-no-href
func (e *HTMLAreaElement) SetNoHref(value bool) {
    C.webkit_dom_html_area_element_set_no_href(e.native(), gboolean(value))
}

// Shape is a wrapper around webkit_dom_html_area_element_get_shape():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-get-shape
func (e *HTMLAreaElement) Shape() string {
    s := C.webkit_dom_html_area_element_get_shape(e.native())
    return C.GoString((*C.char)(s))
}

// SetShape is a wrapper around webkit_dom_html_area_element_set_shape():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-set-shape
func (e *HTMLAreaElement) SetShape(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_area_element_set_shape(e.native(), (*C.gchar)(v))
}

// Target is a wrapper around webkit_dom_html_area_element_get_target():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-get-target
func (e *HTMLAreaElement) Target() string {
    t := C.webkit_dom_html_area_element_get_target(e.native())
    return C.GoString((*C.char)(t))
}

// SetTarget is a wrapper around webkit_dom_html_area_element_set_target():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-set-target
func (e *HTMLAreaElement) SetTarget(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_area_element_set_target(e.native(), (*C.gchar)(v))
}

// Hash is a wrapper around webkit_dom_html_area_element_get_hash():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-get-hash
func (e *HTMLAreaElement) Hash() string {
    h := C.webkit_dom_html_area_element_get_hash(e.native())
    return C.GoString((*C.char)(h))
}

// Host is a wrapper around webkit_dom_html_area_element_get_host():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-get-host
func (e *HTMLAreaElement) Host() string {
    h := C.webkit_dom_html_area_element_get_host(e.native())
    return C.GoString((*C.char)(h))
}

// Hostname is a wrapper around webkit_dom_html_area_element_get_hostname():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-get-hostname
func (e *HTMLAreaElement) Hostname() string {
    h := C.webkit_dom_html_area_element_get_hostname(e.native())
    return C.GoString((*C.char)(h))
}

// Pathname is a wrapper around webkit_dom_html_area_element_get_pathname():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-get-pathname
func (e *HTMLAreaElement) Pathname() string {
    p := C.webkit_dom_html_area_element_get_pathname(e.native())
    return C.GoString((*C.char)(p))
}

// Port is a wrapper around webkit-dom-html-area-element-get-port():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-get-port
func (e *HTMLAreaElement) Port() string {
    p := C.webkit_dom_html_area_element_get_port(e.native())
    return C.GoString((*C.char)(p))
}

// Protocol is a wrapper around webkit_dom_html_area_element_get_protocol():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-get-protocol
func (e *HTMLAreaElement) Protocol() string {
    p := C.webkit_dom_html_area_element_get_protocol(e.native())
    return C.GoString((*C.char)(p))
}

// Search is a wrapper around webkit_dom_html_area_element_get_search():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLAreaElement.html#webkit-dom-html-area-element-get-search
func (e *HTMLAreaElement) Search() string {
    s := C.webkit_dom_html_area_element_get_search(e.native())
    return C.GoString((*C.char)(s))
}


/*
 * HTMLBRElement
 */

type HTMLBRElement struct {
	HTMLElement
}

func marshalHTMLBRElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLBRElement(obj), nil
}

func wrapHTMLBRElement(obj *glib.Object) *HTMLBRElement {
    return &HTMLBRElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLBRElement) native() *C.WebKitDOMHTMLBRElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLBRElement(p)
}

func newHTMLBRElement(p unsafe.Pointer) *HTMLBRElement {
    return wrapHTMLBRElement(newGlibObject(p));
}

// Clear is a wrapper around webkit_dom_html_br_element_get_clear():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBRElement.html#webkit-dom-html-br-element-get-clear
func (e *HTMLBRElement) Clear() string {
    c := C.webkit_dom_html_br_element_get_clear(e.native())
    return C.GoString((*C.char)(c))
}

// SetClear is a wrapper around webkit_dom_html_br_element_set_clear():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBRElement.html#webkit-dom-html-br-element-set-clear
func (e *HTMLBRElement) SetClear(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_br_element_set_clear(e.native(), (*C.gchar)(v))
}

/*
 * HTMLBaseElement
 */

type HTMLBaseElement struct {
	HTMLElement
}

func marshalHTMLBaseElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLBaseElement(obj), nil
}

func wrapHTMLBaseElement(obj *glib.Object) *HTMLBaseElement {
    return &HTMLBaseElement{obj}
}

func (e *HTMLBaseElement) native() *C.WebKitDOMHTMLBaseElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLBaseElement(p)
}

func newHTMLBaseElement(p unsafe.Pointer) *HTMLBaseElement {
    return wrapHTMLBaseElement(newGlibObject(p));
}

// Href is a wrapper around webkit_dom_html_base_element_get_href():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBaseElement.html#webkit-dom-html-base-element-get-href
func (e *HTMLBaseElement) Href() string {
    h := C.webkit_dom_html_base_element_get_href(e.native())
    return C.GoString((*C.char)(h))
}

// SetHref is a wrapper around webkit_dom_html_base_element_set_href():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBaseElement.html#webkit-dom-html-base-element-set-href
func (e *HTMLBaseElement) SetHref(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_base_element_set_href(e.native(), (*C.gchar)(v))
}

// Target is a wrapper around webkit_dom_html_base_element_get_target():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBaseElement.html#webkit-dom-html-base-element-get-target
func (e *HTMLBaseElement) Target() string {
    t := C.webkit_dom_html_base_element_get_target(e.native())
    return C.GoString((*C.char)(t))
}

// SetTarget is a wrapper around webkit_dom_html_base_element_set_target():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBaseElement.html#webkit-dom-html-base-element-set-target
func (e *HTMLBaseElement) SetTarget(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_base_element_set_target(e.native(), (*C.gchar)(v))
}

/*
 * HTMLBaseFontElement
 */

// type HTMLBaseFontElement struct {
// 	HTMLElement
// }

/*
 * HTMLBodyElement
 */

type HTMLBodyElement struct {
	HTMLElement
}

func marshalHTMLBodyElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLBodyElement(obj), nil
}

func wrapHTMLBodyElement(obj *glib.Object) *HTMLBodyElement {
    return &HTMLBodyElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLBodyElement) native() *C.WebKitDOMHTMLBodyElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLBodyElement(p)
}

func newHTMLBodyElement(p unsafe.Pointer) *HTMLBodyElement {
    return wrapHTMLBodyElement(newGlibObject(p));
}

// Alink is a wrapper around webkit_dom_html_body_element_get_a_link():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBodyElement.html#webkit-dom-html-body-element-get-a-link
func (e *HTMLBodyElement) Alink() string {
    cstr := C.webkit_dom_html_body_element_get_a_link(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetAlink is a wrapper around webkit_dom_html_body_element_set_a_link():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBodyElement.html#webkit-dom-html-body-element-set-a-link
func (e *HTMLBodyElement) SetAlink(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_body_element_set_a_link(e.native(), (*C.gchar)(v))
}

// Background is a wrapper around webkit_dom_html_body_element_get_background():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBodyElement.html#webkit-dom-html-body-element-get-background
func (e *HTMLBodyElement) Background() string {
    cstr := C.webkit_dom_html_body_element_get_background(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetBackground is a wrapper around webkit_dom_html_body_element_set_background():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBodyElement.html#webkit-dom-html-body-element-set-background
func (e *HTMLBodyElement) SetBackground(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_body_element_set_background(e.native(), (*C.gchar)(v))
}

// BgColor is a wrapper around webkit_dom_html_body_element_get_bg_color():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBodyElement.html#webkit-dom-html-body-element-get-bg-color
func (e *HTMLBodyElement) BgColor() string {
    cstr := C.webkit_dom_html_body_element_get_bg_color(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetBgColor is a wrapper around webkit_dom_html_body_element_set_background():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBodyElement.html#webkit-dom-html-body-element-set-background
func (e *HTMLBodyElement) SetBgColor(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_body_element_set_bg_color(e.native(), (*C.gchar)(v))
}

// Link is a wrapper around webkit_dom_html_body_element_get_link():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBodyElement.html#webkit-dom-html-body-element-get-link
func (e *HTMLBodyElement) Link() string {
    cstr := C.webkit_dom_html_body_element_get_link(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetLink is a wrapper around webkit_dom_html_body_element_set_link():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBodyElement.html#webkit-dom-html-body-element-set-link
func (e *HTMLBodyElement) SetLink(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_body_element_set_link(e.native(), (*C.gchar)(v))
}

// Text is a wrapper around webkit_dom_html_body_element_get_text():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBodyElement.html#webkit-dom-html-body-element-get-text
func (e *HTMLBodyElement) Text() string {
    cstr := C.webkit_dom_html_body_element_get_text(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetText is a wrapper around webkit_dom_html_body_element_set_text():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBodyElement.html#webkit-dom-html-body-element-set-text
func (e *HTMLBodyElement) SetText(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_body_element_set_text(e.native(), (*C.gchar)(v))
}

// VLink is a wrapper around webkit_dom_html_body_element_get_v_link():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBodyElement.html#webkit-dom-html-body-element-get-v-link
func (e *HTMLBodyElement) VLink() string {
    cstr := C.webkit_dom_html_body_element_get_v_link(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetVLink is a wrapper around webkit_dom_html_body_element_set_v_link():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLBodyElement.html#webkit-dom-html-body-element-set-v-link
func (e *HTMLBodyElement) SetVLink(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_body_element_set_v_link(e.native(), (*C.gchar)(v))
}

/*
 * HTMLButtonElement
 */

type HTMLButtonElement struct {
	HTMLElement
}

func marshalHTMLButtonElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLButtonElement(obj), nil
}

func wrapHTMLButtonElement(obj *glib.Object) *HTMLButtonElement {
    return &HTMLButtonElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLButtonElement) native() *C.WebKitDOMHTMLButtonElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLButtonElement(p)
}

func newHTMLButtonElement(p unsafe.Pointer) *HTMLButtonElement {
    return wrapHTMLButtonElement(newGlibObject(p));
}

// Autofocus is a wrapper around webkit_dom_html_button_element_get_autofocus():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLButtonElement.html#webkit-dom-html-button-element-get-autofocus
func (e *HTMLButtonElement) Autofocus() bool {
    return gobool(C.webkit_dom_html_button_element_get_autofocus(e.native()))
}

// SetAutofocus is a wrapper around webkit_dom_html_button_element_set_autofocus():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLButtonElement.html#webkit-dom-html-button-element-set-autofocus
func (e *HTMLButtonElement) SetAutofocus(value bool) {
    C.webkit_dom_html_button_element_set_autofocus(e.native(), gboolean(value))
}

// Disabled is a wrapper around webkit_dom_html_button_element_get_disabled():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLButtonElement.html#webkit-dom-html-button-element-get-disabled
func (e *HTMLButtonElement) Disabled() bool {
    return gobool(C.webkit_dom_html_button_element_get_disabled(e.native()))
}

// SetDisabled is a wrapper around webkit_dom_html_button_element_set_disabled():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLButtonElement.html#webkit-dom-html-button-element-set-disabled
func (e *HTMLButtonElement) SetDisabled(value bool) {
    C.webkit_dom_html_button_element_set_disabled(e.native(), gboolean(value))
}

// Form is a wrapper around webkit_dom_html_button_element_get_form():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLButtonElement.html#webkit-dom-html-button-element-get-form
func (e *HTMLButtonElement) Form() *HTMLFormElement {
    p := C.webkit_dom_html_button_element_get_form(e.native())
    return newHTMLFormElement(unsafe.Pointer(p))
}

// Name is a wrapper around webkit_dom_html_button_element_get_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLButtonElement.html#webkit-dom-html-button-element-get-name
func (e *HTMLButtonElement) Name() string {
    cstr := C.webkit_dom_html_button_element_get_name(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetName is a wrapper around webkit_dom_html_button_element_set_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLButtonElement.html#webkit-dom-html-button-element-set-name
func (e *HTMLButtonElement) SetName(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_button_element_set_name(e.native(), (*C.gchar)(v))
}

// ButtonType is a wrapper around webkit_dom_html_button_element_get_button_type():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLButtonElement.html#webkit-dom-html-button-element-get-button-type
func (e *HTMLButtonElement) ButtonType() string {
    cstr := C.webkit_dom_html_button_element_get_button_type(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetButtonType is a wrapper around webkit_dom_html_button_element_set_button_type():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLButtonElement.html#webkit-dom-html-button-element-set-button-type
func (e *HTMLButtonElement) SetButtonType(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_button_element_set_button_type(e.native(), (*C.gchar)(v))
}

// Value is a wrapper around webkit_dom_html_button_element_get_value():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLButtonElement.html#webkit-dom-html-button-element-get-value
func (e *HTMLButtonElement) Value() string {
    cstr := C.webkit_dom_html_button_element_get_value(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetValue is a wrapper around webkit_dom_html_button_element_set_value():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLButtonElement.html#webkit-dom-html-button-element-set-value
func (e *HTMLButtonElement) SetValue(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_button_element_set_value(e.native(), (*C.gchar)(v))
}

// WillValidate is a wrapper around webkit_dom_html_button_element_get_will_validate():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLButtonElement.html#webkit-dom-html-button-element-get-will-validate
func (e *HTMLButtonElement) WillValidate() bool {
    return gobool(C.webkit_dom_html_button_element_get_will_validate(e.native()))
}

/*
 * HTMLCanvasElement
 */

// type HTMLCanvasElement struct {
// 	HTMLElement
// }

/*
 * HTMLDListElement
 */

// type HTMLDListElement struct {
// 	HTMLElement
// }

/*
 * HTMLDirectoryElement
 */

// type HTMLDirectoryElement struct {
// 	HTMLElement
// }

/*
 * HTMLDivElement
 */

type HTMLDivElement struct {
	HTMLElement
}

func marshalHTMLDivElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLDivElement(obj), nil
}

func wrapHTMLDivElement(obj *glib.Object) *HTMLDivElement {
    return &HTMLDivElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLDivElement) native() *C.WebKitDOMHTMLDivElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLDivElement(p)
}

func newHTMLDivElement(p unsafe.Pointer) *HTMLDivElement {
    return wrapHTMLDivElement(newGlibObject(p));
}

// Align is a wrapper around webkit_dom_html_div_element_get_align():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDivElement.html#webkit-dom-html-div-element-get-align
func (e *HTMLDivElement) Align() string {
    cstr := C.webkit_dom_html_div_element_get_align(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetAlign is a wrapper around webkit_dom_html_div_element_set_align():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDivElement.html#webkit-dom-html-div-element-set-align
func (e *HTMLDivElement) SetAlign(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_div_element_set_align(e.native(), (*C.gchar)(v))
}

/*
 * HTMLDocument
 */

type HTMLDocument struct {
	Document
}

func marshalHTMLDocument(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLDocument(obj), nil
}

func wrapHTMLDocument(obj *glib.Object) *HTMLDocument {
    return &HTMLDocument{Document{Node{Object{obj}}}}
}

func (e *HTMLDocument) native() *C.WebKitDOMHTMLDocument {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLDocument(p)
}

func newHTMLDocument(p unsafe.Pointer) *HTMLDocument {
    return wrapHTMLDocument(newGlibObject(p));
}

// Close is a wrapper around webkit_dom_html_document_close():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-close
func (d *HTMLDocument) Close() {
    C.webkit_dom_html_document_close(d.native())
}

// Clear is a wrapper around webkit_dom_html_document_clear():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-clear
func (d *HTMLDocument) Clear() {
    C.webkit_dom_html_document_clear(d.native())
}

// CaptureEvents is a wrapper around webkit_dom_html_document_capture_events():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-capture-events
func (d *HTMLDocument) CaptureEvents() {
    C.webkit_dom_html_document_capture_events(d.native())
}

// ReleaseEvents is a wrapper around webkit_dom_html_document_release_events():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-release-events
func (d *HTMLDocument) ReleaseEvents() {
    C.webkit_dom_html_document_release_events(d.native())
}

// Embeds is a wrapper around webkit_dom_html_document_get_embeds():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-get-embeds
func (d *HTMLDocument) Embeds() *HTMLCollection {
    p := C.webkit_dom_html_document_get_embeds(d.native())
    return newHTMLCollection(unsafe.Pointer(p))
}

// Plugins is a wrapper around webkit_dom_html_document_get_plugins():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-get-plugins
func (d *HTMLDocument) Plugins() *HTMLCollection {
    p := C.webkit_dom_html_document_get_plugins(d.native())
    return newHTMLCollection(unsafe.Pointer(p))
}

// Scripts is a wrapper around webkit_dom_html_document_get_scripts():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-get-scripts
func (d *HTMLDocument) Scripts() *HTMLCollection {
    p := C.webkit_dom_html_document_get_scripts(d.native())
    return newHTMLCollection(unsafe.Pointer(p))
}

// Width is a wrapper around webkit_dom_html_document_get_width():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-get-width
func (d *HTMLDocument) Width() int {
    return int(C.webkit_dom_html_document_get_width(d.native()))
}

// Height is a wrapper around webkit_dom_html_document_get_height():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-get-height
func (d *HTMLDocument) Height() int {
    return int(C.webkit_dom_html_document_get_height(d.native()))
}

// Dir is a wrapper around webkit_dom_html_document_get_dir():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-get-dir
func (d *HTMLDocument) Dir() string {
    cstr := C.webkit_dom_html_document_get_dir(d.native())
    return C.GoString((*C.char)(cstr))
}

// SetDir is a wrapper around webkit_dom_html_document_set_dir():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-set-dir
func (d *HTMLDocument) SetDir(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_document_set_dir(d.native(), (*C.gchar)(v))
}

// DesignMode is a wrapper around webkit_dom_html_document_get_design_mode():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-get-design-mode
func (d *HTMLDocument) DesignMode() string {
    cstr := C.webkit_dom_html_document_get_design_mode(d.native())
    return C.GoString((*C.char)(cstr))
}

// SetDesignMode is a wrapper around webkit_dom_html_document_set_design_mode():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-set-design-mode
func (d *HTMLDocument) SetDesignMode(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_document_set_design_mode(d.native(), (*C.gchar)(v))
}

// CompatMode is a wrapper around webkit_dom_html_document_get_compat_mode():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-get-compat-mode
func (d *HTMLDocument) CompatMode() string {
    cstr := C.webkit_dom_html_document_get_compat_mode(d.native())
    return C.GoString((*C.char)(cstr))
}

// BgColor is a wrapper around webkit_dom_html_document_get_bg_color():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-get-bg-color
func (d *HTMLDocument) BgColor() string {
    cstr := C.webkit_dom_html_document_get_bg_color(d.native())
    return C.GoString((*C.char)(cstr))
}

// SetBgColor is a wrapper around webkit_dom_html_document_set_bg_color():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-set-bg-color
func (d *HTMLDocument) SetBgColor(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_document_set_bg_color(d.native(), (*C.gchar)(v))
}

// FgColor is a wrapper around webkit_dom_html_document_get_fg_color():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-get-fg-color
func (d *HTMLDocument) FgColor() string {
    cstr := C.webkit_dom_html_document_get_fg_color(d.native())
    return C.GoString((*C.char)(cstr))
}

// SetFgColor is a wrapper around webkit_dom_html_document_set_fg_color():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-set-fg-color
func (d *HTMLDocument) SetFgColor(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_document_set_fg_color(d.native(), (*C.gchar)(v))
}

// AlinkColor is a wrapper around webkit_dom_html_document_get_alink_color():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-get-alink-color
func (d *HTMLDocument) AlinkColor() string {
    cstr := C.webkit_dom_html_document_get_alink_color(d.native())
    return C.GoString((*C.char)(cstr))
}

// SetAlinkColor is a wrapper around webkit_dom_html_document_set_alink_color():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-set-alink-color
func (d *HTMLDocument) SetAlinkColor(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_document_set_alink_color(d.native(), (*C.gchar)(v))
}

// LinkColor is a wrapper around webkit_dom_html_document_get_link_color():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-get-link-color
func (d *HTMLDocument) LinkColor() string {
    cstr := C.webkit_dom_html_document_get_link_color(d.native())
    return C.GoString((*C.char)(cstr))
}

// SetLinkColor is a wrapper around webkit_dom_html_document_set_link_color():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-set-link-color
func (d *HTMLDocument) SetLinkColor(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_document_set_link_color(d.native(), (*C.gchar)(v))
}

// VlinkColor is a wrapper around webkit_dom_html_document_get_vlink_color():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-get-vlink-color
func (d *HTMLDocument) VlinkColor() string {
    cstr := C.webkit_dom_html_document_get_vlink_color(d.native())
    return C.GoString((*C.char)(cstr))
}

// SetVlinkColor is a wrapper around webkit_dom_html_document_set_vlink_color():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLDocument.html#webkit-dom-html-document-set-vlink-color
func (d *HTMLDocument) SetVlinkColor(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_document_set_vlink_color(d.native(), (*C.gchar)(v))
}

/*
 * HTMLEmbedElement
 */

type HTMLEmbedElement struct {
	HTMLElement
}

func marshalHTMLEmbedElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLEmbedElement(obj), nil
}

func wrapHTMLEmbedElement(obj *glib.Object) *HTMLEmbedElement {
    return &HTMLEmbedElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLEmbedElement) native() *C.WebKitDOMHTMLEmbedElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLEmbedElement(p)
}

func newHTMLEmbedElement(p unsafe.Pointer) *HTMLEmbedElement {
    return wrapHTMLEmbedElement(newGlibObject(p));
}

// Align is a wrapper around webkit_dom_html_embed_element_get_align():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLEmbedElement.html#webkit-dom-html-embed-element-get-align
func (e *HTMLEmbedElement) Align() string {
    cstr := C.webkit_dom_html_embed_element_get_align(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetAlign is a wrapper around webkit_dom_html_embed_element_set_align():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLEmbedElement.html#webkit-dom-html-embed-element-set-align
func (e *HTMLEmbedElement) SetAlign(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_embed_element_set_align(e.native(), (*C.gchar)(cstr))
}

// Height is a wrapper around webkit_dom_html_embed_element_get_height():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLEmbedElement.html#webkit-dom-html-embed-element-get-height
func (e *HTMLEmbedElement) Height() int {
    return int(C.webkit_dom_html_embed_element_get_height(e.native()))
}

// SetHeight is a wrapper around webkit_dom_html_embed_element_set_height():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLEmbedElement.html#webkit-dom-html-embed-element-set-height
func (e *HTMLEmbedElement) SetHeight(value int) {
    C.webkit_dom_html_embed_element_set_height(e.native(), C.glong(value))
}

// Name is a wrapper around webkit_dom_html_embed_element_get_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLEmbedElement.html#webkit-dom-html-embed-element-get-name
func (e *HTMLEmbedElement) Name() string {
    cstr := C.webkit_dom_html_embed_element_get_name(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetName is a wrapper around webkit_dom_html_embed_element_set_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLEmbedElement.html#webkit-dom-html-embed-element-set-name
func (e *HTMLEmbedElement) SetName(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_embed_element_set_name(e.native(), (*C.gchar)(cstr))
}

// Src is a wrapper around webkit_dom_html_embed_element_get_src():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLEmbedElement.html#webkit-dom-html-embed-element-get-src
func (e *HTMLEmbedElement) Src() string {
    cstr := C.webkit_dom_html_embed_element_get_src(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetSrc is a wrapper around webkit_dom_html_embed_element_set_src():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLEmbedElement.html#webkit-dom-html-embed-element-set-src
func (e *HTMLEmbedElement) SetSrc(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_embed_element_set_src(e.native(), (*C.gchar)(cstr))
}

// TypeAttr is a wrapper around webkit_dom_html_embed_element_get_type_attr():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLEmbedElement.html#webkit-dom-html-embed-element-get-type-attr
func (e *HTMLEmbedElement) TypeAttr() string {
    cstr := C.webkit_dom_html_embed_element_get_type_attr(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetTypeAttr is a wrapper around webkit_dom_html_embed_element_set_type_attr():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLEmbedElement.html#webkit-dom-html-embed-element-set-type-attr
func (e *HTMLEmbedElement) SetTypeAttr(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_embed_element_set_type_attr(e.native(), (*C.gchar)(cstr))
}

// Width is a wrapper around webkit_dom_html_embed_element_get_width():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLEmbedElement.html#webkit-dom-html-embed-element-get-width
func (e *HTMLEmbedElement) Width() int {
    return int(C.webkit_dom_html_embed_element_get_width(e.native()))
}

// SetWidth is a wrapper around webkit_dom_html_embed_element_set_width():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLEmbedElement.html#webkit-dom-html-embed-element-set-width
func (e *HTMLEmbedElement) SetWidth(value int) {
    C.webkit_dom_html_embed_element_set_width(e.native(), C.glong(value))
}

/*
 * HTMLFieldSetElement
 */

type HTMLFieldSetElement struct {
	HTMLElement
}

func marshalHTMLFieldSetElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLFieldSetElement(obj), nil
}

func wrapHTMLFieldSetElement(obj *glib.Object) *HTMLFieldSetElement {
    return &HTMLFieldSetElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLFieldSetElement) native() *C.WebKitDOMHTMLFieldSetElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLFieldSetElement(p)
}

func newHTMLFieldSetElement(p unsafe.Pointer) *HTMLFieldSetElement {
    return wrapHTMLFieldSetElement(newGlibObject(p));
}

// Form is a wrapper around webkit_dom_html_field_set_element_get_form():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFieldSetElement.html#webkit-dom-html-field-set-element-get-form
func (e *HTMLFieldSetElement) Form() *HTMLFormElement {
    p := C.webkit_dom_html_field_set_element_get_form(e.native())
    return newHTMLFormElement(unsafe.Pointer(p))
}

/*
 * HTMLFontElement
 */

// type HTMLFontElement struct {
// 	HTMLElement
// }

/*
 * HTMLFormElement
 */

type HTMLFormElement struct {
	HTMLElement
}

func marshalHTMLFormElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLFormElement(obj), nil
}

func wrapHTMLFormElement(obj *glib.Object) *HTMLFormElement {
    return &HTMLFormElement{obj}
}

func (e *HTMLFormElement) native() *C.WebKitDOMHTMLFormElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLFormElement(p)
}

func newHTMLFormElement(p unsafe.Pointer) *HTMLFormElement {
    return wrapHTMLFormElement(newGlibObject(p));
}

// Submit is a wrapper around webkit_dom_html_form_element_submit():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-submit
func (e *HTMLFormElement) Submit() {
    C.webkit_dom_html_form_element_submit(e.native())
}

// Reset is a wrapper around webkit_dom_html_form_element_reset():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-reset
func (e *HTMLFormElement) Reset() {
    C.webkit_dom_html_form_element_reset(e.native())
}

// AcceptCharset is a wrapper around webkit_dom_html_form_element_get_accept_charset():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-get-accept-charset
func (e *HTMLFormElement) AcceptCharset() string {
    c := C.webkit_dom_html_form_element_get_accept_charset(e.native())
    return C.GoString((*C.char)(c))
}

// SetAcceptCharset is a wrapper around webkit_dom_html_form_element_set_accept_charset():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-set-accept-charset
func (e *HTMLFormElement) SetAcceptCharset(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_form_element_set_accept_charset(e.native(), (*C.gchar)(v))
}

// Action is a wrapper around webkit_dom_html_form_element_get_action():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-get-action
func (e *HTMLFormElement) Action() string {
    a := C.webkit_dom_html_form_element_get_action(e.native())
    return C.GoString((*C.char)(a))
}

// SetAction is a wrapper around webkit_dom_html_form_element_set_action():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-set-action
func (e *HTMLFormElement) SetAction(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_form_element_set_action(e.native(), (*C.gchar)(v))
}

// Enctype is a wrapper around webkit_dom_html_form_element_get_enctype():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-get-enctype
func (e *HTMLFormElement) Enctype() string {
    e := C.webkit_dom_html_form_element_get_enctype(e.native())
    return C.GoString((*C.char)(e))
}

// SetEnctype is a wrapper around webkit_dom_html_form_element_set_enctype():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-set-enctype
func (e *HTMLFormElement) SetEnctype(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_form_element_set_enctype(e.native(), (*C.gchar)(v))
}

// Encoding is a wrapper around webkit_dom_html_form_element_get_encoding():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-get-encoding
func (e *HTMLFormElement) Encoding() string {
    e := C.webkit_dom_html_form_element_get_encoding(e.native())
    return C.GoString((*C.char)(e))
}

// SetEncoding is a wrapper around webkit_dom_html_form_element_set_encoding():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-set-encoding
func (e *HTMLFormElement) SetEncoding(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_form_element_set_encoding(e.native(), (*C.gchar)(v))
}

// Method is a wrapper around webkit_dom_html_form_element_get_method():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-get-method
func (e *HTMLFormElement) Method() string {
    m := C.webkit_dom_html_form_element_get_method(e.native())
    return C.GoString((*C.char)(m))
}

// SetMethod is a wrapper around webkit_dom_html_form_element_set_method():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-set-method
func (e *HTMLFormElement) SetMethod(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_form_element_set_method(e.native(), (*C.gchar)(v))
}

// Name is a wrapper around webkit_dom_html_form_element_get_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-get-name
func (e *HTMLFormElement) Name() string {
    n := C.webkit_dom_html_form_element_get_name(e.native())
    return C.GoString((*C.char)(n))
}

// SetName is a wrapper around webkit_dom_html_form_element_set_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-set-name
func (e *HTMLFormElement) SetName(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_form_element_set_name(e.native(), (*C.gchar)(v))
}

// Target is a wrapper around webkit_dom_html_form_element_get_target():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-get-target
func (e *HTMLFormElement) Target() string {
    t := C.webkit_dom_html_form_element_get_target(e.native())
    return C.GoString((*C.char)(t))
}

// SetTarget is a wrapper around webkit_dom_html_form_element_set_target():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-set-target
func (e *HTMLFormElement) SetTarget(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_form_element_set_target(e.native(), (*C.gchar)(v))
}

// Elements is a wrapper around webkit_dom_html_form_element_get_elements():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-get-elements
func (e *HTMLFormElement) Elements() *HTMLCollection {
    c := C.webkit_dom_html_form_element_get_elements(e.native())
    return newHTMLCollection(unsafe.Pointer(c))
}

// Len is a wrapper around webkit_dom_html_form_element_get_length():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFormElement.html#webkit-dom-html-form-element-get-length
func (e *HTMLFormElement) Len() int {
    return int(C.webkit_dom_html_form_element_get_length(e.native()))
}

/*
 * HTMLFrameElement
 */

type HTMLFrameElement struct {
	HTMLElement
}

func marshalHTMLFrameElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLFrameElement(obj), nil
}

func wrapHTMLFrameElement(obj *glib.Object) *HTMLFrameElement {
    return &HTMLFrameElement{obj}
}

func (e *HTMLFrameElement) native() *C.WebKitDOMHTMLFrameElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLFrameElement(p)
}

func newHTMLFrameElement(p unsafe.Pointer) *HTMLFrameElement {
    return wrapHTMLFrameElement(newGlibObject(p));
}

// FrameBorder is a wrapper around webkit_dom_html_frame_element_get_frame_border():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-get-frame-border
func (e *HTMLFrameElement) FrameBorder() string {
    cstr := C.webkit_dom_html_frame_element_get_frame_border(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetFrameBorder is a wrapper around webkit_dom_html_frame_element_set_frame_border():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-set-frame-border
func (e *HTMLFrameElement) SetFrameBorder(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_frame_element_set_frame_border(e.native(), (*C.gchar)(cstr))
}

// LongDesc is a wrapper around webkit_dom_html_frame_element_get_long_desc():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-get-long-desc
func (e *HTMLFrameElement) LongDesc() string {
    cstr := C.webkit_dom_html_frame_element_get_long_desc(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetLongDesc is a wrapper around webkit_dom_html_frame_element_set_long_desc():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-set-long-desc
func (e *HTMLFrameElement) SetLongDesc(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_frame_element_set_long_desc(e.native(), (*C.gchar)(cstr))
}

// MarginHeight is a wrapper around webkit_dom_html_frame_element_get_margin_height():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-get-margin-height
func (e *HTMLFrameElement) MarginHeight() string {
    cstr := C.webkit_dom_html_frame_element_get_margin_height(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetMarginHeight is a wrapper around webkit_dom_html_frame_element_set_margin_height():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-set-margin-height
func (e *HTMLFrameElement) SetMarginHeight(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_frame_element_set_margin_height(e.native(), (*C.gchar)(cstr))
}

// MarginWidth is a wrapper around webkit_dom_html_frame_element_get_margin_width():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-get-margin-width
func (e *HTMLFrameElement) MarginWidth() string {
    cstr := C.webkit_dom_html_frame_element_get_margin_width(e.native())
    return C.GoString((*C.char)(cstr))
}

// Name is a wrapper around webkit_dom_html_frame_element_get_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-get-name
func (e *HTMLFrameElement) Name() string {
    cstr := C.webkit_dom_html_frame_element_get_name(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetName is a wrapper around webkit_dom_html_frame_element_set_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-set-name
func (e *HTMLFrameElement) SetName(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_frame_element_set_name(e.native(), (*C.gchar)(cstr))
}

// NoResize is a wrapper around webkit_dom_html_frame_element_get_no_resize():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-get-no-resize
func (e *HTMLFrameElement) NoResize() bool {
    return gobool(C.webkit_dom_html_frame_element_get_no_resize(e.native()))
}

// SetNoResize is a wrapper around webkit_dom_html_frame_element_set_no_resize():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-set-no-resize
func (e *HTMLFrameElement) SetNoResize(value bool) {
    C.webkit_dom_html_frame_element_set_no_resize(e.native(), gboolean(value))
}

// Scrolling is a wrapper around webkit_dom_html_frame_element_get_scrolling():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-get-scrolling
func (e *HTMLFrameElement) Scrolling() string {
    cstr := C.webkit_dom_html_frame_element_get_scrolling(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetScrolling is a wrapper around webkit_dom_html_frame_element_set_scrolling():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-set-scrolling
func (e *HTMLFrameElement) SetScrolling(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_frame_element_set_scrolling(e.native(), (*C.gchar)(cstr))
}

// Src is a wrapper around webkit_dom_html_frame_element_get_src():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-get-src
func (e *HTMLFrameElement) Src() string {
    cstr := C.webkit_dom_html_frame_element_get_src(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetSrc is a wrapper around webkit_dom_html_frame_element_set_src():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-set-src
func (e *HTMLFrameElement) SetSrc(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_frame_element_set_src(e.native(), (*C.gchar)(cstr))
}

// ContentDocument is a wrapper around webkit_dom_html_frame_element_get_content_document():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-get-content-document
func (e *HTMLFrameElement) ContentDocument() *Document {
    p := C.webkit_dom_html_frame_element_get_content_document(e.native())
    return newDocument(unsafe.Pointer(p))
}

// ContentWindow is a wrapper around webkit_dom_html_frame_element_get_content_window():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-get-content-window
func (e *HTMLFrameElement) ContentWindow() *Window {
    p := C.webkit_dom_html_frame_element_get_content_window(e.native())
    return newWindow(unsafe.Pointer(p))
}

// Width is a wrapper around webkit_dom_html_frame_element_get_width():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-get-width
func (e *HTMLFrameElement) Width() int {
    return int(C.webkit_dom_html_frame_element_get_width(e.native()))
}

// Height is a wrapper around webkit_dom_html_frame_element_get_height():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameElement.html#webkit-dom-html-frame-element-get-height
func (e *HTMLFrameElement) Height() int {
    return int(C.webkit_dom_html_frame_element_get_height(e.native()))
}

/*
 * HTMLFrameSetElement
 */

type HTMLFrameSetElement struct {
	HTMLElement
}

func marshalHTMLFrameSetElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLFrameSetElement(obj), nil
}

func wrapHTMLFrameSetElement(obj *glib.Object) *HTMLFrameSetElement {
    return &HTMLFrameSetElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLFrameSetElement) native() *C.WebKitDOMHTMLFrameSetElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLFrameSetElement(p)
}

func newHTMLFrameSetElement(p unsafe.Pointer) *HTMLFrameSetElement {
    return wrapHTMLFrameSetElement(newGlibObject(p));
}

// Cols is a wrapper around webkit_dom_html_frame_set_element_get_cols():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameSetElement.html#webkit-dom-html-frame-set-element-get-cols
func (e *HTMLFrameSetElement) Cols() string {
    cstr := C.webkit_dom_html_frame_set_element_get_cols(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetCols is a wrapper around webkit_dom_html_frame_set_element_set_cols():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameSetElement.html#webkit-dom-html-frame-set-element-set-cols
func (e *HTMLFrameSetElement) SetCols(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_frame_set_element_set_cols(e.native(), (*C.gchar)(cstr))
}

// Rows is a wrapper around webkit_dom_html_frame_set_element_get_rows():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameSetElement.html#webkit-dom-html-frame-set-element-get-rows
func (e *HTMLFrameSetElement) Rows() string {
    cstr := C.webkit_dom_html_frame_set_element_get_rows(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetRows is a wrapper around webkit_dom_html_frame_set_element_set_rows():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLFrameSetElement.html#webkit-dom-html-frame-set-element-set-rows
func (e *HTMLFrameSetElement) SetRows(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_frame_set_element_set_rows(e.native(), (*C.gchar)(cstr))
}

/*
 * HTMLHRElement
 */

// type HTMLHRElement struct {
// 	HTMLElement
// }

/*
 * HTMLHeadElement
 */

type HTMLHeadElement struct {
	HTMLElement
}

func marshalHTMLHeadElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLHeadElement(obj), nil
}

func wrapHTMLHeadElement(obj *glib.Object) *HTMLHeadElement {
    return &HTMLHeadElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLHeadElement) native() *C.WebKitDOMHTMLHeadElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLHeadElement(p)
}

func newHTMLHeadElement(p unsafe.Pointer) *HTMLHeadElement {
    return wrapHTMLHeadElement(newGlibObject(p));
}

// Profile is a wrapper around webkit_dom_html_head_element_get_profile():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLHeadElement.html#webkit-dom-html-head-element-get-profile
func (e *HTMLHeadElement) Profile() string {
    cstr := C.webkit_dom_html_head_element_get_profile(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetProfile is a wrapper around webkit_dom_html_head_element_set_profile():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLHeadElement.html#webkit-dom-html-head-element-set-profile
func (e *HTMLHeadElement) SetProfile(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_head_element_set_profile(e.native(), (*C.gchar)(cstr))
}

/*
 * HTMLHeadingElement
 */

// type HTMLHeadingElement struct {
// 	HTMLElement
// }

/*
 * HTMLHtmlElement
 */

// type HTMLHtmlElement struct {
// 	HTMLElement
// }

/*
 * HTMLIFrameElement
 */

// type HTMLIFrameElement struct {
// 	HTMLElement
// }

/*
 * HTMLImageElement
 */

type HTMLImageElement struct {
	HTMLElement
}

func marshalHTMLImageElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLImageElement(obj), nil
}

func wrapHTMLImageElement(obj *glib.Object) *HTMLImageElement {
    return &HTMLImageElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLImageElement) native() *C.WebKitDOMHTMLImageElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLImageElement(p)
}

func newHTMLImageElement(p unsafe.Pointer) *HTMLImageElement {
    return wrapHTMLImageElement(newGlibObject(p));
}

// Name is a wrapper around webkit_dom_html_image_element_get_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-name
func (e *HTMLImageElement) Name() string {
    cstr := C.webkit_dom_html_image_element_get_name(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetName is a wrapper around webkit_dom_html_image_element_set_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-set-name
func (e *HTMLImageElement) SetName(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_image_element_set_name(e.native(), (*C.gchar)(cstr))
}

// Align is a wrapper around webkit_dom_html_image_element_get_align():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-align
func (e *HTMLImageElement) Align() string {
    cstr := C.webkit_dom_html_image_element_get_align(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetAlign is a wrapper around webkit_dom_html_image_element_set_align():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-set-align
func (e *HTMLImageElement) SetAlign(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_image_element_set_align(e.native(), (*C.gchar)(cstr))
}

// Alt is a wrapper around webkit_dom_html_image_element_get_alt():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-alt
func (e *HTMLImageElement) Alt() string {
    cstr := C.webkit_dom_html_image_element_get_alt(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetAlt is a wrapper around webkit_dom_html_image_element_set_alt():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-set-alt
func (e *HTMLImageElement) SetAlt(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_image_element_set_alt(e.native(), (*C.gchar)(cstr))
}

// Border is a wrapper around webkit_dom_html_image_element_get_border():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-border
func (e *HTMLImageElement) Border() string {
    cstr := C.webkit_dom_html_image_element_get_border(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetBorder is a wrapper around webkit_dom_html_image_element_set_border():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-set-border
func (e *HTMLImageElement) SetBorder(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_image_element_set_border(e.native(), (*C.gchar)(cstr))
}

// Height is a wrapper around webkit_dom_html_image_element_get_height():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-height
func (e *HTMLImageElement) Height() int {
    return int(C.webkit_dom_html_image_element_get_height(e.native()))
}

// SetHeight is a wrapper around webkit_dom_html_image_element_set_height():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-set-height
func (e *HTMLImageElement) SetHeight(value int) {
    C.webkit_dom_html_image_element_set_height(e.native(), C.glong(value))
}

// Hspace is a wrapper around webkit_dom_html_image_element_get_hspace():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-hspace
func (e *HTMLImageElement) Hspace() int {
    return uint(C.webkit_dom_html_image_element_get_hspace(e.native()))
}

// SetHspace is a wrapper around webkit_dom_html_image_element_set_hspace():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-set-hspace
func (e *HTMLImageElement) SetHspace(value int) {
    C.webkit_dom_html_image_element_set_hspace(e.native(), C.glong(value))
}

// IsMap is a wrapper around webkit_dom_html_image_element_get_is_map():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-is-map
func (e *HTMLImageElement) IsMap() bool {
    return gobool(C.webkit_dom_html_image_element_get_is_map(e.native()))
}

// SetIsMap is a wrapper around webkit_dom_html_image_element_set_is_map():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-set-is-map
func (e *HTMLImageElement) SetIsMap(value bool) {
    C.webkit_dom_html_image_element_set_is_map(e.native(), gboolean(value))
}

// LongDesc is a wrapper around webkit_dom_html_image_element_get_long_desc():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-long-desc
func (e *HTMLImageElement) LongDesc() string {
    cstr := C.webkit_dom_html_image_element_get_long_desc(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetLongDesc is a wrapper around webkit_dom_html_image_element_set_long_desc():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-set-long-desc
func (e *HTMLImageElement) SetLongDesc(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_image_element_set_long_desc(e.native(), (*C.gchar)(cstr))
}

// Src is a wrapper around webkit_dom_html_image_element_get_src():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-src
func (e *HTMLImageElement) Src() string {
    cstr := C.webkit_dom_html_image_element_get_src(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetSrc is a wrapper around webkit_dom_html_image_element_set_src():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-set-src
func (e *HTMLImageElement) SetSrc(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_image_element_set_src(e.native(), (*C.gchar)(cstr))
}

// UseMap is a wrapper around webkit_dom_html_image_element_get_use_map():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-use-map
func (e *HTMLImageElement) UseMap() string {
    cstr := C.webkit_dom_html_image_element_get_use_map(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetUseMap is a wrapper around webkit_dom_html_image_element_set_use_map():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-set-use-map
func (e *HTMLImageElement) SetUseMap(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_image_element_set_use_map(e.native(), (*C.gchar)(cstr))
}

// Vspace is a wrapper around webkit_dom_html_image_element_get_vspace():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-vspace
func (e *HTMLImageElement) Vspace() int {
    return int(C.webkit_dom_html_image_element_get_vspace(e.native()))
}

// SetVspace is a wrapper around webkit_dom_html_image_element_set_vspace():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-set-vspace
func (e *HTMLImageElement) SetVspace(value int) {
    C.webkit_dom_html_image_element_set_vspace(e.native(), C.glong(value))
}

// Width is a wrapper around webkit_dom_html_image_element_get_width():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-width
func (e *HTMLImageElement) Width() int {
    return int(C.webkit_dom_html_image_element_get_width(e.native()))
}

// SetWidth is a wrapper around webkit_dom_html_image_element_set_width():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-set-width
func (e *HTMLImageElement) SetWidth(value int) {
    C.webkit_dom_html_image_element_set_width(e.native(), C.glong(value))
}

// Complete is a wrapper around webkit_dom_html_image_element_get_complete():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-complete
func (e *HTMLImageElement) Complete() bool {
    return gobool(C.webkit_dom_html_image_element_get_complete(e.native()))
}

// Lowsrc is a wrapper around webkit_dom_html_image_element_get_lowsrc():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-lowsrc
func (e *HTMLImageElement) Lowsrc() string {
    cstr := C.webkit_dom_html_image_element_get_lowsrc(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetLowsrc is a wrapper around webkit_dom_html_image_element_set_lowsrc():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-set-lowsrc
func (e *HTMLImageElement) SetLowsrc(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_image_element_set_lowsrc(e.native(), (*C.gchar)(cstr))
}

// NaturalHeight is a wrapper around webkit_dom_html_image_element_get_natural_height():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-natural-height
func (e *HTMLImageElement) NaturalHeight() int {
    return int(C.webkit_dom_html_image_element_get_natural_height(e.native()))
}

// NaturalWidth is a wrapper around webkit_dom_html_image_element_get_natural_width():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-natural-width
func (e *HTMLImageElement) NaturalWidth() int {
    return int(C.webkit_dom_html_image_element_get_natural_width(e.native()))
}

// X is a wrapper around webkit_dom_html_image_element_get_x():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-x
func (e *HTMLImageElement) X() int {
    return int(C.webkit_dom_html_image_element_get_x(e.native()))
}

// Y is a wrapper around webkit_dom_html_image_element_get_y():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLImageElement.html#webkit-dom-html-image-element-get-y
func (e *HTMLImageElement) Y() int {
    return int(C.webkit_dom_html_image_element_get_y(e.native()))
}

/*
 * HTMLInputElement
 */

type HTMLInputElement struct {
	HTMLElement
}

func marshalHTMLInputElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLInputElement(obj), nil
}

func wrapHTMLInputElement(obj *glib.Object) *HTMLInputElement {
    return &HTMLInputElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLInputElement) native() *C.WebKitDOMHTMLInputElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLInputElement(p)
}

func newHTMLInputElement(p unsafe.Pointer) *HTMLInputElement {
    return wrapHTMLInputElement(newGlibObject(p));
}

// Select is a wrapper around webkit_dom_html_input_element_select():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-select
func (e *HTMLInputElement) Select() {
    C.webkit_dom_html_input_element_select(e.native())
}

// Accept is a wrapper around webkit_dom_html_input_element_get_accept():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-accept
func (e *HTMLInputElement) Accept() string {
    cstr := C.webkit_dom_html_input_element_get_accept(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetAccept is a wrapper around webkit_dom_html_input_element_set_accept():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-accept
func (e *HTMLInputElement) SetAccept(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_input_element_set_accept(e.native(), (*C.gchar)(cstr))
}

// Alt is a wrapper around webkit_dom_html_input_element_get_alt():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-alt
func (e *HTMLInputElement) Alt() string {
    cstr := C.webkit_dom_html_input_element_get_alt(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetAlt is a wrapper around webkit_dom_html_input_element_set_alt():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-alt
func (e *HTMLInputElement) SetAlt(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_input_element_set_alt(e.native(), (*C.gchar)(cstr))
}

// Autofocus is a wrapper around webkit_dom_html_input_element_get_autofocus():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-autofocus
func (e *HTMLInputElement) Autofocus() bool {
    return gobool(C.webkit_dom_html_input_element_get_autofocus(e.native()))
}

// SetAutofocus is a wrapper around webkit_dom_html_input_element_set_autofocus():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-autofocus
func (e *HTMLInputElement) SetAutofocus(value bool) {
    C.webkit_dom_html_input_element_set_autofocus(e.native(), gboolean(value))
}

// DefaultChecked is a wrapper around webkit_dom_html_input_element_get_default_checked():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-default-checked
func (e *HTMLInputElement) DefaultChecked() bool {
    return gobool(C.webkit_dom_html_input_element_get_default_checked(e.native()))
}

// Checked is a wrapper around webkit_dom_html_input_element_get_checked():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-checked
func (e *HTMLInputElement) Checked() bool {
    return gobool(C.webkit_dom_html_input_element_get_checked(e.native()))
}

// SetChecked is a wrapper around webkit_dom_html_input_element_set_checked():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-checked
func (e *HTMLInputElement) SetChecked(value bool) {
    C.webkit_dom_html_input_element_set_checked(e.native(), gboolean(value))
}

// Disabled is a wrapper around webkit_dom_html_input_element_get_disabled():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-disabled
func (e *HTMLInputElement) Disabled() bool {
    return gobool(C.webkit_dom_html_input_element_get_disabled(e.native()))
}

// SetDisabled is a wrapper around webkit_dom_html_input_element_set_disabled():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-disabled
func (e *HTMLInputElement) SetDisabled(value bool) {
    C.webkit_dom_html_input_element_set_disabled(e.native(), gboolean(value))
}

// Form is a wrapper around webkit_dom_html_input_element_get_form():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-form
func (e *HTMLInputElement) Form() *HTMLFormElement {
    p := C.webkit_dom_html_input_element_get_form(e.native())
    return newHTMLFormElement(unsafe.Pointer(p))
}

// Files is a wrapper around webkit_dom_html_input_element_get_files():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-files
func (e *HTMLInputElement) Files() *FileList {
    p := C.webkit_dom_html_input_element_get_files(e.native())
    return newFileList(unsafe.Pointer(p))
}

// SetFiles is a wrapper around webkit_dom_html_input_element_set_files():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-files
func (e *HTMLInputElement) SetFiles(value *FileList) {
    C.webkit_dom_html_input_element_set_files(e.native(), value.native())
}

// Height is a wrapper around webkit_dom_html_input_element_get_height():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-height
func (e *HTMLInputElement) Height() uint {
    return uint(C.webkit_dom_html_input_element_get_height(e.native()))
}

// SetHeight is a wrapper around webkit_dom_html_input_element_set_height():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-height
func (e *HTMLInputElement) SetHeight(value uint) {
    C.webkit_dom_html_input_element_set_height(e.native(), C.gulong(value))
}

// Indeterminate is a wrapper around webkit_dom_html_input_element_get_indeterminate():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-indeterminate
func (e *HTMLInputElement) Indeterminate() bool {
    return gobool(C.webkit_dom_html_input_element_get_indeterminate(e.native()))
}

// SetIndeterminate is a wrapper around webkit_dom_html_input_element_set_indeterminate():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-indeterminate
func (e *HTMLInputElement) SetIndeterminate(value bool) {
    C.webkit_dom_html_input_element_set_indeterminate(e.native(), gboolean(value))
}

// MaxLen is a wrapper around webkit_dom_html_input_element_get_max_length():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-max-length
func (e *HTMLInputElement) MaxLen() int {
    return int(C.webkit_dom_html_input_element_get_max_length(e.native()))
}

// SetMaxLen is a wrapper around webkit_dom_html_input_element_set_max_length():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-max-length
func (e *HTMLInputElement) SetMaxLen(value int) error {
    var err *C.GError
    C.webkit_dom_html_input_element_set_max_length(e.native(), C.glong(int), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

// Multiple is a wrapper around webkit_dom_html_input_element_get_multiple():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-multiple
func (e *HTMLInputElement) Multiple() bool {
    return gobool(C.webkit_dom_html_input_element_get_multiple(e.native()))
}

// SetMultiple is a wrapper around webkit_dom_html_input_element_set_multiple():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-multiple
func (e *HTMLInputElement) SetMultiple(value bool) {
    C.webkit_dom_html_input_element_set_multiple(e.native(), gboolean(value))
}

// Name is a wrapper around webkit_dom_html_input_element_get_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-name
func (e *HTMLInputElement) Name() string {
    cstr := C.webkit_dom_html_input_element_get_name(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetName is a wrapper around webkit_dom_html_input_element_set_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-name
func (e *HTMLInputElement) SetName(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_input_element_set_name(e.native(), (*C.gchar)(cstr))
}

// SetReadOnly is a wrapper around webkit_dom_html_input_element_set_read_only():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-read-only
func (e *HTMLInputElement) SetReadOnly(value bool) {
    C.webkit_dom_html_input_element_set_read_only(e.native(), gboolean(value))
}

// Size is a wrapper around webkit_dom_html_input_element_get_size():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-size
func (e *HTMLInputElement) Size() uint {
    return uint(C.webkit_dom_html_input_element_get_size(e.native()))
}

// SetSize is a wrapper around webkit_dom_html_input_element_set_size():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-size
func (e *HTMLInputElement) SetSize(uint value) error {
    var err *C.GError
    C.webkit_dom_html_input_element_set_size(e.native(), C.gulong(value), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

// Src is a wrapper around webkit_dom_html_input_element_get_src():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-src
func (e *HTMLInputElement) Src() string {
    cstr := C.webkit_dom_html_input_element_get_src(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetSrc is a wrapper around webkit_dom_html_input_element_set_src():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-src
func (e *HTMLInputElement) SetSrc(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_input_element_set_src(e.native(), (*C.gchar)(cstr))
}

// InputType is a wrapper around webkit_dom_html_input_element_get_input_type():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-input-type
func (e *HTMLInputElement) InputType() string {
    cstr := C.webkit_dom_html_input_element_get_input_type(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetInputType is a wrapper around webkit_dom_html_input_element_set_input_type():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-input-type
func (e *HTMLInputElement) SetInputType(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_input_element_set_input_type(e.native(), (*C.gchar)(cstr))
}

// DefaultValue is a wrapper around webkit_dom_html_input_element_get_default_value():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-default-value
func (e *HTMLInputElement) DefaultValue() string {
    cstr := C.webkit_dom_html_input_element_get_default_value(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetDefaultValue is a wrapper around webkit_dom_html_input_element_set_default_value():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-default-value
func (e *HTMLInputElement) SetDefaultValue(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_input_element_set_default_value(e.native(), (*C.gchar)(cstr))
}

// Value is a wrapper around webkit_dom_html_input_element_get_value():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-value
func (e *HTMLInputElement) Value() string {
    cstr := C.webkit_dom_html_input_element_get_value(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetValue is a wrapper around webkit_dom_html_input_element_set_value():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-value
func (e *HTMLInputElement) SetValue(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_input_element_set_value(e.native(), (*C.gchar)(cstr))
}

// Width is a wrapper around webkit_dom_html_input_element_get_width():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-width
func (e *HTMLInputElement) Width() uint {
    return uint(C.webkit_dom_html_input_element_get_width(e.native()))
}

// SetWidth is a wrapper around webkit_dom_html_input_element_set_width():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-width
func (e *HTMLInputElement) SetWidth(value uint) {
    C.webkit_dom_html_input_element_set_width(e.native(), C.gulong(value))
}

// WillValidate is a wrapper around webkit_dom_html_input_element_get_will_validate():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-will-validate
func (e *HTMLInputElement) WillValidate() bool {
    return gobool(C.webkit_dom_html_input_element_get_will_validate(e.native()))
}

// Align is a wrapper around webkit_dom_html_input_element_get_align():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-align
func (e *HTMLInputElement) Align() string {
    cstr := C.webkit_dom_html_input_element_get_align(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetAlign is a wrapper around webkit_dom_html_input_element_set_align():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-align
func (e *HTMLInputElement) SetAlign(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_input_element_set_align(e.native(), (*C.gchar)(cstr))
}

// UseMap is a wrapper around webkit_dom_html_input_element_get_use_map():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-use-map
func (e *HTMLInputElement) UseMap() string {
    cstr := C.webkit_dom_html_input_element_get_use_map(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetUseMap is a wrapper around webkit_dom_html_input_element_set_use_map():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-set-use-map
func (e *HTMLInputElement) SetUseMap(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_input_element_set_use_map(e.native(), (*C.gchar)(cstr))
}

// Capture is a wrapper around webkit_dom_html_input_element_get_capture():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLInputElement.html#webkit-dom-html-input-element-get-capture
func (e *HTMLInputElement) Capture() bool {
    return gobool(C.webkit_dom_html_input_element_get_capture(e.native()))
}

/*
 * HTMLLIElement
 */

// type HTMLLIElement struct {
// 	HTMLElement
// }

/*
 * HTMLLabelElement
 */

type HTMLLabelElement struct {
	HTMLElement
}

func marshalHTMLLabelElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLLabelElement(obj), nil
}

func wrapHTMLLabelElement(obj *glib.Object) *HTMLLabelElement {
    return &HTMLLabelElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLLabelElement) native() *C.WebKitDOMHTMLLabelElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLLabelElement(p)
}

func newHTMLLabelElement(p unsafe.Pointer) *HTMLLabelElement {
    return wrapHTMLLabelElement(newGlibObject(p));
}

// Form is a wrapper around webkit_dom_html_label_element_get_form():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLabelElement.html#webkit-dom-html-label-element-get-form
func (e *HTMLLabelElement) Form() *HTMLFormElement {
    p := C.webkit_dom_html_label_element_get_form(e.native())
    return newHTMLFormElement(unsafe.Pointer(p))
}

// HTMLFor is a wrapper around webkit_dom_html_label_element_get_html_for():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLabelElement.html#webkit-dom-html-label-element-get-html-for
func (e *HTMLLabelElement) HTMLFor() string {
    cstr := C.webkit_dom_html_label_element_get_html_for(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetHTMLFor is a wrapper around webkit_dom_html_label_element_set_html_for():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLabelElement.html#webkit-dom-html-label-element-set-html-for
func (e *HTMLLabelElement) SetHTMLFor(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_label_element_set_html_for(e.native(), (*C.gchar)(cstr))
}

/*
 * HTMLLegendElement
 */

// type HTMLLegendElement struct {
// 	HTMLElement
// }

/*
 * HTMLLinkElement
 */

type HTMLLinkElement struct {
	HTMLElement
}

func marshalHTMLLinkElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLLinkElement(obj), nil
}

func wrapHTMLLinkElement(obj *glib.Object) *HTMLLinkElement {
    return &HTMLLinkElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLLinkElement) native() *C.WebKitDOMHTMLLinkElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLLinkElement(p)
}

func newHTMLLinkElement(p unsafe.Pointer) *HTMLLinkElement {
    return wrapHTMLLinkElement(newGlibObject(p));
}

// Disabled is a wrapper around webkit_dom_html_link_element_get_disabled():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-get-disabled
func (e *HTMLLinkElement) Disabled() bool {
    return gobool(C.webkit_dom_html_link_element_get_disabled(e.native()))
}

// SetDisabled is a wrapper around webkit_dom_html_link_element_set_disabled():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-set-disabled
func (e *HTMLLinkElement) SetDisabled(value bool) {
    C.webkit_dom_html_link_element_set_disabled(e.native(), gboolean(value))
}

// Charset is a wrapper around webkit_dom_html_link_element_get_charset():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-get-charset
func (e *HTMLLinkElement) Charset() string {
    cstr := C.webkit_dom_html_link_element_get_charset(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetCharset is a wrapper around webkit_dom_html_link_element_set_charset():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-set-charset
func (e *HTMLLinkElement) SetCharset(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_link_element_set_charset(e.native(), (*C.gchar)(cstr))
}

// Href is a wrapper around webkit_dom_html_link_element_get_href():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-get-href
func (e *HTMLLinkElement) Href() string {
    cstr := C.webkit_dom_html_link_element_get_href(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetHref is a wrapper around webkit_dom_html_link_element_set_href():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-set-href
func (e *HTMLLinkElement) SetHref(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_link_element_set_href(e.native(), (*C.gchar)(cstr))
}

// Hreflang is a wrapper around webkit_dom_html_link_element_get_hreflang():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-get-hreflang
func (e *HTMLLinkElement) Hreflang() string {
    cstr := C.webkit_dom_html_link_element_get_hreflang(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetHreflang is a wrapper around webkit_dom_html_link_element_set_hreflang():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-set-hreflang
func (e *HTMLLinkElement) SetHreflang(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_link_element_set_hreflang(e.native(), (*C.gchar)(cstr))
}

// Media is a wrapper around webkit_dom_html_link_element_get_media():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-get-media
func (e *HTMLLinkElement) Media() string {
    cstr := C.webkit_dom_html_link_element_get_media(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetMedia is a wrapper around webkit_dom_html_link_element_set_media():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-set-media
func (e *HTMLLinkElement) SetMedia(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_link_element_set_media(e.native(), (*C.gchar)(cstr))
}

// Rel is a wrapper around webkit_dom_html_link_element_get_rel():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-get-rel
func (e *HTMLLinkElement) Rel() string {
    cstr := C.webkit_dom_html_link_element_get_rel(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetRel is a wrapper around webkit_dom_html_link_element_set_rel():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-set-rel
func (e *HTMLLinkElement) SetRel(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_link_element_set_rel(e.native(), (*C.gchar)(cstr))
}

// Rev is a wrapper around webkit_dom_html_link_element_get_rev():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-get-rev
func (e *HTMLLinkElement) Rev() string {
    cstr := C.webkit_dom_html_link_element_get_rev(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetRev is a wrapper around webkit_dom_html_link_element_set_rev():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-set-rev
func (e *HTMLLinkElement) SetRev(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_link_element_set_rev(e.native(), (*C.gchar)(cstr))
}

// Target is a wrapper around webkit_dom_html_link_element_get_target():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-get-target
func (e *HTMLLinkElement) Target() string {
    cstr := C.webkit_dom_html_link_element_get_target(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetTarget is a wrapper around webkit_dom_html_link_element_set_target():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-set-target
func (e *HTMLLinkElement) SetTarget(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_link_element_set_target(e.native(), (*C.gchar)(cstr))
}

// TypeAttr is a wrapper around webkit_dom_html_link_element_get_target():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-get-target
func (e *HTMLLinkElement) TypeAttr() string {
    cstr := C.webkit_dom_html_link_element_get_target(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetTypeAttr is a wrapper around webkit_dom_html_link_element_set_type_attr():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-set-type-attr
func (e *HTMLLinkElement) SetTypeAttr(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_link_element_set_type_attr(e.native(), (*C.gchar)(cstr))
}

// Sheet is a wrapper around webkit_dom_html_link_element_get_sheet():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLLinkElement.html#webkit-dom-html-link-element-get-sheet
func (e *HTMLLinkElement) Sheet() *StyleSheet {
    p := C.webkit_dom_html_link_element_get_sheet(e.native())
    return newStyleSheet(unsafe.Pointer(p))
}

/*
 * HTMLMapElement
 */

// type HTMLMapElement struct {
// 	HTMLElement
// }

/*
 * HTMLMarqueeElement
 */

// type HTMLMarqueeElement struct {
// 	HTMLElement
// }

/*
 * HTMLMenuElement
 */

// type HTMLMenuElement struct {
// 	HTMLElement
// }

/*
 * HTMLMetaElement
 */

type HTMLMetaElement struct {
	HTMLElement
}

func marshalHTMLMetaElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLMetaElement(obj), nil
}

func wrapHTMLMetaElement(obj *glib.Object) *HTMLMetaElement {
    return &HTMLMetaElement{HTMLElement{Element{Node{Objet{obj}}}}}
}

func (e *HTMLMetaElement) native() *C.WebKitDOMHTMLMetaElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLMetaElement(p)
}

func newHTMLMetaElement(p unsafe.Pointer) *HTMLMetaElement {
    return wrapHTMLMetaElement(newGlibObject(p));
}

// Content is a wrapper around webkit_dom_html_meta_element_get_content():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLMetaElement.html#webkit-dom-html-meta-element-get-content
func (e *HTMLMetaElement) Content() string {
    cstr := C.webkit_dom_html_meta_element_get_content(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetContent is a wrapper around webkit_dom_html_meta_element_set_content():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLMetaElement.html#webkit-dom-html-meta-element-set-content
func (e *HTMLMetaElement) SetContent(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_meta_element_set_content(e.native(), (*C.gchar)(cstr))
}

// HTTPEquiv is a wrapper around webkit_dom_html_meta_element_get_http_equiv():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLMetaElement.html#webkit-dom-html-meta-element-get-http-equiv
func (e *HTMLMetaElement) HTTPEquiv() string {
    cstr := C.webkit_dom_html_meta_element_get_http_equiv(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetHTTPEquiv is a wrapper around webkit_dom_html_meta_element_set_http_equiv():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLMetaElement.html#webkit-dom-html-meta-element-set-http-equiv
func (e *HTMLMetaElement) SetHTTPEquiv(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_meta_element_set_http_equiv(e.native(), (*C.gchar)(cstr))
}

// Name is a wrapper around webkit_dom_html_meta_element_get_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLMetaElement.html#webkit-dom-html-meta-element-get-name
func (e *HTMLMetaElement) Name() string {
    cstr := C.webkit_dom_html_meta_element_get_name(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetName is a wrapper around webkit_dom_html_meta_element_set_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLMetaElement.html#webkit-dom-html-meta-element-set-name
func (e *HTMLMetaElement) SetName(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_meta_element_set_name(e.native(), (*C.gchar)(cstr))
}

// Scheme is a wrapper around webkit_dom_html_meta_element_get_scheme():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLMetaElement.html#webkit-dom-html-meta-element-get-scheme
func (e *HTMLMetaElement) Scheme() string {
    cstr := C.webkit_dom_html_meta_element_get_scheme(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetScheme is a wrapper around webkit_dom_html_meta_element_set_scheme():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLMetaElement.html#webkit-dom-html-meta-element-set-scheme
func (e *HTMLMetaElement) SetScheme(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_meta_element_set_scheme(e.native(), (*C.gchar)(cstr))
}

/*
 * HTMLModElement
 */

// type HTMLModElement struct {
// 	HTMLElement
// }

/*
 * HTMLOListElement
 */

// type HTMLOListElement struct {
// 	HTMLElement
// }

/*
 * HTMLObjectElement
 */

type HTMLObjectElement struct {
	HTMLElement
}

func marshalHTMLObjectElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLObjectElement(obj), nil
}

func wrapHTMLObjectElement(obj *glib.Object) *HTMLObjectElement {
    return &HTMLObjectElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLObjectElement) native() *C.WebKitDOMHTMLObjectElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLObjectElement(p)
}

func newHTMLObjectElement(p unsafe.Pointer) *HTMLObjectElement {
    return wrapHTMLObjectElement(newGlibObject(p));
}

// Form is a wrapper around webkit_dom_html_object_element_get_form():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-get-form
func (e *HTMLObjectElement) Form() *HTMLFormElement {
    p := C.webkit_dom_html_object_element_get_form(e.native())
    return newHTMLFormElement(unsafe.Pointer(p))
}

// Code is a wrapper around webkit_dom_html_object_element_get_code():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-get-code
func (e *HTMLObjectElement) Code() string {
    cstr := C.webkit_dom_html_object_element_get_code(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetCode is a wrapper around webkit_dom_html_object_element_set_code():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-set-code
func (e *HTMLObjectElement) SetCode(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_object_element_set_code(e.native(), (*C.gchar)(cstr))
}

// Align is a wrapper around webkit_dom_html_object_element_get_align():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-get-align
func (e *HTMLObjectElement) Align() string {
    cstr := C.webkit_dom_html_object_element_get_align(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetAlign is a wrapper around webkit_dom_html_object_element_set_align():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-set-align
func (e *HTMLObjectElement) SetAlign(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_object_element_set_align(e.native(), (*C.gchar)(cstr))
}

// Archive is a wrapper around webkit_dom_html_object_element_set_archive():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-set-archive
func (e *HTMLObjectElement) Archive() string {
    cstr := C.webkit_dom_html_object_element_set_archive(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetArchive is a wrapper around webkit_dom_html_object_element_set_archive():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-set-archive
func (e *HTMLObjectElement) SetArchive(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_object_element_set_archive(e.native(), (*C.gchar)(cstr))
}

// Border is a wrapper around webkit_dom_html_object_element_get_border():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-get-border
func (e *HTMLObjectElement) Border() string {
    cstr := C.webkit_dom_html_object_element_get_border(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetBorder is a wrapper around webkit_dom_html_object_element_set_border():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-set-border
func (e *HTMLObjectElement) SetBorder(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_object_element_set_border(e.native(), (*C.gchar)(cstr))
}

// CodeBase is a wrapper around webkit_dom_html_object_element_get_code_base():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-get-code-base
func (e *HTMLObjectElement) CodeBase() string {
    cstr := C.webkit_dom_html_object_element_get_code_base(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetCodeBase is a wrapper around webkit_dom_html_object_element_set_code_base():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-set-code-base
func (e *HTMLObjectElement) SetCodeBase(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_object_element_set_code_base(e.native(), (*C.gchar)(cstr))
}

// CodeType is a wrapper around webkit_dom_html_object_element_get_code_type():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-get-code-type
func (e *HTMLObjectElement) CodeType() string {
    cstr := C.webkit_dom_html_object_element_get_code_type(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetCodeType is a wrapper around webkit_dom_html_object_element_set_code_type():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-set-code-type
func (e *HTMLObjectElement) SetCodeType(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_object_element_set_code_type(e.native(), (*C.gchar)(cstr))
}

// Data is a wrapper around webkit_dom_html_object_element_get_data():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-get-data
func (e *HTMLObjectElement) Data() string {
    cstr := C.webkit_dom_html_object_element_get_data(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetData is a wrapper around webkit_dom_html_object_element_set_data():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-set-data
func (e *HTMLObjectElement) SetData(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_object_element_set_data(e.native(), (*C.gchar)(cstr))
}

// Declare is a wrapper around webkit_dom_html_object_element_get_declare():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-get-declare
func (e *HTMLObjectElement) Declare() bool {
    return gobool(C.webkit_dom_html_object_element_get_declare(e.native()))
}

// SetDeclare is a wrapper around webkit_dom_html_object_element_set_declare():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-set-declare
func (e *HTMLObjectElement) SetDeclare(value bool) {
    C.webkit_dom_html_object_element_set_declare(e.native(), gboolean(value))
}

// Height is a wrapper around webkit_dom_html_object_element_get_height():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-get-height
func (e *HTMLObjectElement) Height() string {
    cstr := C.webkit_dom_html_object_element_get_height(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetHeight is a wrapper around webkit_dom_html_object_element_set_height():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-set-height
func (e *HTMLObjectElement) SetHeight(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_object_element_set_height(e.native(), (*C.gchar)(cstr))
}

// Width is a wrapper around webkit_dom_html_object_element_get_width():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-get-width
func (e *HTMLObjectElement) Width() string {
    cstr := C.webkit_dom_html_object_element_get_width(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetWidth is a wrapper around webkit_dom_html_object_element_set_width():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-set-width
func (e *HTMLObjectElement) SetWidth(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_object_element_set_width(e.native(), (*C.gchar)(cstr))
}

// Hspace is a wrapper around webkit_dom_html_object_element_get_hspace():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-get-hspace
func (e *HTMLObjectElement) Hspace() int {
    return int(C.webkit_dom_html_object_element_get_hspace(e.native()))
}

// SetHspace is a wrapper around webkit_dom_html_object_element_set_hspace():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-set-hspace
func (e *HTMLObjectElement) SetHspace(value int) {
    C.webkit_dom_html_object_element_set_hspace(e.native(), C.glong(value))
}

// Vspace is a wrapper around webkit_dom_html_object_element_get_vspace():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-get-vspace
func (e *HTMLObjectElement) Vspace() int {
    return int(C.webkit_dom_html_object_element_get_vspace(e.native()))
}

// SetVspace is a wrapper around webkit_dom_html_object_element_set_vspace():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-set-vspace
func (e *HTMLObjectElement) SetVspace(value int) {
    C.webkit_dom_html_object_element_set_vspace(e.native(), C.glong(value))
}

// UseMap is a wrapper around webkit_dom_html_object_element_get_use_map():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-get-use-map
func (e *HTMLObjectElement) UseMap() string {
    cstr := C.webkit_dom_html_object_element_get_use_map(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetUseMap is a wrapper around webkit_dom_html_object_element_set_use_map():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-set-use-map
func (e *HTMLObjectElement) SetUseMap(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_object_element_set_use_map(e.native(), (*C.gchar)(cstr))
}

// Name is a wrapper around webkit_dom_html_object_element_get_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-get-name
func (e *HTMLObjectElement) Name() string {
    cstr := C.webkit_dom_html_object_element_get_name(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetName is a wrapper around webkit_dom_html_object_element_set_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-set-name
func (e *HTMLObjectElement) SetName(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_object_element_set_name(e.native(), (*C.gchar)(cstr))
}

// Standby is a wrapper around webkit_dom_html_object_element_get_standby():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-get-standby
func (e *HTMLObjectElement) Standby() string {
    cstr := C.webkit_dom_html_object_element_get_standby(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetStandby is a wrapper around webkit_dom_html_object_element_set_standby():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-set-standby
func (e *HTMLObjectElement) SetStandby(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_object_element_set_standby(e.native(), (*C.gchar)(cstr))
}

// TypeAttr is a wrapper around webkit_dom_html_object_element_get_type_attr():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-get-type-attr
func (e *HTMLObjectElement) TypeAttr() string {
   cstr := C.webkit_dom_html_object_element_get_type_attr(e.native())
   return C.GoString((*C.char)(cstr))
}

// SetTypeAttr is a wrapper around webkit_dom_html_object_element_set_type_attr():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-set-type-attr
func (e *HTMLObjectElement) SetTypeAttr(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_object_element_set_type_attr(e.native(), (*C.gchar)(cstr))
}

// ContentDocument is a wrapper around webkit_dom_html_object_element_get_content_document():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLObjectElement.html#webkit-dom-html-object-element-get-content-document
func (e *HTMLObjectElement) ContentDocument() *Document {
    p := C.webkit_dom_html_object_element_get_content_document(e.native())
    return newDocument(unsafe.Pointer(p))
}


/*
 * HTMLOptGroupElement
 */

type HTMLOptGroupElement struct {
	HTMLElement
}

func marshalHTMLOptGroupElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLOptGroupElement(obj), nil
}

func wrapHTMLOptGroupElement(obj *glib.Object) *HTMLOptGroupElement {
    return &HTMLOptGroupElement{obj}
}

func (e *HTMLOptGroupElement) native() *C.WebKitDOMHTMLOptGroupElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLOptGroupElement(p)
}

func newHTMLOptGroupElement(p unsafe.Pointer) *HTMLOptGroupElement {
    return wrapHTMLOptGroupElement(newGlibObject(p));
}

// Disabled is a wrapper around webkit_dom_html_opt_group_element_get_disabled():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptGroupElement.html#webkit-dom-html-opt-group-element-get-disabled
func (e *HTMLOptGroupElement) Disabled() bool {
    return gobool(C.webkit_dom_html_opt_group_element_get_disabled(e.native()))
}

// SetDisabled is a wrapper around webkit_dom_html_opt_group_element_set_disabled():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptGroupElement.html#webkit-dom-html-opt-group-element-set-disabled
func (e *HTMLOptGroupElement) SetDisabled(value bool) {
    C.webkit_dom_html_opt_group_element_set_disabled(e.native(), gboolean(value))
}

// Label is a wrapper around webkit_dom_html_opt_group_element_get_label():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptGroupElement.html#webkit-dom-html-opt-group-element-get-label
func (e *HTMLOptGroupElement) Label() string {
    cstr := C.webkit_dom_html_opt_group_element_get_label(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetLabel is a wrapper around webkit_dom_html_opt_group_element_set_label():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptGroupElement.html#webkit-dom-html-opt-group-element-set-label
func (e *HTMLOptGroupElement) SetLabel(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_opt_group_element_set_label(e.native(), (*C.gchar)(cstr))
}

/*
 * HTMLOptionElement
 */

type HTMLOptionElement struct {
	HTMLElement
}

func marshalHTMLOptionElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLOptionElement(obj), nil
}

func wrapHTMLOptionElement(obj *glib.Object) *HTMLOptionElement {
    return &HTMLOptionElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLOptionElement) native() *C.WebKitDOMHTMLOptionElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLOptionElement(p)
}

func newHTMLOptionElement(p unsafe.Pointer) *HTMLOptionElement {
    return wrapHTMLOptionElement(newGlibObject(p));
}

// Disabled is a wrapper around webkit_dom_html_option_element_get_disabled():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptionElement.html#webkit-dom-html-option-element-get-disabled
func (e *HTMLOptionElement) Disabled() bool {
    return gobool(C.webkit_dom_html_option_element_get_disabled(e.native()))
}

// SetDisabled is a wrapper around webkit_dom_html_option_element_set_disabled():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptionElement.html#webkit-dom-html-option-element-set-disabled
func (e *HTMLOptionElement) SetDisabled(value bool) {
    C.webkit_dom_html_option_element_set_disabled(e.native(), gboolean(value))
}

// Form is a wrapper around webkit_dom_html_option_element_get_form():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptionElement.html#webkit-dom-html-option-element-get-form
func (e *HTMLOptionElement) Form() *HTMLFormElement {
    p := C.webkit_dom_html_option_element_get_form(e.native())
    return newHTMLFormElement(unsafe.Pointer(p))
}

// Label is a wrapper around webkit_dom_html_option_element_get_label():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptionElement.html#webkit-dom-html-option-element-get-label
func (e *HTMLOptionElement) Label() string {
    cstr := C.webkit_dom_html_option_element_get_label(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetLabel is a wrapper around webkit_dom_html_option_element_set_label():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptionElement.html#webkit-dom-html-option-element-set-label
func (e *HTMLOptionElement) SetLabel(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_option_element_set_label(e.native(), (*C.gchar)(cstr))
}

// DefaultSelected is a wrapper around webkit_dom_html_option_element_get_default_selected():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptionElement.html#webkit-dom-html-option-element-get-default-selected
func (e *HTMLOptionElement) DefaultSelected() bool {
    return gobool(C.webkit_dom_html_option_element_get_default_selected(e.native()))
}

// SetDefaultSelected is a wrapper around webkit_dom_html_option_element_set_default_selected():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptionElement.html#webkit-dom-html-option-element-set-default-selected
func (e *HTMLOptionElement) SetDefaultSelected(value bool) {
    C.webkit_dom_html_option_element_set_default_selected(e.native(), gboolean(value))
}

// Selected is a wrapper around webkit_dom_html_option_element_get_selected():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptionElement.html#webkit-dom-html-option-element-get-selected
func (e *HTMLOptionElement) Selected() bool {
    return gobool(C.webkit_dom_html_option_element_get_selected(e.native()))
}

// SetSelected is a wrapper around webkit_dom_html_option_element_set_selected():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptionElement.html#webkit-dom-html-option-element-set-selected
func (e *HTMLOptionElement) SetSelected(value bool) {
    C.webkit_dom_html_option_element_set_selected(e.native(), gboolean(value))
}

// Value is a wrapper around webkit_dom_html_option_element_get_value():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptionElement.html#webkit-dom-html-option-element-get-value
func (e *HTMLOptionElement) Value() string {
    cstr := C.webkit_dom_html_option_element_get_value(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetValue is a wrapper around webkit_dom_html_option_element_set_value():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptionElement.html#webkit-dom-html-option-element-set-value
func (e *HTMLOptionElement) SetValue(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_option_element_set_value(e.native(), (*C.gchar)(cstr))
}

// Text is a wrapper around webkit_dom_html_option_element_get_text():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptionElement.html#webkit-dom-html-option-element-get-text
func (e *HTMLOptionElement) Text() string {
    cstr := C.webkit_dom_html_option_element_get_text(e.native())
    return C.GoString((*C.char)(cstr))
}

// Index is a wrapper around webkit_dom_html_option_element_get_index():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptionElement.html#webkit-dom-html-option-element-get-index
func (e *HTMLOptionElement) Index() int {
    return int(C.webkit_dom_html_option_element_get_index(e.native()))
}

/*
 * HTMLOptionsCollection
 */

type HTMLOptionsCollection struct {
	HTMLCollection
}

func marshalHTMLOptionsCollection(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLOptionsCollection(obj), nil
}

func wrapHTMLOptionsCollection(obj *glib.Object) *HTMLOptionsCollection {
    return &HTMLOptionsCollection{HTMLCollection{Object{obj}}}
}

func (c *HTMLOptionsCollection) native() *C.WebKitDOMHTMLOptionsCollection {
    if c == nil || c.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(c.GObject)
    return C.toWebKitDOMHTMLOptionsCollection(p)
}

func newHTMLOptionsCollection(p unsafe.Pointer) *HTMLOptionsCollection {
    return wrapHTMLOptionsCollection(newGlibObject(p));
}

// NamedItem is a wrapper around webkit_dom_html_options_collection_named_item():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptionsCollection.html#webkit-dom-html-options-collection-named-item
func (c *HTMLOptionsCollection) NamedItem(name string) *Node {
    cstr := C.CString(name)
    defer C.free(unsafe.Pointer(cstr))
    p := C.webkit_dom_html_options_collection_named_item(c.native(), (*C.gchar)(cstr))
    return newNode(unsafe.Pointer(p))
}

// SelectedIndex is a wrapper around webkit_dom_html_options_collection_get_selected_index():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptionsCollection.html#webkit-dom-html-options-collection-get-selected-index
func (c *HTMLOptionsCollection) SelectedIndex() int {
    return int(webkit_dom_html_options_collection_get_selected_index(c.native()))
}

// SetSelectedIndex is a wrapper around webkit_dom_html_options_collection_set_selected_index():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptionsCollection.html#webkit-dom-html-options-collection-set-selected-index
func (c *HTMLOptionsCollection) SetSelectedIndex(value int) {
    C.webkit_dom_html_options_collection_set_selected_index(c.native(), C.glong(value))
}

// Len is a wrapper around webkit_dom_html_options_collection_get_length():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLOptionsCollection.html#webkit-dom-html-options-collection-get-length
func (c *HTMLOptionsCollection) Len() uint {
    return uint(C.webkit_dom_html_options_collection_get_length(c.native()))
}

/*
 * HTMLParagraphElement
 */

// type HTMLParagraphElement struct {
// 	HTMLElement
// }

/*
 * HTMLParamElement
 */

// type HTMLParamElement struct {
// 	HTMLElement
// }

/*
 * HTMLPreElement
 */

// type HTMLPreElement struct {
// 	HTMLElement
// }

/*
 * HTMLQuoteElement
 */

// type HTMLQuoteElement struct {
// 	HTMLElement
// }

/*
 * HTMLScriptElement
 */

type HTMLScriptElement struct {
	HTMLElement
}

func marshalHTMLScriptElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLScriptElement(obj), nil
}

func wrapHTMLScriptElement(obj *glib.Object) *HTMLScriptElement {
    return &HTMLScriptElement{HTMLElement{Element{Node{Object{obj}}}}}
}

func (e *HTMLScriptElement) native() *C.WebKitDOMHTMLScriptElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLScriptElement(p)
}

func newHTMLScriptElement(p unsafe.Pointer) *HTMLScriptElement {
    return wrapHTMLScriptElement(newGlibObject(p));
}

// Text is a wrapper around webkit_dom_html_script_element_get_text():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLScriptElement.html#webkit-dom-html-script-element-get-text
func (e *HTMLScriptElement) Text() string {
    cstr := C.webkit_dom_html_script_element_get_text(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetText is a wrapper around webkit_dom_html_script_element_set_text():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLScriptElement.html#webkit-dom-html-script-element-set-text
func (e *HTMLScriptElement) SetText(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_script_element_set_text(e.native(), (*C.gchar)(cstr))
}

// HTMLFor is a wrapper around webkit_dom_html_script_element_get_html_for():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLScriptElement.html#webkit-dom-html-script-element-get-html-for
func (e *HTMLScriptElement) HTMLFor() string {
    cstr := C.webkit_dom_html_script_element_get_html_for(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetHTMLFor is a wrapper around webkit_dom_html_script_element_set_html_for():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLScriptElement.html#webkit-dom-html-script-element-set-html-for
func (e *HTMLScriptElement) SetHTMLFor(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_script_element_set_html_for(e.native(), (*C.gchar)(cstr))
}

// Event is a wrapper around webkit_dom_html_script_element_get_event():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLScriptElement.html#webkit-dom-html-script-element-get-event
func (e *HTMLScriptElement) Event() string {
    cstr := C.webkit_dom_html_script_element_get_event(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetEvent is a wrapper around webkit_dom_html_script_element_set_event():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLScriptElement.html#webkit-dom-html-script-element-set-event
func (e *HTMLScriptElement) SetEvent(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_script_element_set_event(e.native(), (*C.gchar)(cstr))
}

// Charset is a wrapper around webkit_dom_html_script_element_get_charset():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLScriptElement.html#webkit-dom-html-script-element-get-charset
func (e *HTMLScriptElement) Charset() string {
    cstr := C.webkit_dom_html_script_element_get_charset(e.native())
    return C.GoString((*C.char)(cstr))
}

// Defer is a wrapper around webkit_dom_html_script_element_get_defer():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLScriptElement.html#webkit-dom-html-script-element-get-defer
func (e *HTMLScriptElement) Defer() bool {
    return gobool(C.webkit_dom_html_script_element_get_defer(e.native()))
}

// SetDefer is a wrapper around webkit_dom_html_script_element_set_defer():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLScriptElement.html#webkit-dom-html-script-element-set-defer
func (e *HTMLScriptElement) SetDefer(value bool) {
    C.webkit_dom_html_script_element_set_defer(e.native(), gboolean(value))
}

// Src is a wrapper around webkit_dom_html_script_element_get_src():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLScriptElement.html#webkit-dom-html-script-element-get-src
func (e *HTMLScriptElement) Src() string {
    cstr := C.webkit_dom_html_script_element_get_src(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetSrc is a wrapper around webkit_dom_html_script_element_set_src():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLScriptElement.html#webkit-dom-html-script-element-set-src
func (e *HTMLScriptElement) SetSrc(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_script_element_set_src(e.native(), (*C.gchar)(cstr))
}

// TypeAttr is a wrapper around webkit_dom_html_script_element_get_type_attr():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLScriptElement.html#webkit-dom-html-script-element-get-type-attr
func (e *HTMLScriptElement) TypeAttr() string {
    cstr := C.webkit_dom_html_script_element_get_type_attr(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetTypeAttr is a wrapper around webkit_dom_html_script_element_set_type_attr():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLScriptElement.html#webkit-dom-html-script-element-set-type-attr
func (e *HTMLScriptElement) SetTypeAttr(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_script_element_set_type_attr(e.native(), (*C.gchar)(cstr))
}

/*
 * HTMLSelectElement
 */

type HTMLSelectElement struct {
	HTMLElement
}

func marshalHTMLSelectElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLSelectElement(obj), nil
}

func wrapHTMLSelectElement(obj *glib.Object) *HTMLSelectElement {
    return &HTMLSelectElement{obj}
}

func (e *HTMLSelectElement) native() *C.WebKitDOMHTMLSelectElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLSelectElement(p)
}

func newHTMLSelectElement(p unsafe.Pointer) *HTMLSelectElement {
    return wrapHTMLSelectElement(newGlibObject(p));
}

// Item is a wrapper around webkit_dom_html_select_element_item():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-item
func (e *HTMLSelectElement) Item(index uint) *Node {
    p := C.webkit_dom_html_select_element_item(e.native(), C.gulong(index))
    return newNode(unsafe.Pointer(p))
}

// NamedItem is a wrapper around webkit_dom_html_select_element_named_item():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-named-item
func (e *HTMLSelectElement) NamedItem(name string) *Node {
    cstr := C.CString(name)
    defer C.free(unsafe.Pointer(cstr))
    p := C.webkit_dom_html_select_element_named_item(e.native(), (*C.gchar)(cstr))
    return newNode(unsafe.Pointer(p))
}

// Add is a wrapper around webkit_dom_html_select_element_add():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-add
func (e *HTMLSelectElement) Add(element, before *HTMLElement) error {
    var err *C.GError
    C.webkit_dom_html_select_element_add(e.native(), element.native(), before.native(), &err)
    if err != nil {
        defer C.g_error_free(err)
        return errors.New(C.GoString((*C.char)(err.message)))
    }
    return nil
}

// Remove is a wrapper around webkit_dom_html_select_element_remove():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-remove
func (e *HTMLSelectElement) Remove(index int) {
    C.webkit_dom_html_select_element_remove(e.native(), C.glong(index))
}

// Autofocus is a wrapper around webkit_dom_html_select_element_get_autofocus():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-get-autofocus
func (e *HTMLSelectElement) Autofocus() bool {
    return gobool(C.webkit_dom_html_select_element_get_autofocus(e.native()))
}

// SetAutofocus is a wrapper around webkit_dom_html_select_element_set_autofocus():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-set-autofocus
func (e *HTMLSelectElement) SetAutofocus(value bool) {
    C.webkit_dom_html_select_element_set_autofocus(e.native(), gboolean(value))
}

// Disabled is a wrapper around webkit_dom_html_select_element_get_disabled():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-get-disabled
func (e *HTMLSelectElement) Disabled() bool {
    return gobool(C.webkit_dom_html_select_element_get_disabled(e.native()))
}

// SetDisabled is a wrapper around webkit_dom_html_select_element_set_disabled():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-set-disabled
func (e *HTMLSelectElement) SetDisabled(value bool) {
    C.webkit_dom_html_select_element_set_disabled(e.native(), gboolean(value))
}

// Form is a wrapper around webkit_dom_html_select_element_get_form():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-get-form
func (e *HTMLSelectElement) Form() *HTMLFormElement {
    p := C.webkit_dom_html_select_element_get_form(e.native())
    return newHTMLFormElement(unsafe.Pointer(p))
}

// Multiple is a wrapper around webkit_dom_html_select_element_get_multiple():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-get-multiple
func (e *HTMLSelectElement) Multiple() bool {
    return gobool(C.webkit_dom_html_select_element_get_multiple(e.native()))
}

// SetMultiple is a wrapper around webkit_dom_html_select_element_set_multiple():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-set-multiple
func (e *HTMLSelectElement) SetMultiple(value bool) {
    C.webkit_dom_html_select_element_set_multiple(e.native(), gboolean(value))
}

// Name is a wrapper around webkit_dom_html_select_element_get_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-get-name
func (e *HTMLSelectElement) Name() string {
    cstr := C.webkit_dom_html_select_element_get_name(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetName is a wrapper around webkit_dom_html_select_element_set_name():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-set-name
func (e *HTMLSelectElement) SetName(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_select_element_set_name(e.native(), (*C.gchar)(cstr))
}

// Size is a wrapper around webkit_dom_html_select_element_get_size():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-get-size
func (e *HTMLSelectElement) Size() int {
    return int(C.webkit_dom_html_select_element_get_size(e.native()))
}

// SetSize is a wrapper around webkit_dom_html_select_element_set_size():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-set-size
func (e *HTMLSelectElement) SetSize(value int) {
    C.webkit_dom_html_select_element_set_size(e.native(), C.glong(value))
}

// SelectType is a wrapper around webkit_dom_html_select_element_get_select_type():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-get-select-type
func (e *HTMLSelectElement) SelectType() string {
    cstr := C.webkit_dom_html_select_element_get_select_type(e.native())
    return C.GoString((*C.char)(cstr))
}

// Options is a wrapper around webkit_dom_html_select_element_get_options():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-get-options
func (e *HTMLSelectElement) Options() *HTMLOptionsCollection {
    p := C.webkit_dom_html_select_element_get_options(e.native())
    return newHTMLOptionsCollection(unsafe.Pointer(p))
}

// Len is a wrapper around webkit_dom_html_select_element_get_length():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-get-length
func (e *HTMLSelectElement) Len() uint {
    return uint(C.webkit_dom_html_select_element_get_length(e.native()))
}

// SetLen is a wrapper around webkit_dom_html_select_element_set_length():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-set-length
func (e *HTMLSelectElement) SetLen(value uint) {
    C.webkit_dom_html_select_element_set_length(e.native(), C.gulong(value))
}

// SelectedIndex is a wrapper around webkit_dom_html_select_element_get_selected_index():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-get-selected-index
func (e *HTMLSelectElement) SelectedIndex() int {
    return int(C.webkit_dom_html_select_element_get_selected_index(e.native()))
}

// SetSelectedIndex is a wrapper around webkit_dom_html_select_element_set_selected_index():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-set-selected-index
func (e *HTMLSelectElement) SetSelectedIndex(value int) {
    C.webkit_dom_html_select_element_set_selected_index(e.native(), C.glong(value))
}

// Value is a wrapper around webkit_dom_html_select_element_get_value():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-get-value
func (e *HTMLSelectElement) Value() string {
    cstr := C.webkit_dom_html_select_element_get_value(e.native())
    return C.GoString((*C.char)(cstr))
}

// SetValue is a wrapper around webkit_dom_html_select_element_set_value():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-set-value
func (e *HTMLSelectElement) SetValue(value string) {
    cstr := C.CString(value)
    defer C.free(unsafe.Pointer(cstr))
    C.webkit_dom_html_select_element_set_value(e.native(), (*C.gchar)(cstr))
}

// WillValidate is a wrapper around webkit_dom_html_select_element_get_will_validate():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLSelectElement.html#webkit-dom-html-select-element-get-will-validate
func (e *HTMLSelectElement) WillValidate() bool {
    return gobool(C.webkit_dom_html_select_element_get_will_validate(e.native()))
}

/*
 * HTMLStyleElement
 */

// type HTMLStyleElement struct {
// 	HTMLElement
// }

/*
 * HTMLTableCaptionElement
 */

// type HTMLTableCaptionElement struct {
// 	HTMLElement
// }

/*
 * HTMLTableCellElement
 */

// type HTMLTableCellElement struct {
// 	HTMLElement
// }

/*
 * HTMLTableColElement
 */

// type HTMLTableColElement struct {
// 	HTMLElement
// }

/*
 * HTMLTableElement
 */

// type HTMLTableElement struct {
// 	HTMLElement
// }

/*
 * HTMLTableRowElement
 */

// type HTMLTableRowElement struct {
// 	HTMLElement
// }

/*
 * HTMLTableSectionElement
 */

// type HTMLTableSectionElement struct {
// 	HTMLElement
// }

/*
 * HTMLTextAreaElement
 */

type HTMLTextAreaElement struct {
	HTMLElement
}

func marshalHTMLTextAreaElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLTextAreaElement(obj), nil
}

func wrapHTMLTextAreaElement(obj *glib.Object) *HTMLTextAreaElement {
    return &HTMLTextAreaElement{obj}
}

func (e *HTMLTextAreaElement) native() *C.WebKitDOMHTMLTextAreaElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLTextAreaElement(p)
}

func newHTMLTextAreaElement(p unsafe.Pointer) *HTMLTextAreaElement {
    return wrapHTMLTextAreaElement(newGlibObject(p));
}

/*
 * HTMLTitleElement
 */

type HTMLTitleElement struct {
	HTMLElement
}

func marshalHTMLTitleElement(p uintptr) (interface{}, error) {
    c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
    obj := &glib.Object{GObject: glib.ToGObject(unsafe.Pointer(c))}
    return wrapHTMLTitleElement(obj), nil
}

func wrapHTMLTitleElement(obj *glib.Object) *HTMLTitleElement {
    return &HTMLTitleElement{HTMLElement{Element{Node{object{obj}}}}}
}

func (e *HTMLTitleElement) native() *C.WebKitDOMHTMLTitleElement {
    if e == nil || e.GObject == nil {
        return nil
    }
    p := unsafe.Pointer(e.GObject)
    return C.toWebKitDOMHTMLTitleElement(p)
}

func newHTMLTitleElement(p unsafe.Pointer) *HTMLTitleElement {
    return wrapHTMLTitleElement(newGlibObject(p));
}

// Text is a wrapper around webkit_dom_html_title_element_get_text():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLTitleElement.html#webkit-dom-html-title-element-get-text
func (e *HTMLTitleElement) Text() string {
    t := C.webkit_dom_html_title_element_get_text(e.native())
    return C.GoString((*C.char)(t))
}

// SetText is a wrapper around webkit_dom_html_title_element_set_text():
// http://webkitgtk.org/reference/webkitdomgtk/stable/WebKitDOMHTMLTitleElement.html#webkit-dom-html-title-element-set-text
func (e *HTMLTitleElement) SetText(value string) {
    v := C.CString(value)
    defer C.free(unsafe.Pointer(v))
    C.webkit_dom_html_title_element_set_text(e.native, (*C.gchar)(v))
}

/*
 * HTMLUListElement
 */

// type HTMLUListElement struct {
// 	HTMLElement
// }
