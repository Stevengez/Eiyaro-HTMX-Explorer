export const onLangChange = (lang) => {
    // Weird way of redirect the page
    const urlPaths = window.location.href.split('/')
    urlPaths[3] = lang.toLowerCase()
    window.location = urlPaths.join('/')
}