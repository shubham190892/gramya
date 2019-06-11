/* When the user clicks on the button,
toggle between hiding and showing the dropdown content */
function myFunction(e) {
    let d = $("#explore");
    d.css({'left':e.pageX, 'top':e.pageY});
    d.toggleClass("show");
}

// Close the dropdown if the user clicks outside of it
window.onclick = function(event) {
    if (event.target.matches('.person')) {
        myFunction(event)
    }else{
        let explore = document.getElementById("explore")
        if (explore.classList.contains('show')) {
            explore.classList.remove('show');
        }
    }
}