import { Command, Option } from 'clipanion'
import { emptyDir } from '../utils'

const { HOME } = process.env

export default class Tags extends Command {
  static paths = [['tags']]

  static usage = Command.Usage({
    description: 'Generate tags file',
    details: `
      - \`--dir\` (~/env/note)
      - \`--output\` (~/tags)
      - \`--marker\` (∗)
      - \`--cache\` (~/.cache/gnote)
      
      It will empty the cache directory first.
    `,
    examples: [
      ['Generate tags', 'gnote tags'],
      ['Generate tags and watch', 'gnote tags --watch'],
    ],
  })

  dir = Option.String('--dir', `${HOME}/env/note`, {
    description: 'Input note directory',
  })
  output = Option.String('--output', `${HOME}/tags`, {
    description: 'Ouput tags file',
  })
  marker = Option.String('--marker', '∗', { description: 'Marker string' })
  cache = Option.String('--cache', `${HOME}/.cache/gnote`, {
    description: 'Cache directory',
  })
  watch = Option.Boolean('--watch', {
    description: 'Watch mode',
  })

  async execute() {
    emptyDir(this.cache)
  }
}
