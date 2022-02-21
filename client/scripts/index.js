let divs = document.querySelectorAll('.StreamerBio');
for(let i=0;i<divs.length;i++) {
    divs[i].innerHTML = divs[i].innerHTML.substring(0,70);
    divs[i].innerHTML += ' ...'
}