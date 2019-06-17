
document.querySelector("#login2").onclick = (e) => {
   e.preventDefault()

   document.querySelector("#empidu").classList.remove("border-danger")
   document.querySelector("#empidErroru").style.display = "none"
   document.querySelector("#passu").classList.remove("border-danger")
   document.querySelector("#passErroru").style.display = "none"

   axios.post("/s-login", {
      user: document.querySelector("#empidu").value,
      pass: document.querySelector("#passu").value
   })
      .then(data => {
         window.location.href = data.data.SpecialistResponse
         alert("Successful log in")
      })
      .catch(err => {
         var resp = err.response.data
         if (resp.SpecialistUsername == "incorrect") {
            alert("Username does not exist")
            document.querySelector("#empidu").classList.add("border-danger")
            document.querySelector("#empidErroru").style.display = "block"
         }

         if (resp.SpecialistPassword == "incorrect") {
            alert("Incorrect password")
            document.querySelector("#passu").classList.add("border-danger")
            document.querySelector("#passErroru").style.display = "block"
         }
      })

}
