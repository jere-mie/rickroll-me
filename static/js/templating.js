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
let findTemplate=e=>document.querySelector(`template[data-templateName="${e}"`),findContainer=e=>document.querySelector(`[data-containerName="${e}"`),instantiateTemplate=(e,t)=>{let a=document.importNode(e.content,!0);for(let[n,i]of Object.entries(t)){let o=a.querySelector(`[data-fieldName="${n}"]`);if(!o){console.warn(`template does not contain a field named '${n}'`);continue}for(let[l,r]of i)o[l]=r}return a};export{findTemplate,findContainer,instantiateTemplate};
