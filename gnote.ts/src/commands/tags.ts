import fs from 'fs/promises'
import glob from 'fast-glob'
import chokidar from 'chokidar'
import { escapeRegExp } from 'lodash-es'
import { Command, Option } from 'clipanion'
import { exit, writeFileWithMkdir, emptyDir } from '#/utils'

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
      - \`--extension\` (.gnote)
      - \`--watch\` (false)
      
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
    await this.createTags()

    if (this.watch) {
      this.startWatch()
    }
  }

  async createTags(): Promise<void> {
    await this.checksDir()

    await this.emptyCacheDir()

    const files = await this.listFiles()
    for (const file of files) {
      await this.createTagsInCache(file)
    }

    await this.createAllTagsFromCache()

    console.log(`Created ${this.output}`)
  }

  async checksDir(): Promise<void> {
    try {
      await fs.access(this.dir)
    } catch (err) {
      if (err.code === 'ENOENT') {
        exit(`Note directory not found: '${this.dir}'`)
      }
      return err
    }
  }

  /**
   * Empty ~/.cache/gnote directory
   */
  async emptyCacheDir(): Promise<void> {
    await fs.mkdir(this.cache, { recursive: true })
    await emptyDir(this.cache)
  }

  /**
   * List **\/*.gnote
   */
  async listFiles(): Promise<string[]> {
    return await glob(`**/*${this.extension}`, { cwd: this.dir })
  }

  async createTagsInCache(file: string): Promise<void> {
    const tags = await this.extractTagsFromFile(file)
    await this.writeTagsToCache(file, tags)
  }

  /**
   * read a.gnote and return tags format
   */
  async extractTagsFromFile(file: string): Promise<string> {
    const fullFile = `${this.dir}/${file}`
    const text = await fs.readFile(fullFile, 'utf8')
    return extractTagsFromText({ text, path: fullFile, marker: this.marker })
  }

  /**
   * Create ~/.cache/gnote/a.gnote
   */
  async writeTagsToCache(file: string, tags: string): Promise<void> {
    if (tags === '') {
      return
    }
    await writeFileWithMkdir(`${this.cache}/${file}`, tags)
  }

  /**
   * create ~/tags
   */
  async createAllTagsFromCache(): Promise<void> {
    const files = await glob(`${this.cache}/**/*${this.extension}`)
    let allTags = ''
    for (const file of files) {
      const tags = await fs.readFile(file, 'utf8')
      allTags += `\n${tags}`
    }
    allTags = sortTags(allTags)
    const result = `!_TAG_FILE_SORTED	1\n${allTags}`
    await writeFileWithMkdir(this.output, result)
  }

  startWatch(): void {
    const watcher = chokidar.watch(`**/*${this.extension}`, {
      cwd: this.dir,
    })
    watcher.on('change', async function (file: string) {
      console.log(`Changed ${file}`)
      await this.createTagsInCache(file)
      await this.createAllTagsFromCache()
    })
    console.log('Watching for changes...')
  }
}

/**
 * Read text and returns tags format
 *
 * Tags Format
 *   hello	/path/hello.gnote /∗hello∗
 */
export function extractTagsFromText({
  text,
  path,
  marker,
}: {
  text: string
  path: string
  marker: string
}): string {
  const pattern = new RegExp(
    `${escapeRegExp(marker)}([^\n ]+)${escapeRegExp(marker)}`,
    'g'
  )
  const matches = text.matchAll(pattern)
  const ids = Array.from(matches).map((v) => v[1])
  const tags = ids
    .map((id) => {
      const jump = `/${marker}${id}${marker}`
      return `${id}\t${path}\t${jump}`
    })
    .join('\n')
  return tags
}

function sortTags(text: string): string {
  return text
    .split(/\n/)
    .sort((a, b) => a.localeCompare(b))
    .join('\n')
}
