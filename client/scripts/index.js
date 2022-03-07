import {listAppearing} from './animations.js';
import {changeTitle} from './changeTitle.js';

function cutString() {
    const divs = document.querySelectorAll('.streamer-bio');
    for (let i = 0; i < divs.length; i++) {
        if (divs[i].textContent.length > 60) {
            divs[i].textContent = divs[i].textContent.substring(0, 60) + ' ...';
        }
    }
}

function hideStreamers() {
    const dreamSmp = [...document.querySelectorAll('.dream-smp')];
    const hidden = dreamSmp.slice(-4);
    hidden.forEach(item => item.classList.add('dream-hidden'));
}

function clickStreamers() {
    const hidden = document.querySelectorAll('.dream-hidden');

    const more = document.querySelectorAll('.add-more')[0];

    more.onclick = function () {
        if (more.classList.contains('hidden')) {
            hidden.forEach(item => item.style.display = 'flex');
            more.textContent = 'hide...';
            more.classList.remove('hidden');
        } else {
            hidden.forEach(item => item.style.display = 'none');
            more.textContent = 'more...';
            more.classList.add('hidden');
        }
    };
}


window.onload = () => {
    cutString()
    hideStreamers()
    clickStreamers()
    changeTitle()
    listAppearing()
}