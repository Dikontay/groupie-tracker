
//data in json
let availableKeywords = ["apple", "aaaaaple", "apppple"]

const resultsBox = document.querySelector(".search-results")
const inputBox = document.querySelector("#search-input")

inputBox.onkeyup = function(){
    let result = []
    let input = inputBox.value
    if (input.length){
        result = availableKeywords.filter((word)=>{
           return word.toLowerCase().includes(input.toLowerCase())
        });
    }
    //document.write(result[0])
    console.log(result)
}