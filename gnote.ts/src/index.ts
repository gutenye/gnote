import { Cli, Builtins } from 'clipanion'
import { readPkg } from '#/utils'
import help from './commands/help'
import tags from './commands/tags'

const pkg = readPkg()

const cli = new Cli({
  binaryName: pkg.name,
  binaryLabel: pkg.name,
  binaryVersion: pkg.version,
})

cli.register(Builtins.HelpCommand)
cli.register(Builtins.VersionCommand)
cli.register(help)
cli.register(tags)

void cli.runExit(process.argv.slice(2))
