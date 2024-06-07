/**
 * A relatively small (I won't claim tiny, because it isn't) and simple templating system. It works very
 * nicely for small websites that do not want to set up a build system or JavaScript framework such as
 * Svelte or React, but desire some of the templating capabilities of such tools.
 *
 * This is a stand-alone source file with no dependencies. By default, this file is a JS module. Just a few lines
 * below this blurb is the export list. Simply remove it if you wish to use this file as a regular old JS script
 * (not recommended, but you do you).
 *
 * @author Isaac Kilbourne
 * @license 0BSD
 * https://opensource.org/license/0bsd
 */

// If you do not want this to be a module, just remove this line and use the script as normal.
export { findTemplate, findContainer, instantiateTemplate };

/**
 * Maps a field name to an array of attribute names and values.
 *
 * For example, `{"link": [["href", "https://google.com"]]}` would locate an element with
 * attribute `data-fieldName="link"` and set the `href` attribute on that element to `"https://google.com"`.
 *
 * To set the contents of an element, use `innerHTML` as the attribute name.
 * @typedef {{
 *  [keyof: string]: [string, string][]
 * }} TemplateFillParams
 */

/**
 * Finds a <template> element with the data attribute `data-templateName` containing the
 * provided template name.
 *
 * For example, `findTemplate("foo")` would return the first DOM element that matches
 * `<template data-templateName="foo">`.
 * @param {string} templateName the name of the template to find
 * @returns {?HTMLTemplateElement} the first <template> DOM element with matching template name, if it exists
 */
const findTemplate = (templateName) => {
    return document.querySelector(`template[data-templateName="${templateName}"`);
};

/**
 * Finds any element with the data attribute `data-containerName` containing the provided
 * container name. The primary use of this function is to find a container for one or more
 * template instantiations.
 *
 * For example, `findContainer("foo")` would return the first DOM element that matches
 * `<template data-containerName="foo">`.
 * @param {string} containerName the name of the container to find
 * @returns {?HTMLElement} the first DOM element with matching template name, if it exists.
 */
const findContainer = (containerName) => {
    return document.querySelector(`[data-containerName="${containerName}"`);
};

/**
 * The bread and butter of this templating system. Clones the provided <template> element's content,
 * then replaces the contained elements with a `data-fieldName` attribute based on the `fillParams`.
 *
 * A warning is issued for each `fillParams` entry which does not match to an element with `data-fieldName`.
 * This function does not insert the result into the DOM anywhere, but instead returns the `DocumentFragment`
 * in case additional processing is required.
 *
 * @see findTemplate for a handy way to grab a template out of the DOM
 * @see findContainer for a handy way to grab a container element to place the resulting fragment in
 * @see TemplateFillParams for a more complete description of the `fillParams` parameter
 *
 * @param {HTMLTemplateElement} templateElement the template to clone
 * @param {TemplateFillParams} fillParams a mapping of fieldNames to attribute values.
 * @returns {DocumentFragment} a `DocumentFragment` which can be inserted anywhere in the DOM to display
 * the instantiated template on the page.
 */
const instantiateTemplate = (templateElement, fillParams) => {
    /** @type {DocumentFragment} */
    const clone = document.importNode(templateElement.content, true);

    for (const [fieldName, valueFillPatterns] of Object.entries(fillParams)) {
        /** @type {?HTMLElement} */
        const field = clone.querySelector(`[data-fieldName="${fieldName}"]`);
        if (!field) {
            console.warn(`template does not contain a field named '${fieldName}'`);
            continue;
        }

        for (const [attribute, value] of valueFillPatterns) {
            field[attribute] = value;
        }
    }

    return clone;
};