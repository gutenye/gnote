import esbuild from 'esbuild'
import { nodeExternalsPlugin } from 'esbuild-node-externals'

esbuild
  .build({
    entryPoints: ['src/index.ts'],
    outdir: 'build',
    bundle: true,
    format: 'esm',
    platform: 'node',
    logLevel: 'info',
    target: [`node${process.versions.node}`],
    plugins: [nodeExternalsPlugin()],
  })
  .catch(() => {
    process.exit(1)
  })
