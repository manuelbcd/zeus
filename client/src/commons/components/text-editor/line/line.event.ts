function findClickedWord(parentElt, x, y) {
    if (parentElt.nodeName !== '#text') {
        console.log('didn\'t click on text node');
        return null;
    }
    let range = document.createRange();
    let words = parentElt.textContent.split(' ');
    let start = 0;
    let end = 0;
    for (let i = 0; i < words.length; i++) {
        let word = words[i];
        end = start+word.length;
        range.setStart(parentElt, start);
        range.setEnd(parentElt, end);
        // not getBoundingClientRect as word could wrap
        let rects = range.getClientRects();
        let clickedRect = isClickInRects(rects);
        if (clickedRect) {
            return [word, start, clickedRect];
        }
        start = end + 1;
    }

    function isClickInRects(rects) {
        for (let i = 0; i < rects.length; ++i) {
            let r = rects[i]
            if (r.left<x && r.right>x && r.top<y && r.bottom>y) {
                return r;
            }
        }
        return false;
    }
    return null;
}
function onClick(e) {
    let elt = document.getElementById('info');
    let clicked = findClickedWord(e.target.childNodes[0], e.clientX, e.clientY);
    elt.innerHTML = 'Nothing Clicked';
    if (clicked) {
        let word = clicked[0];
        let start = clicked[1];
        let r = clicked[2];
        elt.innerHTML = 'Clicked: ('+r.top+','+r.left+') word:'+word+' at offset '+start;
    }
}