import * as functions from './misc'

Object.keys(functions).forEach((key) => {
    console.log("Key", key)
    window[key] = functions[key]
})