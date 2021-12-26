import { Cli, Builtins } from 'clipanion'
import help from './commands/help'
import tags from './commands/tags'
import pkg from '../package.json'

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
