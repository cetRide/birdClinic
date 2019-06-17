
document.querySelector("#reg").onclick = (e) => {
    e.preventDefault()

    document.querySelector("#username3").classList.remove("border-danger")
    document.querySelector("#usernameError3").style.display = "none"
    document.querySelector("#usernameError4").style.display = "none"
    document.querySelector("#email3").classList.remove("border-danger")
    document.querySelector("#emailError3").style.display = "none"
    document.querySelector("#emailError4").style.display = "none"
    document.querySelector("#district").classList.remove("border-danger")
    document.querySelector("#districtError").style.display = "none"
    document.querySelector("#county").classList.remove("border-danger")
    document.querySelector("#countyError").style.display = "none"
    document.querySelector("#country").classList.remove("border-danger")
    document.querySelector("#countryError").style.display = "none"
    document.querySelector("#first").classList.remove("border-danger")
    document.querySelector("#firstError").style.display = "none"
    document.querySelector("#last").classList.remove("border-danger")
    document.querySelector("#lastError").style.display = "none"
    document.querySelector("#phone").classList.remove("border-danger")
    document.querySelector("#phoneError").style.display = "none"
    document.querySelector("#phoneError2").style.display = "none"
    document.querySelector("#password3").classList.remove("border-danger")
    document.querySelector("#passwordError3").style.display = "none"
    document.querySelector("#copassword3").classList.remove("border-danger")
    document.querySelector("#copasswordError3").style.display = "none"
    document.querySelector("#error").style.display = "none"
    document.querySelector("#error1").style.display = "none"

    axios.post("/s-signup", {
        username: document.querySelector("#username3").value,
        email: document.querySelector("#email3").value,
        district: document.querySelector("#district").value,
        county: document.querySelector("#county").value,
        country: document.querySelector("#country").value,
        first: document.querySelector("#first").value,
        last: document.querySelector("#last").value,
        phone: document.querySelector("#phone").value,
        pass: document.querySelector("#password3").value,
        copassword: document.querySelector("#copassword3").value

    })
        .then(data => {
            window.location.href = data.data.ResS
            alert("Successful account registration.")
        })
        .catch(err => {
            var resp = err.response.data

            if (resp.SFirst_name == "empty") {
                // alert("Fill in empty field")
                document.querySelector("#first").classList.add("border-danger")
                document.querySelector("#firstError").style.display = "block"
            }

            if (resp.SUsername == "empty") {
                // alert("Fill in empty field")
                document.querySelector("#username3").classList.add("border-danger")
                document.querySelector("#usernameError3").style.display = "block"
            }
            if (resp.SLast_name == "empty") {
                // alert("Fill in empty field")
                document.querySelector("#last").classList.add("border-danger")
                document.querySelector("#lastError").style.display = "block"
            }
            if (resp.SDistrict == "empty") {
                // alert("Fill in empty field")
                document.querySelector("#district").classList.add("border-danger")
                document.querySelector("#districtError").style.display = "block"
            }
            if (resp.SCounty == "empty") {
                // alert("Fill in empty field")
                document.querySelector("#county").classList.add("border-danger")
                document.querySelector("#countyError").style.display = "block"
            }
            if (resp.SCountry == "empty") {
                // alert("Fill in empty field")
                document.querySelector("#country").classList.add("border-danger")
                document.querySelector("#countryError").style.display = "block"
            }

            if (resp.SEmail == "empty") {
                // alert("Fill in empty field")
                document.querySelector("#email3").classList.add("border-danger")
                document.querySelector("#emailError5").style.display = "block"
            }
           
           
            if (resp.SPassword == "incorrect") {
                // alert("Password must contain atleast 8 characters composed of both capital,small letters and digits")
                document.querySelector("#password3").classList.add("border-danger")
                // document.querySelector("#copasswordError3").style.display = "block"
                document.querySelector("#copassword3").classList.add("border-danger")
                document.querySelector("#error").style.display = "block"
            }

            if (resp.SPassword == "nomatch" && resp.ScoPassword == "nomatch") {
                // alert("password and confirm password does not match")
                document.querySelector("#password").classList.add("border-danger")
                document.querySelector("#copassword3").classList.add("border-danger")
                // document.querySelector("#passwordError3").style.display = "block"
                document.querySelector("#error1").style.display = "block"

            }         

            if (resp.SPhone == "wrong") {
                // alert("Invalid phone number")
                document.querySelector("#phone").classList.add("border-danger")
                document.querySelector("#phoneError").style.display = "block"
            }

            if (resp.SEmail == "wrong") {
                // alert("Invalid email address")
                document.querySelector("#email3").classList.add("border-danger")
                document.querySelector("#emailError4").style.display = "block"
            }
            if (resp.SEmail == "exist") {
                // alert("Email address already used")
                document.querySelector("#email3").classList.add("border-danger")
                document.querySelector("#emailError3").style.display = "block"
            }

            if (resp.SPhone == "exist") {
                // alert("Phone number already used")
                document.querySelector("#phone").classList.add("border-danger")
                document.querySelector("#phoneError2").style.display = "block"
            }

            if (resp.SUsername == "exist") {
                // alert("Username already taken")
                document.querySelector("#username3").classList.add("border-danger")
                document.querySelector("#usernameError4").style.display = "block"
            }
        })

}
