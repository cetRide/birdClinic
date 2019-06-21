document.querySelector("#send").onclick = (e) => {
    e.preventDefault()

    document.querySelector("#subject").classList.remove("border-danger")
    document.querySelector("#subjectError").style.display = "none"
    document.querySelector("#message").classList.remove("border-danger")
    document.querySelector("#messageError").style.display = "none"

    axios.post("/send-message", {
        subject: document.querySelector("#subject").value,
        message: document.querySelector("#message").value
    })
        .then(data => {
            window.location.href = data.data.Messge
            alert("Message successful send.Thank you!!")
        })
        .catch(err => {
            var resp = err.response.data
           
            if (resp.Body == "empty") {
                document.querySelector("#message").classList.add("border-danger")
                document.querySelector("#messageError").style.display = "block"
            }

            if (resp.Heading == "empty") {
                document.querySelector("#subject").classList.add("border-danger")
                document.querySelector("#subjectError").style.display = "block"
            }
        })

}