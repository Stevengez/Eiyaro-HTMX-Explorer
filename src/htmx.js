import htmx from 'htmx.org'

window.htmx = htmx

document.addEventListener('htmx:beforeRequest', (detail) => {
    console.log("bReq:", detail)
})
document.addEventListener('htmx:afterRequest', (detail) => {
    console.log("aReq:", detail)
})
document.addEventListener('htmx:targetError', (detail) => {
    console.log("badSel:", detail)
})

