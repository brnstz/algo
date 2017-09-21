var NOT_EXISTS_COLOR = "#CC0000";
var EXISTS_COLOR = "#000000";
var WIKI_ROOT = "https://en.wikipedia.org/wiki/";
var SEARCHBOX = document.getElementById("searchbox");
var COMPLETIONS = document.getElementById("completions");

SEARCHBOX.addEventListener("input", inputChange);
SEARCHBOX.addEventListener("keypress", goToFirstLink);
SEARCHBOX.focus();

if (SEARCHBOX.value.length > 0) {
    callAPI(SEARCHBOX.value, SEARCHBOX, COMPLETIONS);
}

function createWikiLink(word, wiki) {
    return '<a href="' + 'https://' + wiki + '.wikipedia.org/wiki/' + encodeURIComponent(word) + '">' + word + ' (' + wiki + ')</a>';
}

function goToFirstLink(e) {
    var key = e.key | e.keyCode | e.which;
    if (key == '13') {
        window.location.href = completions.firstChild.href;
    }
    return false;
}

function callAPI(word, searchbox, completions) {
	var xhr = new XMLHttpRequest();
	var url = "/api/word?word=" + encodeURIComponent(word);

	xhr.open("GET", url);
	xhr.responseType = "json";

	xhr.onload = function() {
        var r = xhr.response;

        // Set color of the input text box
        if (r.exists) {
            searchbox.style.color = EXISTS_COLOR;
        } else {
            searchbox.style.color = NOT_EXISTS_COLOR;
        }

        completions.innerHTML = "";

        // If this word is a word, that counts as a completion
        if (r.exists) {
            for (j = 0; j < r.wikis.length; j++) {
                completions.innerHTML += createWikiLink(word, r.wikis[j]) + "<br>";
            }
        }

        // List our auto-completions
        if (r.completions != null && word.length > 0) {
            for (i = 0; i < r.completions.length; i++) {
                //completions.innerHTML += createWikiLink(r.completions[i]) + "<br>";
                completions.innerHTML += r.completions[i] + "<br>";
            }
        }
        
	};
    xhr.send();
}

function inputChange(e) {
    var word = e.srcElement.value;
    var searchbox = e.srcElement;
    var completions = document.getElementById("completions");
    callAPI(word, searchbox, completions);
}
