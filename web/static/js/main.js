window.addEventListener("load", (event) => {
  let tasks = document.getElementsByClassName("task");

  for (let i = 0; i < tasks.length; i++) {
    tasks[i].addEventListener("click", async (event) => {
      let task = event.target.parentElement.textContent
      await fetch(`/${task}`, {
        method: "DELETE"
      });

      window.location.reload(true);
    })
  }

})

