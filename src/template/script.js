window.onload = function() {
    alert("hello wo")
    console.log("hello")

const modal = document.querySelector('.modal');
const closeBtn = document.querySelector('.close');

closeBtn.onclick = function() {
    modal.style.display = "none";
}


window.onclick = function(event) {
    if (event.target == modal) {
        modal.style.display = "none";
    }
}
//colocar em funçao do botão
//colocar com window.onlklde



}