import { resolve } from "path"
import { defineConfig } from "vite"

export default defineConfig(({mode}) => {
    const publicDirSrc = process.env.PUBLIC_DIR || ''
    
    return {
        build: {
            lib: {
                entry: {
                    htmx: resolve(__dirname, "src/htmx.js"),
                    global: resolve(__dirname, "src/global.js")
                },
                formats: ["es"],
                name: "[name]",
                fileName: "[name]",
            },
            outDir: "fs/static/lib",
            emptyOutDir: false
        },
        server: {
            open: 'fs/templates/layout/global.html',
            port: 3000
        },
        publicDir: publicDirSrc
    }
});