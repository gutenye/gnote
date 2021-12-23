import { Cli, Builtins } from 'clipanion'
import pkg from '../package.json'
import help from './commands/help'
import tags from './commands/tags'

const cli = new Cli({
  binaryName: 'gnote',
  binaryLabel: 'Gnote',
  binaryVersion: pkg.version,
})

cli.register(Builtins.HelpCommand)
cli.register(Builtins.VersionCommand)
cli.register(help)
cli.register(tags)

cli.runExit(process.argv.slice(2))
