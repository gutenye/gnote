import { Command, Option } from 'clipanion'
import {
  TagsContext,
  checksDir,
  listFiles,
  emptyCacheDir,
  extractTagsFromFile,
  writeTagsToCache,
  createAllTagsFromCache,
} from '@/tags'

const { HOME } = process.env as { [key: string]: string }

export default class Tags extends Command {
  static paths = [['tags']]

  static usage = Command.Usage({
    description: 'Generate tags file',
    details: `
      - \`--dir\` (~/env/note)
      - \`--output\` (~/tags)
      - \`--marker\` (∗)
      - \`--cache\` (~/.cache/gnote)
      - \`--watch\` (false)
      - \`--extension\` (.gnote)
      
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

  extension = Option.String('--extension', '.gnote', {
    description: 'Note extension',
  })

  watch = Option.Boolean('--watch', false, {
    description: 'Watch mode',
  })

  async execute(): Promise<void> {
    const context: TagsContext = {
      dir: this.dir,
      output: this.output,
      marker: this.marker,
      cache: this.cache,
      watch: this.watch,
      extension: this.extension,
    }

    await checksDir(context)

    await emptyCacheDir(context)

    const files = await listFiles(context)
    for (const file of files) {
      const tags = await extractTagsFromFile(file, context)
      await writeTagsToCache(file, tags, context)
    }

    await createAllTagsFromCache(context)

    console.log(`Created ${context.output}`)
  }
}
