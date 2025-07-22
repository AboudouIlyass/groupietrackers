const dialog = document.querySelector("dialog");
document.querySelector("#click-here").addEventListener("click", function(){
    dialog.showModal();
});

dialog.querySelector(".closeBtn").addEventListener("click", function(){
    dialog.close();
});
