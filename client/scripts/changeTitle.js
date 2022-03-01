export function changeTitle() {
    const name = document.querySelectorAll('.test-name');
    if (document.querySelector('.streamer-list')) {
        document.title = `${document.querySelector('.search-input').getAttribute('value') || 'Cat in space'} - Results`;
    } else {
        document.title = name[0].innerText;
    }
}