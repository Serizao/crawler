package js

import (
	"strconv"
	"time"

	"github.com/go-rod/rod"
)

const GetLinks = `getLinks();
function absolutePath(href) {
    try {
        var link = document.createElement("a");
        link.href = href;
        return link.href;
    } catch (error) {}
}
function getLinks() {
    var array = [];
    if (!document) return array;
    var allElements = document.querySelectorAll("*");
    for (var el of allElements) {
        if (el.href && typeof el.href === 'string') {
            array.push(el.href);
        } else if (el.src && typeof el.src === 'string') {
            var absolute = absolutePath(el.src);
            array.push(absolute);
        } else if (el.action && typeof el.action === 'string') {
            var absolute = absolutePath(el.src);
            array.push(absolute);
        } else if (el.type && typeof el.type === 'string' && el.type === 'hidden' && el.getAttribute("name") != '' && el.getAttribute("name") != null ) {
                if(el.value && el.value !=""){
                        if(window.location.href.includes('?')) array.push(window.location.href+"&"+el.getAttribute("name")+"="+el.value);
                        else array.push(window.location.href+"?"+el.getAttribute("name")+"="+el.value);
                } else {
                    if(window.location.href.includes('?')) array.push(window.location.href+"&"+el.name+"=xxxx");
                    else array.push(window.location.href+"?"+el.name+"=xxxx");
                }
        }
    }
    return array;
}`

func CreateWaitFunc(d time.Duration) *rod.EvalOptions {
	millis := d / time.Millisecond
	return &rod.EvalOptions{
		ByValue:      true,
		JS:           "new Promise(r => setTimeout(r, " + strconv.Itoa(int(millis)) + "));",
		AwaitPromise: true,
	}
}
