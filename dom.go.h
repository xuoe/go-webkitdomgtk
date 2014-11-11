/*
* Copyright (c) 2014 Alex Diaconu.
* Use of this source code is governed by an ISC
* license that can be found in the LICENSE file.
*/

static WebKitDOMNode* toWebKitDOMNode(void *p) {
    return (WEBKIT_DOM_NODE(p));
}

static WebKitDOMObject* toWebKitDOMObject(void *p) {
    return (WEBKIT_DOM_OBJECT(p));
}

static WebKitDOMNodeList* toWebKitDOMNodeList(void *p) {
    return (WEBKIT_DOM_NODE_LIST(p));
}

static WebKitDOMNamedNodeMap* toWebKitDOMNamedNodeMap(void *p) {
    return (WEBKIT_DOM_NAMED_NODE_MAP(p));
}

static WebKitDOMNodeIterator* toWebKitDOMNodeIterator(void *p) {
    return (WEBKIT_DOM_NODE_ITERATOR(p));
}

static WebKitDOMDocument* toWebKitDOMDocument(void *p) {
    return (WEBKIT_DOM_DOCUMENT(p));
}

static WebKitDOMAttr* toWebKitDOMAttr(void *p) {
    return (WEBKIT_DOM_ATTR(p));
}

static WebKitDOMElement* toWebKitDOMElement(void *p) {
    return (WEBKIT_DOM_ELEMENT(p));
}

static WebKitDOMHTMLElement* toWebKitDOMHTMLElement(void *p) {
    return (WEBKIT_DOM_HTML_ELEMENT(p));
}

static WebKitDOMHTMLHeadElement* toWebKitDOMHTMLHeadElement(void *p) {
    return (WEBKIT_DOM_HTML_HEAD_ELEMENT(p));
}

static WebKitDOMHTMLTitleElement* toWebKitDOMHTMLTitleElement(void *p) {
    return (WEBKIT_DOM_HTML_TITLE_ELEMENT(p));
}

static WebKitDOMHTMLCollection* toWebKitDOMHTMLCollection(void *p) {
    return (WEBKIT_DOM_HTML_COLLECTION(p));
}

static WebKitDOMStyleSheet* toWebKitDOMStyleSheet(void *p) {
    return (WEBKIT_DOM_STYLE_SHEET(p));
}

static WebKitDOMStyleSheetList* toWebKitDOMStyleSheetList(void *p) {
    return (WEBKIT_DOM_STYLE_SHEET_LIST(p));
}

static WebKitDOMMediaList* toWebKitDOMMediaList(void *p) {
    return (WEBKIT_DOM_MEDIA_LIST(p));
}

static WebKitDOMCSSStyleSheet* toWebKitDOMCSSStyleSheet(void *p) {
    return (WEBKIT_DOM_CSS_STYLE_SHEET(p));
}

static WebKitDOMCSSRule* toWebKitDOMCSSRule(void *p) {
    return (WEBKIT_DOM_CSS_RULE(p));
}

static WebKitDOMCSSRuleList* toWebKitDOMCSSRuleList(void *p) {
    return (WEBKIT_DOM_CSS_RULE_LIST(p));
}

static WebKitDOMCSSStyleDeclaration* toWebKitDOMCSSStyleDeclaration(void *p) {
    return (WEBKIT_DOM_CSS_STYLE_DECLARATION(p));
}

static WebKitDOMCSSValue* toWebKitDOMCSSValue(void *p) {
    return (WEBKIT_DOM_CSS_VALUE(p));
}

static WebKitDOMCharacterData* toWebKitDOMCharacterData(void *p) {
    return (WEBKIT_DOM_CHARACTER_DATA(p));
}

static WebKitDOMText* toWebKitDOMText(void *p) {
    return (WEBKIT_DOM_TEXT(p));
}

static WebKitDOMFileList* toWebKitDOMFileList(void *p) {
    return (WEBKIT_DOM_FILE_LIST(p));
}

static WebKitDOMFile* toWebKitDOMFile(void *p) {
    return (WEBKIT_DOM_FILE(p));
}

static WebKitDOMBlob* toWebKitDOMBlob(void *p) {
    return (WEBKIT_DOM_BLOB(p));
}

static WebKitDOMProcessingInstruction* toWebKitDOMProcessingInstruction(void *p) {
    return (WEBKIT_DOM_PROCESSING_INSTRUCTION(p));
}

static WebKitDOMEvent* toWebKitDOMEvent(void *p) {
    return (WEBKIT_DOM_EVENT(p));
}

static WebKitDOMHTMLAnchorElement* toWebKitDOMHTMLAnchorElement(void *p) {
    return (WEBKIT_DOM_HTML_ANCHOR_ELEMENT(p));
}

static WebKitDOMHTMLAppletElement* toWebKitDOMHTMLAppletElement(void *p) {
    return (WEBKIT_DOM_HTML_APPLET_ELEMENT(p));
}

static WebKitDOMHTMLAreaElement* toWebKitDOMHTMLAreaElement(void *p) {
    return (WEBKIT_DOM_HTML_AREA_ELEMENT(p));
}

