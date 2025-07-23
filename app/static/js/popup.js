document.addEventListener("DOMContentLoaded", function () {
  document.querySelectorAll(".artist").forEach((img) => {
    const id = img.dataset.id;
    const dialog = document.getElementById("popup-" + id);

    img.addEventListener("click", () => {
      dialog.showModal();
    });

    dialog.querySelector(".closeBtn").addEventListener("click", () => {
      dialog.close();
    });
  });
});

