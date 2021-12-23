const esbuild = require('esbuild')

esbuild
  .build({
    entryPoints: ['src/index.ts'],
    outdir: 'build',
    platform: 'node',
    logLevel: 'info',
    target: [`node${process.versions.node}`],
    bundle: true,
    format: 'cjs',
  })
  .catch(() => {
    process.exit(1)
  })
