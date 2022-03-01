cutString = () => {
    const divs = document.querySelectorAll('.streamer-bio');
    for (let i = 0; i < divs.length; i++) {
        if (divs[i].textContent.length > 60) {
            divs[i].textContent = divs[i].textContent.substring(0, 60);
            divs[i].textContent += ' ...'
        }
    }
}


hideStreamers = () => {
    let dreamsmp = document.querySelectorAll('.dream-smp');
    let dreamSmp = Array.prototype.slice.call(dreamsmp);
    let hidden = dreamSmp.slice(-3);
    for (let i = 0; i < hidden.length; i++) {
        hidden[i].classList.add('dream-hidden');
    }
}


clickStreamers = () => {
    let dreamsmp = document.querySelectorAll('.dream-smp');
    let dreamSmp = Array.prototype.slice.call(dreamsmp);
    let hidden = dreamSmp.slice(-3);

    let more = document.querySelectorAll('.add-more')[0];

    more.onclick = function() {
        if (more.classList.contains('hidden')) {
            for (let i = 0; i < hidden.length; i++) {
                hidden[i].style.display = "flex";
            }
            more.textContent = "hide...";
            more.classList.remove('hidden');
        } else {
            for (let i = 0; i < hidden.length; i++) {
                hidden[i].style.display = "none";
            }
            more.textContent = "more..."
            more.classList.add('hidden');
        }
    }
}


window.onload = () => {
    cutString()
    hideStreamers()
    clickStreamers()
    changeTitle()
}