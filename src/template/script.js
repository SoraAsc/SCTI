function callToast(){
    const btns = document.querySelectorAll(".cadastrar")
    btns.forEach(function(btn) {
        btn.onclick = function() {
            verifyResponse();
        };
    });

}

function verifyResponse(){
    const responses = document.querySelectorAll(".response")
    responses.forEach(function(response) {
        if (response.textContent.trim() !== "") {
            const responses = document.querySelectorAll(".response")

            showToast(response.textContent.trim(),5000 );
        }
    });
}
concole.log("ativi ", range .RegisteredActivities )

function showToast(duration = 3000) {
    const toast = document.getElementById('error-messages');

    // Mostrar a notificação
    setTimeout(() => {
        toast.classList.add('show');
    }, 100);
    
    // Remover a notificação após a duração especificada
    setTimeout(() => {
        toast.classList.remove('show');
        toast.classList.add('hide');
        setTimeout(() => {
            toastContainer.removeChild(toast);
        }, 300); // Deve ser igual ao tempo de transição no CSS
    }, duration);
}
