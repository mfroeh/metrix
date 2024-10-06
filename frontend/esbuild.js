const esbuild = require("esbuild");

esbuild
    .build({
        entryPoints: ["./Application.tsx"],
        outdir: ".",
        bundle: true,
        minify: false,
        plugins: [],
    })
    .then(() => console.log("⚡ Build complete! ⚡"))
    .catch(() => process.exit(1));