const username = document.getElementById("username");
const submitBtn = document.getElementById("submitBtn");

submitBtn.addEventListener("click", () => {
    sessionStorage.setItem("username", username.value)
})

username.addEventListener("input", () => {
    sessionStorage.setItem("username", username.value)
})

const u = {name: "rahul"}

sessionStorage.setItem("users", JSON.stringify(u))

const us = sessionStorage.getItem("users")

console.log(JSON.parse(us))

// URL and Fetch API 

// protocol -: als called schema
 
// protocol://subdomain.domain.TLD/path?key=value&key1=value1&...

fetch("https://jsonplaceholder.typicode.com/posts",{
    method: "GET"
})


