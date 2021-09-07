const registerFormForm = document.querySelector("#register-form")
const usernameInput = document.querySelector("#username")
const emailInput = document.querySelector("#email")
const passwordInput = document.querySelector("#password")
const submitButton = document.querySelector("#submit")

const REGISTER_SUCCESS = "register successful"
const REGISTER_FAIL = "register failed"

function register() {
    const body = {
        username: usernameInput.value,
        email: emailInput.value,
        password: passwordInput.value
    }
    console.log(body)
    if (checkRegisterBody(body)) {
        return REGISTER_SUCCESS
    } 
    return REGISTER_FAIL
}

function checkRegisterBody(body) {
    if (body.username && body.email && body.password) {
        if (body.username.length > 4 && body.email.includes("@") && body.password.length > 4) {
            return true
        }
    }
    return false
}

function redirect(endpoint) {    
    window.location.href = window.location.href.replace("index.html", `${endpoint}.html`)
}

registerFormForm.addEventListener("submit", (e) => {
    e.preventDefault()
    if (register() === REGISTER_SUCCESS) {
        redirect("login")
    }
})