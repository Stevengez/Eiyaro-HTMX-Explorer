import { resolve } from "path"
import { defineConfig } from "vite"

export default defineConfig(({mode}) => {
    const publicDirSrc = process.env.PUBLIC_DIR || ''
    
    return {
        build: {
            lib: {
                entry: [resolve(__dirname, "src/htmx.js")],
                formats: ["es"],
                name: "[name]",
                fileName: "[name]",
            },
            outDir: "static/lib",
            emptyOutDir: false
        },
        server: {
            open: 'templates/index.html',
            port: 3000
        },
        publicDir: publicDirSrc
    }
});