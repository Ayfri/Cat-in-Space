export function changeTitle() {
    const name = document.querySelector('.test-name');
    if (document.querySelector('.streamer-list')) {
        const search = document.querySelector('.search-input').getAttribute('value');
        document.title = search ? `${search} - Results` : 'Cat in space';
    } else {
        document.title = name.innerText || 'Cat in space';
    }
}