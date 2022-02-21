let divs = document.querySelectorAll('.StreamerBio');
for (let i = 0; i < divs.length; i++) {
    if (divs[i].textContent.length > 60) {
        divs[i].textContent = divs[i].textContent.substring(0, 60);
        divs[i].textContent += ' ...'
    }
}