
document.querySelector("#farmer_reg").onclick = (e) => {
    e.preventDefault()


    document.querySelector("#username").classList.remove("border-danger")
    document.querySelector("#usernameError").style.display = "none"
    document.querySelector("#email").classList.remove("border-danger")
    document.querySelector("#emailError").style.display = "none"
    document.querySelector("#confirm-password").classList.remove("border-danger")
    document.querySelector("#cpasswordError").style.display = "none"
    document.querySelector("#passwor").classList.remove("border-danger")
    document.querySelector("#passwordErro").style.display = "none"


    axios.post("/f-signup", {
        username: document.querySelector("#username").value,
        email: document.querySelector("#email").value,
        password: document.querySelector("#passwor").value,
        cpassword: document.querySelector("#confirm-password").value
    })
        .then(data => {
            window.location.href = data.data.Feedback
            alert("Successful account registration.To continue using iReferral log in")
        })
        .catch(err => {
            var resp = err.response.data

            if (resp.Username == "exist") {
                alert("Username already used")
                document.querySelector("#username").classList.add("border-danger")
                document.querySelector("#usernameError").style.display = "block"
            }

            if (resp.Fpassword == "incorrect") {
                alert("At least 8 characters and 1 digit")
                document.querySelector("#passwor").classList.add("border-danger")
                document.querySelector("#confirm-password").classList.add("border-danger")
                document.querySelector("#cpasswordError").style.display = "block"
            }

            if (resp.Fpassword == "nomatch" && resp.Cpassword == "nomatch") {
                alert("password and confirm password does not match")
                document.querySelector("#confirm-password").classList.add("border-danger")
                document.querySelector("#cpasswordError2").style.display = "block"
                document.querySelector("#passwor").classList.add("border-danger")

            }

            if (resp.Email == "invalid") {
                alert("Invalid email address")
                document.querySelector("#email").classList.add("border-danger")
                document.querySelector("#emailError").style.display = "block"
            }

            if (resp.Email == "exist") {
                alert("Use a different email-already taken!!")
                document.querySelector("#email").classList.add("border-danger")
                document.querySelector("#emailError2").style.display = "block"
            }


        })

}