static WebKitDOMHTMLBRElement* toWebKitDOMHTMLBRElement(void *p) {
    return (WEBKIT_DOM_HTML_BR_ELEMENT(p));
}

static WebKitDOMHTMLBaseElement* toWebKitDOMHTMLBaseElement(void *p) {
    return (WEBKIT_DOM_HTML_BASE_ELEMENT(p));
}

static WebKitDOMHTMLTextAreaElement* toWebKitDOMHTMLTextAreaElement(void *p) {
    return (WEBKIT_DOM_HTML_TEXT_AREA_ELEMENT(p));
}

static WebKitDOMHTMLFormElement* toWebKitDOMHTMLFormElement(void *p) {
    return (WEBKIT_DOM_HTML_FORM_ELEMENT(p));
}

static WebKitDOMHTMLBodyElement* toWebKitDOMHTMLBodyElement(void *p) {
    return (WEBKIT_DOM_HTML_BODY_ELEMENT(p));
}

static WebKitDOMHTMLButtonElement* toWebKitDOMHTMLButtonElement(void *p) {
    return (WEBKIT_DOM_HTML_BUTTON_ELEMENT(p));
}

static WebKitDOMHTMLDivElement* toWebKitDOMHTMLDivElement(void *p) {
    return (WEBKIT_DOM_HTML_DIV_ELEMENT(p));
}

static WebKitDOMHTMLDocument* toWebKitDOMHTMLDocument(void *p) {
    return (WEBKIT_DOM_HTML_DOCUMENT(p));
}

static WebKitDOMHTMLEmbedElement* toWebKitDOMHTMLEmbedElement(void *p) {
    return (WEBKIT_DOM_HTML_EMBED_ELEMENT(p));
}

static WebKitDOMHTMLFieldSetElement* toWebKitDOMHTMLFieldSetElement(void *p) {
    return (WEBKIT_DOM_HTML_FIELD_SET_ELEMENT(p));
}

static WebKitDOMHTMLOptionsCollection* toWebKitDOMHTMLOptionsCollection(void *p) {
    return (WEBKIT_DOM_HTML_OPTIONS_COLLECTION(p));
}

static WebKitDOMHTMLSelectElement* toWebKitDOMHTMLSelectElement(void *p) {
    return (WEBKIT_DOM_HTML_SELECT_ELEMENT(p));
}

static WebKitDOMHTMLScriptElement* toWebKitDOMHTMLScriptElement(void *p) {
    return (WEBKIT_DOM_HTML_SCRIPT_ELEMENT(p));
}

static WebKitDOMHTMLOptionElement* toWebKitDOMHTMLOptionElement(void *p) {
    return (WEBKIT_DOM_HTML_OPTION_ELEMENT(p));
}

static WebKitDOMHTMLOptGroupElement* toWebKitDOMHTMLOptGroupElement(void *p) {
    return (WEBKIT_DOM_HTML_OPT_GROUP_ELEMENT(p));
}

static WebKitDOMHTMLMetaElement* toWebKitDOMHTMLMetaElement(void *p) {
    return (WEBKIT_DOM_HTML_META_ELEMENT(p));
}

static WebKitDOMHTMLObjectElement* toWebKitDOMHTMLObjectElement(void *p) {
    return (WEBKIT_DOM_HTML_OBJECT_ELEMENT(p));
}

static WebKitDOMXPathExpression* toWebKitDOMXPathExpression(void *p) {
    return (WEBKIT_DOM_XPATH_EXPRESSION(p));
}

static WebKitDOMXPathResult* toWebKitDOMXPathResult(void *p) {
    return (WEBKIT_DOM_XPATH_RESULT(p));
}

static WebKitDOMDOMWindow* toWebKitDOMDOMWindow(void *p) {
    return (WEBKIT_DOM_WINDOW(p));
}

static WebKitDOMHTMLFrameElement* toWebKitDOMHTMLFrameElement(void *p) {
    return (WEBKIT_DOM_HTML_FRAME_ELEMENT(p));
}

static WebKitDOMHTMLFrameSetElement* toWebKitDOMHTMLFrameSetElement(void *p) {
    return (WEBKIT_DOM_HTML_FRAME_SET_ELEMENT(p));
}

static WebKitDOMHTMLHeadElement* toWebKitDOMHTMLHeadElement(void *p) {
    return (WEBKIT_DOM_HTML_HEAD_ELEMENT(p));
}

static WebKitDOMHTMLLinkElement* toWebKitDOMHTMLLinkElement(void *p) {
    return (WEBKIT_DOM_HTML_LINK_ELEMENT(p));
}

static WebKitDOMHTMLImageElement* toWebKitDOMHTMLImageElement(void *p) {
    return (WEBKIT_DOM_HTML_IMAGE_ELEMENT(p));
}

static WebKitDOMHTMLInputElement* toWebKitDOMHTMLInputElement(void *p) {
    return (WEBKIT_DOM_HTML_INPUT_ELEMENT(p));
}

static WebKitDOMHTMLLabelElement* toWebKitDOMHTMLLabelElement(void *p) {
    return (WEBKIT_DOM_HTML_LABEL_ELEMENT(p));
}
