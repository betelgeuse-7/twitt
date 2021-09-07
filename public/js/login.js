const loginFormForm = document.querySelector("#login-form")
const emailInput = document.querySelector("#email")
const passwordInput = document.querySelector("#password")
const loginButton = document.querySelector("#login")

const LOGIN_SUCCESS = "logged in"
const LOGIN_FAIL = "cant log in"

function login() {
    const body = {
        email: emailInput.value,
        password: passwordInput.value
    }
    if (checkLoginBody(body)) {
        return LOGIN_SUCCESS
    }
    return LOGIN_FAIL
}

function checkLoginBody(body) {
    if (body.email && body.password) {
        if (body.email.includes("@") && body.password.length > 4) {
            return true
        }
    }
    return false
}

loginFormForm.addEventListener("submit", (e) => {
    e.preventDefault()
    if (login() === LOGIN_SUCCESS) {
        console.log("logged in")
    } else {
        console.log("cant log in" );
    }
})