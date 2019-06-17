
   document.querySelector("#farmerlogin").onclick = (e) => {
    e.preventDefault()

    document.querySelector("#usernameA").classList.remove("border-danger")
    document.querySelector("#usernameErrorA").style.display = "none"
    document.querySelector("#passwordA").classList.remove("border-danger")
    document.querySelector("#passwordErrorA").style.display = "none"

    axios.post("/f-login", {
      username: document.querySelector("#usernameA").value,
       password: document.querySelector("#passwordA").value
    })
       .then(data => {
          window.location.href = data.data.Messages
          alert("Successful log in")
       })
       .catch(err => {
          var respo = err.response.data

          if (respo.FarmerUsername == "invalid") {
             alert("Username does not exist")
             document.querySelector("#usernameA").classList.add("border-danger")
             document.querySelector("#usernameErrorA").style.display = "block"
          }

          if (respo.FarmerPassword == "invalid") {
             alert("Incorrect password")
             document.querySelector("#passwordA").classList.add("border-danger")
             document.querySelector("#passwordErrorA").style.display = "block"
          }
       })

 }
