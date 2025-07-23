
  const artists = document.querySelectorAll(".artist");

  artists.forEach((img) => {
    const id = img.dataset.id;
    const dialog = document.getElementById("popup-" + id);
    const content = dialog.querySelector(".dialog-content");

    img.addEventListener("click", () => {
      dialog.showModal();
    });

    dialog.addEventListener("click", (e) => {
      if (!content.contains(e.target)) {
        dialog.close();
      }
    });
  });














// const dialog = document.querySelector("dialog");
// document.querySelector("#click-here").addEventListener("click", function(){
//     dialog.showModal();
// });

// dialog.querySelector(".closeBtn").addEventListener("click", function(){
//     dialog.close();
// });
