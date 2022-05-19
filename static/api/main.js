let cities = document.getElementsByClassName("city")
let dates = document.getElementsByClassName("date")


function makeCapitalize(text) {
    text = text.charAt(0).toUpperCase() + text.slice(1)
    for(let i = 1; i < text.length; i++) {
        if (text[i] == "-" || text[i] == "_") {
            text = text.slice(0, i+1) + text.charAt(i+1).toUpperCase() + text.slice(i+2)
        }
    }
    return text
}


function deleteBrackets(text) {
    text = text.slice(1)
    text = text.slice(0, text.length-1)
    return text
}

for (var i = 0; i < cities.length; i++) {
    let content = cities[i].textContent
    let result = makeCapitalize(content)
    cities[i].textContent = result
}

for (var i = 0; i < dates.length; i++) {
    let result = deleteBrackets(dates[i].textContent)
    dates[i].textContent = result 
}

