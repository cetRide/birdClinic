document.querySelector("#btnContactUs").onclick = (e) => {
    e.preventDefault()

    document.querySelector("#firstname").classList.remove("border-danger")
    document.querySelector("#firstnameError").style.display = "none"
    document.querySelector("#lastname").classList.remove("border-danger")
    document.querySelector("#lastnameError").style.display = "none"
    document.querySelector("#email").classList.remove("border-danger")
    document.querySelector("#emailError").style.display = "none"
    document.querySelector("#emailError2").style.display = "none"
    document.querySelector("#phone").classList.remove("border-danger")
    document.querySelector("#phoneError").style.display = "none"
    document.querySelector("#phoneError2").style.display = "none"
    document.querySelector("#subject").classList.remove("border-danger")
    document.querySelector("#subjectError").style.display = "none"
    document.querySelector("#message").classList.remove("border-danger")
    document.querySelector("#messageError").style.display = "none"

    axios.post("/contact_us", {
        f_name: document.querySelector("#firstname").value,
        l_name: document.querySelector("#lastname").value,
        email: document.querySelector("#email").value,
        phone: document.querySelector("#phone").value,
        subject: document.querySelector("#subject").value,
        message: document.querySelector("#message").value
    })
        .then(data => {
            window.location.href = data.data.ContactUs
            alert("Message successful send")
        })
        .catch(err => {
            var resp = err.response.data

            if (resp.ContactUsEmail == "empty") {
                document.querySelector("#email").classList.add("border-danger")
                document.querySelector("#emailError").style.display = "block"
            }

            if (resp.ContactUsEmail == "invalid") {
                document.querySelector("#email").classList.add("border-danger")
                document.querySelector("#emailError2").style.display = "block"
            }

            if (resp.ContactUsPhone == "empty") {
                document.querySelector("#phone").classList.add("border-danger")
                document.querySelector("#phoneError").style.display = "block"
            }

            if (resp.ContactUsPhone == "invalid") {
                document.querySelector("#phone").classList.add("border-danger")
                document.querySelector("#phoneError2").style.display = "block"
            }

            if (resp.ContactUsFName == "empty") {
                document.querySelector("#firstname").classList.add("border-danger")
                document.querySelector("#firstnameError").style.display = "block"
            }
            if (resp.ContactUsLName == "empty") {
                document.querySelector("#lastname").classList.add("border-danger")
                document.querySelector("#lastnameError").style.display = "block"
            }
            if (resp.ContactUsMessage == "empty") {
                document.querySelector("#message").classList.add("border-danger")
                document.querySelector("#messageError").style.display = "block"
            }

            if (resp.ContactUsSubject == "empty") {
                document.querySelector("#subject").classList.add("border-danger")
                document.querySelector("#subjectError").style.display = "block"
            }
        })

}